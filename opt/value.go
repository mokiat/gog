package opt

// T represents an optional type of generic type D.
type T[D any] struct {

	// Specified indicates whether Value can be used.
	Specified bool

	// Value holds the actual value.
	Value D
}

// Unwrap is a convenience method that returns this optional in the standard
// Go tuple construct of (value, bool).
func (t T[D]) Unwrap() (D, bool) {
	return t.Value, t.Specified
}

// ValueOrDefault returns the value held in this optional if it is specified,
// otherwise it returns the given fallback value.
func (t T[D]) ValueOrDefault(fallback D) D {
	if t.Specified {
		return t.Value
	}
	return fallback
}

// ToPtr returns a pointer-based representation of the value held
// in this optional. If this optional is not specified, then nil
// is returned.
func (t T[D]) ToPtr() *D {
	if t.Specified {
		return &t.Value
	}
	return nil
}

// V returns a specified value of the generic type.
func V[D any](value D) T[D] {
	return T[D]{
		Specified: true,
		Value:     value,
	}
}

// Unspecified returns an unspecified value of the generic type.
func Unspecified[D any]() T[D] {
	return T[D]{
		Specified: false,
	}
}

// FromPtr returns an optional T value which is specified
// depending on whether the pointer can be dereferenced or not.
func FromPtr[D any](value *D) T[D] {
	result := T[D]{}
	if value != nil {
		result.Specified = true
		result.Value = *value
	}
	return result
}

// Wrap creates an optional from the standard Go tuple (value, bool).
//
// This can be useful if a function returns such a tuple and one wants to
// quickly convert it into an optional.
func Wrap[D any](value D, specified bool) T[D] {
	return T[D]{
		Value:     value,
		Specified: specified,
	}
}
