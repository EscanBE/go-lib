package types

import (
	"math"
	"time"
)

// CancellationTokenSource holds an expiry condition, you can either provide expiry time or reduce counter or manually cancel it via RequestCancellation method
type CancellationTokenSource struct {
	expiry  *time.Time
	counter int64
}

// NewCancellationTokenSourceWithTimeoutMilliseconds returns an instance of CancellationTokenSource
// which returns true when ask for IsExpired after amount of time passed
func NewCancellationTokenSourceWithTimeoutMilliseconds(timeoutMilliseconds int64) *CancellationTokenSource {
	expiry := time.UnixMilli(time.Now().UnixMilli() + timeoutMilliseconds)
	return NewCancellationTokenSourceWithExpiry(&expiry)
}

func NewCancellationTokenSourceWithTimeoutMillisecondsAndCounter(timeoutMilliseconds int64, counter int64) *CancellationTokenSource {
	expiry := time.UnixMilli(time.Now().UnixMilli() + timeoutMilliseconds)
	return NewCancellationTokenSourceWithExpiryAndCounter(&expiry, counter)
}

// NewCancellationTokenSourceWithTimeoutSeconds returns an instance of CancellationTokenSource
// which returns true when ask for IsExpired after amount of time passed
func NewCancellationTokenSourceWithTimeoutSeconds(timeoutSecs int64) *CancellationTokenSource {
	return NewCancellationTokenSourceWithTimeoutMilliseconds(timeoutSecs * 1000)
}

func NewCancellationTokenSourceWithTimeoutSecondsAndCounter(timeoutSecs, counter int64) *CancellationTokenSource {
	return NewCancellationTokenSourceWithTimeoutMillisecondsAndCounter(timeoutSecs*1000, counter)
}

// NewCancellationTokenSourceWithTimeoutDuration returns an instance of CancellationTokenSource
// which returns true when ask for IsExpired after amount of time passed
func NewCancellationTokenSourceWithTimeoutDuration(duration time.Duration) *CancellationTokenSource {
	expiry := time.Now().Add(duration)
	return NewCancellationTokenSourceWithExpiry(&expiry)
}

func NewCancellationTokenSourceWithTimeoutDurationAndCounter(duration time.Duration, counter int64) *CancellationTokenSource {
	expiry := time.Now().Add(duration)
	return NewCancellationTokenSourceWithExpiryAndCounter(&expiry, counter)
}

// NewCancellationTokenSourceWithExpiry returns an instance of CancellationTokenSource
// which returns true when ask for IsExpired after specific time passed
func NewCancellationTokenSourceWithExpiry(expiry *time.Time) *CancellationTokenSource {
	return NewCancellationTokenSourceWithExpiryAndCounter(expiry, math.MaxInt64)
}

func NewCancellationTokenSourceWithExpiryAndCounter(expiry *time.Time, counter int64) *CancellationTokenSource {
	return &CancellationTokenSource{
		expiry:  expiry,
		counter: counter,
	}
}

// NewCancellationTokenSource returns an instance of CancellationTokenSource
// which returns always return false when ask for IsExpired, but you can cancel it by calling RequestCancellation
func NewCancellationTokenSource() *CancellationTokenSource {
	expiry := time.Date(9999, time.September, 9, 9, 9, 9, 9999, time.UTC)
	return NewCancellationTokenSourceWithExpiryAndCounter(&expiry, math.MaxInt64)
}

// GetCancellationToken returns an instance of CancellationToken, which can check if expired but does not have ability to request cancellation
func (cs *CancellationTokenSource) GetCancellationToken() CancellationToken {
	return CancellationToken{
		cancellationTokenSource: cs,
	}
}

// RequestCancellation requests for CancellationTokenSource to be expired
func (cs *CancellationTokenSource) RequestCancellation() {
	cs.expiry = nil
	cs.counter = 0
}

func (cs *CancellationTokenSource) ReduceCounter() {
	cs.counter--
}

// IsExpired returns `true` if expired
func (cs *CancellationTokenSource) IsExpired() bool {
	return cs.expiry == nil || time.Now().After(*cs.expiry) || cs.counter < 1
}

// CancellationToken is child of CancellationTokenSource, it can check if the parent token source is expired
type CancellationToken struct {
	cancellationTokenSource *CancellationTokenSource
}

// IsExpired returns `true` if expired
func (ct CancellationToken) IsExpired() bool {
	return ct.cancellationTokenSource.IsExpired()
}
