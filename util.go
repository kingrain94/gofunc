package gofunc

// Must returns the value if error is nil, otherwise panics with the error.
// This is useful for cases where you want to convert error-returning functions
// into panic-on-error functions, typically during initialization.
//
// Example:
//
//	value := gofunc.Must(strconv.Atoi("42"))
//	// value is 42, or panics if conversion fails
func Must[T any](v T, e error) T {
	if e != nil {
		panic(e)
	}

	return v
}
