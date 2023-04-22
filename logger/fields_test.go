package logger_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/tarusov/toolkit/logger"
)

func TestLoggerFields(t *testing.T) {

	var buf = bytes.NewBuffer(nil)

	var l = logger.New(
		logger.WithOutput(buf),
		logger.WithTimestampEnabled(false),
	)

	l = l.WithFields(logger.Fields{
		"": "unknown_value",
	})

	l = l.WithField("k", "v")
	l = l.WithError(errors.New("err_message"))

	l.Info(MessageText)

	if buf.String() != `{"level":"info","unknown_field":"unknown_value","k":"v","error":"err_message","message":"message"}`+"\n" {
		t.Fatal("test_logger_fields: unexpected logger behaviour")
	}
}

func TestSkipNilErr(t *testing.T) {

	var buf = bytes.NewBuffer(nil)

	var l = logger.New(
		logger.WithOutput(buf),
		logger.WithTimestampEnabled(false),
	)

	l = l.WithError(nil)

	l.Info(MessageText)

	if buf.String() != `{"level":"info","message":"message"}`+"\n" {
		t.Fatal("test_logger_fields: unexpected logger behaviour")
	}
}
