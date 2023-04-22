package logger

import (
	"fmt"
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
		tsEnabled:   true,
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
		zl = zerolog.New(wr).Level(lvl)
	)

	if p.tsEnabled {
		zl = zl.With().Timestamp().Logger()
	}

	return &Logger{
		Logger: zl,
	}
}

// Debug
func (l *Logger) Debug(msg string) {
	l.Logger.Debug().Msg(msg)
}

// Info
func (l *Logger) Info(msg string) {
	l.Logger.Info().Msg(msg)
}

// Warn
func (l *Logger) Warn(msg string) {
	l.Logger.Warn().Msg(msg)
}

// Error
func (l *Logger) Error(msg string) {
	l.Logger.Error().Msg(msg)
}

// Fatal
func (l *Logger) Fatal(msg string) {
	l.Logger.Fatal().Msg(msg)
}

// Debugf
func (l *Logger) Debugf(format string, args ...any) {
	l.Logger.Debug().Msg(fmt.Sprintf(format, args...))
}

// Info
func (l *Logger) Infof(format string, args ...any) {
	l.Logger.Info().Msg(fmt.Sprintf(format, args...))
}

// Warn
func (l *Logger) Warnf(format string, args ...any) {
	l.Logger.Warn().Msg(fmt.Sprintf(format, args...))
}

// Error
func (l *Logger) Errorf(format string, args ...any) {
	l.Logger.Error().Msg(fmt.Sprintf(format, args...))
}

// Fatal
func (l *Logger) Fatalf(format string, args ...any) {
	l.Logger.Fatal().Msg(fmt.Sprintf(format, args...))
}
