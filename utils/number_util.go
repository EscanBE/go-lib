package utils

import (
	"fmt"
	"math/big"
	"strings"
)

// ConditionalInt64 returns int64 based on input expression
func ConditionalInt64(expression bool, whenTrue, whenFalse int64) int64 {
	if expression {
		return whenTrue
	} else {
		return whenFalse
	}
}

// ConditionalInt returns int based on input expression
func ConditionalInt(expression bool, whenTrue, whenFalse int) int {
	if expression {
		return whenTrue
	} else {
		return whenFalse
	}
}

//goland:noinspection SpellCheckingInspection
const validHexCharacters = "0123456789abcdefABCDEF"

// IsValidHexNumber returns true if input is a valid hex
func IsValidHexNumber(hexNumberStr string) bool {
	startIdx := 0
	if len(hexNumberStr) < 1 {
		return false
	} else if len(hexNumberStr) == 1 {
		return strings.Index(validHexCharacters, hexNumberStr) >= 0
	} else if len(hexNumberStr) == 2 {
		return strings.Index(validHexCharacters, hexNumberStr[0:1]) >= 0 && strings.Index(validHexCharacters, hexNumberStr[1:]) >= 0
	} else if hexNumberStr[1] == 'x' || hexNumberStr[1] == 'X' {
		if hexNumberStr[0] != '0' {
			return false
		}

		startIdx = 2
	}

	for i := startIdx; i < len(hexNumberStr); i++ {
		c := hexNumberStr[i]
		if c >= 48 && c <= 57 { // 0-9
			continue
		}
		if c >= 65 && c <= 70 { // A-F
			continue
		}
		if c >= 97 && c <= 102 { // a-f
			continue
		}
		return false
	}

	return true
}

// ConvertFromHexNumberStringToDecimalString returns decimal representation in string value of input hex value
func ConvertFromHexNumberStringToDecimalString(hexNumberStr string) (string, error) {
	return convertFromHexNumberStringToDecimalString(hexNumberStr, false)
}

// convertFromHexNumberStringToDecimalString returns decimal representation in string value of input hex value
func convertFromHexNumberStringToDecimalString(hexNumberStr string, bypassValidation bool) (string, error) {
	if !bypassValidation {
		if len(hexNumberStr) > 0 && strings.Index(hexNumberStr, "-") == 0 {
			return "", fmt.Errorf("support positive number only")
		}
		if !IsValidHexNumber(hexNumberStr) {
			return "", fmt.Errorf("not a valid hex")
		}
	}

	bi := new(big.Int)
	var success bool

	if len(hexNumberStr) > 1 && (hexNumberStr[1] == 'x' || hexNumberStr[1] == 'X') {
		bi, success = bi.SetString(hexNumberStr[2:], 16)
	} else {
		bi, success = bi.SetString(hexNumberStr, 16)
	}

	if !success {
		return "", fmt.Errorf("unable to set hex number [%s] to *big.Int", hexNumberStr)
	}

	return bi.String(), nil
}
