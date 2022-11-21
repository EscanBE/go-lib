package utils

//goland:noinspection SpellCheckingInspection
import (
	"fmt"
	neturl "net/url"
	"strings"
)

// ExtractHostAndPort return the host name (domain) with port number from input uri
func ExtractHostAndPort(rawUrl string) (string, error) {
	url, err := neturl.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	host := strings.TrimPrefix(url.Hostname(), "www.")
	port := url.Port()
	if IsBlank(host) {
		return "", fmt.Errorf("unable to extract host")
	}
	if IsBlank(port) {
		return host, nil
	}
	return fmt.Sprintf("%s:%s", host, port), nil
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
