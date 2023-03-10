package ienumerable

import (
	"fmt"
	"github.com/EscanBE/go-lib/utils"
	"github.com/pkg/errors"
)

func (e IEnumerable[T]) Chunk(size int) [][]T {
	if size < 1 {
		panic(fmt.Errorf("size can not be lower than 1"))
	}
	if len(e.data) < 1 {
		return make([][]T, 0)
	}
	paged, err := utils.Paging(e.data, size)
	if err != nil {
		panic(errors.Wrap(err, "failed to paging elements"))
	}
	return paged
}
