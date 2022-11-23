package utils

import "testing"

func TestConditionalInt64(t *testing.T) {
	tests := []struct {
		name       string
		expression bool
		whenTrue   int64
		whenFalse  int64
		want       int64
	}{
		{
			name:       "true",
			expression: true,
			whenTrue:   99,
			whenFalse:  9,
			want:       99,
		},
		{
			name:       "false",
			expression: false,
			whenTrue:   99,
			whenFalse:  9,
			want:       9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConditionalInt64(tt.expression, tt.whenTrue, tt.whenFalse); got != tt.want {
				t.Errorf("ConditionalInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConditionalInt(t *testing.T) {
	tests := []struct {
		name       string
		expression bool
		whenTrue   int
		whenFalse  int
		want       int
	}{
		{
			name:       "true",
			expression: true,
			whenTrue:   99,
			whenFalse:  9,
			want:       99,
		},
		{
			name:       "false",
			expression: false,
			whenTrue:   99,
			whenFalse:  9,
			want:       9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConditionalInt(tt.expression, tt.whenTrue, tt.whenFalse); got != tt.want {
				t.Errorf("ConditionalInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
