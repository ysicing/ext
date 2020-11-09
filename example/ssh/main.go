// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package main

import (
	"github.com/ysicing/ext/logger"
	"github.com/ysicing/ext/sshutil"
)

func init() {
	cfg := logger.Config{
		Simple:      true,
		ConsoleOnly: true,
		//ConsoleOnly: false,
		//LogConfig: logger.LogConfig{
		//	LogPath:    "/tmp/sshlog",
		//},
	}
	logger.InitLogger(&cfg)
}

func main() {
	ssh := sshutil.SSH{
		User:   "root",
		PkFile: "/Users/ysicing/.ssh/id_rsa",
	}
	if err := ssh.CmdAsync("172.16.16.55:22", "w"); err != nil {
		logger.Slog.Error(err.Error())
	}
	if err := ssh.CmdAsync("172.16.16.55:22", "w", true); err != nil {
		logger.Slog.Error(err.Error())
	}
	//ssh2 := sshutil.SSH{
	//	User:     "root",
	//	Password: "vagrant",
	//}
	//if err := ssh2.CmdAsync("11.11.11.11:22", "w"); err != nil {
	//	logger.Slog.Error(err.Error())
	//}
	if err := ssh.CmdAsync("172.16.16.55:22", "docker system prune", false); err != nil {
		logger.Slog.Error(err.Error())
	}
}
