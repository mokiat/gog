package gog_test

import (
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mokiat/gog"
)

var _ = Describe("Slice", func() {

	Describe("Map", func() {
		mapFunc := func(v int) string {
			return strconv.Itoa(v)
		}

		It("converts from one slice type to another", func() {
			source := []int{1, 2, 3}
			target := gog.Map(source, mapFunc)
			Expect(target).To(Equal([]string{
				"1", "2", "3",
			}))
		})

		It("preserves the nil slice", func() {
			Expect(gog.Map(nil, mapFunc)).To(Equal([]string(nil)))
		})
	})

	Describe("Reduce", func() {
		It("reduces a slice to a single value", func() {
			source := []int{1, 2, 3}
			target := gog.Reduce(source, ">", func(accum string, value int) string {
				return accum + strconv.Itoa(value)
			})
			Expect(target).To(Equal(">123"))
		})
	})

	Describe("Select", func() {
		selectFunc := func(v int) bool {
			return v%2 == 0
		}

		It("returns a slice of desired elements", func() {
			source := []int{0, 1, 2, 3, 4, 5, 6}
			target := gog.Select(source, selectFunc)
			Expect(target).To(Equal([]int{0, 2, 4, 6}))
		})

		It("preserves the nil slice", func() {
			Expect(gog.Select(nil, selectFunc)).To(Equal([]int(nil)))
		})
	})

	Describe("Partition", func() {
		It("partitions a slice", func() {
			source := []int{0, 1, 2, 3, 4, 5, 6}
			target := gog.Partition(source, func(v int) string {
				if v%2 == 0 {
					return "even"
				} else {
					return "odd"
				}
			})
			Expect(target).To(Equal(map[string][]int{
				"even": {0, 2, 4, 6},
				"odd":  {1, 3, 5},
			}))
		})
	})

	Describe("Dedupe", func() {
		It("returns a slice of distinct elements", func() {
			source := []int{0, 0, 2, 3, 3, 5, 6, 6}
			target := gog.Dedupe(source)
			Expect(target).To(Equal([]int{
				0, 2, 3, 5, 6,
			}))
		})

		It("preserves the nil slice", func() {
			Expect(gog.Dedupe[int](nil)).To(Equal([]int(nil)))
		})
	})

	Describe("Flatten", func() {
		It("returns a flat slice", func() {
			source := [][]int{
				{1, 2, 5, 8, 8},
				{1, 11},
			}
			target := gog.Flatten(source)
			Expect(target).To(Equal([]int{
				1, 2, 5, 8, 8, 1, 11,
			}))
		})

		It("preserves the nil slice", func() {
			Expect(gog.Flatten[int](nil)).To(Equal([]int(nil)))
		})
	})

})
