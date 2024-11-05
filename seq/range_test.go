package seq_test

import (
	"slices"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mokiat/gog/seq"
)

var _ = Describe("Range", func() {

	Describe("None", func() {
		It("yields no values", func() {
			iter := seq.None[int]()
			result := slices.Collect(iter)
			Expect(result).To(BeEmpty())
		})
	})

	Describe("Times", func() {
		It("yields the correct sequence of numbers", func() {
			iter := seq.Times(3)
			result := slices.Collect(iter)
			Expect(result).To(Equal([]int{0, 1, 2}))
		})
	})

	Describe("Range", func() {
		It("yields the correct sequence of numbers", func() {
			iter := seq.Range(1, 3)
			result := slices.Collect(iter)
			Expect(result).To(Equal([]int{1, 2, 3}))
		})

		It("works in reverse as well", func() {
			iter := seq.Range(3, 1)
			result := slices.Collect(iter)
			Expect(result).To(Equal([]int{3, 2, 1}))
		})
	})

})
