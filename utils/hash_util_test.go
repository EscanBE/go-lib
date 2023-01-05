package utils

import "testing"

func TestKeccak256Hash(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "abcd",
			want:  "48bed44d1bcd124a28c27f343a817e5f5243190d3c52bf347daf876de1dbbf77",
		},
		{
			input: "",
			want:  "c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := Keccak256Hash(tt.input); got != tt.want {
				t.Errorf("Keccak256Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}
