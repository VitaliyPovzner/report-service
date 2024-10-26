package models

type TestMetricsConfig struct{}

func (t TestMetricsConfig) GetTableName() string {
    return "test_table"
}

func (t TestMetricsConfig) GetMetrics() map[string]string {
    return map[string]string{"metric1": "SUM(metric1)"}
}

func (t TestMetricsConfig) GetCustomDimensions() map[string]string {
    return map[string]string{"dimension100": "dimension100"}
}

func (t TestMetricsConfig) GetBreakdownByDate(breakdown string) string {
	switch breakdown {
	case "hourly":
		return "date"
	case "daily":
		return "DATE(date)"
	case "monthly":
		return "DATE_FORMAT(date, '%Y-%m')"
	default:
		return ""
	}
}