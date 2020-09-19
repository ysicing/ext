// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exos

import (
	"os"
)

// GetEnv 获取环境变量
func GetEnv(envstr, fallback string) string {
	if e := os.Getenv(envstr); e != "" {
		return e
	}
	return fallback
}
