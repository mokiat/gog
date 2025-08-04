package filter_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mokiat/gog/filter"
)

var _ = Describe("Func", func() {

	Describe("True", func() {
		It("always returns true", func() {
			fltr := filter.True[string]()
			Expect(fltr("test")).To(BeTrue())
		})
	})

	Describe("False", func() {
		It("always returns false", func() {
			fltr := filter.False[string]()
			Expect(fltr("test")).To(BeFalse())
		})
	})

	Describe("Not", func() {
		It("flips the result", func() {
			fltr := filter.Not(filter.True[string]())
			Expect(fltr("test")).To(BeFalse())
		})
	})

	Describe("Equal", func() {
		It("returns true on matching value", func() {
			fltr := filter.Equal("a")
			Expect(fltr("a")).To(BeTrue())
		})

		It("returns false on non-matching value", func() {
			fltr := filter.Equal("a")
			Expect(fltr("b")).To(BeFalse())
		})
	})

	Describe("OneOf", func() {
		It("returns true on matching value", func() {
			fltr := filter.OneOf("a", "b", "c")
			Expect(fltr("b")).To(BeTrue())
		})

		It("returns false on non-matching value", func() {
			fltr := filter.OneOf("a", "b", "c")
			Expect(fltr("z")).To(BeFalse())
		})
	})

	Describe("Or", func() {
		It("returns true if the input value matches one of the conditions", func() {
			fltr := filter.Or(filter.Equal("a"), filter.Equal("b"))
			Expect(fltr("a")).To(BeTrue())
			Expect(fltr("b")).To(BeTrue())
		})

		It("returns false if the input value does not match on of the conditions", func() {
			fltr := filter.Or(filter.Equal("a"), filter.Equal("b"))
			Expect(fltr("c")).To(BeFalse())
		})

		It("returns true if the list of conditions is empty", func() {
			fltr := filter.Or[string]()
			Expect(fltr("irrelevant")).To(BeTrue())
		})
	})

	Describe("And", func() {
		It("returns true if all conditions match", func() {
			fltr := filter.And(filter.Equal("a"), filter.True[string]())
			Expect(fltr("a")).To(BeTrue())
		})

		It("returns false if one of the conditions does not match", func() {
			fltr := filter.And(filter.Equal("a"), filter.True[string]())
			Expect(fltr("b")).To(BeFalse())
		})

		It("returns true if the list of conditions is empty", func() {
			fltr := filter.And[string]()
			Expect(fltr("irrelevant")).To(BeTrue())
		})
	})

	Describe("Slice", func() {
		It("return a slice that contains only matching entries", func() {
			src := []string{
				"a", "1", "c", "b", "d",
			}
			fltr := filter.Or(
				filter.Equal("a"),
				filter.Equal("b"),
			)
			result := filter.Slice(src, fltr)
			Expect(result).To(Equal([]string{"a", "b"}))
		})
	})

})
