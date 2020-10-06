// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package main

import (
	"github.com/ysicing/ext/logger"
	"github.com/ysicing/ext/sshutil"
)

func init() {
	cfg := logger.LogConfig{Simple: true}
	logger.InitLogger(&cfg)
}

func main() {
	ssh := sshutil.SSH{
		User:   "root",
		PkFile: "/Users/ysicing/.ssh/id_rsa",
	}
	if err := ssh.CmdAsync("10.147.20.45:22", "w"); err != nil {
		logger.Slog.Error(err.Error())
	}
	if err := ssh.CmdAsync("10.147.20.45:22", "w", true); err != nil {
		logger.Slog.Error(err.Error())
	}
}
