package app

import (
	"github.com/EscanBE/go-lib/logging"
	"testing"
)

func TestRegisterExitFunction(t *testing.T) {
	appExitFunction = nil

	value := 0
	RegisterExitFunction(func(params ...any) {
		value = 1
	})
	if appExitFunction == nil {
		t.Errorf("RegisterExitFunction() failed to register exit function")
	}
	RegisterExitFunction(func(params ...any) {
		value = 2
	})
	if appExitFunction == nil {
		t.Errorf("RegisterExitFunction() failed to register exit function")
	}

	appExitFunction()

	if value != 2 {
		t.Errorf("RegisterExitFunction() registered exit function but it does not work as expected")
	}
}

func TestExecuteExitFunction(t *testing.T) {
	value := 0

	RegisterExitFunction(func(params ...any) {
		value = 1
	})

	RegisterExitFunction(func(params ...any) {
		value = 2
	})

	ExecuteExitFunction("3", "4", "5")

	if value != 2 {
		t.Errorf("RegisterExitFunction() registered exit function but it does not work as expected")
	}

	appExitFunction = nil
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	ExecuteExitFunction("3", "4", "5")
}

func TestTryRecoverAndExecuteExitFunctionIfRecovered(t *testing.T) {
	num := 0

	// multiple defer, last in first out

	// so final defer to be run should be declared first
	defer func() {
		r := recover()

		if r == nil {
			t.Errorf("expect panic (re-throw)")
			return
		}

		const want = 1
		got := num
		if got != want {
			t.Errorf("TryRecoverAndExecuteExitFunctionIfRecovered() executed wrongly. Got %d but want %d", got, want)
		}
	}()
	defer TryRecoverAndExecuteExitFunctionIfRecovered(logging.NewDefaultLogger())

	appExitFunction = nil

	RegisterExitFunction(func(params ...any) {
		num += 1
	})

	panic("fake")
}
