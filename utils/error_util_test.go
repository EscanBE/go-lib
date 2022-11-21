package utils

import (
	"fmt"
	"testing"
)

func TestExitIfErr(t *testing.T) {
	type args struct {
		err error
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "will not exit if non-error",
			args: args{
				err: nil,
				msg: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ExitIfErr(tt.args.err, tt.args.msg)
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
