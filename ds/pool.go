package ds

// NewPool creates a new Pool instance.
func NewPool[T any]() *Pool[T] {
	return &Pool[T]{
		items: NewStack[*T](0),
	}
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
