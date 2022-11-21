package utils

// GetKeys returns keys of map as a slice
func GetKeys[K comparable, V any](myMap map[K]V) []K {
	keys := make([]K, 0)
	for key := range myMap {
		keys = append(keys, key)
	}
	return keys
}

// GetKeysOfTrue returns slide of keys which has value is true
func GetKeysOfTrue[K comparable](myMap map[K]bool) []K {
	return GetKeysOf(myMap, true)
}

// GetKeysOf returns slide of keys which has value equals with expected value
func GetKeysOf[K, V comparable](myMap map[K]V, expectedValue V) []K {
	keys := make([]K, 0)
	for key, value := range myMap {
		if value == expectedValue {
			keys = append(keys, key)
		}
	}
	return keys
}

// SoftCloneMap returns a cloned map, the value will be assigned to the new map
func SoftCloneMap[K comparable, V any](myMap map[K]V) map[K]V {
	result := make(map[K]V)
	for key, value := range myMap {
		result[key] = value
	}
	return result
}
