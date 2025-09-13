package gofunc_test

import (
	"fmt"

	"github.com/kingrain94/gofunc"
)

func ExampleToSet() {
	numbers := []int{1, 2, 2, 3, 1, 4}
	unique := gofunc.ToSet(numbers)
	fmt.Println(unique)
	// Output: [1 2 3 4]
}

func ExampleToSetPred() {
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Alice", 31}, // Different age, same name
	}

	unique := gofunc.ToSetPred(people, func(p Person) string {
		return p.Name
	})

	fmt.Printf("Unique people: %d\n", len(unique))
	for _, p := range unique {
		fmt.Printf("%s (%d)\n", p.Name, p.Age)
	}
	// Output: Unique people: 2
	// Alice (30)
	// Bob (25)
}

func ExampleContains() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(gofunc.Contains(numbers, 3))
	fmt.Println(gofunc.Contains(numbers, 6))
	// Output: true
	// false
}

func ExampleContainsPred() {
	numbers := []int{1, 3, 5, 7, 9}
	hasEven := gofunc.ContainsPred(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(hasEven)
	// Output: false
}

func ExampleFindPred() {
	numbers := []int{1, 3, 4, 7, 8}
	even, found := gofunc.FindPred(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("Even number: %d, found: %t\n", even, found)
	// Output: Even number: 4, found: true
}

func ExampleIndexOf() {
	colors := []string{"red", "green", "blue", "green"}
	fmt.Println(gofunc.IndexOf(colors, "green"))
	fmt.Println(gofunc.IndexOf(colors, "yellow"))
	// Output: 1
	// -1
}

func ExampleLastIndexOf() {
	colors := []string{"red", "green", "blue", "green"}
	fmt.Println(gofunc.LastIndexOf(colors, "green"))
	// Output: 3
}

func ExampleChunkSlice() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	chunks := gofunc.ChunkSlice(numbers, 3)
	for i, chunk := range chunks {
		fmt.Printf("Chunk %d: %v\n", i+1, chunk)
	}
	// Output: Chunk 1: [1 2 3]
	// Chunk 2: [4 5 6]
	// Chunk 3: [7 8]
}

func ExampleConcatSlices() {
	slice1 := []int{1, 2}
	slice2 := []int{3, 4}
	slice3 := []int{5, 6}
	result := gofunc.ConcatSlices(slice1, slice2, slice3)
	fmt.Println(result)
	// Output: [1 2 3 4 5 6]
}

func ExampleMapKeys() {
	m := map[string]int{"apple": 5, "banana": 3}
	keys := gofunc.MapKeys(m)
	fmt.Printf("Keys count: %d\n", len(keys))
	// Output: Keys count: 2
}

func ExampleMapValues() {
	m := map[string]int{"apple": 5, "banana": 3}
	values := gofunc.MapValues(m)
	fmt.Printf("Values count: %d\n", len(values))
	// Output: Values count: 2
}

func ExampleMapGet() {
	config := map[string]string{
		"host": "localhost",
		"port": "8080",
	}

	host := gofunc.MapGet(config, "host", "127.0.0.1")
	timeout := gofunc.MapGet(config, "timeout", "30s")

	fmt.Printf("Host: %s\n", host)
	fmt.Printf("Timeout: %s\n", timeout)
	// Output: Host: localhost
	// Timeout: 30s
}

func ExampleMapSetDefault() {
	config := map[string]string{"host": "localhost"}

	// Key exists
	host, existed := gofunc.MapSetDefault(config, "host", "127.0.0.1")
	fmt.Printf("Host: %s, existed: %t\n", host, existed)

	// Key doesn't exist
	port, existed := gofunc.MapSetDefault(config, "port", "8080")
	fmt.Printf("Port: %s, existed: %t\n", port, existed)

	fmt.Printf("Config now has %d keys\n", len(config))
	// Output: Host: localhost, existed: true
	// Port: 8080, existed: false
	// Config now has 2 keys
}

func ExampleMapUpdate() {
	defaults := map[string]string{
		"host":    "localhost",
		"port":    "8080",
		"timeout": "30s",
	}

	overrides := map[string]string{
		"host": "production.com",
		"ssl":  "true",
	}

	result := gofunc.MapUpdate(defaults, overrides)
	fmt.Printf("Host: %s\n", result["host"])
	fmt.Printf("SSL: %s\n", result["ssl"])
	fmt.Printf("Timeout: %s\n", result["timeout"])
	// Output: Host: production.com
	// SSL: true
	// Timeout: 30s
}

func ExampleMin() {
	min, err := gofunc.Min(3, 1, 4, 1, 5, 9)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Minimum: %d\n", min)
	// Output: Minimum: 1
}

func ExampleMax() {
	max, err := gofunc.Max(3, 1, 4, 1, 5, 9)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Maximum: %d\n", max)
	// Output: Maximum: 9
}

func ExampleAbs() {
	fmt.Println(gofunc.Abs(-42))
	fmt.Println(gofunc.Abs(42))
	fmt.Println(gofunc.Abs(0))
	// Output: 42
	// 42
	// 0
}

func ExampleSort() {
	numbers := []int{3, 1, 4, 1, 5, 9}
	sorted := gofunc.Sort(numbers)
	fmt.Println(sorted)
	// Output: [1 1 3 4 5 9]
}

func ExampleSortPred() {
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Bob", 30},
		{"Alice", 25},
		{"Charlie", 35},
	}

	// Sort by age
	gofunc.SortPred(people, func(p Person) int {
		return p.Age
	})

	for _, p := range people {
		fmt.Printf("%s (%d)\n", p.Name, p.Age)
	}
	// Output: Alice (25)
	// Bob (30)
	// Charlie (35)
}

func ExampleIf() {
	score := 85
	grade := gofunc.If(score >= 90, "A", gofunc.If(score >= 80, "B", "C"))
	fmt.Printf("Score %d gets grade: %s\n", score, grade)
	// Output: Score 85 gets grade: B
}

func ExampleMust() {
	// This would panic if there was an error
	value := gofunc.Must(42, nil)
	fmt.Printf("Value: %d\n", value)
	// Output: Value: 42
}

func ExampleToInterfaceSlice() {
	numbers := []int{1, 2, 3}
	interfaces := gofunc.ToInterfaceSlice(numbers)
	fmt.Printf("Length: %d, Type: %T\n", len(interfaces), interfaces)
	// Output: Length: 3, Type: []interface {}
}

func ExampleToStringSlice() {
	type UserID string
	ids := []UserID{"user1", "user2", "user3"}
	strings := gofunc.ToStringSlice(ids)
	fmt.Printf("Strings: %v\n", strings)
	// Output: Strings: [user1 user2 user3]
}

func ExampleToNumberSlice() {
	type OrderID uint
	orders := []OrderID{1, 2, 3}
	numbers := gofunc.ToNumberSlice[uint](orders)
	fmt.Printf("Numbers: %v\n", numbers)
	// Output: Numbers: [1 2 3]
}
