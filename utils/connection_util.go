package utils

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

// CheckPortIsTcpOpen returns if port is not open for TCP mode
func CheckPortIsTcpOpen(host string, port int, timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, strconv.Itoa(port)), timeout)
	if err != nil {
		if conn != nil {
			//goland:noinspection GoUnhandledErrorResult
			defer conn.Close()
		}
		return err
	} else if conn == nil {
		return fmt.Errorf("enable to open a connection")
	}

	//goland:noinspection GoUnhandledErrorResult
	defer conn.Close()
	return nil
}
