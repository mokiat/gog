// Package filter adds helper functions for performing filtering
// over a data set.
//
// There are times when one needs to filter a slice of values but the exact
// filtering logic cannot be hardcoded in advance and is conditional on
// some external flags or parameters. In such cases, it is useful to be able to
// dynamically construct a filtering expression that can pick values from some
// data set.
//
// This can be achieved through the usage of the Slice function
// and the set of predefined filtering expressions like Or, And, Not and more,
// where custom ones can be plugged in as well.
package filter
