// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kunnos/zap/zapcore"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ysicing/ext/e"
	"github.com/ysicing/ext/ginmid"
	"github.com/ysicing/ext/logger"
	"log"
	"time"
)

func demohook() func(entry zapcore.Entry) error {
	return func(entry zapcore.Entry) error {
		if entry.Level >= zapcore.ErrorLevel {
			log.Println("err hook")
			return nil
		}
		if entry.Level == zapcore.HookLevel {
			log.Println("hook hook")
			return nil
		}
		return nil
	}
}

func init() {
	cfg := logger.LogConfig{Simple: true, HookFunc: demohook()}
	logger.InitLogger(&cfg)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	// gin.DisableConsoleColor()
	r := gin.New()

	r.Use(ginmid.RequestID(), ginmid.PromMiddleware(nil), ginmid.Log(), ginmid.Recovery())

	// Example ping request.
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, e.Done("pong "+fmt.Sprint(time.Now().Unix())))
	})

	// Example / request.
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, e.Done("id:"+ginmid.GetRequestID(c)))
	})

	// Example /metrics
	r.GET("/metrics", ginmid.PromHandler(promhttp.Handler()))

	// Example status 500
	r.GET("/err500", func(c *gin.Context) {
		c.JSON(500, e.Done("id:"+ginmid.GetRequestID(c)))
	})

	r.GET("/panic", func(c *gin.Context) {
		panic(1)
		c.JSON(500, e.Done("id:"+ginmid.GetRequestID(c)))
	})

	// Listen and Server in 0.0.0.0:8080
	logger.Slog.Hook("startup")
	r.Run(":8080")
}
