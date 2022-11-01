package ds_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mokiat/gog/ds"
)

var _ = Describe("Heap", func() {
	var heap *ds.Heap[int]

	BeforeEach(func() {
		smallerInt := func(a, b int) bool {
			return a < b
		}
		heap = ds.NewHeap(0, smallerInt)
	})

	It("is empty by default", func() {
		Expect(heap.IsEmpty()).To(BeTrue())
	})

	It("has zero size by default", func() {
		Expect(heap.Size()).To(BeZero())
	})

	When("items are added", func() {
		BeforeEach(func() {
			heap.Push(21)
			heap.Push(10)
			heap.Push(2)
			heap.Push(15)
			heap.Push(6)
			heap.Push(15)
		})

		It("is no longer empty", func() {
			Expect(heap.IsEmpty()).To(BeFalse())
		})

		It("has the correct size", func() {
			Expect(heap.Size()).To(Equal(6))
		})

		It("is possible to peek best item", func() {
			Expect(heap.Peek()).To(Equal(2))
		})

		It("is possible to fetch all items in order", func() {
			Expect(heap.Pop()).To(Equal(2))
			Expect(heap.Pop()).To(Equal(6))
			Expect(heap.Pop()).To(Equal(10))
			Expect(heap.Pop()).To(Equal(15))
			Expect(heap.Pop()).To(Equal(15))
			Expect(heap.Pop()).To(Equal(21))
			Expect(heap.IsEmpty()).To(BeTrue())
		})

		When("clipped", func() {
			BeforeEach(func() {
				heap.Clip()
			})

			It("still contains the same items", func() {
				Expect(heap.Pop()).To(Equal(2))
				Expect(heap.Pop()).To(Equal(6))
				Expect(heap.Pop()).To(Equal(10))
				Expect(heap.Pop()).To(Equal(15))
				Expect(heap.Pop()).To(Equal(15))
				Expect(heap.Pop()).To(Equal(21))
			})
		})

		When("cleared", func() {
			BeforeEach(func() {
				heap.Clear()
			})

			It("becomes empty", func() {
				Expect(heap.IsEmpty()).To(BeTrue())
			})

			It("changes its size to zero", func() {
				Expect(heap.Size()).To(BeZero())
			})
		})

		When("some items are removed and more are added", func() {
			BeforeEach(func() {
				heap.Pop() // 2
				heap.Pop() // 6
				heap.Pop() // 10
				heap.Pop() // 15
				Expect(heap.IsEmpty()).To(BeFalse())

				heap.Push(1)
				heap.Push(31)
				heap.Push(1)
				heap.Push(12)
				heap.Push(12)
			})

			It("changes its size accordingly", func() {
				Expect(heap.Size()).To(Equal(7))
			})

			It("is possible to fetch new and remaining items in order", func() {
				Expect(heap.Pop()).To(Equal(1))
				Expect(heap.Pop()).To(Equal(1))
				Expect(heap.Pop()).To(Equal(12))
				Expect(heap.Pop()).To(Equal(12))
				Expect(heap.Pop()).To(Equal(15))
				Expect(heap.Pop()).To(Equal(21))
				Expect(heap.Pop()).To(Equal(31))
				Expect(heap.IsEmpty()).To(BeTrue())
			})
		})
	})
})
