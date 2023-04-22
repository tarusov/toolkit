package locker

import "time"

// option func is locker ctor modifier func.
type option func(*Locker)

// WithRetryCount set retry count for locker.
func WithRetryCount(c int) option {
	return func(l *Locker) {
		l.retryCount = c
	}
}

// WithRetryDelay set pause between retries.
func WithRetryDelay(d time.Duration) option {
	return func(l *Locker) {
		l.retryDelay = d
	}
}
