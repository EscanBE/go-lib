package ienumerable

func (e IEnumerable[T]) Select(fCast func(t T) any) IEnumerable[any] {
	size := len(e.data)
	result := make([]any, size)
	if size > 0 {
		for i, data := range e.data {
			result[i] = fCast(data)
		}
	}
	return AsIEnumerable[any](result...)
}

func (e IEnumerable[T]) SelectBool(fCast func(t T) bool) IEnumerable[bool] {
	size := len(e.data)
	result := make([]bool, size)
	if size > 0 {
		for i, data := range e.data {
			result[i] = fCast(data)
		}
	}
	return AsIEnumerable[bool](result...)
}

func (e IEnumerable[T]) SelectByte(fCast func(t T) byte) IEnumerable[byte] {
	size := len(e.data)
	result := make([]byte, size)
	if size > 0 {
		for i, data := range e.data {
			result[i] = fCast(data)
		}
	}
	return AsIEnumerable[byte](result...)
}

func (e IEnumerable[T]) SelectInt(fCast func(t T) int) IEnumerable[int] {
	size := len(e.data)
	result := make([]int, size)
	if size > 0 {
		for i, data := range e.data {
			result[i] = fCast(data)
		}
	}
	return AsIEnumerable[int](result...)
}

func (e IEnumerable[T]) SelectInt64(fCast func(t T) int64) IEnumerable[int64] {
	size := len(e.data)
	result := make([]int64, size)
	if size > 0 {
		for i, data := range e.data {
			result[i] = fCast(data)
		}
	}
	return AsIEnumerable[int64](result...)
}

func (e IEnumerable[T]) SelectFloat64(fCast func(t T) float64) IEnumerable[float64] {
	size := len(e.data)
	result := make([]float64, size)
	if size > 0 {
		for i, data := range e.data {
			result[i] = fCast(data)
		}
	}
	return AsIEnumerable[float64](result...)
}

func (e IEnumerable[T]) SelectString(fCast func(t T) string) IEnumerable[string] {
	size := len(e.data)
	result := make([]string, size)
	if size > 0 {
		for i, data := range e.data {
			result[i] = fCast(data)
		}
	}
	return AsIEnumerable[string](result...)
}

func (e IEnumerable[T]) Cast(fCast func(t T) any) IEnumerable[any] {
	return e.Select(fCast)
}

func (e IEnumerable[T]) CastBool(fCast func(t T) bool) IEnumerable[bool] {
	return e.SelectBool(fCast)
}

func (e IEnumerable[T]) CastByte(fCast func(t T) byte) IEnumerable[byte] {
	return e.SelectByte(fCast)
}

func (e IEnumerable[T]) CastInt(fCast func(t T) int) IEnumerable[int] {
	return e.SelectInt(fCast)
}

func (e IEnumerable[T]) CastInt64(fCast func(t T) int64) IEnumerable[int64] {
	return e.SelectInt64(fCast)
}

func (e IEnumerable[T]) CastFloat64(fCast func(t T) float64) IEnumerable[float64] {
	return e.SelectFloat64(fCast)
}

func (e IEnumerable[T]) CastString(fCast func(t T) string) IEnumerable[string] {
	return e.SelectString(fCast)
}
