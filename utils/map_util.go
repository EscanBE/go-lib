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

// GetMapValues returns slide of all values of the map
func GetMapValues[K comparable, V any](myMap map[K]V) []V {
	values := make([]V, len(myMap))
	cnt := 0
	for _, value := range myMap {
		values[cnt] = value
		cnt++
	}
	return values
}

// SoftCloneMap returns a cloned map, the value will be assigned to the new map
func SoftCloneMap[K comparable, V any](myMap map[K]V) map[K]V {
	result := make(map[K]V)
	for key, value := range myMap {
		result[key] = value
	}
	return result
}

// SliceToTracker converts the slice into a map[K]bool with all values are `true`
func SliceToTracker[K comparable](slice []K) map[K]bool {
	return SliceToMap(slice, true)
}

// SliceToMap converts the slice into a map[K]V with value as provided
func SliceToMap[K comparable, V any](slice []K, value V) map[K]V {
	tracker := make(map[K]V)
	for _, ele := range slice {
		tracker[ele] = value
	}
	return tracker
}

// PutToMapAsKeyBehavior defines behavior when put slice element to map
type PutToMapAsKeyBehavior byte

const (
	// RejectAllWhenAnyDuplicatedKey means if any element of slice already exists on map, no key will be added
	RejectAllWhenAnyDuplicatedKey PutToMapAsKeyBehavior = 1

	// SkipDuplicatedKeys means all slice elements will be put into map as key, except elements which were already exists
	SkipDuplicatedKeys PutToMapAsKeyBehavior = 2

	// AcceptAllAndOverrideDuplicatedKeys means all slice elements will be put into map as key and the value will be overridden by provided default value
	AcceptAllAndOverrideDuplicatedKeys PutToMapAsKeyBehavior = 3

	// AcceptOnlyDuplicatedKeysAndOverrideThem means will take only slice elements which were already exists, and override their value with provided default value. While remaining elements (not exists on map) will be skipped
	AcceptOnlyDuplicatedKeysAndOverrideThem PutToMapAsKeyBehavior = 4
)

// PutAllToMapAsKeys puts all elements from slice into the map, override duplicated elements
func PutAllToMapAsKeys[K comparable, V any](_map map[K]V, slice []K, defaultValue V) map[K]V {
	res, err := PutToMapAsKeys(_map, slice, defaultValue, AcceptAllAndOverrideDuplicatedKeys)
	PanicIfErr(err, "should never happens")
	return res
}

// PutToMapAsKeys puts elements from slice into the map
func PutToMapAsKeys[K comparable, V any](_map map[K]V, slice []K, defaultValue V, behavior PutToMapAsKeyBehavior) (map[K]V, error) {
	if _map == nil {
		_map = make(map[K]V)
	}
	if len(slice) < 1 {
		return _map, nil
	}
	switch behavior {
	case RejectAllWhenAnyDuplicatedKey:
		for _, k := range slice {
			if _, exists := _map[k]; exists {
				return _map, fmt.Errorf("duplicated key %v, rejected all", k)
			}
		}
		for _, k := range slice {
			_map[k] = defaultValue
		}
		return _map, nil
	case SkipDuplicatedKeys:
		for _, k := range slice {
			if _, exists := _map[k]; exists {
				continue
			}
			_map[k] = defaultValue
		}
		return _map, nil
	case AcceptAllAndOverrideDuplicatedKeys:
		for _, k := range slice {
			_map[k] = defaultValue
		}
		return _map, nil
	case AcceptOnlyDuplicatedKeysAndOverrideThem:
		for _, k := range slice {
			if _, exists := _map[k]; exists {
				_map[k] = defaultValue
			}
		}
		return _map, nil
	default:
		return _map, fmt.Errorf("not supported behavior %v", behavior)
	}
}
