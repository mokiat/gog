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
// the values stored in the partition buckets.
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

// Mutate iterates over the slice and calls the provided closure function,
// passing a pointer to each element. This allows the user to easily modify
// all the elements of a slice.
func Mutate[T any](slice []T, fn func(e *T)) {
	for i := range slice {
		fn(&slice[i])
	}
}

// FindFunc iterates over the slice and uses the provided closure function
// to check whether the elements match a user-provided condition. The first
// value that matches is returned as well as a true flag. Otherwise a
// false flag is returned.
func FindFunc[T any](slice []T, fn func(e T) bool) (T, bool) {
	for _, value := range slice {
		if fn(value) {
			return value, true
		}
	}
	var zeroT T
	return zeroT, false
}

// FindFuncPtr iterates over the slice and uses the provided closure function
// to check whether the elements match a user-provided condition. A pointer
// to the first element that passes the condition is returned. If no value
// is applicable, then nil is returned.
func FindFuncPtr[T any](slice []T, fn func(e T) bool) *T {
	for i := range slice {
		if fn(slice[i]) {
			return &slice[i]
		}
	}
	return nil
}

// RefElements converts a slice of elements into a slice of pointers to
// those same elements.
func RefElements[T any](slice []T) []*T {
	result := make([]*T, len(slice))
	for i := range slice {
		result[i] = &slice[i]
	}
	return result
}

// DerefElements converts a slice of element pointers into a slice of
// those elements.
func DerefElements[T any](slice []*T) []T {
	result := make([]T, len(slice))
	for i, element := range slice {
		result[i] = *element
	}
	return result
}
