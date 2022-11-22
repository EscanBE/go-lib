package types

import (
	"strings"
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
		wantErr    bool
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
			wantErr:    false,
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
			wantErr:    true,
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
			wantErr:    true,
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
			wantErr:    true,
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
			wantErr:    true,
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
			wantErr:    true,
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
			wantErr: false,
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
			wantErr:    true,
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
			wantErr:    true,
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
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && !strings.Contains(err.Error(), tt.wantErrMsg) {
				t.Errorf("Validate() expected error message [%s] must contains [%s]", err.Error(), tt.wantErrMsg)
			}
		})
	}
}
