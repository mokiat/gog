package ds

import "golang.org/x/exp/maps"

// NewSet creates a new Set instance with the specified initial capacity,
// which is only used to preallocate memory and does not act as an upper bound.
func NewSet[T comparable](initialCapacity int) *Set[T] {
	return &Set[T]{
		items: make(map[T]struct{}, initialCapacity),
	}
}

// Set represents a set data structure, where only one instance
// of a given item is stored.
//
// Note: Using the standard map[T]struct{} approach will likely yield faster
// performance and should be preferred in performance-critical code. This
// type makes usage of sets more human-readable.
type Set[T comparable] struct {
	items map[T]struct{}
}

// IsEmpty returns whether this Set is empty.
func (s *Set[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items contained in this Set.
func (s *Set[T]) Size() int {
	return len(s.items)
}

// Add adds the specified item to this Set if it was not present already.
// This method returns true if the operation was performed and false if
// the item was aleady present.
func (s *Set[T]) Add(item T) bool {
	if s.Contains(item) {
		return false
	}
	s.items[item] = struct{}{}
	return true
}

// Remove removes the specified item from this Set.
// This method returns true if there was in fact such an item to be removed
// and false otherwise.
func (s *Set[T]) Remove(item T) bool {
	if !s.Contains(item) {
		return false
	}
	delete(s.items, item)
	return true
}

// Contains returns whether this Set holds the specified item.
func (s *Set[T]) Contains(item T) bool {
	_, ok := s.items[item]
	return ok
}

// Items returns a slice containing all of the items from this Set.
//
// Note: The items are returned in a random order which can differ
// between subsequent calls.
func (s *Set[T]) Items() []T {
	result := make([]T, 0, len(s.items))
	for v := range s.items {
		result = append(result, v)
	}
	return result
}

// Clear removes all items from this Set.
func (s *Set[T]) Clear() {
	for v := range s.items {
		delete(s.items, v)
	}
}

// Clip removes unused capacity from the Set.
func (s *Set[T]) Clip() {
	s.items = maps.Clone(s.items)
}
