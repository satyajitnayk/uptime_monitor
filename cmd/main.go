package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/satyajitnayk/uptime_monitor/internal/database"
	"github.com/satyajitnayk/uptime_monitor/pkg/health"
)

var startTime time.Time

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	return port
}

func main() {
	startTime = time.Now()
	godotenv.Load()

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		health.HealthHandler(w, r, db, startTime)
	})

	port := GetPort()
	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(addr, nil); nil != err {
		fmt.Printf("Error starting the server: %v\n", err)
	}
}
