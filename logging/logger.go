package logging

import logtypes "github.com/EscanBE/go-lib/logging/types"

// Logger defines a function that takes an error and logs it.
type Logger interface {
	// SetLogLevel changes the log level, valid values are: info, error and debug
	SetLogLevel(level string) error

	// SetLogFormat changes the log format, valid values are: json and text
	SetLogFormat(format string) error

	// Info does log the input message at level Info
	Info(msg string, keyvals ...interface{})

	// Debug does log the input message at level Debug
	Debug(msg string, keyvals ...interface{})

	// Error does log the input message at level Error
	Error(msg string, keyvals ...interface{})

	// ApplyConfig applies provided configuration
	ApplyConfig(config logtypes.LoggingConfig) error
}
