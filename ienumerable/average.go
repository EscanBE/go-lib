package ienumerable

import "fmt"

func (e IEnumerable[T]) Average(selector func(t T) float64) (float64, error) {
	if len(e.data) < 1 {
		return 0, fmt.Errorf("sequence contains no elements")
	}
	var sum float64 = 0
	for _, t := range e.data {
		selected := selector(t)
		sum = sum + selected
	}
	return sum / float64(len(e.data)), nil
}

func (e IEnumerable[T]) AverageNullable(selector func(t T) *float64) (*float64, error) {
	if len(e.data) < 1 {
		return nil, fmt.Errorf("sequence contains no elements")
	}
	var sumNonNil float64 = 0
	var nonNilCount int = 0
	for _, t := range e.data {
		selected := selector(t)
		if selected != nil {
			sumNonNil = sumNonNil + *selected
			nonNilCount++
		}
	}
	if nonNilCount < 1 {
		return nil, nil
	}
	avg := sumNonNil / float64(nonNilCount)
	return &avg, nil
}

func (e IEnumerable[T]) AverageInt(selector func(t T) int) (float64, error) {
	if len(e.data) < 1 {
		return 0, fmt.Errorf("sequence contains no elements")
	}
	var sum = 0
	for _, t := range e.data {
		selected := selector(t)
		sum = sum + selected
	}
	return float64(sum) / float64(len(e.data)), nil
}

func (e IEnumerable[T]) AverageInt64(selector func(t T) int64) (float64, error) {
	if len(e.data) < 1 {
		return 0, fmt.Errorf("sequence contains no elements")
	}
	var sum int64 = 0
	for _, t := range e.data {
		selected := selector(t)
		sum = sum + selected
	}
	return float64(sum) / float64(len(e.data)), nil
}
