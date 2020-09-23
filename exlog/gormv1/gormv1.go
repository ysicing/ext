// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package gormv1

import (
	"github.com/kunnos/zap"
	"time"
)

type GormV1Logger struct {
	zap *zap.SugaredLogger
}

func NewGormV1(logger *zap.SugaredLogger) GormV1Logger {
	return GormV1Logger{zap: logger}
}

func (gmv1 GormV1Logger) Print(values ...interface{}) {
	if len(values) < 2 {
		return
	}

	switch values[0] {
	case "sql":
		gmv1.zap.Debug("gorm.debug.sql",
			zap.String("query", values[3].(string)),
			zap.Any("values", values[4]),
			zap.Duration("duration", values[2].(time.Duration)),
			zap.Int64("affected-rows", values[5].(int64)),
			zap.String("source", values[1].(string)), // if AddCallerSkip(6) is well defined, we can safely remove this field
		)
	default:
		gmv1.zap.Debug("gorm.debug.other",
			zap.Any("values", values[2:]),
			zap.String("source", values[1].(string)), // if AddCallerSkip(6) is well defined, we can safely remove this field
		)
	}
}
