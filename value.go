package gog

// Zero returns the zero value of the generic type T.
func Zero[T any]() T {
	var zero T
	return zero
}

// Must is a helper function that can be used to quickly extract a value from a
// function that returns a value and an error. It panics if the error is not
// nil.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// Ternary is a generic ternary operator. It returns trueValue if condition is
// true, and falseValue otherwise.
//
// BEWARE: Unlike a traditional ternary operator, both trueValue and falseValue
// are evaluated before the condition is checked. Using it in the following way
// would be wrong (as it would cause a panic if a is nil):
//
//	Ternary(a != nil, a.Value(), 0)
//
// Instead, it is intended to simplify basic conditional assignments,
// for example:
//
//	color := Ternary(element.IsSelected, "0xFFFFFF", "0x000000")
//
// While it can greatly reduce boilerplate code, it should be used with care.
func Ternary[T any](condition bool, trueValue T, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}
