// Package constr provides a set of constraints for Go generics.
package constr

// Integer represents a type that is a signed or unsigned integer.
type Integer interface {
	Signed | Unsigned
}

// Signed represents a type that is a signed integer.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned represents a type that is an unsigned integer.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Float represents a type that is a floating-point number.
type Float interface {
	~float32 | ~float64
}

// Numeric represents a type that is a number.
type Numeric interface {
	Integer | Float
}
