package gofunc

// If returns the first value if the condition is true, otherwise returns the second value.
// This provides a ternary-like operator for Go, useful for simple conditional expressions.
//
// Example:
//
//	result := gofunc.If(score >= 90, "A", "B")
//	// result is "A" if score >= 90, otherwise "B"
func If[T any](cond bool, a T, b T) T {
	if cond {
		return a
	}
	return b
}
