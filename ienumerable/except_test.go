package ienumerable

import (
	"github.com/EscanBE/go-lib/test_utils"
	"testing"
)

func TestIEnumerable_Except(t *testing.T) {
	tests := []struct {
		name    string
		source  []int
		another []int
		want    []int
	}{
		{
			name:    "except not any",
			source:  []int{1, 2, 3},
			another: []int{4, 5, 6, 7},
			want:    []int{1, 2, 3},
		},
		{
			name:    "except one",
			source:  []int{1, 2, 3, 4},
			another: []int{4, 5, 6, 7},
			want:    []int{1, 2, 3},
		},
		{
			name:    "except some",
			source:  []int{1, 2, 3, 5, 6},
			another: []int{4, 5, 6, 7},
			want:    []int{1, 2, 3},
		},
		{
			name:    "except when source empty",
			source:  []int{},
			another: []int{4, 5, 6, 7},
			want:    []int{},
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
			want:    []int{},
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
			eExcept := eSource.Except(eAnother, fCompare)
			test_utils.AssertSlicesEquals[int, int](tt.want, eExcept.data, fCompare, t)

			// Original
			test_utils.AssertSlicesEquals[int, int](originalSource, eSource.data, fCompare, t)
			test_utils.AssertSlicesEquals[int, int](originalAnother, eAnother.data, fCompare, t)
		})
	}
}
