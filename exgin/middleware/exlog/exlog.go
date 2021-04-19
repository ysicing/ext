package exlog

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ysicing/ext/color"
	"github.com/ysicing/ext/exgin/middleware/rid"
	"github.com/ysicing/ext/file"
	"github.com/ysicing/ext/logger/zlog"
	"github.com/ysicing/ext/ztime"
	"time"
)

// ExLog exlog middleware
func ExLog() gin.HandlerFunc {
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
		if latency > time.Second*1 {
			zlog.Warn("[msg] api %v query %v", path, latency)
		}
		if len(c.Errors) > 0 || c.Writer.Status() >= 500 {
			msg := fmt.Sprintf("requestid %v => %v | %v | %v | %v | %v | %v <= err: %v", rid.GetRID(c), color.SRed("%v", c.Writer.Status()), c.ClientIP(), c.Request.Method, path, query, latency, c.Errors.String())
			zlog.Warn(msg)
			go file.Writefile(fmt.Sprintf("/tmp/%v.errreq.txt", ztime.NowDay()), msg)
		} else {
			zlog.Info("requestid %v => %v | %v | %v | %v | %v | %v ", rid.GetRID(c), color.SGreen("%v", c.Writer.Status()), c.ClientIP(), c.Request.Method, path, query, latency)
		}
	}
}
