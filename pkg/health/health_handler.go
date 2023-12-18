package health

import (
	"encoding/json"
	"net/http"
	"time"
)

type HealthResponse struct {
	Status string `json:"status"`
	Uptime string `json:"uptime"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request, startTime time.Time) {
	response := HealthResponse{
		Status: "ok",
		Uptime: time.Since(startTime).String(),
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
