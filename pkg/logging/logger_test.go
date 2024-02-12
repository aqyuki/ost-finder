package logging

import (
	"context"
	"log/slog"
	"testing"
)

func Test_NewLogger(t *testing.T) {
	t.Run("NewLoggerFromEnv should return a logger in develop mode", func(t *testing.T) {
		t.Setenv(EnvLoggingMode, ModeDevelopment)
		actual := NewLoggerFromEnv()
		if actual == nil {
			t.Error("NewLoggerFromEnv should return a logger but received nil")
		}
	})

	t.Run("NewLoggerFromEnv should return a logger in production mode", func(t *testing.T) {
		t.Setenv(EnvLoggingMode, "")
		actual := NewLoggerFromEnv()
		if actual == nil {
			t.Error("NewLoggerFromEnv should return a logger but received nil")
		}
	})

	t.Run("NewLoggerWithConfig should return a logger", func(t *testing.T) {
		actual := NewLoggerWithConfig(true, "debug")
		if actual == nil {
			t.Error("NewLoggerWithConfig should return a logger but received nil")
		}
	})
}

func Test_convertLoggingLevel(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		level    string
		excepted slog.Level
	}{
		{
			name:     "should return debug level",
			level:    LevelDebug,
			excepted: slog.LevelDebug,
		},
		{
			name:     "should return info level",
			level:    LevelInfo,
			excepted: slog.LevelInfo,
		},
		{
			name:     "should return warn level",
			level:    LevelWarn,
			excepted: slog.LevelWarn,
		},
		{
			name:     "should return error level",
			level:    LevelError,
			excepted: slog.LevelError,
		},
		{
			name:     "should return info level when unknown strings",
			level:    "",
			excepted: slog.LevelInfo,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := convertLoggingLevel(tt.level)
			if actual != tt.excepted {
				t.Errorf("convertLoggingLevel should return %s but received %s", tt.excepted, actual)
			}
		})
	}
}

func Test_Context(t *testing.T) {
	t.Parallel()

	t.Run("Context should return a logger which is same as given logger", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		logger := NewLoggerWithConfig(true, LevelDebug)

		actual := WithLogger(ctx, logger)
		if actual == nil {
			t.Errorf("WithLogger should return a context but received nil")
		}

		if UnwrapContext(actual) != logger {
			t.Errorf("UnwrapContext should return a logger which is same as given logger")
		}
	})

	t.Run("Context should returns a default logger", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		actual := UnwrapContext(ctx)
		if actual == nil {
			t.Errorf("UnwrapContext should return a logger but received nil")
		} else if actual != DefaultLogger() {
			t.Errorf("UnwrapContext should return a default logger")
		}
	})
}
