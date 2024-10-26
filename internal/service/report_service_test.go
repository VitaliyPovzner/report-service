package service

import (
    "testing"
    "time"
)

func TestGenerateReport(t *testing.T) {
    params := QueryParams{
        Dimensions: []string{"dimension1"},
        Metrics:    []string{"metric1"},
        DateFrom:   time.Now().AddDate(0, 0, -1), // 1 day ago
        DateTo:     time.Now(),
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
