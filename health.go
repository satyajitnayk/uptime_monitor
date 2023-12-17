// health.go
package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// HealthResponse represents the JSON response for the health endpoint
type HealthResponse struct {
	Status      string  `json:"status"`
	UptimeInSec float64 `json:"uptimeInSec"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startTime).Seconds()

	response := HealthResponse{
		Status:      "ok",
		UptimeInSec: uptime,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	SetJSONContentType(w)
	w.Write(jsonResponse)
}

// SetJSONContentType sets the content type header to JSON
func SetJSONContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
