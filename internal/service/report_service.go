package service

import (
    "fmt"
    "report-service/internal/models"
    "report-service/internal/configuration"
    "strings"
    "report-service/internal/utils"
)

func GenerateReport(params models.AggregationRequest, config configuration.ReportConfig) ([]map[string]interface{}, error) {
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


func generateSQL(params models.AggregationRequest, config configuration.ReportConfig) (string, error) {
    tableName:= config.GetTableName()
    selectClause := generateSelectClause(params, config)
    groupByClause := generateGroupByClause(params, config)
    havingClause := generateHavingClause(params) 

    // Build the full SQL query
    query := fmt.Sprintf(
        "SELECT %s FROM %s WHERE datetime BETWEEN '%s' AND '%s' %s %s",
        selectClause,
        tableName, 
        params.DateFrom.Format("2006-01-02"),
        params.DateTo.Format("2006-01-02"),
        groupByClause,
        havingClause,
    )

    return query, nil
}



func generateGroupByClause(params models.AggregationRequest, config configuration.ReportConfig) string {
    groupByClauses := []string{}

    breakdown,err:=config.GetBreakdownByDate(params.Breakdown)
    if err!=nil {
        return ""
    }

    if breakdownSQL := breakdown; breakdownSQL != "" {
        groupByClauses = append(groupByClauses, breakdownSQL)
    }

    return "GROUP BY " + strings.Join(groupByClauses, ", ")
}


func generateHavingClause(params models.AggregationRequest) string {
    havingClauses := []string{}

    operatorMap := map[string]string{
        "eq":       "=",
        "not_eq":   "<>",
        "less_eq":  "<=",
        "less":     "<",
        "gr_eq":    ">=",
        "gr":       ">",
        "cont":     "LIKE",
        "not_cont": "NOT LIKE",
        "starts":   "LIKE",
        "in":       "IN",
        "not_in":   "NOT IN",
    }

    for _, filter := range params.Filters {
        condition := utils.BuildFilterCondition(filter.Operand, filter.Operator, filter.Value, operatorMap)
        havingClauses = append(havingClauses, condition)
    }

    if len(havingClauses) > 0 {
        return "HAVING " + strings.Join(havingClauses, " AND ")
    }

    return ""
}


func generateSelectClause(params models.AggregationRequest, config configuration.ReportConfig) string {
    selectClauses := []string{}

    // Add metrics
    metricsMap, _ := config.GetMetrics()
   
    for _, metric := range params.Metrics {
        if sql, exists := metricsMap[metric]; exists {
            selectClauses = append(selectClauses, sql+" AS "+metric)
        }
    }

    // Add dimensions
    customDimensions, _ := config.GetCustomDimensions()
 
    for _, dimension := range params.Dimensions {
        if sql, exists := customDimensions[dimension]; exists {
            selectClauses = append(selectClauses, sql+" AS "+dimension)
        } else {
            selectClauses = append(selectClauses, dimension)
        }
    }
    

    return strings.Join(selectClauses, ", ")
}