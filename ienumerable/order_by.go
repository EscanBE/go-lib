package ienumerable

import "sort"

func (e IEnumerable[T]) OrderBy(fLess func(t1, t2 T) bool) IEnumerable[T] {
	copied := copySlice(e.data)
	sort.SliceStable(copied, func(i, j int) bool {
		return fLess(copied[i], copied[j])
	})
	return IEnumerable[T]{
		data: copied,
	}
}

func (e IEnumerable[T]) OrderByDescending(fLess func(t1, t2 T) bool) IEnumerable[T] {
	copied := copySlice(e.data)
	sort.SliceStable(copied, func(i, j int) bool {
		return fLess(copied[j], copied[i])
	})
	return IEnumerable[T]{
		data: copied,
	}
}
