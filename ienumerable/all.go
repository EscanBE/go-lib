package ienumerable

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
