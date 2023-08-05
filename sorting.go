package gofunc

import "sort"

// Sort sort slice values in ascending order
func Sort[T Number | ~string](s []T) []T {
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	return s
}

// SortPred sort slice values in ascending order by predicate
func SortPred[T any, K Number | ~string](s []T, keyFunc func(t T) K) []T {
	sort.Slice(s, func(i, j int) bool {
		x := keyFunc(s[i])
		y := keyFunc(s[j])
		return x < y
	})
	return s
}
