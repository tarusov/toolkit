package logger_test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/tarusov/toolkit/logger"
)

func TestLogger_NoOutputs(t *testing.T) {

	var (
		log = logger.New(
			logger.WithOutput(),
		)
	)

	log.Printf(testFmt, testMsg)
}

func TestLogger_MultipleOutput(t *testing.T) {

	var (
		buf = bytes.NewBuffer(nil)
		log = logger.New(
			logger.WithFormat(logger.FormatConsole),
			logger.WithLevel(logger.LevelDebug),
			logger.WithOutput(
				buf,
				os.Stdout,
				os.Stderr,
				io.Discard,
			),
		)
	)

	log.Printf(testFmt, testMsg)
	if len(buf.String()) == 0 {
		t.Error("TestLogger_MultipleOutput: unexpected behaviour")
	}
}

func TestLogger_TimestampFormat(t *testing.T) {

	var (
		ts  = time.Now().Format(time.DateOnly)
		buf = bytes.NewBuffer(nil)
		log = logger.New(
			logger.WithFormat(logger.FormatJSON),
			logger.WithLevel(logger.LevelDebug),
			logger.WithOutput(buf),
			logger.WithTimestampFormat(time.DateOnly),
			logger.WithTimestampName("Ts"),
		)
	)

	log.Printf(testFmt, testMsg)

	if !strings.Contains(buf.String(), ts) {
		t.Error("TestLogger_TimestampFormat: unexpected behaviour")
	}

	if !strings.Contains(buf.String(), "Ts") {
		t.Error("TestLogger_TimestampFormat: unexpected behaviour")
	}

}
