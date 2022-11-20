package app

import "fmt"

// AppExitFunction is an alias of function which receives params
//goland:noinspection GoNameStartsWithPackageName
type AppExitFunction func(params ...any)

// appExitFunction is handler
var appExitFunction AppExitFunction = nil

// RegisterExitFunction registers a handle which should be executed before application exit, by calling ExecuteExitFunction
func RegisterExitFunction(f AppExitFunction) {
	appExitFunction = f
}

// ExecuteExitFunction invokes the registered function, with supplied params. Will panic if no handle was registered before
func ExecuteExitFunction(params ...any) {
	if appExitFunction == nil {
		panic(fmt.Errorf("app exit function was not registered"))
	}
	appExitFunction(params...)
}
