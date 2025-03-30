// Package overflow provides arithmetic operations for int32 values with overflow detection.
// Each operation returns both the result and a boolean indicating whether the operation
// was successful (true) or resulted in an overflow (false).
package overflow

import "math"

// Add32 performs addition of two int32 values with overflow detection.
// It returns the result of left + right and a boolean indicating whether the operation
// was successful (true) or resulted in an overflow (false).
//
// Example:
//
//	sum, ok := Add32(2147483647, 1)  // Returns MinInt32, false (overflow)
//	sum, ok := Add32(1, 1)           // Returns 2, true (no overflow)
func Add32(left, right int32) (int32, bool) {
	result := left + right
	if (result > left) == (right > 0) {
		return result, true
	}
	return result, false
}

// Sub32 performs subtraction of two int32 values with overflow detection.
// It returns the result of left - right and a boolean indicating whether the operation
// was successful (true) or resulted in an overflow (false).
//
// Example:
//
//	diff, ok := Sub32(-2147483648, 1)  // Returns MaxInt32, false (overflow)
//	diff, ok := Sub32(5, 3)            // Returns 2, true (no overflow)
func Sub32(left, right int32) (int32, bool) {
	result := left - right
	if (result < left) == (right > 0) {
		return result, true
	}
	return result, false
}

// Mul32 performs multiplication of two int32 values with overflow detection.
// It returns the result of left * right and a boolean indicating whether the operation
// was successful (true) or resulted in an overflow (false).
// Special cases:
//   - If either operand is 0, returns (0, true)
//   - For non-zero operands, checks if the result would overflow int32 range
//
// Example:
//
//	prod, ok := Mul32(2147483647, 2)  // Returns -2, false (overflow)
//	prod, ok := Mul32(5, 3)           // Returns 15, true (no overflow)
//	prod, ok := Mul32(0, 5)           // Returns 0, true (zero case)
func Mul32(left, right int32) (int32, bool) {
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

// Div32 performs division of two int32 values with overflow detection.
// It returns the result of left / right and a boolean indicating whether the operation
// was successful (true) or resulted in an invalid operation (false).
// Special cases:
//   - If right is 0, returns (0, false) to indicate division by zero
//   - If left is MinInt32 and right is -1, returns (MinInt32, false) to prevent overflow
//
// Example:
//
//	quot, ok := Div32(10, 2)           // Returns 5, true
//	quot, ok := Div32(5, 0)            // Returns 0, false (division by zero)
//	quot, ok := Div32(-2147483648, -1) // Returns -2147483648, false (overflow)
func Div32(left, right int32) (int32, bool) {
	if right == 0 {
		return 0, false
	}
	if left == math.MinInt32 && right == -1 {
		return math.MinInt32, false
	}
	result := left / right
	return result, (result < 0) == ((left < 0) != (right < 0))
}

// Mod32 performs modulo operation of two int32 values with overflow detection.
// It returns the result of left % right and a boolean indicating whether the operation
// was successful (true) or resulted in an invalid operation (false).
// Special cases:
//   - If right is 0, returns (0, false) to indicate division by zero
//   - If left is MinInt32 and right is -1, returns (0, false) to prevent overflow
//
// Example:
//
//	rem, ok := Mod32(10, 3)           // Returns 1, true
//	rem, ok := Mod32(5, 0)            // Returns 0, false (division by zero)
//	rem, ok := Mod32(-2147483648, -1) // Returns 0, false (overflow)
func Mod32(left, right int32) (int32, bool) {
	if right == 0 || (left == math.MinInt32 && right == -1) {
		return 0, false
	}
	return left % right, true
}
