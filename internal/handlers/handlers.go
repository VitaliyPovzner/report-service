package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"report-service/internal/configuration"
	"report-service/internal/models"
	"report-service/internal/service"
)

func ReportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.AggregationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("Decode error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Request:", req)

	vars := mux.Vars(r)
	reportType := vars["reportType"]

	config, err := configuration.NewReportConfig(reportType)
	if err != nil {
		http.Error(w, "Invalid report type: "+err.Error(), http.StatusBadRequest)
		return
	}
	result, err := service.GenerateReport(req, config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
