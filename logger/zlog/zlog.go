// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package zlog

import "go.uber.org/zap"

var (
	// Log log
	Log *zap.Logger
	// Zlog log sugar
	Zlog *zap.SugaredLogger
)

// InitZlog 初始化日志
func InitZlog(cfg *Config) {
	Log = zap.New(cfg.getCores()).WithOptions(cfg.debugMode()...)
	Zlog = Log.Sugar()
}

// Debug debug
func Debug(f string, args ...interface{}) {
	Zlog.Debugf(f, args...)
}

// Info info
func Info(f string, args ...interface{}) {
	Zlog.Infof(f, args...)
}

// Warn warn
func Warn(f string, args ...interface{}) {
	Zlog.Warnf(f, args...)
}

// Error error
func Error(f string, args ...interface{}) {
	Zlog.Errorf(f, args...)
}

// DPanic dpanic
func DPanic(f string, args ...interface{}) {
	Zlog.DPanicf(f, args...)
}

// Panic panic
func Panic(f string, args ...interface{}) {
	Zlog.Panicf(f, args...)
}

// Fatal fatal
func Fatal(f string, args ...interface{}) {
	Zlog.Panicf(f, args...)
}
