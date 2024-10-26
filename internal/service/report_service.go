package service

import (
    "report-service/internal/models"
)


func GenerateReport(params models.AggregationRequest) ([]map[string]interface{}, error) {
	// Returning dummy data for now
    return []map[string]interface{}{
        {"dimension1": "value1", "metric1": 100},
    }, nil
}
