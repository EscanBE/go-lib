package utils

import (
	"fmt"
	"testing"
)

func TestExitIfErr(t *testing.T) {
	tests := []struct {
		name string
		err  error
		msg  string
	}{
		{
			name: "will not exit if non-error",
			err:  nil,
			msg:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ExitIfErr(tt.err, tt.msg)
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
			defer func() {
				r := recover()
				if r == nil && tt.wantPanic {
					t.Errorf("The code did not panic")
				} else if r != nil && !tt.wantPanic {
					t.Errorf("The code should panic")
				}
			}()
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
