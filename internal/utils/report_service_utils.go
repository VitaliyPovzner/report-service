package utils

import (
	"strconv"
	"strings"
	"fmt"
)

// BuildFilterCondition constructs a SQL filter condition based on the operand, operator, and value.
// It handles both single values and arrays, delegating to specific functions as needed.
func BuildFilterCondition(operand, operator, value string, operatorMap map[string]string) string {
	if isArray(value) && operator != "in" && operator != "not_in" {
		return buildFilterConditionForArray(operand, operator, value, operatorMap)
	}
	return build(operand, operator, value, operatorMap)
}

// isNumeric checks if a given string can be parsed as a numeric value.
func isNumeric(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}

// isArray determines if a string represents an array by checking for commas.
func isArray(value string) bool {
	return strings.Contains(value, ",")
}

// formatValues formats a comma-separated string into a SQL-friendly list of quoted values.
func formatValues(value string) string {
	values := strings.Split(value, ",")
	for i, v := range values {
		values[i] = fmt.Sprintf("'%s'", strings.TrimSpace(v))
	}
	return strings.Join(values, ", ")
}

// buildFilterConditionForArray constructs a SQL condition for array values, using OR logic.
func buildFilterConditionForArray(operand, operator, value string, operatorMap map[string]string) string {
	values := strings.Split(value, ",")
	mappedVals := []string{}
	for _, v := range values {
		mappedVals = append(mappedVals, build(operand, operator, v, operatorMap))
	}
	return fmt.Sprintf("(%s)", strings.Join(mappedVals, " OR "))
}

// build constructs a SQL condition for a single value, using the appropriate operator.
func build(operand, operator, value string, operatorMap map[string]string) string {
	var condition string

	switch operator {
	case "cont", "not_cont":
		condition = fmt.Sprintf("%s %s '%%%s%%'", operand, operatorMap[operator], value)
	case "starts":
		condition = fmt.Sprintf("%s %s '%s%%'", operand, operatorMap[operator], value)
	case "in", "not_in":
		formattedValues := formatValues(value)
		condition = fmt.Sprintf("%s %s (%s)", operand, operatorMap[operator], formattedValues)
	default:
		if isNumeric(value) {
			condition = fmt.Sprintf("%s %s %s", operand, operatorMap[operator], value)
		} else {
			condition = fmt.Sprintf("%s %s '%s'", operand, operatorMap[operator], value)
		}
	}

	return condition
}