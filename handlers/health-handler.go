package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
)

type HealthResponse struct {
	Status string `json:"server_status"`
	Uptime string `json:"uptime"`
	DB     string `json:"db_status"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB, startTime time.Time) {
	// Ping the database to check its status
	dbStatus := "ok"
	if err := db.DB().Ping(); err != nil {
		dbStatus = "error"
	}

	response := HealthResponse{
		Status: "ok",
		Uptime: time.Since(startTime).String(),
		DB:     dbStatus,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
