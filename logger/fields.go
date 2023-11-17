package logger

// Fields is aux type for logging fields.
type Fields map[string]any

// WithField create copy of logger with custom field.
func (l *Logger) WithField(k string, v any) *Logger {
	return &Logger{
		Logger: l.Logger.With().Interface(k, v).Logger(),
	}
}

// WithError create copy of logger with error.
// If err is nil return logger without changes.
func (l *Logger) WithError(err error) *Logger {
	if err != nil {
		return &Logger{
			Logger: l.Logger.With().Err(err).Logger(),
		}
	}
	return l
}

// WithFields create copy of logger with custom fields.
func (l *Logger) WithFields(fields Fields) *Logger {
	var ctx = l.Logger.With()
	for k, v := range fields {
		ctx = ctx.Interface(k, v)
	}
	return &Logger{Logger: ctx.Logger()}
}
