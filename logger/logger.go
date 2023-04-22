package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// Format is logging output format.
type Format string

// List of possible logging format.
const (
	FormatConsole Format = "console"
	FormatJSON    Format = "json"
	FormatText    Format = "text"
)

// Level is logging severnity level.
type Level string

// List of possible logging levels.
const (
	LevelDebug Level = "debug"
	LevelInfo  Level = "info"
	LevelWarn  Level = "warn"
	LevelError Level = "error"
	LevelFatal Level = "fatal"
)

type (
	// Logger struct
	Logger struct {
		zerolog.Logger
	}
)

// New create new logger instance.
func New(opts ...option) *Logger {

	var p = &params{
		outputs:     []io.Writer{os.Stderr},
		format:      FormatJSON,
		level:       LevelInfo,
		tsFieldName: "time",
		tsFormat:    time.RFC3339,
	}

	for _, opt := range opts {
		opt(p)
	}

	var lvl, err = zerolog.ParseLevel(string(p.level))
	if err != nil {
		// TODO: Add message
		lvl = zerolog.DebugLevel
	}

	if p.format != FormatJSON {
		for i, out := range p.outputs {
			if _, ok := out.(*os.File); ok {
				p.outputs[i] = zerolog.ConsoleWriter{
					Out:        out,
					NoColor:    (p.format == FormatText),
					TimeFormat: p.tsFormat,
				}
			}
		}
	}

	zerolog.TimestampFieldName = p.tsFieldName
	zerolog.TimeFieldFormat = p.tsFormat

	var (
		wr = zerolog.MultiLevelWriter(p.outputs...)
		zl = zerolog.New(wr).With().Timestamp().Logger().Level(lvl)
	)

	return &Logger{
		Logger: zl,
	}
}
