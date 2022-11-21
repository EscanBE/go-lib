package utils

import (
	"testing"
)

func TestAbsInt(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		{args: 1, want: 1},
		{args: 99, want: 99},
		{args: -99, want: 99},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbsInt(tt.args); got != tt.want {
				t.Errorf("AbsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbsInt64(t *testing.T) {
	tests := []struct {
		name string
		args int64
		want int64
	}{
		{args: 1, want: 1},
		{args: 99, want: 99},
		{args: -99, want: 99},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbsInt64(tt.args); got != tt.want {
				t.Errorf("AbsInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatInt(t *testing.T) {
	tests := []struct {
		args int
		want string
	}{
		{args: 0, want: "0"},
		{args: 99, want: "99"},
		{args: 999, want: "999"},
		{args: 1_000, want: "1,000"},
		{args: 9_999, want: "9,999"},
		{args: 10_000, want: "10,000"},
		{args: 999_000, want: "999,000"},
		{args: 1_000_000, want: "1,000,000"},
		{args: 999_999_999, want: "999,999,999"},
		{args: 1_000_000_000, want: "1,000,000,000"},
		{args: 999_999_999_999, want: "999,999,999,999"},
		{args: 999_999_999_999_999, want: "999,999,999,999,999"},
		{args: -99, want: "-99"},
		{args: -999, want: "-999"},
		{args: -1_000, want: "-1,000"},
		{args: -9_999, want: "-9,999"},
		{args: -10_000, want: "-10,000"},
		{args: -999_000, want: "-999,000"},
		{args: -1_000_000, want: "-1,000,000"},
		{args: -999_999_999, want: "-999,999,999"},
		{args: -1_000_000_000, want: "-1,000,000,000"},
		{args: -999_999_999_999, want: "-999,999,999,999"},
		{args: -999_999_999_999_999, want: "-999,999,999,999,999"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := FormatInt(tt.args); got != tt.want {
				t.Errorf("FormatInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatInt64(t *testing.T) {
	tests := []struct {
		args int64
		want string
	}{
		{args: 0, want: "0"},
		{args: 99, want: "99"},
		{args: 999, want: "999"},
		{args: 1_000, want: "1,000"},
		{args: 9_999, want: "9,999"},
		{args: 10_000, want: "10,000"},
		{args: 999_000, want: "999,000"},
		{args: 1_000_000, want: "1,000,000"},
		{args: 999_999_999, want: "999,999,999"},
		{args: 1_000_000_000, want: "1,000,000,000"},
		{args: 999_999_999_999, want: "999,999,999,999"},
		{args: 999_999_999_999_999, want: "999,999,999,999,999"},
		{args: -99, want: "-99"},
		{args: -999, want: "-999"},
		{args: -1_000, want: "-1,000"},
		{args: -9_999, want: "-9,999"},
		{args: -10_000, want: "-10,000"},
		{args: -999_000, want: "-999,000"},
		{args: -1_000_000, want: "-1,000,000"},
		{args: -999_999_999, want: "-999,999,999"},
		{args: -1_000_000_000, want: "-1,000,000,000"},
		{args: -999_999_999_999, want: "-999,999,999,999"},
		{args: -999_999_999_999_999, want: "-999,999,999,999,999"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := FormatInt64(tt.args); got != tt.want {
				t.Errorf("FormatInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64ToString(t *testing.T) {
	tests := []struct {
		args int64
		want string
	}{
		{args: -1_000_000, want: "-1000000"},
		{args: -999, want: "-999"},
		{args: -99, want: "-99"},
		{args: 0, want: "0"},
		{args: 99, want: "99"},
		{args: 999, want: "999"},
		{args: 1_000_000, want: "1000000"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := Int64ToString(tt.args); got != tt.want {
				t.Errorf("Int64ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntToString(t *testing.T) {
	tests := []struct {
		name string
		args int
		want string
	}{
		{args: -1_000_000, want: "-1000000"},
		{args: -999, want: "-999"},
		{args: -99, want: "-99"},
		{args: 0, want: "0"},
		{args: 99, want: "99"},
		{args: 999, want: "999"},
		{args: 1_000_000, want: "1000000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToString(tt.args); got != tt.want {
				t.Errorf("IntToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt(t *testing.T) {
	type args struct {
		n1 int
		n2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n1: 1,
				n2: 2,
			},
			want: 2,
		},
		{
			args: args{
				n1: -9999,
				n2: 2,
			},
			want: 2,
		},
		{
			args: args{
				n1: 3,
				n2: -9999,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt(tt.args.n1, tt.args.n2); got != tt.want {
				t.Errorf("MaxInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt64(t *testing.T) {
	type args struct {
		n1 int64
		n2 int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				n1: 1,
				n2: 2,
			},
			want: 2,
		},
		{
			args: args{
				n1: -9999,
				n2: 2,
			},
			want: 2,
		},
		{
			args: args{
				n1: 3,
				n2: -9999,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt64(tt.args.n1, tt.args.n2); got != tt.want {
				t.Errorf("MaxInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinInt(t *testing.T) {
	type args struct {
		n1 int
		n2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n1: 1,
				n2: 2,
			},
			want: 1,
		},
		{
			args: args{
				n1: -2,
				n2: 9999,
			},
			want: -2,
		},
		{
			args: args{
				n1: 9999,
				n2: -3,
			},
			want: -3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinInt(tt.args.n1, tt.args.n2); got != tt.want {
				t.Errorf("MinInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinInt64(t *testing.T) {
	type args struct {
		n1 int64
		n2 int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				n1: 1,
				n2: 2,
			},
			want: 1,
		},
		{
			args: args{
				n1: -2,
				n2: 9999,
			},
			want: -2,
		},
		{
			args: args{
				n1: 9999,
				n2: -3,
			},
			want: -3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinInt64(tt.args.n1, tt.args.n2); got != tt.want {
				t.Errorf("MinInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
