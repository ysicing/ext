// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package gormv2

import (
	"context"
	"github.com/kunnos/zap"
	"gorm.io/gorm/logger"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type Logger struct {
	ZapLogger        *zap.Logger
	LogLevel         logger.LogLevel
	SlowThreshold    time.Duration
	SkipCallerLookup bool
}

func New(zapLogger *zap.Logger) Logger {
	return Logger{
		ZapLogger:        zapLogger,
		LogLevel:         logger.Warn,
		SlowThreshold:    100 * time.Millisecond,
		SkipCallerLookup: false,
	}
}

func (l Logger) SetAsDefault() {
	logger.Default = l
}

func (l Logger) LogMode(level logger.LogLevel) logger.Interface {
	return Logger{
		ZapLogger:        l.ZapLogger,
		SlowThreshold:    l.SlowThreshold,
		LogLevel:         level,
		SkipCallerLookup: l.SkipCallerLookup,
	}
}

func (l Logger) Info(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < logger.Info {
		return
	}
	l.logger().Sugar().Debugf(str, args...)
}

func (l Logger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < logger.Warn {
		return
	}
	l.logger().Sugar().Warnf(str, args...)
}

func (l Logger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < logger.Error {
		return
	}
	l.logger().Sugar().Errorf(str, args...)
}

func (l Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= 0 {
		return
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= logger.Error:
		sql, rows := fc()
		l.logger().Error("trace", zap.Error(err), zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	case l.SlowThreshold != 0 && elapsed > l.SlowThreshold && l.LogLevel >= logger.Warn:
		sql, rows := fc()
		l.logger().Warn("trace", zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	case l.LogLevel >= logger.Info:
		sql, rows := fc()
		l.logger().Debug("trace", zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	}
}

var (
	gormPackage    = filepath.Join("gorm.io", "gorm")
	zapgormPackage = filepath.Join("github.com/ysicing/ext/exlog", "gormv2")
)

func (l Logger) logger() *zap.Logger {
	for i := 2; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		switch {
		case !ok:
		case strings.HasSuffix(file, "_test.go"):
		case strings.Contains(file, gormPackage):
		case strings.Contains(file, zapgormPackage):
		default:
			return l.ZapLogger.WithOptions(zap.AddCallerSkip(i))
		}
	}
	return l.ZapLogger
}
