package main

import (
	"log"
	"net/http"
	"report-service/internal/database"
	"report-service/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	if err := database.ConnectDB("localhost", "5432", "user", "password", "db"); err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/report/{reportType}", handlers.ReportHandler)

	log.Println("Starting server on :3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}