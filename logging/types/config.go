package types

import "fmt"

// LoggingConfig will be used to apply logging config, it also provides utilities
type LoggingConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

// Validate performs validation on the LoggingConfig instance
func (c LoggingConfig) Validate() error {
	if len(c.Level) < 1 {
		// OK
	} else if c.Level == LOG_LEVEL_DEBUG {
		// OK
	} else if c.Level == LOG_LEVEL_INFO {
		// OK
	} else if c.Level == LOG_LEVEL_ERROR {
		// OK
	} else {
		return fmt.Errorf("invalid log level %s", c.Level)
	}

	if len(c.Format) < 1 {
		// OK
	} else if c.Format == LOG_FORMAT_TEXT {
		// OK
	} else if c.Format == LOG_FORMAT_JSON {
		// OK
	} else {
		return fmt.Errorf("invalid log format %s", c.Format)
	}

	return nil
}
