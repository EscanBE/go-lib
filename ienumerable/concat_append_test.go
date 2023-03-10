package ienumerable

import (
	"github.com/EscanBE/go-lib/test_utils"
	"testing"
)

func TestIEnumerable_ConcatAppend(t *testing.T) {
	tests := []struct {
		name    string
		source  []int
		another []int
		want    []int
	}{
		{
			name:    "normal",
			source:  []int{1, 2, 3},
			another: []int{4, 5, 6, 7},
			want:    []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:    "normal",
			source:  []int{},
			another: []int{4, 5, 6, 7},
			want:    []int{4, 5, 6, 7},
		},
		{
			name:    "normal",
			source:  []int{1, 2, 3},
			another: []int{},
			want:    []int{1, 2, 3},
		},
		{
			name:    "normal with nil src",
			source:  nil,
			another: []int{4, 5, 6, 7},
			want:    []int{4, 5, 6, 7},
		},
		{
			name:    "normal with nil another",
			source:  []int{1, 2, 3},
			another: nil,
			want:    []int{1, 2, 3},
		},
		{
			name:    "normal with both nil",
			source:  nil,
			another: nil,
			want:    []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalSource := copySlice[int](tt.source)
			originalAnother := copySlice[int](tt.another)
			eSource := AsIEnumerable[int](tt.source...)
			eAnother := AsIEnumerable[int](tt.another...)

			fCompare := func(i1, i2 int) bool {
				return i1 == i2
			}

			// Concat
			eConcat := eSource.Concat(eAnother)
			test_utils.AssertSlicesEquals[int, int](tt.want, eConcat.data, fCompare, t)

			// Append
			eAppend := eSource.Append(eAnother)
			test_utils.AssertSlicesEquals[int, int](tt.want, eAppend.data, fCompare, t)

			// Original
			test_utils.AssertSlicesEquals[int, int](originalSource, eSource.data, fCompare, t)
			test_utils.AssertSlicesEquals[int, int](originalAnother, eAnother.data, fCompare, t)
		})
	}
}
