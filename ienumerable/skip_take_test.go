package ienumerable

import (
	"github.com/EscanBE/go-lib/test_utils"
	"testing"
)

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
