package api

import (
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"time"
)

type CacheClient struct {
	Client  *resty.Client
	BaseURL string
	logger  *zap.Logger
}

func NewCacheClient(baseURL string, logger *zap.Logger) *CacheClient {
	return &CacheClient{
		Client: resty.New().
			SetTimeout(10 * time.Second).
			SetRetryCount(3).
			SetRetryWaitTime(500 * time.Millisecond),
		BaseURL: baseURL,
		logger:  logger,
	}
}

func (cache *CacheClient) WriteToCache(records map[string]interface{}) {
	var cacheResponse map[string]interface{}
	resp, err := cache.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(records).
		SetResult(&cacheResponse).
		Post(cache.BaseURL + "/cache/some/path")

	if err != nil || resp.IsError() {
		cache.logger.Error("Failed to call API 2",
			zap.Error(err),
			zap.Int("status", resp.StatusCode()),
		)
		// return some error here
		return
	}
}
