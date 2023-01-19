package opt

// T represents an optional type.
//
// While it is possible to represent an optional through a pointer type in
// Go, there are a number of drawbacks:
//
//     - The value might be allocated on the heap due to the pointer reference
//     - The value or structs of such values will not be comparable
//
// The type T allows one to overcome such limitations of pointer references.
// Structs that are composed of such opt fields can often be safely used
// as map keys (as long as the underlying values are comparable).

type T[D any] struct {

	// Specified indicates whether Value can be used.
	Specified bool

	// Value holds the actual value.
	Value D
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
