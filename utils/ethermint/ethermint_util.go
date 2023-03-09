package ethermint

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

var (
	regexChainID          = `[a-z]{1,}`
	regexEIP155Separator  = `_{1}`
	regexEIP155           = `[1-9][0-9]*`
	regexEpochSeparator   = `-{1}`
	regexEpoch            = `[1-9][0-9]*`
	regexEthermintChainID = regexp.MustCompile(fmt.Sprintf(`^(%s)%s(%s)%s(%s)$`,
		regexChainID,
		regexEIP155Separator,
		regexEIP155,
		regexEpochSeparator,
		regexEpoch))
)

// IsValidEthermintChainID returns false if the given chain identifier is incorrectly formatted.
//
// This function was cloned 1-1 from ethermint lib because go-lib does not link with ethermint lib.
func IsValidEthermintChainID(chainID string) bool {
	if len(chainID) > 48 {
		return false
	}

	return regexEthermintChainID.MatchString(chainID)
}

// ParseEthermintChainId parses a string chain identifier's epoch to an Ethereum-compatible
// chain-id in *big.Int format. The function returns an error if the chain-id has an invalid format.
//
// This function, was cloned 1-1 from ethermint lib because go-lib does not link with ethermint lib,
// to be used to provide additional mapping to other chains that chain-id not follow convention of ethermint
func ParseEthermintChainId(chainID string) (*big.Int, error) {
	chainID = strings.TrimSpace(chainID)
	if len(chainID) > 48 {
		return nil, fmt.Errorf("chain-id '%s' cannot exceed 48 chars", chainID)
	}

	matches := regexEthermintChainID.FindStringSubmatch(chainID)
	if matches == nil || len(matches) != 4 || matches[1] == "" {
		return nil, fmt.Errorf("%s: %v", chainID, matches)
	}

	// verify that the chain-id entered is a base 10 integer
	chainIDInt, ok := new(big.Int).SetString(matches[2], 10)
	if !ok {
		return nil, fmt.Errorf("epoch %s must be base-10 integer format", matches[2])
	}

	return chainIDInt, nil
}
