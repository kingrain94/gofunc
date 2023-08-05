package gofunc

// Min find the minimum value in the list
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

// Max find the maximum value in the list
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
