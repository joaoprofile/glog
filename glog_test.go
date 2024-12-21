package glog

import (
	"bytes"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoggingFunctions(t *testing.T) {
	logger := New("test_service")
	assert.NotNil(t, logger, "Logger instance should not be nil")
	assert.NotNil(t, logger.log, "Logger should have a slog.Logger instance")

	Info("This is an info message")

	Error("This is an error message")

	Warn("This is a warning message")

	Debug("This is a debug message")

	defer func() {
		if r := recover(); r != nil {
			assert.Contains(t, r.(string), "This is a fatal message", "Panic message should contain expected message")
		}
	}()
	Fatal("This is a fatal message")

}

func TestConvertLogLevel(t *testing.T) {
	assert.Equal(t, convertLogLevel(LogLevelError), slog.LevelError, "LogLevelError should map to slog.LevelError")
	assert.Equal(t, convertLogLevel(LogLevelWarn), slog.LevelWarn, "LogLevelWarn should map to slog.LevelWarn")
	assert.Equal(t, convertLogLevel(LogLevelDebug), slog.LevelDebug, "LogLevelDebug should map to slog.LevelDebug")
	assert.Equal(t, convertLogLevel(LogLevelInfo), slog.LevelInfo, "LogLevelInfo should map to slog.LevelInfo")
	assert.Equal(t, convertLogLevel("unknown"), slog.LevelInfo, "Unknown log level should default to slog.LevelInfo")
}

func TestLoggerInstance(t *testing.T) {
	logger := New("test_service")
	assert.NotNil(t, logger, "Logger instance should not be nil")
	assert.NotNil(t, logger.log, "Logger should have a slog.Logger instance")

	secondLogger := New("another_service")
	assert.Equal(t, logger, secondLogger, "Logger should maintain singleton instance")
}

func TestDefaultLoggerConfiguration(t *testing.T) {
	New("default_service")

	assert.Equal(t, defaultLogLevel, LogLevelInfo, "Default log level should be info")
}

func TestLoggingWithCustomConfig(t *testing.T) {
	buffer := &bytes.Buffer{}

	config := &LogConfig{
		Level:  LogLevelDebug,
		Output: buffer,
	}

	New("custom_service", config)

	defer func() {
		if r := recover(); r != nil {
			assert.Contains(t, r.(string), "Fatal message", "Panic message should contain expected message")
		}
	}()
	Fatal("Fatal message", "key", "value")
}
