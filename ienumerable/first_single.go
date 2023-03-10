package ienumerable

import "fmt"

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
