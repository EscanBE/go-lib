package types

import (
	"github.com/EscanBE/go-lib/test_utils"
	"testing"
)

func TestPostgresDatabaseConfig_Validate(t *testing.T) {
	type fields struct {
		Host                   string
		Port                   uint16
		Username               string
		Password               string
		Name                   string
		Schema                 string
		EnableSsl              bool
		MaxOpenConnectionCount int16
		MaxIdleConnectionCount int16
	}
	tests := []struct {
		fields     fields
		wantErrMsg string
	}{
		{
			fields: fields{
				Host:                   "127.0.0.1",
				Port:                   5432,
				Username:               "postgres",
				Password:               "postgres",
				Name:                   "postgres",
				Schema:                 "public",
				EnableSsl:              false,
				MaxOpenConnectionCount: 20,
				MaxIdleConnectionCount: 20,
			},
			wantErrMsg: "",
		},
		{
			fields: fields{
				Host:                   "",
				Port:                   5432,
				Username:               "postgres",
				Password:               "postgres",
				Name:                   "postgres",
				Schema:                 "public",
				EnableSsl:              false,
				MaxOpenConnectionCount: 20,
				MaxIdleConnectionCount: 20,
			},
			wantErrMsg: "host",
		},
		{
			fields: fields{
				Host:                   "127.0.0.1",
				Port:                   0,
				Username:               "postgres",
				Password:               "postgres",
				Name:                   "postgres",
				Schema:                 "public",
				EnableSsl:              false,
				MaxOpenConnectionCount: 20,
				MaxIdleConnectionCount: 20,
			},
			wantErrMsg: "port",
		},
		{
			fields: fields{
				Host:                   "127.0.0.1",
				Port:                   5432,
				Username:               "",
				Password:               "postgres",
				Name:                   "postgres",
				Schema:                 "public",
				EnableSsl:              false,
				MaxOpenConnectionCount: 20,
				MaxIdleConnectionCount: 20,
			},
			wantErrMsg: "username",
		},
		{
			fields: fields{
				Host:                   "127.0.0.1",
				Port:                   5432,
				Username:               "postgres",
				Password:               "",
				Name:                   "postgres",
				Schema:                 "public",
				EnableSsl:              false,
				MaxOpenConnectionCount: 20,
				MaxIdleConnectionCount: 20,
			},
			wantErrMsg: "password",
		},
		{
			fields: fields{
				Host:                   "127.0.0.1",
				Port:                   5432,
				Username:               "postgres",
				Password:               "postgres",
				Name:                   "",
				Schema:                 "public",
				EnableSsl:              false,
				MaxOpenConnectionCount: 20,
				MaxIdleConnectionCount: 20,
			},
			wantErrMsg: "database name",
		},
		{
			fields: fields{
				Host:                   "127.0.0.1",
				Port:                   5432,
				Username:               "postgres",
				Password:               "postgres",
				Name:                   "postgres",
				Schema:                 "",
				EnableSsl:              false,
				MaxOpenConnectionCount: 20,
				MaxIdleConnectionCount: 20,
			},
		},
		{
			fields: fields{
				Host:                   "127.0.0.1",
				Port:                   5432,
				Username:               "postgres",
				Password:               "postgres",
				Name:                   "postgres",
				Schema:                 "public",
				EnableSsl:              false,
				MaxOpenConnectionCount: -2,
				MaxIdleConnectionCount: 20,
			},
			wantErrMsg: "max-open",
		},
		{
			fields: fields{
				Host:                   "127.0.0.1",
				Port:                   5432,
				Username:               "postgres",
				Password:               "postgres",
				Name:                   "postgres",
				Schema:                 "public",
				EnableSsl:              false,
				MaxOpenConnectionCount: 20,
				MaxIdleConnectionCount: -2,
			},
			wantErrMsg: "max-idle",
		},
	}
	for _, tt := range tests {
		t.Run(tt.wantErrMsg, func(t *testing.T) {
			c := PostgresDatabaseConfig{
				Host:                   tt.fields.Host,
				Port:                   tt.fields.Port,
				Username:               tt.fields.Username,
				Password:               tt.fields.Password,
				Name:                   tt.fields.Name,
				Schema:                 tt.fields.Schema,
				EnableSsl:              tt.fields.EnableSsl,
				MaxOpenConnectionCount: tt.fields.MaxOpenConnectionCount,
				MaxIdleConnectionCount: tt.fields.MaxIdleConnectionCount,
			}
			err := c.Validate()
			wantErr := len(tt.wantErrMsg) > 0
			if (err != nil) != wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, wantErr)
				return
			}
			if !test_utils.WantErrorContainsStringIfNonEmptyOtherWiseNoError(t, err, tt.wantErrMsg) {
				return
			}
		})
	}
}
