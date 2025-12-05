package overflow

// Signed captures all Go signed integer types.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Add performs addition with overflow detection.
func Add[T Signed](left, right T) (T, bool) {
	result := left + right
	if (result > left) == (right > 0) {
		return result, true
	}
	return result, false
}

// Sub performs subtraction with overflow detection.
func Sub[T Signed](left, right T) (T, bool) {
	result := left - right
	if (result < left) == (right > 0) {
		return result, true
	}
	return result, false
}

// Mul performs multiplication with overflow detection.
func Mul[T Signed](left, right T) (T, bool) {
	if left == 0 || right == 0 {
		var zero T
		return zero, true
	}
	result := left * right
	if (result < 0) == ((left < 0) != (right < 0)) && result/right == left {
		return result, true
	}
	return result, false
}

// Div performs division with overflow detection (including division by zero and MinInt / -1 cases).
func Div[T Signed](left, right T) (T, bool) {
	var zero T
	if right == 0 {
		return zero, false
	}
	if right == -1 && left != 0 && left == -left {
		return left, false
	}
	result := left / right
	return result, (result < 0) == ((left < 0) != (right < 0))
}

// Mod performs modulo with overflow detection.
func Mod[T Signed](left, right T) (T, bool) {
	var zero T
	if right == 0 || (right == -1 && left != 0 && left == -left) {
		return zero, false
	}
	return left % right, true
}
