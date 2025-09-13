package gofunc

// ToInterfaceSlice converts a slice of any type to a slice of interface{}.
// This is useful when you need to work with heterogeneous data or when
// interfacing with APIs that expect []interface{}.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	interfaces := gofunc.ToInterfaceSlice(numbers)
//	// interfaces is []interface{}{1, 2, 3}
func ToInterfaceSlice[T any](slice []T) []interface{} {
	result := make([]interface{}, len(slice))
	for i := range slice {
		result[i] = slice[i]
	}
	return result
}

// ToStringSlice converts a slice of string-like types to a slice of string.
// This is useful when working with custom string types that need to be
// converted to standard strings.
//
// Example:
//
//	type UserID string
//	ids := []UserID{"user1", "user2"}
//	strings := gofunc.ToStringSlice(ids)
//	// strings is []string{"user1", "user2"}
func ToStringSlice[T ~string](slice []T) []string {
	result := make([]string, len(slice))
	for i := range slice {
		result[i] = string(slice[i])
	}
	return result
}

// ToStringDerivedSlice converts a string slice to a slice of string-derived type.
// This is the inverse of ToStringSlice, useful for converting standard strings
// back to custom string types.
//
// Example:
//
//	type UserID string
//	strings := []string{"user1", "user2"}
//	ids := gofunc.ToStringDerivedSlice[UserID](strings)
//	// ids is []UserID{"user1", "user2"}
func ToStringDerivedSlice[U, T ~string](slice []T) []U {
	result := make([]U, len(slice))
	for i := range slice {
		result[i] = U(slice[i])
	}
	return result
}

// ToNumberSlice converts a slice of number-like types to another numeric type.
// This is useful when working with custom numeric types that need conversion.
//
// Example:
//
//	type OrderID uint
//	orders := []OrderID{1, 2, 3}
//	numbers := gofunc.ToNumberSlice[uint](orders)
//	// numbers is []uint{1, 2, 3}
func ToNumberSlice[U, T Number](slice []T) []U {
	result := make([]U, len(slice))
	for i := range slice {
		result[i] = U(slice[i])
	}
	return result
}

// ToSet removes duplicate elements from a slice, returning a new slice
// with unique elements in their first occurrence order.
// The input slice is not modified.
//
// Example:
//
//	numbers := []int{1, 2, 2, 3, 1}
//	unique := gofunc.ToSet(numbers)
//	// unique is []int{1, 2, 3}
func ToSet[T comparable](s []T) []T {
	var (
		length = len(s)
		seen   = make(map[T]struct{}, length)
		result = make([]T, 0, length)
	)

	for i := 0; i < length; i++ {
		v := s[i]
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		result = append(result, v)
	}

	return result
}

// ToSetPred removes duplicate elements from a slice using a key function
// to determine uniqueness. Elements are considered equal if their keys are equal.
// Returns a new slice with unique elements in their first occurrence order.
//
// Example:
//
//	type Person struct { Name string; Age int }
//	people := []Person{{"Alice", 30}, {"Bob", 25}, {"Alice", 31}}
//	unique := gofunc.ToSetPred(people, func(p Person) string { return p.Name })
//	// unique contains only first Alice and Bob
func ToSetPred[T any, K comparable](s []T, keyFunc func(t T) K) []T {
	var (
		length = len(s)
		seen   = make(map[K]struct{}, length)
		result = make([]T, 0, length)
	)

	for i := 0; i < length; i++ {
		v := s[i]
		k := keyFunc(v)
		if _, ok := seen[k]; ok {
			continue
		}
		seen[k] = struct{}{}
		result = append(result, v)
	}

	return result
}

// Contains checks if a slice contains a specific item.
// Returns true if the item is found, false otherwise.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	found := gofunc.Contains(numbers, 3)
//	// found is true
func Contains[T comparable](a []T, b T) bool {
	for i := range a {
		if a[i] == b {
			return true
		}
	}
	return false
}

// ContainsPred checks if a slice contains an item matching the given predicate.
// Returns true if any element satisfies the predicate, false otherwise.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	hasEven := gofunc.ContainsPred(numbers, func(n int) bool { return n%2 == 0 })
//	// hasEven is true
func ContainsPred[T any](a []T, pred func(b T) bool) bool {
	for i := range a {
		if pred(a[i]) {
			return true
		}
	}
	return false
}

// FindPred finds the first item in a slice that matches the given predicate.
// Returns the found item and true, or zero value and false if not found.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	even, found := gofunc.FindPred(numbers, func(n int) bool { return n%2 == 0 })
//	// even is 2, found is true
func FindPred[T any](a []T, pred func(b T) bool) (T, bool) {
	for i := range a {
		if pred(a[i]) {
			return a[i], true
		}
	}
	var zeroT T
	return zeroT, false
}

// IndexOf returns the index of the first occurrence of an item in a slice.
// Returns -1 if the item is not found.
//
// Example:
//
//	numbers := []int{1, 2, 3, 2, 4}
//	index := gofunc.IndexOf(numbers, 2)
//	// index is 1 (first occurrence)
func IndexOf[T comparable](a []T, t T) int {
	for i := range a {
		if a[i] == t {
			return i
		}
	}
	return -1
}

// LastIndexOf returns the index of the last occurrence of an item in a slice.
// Returns -1 if the item is not found.
//
// Example:
//
//	numbers := []int{1, 2, 3, 2, 4}
//	index := gofunc.LastIndexOf(numbers, 2)
//	// index is 3 (last occurrence)
func LastIndexOf[T comparable](a []T, t T) int {
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == t {
			return i
		}
	}
	return -1
}

// IndexOfSlice returns the index of the first occurrence of a sub-slice within a slice.
// Returns -1 if the sub-slice is not found or if sub-slice is empty.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	pattern := []int{3, 4}
//	index := gofunc.IndexOfSlice(numbers, pattern)
//	// index is 2
func IndexOfSlice[T comparable](a []T, sub []T) int {
	lengthA := len(a)
	lengthSub := len(sub)
	if lengthSub == 0 || lengthA < lengthSub {
		return -1
	}
	sub1st := sub[0]
	for i, max := 0, lengthA-lengthSub; i <= max; i++ {
		if a[i] == sub1st {
			found := true
			for j := 1; j < lengthSub; j++ {
				if a[i+j] != sub[j] {
					found = false
					break
				}
			}
			if found {
				return i
			}
		}
	}
	return -1
}

// ChunkSlice splits a slice into smaller slices of specified size.
// The last chunk may be smaller if the slice length is not evenly divisible.
// Returns an empty slice if the input slice is empty.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5, 6, 7}
//	chunks := gofunc.ChunkSlice(numbers, 3)
//	// chunks is [][]int{{1, 2, 3}, {4, 5, 6}, {7}}
func ChunkSlice[T any](slice []T, chunkSize int) [][]T {
	total := len(slice)
	if total == 0 {
		return [][]T{}
	}
	if total <= chunkSize {
		return [][]T{slice}
	}

	chunks := make([][]T, 0, total/chunkSize+1)
	for {
		if len(slice) == 0 {
			break
		}

		// necessary check to avoid slicing beyond slice capacity
		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}

		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}

	return chunks
}

// ConcatSlices concatenates multiple slices into a single slice.
// The resulting slice contains all elements from the input slices in order.
// Returns an empty slice if no slices are provided.
//
// Example:
//
//	slice1 := []int{1, 2}
//	slice2 := []int{3, 4}
//	slice3 := []int{5}
//	result := gofunc.ConcatSlices(slice1, slice2, slice3)
//	// result is []int{1, 2, 3, 4, 5}
func ConcatSlices[T any](slices ...[]T) []T {
	capacity := 0
	for _, s := range slices {
		capacity += len(s)
	}
	result := make([]T, 0, capacity)
	for _, s := range slices {
		result = append(result, s...)
	}
	return result
}
