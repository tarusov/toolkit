package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type (
	// options is list of settings for logger ctor.
	options struct {
		format   Format
		level    Level
		output   []io.Writer
		tsFormat string
		tsName   string
	}

	// Option if type of modifing func for ctor.
	Option func(*options)
)

// getDefaultOptions return default settings for logger.
func getDefaultOptions() *options {
	return &options{
		format:   FormatJSON,
		level:    LevelInfo,
		output:   []io.Writer{os.Stderr},
		tsFormat: time.RFC3339,
		tsName:   zerolog.TimestampFieldName,
	}
}

// WithFormat option setup logging output format.
// Default format is "json"
func WithFormat(f Format) Option {
	return func(o *options) {
		o.format = f
	}
}

// WithLevel option setup custom logging level.
// Default level is "info".
func WithLevel(l Level) Option {
	return func(o *options) {
		o.level = l
	}
}

// WithOutput option setup custom output target.
// Default is os.Stderr.
func WithOutput(w ...io.Writer) Option {
	return func(o *options) {
		o.output = w
	}
}

// WithTimestampFormat option setup custom timestamp format.
// Default is time.RFC3339
func WithTimestampFormat(f string) Option {
	return func(o *options) {
		o.tsFormat = f
	}
}

// WithTimestampName option setup custom timestamp name.
// Default is "time".
func WithTimestampName(n string) Option {
	return func(o *options) {
		o.tsName = n
	}
}
