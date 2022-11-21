package utils

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"math"
	"strconv"
)

// MaxInt64 returns the greater int64
func MaxInt64(n1, n2 int64) int64 {
	if n1 > n2 {
		return n1
	}
	return n2
}

// MaxInt returns the greater int32
func MaxInt(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

// AbsInt64 returns the Absolute value of the input int64
func AbsInt64(n int64) int64 {
	return int64(math.Abs(float64(n)))
}

// AbsInt returns the Absolute value of the input int32
func AbsInt(n int) int {
	return int(math.Abs(float64(n)))
}

var printerEnglish = message.NewPrinter(language.English)

// FormatInt64 formats the number with comma separates thousands, eg 1000 => 1,000
func FormatInt64(n int64) string {
	return printerEnglish.Sprintf("%d", n)
}

// FormatInt formats the number with comma separates thousands, eg 1000 => 1,000
func FormatInt(n int) string {
	return printerEnglish.Sprintf("%d", n)
}

// Int64ToString converts the input int64 into a string
func Int64ToString(number int64) string {
	return strconv.FormatInt(number, 10)
}

// IntToString converts the input int32 into a string
func IntToString(number int) string {
	return fmt.Sprintf("%d", number)
}
