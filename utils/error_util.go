package utils

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// ExitIfErr exit the application with error code = 1. If the `err` is nil, this method does nothing
func ExitIfErr(err error, msg string) {
	if err == nil {
		return
	}
	cobra.CheckErr(errors.Wrap(err, msg))
}

// NilOrWrapIfError returns wrapped input error and message. If the `err` is nil, this method also returns `nil`
func NilOrWrapIfError(err error, msg string) error {
	if err == nil {
		return nil
	}
	return errors.Wrap(err, msg)
}
