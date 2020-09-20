// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package main

import (
	"github.com/ysicing/ext/logger"
)

func main() {
	status, err := RCClient.RCSet("ha", "hahaha")
	logger.Slog.Debug(status, err)
	status1, err := RCClient.RCGet("hahaha")
	logger.Slog.Info(string(status1), err)
	status, err = RCClient.RCDel("hahax")
	logger.Slog.Error(status, err)
	status, err = RCClient.RCReSet("hahaha", "h", 10)
	logger.Slog.Debug(status, err)
	status, err = RCClient.RCDel("ha")
	logger.Slog.Exit(status, err)
	defer RCClient.Close()
}
