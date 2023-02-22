package ds

import "golang.org/x/exp/maps"

// NewSet creates a new Set instance with the specified initial capacity,
// which is only used to preallocate memory and does not act as an upper bound.
func NewSet[T comparable](initialCapacity int) *Set[T] {
	return &Set[T]{
		items: make(map[T]struct{}, initialCapacity),
	}
}

// SetFromSlice creates a new Set instance based on the elements contained
// in the provided slice.
func SetFromSlice[T comparable](slice []T) *Set[T] {
	result := NewSet[T](len(slice))
	for _, item := range slice {
		result.items[item] = struct{}{}
	}
	return result
}

// SetFromMapKeys creates a new Set instance based on the keys of the
// provided map.
func SetFromMapKeys[T comparable, V any](m map[T]V) *Set[T] {
	result := NewSet[T](len(m))
	for key := range m {
		result.items[key] = struct{}{}
	}
	return result
}

// SetFromMapValues creates a new Set instance based on the values of the
// provided map.
func SetFromMapValues[K comparable, V comparable](m map[K]V) *Set[V] {
	result := NewSet[V](len(m))
	for _, value := range m {
		result.items[value] = struct{}{}
	}
	return result
}

// SetUnion creates a new Set that is the union of the specified sets.
func SetUnion[T comparable](first, second *Set[T]) *Set[T] {
	result := NewSet[T](first.Size() + second.Size())
	for item := range first.items {
		result.items[item] = struct{}{}
	}
	for item := range second.items {
		result.items[item] = struct{}{}
	}
	return result
}

// SetDifference creates a new Set that holds the difference between the
// first and the second specified sets.
func SetDifference[T comparable](first, second *Set[T]) *Set[T] {
	result := NewSet[T](first.Size())
	for item := range first.items {
		if _, ok := second.items[item]; !ok {
			result.items[item] = struct{}{}
		}
	}
	return result
}

// SetIntersection creates a new Set that holds the intersection of the
// items of the two specified sets.
func SetIntersection[T comparable](first, second *Set[T]) *Set[T] {
	result := NewSet[T](0)
	for item := range first.items {
		if _, ok := second.items[item]; ok {
			result.items[item] = struct{}{}
		}
	}
	return result
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

// AddSet adds the items of another Set to this Set.
// The operation returns true if the operation resulted in a change and
// false otherwise.
func (s *Set[T]) AddSet(other *Set[T]) bool {
	var changed bool
	for item := range other.items {
		changed = changed || s.Add(item)
	}
	return changed
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

// RemoveSet removes the items of another Set from this Set.
// The operation returns true if the operation resulted in a change and
// false otherwise.
func (s *Set[T]) RemoveSet(other *Set[T]) bool {
	var changed bool
	for item := range other.items {
		changed = changed || s.Remove(item)
	}
	return changed
}

// Contains returns whether this Set holds the specified item.
func (s *Set[T]) Contains(item T) bool {
	_, ok := s.items[item]
	return ok
}

// ContainsSet returns whether this set fully contains another set.
func (s *Set[T]) ContainsSet(other *Set[T]) bool {
	for item := range other.items {
		if !s.Contains(item) {
			return false
		}
	}
	return true
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
