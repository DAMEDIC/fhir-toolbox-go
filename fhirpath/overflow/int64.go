// Package overflow provides arithmetic operations for int64 values with overflow detection.
package overflow

import "math"

// Add64 adds two int64 values and returns the result along with a flag indicating
// whether the operation completed without overflow.
func Add64(left, right int64) (int64, bool) {
	result := left + right
	if (result > left) == (right > 0) {
		return result, true
	}
	return result, false
}

// Sub64 subtracts right from left and returns the result along with a flag indicating
// whether the operation completed without overflow.
func Sub64(left, right int64) (int64, bool) {
	result := left - right
	if (result < left) == (right > 0) {
		return result, true
	}
	return result, false
}

// Mul64 multiplies two int64 values and returns the result along with a flag indicating
// whether the operation completed without overflow.
func Mul64(left, right int64) (int64, bool) {
	if left == 0 || right == 0 {
		return 0, true
	}
	result := left * right
	if (result < 0) == ((left < 0) != (right < 0)) {
		if result/right == left {
			return result, true
		}
	}
	return result, false
}

// Div64 performs integer division of two int64 values and returns the result along with
// a flag indicating whether the operation completed without overflow or division-by-zero.
func Div64(left, right int64) (int64, bool) {
	if right == 0 {
		return 0, false
	}
	if left == math.MinInt64 && right == -1 {
		return math.MinInt64, false
	}
	result := left / right
	return result, (result < 0) == ((left < 0) != (right < 0))
}

// Mod64 performs modulo of two int64 values and returns the result along with a flag
// indicating whether the operation completed without overflow or division-by-zero.
func Mod64(left, right int64) (int64, bool) {
	if right == 0 || (left == math.MinInt64 && right == -1) {
		return 0, false
	}
	return left % right, true
}
