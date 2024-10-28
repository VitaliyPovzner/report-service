package configuration

import "fmt"

type TestConfig struct{}

func (t TestConfig) GetTableName() string {
	return "test_table"
}

func (t TestConfig) GetMetrics() (map[string]string, error) {
	return map[string]string{"metric1": "SUM(metric1)"}, nil
}

func (t TestConfig) GetCustomDimensions() (map[string]string, error) {
	return map[string]string{"dimension100": "dimension100"}, nil
}

func (t TestConfig) GetBreakdownByDate(breakdown string) (string, error) {
	switch breakdown {
	case "hourly":
		return "date", nil
	case "daily":
		return "DATE(date)", nil
	case "monthly":
		return "DATE_FORMAT(date, '%Y-%m')", nil
	default:
		return "", fmt.Errorf("unknown breakdown type: %s", breakdown)
	}
}
