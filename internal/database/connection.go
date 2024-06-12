package database

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/nilemarezz/go-init-template-batch/internal/config"
	"github.com/nilemarezz/go-init-template-batch/internal/logger"
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
			logger.Logger.Info(fmt.Sprintf("Attempt %d: Failed to connect to database: %v", attempt, err))
			time.Sleep(retryInterval)
			retryInterval *= 2 // exponential backoff
			continue
		}
		if err := db.Ping(); err != nil {
			logger.Logger.Info(fmt.Sprintf("Attempt %d: Failed to connect to database: %v", attempt, err))
			db.Close()
			time.Sleep(retryInterval)
			retryInterval *= 2 // exponential backoff
			continue
		}
		logger.Logger.Info("Connect database success")
		return db, nil
	}
	return nil, err
}
