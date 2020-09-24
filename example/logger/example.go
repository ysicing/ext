// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package main

import (
	"github.com/kunnos/zap/zapcore"
	"github.com/ysicing/ext/logger"
	"github.com/ysicing/ext/utils/extime"
	"log"
)

func demohook() func(entry zapcore.Entry) error {
	return func(entry zapcore.Entry) error {
		if entry.Level == zapcore.ErrorLevel {
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
	cfg := logger.LogConfig{Simple: false, HookFunc: demohook(), JsonFormat: true}
	logger.InitLogger(&cfg)
}

func main() {
	logger.Slog.Debug("debug")
	logger.Slog.Debugf("", "1", 2, 3, extime.GetToday())
	logger.Log.Sugar().Debug("1", 2, 3, extime.GetToday())
	logger.Slog.Info("info")
	logger.Slog.Error("error")
	logger.Slog.Hook("hook")
	//logger.Slog.Exit(0, "exit")
	logger.Slog.Exit(-1, "exit")
}
