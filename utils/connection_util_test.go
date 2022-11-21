package utils

import (
	"testing"
	"time"
)

func TestCheckPortIsTcpOpen(t *testing.T) {
	type args struct {
		host    string
		port    int
		timeout time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				host:    "google.com",
				port:    443,
				timeout: 30 * time.Second,
			},
			wantErr: false,
		},
		{
			args: args{
				host:    "google.com",
				port:    1,
				timeout: 3 * time.Second,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckPortIsTcpOpen(tt.args.host, tt.args.port, tt.args.timeout); (err != nil) != tt.wantErr {
				t.Errorf("CheckPortIsTcpOpen() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
