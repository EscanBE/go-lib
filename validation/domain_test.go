package validation

import "testing"

func TestIsValidHostname(t *testing.T) {
	tests := []struct {
		hostname string
		want     bool
	}{
		{
			hostname: "localhost",
			want:     true,
		},
		{
			hostname: "g",
			want:     true,
		},
		{
			hostname: "google.com",
			want:     true,
		},
		{
			hostname: "https://google.com",
			want:     false,
		},
		{
			hostname: "google.com:443",
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.hostname, func(t *testing.T) {
			if got := IsValidHostname(tt.hostname); got != tt.want {
				t.Errorf("IsValidHostname() = %v, want %v", got, tt.want)
			}
		})
	}
}
