package ds

import "golang.org/x/exp/slices"

// NewList creates a new List with the given capacity. The capacity can be
// used as a form of optimization. Regardless of the value, the initial size
// of the list is zero and the list can grow past the specified capacity.
func NewList[T comparable](initialCapacity int) *List[T] {
	return &List[T]{}
}

// ListFromSlice constructs a new List that is based on the items from the
// specified slice.
//
// It is safe to modify the slice afterwards, as the list creates its own
// internal copy.
func ListFromSlice[T comparable](items []T) *List[T] {
	return &List[T]{
		items: slices.Clone(items),
	}
}

// List represents a sequence of items.
//
// Currently a List can only store comparable items. This restriction
// allows for Remove, IndexOf and Contains operations.
type List[T comparable] struct {
	items []T
}

// Size returns the number of items contained in this List.
func (l *List[T]) Size() int {
	return len(l.items)
}

// IsEmpty returns whether this list has no elements.
func (l *List[T]) IsEmpty() bool {
	return len(l.items) == 0
}

// Add appends the specified item to the List.
func (l *List[T]) Add(item T) {
	l.items = append(l.items, item)
}

// Remove removes the specified item from this List and returns true. If the
// item is not contained by this List, then false is returned.
func (l *List[T]) Remove(item T) bool {
	index := l.IndexOf(item)
	if index < 0 {
		return false
	}
	l.items = slices.Delete(l.items, index, index+1)
	return true
}

// Get returns the item in this list that is located at the specified index
// (starting from zero).
//
// This method will panic if the index is outside the list bounds.
func (l *List[T]) Get(index int) T {
	return l.items[index]
}

// Set modifies the item at the specified index.
//
// This method will panic if the index is outside the list bounds.
func (l *List[T]) Set(index int, value T) {
	l.items[index] = value
}

// Unbox provides direct access to the inner representation of the list.
// The returned slice should not be modified, otherwise there is a risk that
// the List might not work correctly afterwards. Even if it works now, a future
// version might break that behavior.
//
// This method should only be used when performance is critical and memory
// allocation is not desired.
func (l *List[T]) Unbox() []T {
	return l.items
}

// Items returns all items stored in this List as a slice. It is safe to mutate
// the returned slice as it is a copy of the inner representation.
//
// If performance is needed, consider using Unbox method instead.
func (l *List[T]) Items() []T {
	return slices.Clone(l.items)
}

// Contains checks whether this List has the specified item and returns true
// if it is contained and false otherwise.
func (l *List[T]) Contains(item T) bool {
	return l.IndexOf(item) >= 0
}

// IndexOf returns the index where the specified item is located in this List.
// If the item is not part of this list, this method returns -1.
func (l *List[T]) IndexOf(item T) int {
	return slices.Index(l.items, item)
}

// Each is a helper method allows one to iterate over all items in this List
// through a closure function.
func (l *List[T]) Each(iterator func(item T)) {
	for _, item := range l.items {
		iterator(item)
	}
}

// Equals returns whether this list matches exactly the provided list.
func (l *List[T]) Equals(other *List[T]) bool {
	return slices.Equal(l.items, other.items)
}

// Clear removes all items from this List.
func (l *List[T]) Clear() {
	l.items = l.items[:0]
}

// Clip removes unused capacity from the List.
func (l *List[T]) Clip() {
	l.items = slices.Clip(l.items)
}
