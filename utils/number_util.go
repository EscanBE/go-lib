package utils

// ConditionalInt64 returns int64 based on input expression
func ConditionalInt64(expression bool, whenTrue, whenFalse int64) int64 {
	if expression {
		return whenTrue
	} else {
		return whenFalse
	}
}

// ConditionalInt returns int based on input expression
func ConditionalInt(expression bool, whenTrue, whenFalse int) int {
	if expression {
		return whenTrue
	} else {
		return whenFalse
	}
}
