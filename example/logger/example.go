// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package main

import (
	"github.com/ysicing/ext/logger"
	"github.com/ysicing/ext/utils/extime"
)

func init() {
	cfg := logger.Config{
		Simple:      false,
		HookFunc:    logger.Defaulthook(),
		JsonFormat:  true,
		CallerSkip:  false,
		ConsoleOnly: false,
		LogConfig:   logger.LogConfig{},
	}
	logger.InitLogger(&cfg)
}

func main() {
	logger.Slog.Debug("debug")
	logger.Slog.Debugf("", "1", 2, 3, extime.GetToday())
	logger.Log.Sugar().Debug("1", 2, 3, extime.GetToday())
	logger.Slog.Info("info")
	logger.Slog.Error("error")
	logger.Slog.Hook("hook")
	logger.Slog.Fatal("exit")
}
