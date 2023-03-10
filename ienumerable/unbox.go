package ienumerable

import "fmt"

func (e IEnumerable[T]) UnboxBool() (eResult IEnumerable[bool], err error) {
	size := len(e.data)
	result := make([]bool, size)
	if size > 0 {
		for i, t := range e.data {
			if cast, ok := any(t).(bool); ok {
				result[i] = cast
			} else {
				err = fmt.Errorf("%v can not be casted to bool", t)
				return
			}
		}
	}

	eResult = AsIEnumerable[bool](result...)
	err = nil
	return
}

func (e IEnumerable[T]) UnsafeUnboxBool() IEnumerable[bool] {
	result, err := e.UnboxBool()
	if err != nil {
		panic(err)
	}
	return result
}

func (e IEnumerable[T]) UnboxByte() (eResult IEnumerable[byte], err error) {
	size := len(e.data)
	result := make([]byte, size)
	if size > 0 {
		for i, t := range e.data {
			if cast, ok := any(t).(byte); ok {
				result[i] = cast
			} else {
				err = fmt.Errorf("%v can not be casted to byte", t)
				return
			}
		}
	}

	eResult = AsIEnumerable[byte](result...)
	err = nil
	return
}

func (e IEnumerable[T]) UnsafeUnboxByte() IEnumerable[byte] {
	result, err := e.UnboxByte()
	if err != nil {
		panic(err)
	}
	return result
}

func (e IEnumerable[T]) UnboxInt() (eResult IEnumerable[int], err error) {
	size := len(e.data)
	result := make([]int, size)
	if size > 0 {
		for i, t := range e.data {
			if cast, ok := any(t).(int); ok {
				result[i] = cast
			} else {
				err = fmt.Errorf("%v can not be casted to int", t)
				return
			}
		}
	}

	eResult = AsIEnumerable[int](result...)
	err = nil
	return
}

func (e IEnumerable[T]) UnsafeUnboxInt() IEnumerable[int] {
	result, err := e.UnboxInt()
	if err != nil {
		panic(err)
	}
	return result
}

func (e IEnumerable[T]) UnboxInt64() (eResult IEnumerable[int64], err error) {
	size := len(e.data)
	result := make([]int64, size)
	if size > 0 {
		for i, t := range e.data {
			if cast, ok := any(t).(int64); ok {
				result[i] = cast
			} else {
				err = fmt.Errorf("%v can not be casted to int64", t)
				return
			}
		}
	}

	eResult = AsIEnumerable[int64](result...)
	err = nil
	return
}

func (e IEnumerable[T]) UnsafeUnboxInt64() IEnumerable[int64] {
	result, err := e.UnboxInt64()
	if err != nil {
		panic(err)
	}
	return result
}

func (e IEnumerable[T]) UnboxFloat64() (eResult IEnumerable[float64], err error) {
	size := len(e.data)
	result := make([]float64, size)
	if size > 0 {
		for i, t := range e.data {
			if cast, ok := any(t).(float64); ok {
				result[i] = cast
			} else {
				err = fmt.Errorf("%v can not be casted to float64", t)
				return
			}
		}
	}

	eResult = AsIEnumerable[float64](result...)
	err = nil
	return
}

func (e IEnumerable[T]) UnsafeUnboxFloat64() IEnumerable[float64] {
	result, err := e.UnboxFloat64()
	if err != nil {
		panic(err)
	}
	return result
}

func (e IEnumerable[T]) UnboxString() (eResult IEnumerable[string], err error) {
	size := len(e.data)
	result := make([]string, size)
	if size > 0 {
		for i, t := range e.data {
			if cast, ok := any(t).(string); ok {
				result[i] = cast
			} else {
				err = fmt.Errorf("%v can not be casted to string", t)
				return
			}
		}
	}

	eResult = AsIEnumerable[string](result...)
	err = nil
	return
}

func (e IEnumerable[T]) UnsafeUnboxString() IEnumerable[string] {
	result, err := e.UnboxString()
	if err != nil {
		panic(err)
	}
	return result
}
