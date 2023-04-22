package logger

type (
	// Fields is logging fields type.
	Fields map[string]any
)

// WithField - create a copy of logger with target field.
func (l *Logger) WithField(name string, value any) *Logger {
	var zl = l.Logger.With().Interface(name, value).Logger()
	return &Logger{Logger: zl}
}

// WithFields - create a copy of logger with target fields.
func (l *Logger) WithFields(fields Fields) *Logger {
	var zl = l.Logger
	for k, v := range fields {
		zl = zl.With().Interface(k, v).Logger()
	}
	return &Logger{Logger: zl}
}
