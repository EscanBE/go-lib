package utils

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
