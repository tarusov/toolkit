package sentry

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
	"github.com/tarusov/toolkit/logger"
)

type (
	// SentryNotifier
	SentryNotifier struct {
		client *sentry.Client
		level  zerolog.Level
	}
)

// New create new sentry notifier instance.
func New(dsn string, opts ...option) (*SentryNotifier, error) {

	// Get default ctor options & modify it.
	var options = getDefaultOptions()
	for _, set := range opts {
		set(options)
	}

	client, err := sentry.NewClient(sentry.ClientOptions{
		Dsn:              dsn,
		Debug:            options.level == logger.LevelDebug,
		AttachStacktrace: options.stacktrace,
	})
	if err != nil {
		return nil, fmt.Errorf("new sentry client: %w", err)
	}

	return &SentryNotifier{
		client: client,
		level:  logger.GetZerologLevel(options.level),
	}, nil
}

// Write implements io.Writer method for use in logger.
func (s *SentryNotifier) Write(p []byte) (int, error) {
	return 0, nil
}

// WriteLevel implements zerolog.LevelWriter for use in logger.
func (s *SentryNotifier) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {

	var lp = len(p)

	if s.level > level {
		return lp, nil
	}

	return 0, nil
}
