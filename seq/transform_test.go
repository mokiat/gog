package seq_test

import (
	"slices"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mokiat/gog/seq"
)

var _ = Describe("Transform", func() {

	Describe("Map", func() {
		It("transforms the sequence", func() {
			source := slices.Values([]int{0, 1, 2, 3})
			target := seq.Map(source, strconv.Itoa)
			result := slices.Collect(target)
			Expect(result).To(Equal([]string{"0", "1", "2", "3"}))
		})
	})

	Describe("BatchSlice", func() {
		intEq := func(a, b int) bool {
			return a == b
		}

		It("groups the elements into batches", func() {
			source := []int{0, 0, 1, 1, 1, 2, 3, 3, 3, 3}
			target := seq.BatchSlice(source, intEq, 0)
			result := slices.Collect(target)
			Expect(result).To(Equal([][]int{
				{0, 0},
				{1, 1, 1},
				{2},
				{3, 3, 3, 3},
			}))
		})

		It("respects the max size", func() {
			source := []int{0, 0, 1, 1, 1, 2, 3, 3, 3, 3}
			target := seq.BatchSlice(source, intEq, 2)
			result := slices.Collect(target)
			Expect(result).To(Equal([][]int{
				{0, 0},
				{1, 1}, {1},
				{2},
				{3, 3}, {3, 3},
			}))
		})

		It("handles empty source", func() {
			source := []int{}
			target := seq.BatchSlice(source, intEq, 0)
			result := slices.Collect(target)
			Expect(result).To(Equal([][]int(nil)))
		})

		It("handles single element source", func() {
			source := []int{5}
			target := seq.BatchSlice(source, intEq, 0)
			result := slices.Collect(target)
			Expect(result).To(Equal([][]int{{5}}))
		})

		It("handles unit max size", func() {
			source := []int{0, 0, 1, 1, 1, 2, 3, 3, 3, 3}
			target := seq.BatchSlice(source, intEq, 1)
			result := slices.Collect(target)
			Expect(result).To(Equal([][]int{
				{0}, {0},
				{1}, {1}, {1},
				{2},
				{3}, {3}, {3}, {3},
			}))
		})

		It("handles max size larger than source", func() {
			source := []int{0, 0, 1, 1, 1, 2, 3, 3, 3, 3}
			target := seq.BatchSlice(source, intEq, 5000)
			result := slices.Collect(target)
			Expect(result).To(Equal([][]int{
				{0, 0},
				{1, 1, 1},
				{2},
				{3, 3, 3, 3},
			}))
		})
	})

	Describe("Reduce", func() {
		It("reduces a sequence to a single value", func() {
			source := slices.Values([]int{1, 2, 3})
			result := seq.Reduce(source, ">", func(accum string, value int) string {
				return accum + strconv.Itoa(value)
			})
			Expect(result).To(Equal(">123"))
		})
	})

	Describe("Sum", func() {
		type IntWrapper int

		It("calculates the sum of the sequence", func() {
			source := slices.Values([]IntWrapper{0, 1, 2, 3})
			result := seq.Sum(source)
			Expect(result).To(Equal(IntWrapper(6)))
		})

		It("returns the zero value for empty sequence", func() {
			source := slices.Values([]int{})
			result := seq.Sum(source)
			Expect(result).To(Equal(0))
		})
	})
})
