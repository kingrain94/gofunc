package gofunc

// Int represents all signed integer types, including custom types based on signed integers.
// This constraint allows functions to work with any signed integer type.
type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// UInt represents all unsigned integer types, including custom types based on unsigned integers.
// This constraint allows functions to work with any unsigned integer type.
type UInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Float represents all floating-point types, including custom types based on floats.
// This constraint allows functions to work with any floating-point type.
type Float interface {
	~float32 | ~float64
}

// Number represents all numeric types (integers and floating-point).
// This is a union of Int, UInt, and Float constraints, allowing functions
// to work with any numeric type.
type Number interface {
	Int | UInt | Float
}
