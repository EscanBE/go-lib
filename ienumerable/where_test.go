package ienumerable

import (
	"github.com/EscanBE/go-lib/test_utils"
	"testing"
)

func TestIEnumerable_Where(t *testing.T) {
	tests := []struct {
		name   string
		source []int
		filter func(e int) bool
		want   []int
	}{
		{
			name:   "nil source ok",
			source: nil,
			filter: func(e int) bool {
				return e == 5
			},
			want: []int{},
		},
		{
			name:   "empty source ok",
			source: []int{},
			filter: func(e int) bool {
				return e == 5
			},
			want: []int{},
		},
		{
			name:   "filter equals",
			source: []int{5, 4, 6, 8, 7, 9, 6},
			filter: func(e int) bool {
				return e == 6
			},
			want: []int{6, 6},
		},
		{
			name:   "filter not equals and keep original order",
			source: []int{5, 4, 6, 8, 7, 9, 6},
			filter: func(e int) bool {
				return e != 6
			},
			want: []int{5, 4, 8, 7, 9},
		},
		{
			name:   "keep original order",
			source: []int{5, 4, 6, 8, 7, 9, 6},
			filter: func(e int) bool {
				return e < 9
			},
			want: []int{5, 4, 6, 8, 7, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalSource := copySlice[int](tt.source)
			eSource := AsIEnumerable[int](tt.source...)

			fCompare := func(e1, e2 int) bool {
				return e1 == e2
			}

			eGot := eSource.Where(tt.filter)
			test_utils.AssertSlicesEquals[int, int](tt.want, eGot.data, fCompare, t)

			// Original
			test_utils.AssertSlicesEquals[int, int](originalSource, eSource.data, fCompare, t)
		})
	}
}
