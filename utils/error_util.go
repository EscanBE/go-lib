package utils

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// ExitIfErr exit the application with error code = 1. If the `err` is nil, this method does nothing
func ExitIfErr(err error, msg string) {
	if err == nil {
		return
	}
	fmt.Printf("Exit with error: %s\n", msg)
	cobra.CheckErr(err)
}

// PanicIfErr raises a panic. If the `err` is nil, this method does nothing
func PanicIfErr(err error, msg string) {
	if err == nil {
		return
	}
	fmt.Printf("Exit with error: %s\n", msg)
	panic(err)
}

// NilOrWrapIfError returns wrapped input error and message. If the `err` is nil, this method also returns `nil`
func NilOrWrapIfError(err error, msg string) error {
	if err == nil {
		return nil
	}
	return errors.Wrap(err, msg)
}
