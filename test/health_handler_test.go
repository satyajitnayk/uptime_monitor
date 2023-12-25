package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/satyajitnayk/uptime_monitor/handlers"
)

func createMockDB() *gorm.DB {
	// Connect to an in-memory SQLite database
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		fmt.Println("Failed to connect to SQLite database:", err)
		panic(err)
	}
	return db
}

func TestHealthHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := createMockDB()
		startTime := time.Now()
		handlers.HealthHandler(w, r, db, startTime)
	}))
	if server != nil {
		defer server.Close()
	}
	response, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", response.StatusCode)
	}
}
