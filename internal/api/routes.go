package api

import (
	"datasource-adapter/internal/adapter"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupRoutes(router *gin.Engine, log *zap.Logger) {
	router.POST("/page", func(c *gin.Context) {
		adapter.PageHandler(c, log)
	})
}
