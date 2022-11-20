package utils

// GetKeys returns keys of map as a slice
func GetKeys[V any](myMap map[string]V) []string {
	keys := make([]string, 0)
	for key := range myMap {
		keys = append(keys, key)
	}
	return keys
}

// GetKeysOfTrue returns slide of keys which has value is true
func GetKeysOfTrue(myMap map[string]bool) []string {
	keys := make([]string, 0)
	for key, value := range myMap {
		if value {
			keys = append(keys, key)
		}
	}
	return keys
}

// CloneMap returns a cloned map, the value is assigned to the new map
func CloneMap[V any](myMap map[string]V) map[string]V {
	result := make(map[string]V)
	for key, value := range myMap {
		result[key] = value
	}
	return result
}

// GetInt64Keys returns keys of map as a slice
func GetInt64Keys[V any](myMap map[int64]V) []int64 {
	keys := make([]int64, 0)
	for key := range myMap {
		keys = append(keys, key)
	}
	return keys
}

// GetInt64KeysOfTrue returns slide of keys which has value is true
func GetInt64KeysOfTrue(myMap map[int64]bool) []int64 {
	keys := make([]int64, 0)
	for key, value := range myMap {
		if value {
			keys = append(keys, key)
		}
	}
	return keys
}

// CloneInt64Map returns a cloned map, the value is assigned to the new map
func CloneInt64Map[V any](myMap map[int64]V) map[int64]V {
	result := make(map[int64]V)
	for key, value := range myMap {
		result[key] = value
	}
	return result
}
