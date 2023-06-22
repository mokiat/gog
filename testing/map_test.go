package gog_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mokiat/gog"
)

var _ = Describe("Map", func() {

	Describe("Entries", func() {
		It("returns the map entries", func() {
			result := gog.Entries(map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
			})
			Expect(result).To(ConsistOf(
				gog.KV[string, int]{
					Key:   "one",
					Value: 1,
				},
				gog.KV[string, int]{
					Key:   "two",
					Value: 2,
				},
				gog.KV[string, int]{
					Key:   "three",
					Value: 3,
				},
			))
		})
	})

})
