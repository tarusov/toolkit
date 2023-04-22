package sentry

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
	"github.com/tarusov/toolkit/logger"
)

type (
	// SentryNotifier is logger webhook for sentry.
	SentryNotifier struct {
		hub        *sentry.Hub
		level      zerolog.Level
		timeout    time.Duration
		stackTrace bool
	}
)

// mapping zerolog to sentry levels.
var zerologToSentryLevels = map[zerolog.Level]sentry.Level{
	zerolog.DebugLevel: sentry.LevelDebug,
	zerolog.InfoLevel:  sentry.LevelInfo,
	zerolog.WarnLevel:  sentry.LevelWarning,
	zerolog.ErrorLevel: sentry.LevelError,
	zerolog.FatalLevel: sentry.LevelFatal,
}

// New create new sentry notifier entry.
func New(dsn string, opts ...option) (*SentryNotifier, error) {

	if dsn == "" {
		return nil, errors.New("sentry dsn is empty")
	}

	// map of levels.
	var levels = map[logger.Level]zerolog.Level{
		logger.LevelDebug: zerolog.DebugLevel,
		logger.LevelInfo:  zerolog.InfoLevel,
		logger.LevelWarn:  zerolog.WarnLevel,
		logger.LevelError: zerolog.ErrorLevel,
		logger.LevelFatal: zerolog.FatalLevel,
	}

	var p = &params{
		level:      logger.LevelError,
		timeout:    time.Second * 5,
		stackTrace: true,
	}

	for _, opt := range opts {
		opt(p)
	}

	lvl, ok := levels[p.level]
	if !ok {
		return nil, fmt.Errorf("failed to create sentry client: unknown severnity level")
	}

	if p.timeout == 0 {
		return nil, fmt.Errorf("failed to create sentry client: flush timeout is null")
	}

	var client, err = sentry.NewClient(sentry.ClientOptions{
		Dsn:              dsn,
		Debug:            p.level == logger.LevelDebug,
		AttachStacktrace: p.stackTrace,
		ServerName:       p.serverName,
		Environment:      p.env,
		Release:          p.release,
		DebugWriter:      zerolog.New(os.Stderr).With().Str("pkg", "sentry").Logger(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create sentry client: %w", err)
	}

	var scope = sentry.NewScope()
	var hub = sentry.NewHub(client, scope)

	return &SentryNotifier{
		hub:        hub,
		level:      lvl,
		timeout:    p.timeout,
		stackTrace: p.stackTrace,
	}, nil
}

// Write is implemens io.Writer Write method.
func (sn *SentryNotifier) Write(p []byte) (int, error) {
	sn.capture(string(p), sn.level, nil)
	return len(p), nil
}

// WriteLevel is implements zerolog.LevelWriter method.
func (sn *SentryNotifier) WriteLevel(zl zerolog.Level, p []byte) (int, error) {

	var count = len(p)

	if zl < sn.level {
		return count, nil
	}

	var extra logger.Fields
	var err = json.Unmarshal(p, &extra)
	if err != nil {
		return 0, fmt.Errorf("unmarshal message: %w", err)
	}

	var msgField, msg = sn.getErrMessage(extra)

	delete(extra, msgField)
	delete(extra, zerolog.MessageFieldName)
	delete(extra, zerolog.ErrorFieldName)
	delete(extra, zerolog.LevelFieldName)
	delete(extra, zerolog.TimestampFieldName)

	sn.capture(msg, zl, extra)

	return count, nil
}

// capture sends final message to server.
func (sn *SentryNotifier) capture(msg string, level zerolog.Level, extra logger.Fields) {

	var e = sentry.NewEvent()
	var stacktrace *sentry.Stacktrace
	if sn.stackTrace {
		stacktrace = sn.getStacktrace(msg)
	}

	e.Message = msg
	e.Level = zerologToSentryLevels[level]
	e.Timestamp = time.Now()
	e.Exception = []sentry.Exception{{
		Value:      msg,
		Stacktrace: stacktrace,
	}}

	_ = sn.hub.CaptureEvent(e)
}

// getErrMessage from error extra data.
func (sn *SentryNotifier) getErrMessage(extra logger.Fields) (field string, msg string) {
	var ok bool
	if msg, ok = extra["error"].(string); ok {
		return "error", msg
	}

	if msg, ok = extra["message"].(string); ok {
		return "message", msg
	}

	return "message", "undefined error"
}

// getStacktrace extract stacktrace from error.
func (sn *SentryNotifier) getStacktrace(msg string) *sentry.Stacktrace {

	var stacktrace = sentry.ExtractStacktrace(errors.New(msg))
	if stacktrace == nil {
		stacktrace = sentry.NewStacktrace()
	}
	if stacktrace == nil {
		return nil
	}

	var frames = make([]sentry.Frame, 0, len(stacktrace.Frames))
	for _, frame := range stacktrace.Frames {
		// Skip tracing into logger files.
		if strings.HasPrefix(frame.Module, "github.com/rs/zerolog") ||
			strings.HasSuffix(frame.Filename, "logger.go") ||
			strings.HasSuffix(frame.Filename, "sentry.go") {
			continue
		}

		frames = append(frames, frame)
	}

	stacktrace.Frames = frames

	return stacktrace
}

// Close implements io.Close method for sentry notifier.
func (sn *SentryNotifier) Close() error {
	if !sn.hub.Flush(sn.timeout) {
		return errors.New("sentry flush timeout")
	}
	return nil
}
