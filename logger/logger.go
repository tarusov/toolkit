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

// map of formats.
var formats = map[Format]struct{}{
	FormatConsole: {},
	FormatJSON:    {},
	FormatText:    {},
}

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

// map of levels.
var levels = map[Level]zerolog.Level{
	LevelDebug: zerolog.DebugLevel,
	LevelInfo:  zerolog.InfoLevel,
	LevelWarn:  zerolog.WarnLevel,
	LevelError: zerolog.ErrorLevel,
	LevelFatal: zerolog.FatalLevel,
}

// supportColors is map of output with color support.
var supportColors = map[*os.File]struct{}{
	os.Stdout: {},
	os.Stderr: {},
}

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

	if _, ok := formats[p.format]; !ok {
		printWarn("unknown format '%s'. json format is set", p.format)
		p.format = FormatJSON
	}

	var lvl, ok = levels[p.level]
	if !ok {
		printWarn("unknown logging level '%s'. debug level is set", p.level)
		lvl = zerolog.DebugLevel
	}

	if p.tsEnabled && p.tsFieldName == "" {
		printWarn("timestamp name is empty. 'time' value is set")
		p.tsFieldName = "time"
	}

	if p.tsEnabled && p.tsFormat == "" {
		printWarn("time format is empty. time.RFC3339 is set")
		p.tsFormat = time.RFC3339
	}

	if p.format != FormatJSON {
		for i, out := range p.outputs {
			if v, ok := out.(*os.File); ok {

				var colorful = false
				if p.format == FormatConsole {
					_, colorful = supportColors[v]
				}

				p.outputs[i] = zerolog.ConsoleWriter{
					Out:        out,
					NoColor:    !colorful,
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

// printWarn write warning while logger create.
func printWarn(format string, args ...any) {
	var tmp = zerolog.New(os.Stderr).Level(zerolog.WarnLevel).With().Logger()
	tmp.Warn().Msgf(format, args...)
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
