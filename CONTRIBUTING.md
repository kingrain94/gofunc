# Contributing to GoFunc

Thank you for your interest in contributing to GoFunc! This document provides guidelines and information for contributors.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [How to Contribute](#how-to-contribute)
- [Coding Standards](#coding-standards)
- [Testing Guidelines](#testing-guidelines)
- [Documentation Guidelines](#documentation-guidelines)
- [Submitting Changes](#submitting-changes)
- [Release Process](#release-process)

## Code of Conduct

This project adheres to a code of conduct that we expect all participants to uphold. Please be respectful and constructive in all interactions.

## Getting Started

### Prerequisites

- Go 1.20 or higher
- Git
- Make (optional, but recommended)

### Development Setup

1. Fork the repository on GitHub
2. Clone your fork locally:
```bash
git clone https://github.com/YOUR_USERNAME/gofunc.git
cd gofunc
```

3. Add the original repository as upstream:
```bash
git remote add upstream https://github.com/kingrain94/gofunc.git
```

4. Create a new branch for your feature or bugfix:
```bash
git checkout -b feature/your-feature-name
```

## How to Contribute

### Types of Contributions

We welcome several types of contributions:

- **Bug fixes**: Fix issues in existing code
- **New features**: Add new utility functions
- **Documentation**: Improve documentation, examples, or comments
- **Performance improvements**: Optimize existing functions
- **Tests**: Add or improve test coverage

### Finding Issues to Work On

- Look for issues labeled `good first issue` for beginner-friendly tasks
- Check issues labeled `help wanted` for tasks that need contributors
- Feel free to propose new features by opening an issue first

## Coding Standards

### General Guidelines

- Follow standard Go conventions and idioms
- Use `gofmt` to format your code
- Use `golint` and `go vet` to check for issues
- Write clear, readable code with meaningful variable names
- Add comments for complex logic

### Function Design

- All exported functions must have GoDoc comments
- Use generics appropriately for type safety
- Follow the existing naming conventions
- Functions should be focused and do one thing well
- Prefer composition over complex single functions

### Example Function Template

```go
// FunctionName performs a specific operation on the input.
// It returns the result and an error if the operation fails.
//
// Example:
//  result, err := FunctionName(input)
//  if err != nil {
//      log.Fatal(err)
//  }
//  fmt.Println(result)
func FunctionName[T any](input T) (T, error) {
    // Implementation here
}
```

### Error Handling

- Use Go's standard error handling patterns
- Create specific error variables in `errors.go` for common errors
- Provide meaningful error messages
- Document error conditions in function comments

## Testing Guidelines

### Test Requirements

- All new functions must have comprehensive tests
- Maintain or improve the current test coverage (99.3%+)
- Write both positive and negative test cases
- Include edge cases and boundary conditions

### Test Structure

```go
func TestFunctionName(t *testing.T) {
    t.Run("positive case", func(t *testing.T) {
        // Test normal operation
        result, err := FunctionName(validInput)
        assert.NoError(t, err)
        assert.Equal(t, expectedResult, result)
    })

    t.Run("edge case", func(t *testing.T) {
        // Test edge conditions
        result, err := FunctionName(edgeInput)
        assert.NoError(t, err)
        assert.Equal(t, expectedEdgeResult, result)
    })

    t.Run("error case", func(t *testing.T) {
        // Test error conditions
        _, err := FunctionName(invalidInput)
        assert.Error(t, err)
        assert.Equal(t, ExpectedError, err)
    })
}
```

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make coverage

# Run specific test
go test -run TestFunctionName

# Run benchmarks
make bench
```

### Benchmark Tests

For performance-critical functions, include benchmark tests:

```go
func BenchmarkFunctionName(b *testing.B) {
    input := prepareTestData()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        FunctionName(input)
    }
}
```

## Documentation Guidelines

### GoDoc Comments

- Start with the function name
- Explain what the function does
- Document parameters and return values
- Include usage examples
- Mention any special behavior or limitations

### Examples

Include testable examples in your documentation:

```go
func ExampleFunctionName() {
    result, err := FunctionName(input)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result)
    // Output: expected output
}
```

### README Updates

When adding new functions:
- Update the API documentation section
- Add usage examples if appropriate
- Update the feature list if adding significant functionality

## Submitting Changes

### Before Submitting

1. Ensure all tests pass:
```bash
make test
```

2. Check code formatting:
```bash
make fmt
```

3. Run linting:
```bash
make lint
```

4. Ensure documentation is up to date

### Pull Request Process

1. Push your changes to your fork
2. Create a pull request against the main branch
3. Fill out the pull request template completely
4. Ensure all CI checks pass
5. Respond to review feedback promptly

### Pull Request Guidelines

- Use a clear, descriptive title
- Provide a detailed description of changes
- Reference any related issues
- Include test results and coverage information
- Keep PRs focused and reasonably sized

### Commit Message Format

Use conventional commit format:

```
type(scope): description

[optional body]

[optional footer]
```

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `test`: Adding or updating tests
- `refactor`: Code refactoring
- `perf`: Performance improvements
- `chore`: Maintenance tasks

Examples:
```
feat(slice): add ChunkSlice function for splitting slices

Add a new utility function that splits a slice into smaller chunks
of specified size. Includes comprehensive tests and documentation.

Closes #123
```

## Release Process

### Versioning

We follow [Semantic Versioning](https://semver.org/):
- MAJOR: Incompatible API changes
- MINOR: New functionality in a backwards compatible manner
- PATCH: Backwards compatible bug fixes

### Release Checklist

1. Update CHANGELOG.md
2. Update version in documentation
3. Ensure all tests pass
4. Create and push version tag
5. GitHub Actions will handle the release

## Development Commands

The project includes a Makefile with common development tasks:

```bash
make help        # Show available commands
make test        # Run tests
make coverage    # Run tests with coverage
make bench       # Run benchmarks
make lint        # Run linters
make fmt         # Format code
make clean       # Clean build artifacts
```

## Getting Help

- Open an issue for bugs or feature requests
- Start a discussion for questions or ideas
- Check existing issues and discussions first
- Be specific and provide examples when asking for help

## Recognition

Contributors will be recognized in:
- GitHub contributors list
- Release notes for significant contributions
- Special mentions for major features or fixes

Thank you for contributing to GoFunc! 🎉
