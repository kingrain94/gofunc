package main

import (
	"fmt"
	"log"

	"github.com/kingrain94/gofunc"
)

func main() {
	fmt.Println("=== GoFunc Basic Usage Examples ===")
	fmt.Println()

	// Slice operations
	fmt.Println("1. Slice Operations:")
	numbers := []int{1, 2, 3, 2, 4, 1, 5}
	fmt.Printf("Original slice: %v\n", numbers)

	// Remove duplicates
	unique := gofunc.ToSet(numbers)
	fmt.Printf("Unique elements: %v\n", unique)

	// Check if slice contains an element
	contains := gofunc.Contains(numbers, 3)
	fmt.Printf("Contains 3: %t\n", contains)

	// Find index of element
	index := gofunc.IndexOf(numbers, 2)
	fmt.Printf("Index of 2: %d\n", index)

	// Chunk slice
	chunks := gofunc.ChunkSlice(numbers, 3)
	fmt.Printf("Chunked by 3: %v\n", chunks)

	fmt.Println()

	// Map operations
	fmt.Println("2. Map Operations:")
	userRoles := map[string]string{
		"alice": "admin",
		"bob":   "user",
		"carol": "moderator",
	}
	fmt.Printf("User roles: %v\n", userRoles)

	// Get map keys
	users := gofunc.MapKeys(userRoles)
	fmt.Printf("Users: %v\n", users)

	// Get map values
	roles := gofunc.MapValues(userRoles)
	fmt.Printf("Roles: %v\n", roles)

	// Get value with default
	role := gofunc.MapGet(userRoles, "david", "guest")
	fmt.Printf("David's role (with default): %s\n", role)

	fmt.Println()

	// Mathematical operations
	fmt.Println("3. Mathematical Operations:")
	values := []int{10, 5, 8, 3, 15}
	fmt.Printf("Values: %v\n", values)

	min, err := gofunc.Min(values...)
	if err != nil {
		log.Printf("Error finding min: %v", err)
	} else {
		fmt.Printf("Minimum: %d\n", min)
	}

	max, err := gofunc.Max(values...)
	if err != nil {
		log.Printf("Error finding max: %v", err)
	} else {
		fmt.Printf("Maximum: %d\n", max)
	}

	// Absolute value
	abs := gofunc.Abs(-42)
	fmt.Printf("Absolute value of -42: %d\n", abs)

	fmt.Println()

	// Sorting operations
	fmt.Println("4. Sorting Operations:")
	unsorted := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Unsorted: %v\n", unsorted)

	sorted := gofunc.Sort(append([]int(nil), unsorted...)) // Copy slice before sorting
	fmt.Printf("Sorted: %v\n", sorted)

	fmt.Println()

	// Conditional operations
	fmt.Println("5. Conditional Operations:")
	score := 85
	grade := gofunc.If(score >= 90, "A", gofunc.If(score >= 80, "B", "C"))
	fmt.Printf("Score %d gets grade: %s\n", score, grade)

	fmt.Println()

	// Utility operations
	fmt.Println("6. Utility Operations:")
	// Must function - panics on error, returns value otherwise
	result := gofunc.Must(fmt.Sprintf("Hello, %s!", "World"), nil)
	fmt.Printf("Must result: %s\n", result)

	fmt.Println()
	fmt.Println("=== Examples completed successfully! ===")
}
