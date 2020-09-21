// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package ginmid

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ysicing/ext/logger"
)

const (
	errLogFormat = "requestid %v => %v | %v | %v | %v | %v | %v <= err: %v"
	logFormat    = "requestid %v => %v | %v | %v | %v | %v | %v "
)

// Log log
func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		if len(c.Errors) > 0 || c.Writer.Status() >= 500 {
			logger.Slog.Error(fmt.Sprintf(errLogFormat, GetRequestID(c), c.Writer.Status(), c.ClientIP(), c.Request.Method, path, query, latency, c.Errors.String()))
		} else {
			logger.Slog.Debug(fmt.Sprintf(logFormat, GetRequestID(c), c.Writer.Status(), c.ClientIP(), c.Request.Method, path, query, latency))
		}
	}
}
