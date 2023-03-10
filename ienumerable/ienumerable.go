package ienumerable

// IEnumerable is inherited from C#.
//
// Contract: with every IEnumerable yields from any function, a soft copy of the inner data should be made and original data from original IEnumerable must be keep without any change (soft copy of inner data)
type IEnumerable[T any] struct {
	data []T
}

func AsIEnumerable[T any](elements ...T) IEnumerable[T] {
	return IEnumerable[T]{
		data: elements,
	}
}

func AsIEnumerableWithCast[T any](elements []any, fCast func(input any) T) IEnumerable[T] {
	size := len(elements)
	if size < 1 {
		return AsIEnumerable[T]()
	}
	casted := make([]T, size)
	for i, element := range elements {
		casted[i] = fCast(element)
	}
	return AsIEnumerable[T](casted...)
}

func (e IEnumerable[T]) Len() int {
	return len(e.data)
}

func (e IEnumerable[T]) Select(convert func(t T) any) IEnumerable[any] {
	converted := make([]any, len(e.data))
	for i, t := range e.data {
		converted[i] = convert(t)
	}
	return IEnumerable[any]{
		data: converted,
	}
}

func (e IEnumerable[T]) ForEach(action func(t T)) {
	if len(e.data) > 0 {
		for _, t := range e.data {
			action(t)
		}
	}
}

func (e IEnumerable[T]) ToSlice() []T {
	return copySlice(e.data)
}

func (e IEnumerable[T]) ToArray() []T {
	return e.ToSlice()
}

func copySlice[T any](source []T) []T {
	dst := make([]T, len(source))
	if len(source) > 0 {
		copy(dst, source)
	}
	return dst
}
