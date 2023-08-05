package gofunc

import (
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Map_MapUpdate(t *testing.T) {
	// Update a nil map
	assert.Equal(t, map[uint8]string{1: "one", 2: "two"},
		MapUpdate(nil, map[uint8]string{1: "one", 2: "two"}))
	// Update a nil map with a nil map
	assert.Equal(t, map[uint8]string{}, MapUpdate[uint8, string](nil, nil))
	// Update a nil map with empty map
	assert.Equal(t, map[uint8]string{}, MapUpdate(nil, map[uint8]string{}))

	// Update with a nil map
	assert.Equal(t, map[int]bool{1: false, 2: true},
		MapUpdate(map[int]bool{1: false, 2: true}, nil))
	// Update with an empty map
	assert.Equal(t, map[int]bool{1: false, 2: true},
		MapUpdate(map[int]bool{1: false, 2: true}, map[int]bool{}))

	// Merge 2 maps
	assert.Equal(t, map[int]string{1: "one", 2: "two", 3: "three"},
		MapUpdate(map[int]string{2: "two"}, map[int]string{1: "one", 3: "three"}))
	// Merge 2 maps with override
	assert.Equal(t, map[int]string{1: "one", 2: "TWO", 3: "three"},
		MapUpdate(map[int]string{2: "two"}, map[int]string{1: "one", 2: "TWO", 3: "three"}))
}

func Test_Map_MapKeys(t *testing.T) {
	// Keys with an empty map
	assert.Equal(t, []int{}, MapKeys(map[int]bool{}))

	// Keys normal Map
	keys := MapKeys(map[int]string{1: "one", 2: "TWO", 3: "three"})
	sort.Ints(keys)
	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, keys))
}

func Test_Map_MapValues(t *testing.T) {
	// Values with an empty map
	assert.Equal(t, []int{}, MapValues(map[int]int{}))

	// Values normal Map
	values := MapValues(map[string]int{"one": 1, "two": 2, "three": 3})
	sort.Ints(values)
	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, values))
}

func Test_Map_MapSetDefault(t *testing.T) {
	// test with not existed key
	testMap := make(map[string]int, 10)
	val, _ := MapSetDefault(testMap, "key1", 100)
	assert.Equal(t, 100, val)

	// test with existed key
	testMap["key2"] = 200
	val2, _ := MapSetDefault(testMap, "key2", 100)
	assert.Equal(t, 200, val2)
}

func Test_Map_MapGet(t *testing.T) {
	// Get value of an empty map
	assert.Equal(t, 10, MapGet(map[string]int{}, "", 10))

	// Get value of normal map
	assert.Equal(t, 1, MapGet(map[string]int{"one": 1, "two": 2}, "one", 2))
}
