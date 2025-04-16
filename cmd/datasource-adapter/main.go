package main

import (
	"fmt"
	"os"

	"datasource-adapter/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Configuration error:", err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", cfg)

	//log := logger.SetupLogger(cfg.LogLevel, cfg.IsProduction)
	//defer log.Sync()
	//
	//// Set Gin's mode based on configuration
	//if cfg.IsProduction {
	//	gin.SetMode(gin.ReleaseMode)
	//} else {
	//	gin.SetMode(gin.DebugMode)
	//}
	//
	//// Create a new Gin router
	//router := gin.New()
	//
	//// Add logging middleware (example - you might want a more robust one)
	//router.Use(func(c *gin.Context) {
	//	c.Next()
	//	log.Info("Request",
	//		zap.String("method", c.Request.Method),
	//		zap.String("path", c.Request.URL.Path),
	//		zap.Int("status", c.Writer.Status()),
	//		zap.String("client_ip", c.ClientIP()),
	//		zap.String("user_agent", c.Request.UserAgent()),
	//	)
	//})
	//
	//// Define the /page endpoint and associate it with the handler
	//router.POST("/page", func(c *gin.Context) {
	//	api.PageHandler(c, log)
	//})
	//
	//// Start the Gin server
	//port := fmt.Sprintf(":%d", cfg.ServicePort)
	//log.Info("Starting server", zap.String("address", port))
	//if err := router.Run(port); err != nil {
	//	log.Fatal("Failed to run server", zap.Error(err))
	//}
}
