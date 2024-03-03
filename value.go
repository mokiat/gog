package gog

// Zero returns the zero value of the generic type T.
func Zero[T any]() T {
	var zero T
	return zero
}
