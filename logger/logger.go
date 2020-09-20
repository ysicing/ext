// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package logger

import (
	"fmt"
	"github.com/ysicing/ext/utils/exos"
	"github.com/ysicing/ext/utils/extime"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

const (
	DefaultMaxSize  = 10 // MB
	DefaultBackups  = 3
	DefaultMaxAge   = 7 // days
	DefaultCompress = true
)

var (
	Log  *zap.Logger
	Slog *zap.SugaredLogger
)

func InitLogger(cfg *LogConfig) {
	encoder := getEncoder(cfg) // 编码器
	errPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zap.ErrorLevel
	})
	debugPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level < zap.InfoLevel
	})
	customPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zap.InfoLevel && level < zap.ErrorLevel
	})
	if cfg.Simple {
		writeSyncer := getLogWriterSimple()                                                                                        // 写日志
		core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.AddSync(os.Stdout)), zapcore.DebugLevel) // 如何写，写到哪, 什么级别写
		Log = zap.New(zapcore.NewTee(core)).WithOptions(cfg.debugMode()...)
		Slog = Log.Sugar()
		return
	}
	//writeSyncer := getLogWriter() // 写日志
	errCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(getLogWriter("err"), zapcore.AddSync(os.Stdout)), errPriority)
	debugCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(getLogWriter("debug"), zapcore.AddSync(os.Stdout)), debugPriority)
	customCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(getLogWriter("custom"), zapcore.AddSync(os.Stdout)), customPriority)
	Log = zap.New(zapcore.NewTee(debugCore, customCore, errCore)).WithOptions(cfg.debugMode()...)
	Slog = Log.Sugar()
}

func getEncoder(cfg *LogConfig) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = timeEncoder //zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	//encoderConfig.LineEnding = zapcore.DefaultLineEnding
	if cfg.JsonFormat {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func getLogWriter(loglevel string) zapcore.WriteSyncer {
	var logpath string
	if exos.IsLinux() {
		logpath = fmt.Sprintf("/var/log/gologger/%v/%v.log", extime.GetToday(), loglevel)
	} else {
		logpath = fmt.Sprintf("/tmp/gologger/%s/%v.log", extime.GetTodayHour(), loglevel)
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   logpath,
		MaxSize:    DefaultMaxSize,
		MaxBackups: DefaultBackups,
		MaxAge:     DefaultMaxAge,
		Compress:   DefaultCompress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getLogWriterSimple() zapcore.WriteSyncer {
	var logpath string
	if exos.IsLinux() {
		logpath = fmt.Sprintf("/var/log/gologger/%v/debug.log", extime.GetToday())
	} else {
		logpath = fmt.Sprintf("/tmp/gologger/%s/debug.log", extime.GetTodayHour())
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   logpath,
		MaxSize:    DefaultMaxSize,
		MaxBackups: DefaultBackups,
		MaxAge:     DefaultMaxAge,
		Compress:   DefaultCompress,
	}
	return zapcore.AddSync(lumberJackLogger)
}
