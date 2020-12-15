// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ysicing/ext/e"
	"github.com/ysicing/ext/ginmid"
	"github.com/ysicing/ext/httputil"
	"github.com/ysicing/ext/logger"
	"log"
	"net/http"
	"time"
)

func demohook() func(entry zapcore.Entry) error {
	return func(entry zapcore.Entry) error {
		if entry.Level >= zapcore.ErrorLevel {
			log.Println("err hook")
			return nil
		}
		if entry.Level == zapcore.WarnLevel {
			log.Println("warn log")
			return nil
		}
		return nil
	}
}

func init() {
	cfg := logger.Config{
		Simple:      true,
		HookFunc:    logger.Defaulthook(),
		ConsoleOnly: false,
		JsonFormat:  false,
		LogConfig: logger.LogConfig{
			LogPath: "/tmp/ginlogs",
		},
	}
	logger.InitLogger(&cfg)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	// gin.DisableConsoleColor()
	r := gin.New()

	r.Use(ginmid.RequestID(), ginmid.Log(), ginmid.Recovery(), ginmid.Ginprom())

	// Example ping request.
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, e.Done("pong "+fmt.Sprint(time.Now().Unix())))
	})

	// Example / request.
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, e.Done("id:"+ginmid.GetRequestID(c)))
	})

	// Example /metrics
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Example status 500
	r.GET("/err500", func(c *gin.Context) {
		c.JSON(500, e.Done("id:"+ginmid.GetRequestID(c)))
	})

	r.GET("/panic", func(c *gin.Context) {
		panic(1)
		c.JSON(500, e.Done("id:"+ginmid.GetRequestID(c)))
	})

	// Listen and Server in 0.0.0.0:8080
	logger.Slog.Warn("startup")
	//r.Run(":8080")
	srv := http.Server{Addr: ":8080", Handler: r}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Slog.Fatal(err)
		}
	}()
	httputil.SetupGracefulStop(&srv)
}
