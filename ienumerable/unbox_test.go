package ienumerable

import (
	"testing"
)

func TestIEnumerable_UnboxInt(t *testing.T) {
	eSource := AsIEnumerable[string]("a", "bc", "def", "ghij")
	eExpect := AsIEnumerable[int](1, 2, 3, 4).ToSlice()
	casted := eSource.Select(func(ele string) any {
		return len(ele)
	}).UnsafeUnboxInt().ToSlice()

	if len(casted) != len(eExpect) {
		t.Errorf("expected len %d, got %d", len(eExpect), len(casted))
		return
	}

	for i, c := range casted {
		if c != eExpect[i] {
			t.Errorf("expect [%d] = %d, got %d", i, eExpect[i], c)
		}
	}
}
