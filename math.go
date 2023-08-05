package gofunc

import "math"

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
