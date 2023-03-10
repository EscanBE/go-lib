package ienumerable

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
