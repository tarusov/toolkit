package sentry

import (
	"time"

	"github.com/tarusov/toolkit/logger"
)

type (
	// params is sentry ctor parameters.
	params struct {
		level      logger.Level
		timeout    time.Duration
		serverName string
		env        string
		dist       string
		release    string
		stackTrace bool
	}

	// option
	option func(*params)
)

// WithTimeout setup sentry flush timeout. Default is 5 seconds.
func WithTimeout(t time.Duration) option {
	return func(p *params) {
		p.timeout = t
	}
}

// WithLevel setup sentry severnity level. Default is error.
func WithLevel(l logger.Level) option {
	return func(p *params) {
		p.level = l
	}
}

// WithServerName setup server name.
func WithServerName(sn string) option {
	return func(p *params) {
		p.serverName = sn
	}
}

// WithEnv setup target env name.
func WithEnv(env string) option {
	return func(p *params) {
		p.env = env
	}
}

// WithDist setup target dist name.
func WithDist(dist string) option {
	return func(p *params) {
		p.dist = dist
	}
}

// WithRelease setup target release name.
func WithRelease(release string) option {
	return func(p *params) {
		p.release = release
	}
}

// WithStackTraceEnabled setup stacktrace logging. Default is true.
func WithStackTraceEnabled(e bool) option {
	return func(p *params) {
		p.stackTrace = e
	}
}
