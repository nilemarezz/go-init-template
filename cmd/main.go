package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nilemarezz/go-init-template/pkg/config"
	"github.com/nilemarezz/go-init-template/pkg/database"
	"github.com/nilemarezz/go-init-template/pkg/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/nilemarezz/go-init-template/internal/author"
	ginSwagger "github.com/swaggo/gin-swagger"

	// gin-swagger middleware
	swaggerFiles "github.com/swaggo/files"

	_ "github.com/nilemarezz/go-init-template/docs"
)

// @title           Golang Testing Project
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	var env string
	flag.StringVar(&env, "env", "dev", "Environment (dev, staging, prod)")
	flag.Parse()

	// Load config from config file
	config, err := config.LoadConfig(env)
	if err != nil {
		panic(err)
	}

	// Initialize logger
	if err := logger.InitLogger(&config); err != nil {
		panic(err)
	}

	// Initialize database
	db, err := database.ConnectDB(&config)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Initialize  /metrics routes for prometheus metrics
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Init routes
	author.SetupRouter(router, db)

	// Initialize web service
	s := fmt.Sprintf(":%s", config.App.Port)
	logger.Info(s)
	router.Run(s)

}
