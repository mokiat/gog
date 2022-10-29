package opt

// T represents an optional type.
type T[D any] struct {

	// Specified indicates whether Value can be used.
	Specified bool

	// Value holds the actual value.
	Value D
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
