package logger_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/tarusov/toolkit/logger"
)

func TestLoggerLevel(t *testing.T) {

	const MessageText = "message"

	tc := []struct {
		Name         string
		Format       logger.Format
		LoggerLevel  logger.Level
		MessageLevel logger.Level
		Output       string
	}{{
		Name:         "json_debug_ok",
		Format:       logger.FormatJSON,
		LoggerLevel:  logger.LevelDebug,
		MessageLevel: logger.LevelDebug,
		Output:       `{"level":"debug","message":"message"}`,
	}, {
		Name:         "json_info_ok",
		Format:       logger.FormatJSON,
		LoggerLevel:  logger.LevelDebug,
		MessageLevel: logger.LevelInfo,
		Output:       `{"level":"info","message":"message"}`,
	}, {
		Name:         "json_warn_ok",
		Format:       logger.FormatJSON,
		LoggerLevel:  logger.LevelDebug,
		MessageLevel: logger.LevelWarn,
		Output:       `{"level":"warn","message":"message"}`,
	}, {
		Name:         "json_error_ok",
		Format:       logger.FormatJSON,
		LoggerLevel:  logger.LevelDebug,
		MessageLevel: logger.LevelError,
		Output:       `{"level":"error","message":"message"}`,
	}, {
		Name:         "json_debug_no_msg_ok",
		Format:       logger.FormatJSON,
		LoggerLevel:  logger.LevelFatal,
		MessageLevel: logger.LevelDebug,
	}, {
		Name:         "json_info_no_msg_ok",
		Format:       logger.FormatJSON,
		LoggerLevel:  logger.LevelFatal,
		MessageLevel: logger.LevelInfo,
	}, {
		Name:         "json_warn_no_msg_ok",
		Format:       logger.FormatJSON,
		LoggerLevel:  logger.LevelFatal,
		MessageLevel: logger.LevelWarn,
	}, {
		Name:         "json_error_no_msg_ok",
		Format:       logger.FormatJSON,
		LoggerLevel:  logger.LevelFatal,
		MessageLevel: logger.LevelError,
	}}

	for _, c := range tc {
		t.Run(c.Name, func(t *testing.T) {

			buf := bytes.NewBuffer(nil)

			l := logger.New(
				logger.WithOutput(buf),
				logger.WithFormat(c.Format),
				logger.WithLevel(c.LoggerLevel),
				logger.WithTimestampEnabled(false),
			)

			switch c.MessageLevel {
			case logger.LevelDebug:
				l.Debug(string(MessageText))

			case logger.LevelInfo:
				l.Info(string(MessageText))

			case logger.LevelWarn:
				l.Warn(string(MessageText))

			case logger.LevelError:
				l.Error(string(MessageText))
			}

			// TODO: remove
			fmt.Println(buf.String())

			if c.Output != "" && !strings.Contains(buf.String(), c.Output) {
				t.Errorf("%s: unexpected logging behaviour", c.Name)
			}

		})
	}
}
