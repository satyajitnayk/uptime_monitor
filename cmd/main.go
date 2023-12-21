package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/satyajitnayk/uptime_monitor/handlers"
	health "github.com/satyajitnayk/uptime_monitor/handlers"
	"github.com/satyajitnayk/uptime_monitor/internal/database"
	"github.com/satyajitnayk/uptime_monitor/internal/models"
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

	fmt.Println("Connected to the PostgreSQL database")

	// Create tables and migrate schemas if necessary
	if err := db.AutoMigrate(&models.User{}).Error; err != nil {
		log.Fatal(err)
	}

	// Create a new router
	r := mux.NewRouter()

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		health.HealthHandler(w, r, db, startTime)
	}).Methods("GET")

	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterUserHandler(w, r, db)
	}).Methods("POST")

	port := GetPort()
	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(addr, r); nil != err {
		fmt.Printf("Error starting the server: %v\n", err)
	}
}
