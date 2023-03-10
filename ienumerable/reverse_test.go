package ienumerable

import (
	"github.com/EscanBE/go-lib/test_utils"
	"testing"
)

func TestIEnumerable_Reverse(t *testing.T) {
	tests := []struct {
		name   string
		source []int
		want   []int
	}{
		{
			name:   "nil ok",
			source: nil,
			want:   []int{},
		},
		{
			name:   "empty ok",
			source: []int{},
			want:   []int{},
		},
		{
			name:   "single ok",
			source: []int{2},
			want:   []int{2},
		},
		{
			name:   "reverse",
			source: []int{2, 3, 4, 5, 6},
			want:   []int{6, 5, 4, 3, 2},
		},
		{
			name:   "reverse",
			source: []int{2, 3, 4, 5, 6, 9, 8, 7, 6, 5},
			want:   []int{5, 6, 7, 8, 9, 6, 5, 4, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalSource := copySlice[int](tt.source)
			eSource := AsIEnumerable[int](tt.source...)

			fCompare := func(e1, e2 int) bool {
				return e1 == e2
			}

			eGot := eSource.Reverse()
			test_utils.AssertSlicesEquals[int, int](tt.want, eGot.data, fCompare, t)

			// Original
			test_utils.AssertSlicesEquals[int, int](originalSource, eSource.data, fCompare, t)
		})
	}
}
