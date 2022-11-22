package types

import (
	"fmt"
	"github.com/EscanBE/go-lib/test_utils"
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
		wantErrMsgContains string
	}{
		{
			name:   "success",
			level:  LOG_LEVEL_DEFAULT,
			format: LOG_FORMAT_DEFAULT,
		},
		{
			name:   "empty ok",
			level:  "",
			format: "",
		},
		{
			name:               "invalid level",
			level:              LOG_LEVEL_DEFAULT + "-invalid",
			format:             LOG_FORMAT_DEFAULT,
			wantErrMsgContains: "level",
		},
		{
			name:               "invalid format",
			level:              LOG_LEVEL_DEFAULT,
			format:             LOG_FORMAT_DEFAULT + "-invalid",
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
			wantErr := len(tt.wantErrMsgContains) > 0
			if (err != nil) != wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, wantErr)
				return
			}
			if !test_utils.WantErrorContainsStringIfNonEmptyOtherWiseNoError(t, err, tt.wantErrMsgContains) {
				return
			}
		})
	}
}
