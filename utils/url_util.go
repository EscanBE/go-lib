package utils

//goland:noinspection SpellCheckingInspection
import (
	"fmt"
	neturl "net/url"
	"strings"
)

// ExtractHostAndPort return the host name (domain) with port number from input uri
func ExtractHostAndPort(url string) (string, error) {
	host, err := neturl.Parse(url)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s", strings.TrimPrefix(host.Hostname(), "www."), host.Port()), nil
}

// ExtractHostAndPortOrKeep return the host name (domain) with port number from input uri. Return original input url if failed to parse
func ExtractHostAndPortOrKeep(url string) string {
	result, err := ExtractHostAndPort(url)
	if err != nil {
		return url
	} else {
		return result
	}
}
