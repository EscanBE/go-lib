package types

import (
	"fmt"
	"strings"
	"testing"
)

func TestLoggingConfig_Validate(t *testing.T) {
	for _, level := range []string{LOG_LEVEL_DEBUG, LOG_LEVEL_INFO, LOG_LEVEL_ERROR} {
		for _, format := range []string{LOG_FORMAT_JSON, LOG_FORMAT_TEXT} {
			t.Run(fmt.Sprintf("L=%s-F=%s", level, format), func(t *testing.T) {
				c := LoggingConfig{
					Level:  level,
					Format: format,
				}
				err := c.Validate()
				if err != nil {
					t.Errorf("Validate() error = %v, do not want err", err)
				}
			})
		}
	}

	tests := []struct {
		name               string
		level              string
		format             string
		wantErr            bool
		wantErrMsgContains string
	}{
		{
			name:    "success",
			level:   LOG_LEVEL_DEFAULT,
			format:  LOG_FORMAT_DEFAULT,
			wantErr: false,
		},
		{
			name:    "empty ok",
			level:   "",
			format:  "",
			wantErr: false,
		},
		{
			name:               "invalid level",
			level:              LOG_LEVEL_DEFAULT + "-invalid",
			format:             LOG_FORMAT_DEFAULT,
			wantErr:            true,
			wantErrMsgContains: "level",
		},
		{
			name:               "invalid format",
			level:              LOG_LEVEL_DEFAULT,
			format:             LOG_FORMAT_DEFAULT + "-invalid",
			wantErr:            true,
			wantErrMsgContains: "format",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := LoggingConfig{
				Level:  tt.level,
				Format: tt.format,
			}
			err := c.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				if len(tt.wantErrMsgContains) < 1 {
					t.Errorf("setup test wrongly. Partial err msg is required")
				}
				if !strings.Contains(err.Error(), tt.wantErrMsgContains) {
					t.Errorf("Validate() error = %s, want contains %s", err.Error(), tt.wantErrMsgContains)
				}
			}
		})
	}
}
