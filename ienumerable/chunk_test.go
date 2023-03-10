package ienumerable

import (
	"github.com/EscanBE/go-lib/test_utils"
	"testing"
)

func TestIEnumerable_Chunk(t *testing.T) {
	tests := []struct {
		name      string
		source    []int
		chunkSize int
		want      [][]int
		wantPanic bool
	}{
		{
			name:      "nil input",
			source:    nil,
			chunkSize: 1,
			want:      [][]int{},
		},
		{
			name:      "bad chunk size",
			source:    nil,
			chunkSize: 0,
			wantPanic: true,
		},
		{
			name:      "full",
			source:    []int{1, 2, 3, 4, 5, 6},
			chunkSize: 6,
			want:      [][]int{{1, 2, 3, 4, 5, 6}},
		},
		{
			name:      "half",
			source:    []int{1, 2, 3, 4, 5, 6},
			chunkSize: 3,
			want:      [][]int{{1, 2, 3}, {4, 5, 6}},
		},
		{
			name:      "1/3",
			source:    []int{1, 2, 3, 4, 5, 6},
			chunkSize: 2,
			want:      [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name:      "lonely",
			source:    []int{1, 2, 3, 4, 5, 6},
			chunkSize: 1,
			want:      [][]int{{1}, {2}, {3}, {4}, {5}, {6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalSource := copySlice(tt.source)
			defer test_utils.DeferWantPanicDepends(t, tt.wantPanic)
			eSource := AsIEnumerable[int](tt.source...)
			got := eSource.Chunk(tt.chunkSize)

			if got == nil {
				t.Errorf("output can not be nil")
			}

			if len(tt.source) < 1 && len(got) != 0 {
				t.Errorf("expected empty output")
				return
			}

			if len(got) != len(tt.want) {
				t.Errorf("want %d elements, got %d", len(tt.want), len(got))
				return
			}

			if len(tt.want) < 1 {
				return
			}

			for i, wantSlice := range tt.want {
				gotSlice := got[i]
				if len(wantSlice) < 1 {
					t.Errorf("result can not contains empty inner slice")
					return
				}
				if len(gotSlice) != len(wantSlice) {
					t.Errorf("at [%d] slice, want %d elements, got %d", i, len(wantSlice), len(gotSlice))
					return
				}
				for i2, wantInner := range wantSlice {
					gotInner := gotSlice[i2]
					if wantInner != gotInner {
						t.Errorf("want[%d][%d] = %d, got %d", i, i2, wantInner, gotInner)
					}
				}
			}

			// Original
			test_utils.AssertSlicesEquals[int, int](originalSource, eSource.data, func(i1, i2 int) bool {
				return i1 == i2
			}, t)
		})
	}
}
