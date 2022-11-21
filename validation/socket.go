package validation

import "net"

// IsValidIP returns true if the provided IP is a valid IP address (IPv4 + IPv6)
func IsValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

// IsValidPort returns true if the port is in-range 1 to 65535
func IsValidPort(port int) bool {
	return port >= 1 && port <= 65535
}
