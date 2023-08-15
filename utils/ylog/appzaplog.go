package ylog

import (
	"errors"
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	appInnerLog *zap.Logger
	logLevel    zap.AtomicLevel
)

// InitAppLog 根据options的设置,初始化日志系统。
// 注意默认是测试环境模式,需要设置线上模式的需要设置TestEnv(false)
func InitAppLog(options ...appZapOption) error {
	var (
		err error
	)
	config := defaultLogOptions
	for _, option := range options {
		option.apply(&config)
	}

	if logLevel, appInnerLog, err = zapLogInit(&config); err != nil {
		fmt.Printf("ZapLogInit err:%v", err)
		return err
	}

	appInnerLog = appInnerLog.WithOptions(zap.AddCallerSkip(1))

	fmt.Printf("ZapLogInit finished. level=%s", logLevel)
	return nil
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Debug(msg string, fields ...zapcore.Field) {
	appInnerLog.Debug(msg, fields...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Info(msg string, fields ...zapcore.Field) {
	appInnerLog.Info(msg, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Warn(msg string, fields ...zapcore.Field) {
	appInnerLog.Warn(msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Error(msg string, fields ...zapcore.Field) {
	appInnerLog.Error(msg, fields...)
}

// DPanic logs a message at DPanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// If the logger is in development mode, it then panics (DPanic means
// "development panic"). This is useful for catching errors that are
// recoverable, but shouldn't ever happen.
func DPanic(msg string, fields ...zapcore.Field) {
	appInnerLog.DPanic(msg, fields...)
}

// Panic logs a message at PanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then panics, even if logging at PanicLevel is disabled.
func Panic(msg string, fields ...zapcore.Field) {
	appInnerLog.Panic(msg, fields...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is disabled.
func Fatal(msg string, fields ...zapcore.Field) {
	appInnerLog.Fatal(msg, fields...)
}

func Sync() error {
	return appInnerLog.Sync()
}

func SetLogLevel(input string) error {
	level := strings.ToLower(input)
	switch level {
	case "debug", "info", "warn", "error", "fatal":
		// do nothing
	case "all":
		level = "debug"
	case "off", "none":
		level = "fatal"
	default:
		return errors.New("not support level: " + level)
	}
	err := logLevel.UnmarshalText([]byte(level))
	if err != nil {
		return fmt.Errorf("SetLogLevel: %v", err)
	}
	fmt.Println("ulog, SetLogLevel level " + level)
	return nil
}
