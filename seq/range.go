package seq

import "iter"

// None returns an empty sequence.
func None[T any]() iter.Seq[T] {
	return func(yield func(T) bool) {}
}

// Times returns a sequence of integers from 0 to count.
func Times(count int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := range count {
			if !yield(i) {
				return
			}
		}
	}
}

// Range returns a sequence of integers from from (inclusive) to to (inclusive).
//
// If from is greater than to, the sequence will be in descending order.
func Range(from, to int) iter.Seq[int] {
	return func(yield func(int) bool) {
		if from < to {
			for i := from; i <= to; i++ {
				if !yield(i) {
					return
				}
			}
		} else {
			for i := from; i >= to; i-- {
				if !yield(i) {
					return
				}
			}
		}
	}
}
