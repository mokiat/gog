package seq

import "iter"

// Map applies the given transformation function to each element of the source
// sequence and returns a new sequence with the results.
func Map[T any, S any](src iter.Seq[S], fn func(S) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range src {
			if !yield(fn(v)) {
				return
			}
		}
	}
}

// BatchSlice groups the elements of the source sequence into batches with the
// same key, as determined by the key function. In order for this function to
// work best, it is assumed that items are already sorted acoording to the key
// function.
//
// The maxSize parameter can be used to limit the size of the batches. If the
// maxSize is 0 or negative, then the batches will be of max possible size.
func BatchSlice[T any](items []T, eqFunc func(a, b T) bool, maxSize int) iter.Seq[[]T] {
	return BatchSliceFast(items, func(items []T, i, j int) bool {
		return eqFunc(items[i], items[j])
	}, maxSize)
}

// BatchSliceFast is the same as BatchSlice, but it uses a much more performant
// equality function, which allows one to work with references to items in the
// slice instead of copies. This can have a huge impact when the items are
// large.
func BatchSliceFast[T any](items []T, eqFunc func(items []T, i, j int) bool, maxSize int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		var (
			batchOffset int
			batchSize   int
		)
		flush := func() bool {
			output := items[batchOffset : batchOffset+batchSize]
			batchOffset += batchSize
			batchSize = 0
			return yield(output)
		}
		for i := range items {
			if batchSize == 0 || eqFunc(items, i, i-1) {
				batchSize++
				if maxSize > 0 && batchSize >= maxSize {
					if !flush() {
						return
					}
				}
			} else {
				if !flush() {
					return
				}
				batchSize++ // the i-th element has been read
			}
		}
		if batchSize > 0 {
			if !flush() {
				return
			}
		}
	}
}
