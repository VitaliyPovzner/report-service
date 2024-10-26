package models


type MetricsConfig interface {
	GetMetrics() map[string]string
	GetCustomDimensions() map[string]string
	GetBreakdownByDate(breakdown string) string
	GetTableName() string
}

type GenericConfig struct{}


func (g *GenericConfig) GetMetrics() map[string]string {
	return map[string]string{
		"total_value": "SUM(value)",
		"total_count": "COUNT(*)",
	}
}


func (g *GenericConfig) GetCustomDimensions() map[string]string {
	return map[string]string{
		"category": "CASE when category like '%-mob-%' then 'Mobile' when category like '%-d-%' then 'Desktop' END",
	}
}
func (g * GenericConfig) GetTableName() string{
	return "my_example_table"
}
func (g *GenericConfig) GetBreakdownByDate(breakdown string) string{
	switch breakdown {
	case "hourly":
		return "est_datetime"
	case "daily":
		return "DATE(est_datetime)"
	case "monthly":
		return "DATE_FORMAT(est_datetime, '%Y-%m')"
	default:
		return ""
	}
}
