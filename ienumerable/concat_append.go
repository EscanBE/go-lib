package ienumerable

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

func (e IEnumerable[T]) Append(another IEnumerable[T]) IEnumerable[T] {
	return e.Concat(another)
}
