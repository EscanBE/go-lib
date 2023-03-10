package ienumerable

func (e IEnumerable[T]) Any(filter func(t T) bool) bool {
	return e.Where(filter).Len() > 0
}
