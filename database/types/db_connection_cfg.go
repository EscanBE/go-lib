package types

import "fmt"

// PostgresDatabaseConfig holds configuration needed to connect to the postgres database
type PostgresDatabaseConfig struct {
	Host                   string `mapstructure:"host"`
	Port                   uint16 `mapstructure:"port"`
	Username               string `mapstructure:"username"`
	Password               string `mapstructure:"password"`
	Name                   string `mapstructure:"name"`
	Schema                 string `mapstructure:"schema,omitempty"`
	EnableSsl              bool   `mapstructure:"enable-ssl,omitempty"`
	MaxOpenConnectionCount int16  `mapstructure:"max-open-connection-count"`
	MaxIdleConnectionCount int16  `mapstructure:"max-idle-connection-count"`
}

func (c PostgresDatabaseConfig) Validate() error {
	if len(c.Host) < 1 {
		return fmt.Errorf("missing database's host")
	}
	if c.Port < 1 {
		return fmt.Errorf("missing database's port")
	}
	if len(c.Username) < 1 {
		return fmt.Errorf("missing database's username")
	}
	if len(c.Password) < 1 {
		return fmt.Errorf("missing database's password")
	}
	if len(c.Name) < 1 {
		return fmt.Errorf("missing database name")
	}
	if c.MaxOpenConnectionCount < 1 && c.MaxOpenConnectionCount != -1 {
		return fmt.Errorf("invalid database's max-open-connection-count")
	}
	if c.MaxIdleConnectionCount < 1 && c.MaxIdleConnectionCount != -1 {
		return fmt.Errorf("invalid database's max-idle-connection-count")
	}
	return nil
}
