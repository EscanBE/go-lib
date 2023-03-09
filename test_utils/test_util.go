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

//goland:noinspection SpellCheckingInspection
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

func AssertSlicesEquals[T1 any, T2 any](expected []T1, got []T2, fCompare func(l T1, r T2) bool, t *testing.T) {
	if len(expected) != len(got) {
		t.Errorf("slices are not equals, expected len %d, got %d", len(expected), len(got))
		return
	}
	if len(expected) > 0 {
		for i, e := range expected {
			if !fCompare(e, got[i]) {
				t.Errorf("slices are not equals, expected[%d] = %v, got %v", i, e, got[i])
				return
			}
		}
	}
}
