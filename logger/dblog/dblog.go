// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package dblog

import (
	"context"
	"fmt"
	"github.com/ysicing/ext/file"
	"github.com/ysicing/ext/ztime"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

const (
	traceStr     = "%s [%.3fms] [rows:%v] %s"
	traceWarnStr = "%s %s [%.3fms] [rows:%v] %s"
	traceErrStr  = "%s %s [%.3fms] [rows:%v] %s"
)

// Logger logger
type Logger struct {
	Zlog          *zap.SugaredLogger
	Loglevel      logger.LogLevel
	SlowThreshold time.Duration
}

// New new logger
func New(zlog *zap.SugaredLogger, debug ...bool) Logger {
	if len(debug) != 0 && debug[0] {
		return Logger{
			Zlog:          zlog,
			Loglevel:      logger.Info,
			SlowThreshold: 1 * time.Second,
		}
	}
	return Logger{
		Zlog:          zlog,
		Loglevel:      logger.Warn,
		SlowThreshold: 1 * time.Second,
	}
}

// LogMode logmode
func (l Logger) LogMode(loglevel logger.LogLevel) logger.Interface {
	return Logger{
		Zlog:     l.Zlog,
		Loglevel: loglevel,
	}
}

// Info info
func (l Logger) Info(ctx context.Context, str string, args ...interface{}) {
	l.Zlog.Debugf(str, args...)
}

// Warn warn
func (l Logger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.Loglevel == logger.Warn {
		l.Zlog.Warnf("warn "+str, args...)
	}
	l.Zlog.Warnf("warn "+str, args...)
}

// Error error
func (l Logger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.Loglevel == logger.Error {
		l.Zlog.Errorf("err "+str, args...)
	}
	l.Zlog.Errorf("err "+str, args...)
}

// Trace trace
func (l Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.Loglevel > 0 {
		elapsed := time.Since(begin)
		switch {
		case err != nil && l.Loglevel >= logger.Error:
			sql, rows := fc()
			if rows == -1 || rows == 0 || err == gorm.ErrRecordNotFound {
				go file.Writefile(fmt.Sprintf("/tmp/%v.dbnotfound.txt", ztime.GetToday()), sql)
				l.Zlog.Debugf(traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				l.Zlog.Errorf(traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case l.SlowThreshold != 0 && elapsed > l.SlowThreshold && l.Loglevel >= logger.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
			go file.Writefile(fmt.Sprintf("/tmp/%v.slowsql.txt", ztime.GetToday()), sql+" "+slowLog)
			if rows == -1 {
				l.Zlog.Warnf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				l.Zlog.Warnf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case l.Loglevel >= logger.Info:
			sql, rows := fc()
			if rows == -1 {
				l.Zlog.Infof(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				l.Zlog.Infof(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		}
	}
}
