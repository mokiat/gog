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

	Describe("MapIndex", func() {
		mapFunc := func(index, v int) string {
			return strconv.Itoa(v * index)
		}

		It("converts from one slice type to another", func() {
			source := []int{1, 2, 3}
			target := gog.MapIndex(source, mapFunc)
			Expect(target).To(Equal([]string{
				"0", "2", "6",
			}))
		})

		It("preserves the nil slice", func() {
			Expect(gog.MapIndex(nil, mapFunc)).To(Equal([]string(nil)))
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

	Describe("Mapping", func() {
		It("partitions a slice into custom buckets", func() {
			source := []int{0, 1, 2, 3, 4, 5, 6}
			target := gog.Mapping(source, func(v int) (string, string) {
				if v%2 == 0 {
					return "even", strconv.Itoa(v)
				} else {
					return "odd", strconv.Itoa(v)
				}
			})
			Expect(target).To(Equal(map[string][]string{
				"even": {"0", "2", "4", "6"},
				"odd":  {"1", "3", "5"},
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

	Describe("Mutate", func() {
		double := func(v *int) {
			*v *= 2
		}

		It("mutates the items of a slice", func() {
			slice := []int{1, 2, 3, 4}
			gog.Mutate(slice, double)
			Expect(slice).To(Equal([]int{2, 4, 6, 8}))
		})

		It("ignores empty slices", func() {
			slice := []int{}
			gog.Mutate(slice, double)
			Expect(slice).To(Equal([]int{}))
		})

		It("ignores nil slices", func() {
			var slice []int
			gog.Mutate(slice, double)
			Expect(slice).To(Equal([]int(nil)))
		})
	})

	Describe("MutateIndex", func() {
		timesIndex := func(index int, v *int) {
			*v *= index
		}

		It("mutates the items of a slice", func() {
			slice := []int{1, 2, 3, 4}
			gog.MutateIndex(slice, timesIndex)
			Expect(slice).To(Equal([]int{0, 2, 6, 12}))
		})

		It("ignores empty slices", func() {
			slice := []int{}
			gog.MutateIndex(slice, timesIndex)
			Expect(slice).To(Equal([]int{}))
		})

		It("ignores nil slices", func() {
			var slice []int
			gog.MutateIndex(slice, timesIndex)
			Expect(slice).To(Equal([]int(nil)))
		})
	})

	Describe("FindFunc", func() {
		divisibleByFive := func(v int) bool {
			return v%5 == 0
		}

		It("returns the first matching element", func() {
			slice := []int{3, 4, 10, 7, 5, 8}
			element, found := gog.FindFunc(slice, divisibleByFive)
			Expect(found).To(BeTrue())
			Expect(element).To(Equal(10))
		})

		It("returns false if no element matches", func() {
			slice := []int{1, 8, 4, 13}
			_, found := gog.FindFunc(slice, divisibleByFive)
			Expect(found).To(BeFalse())
		})

		It("returns false for empty slices", func() {
			slice := []int{}
			_, found := gog.FindFunc(slice, divisibleByFive)
			Expect(found).To(BeFalse())
		})

		It("returns false for nil slices", func() {
			var slice []int
			_, found := gog.FindFunc(slice, divisibleByFive)
			Expect(found).To(BeFalse())
		})
	})

	Describe("FindFuncPtr", func() {
		divisibleByFive := func(v int) bool {
			return v%5 == 0
		}

		It("returns a pointer to the first matching element", func() {
			slice := []int{3, 4, 10, 7, 5, 8}
			element := gog.FindFuncPtr(slice, divisibleByFive)
			Expect(element).ToNot(BeNil())
			Expect(element).To(Equal(&slice[2]))
			Expect(element).ToNot(Equal(&slice[4]))
		})

		It("returns nil if no element matches", func() {
			slice := []int{1, 8, 4, 13}
			element := gog.FindFuncPtr(slice, divisibleByFive)
			Expect(element).To(BeNil())
		})

		It("returns nil for empty slices", func() {
			slice := []int{}
			element := gog.FindFuncPtr(slice, divisibleByFive)
			Expect(element).To(BeNil())
		})

		It("returns nil for nil slices", func() {
			var slice []int
			element := gog.FindFuncPtr(slice, divisibleByFive)
			Expect(element).To(BeNil())
		})
	})

	Describe("RefElements", func() {
		It("returns a slice of pointers", func() {
			slice := []int{1, 2, 3, 5, 6, 7}
			ptrSlice := gog.RefElements(slice)
			Expect(ptrSlice).To(HaveLen(len(slice)))
			Expect(*ptrSlice[0]).To(Equal(1))
			Expect(*ptrSlice[1]).To(Equal(2))
			Expect(*ptrSlice[2]).To(Equal(3))
			Expect(*ptrSlice[3]).To(Equal(5))
			Expect(*ptrSlice[4]).To(Equal(6))
			Expect(*ptrSlice[5]).To(Equal(7))
		})

		It("returns pointers to the actual elements", func() {
			slice := []int{1}
			ptrSlice := gog.RefElements(slice)
			Expect(ptrSlice).To(HaveLen(1))
			*ptrSlice[0] = 10
			Expect(slice[0]).To(Equal(10))
		})
	})

	Describe("DerefElements", func() {
		It("returns a slice of values", func() {
			first, second, third := 1, 2, 3
			slice := []*int{&first, &second, &third}
			valueSlice := gog.DerefElements(slice)
			Expect(valueSlice).To(HaveLen(len(slice)))
			Expect(valueSlice).To(ConsistOf(first, second, third))
		})
	})

	Describe("Concat", func() {
		It("returns a concatenation of slices", func() {
			first := []int{1, 2, 3}
			second := []int{4}
			third := []int{5, 6, 7}
			result := gog.Concat(first, second, third)
			Expect(result).To(Equal([]int{1, 2, 3, 4, 5, 6, 7}))
			Expect(result).To(HaveCap(7))
		})

		It("returns an empty slice when called without args", func() {
			result := gog.Concat[int]()
			Expect(result).To(Equal([]int{}))
		})
	})

	Describe("Merge", func() {
		It("returns the combination of maps", func() {
			first := map[int]string{
				1: "one",
				2: "two",
				3: "three",
			}
			second := map[int]string{
				1: "first",
				4: "fourth",
			}
			result := gog.Merge(first, second)
			Expect(result).To(Equal(map[int]string{
				1: "first",
				2: "two",
				3: "three",
				4: "fourth",
			}))
		})

		It("returns an empty map when called without args", func() {
			result := gog.Merge[int, string]()
			Expect(result).To(Equal(map[int]string{}))
		})
	})
})
