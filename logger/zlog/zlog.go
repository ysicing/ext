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

func Debug(f string, args ...interface{}) {
	Zlog.Debugf(f, args...)
}

func Info(f string, args ...interface{}) {
	Zlog.Infof(f, args...)
}

func Warn(f string, args ...interface{}) {
	Zlog.Warnf(f, args...)
}

func Error(f string, args ...interface{}) {
	Zlog.Errorf(f, args...)
}

func DPanic(f string, args ...interface{}) {
	Zlog.DPanicf(f, args...)
}

func Panic(f string, args ...interface{}) {
	Zlog.Panicf(f, args...)
}

func Fatal(f string, args ...interface{}) {
	Zlog.Panicf(f, args...)
}
