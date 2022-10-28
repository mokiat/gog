package gog

// Map can be used to transform one slice into another by providing a
// function to do the mapping.
func Map[S, T any](slice []S, fn func(S) T) []T {
	result := make([]T, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Reduce compacts a slice into a single value. The provided function is used
// to perform the reduction starting with the initialValue.
func Reduce[S, T any](slice []S, initialValue T, fn func(accum T, value S) T) T {
	result := initialValue
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

// Select returns a new slice that contains only the elements of the original
// slice that pass the filter function.
func Select[S any](slice []S, fn func(S) bool) []S {
	var result []S
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}
