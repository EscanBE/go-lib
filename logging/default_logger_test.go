package logging

import (
	logtypes "github.com/EscanBE/go-lib/logging/types"
	"github.com/EscanBE/go-lib/test_utils"
	"github.com/rs/zerolog"
	"reflect"
	"testing"
)

func TestNewDefaultLogger(t *testing.T) {
	t.Run("init with level default", func(t *testing.T) {
		_ = NewDefaultLogger()
		level, _ := zerolog.ParseLevel(logtypes.LOG_LEVEL_DEFAULT)
		if zerolog.GlobalLevel().String() != level.String() {
			t.Errorf("NewDefaultLogger() wrong defaul level, got %s, want %s", zerolog.GlobalLevel().String(), logtypes.LOG_LEVEL_DEFAULT)
		}
	})
}

func Test_defaultLogger_SetLogLevel(t *testing.T) {
	logger := NewDefaultLogger()
	tests := []struct {
		name    string
		level   string
		wantErr bool
	}{
		{
			name:    "default",
			level:   logtypes.LOG_LEVEL_DEFAULT,
			wantErr: false,
		},
		{
			name:    "empty",
			level:   "",
			wantErr: false,
		},
		{
			name:    "debug",
			level:   logtypes.LOG_LEVEL_DEBUG,
			wantErr: false,
		},
		{
			name:    "info",
			level:   logtypes.LOG_LEVEL_INFO,
			wantErr: false,
		},
		{
			name:    "error",
			level:   logtypes.LOG_LEVEL_ERROR,
			wantErr: false,
		},
		{
			name:    "invalid",
			level:   logtypes.LOG_LEVEL_DEFAULT + "-invalid",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := logger.SetLogLevel(tt.level)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetLogLevel() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				level, _ := zerolog.ParseLevel(tt.level)
				if zerolog.GlobalLevel().String() != level.String() {
					t.Errorf("SetLogLevel() set wrong level, got %s, want %s", zerolog.GlobalLevel().String(), tt.level)
				}
			}
		})
	}
}

func Test_defaultLogger_SetLogFormat(t *testing.T) {
	logger := NewDefaultLogger()
	tests := []struct {
		name    string
		format  string
		wantErr bool
	}{
		{
			name:    "default",
			format:  logtypes.LOG_FORMAT_DEFAULT,
			wantErr: false,
		},
		{
			name:    "empty",
			format:  "",
			wantErr: true,
		},
		{
			name:    "json",
			format:  logtypes.LOG_FORMAT_JSON,
			wantErr: false,
		},
		{
			name:    "text",
			format:  logtypes.LOG_FORMAT_TEXT,
			wantErr: false,
		},
		{
			name:    "invalid",
			format:  logtypes.LOG_FORMAT_DEFAULT + "-invalid",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := logger.SetLogFormat(tt.format)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetLogFormat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_defaultLogger_Info_Debug_Error(t *testing.T) {
	t.Run("log success", func(t *testing.T) {
		defer test_utils.DeferWantNoPanic(t)

		logger := NewDefaultLogger()
		logger.Info("info")
		logger.Debug("debug")
		logger.Error("error")
		logger.Info("info", "k", "v")
		logger.Debug("debug", "k", "v")
		logger.Error("error", "k", "v")
	})
}

func Test_defaultLogger_ApplyConfig(t *testing.T) {
	tests := []struct {
		name               string
		level              string
		format             string
		wantErrMsgContains string
	}{
		{
			name:   "success",
			level:  logtypes.LOG_LEVEL_DEFAULT,
			format: logtypes.LOG_FORMAT_DEFAULT,
		},
		{
			name:   "empty ok",
			level:  "",
			format: "",
		},
		{
			name:               "wrong level",
			level:              logtypes.LOG_LEVEL_DEFAULT + "-invalid",
			format:             logtypes.LOG_FORMAT_DEFAULT,
			wantErrMsgContains: "level",
		},
		{
			name:               "wrong format",
			level:              logtypes.LOG_LEVEL_DEFAULT,
			format:             logtypes.LOG_FORMAT_DEFAULT + "-invalid",
			wantErrMsgContains: "format",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := NewDefaultLogger()
			err := logger.ApplyConfig(logtypes.LoggingConfig{
				Level:  tt.level,
				Format: tt.format,
			})
			wantErr := len(tt.wantErrMsgContains) > 0
			if (err != nil) != wantErr {
				t.Errorf("ApplyConfig() error = %v, wantErr %v", err, wantErr)
				return
			}
			if !test_utils.WantErrorContainsStringIfNonEmptyOtherWiseNoError(t, err, tt.wantErrMsgContains) {
				return
			}
		})
	}
}

func Test_getLogFields(t *testing.T) {
	tests := []struct {
		name      string
		input     []interface{}
		want      map[string]interface{}
		wantPanic bool
	}{
		{
			name:  "success",
			input: []interface{}{"k", "v"},
			want:  map[string]interface{}{"k": "v"},
		},
		{
			name:      "key must be string",
			input:     []interface{}{666, 999},
			wantPanic: true,
		},
		{
			name:  "empty success return nil",
			input: nil,
			want:  nil,
		},
		{
			name:      "panic when number of args is not an even number",
			input:     []interface{}{"1"},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer test_utils.DeferWantPanicDepends(t, tt.wantPanic)

			if got := getLogFields(tt.input...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLogFields() = %v, want %v", got, tt.want)
			}
		})
	}
}
