// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package main

import (
	"github.com/ysicing/ext/logger"
	"github.com/ysicing/ext/utils/extime"
	"go.uber.org/zap/zapcore"
	"log"
)

func demohook() func(entry zapcore.Entry) error {
	return func(entry zapcore.Entry) error {
		if entry.Level < zapcore.ErrorLevel {
			return nil
		}
		log.Println("debug hook")
		return nil
	}
}

func init() {
	cfg := logger.LogConfig{Simple: false, HookFunc: demohook(), JsonFormat: true}
	logger.InitLogger(&cfg)
}

func main() {
	logger.Slog.Debug("debug")
	logger.Slog.Debugf("", "1", 2, 3, extime.GetToday())
	logger.Log.Sugar().Debug("1", 2, 3, extime.GetToday())
	logger.Slog.Info("info")
	logger.Slog.Error("error")
	logger.Slog.Exit("exit")
}
