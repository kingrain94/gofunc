package gofunc

import "math"

// Abs returns the absolute value of an int64.
// Panics if x is math.MinInt64 since it cannot be represented as a positive int64.
//
// Example:
//
//	result := gofunc.Abs(-42)
//	// result is 42
func Abs(x int64) int64 {
	if x < 0 {
		// The lowest value of int64 has no corresponding positive value
		if x == math.MinInt64 {
			panic("unable to calculate abs of the lowest int64 value")
		}
		return -x
	}
	return x
}
