package ienumerable

import "testing"

func TestIEnumerable_All(t *testing.T) {
	isGe5 := func(i int) bool {
		return i >= 5
	}
	tests := []struct {
		name    string
		source  []int
		fAccept func(i int) bool
		want    bool
	}{
		{
			name:    "nil source, should return true",
			source:  nil,
			fAccept: isGe5,
			want:    true,
		},
		{
			name:    "empty source, should return true",
			source:  []int{},
			fAccept: isGe5,
			want:    true,
		},
		{
			name:    "no match",
			source:  []int{1, 2, 3, 4},
			fAccept: isGe5,
			want:    false,
		},
		{
			name:    "all match",
			source:  []int{5, 6, 7, 8},
			fAccept: isGe5,
			want:    true,
		},
		{
			name:    "partially match",
			source:  []int{4, 5, 6, 7, 8},
			fAccept: isGe5,
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := AsIEnumerable[int](tt.source...)
			if got := e.All(tt.fAccept); got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}
