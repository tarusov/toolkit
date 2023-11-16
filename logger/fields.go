package logger

// Fields is aux type for logging fields.
type Fields map[string]any

// WithField create copy of logger with custom field.
func (l *Logger) WithField(k string, v any) *Logger {
	zl := l.Logger.With().Interface(k, v).Logger()
	return &Logger{Logger: zl}
}

// WithError create copy of logger with error.
// If err is nil return logger without changes.
func (l *Logger) WithError(err error) *Logger {
	if err != nil {
		zl := l.Logger.With().Err(err).Logger()
		return &Logger{Logger: zl}
	}
	return l
}

// WithFields create copy of logger with custom fields.
func (l *Logger) WithFields(fields Fields) *Logger {
	zl := l.Logger
	for k, v := range fields {
		zl = zl.With().Interface(k, v).Logger()
	}
	return &Logger{Logger: zl}
}
