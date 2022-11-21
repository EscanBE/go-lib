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

type AddSliceToTrackerBehavior byte

const (
	RejectAllWhenAnyDuplicatedKey           AddSliceToTrackerBehavior = 1
	SkipDuplicatedKeys                      AddSliceToTrackerBehavior = 2
	AcceptAllAndOverrideDuplicatedKeys      AddSliceToTrackerBehavior = 3
	AcceptOnlyDuplicatedKeysAndOverrideThem AddSliceToTrackerBehavior = 4
)

// AddSliceToTracker puts all elements from slice into the map, with map value is `true`.
func AddSliceToTracker[K comparable](tracker map[K]bool, slice []K, defaultValue bool, behavior AddSliceToTrackerBehavior) error {
	if tracker == nil {
		return fmt.Errorf("tracker is nil")
	}
	if len(slice) < 1 {
		return nil
	}
	switch behavior {
	case RejectAllWhenAnyDuplicatedKey:
		for _, k := range slice {
			if _, exists := tracker[k]; exists {
				return fmt.Errorf("duplicated key %v, rejected all", k)
			}
		}
		for _, k := range slice {
			tracker[k] = defaultValue
		}
		return nil
	case SkipDuplicatedKeys:
		for _, k := range slice {
			if _, exists := tracker[k]; exists {
				continue
			}
			tracker[k] = defaultValue
		}
		return nil
	case AcceptAllAndOverrideDuplicatedKeys:
		for _, k := range slice {
			tracker[k] = defaultValue
		}
		return nil
	case AcceptOnlyDuplicatedKeysAndOverrideThem:
		for _, k := range slice {
			if _, exists := tracker[k]; exists {
				tracker[k] = defaultValue
			}
		}
		return nil
	default:
		return fmt.Errorf("not supported behavior %v", behavior)
	}
}
