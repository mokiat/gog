package ds

// EmptyPool creates a new Pool instance.
func EmptyPool[T any]() *Pool[T] {
	return PreallocatedPool[T](0)
}

// PreallocatedPool creates a new Pool instance with a preallocated capacity,
// which is only used to preallocate memory and does not act as an upper bound.
func PreallocatedPool[T any](initialCapacity int) *Pool[T] {
	pool := &Pool[T]{
		items: PreallocatedStack[*T](initialCapacity),
	}
	for range initialCapacity {
		pool.items.Push(new(T))
	}
	return pool
}

// NewPool creates a new Pool instance.
//
// Deprecated: Use EmptyPool or PreallocatedPool instead.
//
//go:fix inline
func NewPool[T any]() *Pool[T] {
	return EmptyPool[T]()
}

// Pool represents a storage structure that can preserve allocated objects
// for faster reuse.
type Pool[T any] struct {
	items *Stack[*T]
}

// IsEmpty returns true if there is nothing stored for reuse in this pool.
func (p *Pool[T]) IsEmpty() bool {
	return p.items.IsEmpty()
}

// Clear removes any items that were stored for reuse.
func (p *Pool[T]) Clear() {
	p.items.Clear()
}

// Fetch retrieves an available item from the pool or creates a new one
// if one is not available.
func (p *Pool[T]) Fetch() *T {
	if p.items.IsEmpty() {
		return new(T)
	}
	return p.items.Pop()
}

// Restore returns an item to the pool to be reused.
func (p *Pool[T]) Restore(v *T) {
	p.items.Push(v)
}
