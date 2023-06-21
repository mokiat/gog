// Package ds provides various data structure implementations with generics.
//
// Prior to Go 1.18, such data structures could not be implemented in a
// satisfactory way, as one had to use interfaces, which ofter resulted in
// unreadable code, memory allocations and conditions for unexpected panics.
//
// Now that Go has generics, this package provides some common data structures
// that make use of this new language feature.
package ds
