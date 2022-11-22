package utils

import (
	"net"
	"strconv"
	"time"
)

// CheckPortIsTcpOpen returns if port is not open for TCP mode
func CheckPortIsTcpOpen(host string, port int, timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, strconv.Itoa(port)), timeout)
	if err != nil {
		return err
	}

	_ = conn.Close()
	return nil
}
