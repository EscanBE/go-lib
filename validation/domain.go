package validation

import (
	"golang.org/x/net/idna"
	"strings"
)

// IsValidHostname returns true if input hostname is localhost or valid domain name
func IsValidHostname(hostname string) bool {
	if strings.EqualFold(hostname, "localhost") {
		return true
	}
	h, err := idna.Lookup.ToASCII(hostname)
	return err == nil && len(h) > 0
}
