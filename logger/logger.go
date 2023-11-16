package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

type (
	// Logger
	Logger struct {
		zerolog.Logger
	}
)

// New - create new logger instance.
func New(opts ...option) *Logger {

	// Get default ctor options & modify it.
	var options = getDefaultOptions()
	for _, set := range opts {
		set(options)
	}

	// Setup globals.
	zerolog.TimeFieldFormat = options.tsFormat
	zerolog.TimestampFieldName = options.tsName

	// Convert to console writer default outputs.
	if options.format != FormatJSON {
		for i, out := range options.output {
			switch out.(type) {
			case *os.File:
				options.output[i] = &zerolog.ConsoleWriter{
					Out:        out,
					NoColor:    options.format != FormatConsole,
					TimeFormat: options.tsFormat,
				}
			}
		}
	}

	// Get zerolog logger parameters.
	var (
		logger zerolog.Logger
		level  = GetZerologLevel(options.level)
		writer io.Writer
	)

	switch len(options.output) {
	case 0:
		writer = io.Discard
	case 1:
		writer = options.output[0]
	default:
		writer = zerolog.MultiLevelWriter(options.output...)
	}

	logger = zerolog.New(writer).With().Timestamp().Logger().Level(level)

	return &Logger{
		Logger: logger,
	}
}
