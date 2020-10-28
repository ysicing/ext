// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package logger

import (
	"fmt"
	"github.com/kunnos/zap"
	"github.com/kunnos/zap/zapcore"
	"github.com/ysicing/ext/utils/exos"
	"github.com/ysicing/ext/utils/extime"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

type Config struct {
	Simple      bool //
	HookFunc    func(zapcore.Entry) error
	JsonFormat  bool
	CallerSkip  bool
	ConsoleOnly bool // 不写日志文件
	LogConfig   LogConfig
}

type LogConfig struct {
	LogPath    string // 日志路径
	MaxSize    int    // 日志大小
	MaxBackups int    // 备份
	MaxAge     int    // 天数
}

func (cfg *Config) debugMode() []zap.Option {
	// AddCaller 输出文件名和行号
	// AddCallerSkip(1) 调用函数位置
	// AddStacktrace 输出堆栈
	var cfgopts []zap.Option
	if !cfg.Simple {
		cfgopts = append(cfgopts, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	}

	if cfg.CallerSkip {
		cfgopts = append(cfgopts, zap.AddCallerSkip(1))
	}

	if cfg.HookFunc != nil {
		cfgopts = append(cfgopts, zap.Hooks(cfg.HookFunc))
	}
	return cfgopts
}

func (cfg *Config) getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = timeEncoder //zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	if !cfg.ConsoleOnly && cfg.JsonFormat {
		// 写日志且json
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func (cfg *Config) logsyncer(lvl ...string) zapcore.WriteSyncer {
	var wss []zapcore.WriteSyncer
	if !cfg.ConsoleOnly {
		if cfg.Simple {
			wss = append(wss, cfg.getLogWriter())
		} else {
			if len(lvl) == 0 {
				lvl = append(lvl, "debug")
			}
			wss = append(wss, cfg.getLogWriter(lvl[0]))
		}
	}
	wss = append(wss, zapcore.AddSync(os.Stdout))
	return zapcore.NewMultiWriteSyncer(wss...)
}

func (cfg *Config) GetLogConfig() *LogConfig {
	var logcfg LogConfig
	logcfg = cfg.LogConfig
	if logcfg.LogPath == "" {
		if exos.IsLinux() {
			logcfg.LogPath = fmt.Sprintf("/var/log/gologger/%v", extime.GetToday())
		} else {
			logcfg.LogPath = fmt.Sprintf("/tmp/gologger/%v", extime.GetToday())
		}
	}
	if logcfg.MaxAge <= DefaultMaxSize {
		logcfg.MaxAge = DefaultMaxSize
	}
	if logcfg.MaxBackups <= DefaultBackups {
		logcfg.MaxBackups = DefaultBackups
	}
	if logcfg.MaxSize <= DefaultMaxSize {
		logcfg.MaxSize = DefaultMaxSize
	}
	return &logcfg
}

func (cfg *Config) getLogWriter(loglevel ...string) zapcore.WriteSyncer {
	logcfg := cfg.GetLogConfig()
	var logpath string
	if len(loglevel) > 0 {
		logpath = fmt.Sprintf("%v/%v.log", logcfg.LogPath, loglevel[0])
	} else {
		logpath = fmt.Sprintf("%v.log", logcfg.LogPath)
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   logpath,
		MaxSize:    logcfg.MaxSize,
		MaxBackups: logcfg.MaxBackups,
		MaxAge:     logcfg.MaxAge,
		Compress:   DefaultCompress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func (cfg *Config) getCores() zapcore.Core {
	var cors []zapcore.Core
	encoder := cfg.getEncoder() // 编码器
	if cfg.Simple {
		cors = append(cors, zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(cfg.logsyncer()), zapcore.DebugLevel))
	} else {
		errPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			return level >= zap.ErrorLevel
		})
		debugPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			return level < zap.InfoLevel
		})
		customPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			return level >= zap.InfoLevel && level < zap.ErrorLevel
		})
		errCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(cfg.logsyncer("err")), errPriority)
		debugCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(cfg.logsyncer("debug")), debugPriority)
		infoCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(cfg.logsyncer("info")), customPriority)
		cors = append(cors, errCore, debugCore, infoCore)
	}
	return zapcore.NewTee(cors...)
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func Defaulthook() func(entry zapcore.Entry) error {
	return func(entry zapcore.Entry) error {
		if entry.Level == zapcore.ErrorLevel {
			fmt.Println("err hook")
			return nil
		}
		if entry.Level == zapcore.HookLevel {
			fmt.Println("hook hook")
			return nil
		}
		return nil
	}
}
