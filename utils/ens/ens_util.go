package ens

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"golang.org/x/net/idna"
	"strings"
)

// Uts46ToAscii is to convert input into Ascii version follow IDN UTS46 spec.
// Eg: bücher.com => xn--bcher-kva.com
func Uts46ToAscii(input string) (string, error) {
	output, err := idna.Lookup.ToASCII(input)
	if err != nil {
		return "", err
	}
	if strings.Contains(input, "..") {
		return "", fmt.Errorf("bad input")
	}
	return output, nil
}

// Uts46ToUnicode is to convert input into Unicode version follow IDN UTS46 spec.
// Eg: xn--bcher-kva.com => bücher.com
func Uts46ToUnicode(input string) (string, error) {
	output, err := idna.Lookup.ToUnicode(input)
	if err != nil {
		output = ""
	}
	return output, err
}

// ValidateEnsDomain to check whether input domain is valid or not.
// Also check if it has correct top-level-domain part
func ValidateEnsDomain(domain, tld string) error {
	domain = strings.TrimSpace(domain)
	tld = strings.ToLower(strings.TrimSpace(tld))
	if len(tld) > 0 {
		if len(domain) < len(tld)+1 /*dot*/ +3 /*at least 3 characters in first label*/ {
			return fmt.Errorf("too short")
		}
	} else {
		if len(domain) < 1 /*tld*/ +1 /*dot*/ +3 /*at least 3 characters in first label*/ {
			return fmt.Errorf("too short")
		}
	}

	domain, err := Uts46ToAscii(domain)

	if err == nil {
		if domain[0] == '.' {
			return fmt.Errorf("can not begins with a dot")
		}

		if len(tld) > 0 {
			if !strings.HasSuffix(strings.ToLower(domain), "."+tld) {
				return fmt.Errorf("not tld %s", tld)
			}
		}

		if strings.HasSuffix(strings.ToLower(domain), ".addr.reverse") {
			return fmt.Errorf("this probably a reverse addr")
		}

		lastDotIdx := strings.LastIndex(domain, ".")
		if lastDotIdx < 0 {
			return fmt.Errorf("must contains at least one dot")
		} else if lastDotIdx == len(domain)-1 {
			return fmt.Errorf("must not ends with a dot")
		}

		parts := strings.Split(domain, ".")
		domainName := parts[len(parts)-2]
		if len(domainName) < 3 {
			return fmt.Errorf("bad domain")
		}
	}

	return err
}

// NameHash returns keccak256 value of input domain (to be used to procedure a node).
// This implementation follow javascript version
func NameHash(labelOrDomain string) (string, error) {
	node := make([]byte, 32)

	if labelOrDomain != "" {
		labels := strings.Split(labelOrDomain, ".")

		for i := len(labels) - 1; i >= 0; i-- {
			normalisedLabel, err := Uts46ToUnicode(labels[i])
			if err != nil {
				return "", errors.Wrap(err, "failed to normalize label")
			}
			var labelSha = crypto.Keccak256Hash([]byte(normalisedLabel))
			node = crypto.Keccak256Hash(node, labelSha.Bytes()).Bytes()
		}
	}

	return "0x" + hex.EncodeToString(node), nil
}

// IsValidEnsNode returns true if input node is valid lowercase hex 32 byte string with 0x as prefix
func IsValidEnsNode(node string) bool {
	if len(node) != 66 {
		return false
	}

	if node[:2] != "0x" {
		return false
	}

	for _, c := range node[2:] {
		if c >= '0' && c <= '9' {
			continue
		}
		if c >= 'a' && c <= 'f' {
			continue
		}
		// only accept 0-9 and a-f
		return false
	}

	return true
}
