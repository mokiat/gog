// Package opt provides utilities for representing optional types.
//
// While it is possible to represent an optional type in Go through the usage
// of a pointer, there are a number of drawbacks with such an approach:
//
//   - The value might be allocated on the heap due to the pointer reference.
//   - The value or structs of such values will not be comparable.
//   - It becomes unreadable when trying to represent an optional pointer type.
//
// The generic type T in this package allows one to overcome the limitations of
// pointer references. Structs that are composed of such opt fields can often be
// safely used as map keys (as long as the underlying values are comparable).
package opt
