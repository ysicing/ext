// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package logger

import (
	"go.uber.org/zap"
)

const (
	DefaultMaxSize  = 10 // MB
	DefaultBackups  = 3  // 备份
	DefaultMaxAge   = 7  // days
	DefaultCompress = true
)

var (
	Log  *zap.Logger
	Slog *zap.SugaredLogger
)

func InitLogger(cfg *Config) {
	Log = zap.New(cfg.getCores()).WithOptions(cfg.debugMode()...)
	Slog = Log.Sugar()
}
