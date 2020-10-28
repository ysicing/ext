// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package main

import (
	"github.com/ysicing/ext/logger"
	"github.com/ysicing/ext/redis"
)

var RCClient *redis.Client

func init() {
	rediscfg := redis.RedisConfig{
		Maxidle:     31,
		Maxactive:   31,
		IdleTimeout: 200,
		Host:        "10.147.20.44",
		Port:        6379,
		Password:    "ahphu9nah9iuheid1aew2eiPei6Ach",
	}
	cfg := redis.Config{RedisCfg: &rediscfg}
	RCClient = redis.New(&cfg)

	logcfg := logger.Config{
		Simple:      true,
		HookFunc:    nil,
		JsonFormat:  false,
		ConsoleOnly: true,
	}
	logger.InitLogger(&logcfg)
}
