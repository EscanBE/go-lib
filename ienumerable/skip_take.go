package ienumerable

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
