// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package exgin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ysicing/ext/file"
	"github.com/ysicing/ext/gerr"
	"github.com/ysicing/ext/logger/zlog"
	"github.com/ysicing/ext/misc"
	"github.com/ysicing/ext/zos"
	"github.com/ysicing/ext/ztime"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

func ExGin(debug bool) *gin.Engine {
	if debug || zos.IsMacOS() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.DisableConsoleColor()
	return gin.New()
}

// ExCors cors middleware
func ExCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, UPDATE, HEAD, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Access-Control-Request-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Max-Age", "3600")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

// ExLog log middleware
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
			msg := fmt.Sprintf("requestid %v => %v | %v | %v | %v | %v | %v <= err: %v", GetRID(c), misc.SRed("%v", c.Writer.Status()), c.ClientIP(), c.Request.Method, path, query, latency, c.Errors.String())
			zlog.Warn(msg)
			go file.Writefile(fmt.Sprintf("/tmp/%v.errreq.txt", ztime.GetToday()), msg)
		} else {
			zlog.Info("requestid %v => %v | %v | %v | %v | %v | %v ", GetRID(c), misc.SGreen("%v", c.Writer.Status()), c.ClientIP(), c.Request.Method, path, query, latency)
		}
	}
}

// Exrecovery recovery
func Exrecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if res, ok := err.(gerr.GinsError); ok {
					GinsData(c, nil, fmt.Errorf(res.Message))
					c.Abort()
					return
				}
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zlog.Error("Recovery from brokenPipe ---> path: %v, err: %v, request: %v", c.Request.URL.Path, err, string(httpRequest))
					c.AbortWithStatusJSON(200, gin.H{
						"data":      nil,
						"message":   "请求broken",
						"timestamp": ztime.NowUnix(),
						"code":      10500,
					})
					return
				} else {
					zlog.Error("Recovery from panic ---> err: %v, request: %v, stack: %v", err, string(httpRequest), string(debug.Stack()))
					c.AbortWithStatusJSON(200, gin.H{
						"data":      nil,
						"message":   "请求panic",
						"timestamp": ztime.NowUnix(),
						"code":      10500,
					})
					return
				}
			}
		}()
		c.Next()
	}
}

// ExRid rid 请求ID
func ExRid() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get(headerXRequestID)
		if requestID == "" {
			requestID = zos.GenUUID()
		}
		c.Set(headerXRequestID, requestID)
		c.Writer.Header().Set(headerXRequestID, requestID)
		c.Next()
	}
}
