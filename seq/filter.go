package seq

import "iter"

// Select applies the given predicate function to each element of the source
// sequence and returns a new sequence with the elements for which the predicate
// returned true.
func Select[T any](src iter.Seq[T], pred func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for item := range src {
			if !pred(item) {
				continue
			}
			if !yield(item) {
				return
			}
		}
	}
}
