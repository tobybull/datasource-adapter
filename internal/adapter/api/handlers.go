package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// PageRequest represents the expected JSON body for the /page endpoint
type PageRequest struct {
	InputData string `json:"input_data" binding:"required"`
	UserID    int    `json:"user_id"`
}

// PageResponse represents the JSON response for the /page endpoint
type PageResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// PageHandler handles POST requests to /page
func PageHandler(c *gin.Context, logger *zap.Logger) {
	var req PageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("Invalid request body", zap.Error(err))
		c.JSON(http.StatusBadRequest, PageResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
		return
	}

	logger.Info("Received /page request", zap.String("input_data", req.InputData), zap.Int("user_id", req.UserID))

	// Process the request (in a real app, you'd do something more useful)
	processedData := map[string]interface{}{
		"processed_input": req.InputData + "-processed",
		"user_id":         req.UserID,
	}

	response := PageResponse{
		Status:  "success",
		Message: "Page request processed",
		Data:    processedData,
	}

	c.JSON(http.StatusOK, response)
}
