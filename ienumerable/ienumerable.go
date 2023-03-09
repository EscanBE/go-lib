package ienumerable

import (
	"fmt"
	"github.com/EscanBE/go-lib/utils"
	"sort"
)

// IEnumerable is inherited from C#.
//
// Contract: with every IEnumerable yields from any function, a soft copy of the inner data should be made and original data from original IEnumerable must be keep without any change (soft copy of inner data)
type IEnumerable[T comparable] struct {
	data []T
}

func AsIEnumerable[T comparable](elements ...T) IEnumerable[T] {
	return IEnumerable[T]{
		data: elements,
	}
}

func (e IEnumerable[T]) Len() int {
	return len(e.data)
}

func (e IEnumerable[T]) Skip(size int) IEnumerable[T] {
	if size >= len(e.data) {
		return IEnumerable[T]{
			data: make([]T, 0),
		}
	}
	return IEnumerable[T]{
		data: e.data[size:],
	}
}

func (e IEnumerable[T]) Take(size int) IEnumerable[T] {
	if size >= len(e.data) {
		return IEnumerable[T]{
			data: copySlice(e.data),
		}
	}
	return IEnumerable[T]{
		data: e.data[:size],
	}
}

func (e IEnumerable[T]) Concat(another IEnumerable[T]) IEnumerable[T] {
	if len(another.data) < 1 {
		return IEnumerable[T]{
			data: copySlice(e.data),
		}
	}
	if len(e.data) < 1 {
		return IEnumerable[T]{
			data: copySlice(another.data),
		}
	}
	return IEnumerable[T]{
		data: append(copySlice(e.data), another.data...),
	}
}

func (e IEnumerable[T]) ConcatSlice(another []T) IEnumerable[T] {
	if len(another) < 1 {
		return IEnumerable[T]{
			data: copySlice(e.data),
		}
	}
	if len(e.data) < 1 {
		return IEnumerable[T]{
			data: copySlice(another),
		}
	}
	return IEnumerable[T]{
		data: append(copySlice(e.data), another...),
	}
}

func (e IEnumerable[T]) Except(another IEnumerable[T]) IEnumerable[T] {
	if len(another.data) < 1 {
		return IEnumerable[T]{
			data: copySlice(e.data),
		}
	}
	if len(e.data) < 1 {
		return IEnumerable[T]{
			data: make([]T, 0),
		}
	}
	trackerOfAnother := utils.SliceToTracker(another.data)
	result := make([]T, 0)
	for _, d := range e.data {
		if _, found := trackerOfAnother[d]; found {
			continue
		}
		result = append(result, d)
	}
	return AsIEnumerable[T](result...)
}

func (e IEnumerable[T]) ExceptSlice(another []T) IEnumerable[T] {
	if len(another) < 1 {
		return IEnumerable[T]{
			data: copySlice(e.data),
		}
	}
	if len(e.data) < 1 {
		return IEnumerable[T]{
			data: make([]T, 0),
		}
	}
	trackerOfAnother := utils.SliceToTracker(another)
	result := make([]T, 0)
	for _, d := range e.data {
		if _, found := trackerOfAnother[d]; found {
			continue
		}
		result = append(result, d)
	}
	return AsIEnumerable[T](result...)
}

func (e IEnumerable[T]) OrderBy(fCompare func(t1, t2 T) bool) IEnumerable[T] {
	copied := copySlice(e.data)
	sort.SliceStable(copied, func(i, j int) bool {
		return fCompare(copied[i], copied[j])
	})
	return IEnumerable[T]{
		data: copied,
	}
}

func (e IEnumerable[T]) Reverse() IEnumerable[T] {
	size := len(e.data)
	if size < 1 {
		return AsIEnumerable[T]()
	}
	result := make([]T, len(e.data))
	for i, t := range e.data {
		result[size-1-i] = t
	}
	return AsIEnumerable[T](result...)
}

func (e IEnumerable[T]) Where(filter func(t T) bool) IEnumerable[T] {
	result := make([]T, 0)
	if len(e.data) > 0 {
		for _, t := range e.data {
			if filter(t) {
				result = append(result, t)
			}
		}
	}
	return IEnumerable[T]{
		data: result,
	}
}

func (e IEnumerable[T]) ToSlice() []T {
	return copySlice(e.data)
}

func (e IEnumerable[T]) ToArray() []T {
	return e.ToSlice()
}

func (e IEnumerable[T]) Any() bool {
	return len(e.data) > 0
}

func (e IEnumerable[T]) AnyBy(filter func(t T) bool) bool {
	return e.Where(filter).Len() > 0
}

func (e IEnumerable[T]) All(fAccept func(t T) bool) bool {
	if len(e.data) < 1 {
		return true
	}

	for _, t := range e.data {
		if !fAccept(t) {
			return false
		}
	}

	return true
}

func (e IEnumerable[T]) First() (result T, err error) {
	if len(e.data) == 0 {
		err = fmt.Errorf("IEnumerable is empty")
		return
	}

	result = e.data[0]
	// err = nil
	return
}

func (e IEnumerable[T]) FirstBy(fAccept func(t T) bool) (result T, found bool) {
	if len(e.data) > 0 {
		for _, t := range e.data {
			if fAccept(t) {
				result = t
				found = true
				return
			}
		}
	}

	return
}

func (e IEnumerable[T]) Single(fAccept func(t T) bool) (result T, found bool, err error) {
	if len(e.data) > 0 {
		filtered := e.Where(fAccept)
		if len(filtered.data) == 0 {
			// not found without any error
			// found = false
			// err = nil
			return
		}

		result = filtered.data[0]
		found = true

		if len(filtered.data) == 1 {
			// err = nil
			return
		}

		err = fmt.Errorf("found more than one element matches")
		return
	}

	return
}

func copySlice[T any](source []T) []T {
	dst := make([]T, len(source))
	if len(source) > 0 {
		copy(dst, source)
	}
	return dst
}
