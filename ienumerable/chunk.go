package ienumerable

/*
func (e IEnumerable[T]) Chunk(size int) IEnumerable[[]T] {
	if size < 1 {
		panic(fmt.Errorf("size can not be lower than 1"))
	}
	if len(e.data) < 1 {
		return AsIEnumerable[[]T]([]T{})
	}
	paged, err := utils.Paging(e.data, size)
	if err != nil {
		panic(errors.Wrap(err, "failed to paging elements"))
	}
	return AsIEnumerable[[]T](paged...)
}
*/
