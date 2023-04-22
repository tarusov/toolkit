package context

import (
	"context"

	"github.com/tarusov/toolkit/logger"
)

// special type for logger context storage.
type loggerContextKey struct{}

// GetLogger extract logger from context.
func GetLogger(ctx context.Context) *logger.Logger {
	if ctx != nil {
		if v, ok := ctx.Value(loggerContextKey{}).(*logger.Logger); ok {
			return v
		}
	}
	return logger.New()
}

// WithLogger puts logger into context.
func WithLogger(ctx context.Context, l *logger.Logger) context.Context {
	return context.WithValue(ctx, loggerContextKey{}, l)
}
