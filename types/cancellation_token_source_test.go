package types

import (
	"testing"
	"time"
)

func TestCancellationTokenSource_GetCancellationToken(t *testing.T) {
	expiry := time.Now().Add(-1 * time.Second)

	cs := NewCancellationTokenSourceWithExpiry(&expiry)

	ct := cs.GetCancellationToken()

	if cs.IsExpired() != ct.IsExpired() {
		t.Errorf("GetCancellationToken() returns un-match expiration")
	}

	if cs != ct.cancellationTokenSource {
		t.Errorf("GetCancellationToken() returns un-match parent")
	}
}

func TestCancellationTokenSource_IsExpired(t *testing.T) {
	expiry := time.Now().Add(-1 * time.Second)

	cs := NewCancellationTokenSourceWithExpiry(&expiry)

	if !cs.IsExpired() {
		t.Errorf("IsExpired() returns wrong expiration state (1)")
	}

	expiry = time.Now().Add(time.Hour)
	cs.expiry = &expiry

	if cs.IsExpired() {
		t.Errorf("IsExpired() returns wrong expiration state (2)")
	}
}

func TestCancellationTokenSource_RequestCancellation(t *testing.T) {
	expiry := time.Now().Add(time.Hour)

	cs := NewCancellationTokenSourceWithExpiry(&expiry)

	cs.RequestCancellation()

	if !cs.IsExpired() {
		t.Errorf("RequestCancellation() does not working as expected")
	}

	cs = NewCancellationTokenSourceWithExpiryAndCounter(&expiry, 999999)

	cs.RequestCancellation()

	if !cs.IsExpired() {
		t.Errorf("RequestCancellation() does not working as expected")
	}
}

func TestCancellationToken_IsExpired(t *testing.T) {
	expiry := time.Now().Add(-1 * time.Second)

	cs := NewCancellationTokenSourceWithExpiry(&expiry)

	ct := cs.GetCancellationToken()

	if !ct.IsExpired() {
		t.Errorf("IsExpired() returns wrong expiration state (1)")
	}

	expiry = time.Now().Add(time.Hour)
	cs.expiry = &expiry

	if ct.IsExpired() {
		t.Errorf("IsExpired() returns wrong expiration state (2)")
	}
}

func TestNewCancellationTokenSource(t *testing.T) {
	cs := NewCancellationTokenSource()

	if cs.IsExpired() {
		t.Errorf("NewCancellationTokenSource() returns wrong expiration state (%v) (%t)", *cs.expiry, time.Now().After(*cs.expiry))
	}
}

func TestNewCancellationTokenSourceWithExpiry(t *testing.T) {
	expiry := time.Now().Add(-1 * time.Second)
	cs := NewCancellationTokenSourceWithExpiry(&expiry)

	if !cs.IsExpired() {
		t.Errorf("NewCancellationTokenSourceWithExpiry() returns wrong expiration state")
	}

	expiry = time.Now().Add(time.Hour)
	cs = NewCancellationTokenSourceWithExpiry(&expiry)

	if cs.IsExpired() {
		t.Errorf("NewCancellationTokenSourceWithExpiry() returns wrong expiration state")
	}
}

func TestNewCancellationTokenSourceWithExpiryAndCounter(t *testing.T) {
	expiry := time.Now().Add(-1 * time.Second)
	cs := NewCancellationTokenSourceWithExpiryAndCounter(&expiry, 999999)

	if !cs.IsExpired() {
		t.Errorf("NewCancellationTokenSourceWithExpiry() returns wrong expiration state")
	}

	expiry = time.Now().Add(time.Hour)
	cs = NewCancellationTokenSourceWithExpiryAndCounter(&expiry, 0)

	if !cs.IsExpired() {
		t.Errorf("NewCancellationTokenSourceWithExpiry() returns wrong expiration state")
	}

	cs = NewCancellationTokenSourceWithExpiryAndCounter(&expiry, 999999)
	if cs.IsExpired() {
		t.Errorf("NewCancellationTokenSourceWithExpiry() returns wrong expiration state")
	}
}

func TestNewCancellationTokenSourceWithTimeoutDuration(t *testing.T) {
	dur := 300 * time.Millisecond
	cs := NewCancellationTokenSourceWithTimeoutDuration(dur)

	if cs.IsExpired() {
		t.Errorf("NewCancellationTokenSourceWithTimeoutDuration() returns wrong expiration state")
	}

	time.Sleep(dur)
	time.Sleep(50 * time.Millisecond)

	if !cs.IsExpired() {
		t.Errorf("NewCancellationTokenSourceWithTimeoutDuration() returns wrong expiration state")
	}
}

func TestNewCancellationTokenSourceWithTimeoutDurationAndCounter(t *testing.T) {
	cs := NewCancellationTokenSourceWithTimeoutDurationAndCounter(time.Hour, 2)

	if cs.IsExpired() {
		t.Errorf("NewCancellationTokenSourceWithTimeoutDuration() returns wrong expiration state")
	}

	cs.ReduceCounter()
	if cs.IsExpired() {
		t.Errorf("NewCancellationTokenSourceWithTimeoutDuration() returns wrong expiration state")
	}

	cs.ReduceCounter()
	if !cs.IsExpired() {
		t.Errorf("NewCancellationTokenSourceWithTimeoutDuration() returns wrong expiration state")
	}
}

func TestNewCancellationTokenSourceWithTimeoutMilliseconds(t *testing.T) {
	cs := NewCancellationTokenSourceWithTimeoutMilliseconds(300)

	if cs.IsExpired() {
		t.Errorf("NewCancellationTokenSourceWithTimeoutMilliseconds() returns wrong expiration state")
	}

	time.Sleep(301 * time.Millisecond)

	if !cs.IsExpired() {
		t.Errorf("NewCancellationTokenSourceWithTimeoutMilliseconds() returns wrong expiration state")
	}
}

func TestNewCancellationTokenSourceWithTimeoutMillisecondsAndCounter(t *testing.T) {
	cs := NewCancellationTokenSourceWithTimeoutMillisecondsAndCounter(9999999, 2)

	if cs.IsExpired() {
		t.Errorf("NewCancellationTokenSourceWithTimeoutMilliseconds() returns wrong expiration state")
	}

	cs.ReduceCounter()
	if cs.IsExpired() {
		t.Errorf("NewCancellationTokenSourceWithTimeoutMilliseconds() returns wrong expiration state")
	}

	cs.ReduceCounter()
	if !cs.IsExpired() {
		t.Errorf("NewCancellationTokenSourceWithTimeoutMilliseconds() returns wrong expiration state")
	}
}

func TestNewCancellationTokenSourceWithTimeoutSeconds(t *testing.T) {
	cs := NewCancellationTokenSourceWithTimeoutSeconds(1)

	if cs.IsExpired() {
		t.Errorf("NewCancellationTokenSourceWithTimeoutSeconds() returns wrong expiration state")
	}

	time.Sleep(2 * time.Second)

	if !cs.IsExpired() {
		t.Errorf("NewCancellationTokenSourceWithTimeoutSeconds() returns wrong expiration state")
	}
}

func TestNewCancellationTokenSourceWithTimeoutSecondsAndCounter(t *testing.T) {
	cs := NewCancellationTokenSourceWithTimeoutSecondsAndCounter(999999, 2)

	if cs.IsExpired() {
		t.Errorf("NewCancellationTokenSourceWithTimeoutSeconds() returns wrong expiration state")
	}

	cs.ReduceCounter()
	if cs.IsExpired() {
		t.Errorf("NewCancellationTokenSourceWithTimeoutSeconds() returns wrong expiration state")
	}

	cs.ReduceCounter()
	if !cs.IsExpired() {
		t.Errorf("NewCancellationTokenSourceWithTimeoutSeconds() returns wrong expiration state")
	}

	ct := NewCancellationTokenSourceWithTimeoutSecondsAndCounter(999999, 2).GetCancellationToken()

	if ct.IsExpired() {
		t.Errorf("NewCancellationTokenSourceWithTimeoutSeconds() returns wrong expiration state")
	}

	ct.ReduceCounter()
	if ct.IsExpired() {
		t.Errorf("NewCancellationTokenSourceWithTimeoutSeconds() returns wrong expiration state")
	}

	ct.ReduceCounter()
	if !ct.IsExpired() {
		t.Errorf("NewCancellationTokenSourceWithTimeoutSeconds() returns wrong expiration state")
	}
}

func TestReduceCounterOnCancellationTokenSource(t *testing.T) {
	cs := NewCancellationTokenSourceWithTimeoutSecondsAndCounter(999999, 2)
	cs.ReduceCounter()
	if cs.counter != 1 {
		t.Errorf("ReduceCounter() didn't change counter state correctly")
	}

	cs.ReduceCounter()
	if cs.counter != 0 {
		t.Errorf("ReduceCounter() didn't change counter state correctly")
	}

	cs.ReduceCounter()
	if cs.counter != 0 {
		t.Errorf("ReduceCounter() didn't change counter state correctly")
	}

	cs.counter = -1

	cs.ReduceCounter()
	if cs.counter != -1 {
		t.Errorf("ReduceCounter() didn't change counter state correctly")
	}
}

func TestReduceCounterOnCancellationToken(t *testing.T) {
	ct := NewCancellationTokenSourceWithTimeoutSecondsAndCounter(999999, 2).GetCancellationToken()
	ct.ReduceCounter()
	if ct.cancellationTokenSource.counter != 1 {
		t.Errorf("ReduceCounter() didn't change counter state correctly")
	}

	ct.ReduceCounter()
	if ct.cancellationTokenSource.counter != 0 {
		t.Errorf("ReduceCounter() didn't change counter state correctly")
	}

	ct.ReduceCounter()
	if ct.cancellationTokenSource.counter != 0 {
		t.Errorf("ReduceCounter() didn't change counter state correctly")
	}

	ct.cancellationTokenSource.counter = -1

	ct.ReduceCounter()
	if ct.cancellationTokenSource.counter != -1 {
		t.Errorf("ReduceCounter() didn't change counter state correctly")
	}
}
