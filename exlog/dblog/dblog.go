// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package dblog

import (
	"context"
	"fmt"
	"github.com/kunnos/zap"
	glog "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

const (
	traceStr     = "%s [%.3fms] [rows:%v] %s"
	traceWarnStr = "%s %s [%.3fms] [rows:%v] %s"
	traceErrStr  = "%s %s [%.3fms] [rows:%v] %s"
)

const (
	LogInfo   = glog.Info
	LogWarn   = glog.Warn
	LogError  = glog.Error
	LogSilent = glog.Silent
)

type Logger struct {
	Zlog          *zap.SugaredLogger
	Loglevel      glog.LogLevel
	SlowThreshold time.Duration
}

func New(zlog *zap.SugaredLogger, debug ...bool) Logger {
	if len(debug) != 0 && debug[0] {
		return Logger{
			Zlog:          zlog,
			Loglevel:      glog.Info,
			SlowThreshold: 1 * time.Second,
		}
	}
	return Logger{
		Zlog:          zlog,
		Loglevel:      glog.Warn,
		SlowThreshold: 1 * time.Second,
	}
}

func (l Logger) LogMode(loglevel glog.LogLevel) glog.Interface {
	return Logger{
		Zlog:     l.Zlog,
		Loglevel: loglevel,
	}
}

func (l Logger) Info(ctx context.Context, str string, args ...interface{}) {
	//if l.Loglevel == glog.Info {
	//	l.Zlog.Hookf("info", args...)
	//}
	l.Zlog.Debugf(str, args...)
}
func (l Logger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.Loglevel == glog.Warn {
		l.Zlog.Warnf("warn "+str, args...)
	}
	l.Zlog.Warnf("warn "+str, args...)
}
func (l Logger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.Loglevel == glog.Error {
		l.Zlog.Errorf("err "+str, args...)
	}
	l.Zlog.Errorf("err "+str, args...)
}
func (l Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.Loglevel > 0 {
		elapsed := time.Since(begin)
		switch {
		case err != nil && l.Loglevel >= glog.Error:
			sql, rows := fc()
			if rows == -1 {
				l.Zlog.Errorf(traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				l.Zlog.Errorf(traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case l.SlowThreshold != 0 && elapsed > l.SlowThreshold && l.Loglevel >= glog.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
			if rows == -1 {
				l.Zlog.Warnf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				l.Zlog.Warnf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case l.Loglevel >= glog.Info:
			sql, rows := fc()
			if rows == -1 {
				l.Zlog.Infof(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				l.Zlog.Infof(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		}
	}
}
