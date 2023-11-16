package logger

import "github.com/rs/zerolog"

// Level is logger severnity level.
type Level string

// List of supported logger levels.
const (
	LevelDebug   Level = "debug"
	LevelInfo    Level = "info"
	LevelWarning Level = "warn"
	LevelError   Level = "error"
	LevelFatal   Level = "fatal"
	LevelPanic   Level = "panic"
)

// getZerologLevel convert level name to zerolog.Level.
func GetZerologLevel(l Level) zerolog.Level {
	switch l {
	case LevelDebug:
		return zerolog.DebugLevel
	case LevelInfo:
		return zerolog.InfoLevel
	case LevelWarning:
		return zerolog.WarnLevel
	case LevelError:
		return zerolog.ErrorLevel
	case LevelFatal:
		return zerolog.FatalLevel
	case LevelPanic:
		return zerolog.PanicLevel
	default:
		return zerolog.NoLevel
	}
}
