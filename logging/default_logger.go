package logging

import (
	"fmt"
	logtypes "github.com/EscanBE/go-lib/logging/types"
	"os"
	"time"

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
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().UTC()
	}
	result := &defaultLogger{
		Logger: log.Logger,
	}
	_ = result.SetLogLevel(logtypes.LOG_LEVEL_DEFAULT)
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
	validationErr := config.Validate()
	if validationErr != nil {
		return validationErr
	}
	if len(config.Level) > 0 {
		_ = d.SetLogLevel(config.Level) // shouldn't err, validated before
	}
	if len(config.Format) > 0 {
		_ = d.SetLogFormat(config.Format) // shouldn't err, validated before
	}
	return nil
}

func getLogFields(keyVals ...interface{}) map[string]interface{} {
	if len(keyVals) < 1 {
		return nil
	}

	if len(keyVals)%2 != 0 {
		panic(fmt.Errorf("number of argument should be even"))
	}

	fields := make(map[string]interface{})
	for i := 0; i < len(keyVals); i += 2 {
		fields[keyVals[i].(string)] = keyVals[i+1]
	}

	return fields
}
