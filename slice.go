package gofunc

// ToInterfaceSlice convert a slice to a slice of interface
func ToInterfaceSlice[T any](slice []T) []interface{} {
	result := make([]interface{}, len(slice))
	for i := range slice {
		result[i] = slice[i]
	}
	return result
}

// ToStringSlice convert a slice of str-approximate-type to a slice of string
func ToStringSlice[T ~string](slice []T) []string {
	result := make([]string, len(slice))
	for i := range slice {
		result[i] = string(slice[i])
	}
	return result
}

// ToStringDerivedSlice convert a string slice to a slice of string-approximate-type
func ToStringDerivedSlice[U, T ~string](slice []T) []U {
	result := make([]U, len(slice))
	for i := range slice {
		result[i] = U(slice[i])
	}
	return result
}

// ToNumberSlice convert a slice of number-approximate type to a slice of number
// E.g. type OrderID uint, ToNumberSlice[uint]([]OrderID) -> []uint
func ToNumberSlice[U, T Number](slice []T) []U {
	result := make([]U, len(slice))
	for i := range slice {
		result[i] = U(slice[i])
	}
	return result
}

// ToSet convert a slice to a set (no duplicated items)
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

// ToSetPred convert a slice to a set with support key function
// Unlike ToSet(), ToSetPred() can convert a slice of any type
// ToSet() is the special case of ToSetPred() when keyFunc is 'func(t T) T { return t }'
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

// Contains test if a slice contains an item
func Contains[T comparable](a []T, b T) bool {
	for i := range a {
		if a[i] == b {
			return true
		}
	}
	return false
}

// ContainsPred test if a slice contains an item by predicate
func ContainsPred[T any](a []T, pred func(b T) bool) bool {
	for i := range a {
		if pred(a[i]) {
			return true
		}
	}
	return false
}

// FindPred find an item in slice by predicate
func FindPred[T any](a []T, pred func(b T) bool) (T, bool) {
	for i := range a {
		if pred(a[i]) {
			return a[i], true
		}
	}
	var zeroT T
	return zeroT, false
}

// IndexOf get index of item in slice
// Return -1 if not found
func IndexOf[T comparable](a []T, t T) int {
	for i := range a {
		if a[i] == t {
			return i
		}
	}
	return -1
}

// LastIndexOf get index of item from the end in slice
// Return -1 if not found
func LastIndexOf[T comparable](a []T, t T) int {
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == t {
			return i
		}
	}
	return -1
}

// IndexOfSlice get index of sub-slice in slice
// Return -1 if not found
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

// ChunkSlice split a slice into some slices by chunkSize
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

// ConcatSlices merge multiple slices into one
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
