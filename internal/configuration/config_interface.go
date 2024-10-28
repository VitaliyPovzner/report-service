package configuration

type ReportConfig interface {
	GetMetrics() (map[string]string, error)
	GetCustomDimensions() (map[string]string, error)
	GetBreakdownByDate(breakdown string) (string, error)
	GetTableName() string
}
