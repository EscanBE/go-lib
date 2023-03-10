package ienumerable

import (
	"github.com/EscanBE/go-lib/test_utils"
	"testing"
)

func TestIEnumerable_OrderBy(t *testing.T) {
	tests := []struct {
		name   string
		source []int
		fOrder func(e1, e2 int) bool
		want   []int
	}{
		{
			name:   "nil source ok",
			source: nil,
			fOrder: func(e1, e2 int) bool {
				return e1 < e2
			},
			want: []int{},
		},
		{
			name:   "empty source ok",
			source: []int{},
			fOrder: func(e1, e2 int) bool {
				return e1 < e2
			},
			want: []int{},
		},
		{
			name:   "asc",
			source: []int{5, 4, 6, 8, 7, 9, 6},
			fOrder: func(e1, e2 int) bool {
				return e1 < e2
			},
			want: []int{4, 5, 6, 6, 7, 8, 9},
		},
		{
			name:   "desc",
			source: []int{5, 4, 6, 8, 7, 9, 6},
			fOrder: func(e1, e2 int) bool {
				return e1 > e2
			},
			want: []int{9, 8, 7, 6, 6, 5, 4},
		},
		{
			name:   "keep original order",
			source: []int{5, 4, 6, 8, 7, 9, 6},
			fOrder: func(e1, e2 int) bool {
				return e1 == e2
			},
			want: []int{5, 4, 6, 8, 7, 9, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalSource := copySlice[int](tt.source)
			eSource := AsIEnumerable[int](tt.source...)

			fCompare := func(e1, e2 int) bool {
				return e1 == e2
			}

			eGot := eSource.OrderBy(tt.fOrder)
			test_utils.AssertSlicesEquals[int, int](tt.want, eGot.data, fCompare, t)

			// Original
			test_utils.AssertSlicesEquals[int, int](originalSource, eSource.data, fCompare, t)
		})
	}
}

func TestIEnumerable_OrderByDescending(t *testing.T) {
	tests := []struct {
		name   string
		source []int
		fOrder func(e1, e2 int) bool
		want   []int
	}{
		{
			name:   "nil source ok",
			source: nil,
			fOrder: func(e1, e2 int) bool {
				return e1 < e2
			},
			want: []int{},
		},
		{
			name:   "empty source ok",
			source: []int{},
			fOrder: func(e1, e2 int) bool {
				return e1 < e2
			},
			want: []int{},
		},
		{
			name:   "desc",
			source: []int{5, 4, 6, 8, 7, 9, 6},
			fOrder: func(e1, e2 int) bool {
				return e1 < e2
			},
			want: []int{9, 8, 7, 6, 6, 5, 4},
		},
		{
			name:   "desc (by reverse)",
			source: []int{5, 4, 6, 8, 7, 9, 6},
			fOrder: func(e1, e2 int) bool {
				return e1 > e2
			},
			want: []int{4, 5, 6, 6, 7, 8, 9},
		},
		{
			name:   "keep original order",
			source: []int{5, 4, 6, 8, 7, 9, 6},
			fOrder: func(e1, e2 int) bool {
				return e1 == e2
			},
			want: []int{5, 4, 6, 8, 7, 9, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalSource := copySlice[int](tt.source)
			eSource := AsIEnumerable[int](tt.source...)

			fCompare := func(e1, e2 int) bool {
				return e1 == e2
			}

			eGot := eSource.OrderByDescending(tt.fOrder)
			test_utils.AssertSlicesEquals[int, int](tt.want, eGot.data, fCompare, t)

			// Original
			test_utils.AssertSlicesEquals[int, int](originalSource, eSource.data, fCompare, t)
		})
	}
}
