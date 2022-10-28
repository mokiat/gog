package gog_test

import (
	"github.com/mokiat/gog"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pointer", func() {
	Describe("PtrOf", func() {
		It("returns a pointer to the passed value", func() {
			result := gog.PtrOf("hello")
			Expect(result).ToNot(BeNil())
			Expect(*result).To(Equal("hello"))
		})
	})

	Describe("ValueOf", func() {
		var ptr *string

		BeforeEach(func() {
			value := "test"
			ptr = &value
		})

		It("returns the value behind a pointer", func() {
			Expect(gog.ValueOf(ptr, "")).To(Equal("test"))
		})

		It("returns the default value when the pointer is nil", func() {
			Expect(gog.ValueOf(nil, "default")).To(Equal("default"))
		})
	})
})
