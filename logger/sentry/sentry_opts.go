package sentry

import (
	"time"

	"github.com/tarusov/toolkit/logger"
)

type (
	// options is list of settings for sentry ctor.
	options struct {
		environment string
		release     string
		level       logger.Level
		timeout     time.Duration
	}

	// option if type of modifing func for ctor.
	option func(*options)
)

// getDefaultOptions return default settings for logger.
func getDefaultOptions() *options {
	return &options{
		environment: "staging",
		release:     "1",
		level:       logger.LevelError,
		timeout:     time.Second * 3,
	}
}

// WithEnvironment setup target env name.
// Default is "staging".
func WithEnvironment(e string) option {
	return func(o *options) {
		o.environment = e
	}
}

// WithRelease setup target release version.
// Default is "1".
func WithRelease(r string) option {
	return func(o *options) {
		o.release = r
	}
}

// WithLevel option setup custom logging level.
// Default level is "error".
func WithLevel(l logger.Level) option {
	return func(o *options) {
		o.level = l
	}
}

// WithTimeout option setup custom flush timeout.
// Default timeout is 3s.
func WithTimeout(t time.Duration) option {
	return func(o *options) {
		o.timeout = t
	}
}
