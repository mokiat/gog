package ds

import "golang.org/x/exp/slices"

// NewHeap creates a new Heap instance that is configured to use the
// specified better function to order items. When better returns true, the
// first argument will be placed higher in the heap.
// The specified initialCapacity is used to preallocate memory.
func NewHeap[T any](initialCapacity int, better func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		better: better,
		items:  make([]T, 0, initialCapacity),
	}
}

// Heap is a data structure that orders items when inserted according to a
// specified ordering.
type Heap[T any] struct {
	better func(a, b T) bool
	items  []T
}

// IsEmpty returns true if there are no items in this Heap and false otherwise.
func (h *Heap[T]) IsEmpty() bool {
	return len(h.items) == 0
}

// Size returns the number of items stored in this Heap.
func (h *Heap[T]) Size() int {
	return len(h.items)
}

// Push adds a new item to this Heap.
func (h *Heap[T]) Push(value T) {
	h.items = append(h.items, value)
	h.siftUp(value, len(h.items)-1)
}

// Pop removes the top-most item from this Heap and returns it.
// This method panics if the Heap is empty so make sure to use IsEmpty
// beforehand.
func (h *Heap[T]) Pop() T {
	result := h.items[0]
	if len(h.items) > 1 {
		h.items[0] = h.items[len(h.items)-1]
		h.items = h.items[:len(h.items)-1]
		h.siftDown(h.items[0], 0)
	} else {
		h.items = h.items[:len(h.items)-1]
	}
	return result
}

// Peek returns the top-most item from this Heap without removing it.
// This method panics if the Heap is empty so make sure to use IsEmpty
// beforehand.
func (h *Heap[T]) Peek() T {
	return h.items[0]
}

// Clear removes all items from this Heap.
func (h *Heap[T]) Clear() {
	h.items = h.items[:0]
}

// Clip removes unused capacity from the Heap.
func (s *Heap[T]) Clip() {
	s.items = slices.Clip(s.items)
}

func (h *Heap[T]) siftUp(value T, index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		parentValue := h.items[parentIndex]
		if !h.better(value, parentValue) {
			return
		}
		h.items[index] = parentValue
		h.items[parentIndex] = value
		index = parentIndex
	}
}

func (h *Heap[T]) siftDown(value T, index int) {
	leftChildIndex := index*2 + 1
	for leftChildIndex < len(h.items) {
		bestIndex := index

		leftValue := h.items[leftChildIndex]
		if h.better(leftValue, value) {
			bestIndex = leftChildIndex
		}

		rightChildIndex := leftChildIndex + 1
		if rightChildIndex < len(h.items) {
			rightValue := h.items[rightChildIndex]
			if h.better(rightValue, value) && h.better(rightValue, leftValue) {
				bestIndex = rightChildIndex
			}
		}

		if bestIndex == index {
			return
		}

		h.items[index] = h.items[bestIndex]
		h.items[bestIndex] = value

		index = bestIndex
		leftChildIndex = index*2 + 1
	}
}
