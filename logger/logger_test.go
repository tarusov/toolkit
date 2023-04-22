package logger_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/tarusov/toolkit/logger"
)

const (
	MessageText = "message"

	TestFileName = "dear_diary.log"
)

func TestLoggerIvalidFormatAndLevel(t *testing.T) {

	var buf = bytes.NewBuffer(nil)

	var l = logger.New(
		logger.WithOutput(buf),
		logger.WithFormat(logger.Format("everything")),
		logger.WithLevel(logger.Level("everything")),
		logger.WithTimestampEnabled(false),
	)

	l.Debug(MessageText)

	if buf.String() != `{"level":"debug","message":"message"}`+"\n" {
		t.Errorf("invalid_logger_level: unexpected logging behaviour")
	}
}

func TestLoggerIvalidTimestampParams(t *testing.T) {

	var buf = bytes.NewBuffer(nil)

	var l = logger.New(
		logger.WithOutput(buf),
		logger.WithTimestampEnabled(true),
		logger.WithTimestampFormat(""),
		logger.WithTimestampName(""),
	)

	l.Info(MessageText)

	if !strings.Contains(buf.String(), "time") {
		t.Errorf("invalid_timestamp_name: unexpected logging behaviour")
	}
}

func TestLoggerLevel(t *testing.T) {

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
				l.Debug(MessageText)

			case logger.LevelInfo:
				l.Info(MessageText)

			case logger.LevelWarn:
				l.Warn(MessageText)

			case logger.LevelError:
				l.Error(MessageText)
			}

			// TODO: remove
			fmt.Print(buf.String())

			if c.Output != "" {
				c.Output += "\n"
			}

			if buf.String() != c.Output {
				t.Errorf("%s: unexpected logging behaviour", c.Name)
			}
		})
	}
}

func TestLoggerDebugf(t *testing.T) {

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
		Output:       `{"level":"debug","message":"text: message"}`,
	}, {
		Name:         "json_info_ok",
		Format:       logger.FormatJSON,
		LoggerLevel:  logger.LevelDebug,
		MessageLevel: logger.LevelInfo,
		Output:       `{"level":"info","message":"text: message"}`,
	}, {
		Name:         "json_warn_ok",
		Format:       logger.FormatJSON,
		LoggerLevel:  logger.LevelDebug,
		MessageLevel: logger.LevelWarn,
		Output:       `{"level":"warn","message":"text: message"}`,
	}, {
		Name:         "json_error_ok",
		Format:       logger.FormatJSON,
		LoggerLevel:  logger.LevelDebug,
		MessageLevel: logger.LevelError,
		Output:       `{"level":"error","message":"text: message"}`,
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
				l.Debugf("text: %s", MessageText)

			case logger.LevelInfo:
				l.Infof("text: %s", MessageText)

			case logger.LevelWarn:
				l.Warnf("text: %s", MessageText)

			case logger.LevelError:
				l.Errorf("text: %s", MessageText)
			}

			// TODO: remove
			fmt.Print(buf.String())

			if c.Output != "" {
				c.Output += "\n"
			}

			if buf.String() != c.Output {
				t.Errorf("%s: unexpected logging behaviour", c.Name)
			}
		})
	}
}

func TestWriteToFile(t *testing.T) {

	f, err := logger.OpenFile(TestFileName)
	if err != nil {
		t.Fatal("failed to open file to write: ", err)
	}

	var l = logger.New(
		logger.WithOutput(f, os.Stdout),
		logger.WithFormat(logger.FormatConsole),
	)

	l.Error(MessageText)
	_ = f.Close()

	raw, err := os.ReadFile(TestFileName)
	if err != nil && err != io.EOF {
		t.Fatal("failed to open file to read: ", err)
	}

	if !strings.Contains(string(raw), "ERR message") {
		t.Error("write_to_file: unexpected logger behaviour")
	}

	_ = os.Remove(TestFileName)
}
