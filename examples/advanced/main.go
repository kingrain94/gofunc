package main

import (
	"fmt"
	"strings"

	"github.com/kingrain94/gofunc"
)

// Custom types for demonstration
type UserID int
type Username string
type User struct {
	ID   UserID
	Name Username
	Age  int
	Role string
}

func main() {
	fmt.Println("=== GoFunc Advanced Usage Examples ===")
	fmt.Println()

	// Working with custom types
	fmt.Println("1. Custom Type Conversions:")
	userIDs := []UserID{1, 2, 3, 2, 1, 4}
	fmt.Printf("User IDs: %v\n", userIDs)

	// Remove duplicate user IDs
	uniqueIDs := gofunc.ToSet(userIDs)
	fmt.Printf("Unique user IDs: %v\n", uniqueIDs)

	// Convert to regular integers
	intIDs := gofunc.ToNumberSlice[int](userIDs)
	fmt.Printf("As integers: %v\n", intIDs)

	usernames := []Username{"alice", "bob", "charlie"}
	stringUsernames := gofunc.ToStringSlice(usernames)
	fmt.Printf("Usernames as strings: %v\n", stringUsernames)

	fmt.Println()

	// Complex data processing with structs
	fmt.Println("2. Complex Data Processing:")
	users := []User{
		{ID: 1, Name: "Alice", Age: 30, Role: "admin"},
		{ID: 2, Name: "Bob", Age: 25, Role: "user"},
		{ID: 3, Name: "Charlie", Age: 35, Role: "moderator"},
		{ID: 4, Name: "Diana", Age: 28, Role: "user"},
		{ID: 1, Name: "Alice", Age: 30, Role: "admin"}, // Duplicate
	}

	fmt.Printf("Original users (%d):\n", len(users))
	for _, user := range users {
		fmt.Printf("  %+v\n", user)
	}

	// Remove duplicates by ID
	uniqueUsers := gofunc.ToSetPred(users, func(u User) UserID {
		return u.ID
	})
	fmt.Printf("\nUnique users by ID (%d):\n", len(uniqueUsers))
	for _, user := range uniqueUsers {
		fmt.Printf("  %+v\n", user)
	}

	// Find users by predicate
	admins := []User{}
	for _, user := range users {
		if user.Role == "admin" {
			admins = append(admins, user)
		}
	}
	fmt.Printf("\nAdmin users: %+v\n", admins)

	// Find first user over 30
	oldUser, found := gofunc.FindPred(users, func(u User) bool {
		return u.Age > 30
	})
	if found {
		fmt.Printf("First user over 30: %+v\n", oldUser)
	}

	// Sort users by age
	usersByAge := make([]User, len(uniqueUsers))
	copy(usersByAge, uniqueUsers)
	gofunc.SortPred(usersByAge, func(u User) int {
		return u.Age
	})
	fmt.Printf("\nUsers sorted by age:\n")
	for _, user := range usersByAge {
		fmt.Printf("  %s (age %d)\n", user.Name, user.Age)
	}

	fmt.Println()

	// Advanced map operations
	fmt.Println("3. Advanced Map Operations:")

	// Configuration merging
	defaultConfig := map[string]string{
		"host":    "localhost",
		"port":    "8080",
		"timeout": "30s",
		"retries": "3",
		"debug":   "false",
	}

	userConfig := map[string]string{
		"host":  "production.example.com",
		"port":  "443",
		"debug": "true",
		"ssl":   "true", // New setting
	}

	fmt.Printf("Default config: %v\n", defaultConfig)
	fmt.Printf("User config: %v\n", userConfig)

	// Merge configurations (user config overrides defaults)
	finalConfig := gofunc.MapUpdate(
		gofunc.MapUpdate(map[string]string{}, defaultConfig), // Copy defaults
		userConfig, // Apply user overrides
	)
	fmt.Printf("Final config: %v\n", finalConfig)

	// Get configuration values with defaults
	host := gofunc.MapGet(finalConfig, "host", "localhost")
	port := gofunc.MapGet(finalConfig, "port", "8080")
	ssl := gofunc.MapGet(finalConfig, "ssl", "false")

	fmt.Printf("\nConfiguration values:\n")
	fmt.Printf("  Host: %s\n", host)
	fmt.Printf("  Port: %s\n", port)
	fmt.Printf("  SSL: %s\n", ssl)

	fmt.Println()

	// Advanced slice operations
	fmt.Println("4. Advanced Slice Operations:")

	// Working with strings
	words := []string{"hello", "world", "go", "programming", "is", "fun", "go", "rocks"}
	fmt.Printf("Words: %v\n", words)

	// Find long words
	longWords := []string{}
	for _, word := range words {
		if len(word) > 3 {
			longWords = append(longWords, word)
		}
	}
	fmt.Printf("Long words (>3 chars): %v\n", longWords)

	// Find words containing 'o'
	wordsWithO := []string{}
	for _, word := range words {
		if strings.Contains(word, "o") {
			wordsWithO = append(wordsWithO, word)
		}
	}
	fmt.Printf("Words containing 'o': %v\n", wordsWithO)

	// Remove duplicates and sort
	uniqueWords := gofunc.ToSet(words)
	sortedWords := gofunc.Sort(uniqueWords)
	fmt.Printf("Unique words (sorted): %v\n", sortedWords)

	// Chunk words for processing
	wordChunks := gofunc.ChunkSlice(sortedWords, 3)
	fmt.Printf("Word chunks (size 3):\n")
	for i, chunk := range wordChunks {
		fmt.Printf("  Chunk %d: %v\n", i+1, chunk)
	}

	// Concatenate multiple slices
	moreWords := []string{"awesome", "powerful", "simple"}
	evenMoreWords := []string{"efficient", "type-safe"}
	allWords := gofunc.ConcatSlices(sortedWords, moreWords, evenMoreWords)
	fmt.Printf("All words combined: %v\n", allWords)

	fmt.Println()

	// Performance-conscious operations
	fmt.Println("5. Performance Examples:")

	// Large slice operations
	largeSlice := make([]int, 10000)
	for i := range largeSlice {
		largeSlice[i] = i % 100 // Create some duplicates
	}

	fmt.Printf("Large slice length: %d\n", len(largeSlice))

	// Remove duplicates efficiently
	uniqueLarge := gofunc.ToSet(largeSlice)
	fmt.Printf("Unique elements in large slice: %d\n", len(uniqueLarge))

	// Check if contains specific values efficiently
	containsTarget := gofunc.Contains(largeSlice, 42)
	fmt.Printf("Contains 42: %t\n", containsTarget)

	// Find index efficiently
	targetIndex := gofunc.IndexOf(uniqueLarge, 42)
	fmt.Printf("Index of 42 in unique slice: %d\n", targetIndex)

	fmt.Println()
	fmt.Println("=== Advanced examples completed successfully! ===")
}
