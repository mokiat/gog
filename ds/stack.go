package ds

import "slices"

// EmptyStack creates an empty Stack instance.
func EmptyStack[T any]() *Stack[T] {
	return PreallocatedStack[T](0)
}

// PreallocatedStack creates a new Stack instance with the specified initial
// capacity, which is only used to preallocate memory and does not act as
// an upper bound.
func PreallocatedStack[T any](initialCapacity int) *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0, initialCapacity),
	}
}

// NewStack creates a new Stack instance with the specified initial capacity,
// which only serves to preallocate memory. Exceeding the initial capacity is
// allowed.
//
// Deprecated: Use EmptyStack or PreallocatedStack instead.
//
//go:fix inline
func NewStack[T any](initialCapacity int) *Stack[T] {
	return PreallocatedStack[T](initialCapacity)
}

// Stack is an implementation of a stack data structure. The last inserted
// item is the first one to be removed (LIFO - last in, first out).
type Stack[T any] struct {
	items []T
}

// Size returns the number of items stored in this Stack.
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// IsEmpty returns true if there are no more items in this Stack.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Push adds an item to the top of this Stack.
func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

// Pop removes the item from the top of this Stack and returns it.
// This function panics if there are no more items. Use IsEmpty to check
// for that.
func (s *Stack[T]) Pop() T {
	count := len(s.items)
	result := s.items[count-1]
	s.items = s.items[:count-1]
	return result
}

// Peek returns the item that is at the top of the Stack without removing it.
// Make sure that the Stack is not empty, otherwise this method will panic.
func (s *Stack[T]) Peek() T {
	return s.items[len(s.items)-1]
}

// Clear removes all items from this Stack.
func (s *Stack[T]) Clear() {
	s.items = s.items[:0]
}

// Clip removes unused capacity from the Stack.
func (s *Stack[T]) Clip() {
	s.items = slices.Clip(s.items)
}
