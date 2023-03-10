package ienumerable

func (e IEnumerable[T]) Except(another IEnumerable[T], fEquals func(t1, t2 T) bool) IEnumerable[T] {
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
	result := make([]T, 0)
	for _, d := range e.data {
		var foundInAnother bool
		for _, t := range another.data {
			if fEquals(d, t) {
				foundInAnother = true
				break
			}
		}
		if !foundInAnother {
			result = append(result, d)
		}
	}
	return AsIEnumerable[T](result...)
}
