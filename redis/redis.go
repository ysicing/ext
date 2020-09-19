// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
)

type Client struct {
	RClient *redis.Pool
}

type Config struct {
	RedisCfg *RedisConfig
}

type RedisConfig struct {
	Maxidle     int
	Maxactive   int
	IdleTimeout time.Duration
	Host        string
	Port        int
	Password    string
}

const (
	DefaultRedisMaxIdle   = 30
	DefaultRedisMaxactive = 30
	DefaultIdleTimeout    = 200
)

func New(cfg *Config) (client *Client) {
	c := Client{}
	if cfg.RedisCfg != nil {
		if cfg.RedisCfg.Maxidle == 0 {
			cfg.RedisCfg.Maxidle = DefaultRedisMaxIdle
		}
		if cfg.RedisCfg.Maxactive == 0 {
			cfg.RedisCfg.Maxactive = DefaultRedisMaxactive
		}
		if cfg.RedisCfg.IdleTimeout == 0 {
			cfg.RedisCfg.IdleTimeout = DefaultIdleTimeout
		}
		if cfg.RedisCfg.Port == 0 {
			cfg.RedisCfg.Port = 6379
		}
		if len(cfg.RedisCfg.Host) == 0 {
			cfg.RedisCfg.Host = "127.0.0.1"
		}
		pool := &redis.Pool{
			MaxIdle:     cfg.RedisCfg.Maxidle,
			MaxActive:   cfg.RedisCfg.Maxactive,
			IdleTimeout: cfg.RedisCfg.IdleTimeout,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", fmt.Sprintf("%v:%v", cfg.RedisCfg.Host, cfg.RedisCfg.Port))
				if err != nil {
					return nil, err
				}
				if cfg.RedisCfg.Password != "" {
					if _, err := c.Do("AUTH", cfg.RedisCfg.Password); err != nil {
						c.Close()
						return nil, err
					}
				}
				return c, err
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		}
		c.RClient = pool
	}
	return &c
}

// RCSet set key
func (c *Client) RCSet(key string, value interface{}, expiretime ...int64) (bool, error) {
	client := c.RClient.Get()
	defer client.Close()
	data, err := json.Marshal(value)
	if err != nil {
		return false, err
	}
	reply, err := parsereply(client.Do("SET", key, data))
	if len(expiretime) > 0 {
		reply, err = parsereply(client.Do("EXPIRE", key, expiretime[0]))
	}
	return reply, err
}

// RCSet reset key
func (c *Client) RCReSet(key string, value interface{}, expiretime ...int64) (bool, error) {
	client := c.RClient.Get()
	defer client.Close()
	if c.RCExists(key) {
		if _, err := c.RCDel(key); err != nil {
			return false, err
		}
	}
	return c.RCSet(key, value, expiretime...)
}

// RCExists check key is exists
func (c *Client) RCExists(key string) bool {
	client := c.RClient.Get()
	defer client.Close()
	reply, _ := parsereply(client.Do("EXISTS", key))
	return reply
}

// RCDel 删除key
func (c *Client) RCDel(key string) (bool, error) {
	client := c.RClient.Get()
	defer client.Close()
	reply, err := parsereply(client.Do("DEL", key))
	return reply, err
}

// RCLikeDel 模糊删除
func (c *Client) RCLikeDel(key string) error {
	client := c.RClient.Get()
	defer client.Close()
	keys, err := redis.Strings(client.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}
	for _, ikey := range keys {
		_, err = c.RCDel(ikey)
		if err != nil {
			return err
		}
	}
	return nil
}

// RCGet get key
func (c *Client) RCGet(key string) ([]byte, error) {
	client := c.RClient.Get()
	defer client.Close()
	reply, err := redis.Bytes(client.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// RCLPUSH 列表
func (c *Client) RCLPUSH(queue, payload string) error {
	client := c.RClient.Get()
	defer client.Close()
	_, err := client.Do("LPUSH", queue, payload)
	return err
}

func (c *Client) Close() {
	c.RClient.Close()
}

func parsereply(reply interface{}, err error) (bool, error) {
	if err != nil {
		switch reply := reply.(type) {
		case int64:
			return reply != 0, nil
		case []byte:
			return strconv.ParseBool(string(reply))
		case nil:
			return false, errors.New("redigo: nil returned")
		case string:
			return false, errors.New(reply)
		}
		return false, fmt.Errorf("redigo: unexpected type for Bool, got type %T", reply)
	}
	return true, nil
}
