package utils

import "fmt"

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

// SlideToTracker converts the slice into a map[K]bool with all values are `true`
func SlideToTracker[K comparable](slice []K) map[K]bool {
	tracker := make(map[K]bool)
	for _, ele := range slice {
		tracker[ele] = true
	}
	return tracker
}

type PutToMapAsKeyBehavior byte

const (
	RejectAllWhenAnyDuplicatedKey           PutToMapAsKeyBehavior = 1
	SkipDuplicatedKeys                      PutToMapAsKeyBehavior = 2
	AcceptAllAndOverrideDuplicatedKeys      PutToMapAsKeyBehavior = 3
	AcceptOnlyDuplicatedKeysAndOverrideThem PutToMapAsKeyBehavior = 4
)

// PutToMapAsKeys puts all elements from slice into the map
func PutToMapAsKeys[K comparable, V any](_map map[K]V, slice []K, defaultValue V, behavior PutToMapAsKeyBehavior) error {
	if _map == nil {
		return fmt.Errorf("map is nil")
	}
	if len(slice) < 1 {
		return nil
	}
	switch behavior {
	case RejectAllWhenAnyDuplicatedKey:
		for _, k := range slice {
			if _, exists := _map[k]; exists {
				return fmt.Errorf("duplicated key %v, rejected all", k)
			}
		}
		for _, k := range slice {
			_map[k] = defaultValue
		}
		return nil
	case SkipDuplicatedKeys:
		for _, k := range slice {
			if _, exists := _map[k]; exists {
				continue
			}
			_map[k] = defaultValue
		}
		return nil
	case AcceptAllAndOverrideDuplicatedKeys:
		for _, k := range slice {
			_map[k] = defaultValue
		}
		return nil
	case AcceptOnlyDuplicatedKeysAndOverrideThem:
		for _, k := range slice {
			if _, exists := _map[k]; exists {
				_map[k] = defaultValue
			}
		}
		return nil
	default:
		return fmt.Errorf("not supported behavior %v", behavior)
	}
}
