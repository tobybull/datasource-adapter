package handlers

import (
	"datasource-adapter/internal/api"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	CacheAPI *api.CacheClient
	logger   *zap.Logger
}

func NewHandler(cacheAPI *api.CacheClient, logger *zap.Logger) *Handler {
	return &Handler{
		CacheAPI: cacheAPI,
		logger:   logger,
	}
}

// PageResponse represents the JSON response for the /page endpoint
type PageResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (h *Handler) PageHandler(c *gin.Context) {
	var query Query
	if err := c.ShouldBindJSON(&query); err != nil {
		h.logger.Warn("Invalid request body", zap.Error(err))
		c.JSON(http.StatusBadRequest, PageResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
		return
	}

	h.logger.Info("Received /page request", zap.Any("query", query))

	// Build the datasource query

	// Call the datasource

	// Pull page of results

	// Send the results to the cache
	//results := map[string]interface{}{
	//	"key1": "value1",
	//	"key2": "value2",
	//	"key3": "value3",
	//}

	//h.CacheAPI.WriteToCache(results)

	// Return metadata response

	response := PageResponse{
		Status:  "success",
		Message: "Page request processed",
		Data:    query,
	}

	c.JSON(http.StatusOK, response)
}
