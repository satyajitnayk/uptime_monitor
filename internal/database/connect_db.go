package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectDB() (*gorm.DB, error) {
	// Get database connection parameters from environment variables
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	// Construct the connection string
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		host, port, user, dbname, password)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		return nil, err
	}

	// Enable Logger, show detailed log
	db.LogMode(true)

	return db, err
}
