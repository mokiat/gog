package ds_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mokiat/gog/ds"
)

var _ = Describe("Stack", func() {
	var (
		stack *ds.Stack[string]
	)

	BeforeEach(func() {
		stack = ds.NewStack[string](0)
	})

	It("is empty by default", func() {
		Expect(stack.IsEmpty()).To(BeTrue())
	})

	It("has zero size by default", func() {
		Expect(stack.Size()).To(BeZero())
	})

	When("items are added", func() {
		BeforeEach(func() {
			stack.Push("first")
			stack.Push("second")
			stack.Push("third")
		})

		It("is no longer empty", func() {
			Expect(stack.IsEmpty()).To(BeFalse())
		})

		It("has the correct size", func() {
			Expect(stack.Size()).To(Equal(3))
		})

		It("is possible to fetch the items", func() {
			Expect(stack.Pop()).To(Equal("third"))
			Expect(stack.Pop()).To(Equal("second"))
			Expect(stack.Pop()).To(Equal("first"))
		})

		It("is possible to peek the top item", func() {
			Expect(stack.Peek()).To(Equal("third"))
			Expect(stack.Size()).To(Equal(3))
		})

		When("the stack is clipped", func() {
			BeforeEach(func() {
				stack.Clip()
			})

			It("still contains the same items", func() {
				Expect(stack.Pop()).To(Equal("third"))
				Expect(stack.Pop()).To(Equal("second"))
				Expect(stack.Pop()).To(Equal("first"))
			})
		})

		When("the stack is cleared", func() {
			BeforeEach(func() {
				stack.Clear()
			})

			It("becomes empty", func() {
				Expect(stack.IsEmpty()).To(BeTrue())
			})

			It("changes its size to zero", func() {
				Expect(stack.Size()).To(BeZero())
			})
		})

		When("an item is popped", func() {
			BeforeEach(func() {
				stack.Pop()
			})

			It("changes its size accordingly", func() {
				Expect(stack.Size()).To(Equal(2))
			})

			It("no longer contains the top item", func() {
				Expect(stack.Pop()).To(Equal("second"))
				Expect(stack.Pop()).To(Equal("first"))
			})
		})
	})
})
