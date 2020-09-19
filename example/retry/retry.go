// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package main

import (
	"fmt"
	"github.com/ysicing/ext/logger"
	"github.com/ysicing/ext/retry"
	"time"
)

func init() {
	cfg := logger.LogConfig{Simple: true}
	logger.InitLogger(&cfg)
}

func main() {
	i := 1 // lets assume we expect i to be a value of 8
	err := retry.DoFunc(10, 1*time.Second, func() error {
		logger.Slog.Debugf("trying for: %dth time\n", i)
		i++
		if i > 7 {
			return nil
		}
		return fmt.Errorf("i = %d is still low value", i)
	})

	if err != nil {
		logger.Slog.Panic(err)
	}

	logger.Slog.Info("Got our expected result: ", i)
}
