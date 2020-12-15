// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package ginmid

import (
	"github.com/gin-gonic/gin"
	"github.com/ysicing/ext/logger"
	"github.com/ysicing/ext/utils/exmisc"
	"time"
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
		if len(query) == 0 {
			query = " - "
		}
		if latency > time.Second*2 {
			logger.Slog.Warnf("[msg] api %v query %v", path, latency)
		}
		if len(c.Errors) > 0 || c.Writer.Status() >= 500 {
			logger.Slog.Errorf("requestid %v => %v | %v | %v | %v | %v | %v <= err: %v", GetRequestID(c), exmisc.SRed("%v", c.Writer.Status()), c.ClientIP(), c.Request.Method, path, query, latency, c.Errors.String())
		} else {
			logger.Slog.Infof("requestid %v => %v | %v | %v | %v | %v | %v ", GetRequestID(c), exmisc.SGreen("%v", c.Writer.Status()), c.ClientIP(), c.Request.Method, path, query, latency)
		}
	}
}
