package database

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/nilemarezz/go-init-template/pkg/config"
)

func ConnectDB(config *config.Config) (*sqlx.DB, error) {
	// Database connection parameters
	dbURI := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.DBName, config.Database.Password)

	var db *sqlx.DB
	var err error

	// Retry connecting to database with exponential backoff
	retryInterval := 1 * time.Second
	maxRetryAttempts := 5
	for attempt := 1; attempt <= maxRetryAttempts; attempt++ {
		db, err = sqlx.Connect("postgres", dbURI)
		if err != nil {
			log.Printf("Attempt %d: Failed to connect to database: %v", attempt, err)
			time.Sleep(retryInterval)
			retryInterval *= 2 // exponential backoff
			continue
		}
		if err := db.Ping(); err != nil {
			log.Printf("Attempt %d: Failed to ping database: %v", attempt, err)
			db.Close()
			time.Sleep(retryInterval)
			retryInterval *= 2 // exponential backoff
			continue
		}
		log.Printf("Connected to database")
		return db, nil
	}
	return nil, err
}
