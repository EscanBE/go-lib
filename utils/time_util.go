package utils

import "time"

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
