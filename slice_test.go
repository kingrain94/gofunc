package gofunc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Slice_ToInterfaceSlice(t *testing.T) {
	assert.Equal(t, []interface{}{}, ToInterfaceSlice([]int{}))
	assert.Equal(t, []interface{}{"one"}, ToInterfaceSlice([]string{"one"}))
	assert.Equal(t, []interface{}{1, 2, 3}, ToInterfaceSlice([]int{1, 2, 3}))
	assert.Equal(t, []interface{}{float32(1.1), float32(2.2), float32(3.3)}, ToInterfaceSlice([]float32{1.1, 2.2, 3.3}))
	assert.Equal(t, []interface{}{"one", 2, 3.3}, ToInterfaceSlice([]interface{}{"one", 2, 3.3}))
}

func Test_Slice_ToStringSlice(t *testing.T) {
	type TypeX string
	assert.Equal(t, []string{}, ToStringSlice([]TypeX{}))
	assert.Equal(t, []string{"one", "two"}, ToStringSlice([]TypeX{"one", "two"}))
}

func Test_Slice_ToStringDerivedSlice(t *testing.T) {
	type TypeX string
	assert.Equal(t, []TypeX{}, ToStringDerivedSlice[TypeX]([]string{}))
	assert.Equal(t, []TypeX{"one", "two"}, ToStringDerivedSlice[TypeX]([]string{"one", "two"}))
}

func Test_Slice_ToNumberSlice(t *testing.T) {
	type TypeX int8
	assert.Equal(t, []int8{}, ToNumberSlice[int8]([]TypeX{}))
	assert.Equal(t, []int8{1, 2, 3}, ToNumberSlice[int8]([]TypeX{1, 2, 3}))

	type TypeY uint32
	assert.Equal(t, []uint32{}, ToNumberSlice[uint32]([]TypeY{}))
	assert.Equal(t, []uint32{10, 20}, ToNumberSlice[uint32]([]TypeY{10, 20}))

	type TypeZ float64
	assert.Equal(t, []float64{}, ToNumberSlice[float64]([]TypeZ{}))
	assert.Equal(t, []float64{-1.23, 0, 3.45}, ToNumberSlice[float64]([]TypeZ{-1.23, 0, 3.45}))
}

func Test_Slice_ToSet(t *testing.T) {
	assert.Equal(t, []int{}, ToSet([]int{}))
	assert.Equal(t, []string{"one"}, ToSet([]string{"one"}))
	assert.Equal(t, []string{"one", "two", "Two"}, ToSet([]string{"one", "two", "one", "Two"}))
	assert.Equal(t, []int{1, 2, 3}, ToSet([]int{1, 2, 3, 1, 2}))
	assert.Equal(t, []float32{float32(1.1), float32(2.2), float32(3.3), float32(1.11)},
		ToSet([]float32{1.1, 1.1, 2.2, 3.3, 1.11}))

	type st struct {
		I int
		S string
	}
	assert.Equal(t, []st{{1, "one"}, {2, "two"}}, ToSet([]st{{1, "one"}, {2, "two"}, {1, "one"}}))
}

func Test_Slice_ToSetPred(t *testing.T) {
	// Comparable types
	assert.Equal(t, []int{}, ToSetPred([]int{}, func(t int) int { return t }))
	assert.Equal(t, []string{"one"}, ToSetPred([]string{"one"}, func(t string) string { return t }))
	assert.Equal(t, []string{"one", "two", "Two"},
		ToSetPred([]string{"one", "two", "one", "Two"}, func(t string) string { return t }))
	assert.Equal(t, []int{1, 2, 3}, ToSetPred([]int{1, 2, 3, 1, 2}, func(t int) int { return t }))
	assert.Equal(t, []float32{float32(1.1), float32(2.2), float32(3.3), float32(1.11)},
		ToSetPred([]float32{1.1, 1.1, 2.2, 3.3, 1.11}, func(t float32) float32 { return t }))

	// Incomparable types
	assert.Equal(t, []interface{}{},
		ToSetPred([]interface{}{}, func(t interface{}) int { return t.(int) }))
	assert.Equal(t, []interface{}{"one"},
		ToSetPred([]interface{}{"one"}, func(t interface{}) string { return t.(string) }))
	assert.Equal(t, []interface{}{"one", "two", "Two"},
		ToSetPred([]interface{}{"one", "two", "one", "Two"}, func(t interface{}) string { return t.(string) }))
	assert.Equal(t, []interface{}{1, 2, 3},
		ToSetPred([]interface{}{1, 2, 3, 1, 2}, func(t interface{}) int { return t.(int) }))
	assert.Equal(t, []interface{}{1.1, 2.2, 3.3, 1.11},
		ToSetPred([]interface{}{1.1, 1.1, 2.2, 3.3, 1.11}, func(t interface{}) float64 { return t.(float64) }))
}

func Test_Slice_Contains(t *testing.T) {
	assert.False(t, Contains([]int{}, 1))
	assert.False(t, Contains([]string{"one"}, "One"))
	assert.False(t, Contains([]string{"one", "two"}, ""))
	assert.False(t, Contains([]int64{1, 2, 3}, 4))
	assert.False(t, Contains([]float32{1.1, 2.2, 3.3}, 3.35))

	assert.True(t, Contains([]int64{1}, 1))
	assert.True(t, Contains([]uint{1, 2, 3, 1, 2, 3}, 2))
	assert.True(t, Contains([]string{"one", "two"}, "two"))
	assert.True(t, Contains([]string{"one", "two", ""}, ""))
	assert.True(t, Contains([]float64{1.1, 2.2, 3.3}, 2.2))
}

func Test_Slice_ContainsPred(t *testing.T) {
	assert.False(t, ContainsPred([]int{}, func(i int) bool { return i == 1 }))
	assert.False(t, ContainsPred([]string{"one"}, func(i string) bool { return i == "One" }))
	assert.False(t, ContainsPred([]string{"one", "two"}, func(i string) bool { return i == "" }))
	assert.False(t, ContainsPred([]int64{1, 2, 3}, func(i int64) bool { return i == 4 }))
	assert.False(t, ContainsPred([]float32{1.1, 2.2, 3.3}, func(i float32) bool { return i == 3.35 }))

	assert.True(t, ContainsPred([]int64{1}, func(i int64) bool { return i == 1 }))
	assert.True(t, ContainsPred([]uint{1, 2, 3, 1, 2, 3}, func(i uint) bool { return i == 2 }))
	assert.True(t, ContainsPred([]string{"one", "two"}, func(i string) bool { return i == "two" }))
	assert.True(t, ContainsPred([]string{"one", "two", ""}, func(i string) bool { return i == "" }))
	assert.True(t, ContainsPred([]float64{1.1, 2.2, 3.3}, func(i float64) bool { return i == 2.2 }))
}

func Test_Slice_FindPred(t *testing.T) {
	itemInt, exists := FindPred([]int{}, func(i int) bool { return i == 1 })
	assert.True(t, itemInt == 0 && exists == false)
	itemStr, exists := FindPred([]string{"one"}, func(i string) bool { return i == "One" })
	assert.True(t, itemStr == "" && exists == false)
	itemStr, exists = FindPred([]string{"one", "two"}, func(i string) bool { return i == "" })
	assert.True(t, itemStr == "" && exists == false)
	itemInt, exists = FindPred([]int{}, func(i int) bool { return i == 1 })
	assert.True(t, itemInt == 0 && exists == false)
	itemFloat32, exists := FindPred([]float32{1.1, 2.2, 3.3}, func(i float32) bool { return i == 3.35 })
	assert.True(t, itemFloat32 == float32(0) && exists == false)

	itemInt64, exists := FindPred([]int64{1}, func(i int64) bool { return i == 1 })
	assert.True(t, itemInt64 == 1 && exists == true)
	itemUint, exists := FindPred([]uint{1, 2, 3, 1, 2, 3}, func(i uint) bool { return i == 2 })
	assert.True(t, itemUint == 2 && exists == true)
	itemStr, exists = FindPred([]string{"one", "two"}, func(i string) bool { return i == "one" })
	assert.True(t, itemStr == "one" && exists == true)
	itemStr, exists = FindPred([]string{"one", "two", ""}, func(i string) bool { return i == "" })
	assert.True(t, itemStr == "" && exists == true)
	itemFloat64, exists := FindPred([]float64{1.1, 2.2, 3.3}, func(i float64) bool { return i == 3.3 })
	assert.True(t, itemFloat64 == 3.3 && exists == true)
}

func Test_Slice_IndexOf(t *testing.T) {
	assert.Equal(t, -1, IndexOf([]int{}, 1))
	assert.Equal(t, -1, IndexOf([]string{"one"}, "One"))
	assert.Equal(t, -1, IndexOf([]string{"one", "two"}, ""))
	assert.Equal(t, -1, IndexOf([]int64{1, 2, 3}, 4))
	assert.Equal(t, -1, IndexOf([]float32{1.1, 2.2, 3.3}, 3.35))

	assert.Equal(t, 0, IndexOf([]int64{1}, 1))
	assert.Equal(t, 1, IndexOf([]uint{1, 2, 3, 1, 2, 3}, 2))
	assert.Equal(t, 1, IndexOf([]string{"one", "two"}, "two"))
	assert.Equal(t, 2, IndexOf([]string{"one", "two", ""}, ""))
	assert.Equal(t, 2, IndexOf([]float64{1.1, 2.2, 3.3}, 3.3))
}

func Test_Slice_LastIndexOf(t *testing.T) {
	assert.Equal(t, -1, LastIndexOf([]int{}, 1))
	assert.Equal(t, -1, LastIndexOf([]string{"one"}, "One"))
	assert.Equal(t, -1, LastIndexOf([]string{"one", "two"}, ""))
	assert.Equal(t, -1, LastIndexOf([]int64{1, 2, 3}, 4))
	assert.Equal(t, -1, LastIndexOf([]float32{1.1, 2.2, 3.3}, 3.35))

	assert.Equal(t, 0, LastIndexOf([]int64{1}, 1))
	assert.Equal(t, 4, LastIndexOf([]uint{1, 2, 3, 1, 2, 3}, 2))
	assert.Equal(t, 1, LastIndexOf([]string{"one", "two"}, "two"))
	assert.Equal(t, 3, LastIndexOf([]string{"one", "", "two", ""}, ""))
	assert.Equal(t, 0, LastIndexOf([]float64{1.1, 2.2, 3.3}, 1.1))
}

func Test_Slice_IndexOfSlice(t *testing.T) {
	assert.Equal(t, -1, IndexOfSlice([]int{}, nil))
	assert.Equal(t, -1, IndexOfSlice([]string{"one"}, []string{}))
	assert.Equal(t, -1, IndexOfSlice([]string{"one", "two"}, []string{"Two"}))
	assert.Equal(t, -1, IndexOfSlice([]int64{1, 2, 3}, []int64{1, 2, 3, 4}))
	assert.Equal(t, -1, IndexOfSlice([]uint{0, 1, 2, 3, 4, 5}, []uint{3, 4, 5, 6}))
	assert.Equal(t, -1, IndexOfSlice([]float32{1.1, 2.2, 3.3}, []float32{2.2, 3.31}))

	assert.Equal(t, 0, IndexOfSlice([]int{1}, []int{1}))
	assert.Equal(t, 2, IndexOfSlice([]int{0, 1, 2}, []int{2}))
	assert.Equal(t, 0, IndexOfSlice([]int{0, 1, 2, 0, 1, 2, 3}, []int{0, 1, 2}))
	assert.Equal(t, 1, IndexOfSlice([]string{"one", ""}, []string{""}))
	assert.Equal(t, 0, IndexOfSlice([]string{"one", "two", "three"}, []string{"one", "two"}))
	assert.Equal(t, 0, IndexOfSlice([]int64{1, 2, 3}, []int64{1, 2, 3}))
	assert.Equal(t, 1, IndexOfSlice([]uint{0, 1, 1, 1, 1}, []uint{1}))
}

func Test_Slice_ChunkSlice(t *testing.T) {
	strSlice := make([]string, 5)
	chunkStr := ChunkSlice(strSlice, 2)
	// This will be chunked into 3 slices, with length is 2, 2, 1
	assert.True(t, len(chunkStr) == 3 && len(chunkStr[0]) == 2 &&
		len(chunkStr[1]) == 2 && len(chunkStr[2]) == 1)

	intSlice := make([]int, 10000)
	chunkInt := ChunkSlice(intSlice, 3000)
	// This will be chunked into 4 slices, with length is 3000, 3000, 3000, 1000
	assert.True(t, len(chunkInt) == 4 && len(chunkInt[0]) == 3000 &&
		len(chunkInt[1]) == 3000 && len(chunkInt[2]) == 3000 && len(chunkInt[3]) == 1000)

	emptySlice := make([]uint, 0)
	chunkEmpty := ChunkSlice(emptySlice, 500)
	// This will be chunked into empty slice
	assert.True(t, len(chunkEmpty) == 0)
}

func Test_Slice_ConcatSlices(t *testing.T) {
	assert.Equal(t, []int{}, ConcatSlices[int](nil, nil, nil))
	assert.Equal(t, []bool{}, ConcatSlices([]bool{}, []bool{}))
	assert.Equal(t, []float64{1.1}, ConcatSlices([]float64{}, []float64{}, []float64{1.1}))
	assert.Equal(t, []string{"", "1", "2", "3"}, ConcatSlices([]string{""}, []string{"1", "2"}, []string{}, []string{"3"}))
}
