// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package mid

import (
	"github.com/gin-gonic/gin"
	"github.com/ysicing/ext/utils/exid"
)

const headerXRequestID = "X-Request-ID"

// RequestID 请求ID
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get(headerXRequestID)
		if requestID == "" {
			requestID = exid.GenUUID()
		}
		c.Set(headerXRequestID, requestID)
		c.Writer.Header().Set(headerXRequestID, requestID)
		c.Next()
	}
}

// 获取ID
func GetRequestID(c *gin.Context) string {
	return c.Writer.Header().Get(headerXRequestID)
}
