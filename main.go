package main

import (
	"flag"
	"fmt"

	"github.com/nilemarezz/go-init-template-batch/internal/config"
	"github.com/nilemarezz/go-init-template-batch/internal/database"
	"github.com/nilemarezz/go-init-template-batch/internal/logger"
	"go.uber.org/zap"
)

func main() {
	appName := "batch1"
	var env string
	var dateArg string
	flag.StringVar(&env, "env", "", "Environment (dev, staging, prod)")
	flag.StringVar(&dateArg, "date", "", "Date arguments (yyyyMMdd)")
	flag.Parse()

	// Load config from config file
	config, err := config.LoadConfig(env)
	if err != nil {
		panic(err)
	}

	logger.InitLogger(appName)
	logger.Logger.Info(fmt.Sprintf("Start application with args... %v", dateArg))

	if env == "" {
		env = "default"
	}

	logger.Logger.Info(fmt.Sprintf("Config set to '%v' env", env), zap.Reflect("config", config))

	// Initialize database
	db, err := database.ConnectDB(&config)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("Fail to connect database at '%v' ", config.Database.Host))
		panic("Database connection fail")
	}
	defer db.Close()

}
