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

	It("equals an empty set", func() {
		other := ds.NewSet[string](0)
		Expect(set.Equals(other)).To(BeTrue())
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

		It("is possible to unbox the set", func() {
			items := set.Unbox()
			Expect(items).To(Equal(map[string]struct{}{
				"first":  {},
				"second": {},
				"third":  {},
			}))

			delete(items, "first")
			Expect(set.Unbox()).To(Equal(map[string]struct{}{
				"second": {},
				"third":  {},
			}))
		})

		It("is possible to get all items", func() {
			Expect(set.Items()).To(ContainElements("first", "second", "third"))
		})

		It("is possible to check if an item is contained", func() {
			Expect(set.Contains("first")).To(BeTrue())
			Expect(set.Contains("missing")).To(BeFalse())
		})

		It("is possible to check if a whole different set is contained", func() {
			subSet := ds.NewSet[string](2)
			subSet.Add("first")
			subSet.Add("third")
			Expect(set.ContainsSet(subSet)).To(BeTrue())

			mismatchSet := ds.NewSet[string](3)
			mismatchSet.Add("first")
			mismatchSet.Add("third")
			mismatchSet.Add("fifth")
			Expect(set.ContainsSet(mismatchSet)).To(BeFalse())
		})

		It("ignores add operations on existing items", func() {
			Expect(set.Add("second")).To(BeFalse())
			Expect(set.Size()).To(Equal(3))
		})

		It("ignores remove operations on missing items", func() {
			Expect(set.Remove("missing")).To(BeFalse())
		})

		It("equals another set with same items", func() {
			other := ds.SetFromSlice([]string{"first", "second", "third"})
			Expect(set.Equals(other)).To(BeTrue())
		})

		It("does not equal another set with additional items", func() {
			other := ds.SetFromSlice([]string{"first", "second", "third", "extra"})
			Expect(set.Equals(other)).To(BeFalse())
		})

		It("does not equal another set with insufficient items", func() {
			other := ds.SetFromSlice([]string{"first", "third"})
			Expect(set.Equals(other)).To(BeFalse())
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

		When("another set is added", func() {
			BeforeEach(func() {
				other := ds.NewSet[string](3)
				other.Add("second")
				other.Add("third")
				other.Add("fourth")
				other.Add("fifth")
				Expect(set.AddSet(other)).To(BeTrue())
			})

			It("changes its size accordingly", func() {
				Expect(set.Size()).To(Equal(5))
			})

			It("contains the union of the sets", func() {
				Expect(set.Contains("first")).To(BeTrue())
				Expect(set.Contains("second")).To(BeTrue())
				Expect(set.Contains("third")).To(BeTrue())
				Expect(set.Contains("fourth")).To(BeTrue())
				Expect(set.Contains("fifth")).To(BeTrue())
			})

			It("returns false if the set is already contained", func() {
				other := ds.NewSet[string](2)
				other.Add("second")
				other.Add("third")
				Expect(set.AddSet(other)).To(BeFalse())
			})
		})

		When("another set is removed", func() {
			BeforeEach(func() {
				other := ds.NewSet[string](2)
				other.Add("second")
				other.Add("third")
				other.Add("fourth")
				Expect(set.RemoveSet(other)).To(BeTrue())
			})

			It("changes its size accordingly", func() {
				Expect(set.Size()).To(Equal(1))
			})

			It("contains the difference of the sets", func() {
				Expect(set.Contains("first")).To(BeTrue())
			})

			It("returns false if the set is not contained", func() {
				other := ds.NewSet[string](2)
				other.Add("fourth")
				other.Add("fifth")
				Expect(set.RemoveSet(other)).To(BeFalse())
			})
		})
	})

	When("constructed from a slice", func() {
		BeforeEach(func() {
			set = ds.SetFromSlice([]string{"a", "c", "c", "b", "a", "d"})
		})

		It("has the correct size", func() {
			Expect(set.Size()).To(Equal(4))
		})

		It("contains the elements of the slice", func() {
			Expect(set.Contains("a")).To(BeTrue())
			Expect(set.Contains("b")).To(BeTrue())
			Expect(set.Contains("c")).To(BeTrue())
			Expect(set.Contains("d")).To(BeTrue())
		})

		When("the slice is nil", func() {
			BeforeEach(func() {
				var slice []string
				set = ds.SetFromSlice(slice)
			})

			It("is empty", func() {
				Expect(set.IsEmpty()).To(BeTrue())
			})
		})
	})

	When("constructed from a map's keys", func() {
		BeforeEach(func() {
			set = ds.SetFromMapKeys(map[string]int{
				"a": 1,
				"c": 5,
				"b": 13,
				"d": 31,
			})
		})

		It("has the correct size", func() {
			Expect(set.Size()).To(Equal(4))
		})

		It("contains the elements of the slice", func() {
			Expect(set.Contains("a")).To(BeTrue())
			Expect(set.Contains("b")).To(BeTrue())
			Expect(set.Contains("c")).To(BeTrue())
			Expect(set.Contains("d")).To(BeTrue())
		})

		When("the map is nil", func() {
			BeforeEach(func() {
				var m map[string]int
				set = ds.SetFromMapKeys(m)
			})

			It("is empty", func() {
				Expect(set.IsEmpty()).To(BeTrue())
			})
		})
	})

	When("constructed from a map's values", func() {
		BeforeEach(func() {
			set = ds.SetFromMapValues(map[int]string{
				1: "a",
				2: "c",
				3: "c",
				4: "b",
				5: "a",
				6: "d",
			})
		})

		It("has the correct size", func() {
			Expect(set.Size()).To(Equal(4))
		})

		It("contains the elements of the slice", func() {
			Expect(set.Contains("a")).To(BeTrue())
			Expect(set.Contains("b")).To(BeTrue())
			Expect(set.Contains("c")).To(BeTrue())
			Expect(set.Contains("d")).To(BeTrue())
		})

		When("the map is nil", func() {
			BeforeEach(func() {
				var m map[int]string
				set = ds.SetFromMapValues(m)
			})

			It("is empty", func() {
				Expect(set.IsEmpty()).To(BeTrue())
			})
		})
	})

	When("constructed as the union of two sets", func() {
		var (
			firstSet  *ds.Set[string]
			secondSet *ds.Set[string]
		)

		BeforeEach(func() {
			firstSet = ds.NewSet[string](2)
			firstSet.Add("a")
			firstSet.Add("b")

			secondSet = ds.NewSet[string](2)
			secondSet.Add("b")
			secondSet.Add("c")

			set = ds.SetUnion(firstSet, secondSet)
		})

		It("the set has the correct size", func() {
			Expect(set.Size()).To(Equal(3))
		})

		It("contains the items from both sets", func() {
			Expect(set.Contains("a")).To(BeTrue())
			Expect(set.Contains("b")).To(BeTrue())
			Expect(set.Contains("c")).To(BeTrue())
		})
	})

	When("constructed as the difference of two sets", func() {
		var (
			firstSet  *ds.Set[string]
			secondSet *ds.Set[string]
		)

		BeforeEach(func() {
			firstSet = ds.NewSet[string](2)
			firstSet.Add("a")
			firstSet.Add("b")

			secondSet = ds.NewSet[string](2)
			secondSet.Add("b")
			secondSet.Add("c")

			set = ds.SetDifference(firstSet, secondSet)
		})

		It("the set has the correct size", func() {
			Expect(set.Size()).To(Equal(1))
		})

		It("contains the items from both sets", func() {
			Expect(set.Contains("a")).To(BeTrue())
		})
	})

	When("constructed as the intersection of two sets", func() {
		var (
			firstSet  *ds.Set[string]
			secondSet *ds.Set[string]
		)

		BeforeEach(func() {
			firstSet = ds.NewSet[string](3)
			firstSet.Add("a")
			firstSet.Add("b")
			firstSet.Add("c")

			secondSet = ds.NewSet[string](3)
			secondSet.Add("b")
			secondSet.Add("c")
			secondSet.Add("d")

			set = ds.SetIntersection(firstSet, secondSet)
		})

		It("the set has the correct size", func() {
			Expect(set.Size()).To(Equal(2))
		})

		It("contains the items from both sets", func() {
			Expect(set.Contains("b")).To(BeTrue())
			Expect(set.Contains("c")).To(BeTrue())
		})
	})
})
