package locker

import (
	"context"
	"errors"
	"time"

	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
	tkcontext "github.com/tarusov/toolkit/context"
)

type (
	// Locker struct.
	Locker struct {
		cl         redis.UniversalClient
		retryCount int
		retryDelay time.Duration
	}

	// UnlockFn is func to remove lock.
	UnlockFn func() error

	// lock struct.
	lock struct {
		*redislock.Lock
		key string
		ctx context.Context
	}
)

// Aux error types.
var (
	ErrLockNotObtained = errors.New("failed to obtain lock")       // Unable to obtain lock.
	ErrNotLocked       = errors.New("failed to unlock; not exist") // No lock exist.
)

// New creates a new redis-locker instance.
func New(cl redis.UniversalClient, opts ...option) *Locker {

	var l = &Locker{
		cl:         cl,
		retryCount: 5,
		retryDelay: 5 * time.Second,
	}

	for _, opt := range opts {
		opt(l)
	}

	return l
}

// Lock by key.
func (l *Locker) Lock(ctx context.Context, key string, ttl time.Duration) (UnlockFn, error) {

	obtained, err := redislock.Obtain(
		ctx,
		l.cl,
		key,
		ttl,
		&redislock.Options{
			RetryStrategy: redislock.LimitRetry(redislock.LinearBackoff(l.retryDelay), l.retryCount),
		},
	)
	if err != nil {
		if err == redislock.ErrNotObtained {
			return nil, ErrLockNotObtained
		}
		return nil, err
	}

	tkcontext.GetLogger(ctx).WithField("key", key).Debug("lock obtained")

	r := lock{
		Lock: obtained,
		key:  key,
		ctx:  ctx,
	}

	return func() error {
		return r.Unlock()
	}, nil
}

// Unlock method try to release current mutex.
func (l *lock) Unlock() (err error) {
	defer func() {
		tkcontext.GetLogger(l.ctx).WithField("key", l.key).WithError(err).Debug("unlocked")
	}()

	err = l.Lock.Release(l.ctx)
	if err == redislock.ErrLockNotHeld {
		return ErrNotLocked
	}

	return err
}
