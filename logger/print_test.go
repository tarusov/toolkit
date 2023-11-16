package logger_test

import (
	"bytes"
	"testing"

	"github.com/tarusov/toolkit/logger"
)

const (
	testFmt = "%s"
	testMsg = "foo? bar!"
)

func TestLogger_Print(t *testing.T) {

	tests := []struct {
		name     string
		logLevel logger.Level
		msgLevel logger.Level
		empty    bool
	}{{
		name:     "debug_debug_ok",
		logLevel: logger.LevelDebug,
		msgLevel: logger.LevelDebug,
	}, {
		name:     "debug_info_ok",
		logLevel: logger.LevelDebug,
		msgLevel: logger.LevelInfo,
	}, {
		name:     "debug_warn_ok",
		logLevel: logger.LevelDebug,
		msgLevel: logger.LevelWarning,
	}, {
		name:     "debug_error_ok",
		logLevel: logger.LevelDebug,
		msgLevel: logger.LevelError,
	}, {
		name:     "info_debug_ok",
		logLevel: logger.LevelInfo,
		msgLevel: logger.LevelDebug,
		empty:    true,
	}, {
		name:     "info_info_ok",
		logLevel: logger.LevelInfo,
		msgLevel: logger.LevelInfo,
	}, {
		name:     "info_warn_ok",
		logLevel: logger.LevelInfo,
		msgLevel: logger.LevelWarning,
	}, {
		name:     "info_error_ok",
		logLevel: logger.LevelInfo,
		msgLevel: logger.LevelError,
	}, {
		name:     "warn_debug_ok",
		logLevel: logger.LevelWarning,
		msgLevel: logger.LevelDebug,
		empty:    true,
	}, {
		name:     "warn_info_ok",
		logLevel: logger.LevelWarning,
		msgLevel: logger.LevelInfo,
		empty:    true,
	}, {
		name:     "warn_warn_ok",
		logLevel: logger.LevelWarning,
		msgLevel: logger.LevelWarning,
	}, {
		name:     "warn_error_ok",
		logLevel: logger.LevelWarning,
		msgLevel: logger.LevelError,
	}, {
		name:     "error_debug_ok",
		logLevel: logger.LevelError,
		msgLevel: logger.LevelDebug,
		empty:    true,
	}, {
		name:     "error_info_ok",
		logLevel: logger.LevelError,
		msgLevel: logger.LevelInfo,
		empty:    true,
	}, {
		name:     "error_warn_ok",
		logLevel: logger.LevelError,
		msgLevel: logger.LevelWarning,
		empty:    true,
	}, {
		name:     "error_error_ok",
		logLevel: logger.LevelError,
		msgLevel: logger.LevelError,
	}, {
		name:     "debug_print_ok",
		logLevel: logger.LevelDebug,
	}}

	for _, test := range tests {
		t.Run(
			test.name,
			func(t *testing.T) {

				var (
					buf = bytes.NewBuffer(nil)
					log = logger.New(
						logger.WithLevel(test.logLevel),
						logger.WithOutput(buf),
					)
				)
				switch test.msgLevel {
				case logger.LevelDebug:
					log.Debug(testMsg)
				case logger.LevelInfo:
					log.Info(testMsg)
				case logger.LevelWarning:
					log.Warn(testMsg)
				case logger.LevelError:
					log.Error(testMsg)
				case logger.LevelFatal:
					log.Fatal(testMsg)
				case logger.LevelPanic:
					log.Panic(testMsg)
				default:
					log.Print(testMsg)
				}

				if len(buf.String()) > 0 && test.empty {
					t.Errorf("%s: unexpected behaviour", test.name)
				}
			},
		)
	}
}

func TestLogger_Printf(t *testing.T) {

	tests := []struct {
		name     string
		logLevel logger.Level
		msgLevel logger.Level
		empty    bool
	}{{
		name:     "debug_debug_ok",
		logLevel: logger.LevelDebug,
		msgLevel: logger.LevelDebug,
	}, {
		name:     "debug_info_ok",
		logLevel: logger.LevelDebug,
		msgLevel: logger.LevelInfo,
	}, {
		name:     "debug_warn_ok",
		logLevel: logger.LevelDebug,
		msgLevel: logger.LevelWarning,
	}, {
		name:     "debug_error_ok",
		logLevel: logger.LevelDebug,
		msgLevel: logger.LevelError,
	}, {
		name:     "info_debug_ok",
		logLevel: logger.LevelInfo,
		msgLevel: logger.LevelDebug,
		empty:    true,
	}, {
		name:     "info_info_ok",
		logLevel: logger.LevelInfo,
		msgLevel: logger.LevelInfo,
	}, {
		name:     "info_warn_ok",
		logLevel: logger.LevelInfo,
		msgLevel: logger.LevelWarning,
	}, {
		name:     "info_error_ok",
		logLevel: logger.LevelInfo,
		msgLevel: logger.LevelError,
	}, {
		name:     "warn_debug_ok",
		logLevel: logger.LevelWarning,
		msgLevel: logger.LevelDebug,
		empty:    true,
	}, {
		name:     "warn_info_ok",
		logLevel: logger.LevelWarning,
		msgLevel: logger.LevelInfo,
		empty:    true,
	}, {
		name:     "warn_warn_ok",
		logLevel: logger.LevelWarning,
		msgLevel: logger.LevelWarning,
	}, {
		name:     "warn_error_ok",
		logLevel: logger.LevelWarning,
		msgLevel: logger.LevelError,
	}, {
		name:     "error_debug_ok",
		logLevel: logger.LevelError,
		msgLevel: logger.LevelDebug,
		empty:    true,
	}, {
		name:     "error_info_ok",
		logLevel: logger.LevelError,
		msgLevel: logger.LevelInfo,
		empty:    true,
	}, {
		name:     "error_warn_ok",
		logLevel: logger.LevelError,
		msgLevel: logger.LevelWarning,
		empty:    true,
	}, {
		name:     "error_error_ok",
		logLevel: logger.LevelError,
		msgLevel: logger.LevelError,
	}, {
		name:     "debug_print_ok",
		logLevel: logger.LevelDebug,
	}}

	for _, test := range tests {
		t.Run(
			test.name,
			func(t *testing.T) {

				var (
					buf = bytes.NewBuffer(nil)
					log = logger.New(
						logger.WithLevel(test.logLevel),
						logger.WithOutput(buf),
					)
				)
				switch test.msgLevel {
				case logger.LevelDebug:
					log.Debugf(testFmt, testMsg)
				case logger.LevelInfo:
					log.Infof(testFmt, testMsg)
				case logger.LevelWarning:
					log.Warnf(testFmt, testMsg)
				case logger.LevelError:
					log.Errorf(testFmt, testMsg)
				case logger.LevelFatal:
					log.Fatalf(testFmt, testMsg)
				case logger.LevelPanic:
					log.Panicf(testFmt, testMsg)
				default:
					log.Printf(testFmt, testMsg)
				}

				if len(buf.String()) > 0 && test.empty {
					t.Errorf("%s: unexpected behaviour", test.name)
				}
			},
		)
	}
}
