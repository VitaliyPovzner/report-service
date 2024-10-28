package configuration

import "fmt"

type AnotherConfig struct{}

func (a *AnotherConfig) GetMetrics() (map[string]string, error) {
	return map[string]string{
		"average_value": "AVG(value)",
		"max_value":     "MAX(value)",
	}, nil
}

func (a *AnotherConfig) GetCustomDimensions() (map[string]string, error) {
	return map[string]string{
		"region": "CASE when region = 'NA' then 'North America' when region = 'EU' then 'Europe' END",
	}, nil
}

func (a *AnotherConfig) GetTableName() string {
	return "another_example_table"
}

func (a *AnotherConfig) GetBreakdownByDate(breakdown string) (string, error) {
	switch breakdown {
	case "hourly":
		return "HOUR(date)", nil
	case "daily":
		return "DATE(date)", nil
	case "weekly":
		return "WEEK(date)", nil
	default:
		return "", fmt.Errorf("unknown breakdown type: %s", breakdown)
	}
}