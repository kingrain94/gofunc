package gofunc

// MapUpdate update map content with another map
// Not change the target map, only change the source map
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

// MapKeys get map keys as slice
func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

// MapValues get map values as slice
func MapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))
	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
}

// MapGet get the value for the key, if not exist, return the default one
func MapGet[K comparable, V any](m map[K]V, k K, defaultVal V) V {
	if val, ok := m[k]; ok {
		return val
	}
	return defaultVal
}

// MapSetDefault set default value for a key
// If the key exists, return `existing_value, true`, otherwise `default_value, false`.
func MapSetDefault[K comparable, V any](m map[K]V, k K, defaultVal V) (V, bool) {
	if val, ok := m[k]; ok {
		return val, true
	} else {
		m[k] = defaultVal
		return defaultVal, false
	}
}
