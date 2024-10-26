package service

import (
    "fmt"
    "report-service/internal/models"
    "strings"
)

func GenerateReport(params models.AggregationRequest, config models.MetricsConfig) ([]map[string]interface{}, error) {
    sql, err := generateSQL(params, config)
    if err != nil {
        return nil, err
    }
    fmt.Print(sql)
    
    // Here we return dummy data for demonstration
    return []map[string]interface{}{
        {"dimension1": sql, "metric1": 100},
    }, nil
}


func generateSQL(params models.AggregationRequest, config models.MetricsConfig) (string, error) {
    tableName:= config.GetTableName()
    selectClause := generateSelectClause(params, config)
    groupByClause := generateGroupByClause(params, config)
    havingClause := generateHavingClause(params, config) 

    // Build the full SQL query
    query := fmt.Sprintf(
        "SELECT %s FROM %s WHERE datetime BETWEEN '%s' AND '%s' %s %s",
        selectClause,
        tableName, // Use the tableName variable here
        params.DateFrom.Format("2006-01-02"),
        params.DateTo.Format("2006-01-02"),
        groupByClause,
        havingClause,
    )

    return query, nil
}

// generateSelectClause constructs the SELECT part of the SQL query.
func generateSelectClause(params models.AggregationRequest, config models.MetricsConfig) string {
    selectClauses := []string{}

    // Add metrics
    for _, metric := range params.Metrics {
        if sql, exists := config.GetMetrics()[metric]; exists {
            selectClauses = append(selectClauses, sql+" AS "+metric)
        }
    }

    // Add dimensions
    for _, dimension := range params.Dimensions {
        if sql, exists := config.GetCustomDimensions()[dimension]; exists {
            selectClauses = append(selectClauses, sql+" AS "+dimension)
        } else {
            selectClauses = append(selectClauses, dimension)
        }
    }

    return strings.Join(selectClauses, ", ")
}

// generateGroupByClause constructs the GROUP BY part of the SQL query.
func generateGroupByClause(params models.AggregationRequest, config models.MetricsConfig) string {
    groupByClauses := []string{}


    if breakdownSQL := config.GetBreakdownByDate(params.Breakdown); breakdownSQL != "" {
        groupByClauses = append(groupByClauses, breakdownSQL)
    }

    return "GROUP BY " + strings.Join(groupByClauses, ", ")
}


func generateHavingClause(params models.AggregationRequest, config models.MetricsConfig) string {


    return "" 
}
