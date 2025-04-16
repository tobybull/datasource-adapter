package main

import (
	"datasource-adapter/internal/api"
	"datasource-adapter/internal/handlers"
	"datasource-adapter/internal/logger"
	"datasource-adapter/internal/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"

	"datasource-adapter/internal/config"
)

func main() {

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Configuration error:", err)
		os.Exit(1)
	}
	fmt.Printf("The config --> %+v\n", cfg)

	// Set up logger
	log := logger.SetupLogger(cfg.LogLevel, cfg.IsProduction)
	defer func(log *zap.Logger) {
		err := log.Sync()
		if err != nil {
			// No need to do anything here
		}
	}(log)
	log.Info("I am logging a log, using the logger")

	// zap doesn't do string interpolation, instead you create json like this
	//userID := 42
	//log.Info("Example log with json data",
	//	zap.String("requestId", "9739f75a-1aa8-11f0-ae1b-87947694378b"),
	//	zap.Int("user_id", userID),
	//	zap.Strings("items", []string{"book", "laptop", "coffee"}),
	//)

	// Initialize API clients

	// Set Gin's mode based on configuration (tbc what this does)
	if cfg.IsProduction {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Create Cache Client
	cacheAPI := api.NewCacheClient(cfg.CacheURL, log)

	// Create Handler
	handler := handlers.NewHandler(cacheAPI, log)

	// Create a new Gin router
	router := gin.New()

	// Configure the /page endpoint
	routes.SetupRoutes(router, handler)

	// Start the Gin server
	port := fmt.Sprintf(":%d", cfg.ServicePort)
	log.Info("Starting server", zap.String("address", port))
	if err := router.Run(port); err != nil {
		log.Fatal("Failed to run server", zap.Error(err))
	}
}
