package gog_test

import (
	"testing"

	"github.com/mokiat/gog"
)

func BenchmarkIsOneOf(b *testing.B) {
	// NOTE: Using slices.Contains in the implementation is one idea slower
	// but at least it does not allocate.
	for b.Loop() {
		gog.IsOneOf(5, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	}
}
