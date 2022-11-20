package utils

import (
	"strconv"
	"strings"
)

// Int64ToString converts the input int64 into a string
func Int64ToString(number int64) string {
	return strconv.FormatInt(number, 10)
}

// AnyOf returns true if the `inAny` set contains the input `source`
func AnyOf(source string, inAny ...string) bool {
	for _, str := range inAny {
		if source == str {
			return true
		}
	}

	return false
}

// RemoveEmptyString takes a slice and return another slice without blank elements
func RemoveEmptyString(source []string) []string {
	result := make([]string, 0)
	if source == nil {
		return nil
	}

	for _, str := range source {
		normalized := strings.TrimSpace(str)
		if len(normalized) < 1 {
			continue
		}
		result = append(result, normalized)
	}

	return result
}

// ConditionalString returns string based on input expression
func ConditionalString(expression bool, whenTrue, whenFalse string) string {
	if expression {
		return whenTrue
	} else {
		return whenFalse
	}
}

// IsBlank returns true if the string is a blank string (not contains any character other than space and tab)
func IsBlank(text string) bool {
	return strings.TrimSpace(text) == ""
}

// IsTrimmable returns true if the trimmed version of the string is different with the original one
func IsTrimmable(text string) bool {
	return strings.TrimSpace(text) != text
}
