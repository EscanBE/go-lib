package utils

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

var osExit = os.Exit

// ExitIfErr exit the application with error code = 1. If the `err` is nil, this method does nothing
func ExitIfErr(err error, msg string) {
	if err == nil {
		return
	}
	PrintlnStdErr("Exit with error:", msg, "\n", err)
	osExit(1)
}

// PanicIfErr raises a panic. If the `err` is nil, this method does nothing
func PanicIfErr(err error, msg string) {
	if err == nil {
		return
	}
	PrintlnStdErr("Exit with error:", msg)
	panic(err)
}

// NilOrWrapIfError returns wrapped input error and message. If the `err` is nil, this method also returns `nil`
func NilOrWrapIfError(err error, msg string) error {
	if err == nil {
		return nil
	}
	return errors.Wrap(err, msg)
}

// PrintlnStdErr does println to StdErr
func PrintlnStdErr(a ...any) {
	fmt.Fprintln(os.Stderr, a...)
}

// PrintfStdErr does printf to StdErr
func PrintfStdErr(format string, a ...any) {
	fmt.Fprintf(os.Stderr, format, a...)
}

// PrintStdErr does print to StdErr
func PrintStdErr(a ...any) {
	fmt.Fprint(os.Stderr, a...)
}
