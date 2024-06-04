package gomaps

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
