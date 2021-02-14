// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package exhttp

import (
	"context"
	"github.com/ysicing/ext/logger/zlog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// SetupGracefulStop grace stop
func SetupGracefulStop(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zlog.Info("ShutDown Server ...")
	ShutDown(srv)
}

// ShutDown shutdown http
func ShutDown(srv *http.Server) {
	ctx, cancal := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancal()
	if err := srv.Shutdown(ctx); err != nil {
		zlog.Fatal("shutdown err: ", err)
	}
	select {
	case <-ctx.Done():
		zlog.Info("server exit timeout of 5 seconds.")
	default:

	}
	zlog.Info("server exited.")
}
