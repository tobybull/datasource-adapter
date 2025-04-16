package routes

import (
	"datasource-adapter/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, handler *handlers.Handler) {
	router.POST("/page", func(c *gin.Context) {
		handler.PageHandler(c)
	})
}
