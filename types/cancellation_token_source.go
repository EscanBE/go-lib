package types

import (
	"time"
)

// CancellationTokenSource holds an expiry condition, you can either provide expiry time or manually cancel it via RequestCancellation method
type CancellationTokenSource struct {
	expiry *time.Time
}

// NewCancellationTokenSourceWithTimeoutMilliseconds returns an instance of CancellationTokenSource
// which returns true when ask for IsExpired after amount of time passed
func NewCancellationTokenSourceWithTimeoutMilliseconds(timeoutMilliseconds int64) *CancellationTokenSource {
	expiry := time.UnixMilli(time.Now().UnixMilli() + timeoutMilliseconds)
	return NewCancellationTokenSourceWithExpiry(&expiry)
}

// NewCancellationTokenSourceWithTimeoutSeconds returns an instance of CancellationTokenSource
// which returns true when ask for IsExpired after amount of time passed
func NewCancellationTokenSourceWithTimeoutSeconds(timeoutSecs int64) *CancellationTokenSource {
	return NewCancellationTokenSourceWithTimeoutMilliseconds(timeoutSecs * 1000)
}

// NewCancellationTokenSourceWithTimeoutDuration returns an instance of CancellationTokenSource
// which returns true when ask for IsExpired after amount of time passed
func NewCancellationTokenSourceWithTimeoutDuration(duration time.Duration) *CancellationTokenSource {
	expiry := time.Now().Add(duration)
	return NewCancellationTokenSourceWithExpiry(&expiry)
}

// NewCancellationTokenSourceWithExpiry returns an instance of CancellationTokenSource
// which returns true when ask for IsExpired after specific time passed
func NewCancellationTokenSourceWithExpiry(expiry *time.Time) *CancellationTokenSource {
	return &CancellationTokenSource{
		expiry: expiry,
	}
}

// NewCancellationTokenSource returns an instance of CancellationTokenSource
// which returns always return false when ask for IsExpired, but you can cancel it by calling RequestCancellation
func NewCancellationTokenSource() *CancellationTokenSource {
	expiry := time.Date(9999, time.September, 9, 9, 9, 9, 9999, time.UTC)
	return NewCancellationTokenSourceWithExpiry(&expiry)
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
}

// IsExpired returns `true` if expired
func (cs *CancellationTokenSource) IsExpired() bool {
	return cs.expiry == nil || time.Now().After(*cs.expiry)
}

// CancellationToken is child of CancellationTokenSource, it can check if the parent token source is expired
type CancellationToken struct {
	cancellationTokenSource *CancellationTokenSource
}

// IsExpired returns `true` if expired
func (ct CancellationToken) IsExpired() bool {
	return ct.cancellationTokenSource.IsExpired()
}
