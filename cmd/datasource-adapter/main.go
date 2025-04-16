package main

import (
	"datasource-adapter/internal/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"

	"datasource-adapter/internal/api"
	"datasource-adapter/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Configuration error:", err)
		os.Exit(1)
	}
	fmt.Printf("The config --> %+v\n", cfg)

	log := logger.SetupLogger(cfg.LogLevel, cfg.IsProduction)
	defer func(log *zap.Logger) {
		err := log.Sync()
		if err != nil {
			// No need to do anything here
		}
	}(log)

	log.Info("I am logging a log, using the logger")

	// zap doesn't do string interpolation, instead you create json like this
	userID := 42
	log.Info("Received request",
		zap.String("requestId", "9739f75a-1aa8-11f0-ae1b-87947694378b"),
		zap.Int("user_id", userID),
		zap.Strings("items", []string{"book", "laptop", "coffee"}),
	)

	// Set Gin's mode based on configuration (tbc what this does)
	if cfg.IsProduction {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Create a new Gin router
	router := gin.New()

	// Example logging middleware
	router.Use(func(c *gin.Context) {
		c.Next()
		log.Info("Request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.String("client_ip", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
		)
	})

	// Configure the /page endpoint
	api.SetupRoutes(router, log)

	// Start the Gin server
	port := fmt.Sprintf(":%d", cfg.ServicePort)
	log.Info("Starting server", zap.String("address", port))
	if err := router.Run(port); err != nil {
		log.Fatal("Failed to run server", zap.Error(err))
	}
}
