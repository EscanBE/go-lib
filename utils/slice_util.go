package utils

import "fmt"

// GetUniqueElements returns a distinct slide based on input
func GetUniqueElements[T comparable](slice ...T) []T {
	tracker := make(map[T]bool)
	for _, ele := range slice {
		tracker[ele] = true
	}

	result := make([]T, 0)
	for k := range tracker {
		result = append(result, k)
	}

	return result
}

// Paging returns split original slide into multiple batches.
// Will returns error if page-size lower than 1,
// will returns empty if input slide is empty
func Paging[T any](slice []T, pageSize int) ([][]T, error) {
	if pageSize < 1 {
		return nil, fmt.Errorf("bad page size %d", pageSize)
	}

	result := make([][]T, 0)
	tmpSlide := make([]T, 0)
	if len(slice) < 1 {
		result = append(result, tmpSlide)
		return result, nil
	}

	for _, ele := range slice {
		if len(tmpSlide) < pageSize {
			tmpSlide = append(tmpSlide, ele) // just append ele
		} else {
			result = append(result, tmpSlide) // persist previous
			tmpSlide = make([]T, 0)           // remake
			tmpSlide = append(tmpSlide, ele)  // append ele
		}
	}

	if len(tmpSlide) > 0 {
		result = append(result, tmpSlide)
	}

	return result, nil
}
