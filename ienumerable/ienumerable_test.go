package ienumerable

import (
	"github.com/EscanBE/go-lib/test_utils"
	"math/rand"
	"testing"
)

func TestIEnumerable_Len(t *testing.T) {
	testLen := func(expectedLen int, input ...int) {
		got := AsIEnumerable[int](input...).Len()
		if expectedLen != got {
			t.Errorf("expected len = %d, got %d", expectedLen, got)
		}
	}

	testLen(0)
	testLen(2, 5, 6)
	random := rand.Intn(1_000)
	testLen(random, make([]int, random)...)
}

func TestIEnumerable_Skip(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		skip      int
		want      []int
		wantPanic bool
	}{
		{
			name:  "",
			input: nil,
			skip:  0,
			want:  []int{},
		},
		{
			name:  "empty but skip larger",
			input: nil,
			skip:  2,
			want:  []int{},
		},
		{
			name:  "not empty and skip larger",
			input: []int{1, 2},
			skip:  4,
			want:  []int{},
		},
		{
			name:  "skip all",
			input: []int{1, 2, 3},
			skip:  3,
			want:  []int{},
		},
		{
			name:  "skip less",
			input: []int{1, 2, 3},
			skip:  2,
			want:  []int{3},
		},
		{
			name:      "skip negative",
			input:     []int{1, 2, 3},
			skip:      -2,
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer test_utils.DeferWantPanicDepends(t, tt.wantPanic)
			originalInput := copySlice(tt.input)
			e := AsIEnumerable[int](tt.input...)
			got := e.Skip(tt.skip)
			if got.Len() != len(tt.want) {
				t.Errorf("Skip(%d) = %v, want %v", tt.skip, got.data, tt.want)
				return
			}
			if got.Len() < 1 {
				return
			}
			for i, t2 := range got.data {
				if t2 != tt.want[i] {
					t.Errorf("Skip(%d) = %v, want %v", tt.skip, got.data, tt.want)
				}
			}
			// test original must not be changed
			test_utils.AssertSlicesEquals[int, int](originalInput, e.data, func(i1, i2 int) bool {
				return i1 == i2
			}, t)
		})
	}
}

func TestIEnumerable_Take(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		take      int
		want      []int
		wantPanic bool
	}{
		{
			name:  "",
			input: nil,
			take:  0,
			want:  []int{},
		},
		{
			name:  "empty but take larger",
			input: nil,
			take:  2,
			want:  []int{},
		},
		{
			name:  "not empty and take larger",
			input: []int{1, 2},
			take:  4,
			want:  []int{1, 2},
		},
		{
			name:  "take all",
			input: []int{1, 2, 3},
			take:  3,
			want:  []int{1, 2, 3},
		},
		{
			name:  "take less",
			input: []int{1, 2, 3},
			take:  2,
			want:  []int{1, 2},
		},
		{
			name:      "take negative",
			input:     []int{1, 2, 3},
			take:      -2,
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer test_utils.DeferWantPanicDepends(t, tt.wantPanic)
			originalInput := copySlice(tt.input)
			e := AsIEnumerable[int](tt.input...)
			got := e.Take(tt.take)
			if got.Len() != len(tt.want) {
				t.Errorf("Take(%d) = %v, want %v", tt.take, got.data, tt.want)
				return
			}
			if got.Len() < 1 {
				return
			}
			for i, t2 := range got.data {
				if t2 != tt.want[i] {
					t.Errorf("Take(%d) = %v, want %v", tt.take, got.data, tt.want)
				}
			}
			// test original must not be changed
			test_utils.AssertSlicesEquals[int, int](originalInput, e.data, func(i1, i2 int) bool {
				return i1 == i2
			}, t)
		})
	}
}

func TestIEnumerable_Concat(t *testing.T) {
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

			// ConcatSlice
			eConcatSlice := eSource.ConcatSlice(tt.another)
			test_utils.AssertSlicesEquals[int, int](tt.want, eConcatSlice.data, fCompare, t)

			// Original
			test_utils.AssertSlicesEquals[int, int](originalSource, eSource.data, fCompare, t)
			test_utils.AssertSlicesEquals[int, int](originalAnother, eAnother.data, fCompare, t)
		})
	}
}

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
			eExcept := eSource.Except(eAnother)
			test_utils.AssertSlicesEquals[int, int](tt.want, eExcept.data, fCompare, t)

			// ConcatSlice
			eExceptSlice := eSource.ExceptSlice(tt.another)
			test_utils.AssertSlicesEquals[int, int](tt.want, eExceptSlice.data, fCompare, t)

			// Original
			test_utils.AssertSlicesEquals[int, int](originalSource, eSource.data, fCompare, t)
			test_utils.AssertSlicesEquals[int, int](originalAnother, eAnother.data, fCompare, t)
		})
	}
}

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
