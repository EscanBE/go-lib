package app

import "testing"

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
}
