// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package zlog

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"testing"
)

func _warnHook() func(entry zapcore.Entry) error {
	return func(entry zapcore.Entry) error {
		if entry.Level == zapcore.DebugLevel {
			fmt.Println("debug")
			return nil
		}
		fmt.Println(entry.Level)
		return nil
	}
}

func TestInitZlog(t *testing.T) {
	hookfunc := append([]func(entry zapcore.Entry) error{}, _warnHook())
	logcfg := Config{Simple: true, ServiceName: "harbor-monitor", HookFunc: hookfunc}
	InitZlog(&logcfg)
	Debug("debug")
	Info("info")
}
