package ds_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mokiat/gog/ds"
)

var _ = Describe("Set", func() {
	var (
		set *ds.Set[string]
	)

	BeforeEach(func() {
		set = ds.NewSet[string](0)
	})

	It("is empty by default", func() {
		Expect(set.IsEmpty()).To(BeTrue())
	})

	It("has zero size by default", func() {
		Expect(set.Size()).To(BeZero())
	})

	It("contains no items", func() {
		Expect(set.Items()).To(BeEmpty())
	})

	When("items are added", func() {
		BeforeEach(func() {
			set.Add("first")
			set.Add("second")
			set.Add("third")
		})

		It("is no longer empty", func() {
			Expect(set.IsEmpty()).To(BeFalse())
		})

		It("has the correct size", func() {
			Expect(set.Size()).To(Equal(3))
		})

		It("is possible to get all items", func() {
			Expect(set.Items()).To(ContainElements("first", "second", "third"))
		})

		It("is possible to check if an item is contained", func() {
			Expect(set.Contains("first")).To(BeTrue())
			Expect(set.Contains("missing")).To(BeFalse())
		})

		It("ignores add operations on existing items", func() {
			Expect(set.Add("second")).To(BeFalse())
			Expect(set.Size()).To(Equal(3))
		})

		It("ignores remove operations on missing items", func() {
			Expect(set.Remove("missing")).To(BeFalse())
		})

		When("clipped", func() {
			BeforeEach(func() {
				set.Clip()
			})

			It("still contains the same items", func() {
				Expect(set.Items()).To(ContainElements("first", "second", "third"))
			})
		})

		When("cleared", func() {
			BeforeEach(func() {
				set.Clear()
			})

			It("becomes empty", func() {
				Expect(set.IsEmpty()).To(BeTrue())
			})

			It("changes its size to zero", func() {
				Expect(set.Size()).To(BeZero())
			})

			It("no longer contains items", func() {
				Expect(set.Items()).To(BeEmpty())
			})
		})

		When("an item is removed", func() {
			BeforeEach(func() {
				Expect(set.Remove("second")).To(BeTrue())
			})

			It("changes its size accordingly", func() {
				Expect(set.Size()).To(Equal(2))
			})

			It("no longer contains the given item", func() {
				Expect(set.Items()).To(ContainElements("first", "third"))
			})
		})
	})
})
