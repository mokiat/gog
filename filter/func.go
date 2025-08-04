package filter

import "slices"

// Func represents a filtering function. The actual semantics of the Func
// depend on the user but in general if true is returned then the value
// that is being checked is accepted.
type Func[T any] func(T) bool

// True creates a filter Func that always returns true, regardless of the
// input value.
func True[T any]() Func[T] {
	return func(T) bool {
		return true
	}
}

// False creates a filter Func that always returns false, regardless of the
// input value.
func False[T any]() Func[T] {
	return func(T) bool {
		return false
	}
}

// Not returns a filter function that returns the opposite of what the specified
// Func in the arguments would return.
func Not[T any](delegate Func[T]) Func[T] {
	return func(item T) bool {
		return !delegate(item)
	}
}

// Equal returns a filter that returns true only if the input value matches
// the specified value.
func Equal[T comparable](expected T) Func[T] {
	return func(item T) bool {
		return item == expected
	}
}

// OneOf returns a filter that returns true only if the input value matches
// one of the specified values.
func OneOf[T comparable](expected ...T) Func[T] {
	return func(item T) bool {
		return slices.Contains(expected, item)
	}
}

// Or returns true if any of the specified Func filters in the arguments
// return true. If no filters are specified as arguments, then the
// returned Func always returns true.
func Or[T any](filters ...Func[T]) Func[T] {
	if len(filters) == 0 {
		return True[T]()
	}
	return func(item T) bool {
		for _, filter := range filters {
			if filter(item) {
				return true
			}
		}
		return false
	}
}

// And returns true if all of the specified Func filters in the arguments
// return true. If no filters are specified as arguments, then the
// returned Func always returns true.
func And[T any](filters ...Func[T]) Func[T] {
	return func(item T) bool {
		for _, filter := range filters {
			if !filter(item) {
				return false
			}
		}
		return true
	}
}

// Slice filters the specified slice of entries and returns a new slice that
// contains only those elements that have passed the filter Func.
func Slice[T any](entries []T, filter Func[T]) []T {
	var result []T
	for _, entry := range entries {
		if filter(entry) {
			result = append(result, entry)
		}
	}
	return result
}
