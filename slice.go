package gog

// Map can be used to transform one slice into another by providing a
// function to do the mapping.
func Map[S, T any](slice []S, fn func(S) T) []T {
	if slice == nil {
		return nil
	}
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

// Partition splits the specified slice into groups, each having a key
// according to the specified function.
func Partition[S any, K comparable](slice []S, fn func(S) K) map[K][]S {
	result := make(map[K][]S)
	for _, v := range slice {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}

// Mapping is similar to Partition, except that it allows one to transform
// the values stored in the patition buckets.
// In essence, it allows the caller to construct an almost arbitrary map
// (it is always of kind map[key][]value, though the types of the keys and the
// values are user controlled) from an arbitrary slice.
func Mapping[S any, K comparable, V any](slice []S, fn func(S) (K, V)) map[K][]V {
	result := make(map[K][]V)
	for _, v := range slice {
		key, value := fn(v)
		result[key] = append(result[key], value)
	}
	return result
}

// Dedupe returns a new slice that contains only distinct elements from
// the original slice.
func Dedupe[T comparable](slice []T) []T {
	if slice == nil {
		return nil
	}
	seen := make(map[T]struct{}, len(slice))
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if _, ok := seen[v]; !ok {
			result = append(result, v)
			seen[v] = struct{}{}
		}
	}
	return result
}

// Flatten returns a new slice that is the result of merging all nested
// slices into a single top-level slice.
func Flatten[T any](slice [][]T) []T {
	var result []T
	for _, subSlice := range slice {
		result = append(result, subSlice...)
	}
	return result
}
