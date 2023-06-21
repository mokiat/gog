package filter_test

import (
	"fmt"
	"strings"

	"github.com/mokiat/gog/filter"
)

func Example() {
	items := []string{
		"hello universe",
		"big world",
		"hello strange world",
		"will be dropped",
		"oh well",
		"let me through",
	}

	hasHelloPrefix := func(candidate string) bool {
		return strings.HasPrefix(candidate, "hello")
	}
	hasWorldSuffix := func(candidate string) bool {
		return strings.HasSuffix(candidate, "world")
	}

	filteredItems := filter.Slice(items,
		filter.Or(
			filter.And(
				hasHelloPrefix,
				hasWorldSuffix,
			),
			filter.Equal("let me through"),
		),
	)
	fmt.Printf("%#v\n", filteredItems)

	// Output:
	// []string{"hello strange world", "let me through"}
}
