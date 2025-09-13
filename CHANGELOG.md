# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Comprehensive README.md with installation and usage examples
- CONTRIBUTING.md with development guidelines
- GitHub Actions for CI/CD pipeline
- Makefile for common development tasks
- Benchmark tests for performance-critical functions
- GoDoc examples for all exported functions
- Examples directory with practical usage scenarios

### Changed
- Improved GoDoc comments with detailed descriptions and examples
- Enhanced error handling patterns

### Fixed
- Minor documentation improvements

## [1.0.0] - 2023-12-XX

### Added
- Initial release of GoFunc utility library
- Slice operations: conversion, set operations, search, manipulation
- Map operations: keys, values, merging, default values
- Mathematical operations: absolute value, min/max functions
- Sorting operations with custom predicates
- Utility functions: conditional expressions, must patterns
- Comprehensive test suite with 99.3% coverage
- MIT License
- Go 1.20+ generics support

### Slice Functions
- `ToInterfaceSlice` - Convert slice to interface slice
- `ToStringSlice` - Convert string-like slice to string slice
- `ToStringDerivedSlice` - Convert string slice to string-derived type
- `ToNumberSlice` - Convert numeric slice to another numeric type
- `ToSet` - Remove duplicates from slice
- `ToSetPred` - Remove duplicates using key function
- `Contains` - Check if slice contains item
- `ContainsPred` - Check if slice contains item matching predicate
- `FindPred` - Find first item matching predicate
- `IndexOf` - Find index of item
- `LastIndexOf` - Find last index of item
- `IndexOfSlice` - Find index of sub-slice
- `ChunkSlice` - Split slice into chunks
- `ConcatSlices` - Concatenate multiple slices

### Map Functions
- `MapUpdate` - Merge two maps
- `MapKeys` - Get map keys as slice
- `MapValues` - Get map values as slice
- `MapGet` - Get value with default
- `MapSetDefault` - Set default value

### Math Functions
- `Abs` - Absolute value for int64
- `Min` - Find minimum value
- `Max` - Find maximum value

### Sorting Functions
- `Sort` - Sort slice in ascending order
- `SortPred` - Sort by key function

### Utility Functions
- `If` - Conditional expression (ternary-like)
- `Must` - Panic on error, return value otherwise

### Type Definitions
- `Int` - Integer type constraint
- `UInt` - Unsigned integer type constraint
- `Float` - Floating point type constraint
- `Number` - Numeric type constraint

---

## Release Notes Format

Each release includes:
- **Added**: New features and functions
- **Changed**: Changes to existing functionality
- **Deprecated**: Soon-to-be removed features
- **Removed**: Features removed in this version
- **Fixed**: Bug fixes
- **Security**: Vulnerability fixes

## Version History

- **v1.0.0**: Initial stable release with core functionality
- **Future releases**: Will follow semantic versioning

## Migration Guides

### From v0.x to v1.0.0
- No migration needed for new users
- First stable release

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for information about contributing to this project and how releases are managed.
