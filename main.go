// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var startTime time.Time

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "3001"
	}
	return port
}

func TestDBConnection(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return fmt.Errorf("failed to test the database connection: %v", err)
	}
	return nil
}

func ConnectDB() (*sql.DB, error) {
	// Get database connection parameters from environment variables
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	// Construct the connection string
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, dbname, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	startTime = time.Now()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the PostgreSQL database
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection
	if err := TestDBConnection(db); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the PostgreSQL database")

	http.HandleFunc("/health", healthHandler)

	port := GetPort()
	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server is running on port %s\n", port)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
