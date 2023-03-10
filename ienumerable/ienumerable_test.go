package ienumerable

import (
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
