package configuration

type BaseConfig struct{}

func (b *BaseConfig) GetBreakdownByDate(breakdown string) (string, error) {
	switch breakdown {
	case "hourly":
		return "datetime", nil
	case "daily":
		return "DATE(datetime)", nil
	case "monthly":
		return "DATE_FORMAT(datetime, '%Y-%m')", nil
	default:
		return "", nil
	}
}

func (b BaseConfig) GetTableName() string {
    return "base_table"
}

func (b BaseConfig) GetMetrics() (map[string]string, error) {
    return map[string]string{"metric1": "SUM(metric1)"},nil
}

func (b BaseConfig) GetCustomDimensions() (map[string]string,error) {
    return map[string]string{"dimension100": "dimension100"},nil
}


