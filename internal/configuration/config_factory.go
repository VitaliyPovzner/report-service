package configuration

import "fmt"

func NewReportConfig(configType string) (ReportConfig, error) {
    switch configType {
    case "example-report":
        return &BaseConfig{}, nil
    case "another-report":
        return &AnotherConfig{}, nil
    case "test-report":
        return &TestConfig{}, nil
    default:
        return nil, fmt.Errorf("unknown configuration type: %s", configType)
    }
}
