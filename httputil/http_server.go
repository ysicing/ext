// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package httputil

import (
	"context"
	"github.com/ysicing/ext/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func SetupGracefulStop(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Slog.Info("ShutDown Server ...")
	ShutDown(srv)
}

func ShutDown(srv *http.Server) {
	ctx, cancal := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancal()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Slog.Fatal("shutdown err: ", err)
	}
	select {
	case <-ctx.Done():
		logger.Slog.Info("server exit timeout of 5 seconds.")
	default:

	}
	logger.Slog.Info("server exited.")
}
