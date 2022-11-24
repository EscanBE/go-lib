package validation

import (
	"fmt"
	"testing"
)

func TestIsValidIP(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		ip   string
		want bool
	}{
		{
			ip:   "127.0.0.1",
			want: true,
		},
		{
			ip:   "::ffff:192.0.2.1",
			want: true,
		},
		{
			ip:   "127.0.0.1:8080",
			want: false,
		},
		{
			ip:   "localhost",
			want: false,
		},
		{
			ip:   "",
			want: false,
		},
		{
			ip:   "127",
			want: false,
		},
		{
			ip:   "127.0",
			want: false,
		},
		{
			ip:   "127.0.0",
			want: false,
		},
		{
			ip:   "127.0.0.1/32",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.ip, func(t *testing.T) {
			if got := IsValidIP(tt.ip); got != tt.want {
				t.Errorf("IsValidIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidPort(t *testing.T) {
	tests := []struct {
		port int
		want bool
	}{
		{
			port: 1,
			want: true,
		},
		{
			port: 65535,
			want: true,
		},
		{
			port: 0,
			want: false,
		},
		{
			port: 65536,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.port), func(t *testing.T) {
			if got := IsValidPort(tt.port); got != tt.want {
				t.Errorf("IsValidPort() = %v, want %v", got, tt.want)
			}
		})
	}
}
