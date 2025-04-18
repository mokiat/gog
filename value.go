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
