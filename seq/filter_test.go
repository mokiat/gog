package seq_test

import (
	"slices"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mokiat/gog/seq"
)

var _ = Describe("Filter", func() {

	Describe("Select", func() {
		isEven := func(value int) bool {
			return value%2 == 0
		}

		It("returns only the elements that match the predicate", func() {
			source := slices.Values([]int{1, 2, 3, 4, 5})
			target := seq.Select(source, isEven)
			items := slices.Collect(target)
			Expect(items).To(Equal([]int{2, 4}))
		})
	})

})
