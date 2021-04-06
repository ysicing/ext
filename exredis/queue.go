// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package exredis

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/ysicing/ext/logger/zlog"
)

// ReadHeader read list header
func ReadHeader(queue string) string {
	rc := RedisConn.Get()
	defer rc.Close()
	reply, err := redis.String(rc.Do("LPOP", queue)) // 头部
	if err != nil {
		if err != redis.ErrNil {
			zlog.Debug("lpop queue:%s failed, err: %v", queue, err)
		}
		return ""
	}
	return reply
}

// ReadEnd read list end
func ReadEnd(queue string) string {
	rc := RedisConn.Get()
	defer rc.Close()
	reply, err := redis.String(rc.Do("RPOP", queue)) // 尾部
	if err != nil {
		if err != redis.ErrNil {
			zlog.Debug("rpop queue:%s failed, err: %v", queue, err)
		}
		return ""
	}
	return reply
}

// WriteHeader 列表
func WriteHeader(queue, message string) error {
	rc := RedisConn.Get()
	defer rc.Close()
	_, err := rc.Do("LPUSH", queue, message) // 头部
	if err != nil {
		zlog.Error("LPUSH:%v fail, message:%v, err: %v", queue, message, err)
	}
	return err
}

// WriteEnd 列表
func WriteEnd(queue, message string) error {
	rc := RedisConn.Get()
	defer rc.Close()
	_, err := rc.Do("RPUSH", queue, message) // 尾部
	if err != nil {
		zlog.Error("RPUSH:%v fail, message:%v, err: %v", queue, message, err)
	}
	return err
}

// WriteData write data to queue
func WriteData(queue string, data interface{}) error {
	if data == nil {
		return fmt.Errorf("message is nil")
	}
	bs, err := json.Marshal(data)
	if err != nil {
		zlog.Error("marshal message failed, message: %+v, err: %v", data, err)
		return err
	}
	zlog.Debug("write message to queue, message:%+v, queue:%s", data, queue)
	return WriteHeader(queue, string(bs))
}
