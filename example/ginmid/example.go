// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ysicing/ext/ginmid"
	"github.com/ysicing/ext/logger"
	"go.uber.org/zap/zapcore"
	"log"
	"net/http"
	"runtime"
	"time"
)

func demohook() func(entry zapcore.Entry) error {
	return func(entry zapcore.Entry) error {
		if entry.Level < zapcore.ErrorLevel {
			return nil
		}
		log.Println(runtime.GOOS)
		return nil
	}
}

func init() {
	cfg := logger.LogConfig{Simple: true, HookFunc: demohook()}
	logger.InitLogger(&cfg)
}

func main() {
	// gin.SetMode(gin.ReleaseMode)
	// gin.DisableConsoleColor()
	r := gin.New()

	r.Use(ginmid.RequestID(), ginmid.PromMiddleware(nil), ginmid.Log(), ginmid.Recovery())

	// Example ping request.
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Example / request.
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "id:"+ginmid.GetRequestID(c))
	})

	// Example /metrics
	r.GET("/metrics", ginmid.PromHandler(promhttp.Handler()))

	// Example status 500
	r.GET("/err500", func(c *gin.Context) {
		c.String(http.StatusBadGateway, "id:"+ginmid.GetRequestID(c))
	})

	r.GET("/panic", func(c *gin.Context) {
		panic(1)
		c.String(http.StatusBadGateway, "id:"+ginmid.GetRequestID(c))
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
