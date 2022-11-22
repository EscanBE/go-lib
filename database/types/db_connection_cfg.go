package types

import (
	"fmt"
	"github.com/EscanBE/go-lib/utils"
)

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

// Validate will return an error if any configuration problem
func (c PostgresDatabaseConfig) Validate() error {
	if utils.IsBlank(c.Host) {
		return fmt.Errorf("missing database host")
	}
	if c.Port == 0 {
		return fmt.Errorf("missing database port")
	}
	if c.Port < 1 || c.Port > 65535 {
		return fmt.Errorf("invalid database port :%d", c.Port)
	}
	if utils.IsBlank(c.Username) {
		return fmt.Errorf("missing database username")
	}
	if utils.IsBlank(c.Password) {
		return fmt.Errorf("missing database password")
	}
	if utils.IsBlank(c.Name) {
		return fmt.Errorf("missing database name")
	}
	if c.MaxOpenConnectionCount < 1 && c.MaxOpenConnectionCount != -1 {
		return fmt.Errorf("invalid database max-open-connection-count")
	}
	if c.MaxIdleConnectionCount < 1 && c.MaxIdleConnectionCount != -1 {
		return fmt.Errorf("invalid database max-idle-connection-count")
	}
	return nil
}
