package logger

import "fmt"

// Fatal
func (l *Logger) Fatal(v ...any) {
	l.Logger.Fatal().Msg(fmt.Sprint(v...))
}

// Fatalf
func (l *Logger) Fatalf(format string, v ...any) {
	l.Logger.Fatal().Msgf(format, v...)
}

// Fatalln
func (l *Logger) Fatalln(v ...any) {
	l.Logger.Fatal().Msg(fmt.Sprint(v...))
}

// Panic
func (l *Logger) Panic(v ...any) {
	l.Logger.Panic().Msg(fmt.Sprint(v...))
}

// Panicf
func (l *Logger) Panicf(format string, v ...any) {
	l.Logger.Panic().Msgf(format, v...)
}

// Panicln
func (l *Logger) Panicln(v ...any) {
	l.Logger.Panic().Msg(fmt.Sprint(v...))
}

// Print
func (l *Logger) Print(v ...any) {
	l.Logger.Print(v...)
}

// Printf
func (l *Logger) Printf(format string, v ...any) {
	l.Logger.Printf(format, v...)
}

// Println
func (l *Logger) Println(v ...any) {
	l.Logger.Print(v...)
}

// Debug
func (l *Logger) Debug(v ...any) {
	l.Logger.Debug().Msg(fmt.Sprint(v...))
}

// Debugf
func (l *Logger) Debugf(format string, v ...any) {
	l.Logger.Debug().Msgf(format, v...)
}

// Info
func (l *Logger) Info(v ...any) {
	l.Logger.Info().Msg(fmt.Sprint(v...))
}

// Infof
func (l *Logger) Infof(format string, v ...any) {
	l.Logger.Info().Msgf(format, v...)
}

// Warn
func (l *Logger) Warn(v ...any) {
	l.Logger.Warn().Msg(fmt.Sprint(v...))
}

// Warnf
func (l *Logger) Warnf(format string, v ...any) {
	l.Logger.Warn().Msgf(format, v...)
}

// Error
func (l *Logger) Error(v ...any) {
	l.Logger.Error().Msg(fmt.Sprint(v...))
}

// Errorf
func (l *Logger) Errorf(format string, v ...any) {
	l.Logger.Error().Msgf(format, v...)
}
