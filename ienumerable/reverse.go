package ienumerable

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
