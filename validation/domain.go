package validation

import (
	"golang.org/x/net/idna"
)

// IsValidHostname returns true if input hostname is a valid domain name.
// with google.com => true,
// with localhost => true,
// with https://google.com => false,
// with google.com:443 => false,
func IsValidHostname(hostname string) bool {
	h, err := idna.Lookup.ToASCII(hostname)
	return err == nil && len(h) > 0
}
