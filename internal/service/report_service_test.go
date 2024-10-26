package service

import (
    "testing"
    "time"
    "report-service/internal/models"
)

func TestGenerateReport(t *testing.T) {
    params := models.AggregationRequest{
        Dimensions: []string{"dimension1"},
        Metrics:    []string{"metric1"},
        DateFrom:   models.FlexibleDateTime(time.Now().AddDate(0, 0, -1)),
        DateTo:     models.FlexibleDateTime(time.Now()),
        Filters:    []models.Filter{{Operand: "metric1", Operator: "gr_eq", Value: "100"}},
    }

    config := models.TestMetricsConfig{}

    report, err := GenerateReport(params, config)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    if len(report) != 1 {
        t.Fatalf("Expected report length of 1, got %d", len(report))
    }

    if report[0]["metric1"] != 100 {
        t.Errorf("Expected metric1 to be 100, got %v", report[0]["metric1"])
    }
}

func TestGenerateSQL(t *testing.T) {
    params := models.AggregationRequest{
        Dimensions: []string{"dimension1"},
        Metrics:    []string{"metric1"},
        DateFrom:   models.FlexibleDateTime(time.Now().AddDate(0, 0, -1)),
        DateTo:     models.FlexibleDateTime(time.Now()),
        Filters:    []models.Filter{{Operand: "metric1", Operator: "gr_eq", Value: "100"}},
		Breakdown: "daily",
    }

    config := models.TestMetricsConfig{}

    sql, err := generateSQL(params, config)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    expectedSQL := "SELECT SUM(metric1) AS metric1, dimension1 AS dimension1 FROM test_table WHERE datetime BETWEEN '" +
        params.DateFrom.Format("2006-01-02") + "' AND '" + params.DateTo.Format("2006-01-02") + "' GROUP BY DATE(datetime) HAVING metric1 >= 100"
    if sql != expectedSQL {
        t.Errorf("Expected SQL: %s, got: %s", expectedSQL, sql)
    }
}

func TestGenerateSelectClause(t *testing.T) {
    params := models.AggregationRequest{
        Dimensions: []string{"dimension1"},
        Metrics:    []string{"metric1"},
    }

    config := models.TestMetricsConfig{}

    selectClause := generateSelectClause(params, config)

    expectedClause := "SUM(metric1) AS metric1, dimension1 AS dimension1"
    if selectClause != expectedClause {
        t.Errorf("Expected SELECT clause: %s, got: %s", expectedClause, selectClause)
    }
}

func TestGenerateGroupByClause(t *testing.T) {
    params := models.AggregationRequest{
        Breakdown: "daily",
    }

    config := models.TestMetricsConfig{}

    groupByClause := generateGroupByClause(params, config)

    expectedClause := "GROUP BY DATE(datetime)"
    if groupByClause != expectedClause {
        t.Errorf("Expected GROUP BY clause: %s, got: %s", expectedClause, groupByClause)
    }
}

func TestGenerateHavingClause(t *testing.T) {
    params := models.AggregationRequest{
        Filters: []models.Filter{{Operand: "metric1", Operator: "gr_eq", Value: "100"}},
    }

    havingClause := generateHavingClause(params)

    expectedClause := "HAVING metric1 >= 100"
    if havingClause != expectedClause {
        t.Errorf("Expected HAVING clause: %s, got: %s", expectedClause, havingClause)
    }
}
func TestGenerateSQLWithComplexParams(t *testing.T) {
    params := models.AggregationRequest{
        Dimensions: []string{"dimension1", "dimension2"},
        Metrics:    []string{"metric1", "metric2"},
        DateFrom:   models.FlexibleDateTime(time.Now().AddDate(0, 0, -7)),
        DateTo:     models.FlexibleDateTime(time.Now()),
        Filters: []models.Filter{
            {Operand: "metric1", Operator: "gr_eq", Value: "100"},
            {Operand: "metric2", Operator: "less", Value: "500"},
        },
        Breakdown: "monthly",
    }

    config := models.TestMetricsConfig{}

    sql, err := generateSQL(params, config)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    expectedSQL := "SELECT SUM(metric1) AS metric1, SUM(metric2) AS metric2, dimension1 AS dimension1, dimension2 AS dimension2 FROM test_table WHERE datetime BETWEEN '" +
        params.DateFrom.Format("2006-01-02") + "' AND '" + params.DateTo.Format("2006-01-02") + "' GROUP BY DATE_FORMAT(date, '%Y-%m') HAVING metric1 >= 100 AND metric2 < 500"

    if sql != expectedSQL {
        t.Errorf("Expected SQL: %s, got: %s", expectedSQL, sql)
    }
}