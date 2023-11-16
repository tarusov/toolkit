package logger

// Format is logging output format.
type Format string

// List of supported logging formats.
const (
	FormatConsole Format = "console" // Colored text
	FormatJSON    Format = "json"    // JSON output
	FormatText    Format = "text"    // Plain text output
)
