# GoFunc - Go Utility Functions Library

[![Go Reference](https://pkg.go.dev/badge/github.com/kingrain94/gofunc.svg)](https://pkg.go.dev/github.com/kingrain94/gofunc)
[![Go Report Card](https://goreportcard.com/badge/github.com/kingrain94/gofunc)](https://goreportcard.com/report/github.com/kingrain94/gofunc)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Coverage Status](https://img.shields.io/badge/coverage-99.3%25-brightgreen.svg)](https://github.com/kingrain94/gofunc)

A comprehensive collection of utility functions for Go, leveraging generics to provide type-safe, reusable operations for common programming tasks.

## Features

- 🚀 **Type-safe**: Built with Go 1.20+ generics for maximum type safety
- 📦 **Zero dependencies**: Pure Go implementation with no external dependencies
- 🧪 **Well tested**: 99.3% test coverage with comprehensive test suites
- 📚 **Well documented**: Extensive GoDoc documentation with examples
- ⚡ **High performance**: Optimized implementations for common operations
- 🔧 **Easy to use**: Simple, intuitive API design

## Installation

```bash
go get github.com/kingrain94/gofunc
```

## Requirements

- Go 1.20 or higher (for generics support)

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/kingrain94/gofunc"
)

func main() {
    // Slice operations
    numbers := []int{1, 2, 3, 2, 4, 1}
    unique := gofunc.ToSet(numbers)
    fmt.Println(unique) // [1, 2, 3, 4]

    // Conditional operations
    result := gofunc.If(len(numbers) > 5, "many", "few")
    fmt.Println(result) // "many"

    // Min/Max operations
    min, _ := gofunc.Min(3, 1, 4, 1, 5)
    max, _ := gofunc.Max(3, 1, 4, 1, 5)
    fmt.Printf("Min: %d, Max: %d\n", min, max) // Min: 1, Max: 5

    // Map operations
    m := map[string]int{"a": 1, "b": 2}
    keys := gofunc.MapKeys(m)
    fmt.Println(keys) // [a, b] (order may vary)
}
```

## API Documentation

### Slice Operations

#### Type Conversions
- `ToInterfaceSlice[T any](slice []T) []interface{}` - Convert slice to interface slice
- `ToStringSlice[T ~string](slice []T) []string` - Convert string-like slice to string slice
- `ToStringDerivedSlice[U, T ~string](slice []T) []U` - Convert string slice to string-derived type
- `ToNumberSlice[U, T Number](slice []T) []U` - Convert numeric slice to another numeric type

#### Set Operations
- `ToSet[T comparable](s []T) []T` - Remove duplicates from slice
- `ToSetPred[T any, K comparable](s []T, keyFunc func(t T) K) []T` - Remove duplicates using key function

#### Search Operations
- `Contains[T comparable](a []T, b T) bool` - Check if slice contains item
- `ContainsPred[T any](a []T, pred func(b T) bool) bool` - Check if slice contains item matching predicate
- `FindPred[T any](a []T, pred func(b T) bool) (T, bool)` - Find first item matching predicate
- `IndexOf[T comparable](a []T, t T) int` - Find index of item
- `LastIndexOf[T comparable](a []T, t T) int` - Find last index of item
- `IndexOfSlice[T comparable](a []T, sub []T) int` - Find index of sub-slice

#### Manipulation Operations
- `ChunkSlice[T any](slice []T, chunkSize int) [][]T` - Split slice into chunks
- `ConcatSlices[T any](slices ...[]T) []T` - Concatenate multiple slices

### Map Operations

- `MapUpdate[K comparable, V any](m1, m2 map[K]V) map[K]V` - Merge two maps
- `MapKeys[K comparable, V any](m map[K]V) []K` - Get map keys as slice
- `MapValues[K comparable, V any](m map[K]V) []V` - Get map values as slice
- `MapGet[K comparable, V any](m map[K]V, k K, defaultVal V) V` - Get value with default
- `MapSetDefault[K comparable, V any](m map[K]V, k K, defaultVal V) (V, bool)` - Set default value

### Mathematical Operations

- `Abs(x int64) int64` - Absolute value for int64
- `Min[T Number | ~string](s ...T) (T, error)` - Find minimum value
- `Max[T Number | ~string](s ...T) (T, error)` - Find maximum value

### Sorting Operations

- `Sort[T Number | ~string](s []T) []T` - Sort slice in ascending order
- `SortPred[T any, K Number | ~string](s []T, keyFunc func(t T) K) []T` - Sort by key function

### Utility Operations

- `If[T any](cond bool, a T, b T) T` - Conditional expression (ternary-like)
- `Must[T any](v T, e error) T` - Panic on error, return value otherwise

### Type Definitions

```go
type Int interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64
}

type UInt interface {
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Float interface {
    ~float32 | ~float64
}

type Number interface {
    Int | UInt | Float
}
```

## Examples

### Working with Custom Types

```go
type UserID int
type Username string

users := []UserID{1, 2, 3, 2, 1}
uniqueUsers := gofunc.ToSet(users) // []UserID{1, 2, 3}

usernames := []Username{"alice", "bob", "charlie"}
stringUsernames := gofunc.ToStringSlice(usernames) // []string{"alice", "bob", "charlie"}
```

### Complex Data Processing

```go
type Person struct {
    Name string
    Age  int
}

people := []Person{
    {"Alice", 30},
    {"Bob", 25},
    {"Charlie", 35},
    {"Alice", 30}, // duplicate
}

// Remove duplicates by name
uniquePeople := gofunc.ToSetPred(people, func(p Person) string {
    return p.Name
})

// Sort by age
gofunc.SortPred(people, func(p Person) int {
    return p.Age
})

// Find person by predicate
person, found := gofunc.FindPred(people, func(p Person) bool {
    return p.Age > 30
})
```

### Map Operations

```go
config := map[string]string{
    "host": "localhost",
    "port": "8080",
}

defaults := map[string]string{
    "port":    "3000",
    "timeout": "30s",
}

// Merge configurations (config takes precedence)
merged := gofunc.MapUpdate(defaults, config)
// Result: {"host": "localhost", "port": "8080", "timeout": "30s"}

// Get with default
timeout := gofunc.MapGet(config, "timeout", "10s") // "10s"
```

## Performance

GoFunc is designed for performance. Here are some benchmark results:

```
BenchmarkToSet-8           1000000      1043 ns/op     896 B/op      2 allocs/op
BenchmarkContains-8       10000000       156 ns/op       0 B/op      0 allocs/op
BenchmarkSort-8            5000000       312 ns/op       0 B/op      0 allocs/op
```

## Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details on how to contribute to this project.

### Development Setup

1. Clone the repository:
```bash
git clone https://github.com/kingrain94/gofunc.git
cd gofunc
```

2. Run tests:
```bash
make test
```

3. Run benchmarks:
```bash
make bench
```

4. Check coverage:
```bash
make coverage
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for a list of changes and version history.

## Support

- 📖 [Documentation](https://pkg.go.dev/github.com/kingrain94/gofunc)
- 🐛 [Issue Tracker](https://github.com/kingrain94/gofunc/issues)
- 💬 [Discussions](https://github.com/kingrain94/gofunc/discussions)

## Related Projects

- [samber/lo](https://github.com/samber/lo) - A Lodash-style Go library based on Go 1.18+ Generics
- [thoas/go-funk](https://github.com/thoas/go-funk) - A modern Go utility library
