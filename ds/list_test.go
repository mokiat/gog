package ds_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mokiat/gog/ds"
)

var _ = Describe("List", func() {
	var (
		list *ds.List[string]
	)

	BeforeEach(func() {
		list = ds.NewList[string](0)
	})

	It("is empty by default", func() {
		Expect(list.IsEmpty()).To(BeTrue())
	})

	It("has zero size by default", func() {
		Expect(list.Size()).To(BeZero())
	})

	It("contains no items", func() {
		Expect(list.Items()).To(BeEmpty())
	})

	When("items are added", func() {
		BeforeEach(func() {
			list.Add("first")
			list.Add("second")
			list.Add("third")
		})

		It("is no longer empty", func() {
			Expect(list.IsEmpty()).To(BeFalse())
		})

		It("has the correct size", func() {
			Expect(list.Size()).To(Equal(3))
		})

		It("is possible to fetch the items by index", func() {
			Expect(list.Get(0)).To(Equal("first"))
			Expect(list.Get(1)).To(Equal("second"))
			Expect(list.Get(2)).To(Equal("third"))
		})

		It("is possible to get all items", func() {
			Expect(list.Items()).To(Equal([]string{
				"first", "second", "third",
			}))
		})

		It("is possible to check if an item is contained", func() {
			Expect(list.Contains("first")).To(BeTrue())
			Expect(list.Contains("missing")).To(BeFalse())
		})

		It("is possible to get the index of an item", func() {
			Expect(list.IndexOf("second")).To(Equal(1))
			Expect(list.IndexOf("missing")).To(Equal(-1))
		})

		It("ignores remove operations on missing items", func() {
			Expect(list.Remove("missing")).To(BeFalse())
		})

		It("is possible to iterate over the list", func() {
			seen := make([]string, 0)
			list.Each(func(item string) {
				seen = append(seen, item)
			})
			Expect(seen).To(Equal([]string{
				"first", "second", "third",
			}))
		})

		When("the list is clipped", func() {
			BeforeEach(func() {
				list.Clip()
			})

			It("still contains the same items", func() {
				Expect(list.Items()).To(Equal([]string{
					"first", "second", "third",
				}))
			})
		})

		When("the list is cleared", func() {
			BeforeEach(func() {
				list.Clear()
			})

			It("becomes empty", func() {
				Expect(list.IsEmpty()).To(BeTrue())
			})

			It("changes its size to zero", func() {
				Expect(list.Size()).To(BeZero())
			})

			It("no longer contains items", func() {
				Expect(list.Items()).To(BeEmpty())
			})
		})

		When("an item is removed", func() {
			BeforeEach(func() {
				Expect(list.Remove("second")).To(BeTrue())
			})

			It("changes its size accordingly", func() {
				Expect(list.Size()).To(Equal(2))
			})

			It("no longer contains the given item", func() {
				Expect(list.Items()).To(Equal([]string{
					"first", "third",
				}))
			})
		})
	})
})
