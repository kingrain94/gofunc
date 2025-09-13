# GoFunc Examples

This directory contains practical examples demonstrating how to use the GoFunc utility library in real-world scenarios.

## Examples Overview

### 1. Basic Usage (`basic/main.go`)

Demonstrates fundamental operations with GoFunc:

- **Slice operations**: Removing duplicates, checking containment, finding indices, chunking
- **Map operations**: Getting keys/values, merging maps, default values
- **Mathematical operations**: Finding min/max values, absolute values
- **Sorting operations**: Sorting slices with built-in comparisons
- **Conditional operations**: Ternary-like conditional expressions
- **Utility operations**: Must pattern for error handling

**Run the example:**
```bash
cd examples/basic
go run main.go
```

### 2. Advanced Usage (`advanced/main.go`)

Shows more complex scenarios and custom type usage:

- **Custom type conversions**: Working with user-defined types
- **Complex data processing**: Struct manipulation, deduplication by custom keys
- **Advanced map operations**: Configuration merging, nested operations
- **Advanced slice operations**: String processing, chunking, concatenation
- **Performance examples**: Efficient operations on large datasets

**Run the example:**
```bash
cd examples/advanced
go run main.go
```

## Running All Examples

You can run all examples at once:

```bash
# From the project root
make run-examples

# Or manually
cd examples/basic && go run main.go
cd examples/advanced && go run main.go
```

## Example Patterns

### Working with Custom Types

```go
type UserID int
type User struct {
    ID   UserID
    Name string
    Age  int
}

users := []User{ /* ... */ }

// Remove duplicates by ID
uniqueUsers := gofunc.ToSetPred(users, func(u User) UserID {
    return u.ID
})

// Sort by age
gofunc.SortPred(users, func(u User) int {
    return u.Age
})
```

### Configuration Management

```go
defaults := map[string]string{
    "host": "localhost",
    "port": "8080",
}

userConfig := map[string]string{
    "host": "production.com",
    "ssl":  "true",
}

// Merge with user config taking precedence
config := gofunc.MapUpdate(defaults, userConfig)

// Get values with fallbacks
host := gofunc.MapGet(config, "host", "localhost")
```

### Data Processing Pipeline

```go
data := []string{"item1", "item2", "item1", "item3"}

// Process: deduplicate → sort → chunk
processed := gofunc.ChunkSlice(
    gofunc.Sort(
        gofunc.ToSet(data),
    ),
    2,
)
```

## Best Practices Demonstrated

1. **Type Safety**: Using generics for compile-time type checking
2. **Performance**: Efficient algorithms for common operations
3. **Composability**: Combining functions for complex operations
4. **Error Handling**: Proper error checking patterns
5. **Memory Efficiency**: Avoiding unnecessary allocations

## Adding Your Own Examples

To add a new example:

1. Create a new `.go` file in this directory
2. Follow the existing naming pattern (`your_example.go`)
3. Include comprehensive comments explaining the use case
4. Add a section to this README describing your example
5. Test your example to ensure it runs correctly

## Integration with Your Projects

These examples show patterns you can use in your own projects:

```go
package main

import "github.com/kingrain94/gofunc"

func processUserData(users []User) []User {
    // Remove duplicates, sort by name, return first 10
    unique := gofunc.ToSetPred(users, func(u User) string {
        return u.Email
    })
    
    sorted := gofunc.SortPred(unique, func(u User) string {
        return u.Name
    })
    
    if len(sorted) > 10 {
        return sorted[:10]
    }
    return sorted
}
```

## Performance Notes

The examples include performance-conscious patterns:

- Using `ToSet` for O(n) deduplication instead of nested loops
- Leveraging `Contains` for O(n) searches instead of manual iteration
- Using `ChunkSlice` for efficient batch processing
- Demonstrating `SortPred` for custom sorting criteria

For more performance tips, see the main [README.md](../README.md) and [CONTRIBUTING.md](../CONTRIBUTING.md).
