package toolkit

import (
	"io"
	"log"
	"os"

	"github.com/tarusov/toolkit/config"
	"github.com/tarusov/toolkit/logger"
	"github.com/tarusov/toolkit/logger/sentry"
)

// NewLogger create new logger based on env defaults.
func NewLogger() *logger.Logger {

	type (
		cfg struct {
			Format      logger.Format `env:"LOGGER_FORMAT" envDefault:"json"`
			Level       logger.Level  `env:"LOGGER_LEVEL"  envDefault:"info"`
			DSN         string        `env:"SENTRY_DSN"`
			Environment string        `env:"SENTRY_ENV"     envDefault:"staging"`
			Release     string        `env:"SENTRY_RELEASE" envDefault:"0.0.1"`
		}
	)

	// Load cfg from ENV.
	var c cfg
	err := config.Load(&c)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var out = []io.Writer{os.Stderr}

	// If Sentry DSN is defined
	// add sentry to output list.
	if c.DSN != "" {
		s, err := sentry.New(
			c.DSN,
			sentry.WithEnvironment(c.Environment),
			sentry.WithRelease(c.Release),
		)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		out = append(out, s)
	}

	// Create new logger with target parameters.
	return logger.New(
		logger.WithFormat(c.Format),
		logger.WithLevel(c.Level),
		logger.WithOutput(out...),
	)
}
