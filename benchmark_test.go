package gofunc

import (
	"fmt"
	"testing"
)

// Benchmark for ToSet function
func BenchmarkToSet(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			// Create slice with some duplicates
			slice := make([]int, size)
			for i := 0; i < size; i++ {
				slice[i] = i % (size / 2) // Creates duplicates
			}
			
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ToSet(slice)
			}
		})
	}
}

// Benchmark for ToSetPred function
func BenchmarkToSetPred(b *testing.B) {
	type testStruct struct {
		ID   int
		Name string
	}
	
	sizes := []int{10, 100, 1000}
	
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			slice := make([]testStruct, size)
			for i := 0; i < size; i++ {
				slice[i] = testStruct{
					ID:   i % (size / 2), // Creates duplicates
					Name: fmt.Sprintf("name-%d", i),
				}
			}
			
			keyFunc := func(t testStruct) int { return t.ID }
			
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ToSetPred(slice, keyFunc)
			}
		})
	}
}

// Benchmark for Contains function
func BenchmarkContains(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			slice := make([]int, size)
			for i := 0; i < size; i++ {
				slice[i] = i
			}
			target := size / 2 // Middle element
			
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				Contains(slice, target)
			}
		})
	}
}

// Benchmark for ContainsPred function
func BenchmarkContainsPred(b *testing.B) {
	sizes := []int{10, 100, 1000}
	
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			slice := make([]int, size)
			for i := 0; i < size; i++ {
				slice[i] = i
			}
			pred := func(x int) bool { return x == size/2 }
			
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ContainsPred(slice, pred)
			}
		})
	}
}

// Benchmark for IndexOf function
func BenchmarkIndexOf(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			slice := make([]int, size)
			for i := 0; i < size; i++ {
				slice[i] = i
			}
			target := size / 2 // Middle element
			
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				IndexOf(slice, target)
			}
		})
	}
}

// Benchmark for ChunkSlice function
func BenchmarkChunkSlice(b *testing.B) {
	sizes := []int{100, 1000, 10000}
	chunkSizes := []int{10, 50, 100}
	
	for _, size := range sizes {
		for _, chunkSize := range chunkSizes {
			b.Run(fmt.Sprintf("size-%d-chunk-%d", size, chunkSize), func(b *testing.B) {
				slice := make([]int, size)
				for i := 0; i < size; i++ {
					slice[i] = i
				}
				
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					ChunkSlice(slice, chunkSize)
				}
			})
		}
	}
}

// Benchmark for ConcatSlices function
func BenchmarkConcatSlices(b *testing.B) {
	sliceCounts := []int{2, 5, 10}
	sliceSize := 1000
	
	for _, count := range sliceCounts {
		b.Run(fmt.Sprintf("slices-%d", count), func(b *testing.B) {
			slices := make([][]int, count)
			for i := 0; i < count; i++ {
				slice := make([]int, sliceSize)
				for j := 0; j < sliceSize; j++ {
					slice[j] = i*sliceSize + j
				}
				slices[i] = slice
			}
			
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ConcatSlices(slices...)
			}
		})
	}
}

// Benchmark for Sort function
func BenchmarkSort(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			// Create a slice in reverse order (worst case for many algorithms)
			slice := make([]int, size)
			for i := 0; i < size; i++ {
				slice[i] = size - i
			}
			
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// Create a copy for each iteration
				testSlice := make([]int, len(slice))
				copy(testSlice, slice)
				Sort(testSlice)
			}
		})
	}
}

// Benchmark for SortPred function
func BenchmarkSortPred(b *testing.B) {
	type testStruct struct {
		Value int
		Name  string
	}
	
	sizes := []int{10, 100, 1000}
	
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			slice := make([]testStruct, size)
			for i := 0; i < size; i++ {
				slice[i] = testStruct{
					Value: size - i, // Reverse order
					Name:  fmt.Sprintf("item-%d", i),
				}
			}
			
			keyFunc := func(t testStruct) int { return t.Value }
			
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// Create a copy for each iteration
				testSlice := make([]testStruct, len(slice))
				copy(testSlice, slice)
				SortPred(testSlice, keyFunc)
			}
		})
	}
}

// Benchmark for MapKeys function
func BenchmarkMapKeys(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			m := make(map[int]string, size)
			for i := 0; i < size; i++ {
				m[i] = fmt.Sprintf("value-%d", i)
			}
			
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				MapKeys(m)
			}
		})
	}
}

// Benchmark for MapValues function
func BenchmarkMapValues(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			m := make(map[int]string, size)
			for i := 0; i < size; i++ {
				m[i] = fmt.Sprintf("value-%d", i)
			}
			
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				MapValues(m)
			}
		})
	}
}

// Benchmark for MapUpdate function
func BenchmarkMapUpdate(b *testing.B) {
	sizes := []int{10, 100, 1000}
	
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			m1 := make(map[int]string, size)
			m2 := make(map[int]string, size/2)
			
			for i := 0; i < size; i++ {
				m1[i] = fmt.Sprintf("value1-%d", i)
			}
			
			for i := 0; i < size/2; i++ {
				m2[i] = fmt.Sprintf("value2-%d", i)
			}
			
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// Create copies for each iteration
				m1Copy := make(map[int]string, len(m1))
				for k, v := range m1 {
					m1Copy[k] = v
				}
				MapUpdate(m1Copy, m2)
			}
		})
	}
}

// Benchmark for Min function
func BenchmarkMin(b *testing.B) {
	sizes := []int{10, 100, 1000}
	
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			values := make([]int, size)
			for i := 0; i < size; i++ {
				values[i] = i
			}
			
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				Min(values...)
			}
		})
	}
}

// Benchmark for Max function
func BenchmarkMax(b *testing.B) {
	sizes := []int{10, 100, 1000}
	
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			values := make([]int, size)
			for i := 0; i < size; i++ {
				values[i] = i
			}
			
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				Max(values...)
			}
		})
	}
}

// Benchmark for type conversion functions
func BenchmarkToInterfaceSlice(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			slice := make([]int, size)
			for i := 0; i < size; i++ {
				slice[i] = i
			}
			
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ToInterfaceSlice(slice)
			}
		})
	}
}

// Benchmark for string conversion functions
func BenchmarkToStringSlice(b *testing.B) {
	type MyString string
	
	sizes := []int{10, 100, 1000}
	
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			slice := make([]MyString, size)
			for i := 0; i < size; i++ {
				slice[i] = MyString(fmt.Sprintf("string-%d", i))
			}
			
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ToStringSlice(slice)
			}
		})
	}
}
