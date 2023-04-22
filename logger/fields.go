package logger

type (
	// Fields is logging fields type.
	Fields map[string]any
)

// WithError - create copy of logger with error field if err != nil
func (l *Logger) WithError(err error) *Logger {
	if err == nil {
		return l
	}
	var zl = l.Logger.With().Err(err).Logger()
	return &Logger{Logger: zl}
}

// WithField - create a copy of logger with target field.
func (l *Logger) WithField(name string, value any) *Logger {
	var k, v = kv(name, value)
	var zl = l.Logger.With().Interface(k, v).Logger()
	return &Logger{Logger: zl}
}

// WithFields - create a copy of logger with target fields.
func (l *Logger) WithFields(fields Fields) *Logger {
	var zl = l.Logger
	for ink, inv := range fields {
		var k, v = kv(ink, inv)
		zl = zl.With().Interface(k, v).Logger()
	}
	return &Logger{Logger: zl}
}

// kv check key-value pair.
func kv(k string, v any) (string, any) {
	if k == "" {
		k = "unknown_field"
	}
	return k, v
}
