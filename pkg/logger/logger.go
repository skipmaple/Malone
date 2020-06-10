// Copyright © 2020. Drew Lee. All rights reserved.

package logger

import (
	"KarlMalone/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

func init() {
	logger = InitLogger(config.Logger.Dir, false)
}

// when parameter isDev set true will print log's locate filename and line number, default set false
func InitLogger(logPath string, isDev bool) *zap.Logger {
	var logger *zap.Logger

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	infoHook := lumberjack.Logger{
		Filename:   logPath + "/" + "info.log",
		MaxSize:    1 << 8,
		MaxBackups: 10,
		MaxAge:     30,
		Compress:   true,
	}
	errHook := lumberjack.Logger{
		Filename:   logPath + "/" + "err.log",
		MaxSize:    1 << 8,
		MaxBackups: 10,
		MaxAge:     30,
		Compress:   true,
	}
	debugHook := lumberjack.Logger{
		Filename:   logPath + "/event/" + "event.log",
		MaxSize:    1 << 8,
		MaxBackups: 10,
		MaxAge:     30,
		Compress:   true,
	}

	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(&infoHook)),
			&EnableLog{zapcore.InfoLevel},
		),
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(&errHook)),
			&EnableLog{zapcore.ErrorLevel},
		),
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(&debugHook)),
			&EnableLog{zapcore.DebugLevel},
		),
	)

	if isDev {
		logger = zap.New(core, zap.Development(), zap.AddCaller())
	} else {
		logger = zap.New(core)
	}
	return logger
}

type EnableLog struct {
	Level zapcore.Level
}

func (e *EnableLog) Enabled(Level zapcore.Level) bool {
	return e.Level == Level
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	logger.DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}
