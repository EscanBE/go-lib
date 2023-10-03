package app

import (
	"fmt"
	"github.com/EscanBE/go-lib/logging"
	"github.com/EscanBE/go-lib/utils"
)

// AppExitFunction is an alias of function which receives params
//
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

// TryRecoverAndExecuteExitFunctionIfRecovered will check if method has exited with panic.
// If recovered, it will execute exit function and then panic again using that error.
// Otherwise, do nothing (when recover is nil)
func TryRecoverAndExecuteExitFunctionIfRecovered(logger logging.Logger, exitFuncParams ...any) {
	err := recover()
	if err != nil {
		if appExitFunction == nil {
			if logger != nil {
				logger.Error("Panic caught", "error", err)
			} else {
				utils.PrintfStdErr("Panic caught, err: %v\n", err)
			}
			panic(err)
		} else {
			if logger != nil {
				logger.Error("Recovered from panic, executing exit function")
			} else {
				utils.PrintfStdErr("Recovered from panic, executing exit function")
			}
			ExecuteExitFunction(exitFuncParams...)
			if logger != nil {
				logger.Error("Executed exit function, going to panic using recovered error")
			} else {
				utils.PrintfStdErr("Executed exit function, going to panic using recovered error")
			}
			panic(err)
		}
	}
}
