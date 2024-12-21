package glog

import (
	"io"
	"log/slog"
	"os"
	"sync"
)

type LogLevel string

const (
	LogLevelError LogLevel = "error"
	LogLevelWarn  LogLevel = "warn"
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
)

type LogConfig struct {
	Level  LogLevel
	Output io.Writer
}

type Logger struct {
	log *slog.Logger
}

var (
	defaultLogLevel  LogLevel  = LogLevelInfo
	defaultLogOutput io.Writer = os.Stdout
	once             sync.Once
	instance         *Logger
)

func New(serviceName string, logConfig ...*LogConfig) *Logger {
	var config *LogConfig

	if len(logConfig) > 0 && logConfig[0] != nil {
		config = logConfig[0]
	} else {
		config = &LogConfig{
			Level:  defaultLogLevel,
			Output: defaultLogOutput,
		}
	}

	once.Do(func() {
		handler := slog.NewJSONHandler(config.Output, &slog.HandlerOptions{
			Level: convertLogLevel(config.Level),
		})
		sLogger := slog.New(handler).With("service_name", serviceName)
		slog.SetDefault(sLogger)

		instance = &Logger{
			log: sLogger,
		}
	})

	return instance
}

func convertLogLevel(lvl LogLevel) slog.Level {
	switch lvl {
	case LogLevelError:
		return slog.LevelError
	case LogLevelWarn:
		return slog.LevelWarn
	case LogLevelDebug:
		return slog.LevelDebug
	default:
		return slog.LevelInfo
	}
}

func Info(message string, args ...interface{}) {
	instance.log.Info(message, args...)
}

func Fatal(message string, args ...interface{}) {
	instance.log.Error(message, args...)
	panic(message)
}

func Error(message string, args ...interface{}) {
	instance.log.Error(message, args...)
}

func Warn(message string, args ...interface{}) {
	instance.log.Warn(message, args...)
}

func Debug(message string, args ...interface{}) {
	instance.log.Debug(message, args...)
}
