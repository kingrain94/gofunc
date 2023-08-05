package gofunc

func If[T any](cond bool, a T, b T) T {
	if cond {
		return a
	}
	return b
}
