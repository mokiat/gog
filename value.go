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
// While it can greatly reduce boilerplate code, it should be used sparingly to
// avoid reducing code readability (e.g. avoid nesting Ternary calls).
func Ternary[T any](condition bool, trueValue T, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}
