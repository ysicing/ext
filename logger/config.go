// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogConfig struct {
	Simple     bool
	HookFunc   func(zapcore.Entry) error
	JsonFormat bool
	CallerSkip bool
}

func (cfg *LogConfig) debugMode() []zap.Option {
	// AddCaller 输出文件名和行号
	// AddCallerSkip(1) 调用函数位置
	// AddStacktrace 输出堆栈
	var cfgopts []zap.Option
	if !cfg.Simple {
		cfgopts = append(cfgopts, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	}

	if cfg.CallerSkip {
		cfgopts = append(cfgopts, zap.AddCallerSkip(1))
	}

	if cfg.HookFunc != nil {
		cfgopts = append(cfgopts, zap.Hooks(cfg.HookFunc))
	}

	return cfgopts
}
