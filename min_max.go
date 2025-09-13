package gofunc

// Min finds the minimum value among the provided arguments.
// Returns an error if no arguments are provided.
// Works with numeric types and strings.
//
// Example:
//
//	min, err := gofunc.Min(3, 1, 4, 1, 5)
//	// min is 1, err is nil
func Min[T Number | ~string](s ...T) (T, error) {
	if len(s) == 0 {
		var zeroT T
		return zeroT, ErrInputRequired
	}
	min := s[0]
	for i := range s {
		if s[i] < min {
			min = s[i]
		}
	}
	return min, nil
}

// Max finds the maximum value among the provided arguments.
// Returns an error if no arguments are provided.
// Works with numeric types and strings.
//
// Example:
//
//	max, err := gofunc.Max(3, 1, 4, 1, 5)
//	// max is 5, err is nil
func Max[T Number | ~string](s ...T) (T, error) {
	if len(s) == 0 {
		var zeroT T
		return zeroT, ErrInputRequired
	}
	max := s[0]
	for i := range s {
		if s[i] > max {
			max = s[i]
		}
	}
	return max, nil
}
