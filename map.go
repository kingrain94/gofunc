package gofunc

// MapUpdate merges two maps, with values from m2 overriding values in m1.
// The first map (m1) is modified and returned. If m1 is nil, a new map is created.
// If m2 is nil, m1 is returned unchanged.
//
// Example:
//
//	defaults := map[string]int{"a": 1, "b": 2}
//	overrides := map[string]int{"b": 3, "c": 4}
//	result := gofunc.MapUpdate(defaults, overrides)
//	// result is map[string]int{"a": 1, "b": 3, "c": 4}
func MapUpdate[K comparable, V any](m1, m2 map[K]V) map[K]V {
	if m1 == nil {
		m1 = make(map[K]V, len(m2))
	}
	if m2 == nil {
		return m1
	}
	for k, v := range m2 {
		m1[k] = v
	}
	return m1
}

// MapKeys returns all keys from a map as a slice.
// The order of keys is not guaranteed to be consistent between calls.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	keys := gofunc.MapKeys(m)
//	// keys contains ["a", "b", "c"] in some order
func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

// MapValues returns all values from a map as a slice.
// The order of values is not guaranteed to be consistent between calls.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	values := gofunc.MapValues(m)
//	// values contains [1, 2, 3] in some order
func MapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))
	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
}

// MapGet retrieves a value from a map by key, returning a default value if the key doesn't exist.
// This is useful for safe map access without checking existence separately.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2}
//	value := gofunc.MapGet(m, "c", 999)
//	// value is 999 (default) since "c" doesn't exist
func MapGet[K comparable, V any](m map[K]V, k K, defaultVal V) V {
	if val, ok := m[k]; ok {
		return val
	}
	return defaultVal
}

// MapSetDefault sets a default value for a key if it doesn't exist in the map.
// Returns the existing value and true if the key exists, or the default value and false if it was set.
// The map is modified if the key doesn't exist.
//
// Example:
//
//	m := map[string]int{"a": 1}
//	val, existed := gofunc.MapSetDefault(m, "b", 2)
//	// val is 2, existed is false, m now contains {"a": 1, "b": 2}
func MapSetDefault[K comparable, V any](m map[K]V, k K, defaultVal V) (V, bool) {
	if val, ok := m[k]; ok {
		return val, true
	} else {
		m[k] = defaultVal
		return defaultVal, false
	}
}
