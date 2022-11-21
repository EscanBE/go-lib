package types

import (
	"testing"
	"time"
)

func TestCancellationTokenSource_GetCancellationToken(t *testing.T) {
	expiry := time.Now().Add(-1 * time.Second)

	cs := &CancellationTokenSource{
		expiry: &expiry,
	}

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

	cs := &CancellationTokenSource{
		expiry: &expiry,
	}

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

	cs := &CancellationTokenSource{
		expiry: &expiry,
	}

	cs.RequestCancellation()

	if !cs.IsExpired() {
		t.Errorf("RequestCancellation() does not working as expected")
	}
}

func TestCancellationToken_IsExpired(t *testing.T) {
	expiry := time.Now().Add(-1 * time.Second)

	cs := &CancellationTokenSource{
		expiry: &expiry,
	}

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
