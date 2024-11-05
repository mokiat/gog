package seq

import (
	"iter"
	"slices"
)

// CollectCap collects values from src into a new slice with the given capacity
// preallocated and returns it.
func CollectCap[T any](src iter.Seq[T], cap int) []T {
	result := make([]T, 0, cap)
	result = slices.AppendSeq(result, src)
	return result
}
