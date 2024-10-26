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
        Breakdown:  "daily",
    }

    result, err := GenerateReport(params)

    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }

    if result == nil {
        t.Fatal("expected non-nil result")
    }
}
