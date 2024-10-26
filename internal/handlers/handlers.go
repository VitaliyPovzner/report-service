package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"report-service/internal/models"
	"report-service/internal/service"

	"github.com/gorilla/mux"
)

func ReportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.AggregationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("Decode error: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Request: ", req)

	vars := mux.Vars(r)
	reportType := vars["reportType"]

	var config models.MetricsConfig
	switch reportType {
	case "example-report":
		config = &models.GenericConfig{}
	case "another-report":
		config = &models.AnotherConfig{}
	default:
		http.Error(w, "Invalid report type", http.StatusBadRequest)
		return
	}

	result, err := service.GenerateReport(req, config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}