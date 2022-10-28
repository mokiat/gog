package gog_test

import (
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mokiat/gog"
)

var _ = Describe("MapReduce", func() {

	Describe("Map", func() {
		It("converts from one slice type to another", func() {
			source := []int{1, 2, 3}
			target := gog.Map(source, func(v int) string {
				return strconv.Itoa(v)
			})
			Expect(target).To(Equal([]string{
				"1", "2", "3",
			}))
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

})
