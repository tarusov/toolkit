package logger_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"

	"github.com/tarusov/toolkit/logger"
)

func TestLogger_WithField(t *testing.T) {

	var (
		buf = bytes.NewBuffer(nil)
		log = logger.New(
			logger.WithLevel(logger.LevelDebug),
			logger.WithFormat(logger.FormatJSON),
			logger.WithOutput(buf),
		)

		obj = struct {
			K string `json:"k"`
		}{}
	)

	log = log.WithField("k", "v")
	log.Printf(testFmt, testMsg)

	err := json.Unmarshal(buf.Bytes(), &obj)
	if err != nil {
		t.Errorf("TestLogger_WithField: %v", err)
	}

	if obj.K != "v" {
		t.Error("TestLogger_WithField: unexpected behaviour")
	}
}

func TestLogger_WithFields(t *testing.T) {

	var (
		buf = bytes.NewBuffer(nil)
		log = logger.New(
			logger.WithLevel(logger.LevelDebug),
			logger.WithFormat(logger.FormatJSON),
			logger.WithOutput(buf),
		)

		obj = struct {
			K string `json:"k"`
			X int    `json:"x"`
		}{}
	)

	log = log.WithFields(logger.Fields{
		"k": "v",
		"x": 100,
	})
	log.Printf(testFmt, testMsg)

	err := json.Unmarshal(buf.Bytes(), &obj)
	if err != nil {
		t.Errorf("TestLogger_WithFields: %v", err)
	}

	if obj.K != "v" || obj.X != 100 {
		t.Error("TestLogger_WithFields: unexpected behaviour")
	}
}

func TestLogger_WithError(t *testing.T) {

	var (
		buf = bytes.NewBuffer(nil)
		log = logger.New(
			logger.WithLevel(logger.LevelDebug),
			logger.WithFormat(logger.FormatJSON),
			logger.WithOutput(buf),
		)

		obj = struct {
			ErrMsg string `json:"error"`
		}{}
	)

	log = log.WithError(errors.New("errmsg"))
	log.Printf(testFmt, testMsg)

	err := json.Unmarshal(buf.Bytes(), &obj)
	if err != nil {
		t.Errorf("TestLogger_WithField: %v", err)
	}

	if obj.ErrMsg != "errmsg" {
		t.Error("TestLogger_WithError: unexpected behaviour")
	}
}
