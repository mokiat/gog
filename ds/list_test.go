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

	It("equals an empty list", func() {
		other := ds.NewList[string](0)
		Expect(list.Equals(other)).To(BeTrue())
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

		It("is possible to unbox the list", func() {
			items := list.Unbox()
			Expect(items).To(Equal([]string{
				"first", "second", "third",
			}))

			items[0] = "modified"
			Expect(list.Items()).To(Equal([]string{
				"modified", "second", "third",
			}))
		})

		It("is possible to get all items", func() {
			items := list.Items()
			Expect(items).To(Equal([]string{
				"first", "second", "third",
			}))

			items[0] = "modified"
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

		It("equals another list with same items", func() {
			other := ds.ListFromSlice([]string{"first", "second", "third"})
			Expect(list.Equals(other)).To(BeTrue())
		})

		It("does not equal another list with reordered items", func() {
			other := ds.ListFromSlice([]string{"second", "first", "third"})
			Expect(list.Equals(other)).To(BeFalse())
		})

		It("does not equal another list with more items", func() {
			other := ds.ListFromSlice([]string{"first", "second", "third", "fourth"})
			Expect(list.Equals(other)).To(BeFalse())
		})

		It("does not equal another list with fewer items", func() {
			other := ds.ListFromSlice([]string{"first", "second"})
			Expect(list.Equals(other)).To(BeFalse())
		})

		When("an item is overwritten", func() {
			BeforeEach(func() {
				list.Set(1, "modified")
			})

			It("is reflected in the items", func() {
				Expect(list.Items()).To(Equal([]string{"first", "modified", "third"}))
			})
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

	When("constructed from a slice", func() {
		BeforeEach(func() {
			list = ds.ListFromSlice([]string{"a", "b", "c"})
		})

		It("has the correct size", func() {
			Expect(list.Size()).To(Equal(3))
		})

		It("contains the elements of the slice", func() {
			Expect(list.Items()).To(Equal([]string{
				"a", "b", "c",
			}))
		})

		When("the slice is nil", func() {
			BeforeEach(func() {
				var slice []string
				list = ds.ListFromSlice(slice)
			})

			It("is empty", func() {
				Expect(list.IsEmpty()).To(BeTrue())
			})
		})
	})
})
