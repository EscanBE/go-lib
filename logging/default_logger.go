package logging

import (
	"fmt"
	logtypes "github.com/EscanBE/go-lib/logging/types"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	_ Logger = &defaultLogger{}
)

// defaultLogger represents the default logger for any kind of error
type defaultLogger struct {
	Logger zerolog.Logger
}

// NewDefaultLogger builds a new defaultLogger instance
func NewDefaultLogger() Logger {
	result := &defaultLogger{
		Logger: log.Logger,
	}
	result.SetLogLevel(logtypes.LOG_LEVEL_DEFAULT)
	return result
}

// SetLogLevel implements Logger
func (d *defaultLogger) SetLogLevel(level string) error {
	logLvl, err := zerolog.ParseLevel(level)
	if err != nil {
		return err
	}

	zerolog.SetGlobalLevel(logLvl)
	return nil
}

// SetLogFormat implements Logger
func (d *defaultLogger) SetLogFormat(format string) error {
	switch format {
	case logtypes.LOG_FORMAT_JSON:
		// JSON is the default logging format
		break

	case logtypes.LOG_FORMAT_TEXT:
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		break

	default:
		return fmt.Errorf("invalid logging format: %s", format)
	}

	return nil
}

// Info implements Logger
func (d *defaultLogger) Info(msg string, keyVals ...interface{}) {
	d.Logger.Info().Fields(getLogFields(keyVals...)).Msg(msg)
}

// Debug implements Logger
func (d *defaultLogger) Debug(msg string, keyVals ...interface{}) {
	d.Logger.Debug().Fields(getLogFields(keyVals...)).Msg(msg)
}

// Error implements Logger
func (d *defaultLogger) Error(msg string, keyVals ...interface{}) {
	d.Logger.Error().Fields(getLogFields(keyVals...)).Msg(msg)
}

// ApplyConfig implements Logger
func (d *defaultLogger) ApplyConfig(config logtypes.LoggingConfig) error {
	if len(config.Level) > 0 {
		err := d.SetLogLevel(config.Level)
		if err != nil {
			return err
		}
	}
	if len(config.Format) > 0 {
		err := d.SetLogFormat(config.Format)
		if err != nil {
			return err
		}
	}
	return nil
}

func getLogFields(keyVals ...interface{}) map[string]interface{} {
	if len(keyVals) < 1 || len(keyVals)%2 != 0 {
		return nil
	}

	fields := make(map[string]interface{})
	for i := 0; i < len(keyVals); i += 2 {
		fields[keyVals[i].(string)] = keyVals[i+1]
	}

	return fields
}
