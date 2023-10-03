package utils

import (
	"fmt"
	"math"
	"time"
)

// NowS returns the current epoch seconds
func NowS() int64 {
	return time.Now().Unix()
}

// NowMs returns the current epoch milliseconds
func NowMs() int64 {
	return time.Now().UnixMilli()
}

// DiffS returns offset amount between the current epoch seconds and supplied epoch from the `previous` arg
func DiffS(previous int64) int64 {
	return NowS() - previous
}

// DiffMs returns offset amount between the current epoch milliseconds and supplied epoch from the `previous` arg
func DiffMs(previous int64) int64 {
	return NowMs() - previous
}

// GetLocationFromUtcTimezone returns location corresponding to specified UTC-based timezone
func GetLocationFromUtcTimezone(utcTimezone int) *time.Location {
	ensureUtcTimezone(utcTimezone)
	return time.FixedZone(GetUtcName(utcTimezone), utcTimezone*60*60)
}

// GetUtcName returns naming convention of UTC timezone. Eg: 7 => UTC+0700
func GetUtcName(utcTimezone int) string {
	ensureUtcTimezone(utcTimezone)
	return fmt.Sprintf("UTC%s", getTimezoneSuffix(utcTimezone))
}

// ensureUtcTimezone will panic if timezone is out of range from -12 to 14
func ensureUtcTimezone(utcTimezone int) {
	if utcTimezone < -12 || utcTimezone > 14 {
		panic(fmt.Errorf("UTC timezone must be in range -12 to 14"))
	}
}

func getTimezoneSuffix(timezone int) string {
	if timezone > 9 {
		return fmt.Sprintf("+%d00", timezone)
	} else if timezone >= 0 {
		return fmt.Sprintf("+0%d00", timezone)
	} else if timezone >= -9 {
		return fmt.Sprintf("-0%d00", int(math.Abs(float64(timezone))))
	} else {
		return fmt.Sprintf("%d00", timezone)
	}
}
