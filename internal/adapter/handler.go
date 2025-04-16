package adapter

import (
	"datasource-adapter/internal/api"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// PageResponse represents the JSON response for the /page endpoint
type PageResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func PageHandler(c *gin.Context, logger *zap.Logger) {
	var query api.Query
	if err := c.ShouldBindJSON(&query); err != nil {
		logger.Warn("Invalid request body", zap.Error(err))
		c.JSON(http.StatusBadRequest, PageResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
		return
	}

	logger.Info("Received /page request", zap.Any("query", query))

	// Build the datasource query

	// Call the datasource

	// Pull page of results

	// Send the results to the cache

	// Return metadata response

	response := PageResponse{
		Status:  "success",
		Message: "Page request processed",
		Data:    "some data",
	}

	c.JSON(http.StatusOK, response)
}
