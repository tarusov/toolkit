package logger

import "io"

type (
	// option is ctor modificator func.
	option func(*params)

	// params is ctor parameters.
	params struct {
		outputs     []io.Writer
		format      Format
		level       Level
		tsEnabled   bool
		tsFieldName string // timestamp field name
		tsFormat    string // time format
	}
)

// WithOutput setup target logging output. Default is os.Stderr
func WithOutput(o ...io.Writer) option {
	return func(p *params) {
		p.outputs = o
	}
}

// WithFormat setup format for logging. Default is JSON.
func WithFormat(f Format) option {
	return func(p *params) {
		p.format = f
	}
}

// WithLevel setup severnity level of logging. Default is "info".
func WithLevel(l Level) option {
	return func(p *params) {
		p.level = l
	}
}

// WithTimestampEnabled setup ts including. Default is true.
func WithTimeStampEnabled(e bool) option {
	return func(p *params) {
		p.tsEnabled = e
	}
}

// WithTimeStampFormat setup format for timestamp. Default is RFC3339.
func WithTimeStampFormat(f string) option {
	return func(p *params) {
		p.tsFormat = f
	}
}

// WithTimeStampName setup timestamp field name. Default is "time".
func WithTimeStampName(n string) option {
	return func(p *params) {
		p.tsFieldName = n
	}
}
