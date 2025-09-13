package gofunc

import "sort"

// Sort sorts a slice in ascending order and returns the modified slice.
// The original slice is modified in place.
// Works with numeric types and strings.
//
// Example:
//
//	numbers := []int{3, 1, 4, 1, 5}
//	sorted := gofunc.Sort(numbers)
//	// sorted and numbers are both []int{1, 1, 3, 4, 5}
func Sort[T Number | ~string](s []T) []T {
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	return s
}

// SortPred sorts a slice in ascending order using a key function and returns the modified slice.
// The original slice is modified in place. The key function extracts a comparable value from each element.
//
// Example:
//
//	type Person struct { Name string; Age int }
//	people := []Person{{"Bob", 30}, {"Alice", 25}}
//	sorted := gofunc.SortPred(people, func(p Person) string { return p.Name })
//	// sorted by name: [{"Alice", 25}, {"Bob", 30}]
func SortPred[T any, K Number | ~string](s []T, keyFunc func(t T) K) []T {
	sort.Slice(s, func(i, j int) bool {
		x := keyFunc(s[i])
		y := keyFunc(s[j])
		return x < y
	})
	return s
}
