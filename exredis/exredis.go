// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package exredis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"time"
)

type RedisConfig struct {
	MaxIdle int
	MaxActive int
	IdleTimeout time.Duration
	RedisAddr string
	RedisPass string
}


var RedisConn *redis.Pool


func (rc *RedisConfig) Check() {}

func InitRedisSdk(cfg *RedisConfig) {
	RedisConn = &redis.Pool{
		MaxIdle:     cfg.MaxIdle,
		MaxActive:   cfg.MaxActive,
		IdleTimeout: cfg.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", cfg.RedisAddr)
			if err != nil {
				return nil, err
			}
			if cfg.RedisPass != "" {
				if _, err := c.Do("AUTH", cfg.RedisPass); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: pingRedis,
	}
}

func pingRedis(c redis.Conn, t time.Time) error {
	_, err := c.Do("PING")
	if err != nil {
		return err
	}
	return nil
}

func CaCheSetEXPIRE(key string, data interface{}, time int64) (bool, error) {
	rediscn := RedisConn.Get()
	defer rediscn.Close()
	value, err := json.Marshal(data)
	if err != nil {
		return false, err
	}
	reply, _ := redis.Bool(rediscn.Do("SET", key, value))
	rediscn.Do("EXPIRE", key, time)
	return reply, err
}

func CaCheSet(key string, data interface{}) (string, error) {
	rediscn := RedisConn.Get()
	defer rediscn.Close()
	value, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	reply, err := redis.String(rediscn.Do("SET", key, value))
	if err != nil {
		return "", err
	}
	return reply, nil
}

func CaCheReSet(key string, data interface{}, time ...int) (string, error) {
	rediscn := RedisConn.Get()
	defer rediscn.Close()
	if CacheExists(key) {
		if _, err := CacheDelete(key); err != nil {
			return "", err
		}
	}

	value, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	reply, err := redis.String(rediscn.Do("SET", key, value))
	if err != nil {
		return "", err
	}
	var ztime = 120
	if len(time) != 0 {
		ztime = time[0]
	}
	rediscn.Do("EXPIRE", key, ztime)
	return reply, err
}

func CacheExists(key string) bool {
	rediscn := RedisConn.Get()
	defer rediscn.Close()
	exists, err := redis.Bool(rediscn.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return exists
}

func CacheGet(key string) ([]byte, error) {
	rediscn := RedisConn.Get()
	defer rediscn.Close()
	reply, err := redis.Bytes(rediscn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func CacheLikeGet(key string) ([]string, error) {
	rediscn := RedisConn.Get()
	defer rediscn.Close()
	keys, err := redis.Strings(rediscn.Do("KEYS", key))
	if err != nil {
		return nil, err
	}
	return keys, nil
}

func CacheDelete(key string) (bool, error) {
	rediscn := RedisConn.Get()
	defer rediscn.Close()
	return redis.Bool(rediscn.Do("DEL", key))
}

func CacheLikeDeletes(key string) error {
	rediscn := RedisConn.Get()
	defer rediscn.Close()
	keys, err := redis.Strings(rediscn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}
	for _, key := range keys {
		_, err = CacheDelete(key)
		if err != nil {
			return err
		}
	}
	return nil
}

func CloseRedis() {
	RedisConn.Close()
}
