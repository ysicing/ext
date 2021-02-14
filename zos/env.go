// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package zos

import "os"

// GetEnv 获取环境变量
func GetEnv(envstr string, fallback ...string) string {
	e := os.Getenv(envstr)
	if e == "" && len(fallback) > 0 {
		e = fallback[0]
	}
	return e
}
