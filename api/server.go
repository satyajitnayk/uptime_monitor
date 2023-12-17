package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var startTime time.Time

func main() {
	startTime = time.Now()

	// Load the .env file in the current directory
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/health", healthHandler)

	port := os.Getenv("PORT")
	fmt.Printf("Server is running on port %s\n", port)
	addr := ":" + port
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startTime).Seconds()

	// struct to represent the JSON response
	response := struct {
		Status      string  `json:"status"`
		UptimeInSec float64 `json:"uptimeInSec"`
	}{
		Status:      "ok",
		UptimeInSec: uptime,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// set content type to json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the response writer
	w.Write(jsonResponse)
}
