package gomaps

import (
	"errors"

	"golang.org/x/exp/constraints"
)

var ErrSequenceContainsNoElements error = errors.New("sequence contains no elements")

type RealNumber interface {
	constraints.Integer | constraints.Float
}

type ComplexNumber interface {
	RealNumber | constraints.Complex
}

type Predicate[Key comparable, Value any] func(key Key, value Value) bool
type ValueSelector[Key comparable, Value any, T any] func(key Key, value Value) T

func Max[Dictionary ~map[Key]Value, Key comparable, Value any, Max constraints.Ordered](dictionary Dictionary, selector ValueSelector[Key, Value, Max]) (max Max, err error) {
	if len(dictionary) == 0 {
		return max, ErrSequenceContainsNoElements
	}

	for key, value := range dictionary {
		max = selector(key, value)
		break
	}

	if len(dictionary) == 1 {
		return max, nil
	}

	for key, value := range dictionary {
		value := selector(key, value)
		if value > max {
			max = value
		}
	}

	return max, nil
}

func Min[Dictionary ~map[Key]Value, Key comparable, Value any, Min constraints.Ordered](dictionary Dictionary, selector ValueSelector[Key, Value, Min]) (min Min, err error) {
	if len(dictionary) == 0 {
		return min, ErrSequenceContainsNoElements
	}

	for key, value := range dictionary {
		min = selector(key, value)
		break
	}

	if len(dictionary) == 1 {
		return min, nil
	}

	for key, value := range dictionary {
		value := selector(key, value)
		if value < min {
			min = value
		}
	}

	return min, nil
}

func Sum[Dictionary ~map[Key]Value, Key comparable, Value any, T ComplexNumber](dictionary Dictionary, selector ValueSelector[Key, Value, T]) (sum T) {
	for key, value := range dictionary {
		sum += selector(key, value)
	}
	return sum
}

func Where[Dictionary ~map[Key]Value, Key comparable, Value any](dictionary Dictionary, predicate Predicate[Key, Value]) Dictionary {
	m := make(Dictionary)
	for key, value := range dictionary {
		if predicate(key, value) {
			m[key] = value
		}
	}
	return m
}

func Any[Dictionary ~map[Key]Value, Key comparable, Value any](dictionary Dictionary, predicate Predicate[Key, Value]) bool {
	for key, value := range dictionary {
		if predicate(key, value) {
			return true
		}
	}
	return false
}

func All[Dictionary ~map[Key]Value, Key comparable, Value any](dictionary Dictionary, predicate Predicate[Key, Value]) bool {
	for key, value := range dictionary {
		if !predicate(key, value) {
			return false
		}
	}
	return true
}
