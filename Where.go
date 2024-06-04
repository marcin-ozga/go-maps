package gomaps

import (
	"errors"

	"golang.org/x/exp/constraints"
)

var ErrSequenceContainsNoElements error = errors.New("sequence contains no elements")

func Max[Key comparable, Value any, Max constraints.Ordered](dictionary map[Key]Value, maxSelector func(key Key, value Value) Max) (max Max, err error) {
	if len(dictionary) == 0 {
		return max, ErrSequenceContainsNoElements
	}

	for key, value := range dictionary {
		max = maxSelector(key, value)
		break
	}

	if len(dictionary) == 1 {
		return max, nil
	}

	for key, value := range dictionary {
		value := maxSelector(key, value)
		if value > max {
			max = value
		}
	}

	return max, nil
}

func Min[Key comparable, Value any, Min constraints.Ordered](dictionary map[Key]Value, maxSelector func(key Key, value Value) Min) (min Min, err error) {
	if len(dictionary) == 0 {
		return min, ErrSequenceContainsNoElements
	}

	for key, value := range dictionary {
		min = maxSelector(key, value)
		break
	}

	if len(dictionary) == 1 {
		return min, nil
	}

	for key, value := range dictionary {
		value := maxSelector(key, value)
		if value < min {
			min = value
		}
	}

	return min, nil
}

func Sum[Key comparable, Value any, ReturnValue constraints.Integer | constraints.Float | constraints.Complex](dictionary map[Key]Value, valueSelector func(key Key, value Value) ReturnValue) (sum ReturnValue) {
	for key, value := range dictionary {
		sum += valueSelector(key, value)
	}
	return sum
}

func Where[Dictionary map[Key]Value, Key comparable, Value any](dictionary Dictionary, keyValueSelector func(key Key, value Value) bool) Dictionary {
	m := make(Dictionary)
	for key, value := range dictionary {
		if keyValueSelector(key, value) {
			m[key] = value
		}
	}
	return m
}

func Any[Key comparable, Value any](dictionary map[Key]Value, keyValueSelector func(key Key, value Value) bool) bool {
	for key, value := range dictionary {
		if keyValueSelector(key, value) {
			return true
		}
	}
	return false
}

func All[Key comparable, Value any](dictionary map[Key]Value, keyValueSelector func(key Key, value Value) bool) bool {
	for key, value := range dictionary {
		if !keyValueSelector(key, value) {
			return false
		}
	}
	return true
}
