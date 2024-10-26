package models

type AnotherConfig struct{}

func (a *AnotherConfig) GetMetrics() map[string]string {
	return map[string]string{
		"average_value": "AVG(value)",
		"max_value":     "MAX(value)",
	}
}

func (a *AnotherConfig) GetCustomDimensions() map[string]string {
	return map[string]string{
		"region": "CASE when region = 'NA' then 'North America' when region = 'EU' then 'Europe' END",
	}
}

func (a *AnotherConfig) GetTableName() string {
	return "another_example_table"
}

func (a *AnotherConfig) GetBreakdownByDate(breakdown string) string {
	switch breakdown {
	case "hourly":
		return "HOUR(date)"
	case "daily":
		return "DATE(date)"
	case "weekly":
		return "WEEK(date)"
	default:
		return ""
	}
}