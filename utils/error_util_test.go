package utils

import (
	"fmt"
	"github.com/EscanBE/go-lib/test_utils"
	"os"
	"testing"
)

func TestExitIfErr(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{
			name: "will not exit if non-error",
			err:  nil,
		},
		{
			name: "will exit with code 1 if error",
			err:  fmt.Errorf("fake"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got int
			myExit := func(code int) {
				got = code
			}
			if tt.err != nil {
				defer func() {
					osExit = os.Exit // restore
				}()
				osExit = myExit
				ExitIfErr(tt.err, "")
				if got != 1 {
					t.Errorf("program should exit with code 1")
				}
			} else {
				ExitIfErr(tt.err, "")
			}
		})
	}
}

func TestPanicIfErr(t *testing.T) {
	tests := []struct {
		name      string
		err       error
		msg       string
		wantPanic bool
	}{
		{
			name: "will not panic if non-error",
			err:  nil,
			msg:  "",
		},
		{
			name:      "will panic if error",
			err:       fmt.Errorf("panic"),
			msg:       "panic",
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer test_utils.DeferWantPanicDepends(t, tt.wantPanic)
			PanicIfErr(tt.err, tt.msg)
		})
	}
}

func TestNilOrWrapIfError(t *testing.T) {
	type args struct {
		err error
		msg string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should nil",
			args: args{
				err: nil,
				msg: "m",
			},
			wantErr: false,
		},
		{
			name: "should err",
			args: args{
				err: fmt.Errorf("error"),
				msg: "m",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NilOrWrapIfError(tt.args.err, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("NilOrWrapIfError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
