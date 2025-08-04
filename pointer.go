package gog

// PtrOf returns a pointer to the passed value.
func PtrOf[T any](v T) *T {
	return &v
}

// ValueOf is the opposite of PtrOf. It takes a pointer and dereferences it.
// If the pointer is nil, then it uses the provided defaultValue.
func ValueOf[T any](v *T, defaultValue T) T {
	if v == nil {
		return defaultValue
	}
	return *v
}
