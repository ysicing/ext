package rid

import (
	"github.com/gin-gonic/gin"
	"github.com/ysicing/ext/exhash"
)

const headerXRequestID = "X-Request-ID"

// ExRid rid 请求ID
func ExRid() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get(headerXRequestID)
		if requestID == "" {
			requestID = exhash.GenUUID()
		}
		c.Set(headerXRequestID, requestID)
		c.Writer.Header().Set(headerXRequestID, requestID)
		c.Next()
	}
}

// GetRID 获取ID
func GetRID(c *gin.Context) string {
	return c.Writer.Header().Get(headerXRequestID)
}
