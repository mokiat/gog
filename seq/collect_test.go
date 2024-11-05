package seq_test

import (
	"slices"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mokiat/gog/seq"
)

var _ = Describe("Collect", func() {

	Describe("CollectCap", func() {

		It("collects values from sequence into new slice with given capacity", func() {
			iter := slices.Values([]int{1, 2, 3})
			result := seq.CollectCap(iter, 5)
			Expect(result).To(Equal([]int{1, 2, 3}))
			Expect(result).To(HaveCap(5))
		})

	})

})
