package test_utils

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func DeferWantNoPanic(t *testing.T) {
	err := recover()
	if err != nil {
		t.Errorf("expect not panic")
	}
}

func DeferWantPanic(t *testing.T) {
	err := recover()
	if err == nil {
		t.Errorf("expect panic")
	}
}

func DeferWantPanicDepends(t *testing.T, wantPanic bool) {
	err := recover()
	if err == nil && wantPanic {
		t.Errorf("expect panic")
	} else if err != nil && !wantPanic {
		t.Errorf("expect not panic")
	}
}

// WantErrorContainsStringIfNonEmptyOtherWiseNoError notify error and return `false` if any err not contains provided text
func WantErrorContainsStringIfNonEmptyOtherWiseNoError(t *testing.T, err error, wantErrMsgContains string) bool {
	wantErr := len(wantErrMsgContains) > 0
	if err != nil {
		if !wantErr {
			t.Errorf("want no error, got %v", err)
			return false
		} else if !strings.Contains(err.Error(), wantErrMsgContains) {
			t.Errorf("want error msg [%s] contains string [%s]", err.Error(), wantErrMsgContains)
			return false
		}
	} else {
		if wantErr {
			t.Errorf("want error but no error (expect error contains [%s])", wantErrMsgContains)
			return false
		}
	}
	return true
}

var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

var RadStr = RandomText

func RandomText(length int) string {
	if length < 1 || length > 1000 {
		panic(fmt.Errorf("invalid length %d", length))
	}
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = letters[rand.Int()%len(letters)]
	}
	return string(bytes)
}
