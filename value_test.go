package gog_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mokiat/gog"
)

var _ = Describe("Value", func() {

	Describe("Zero", func() {
		It("returns the zero value", func() {
			Expect(gog.Zero[int]()).To(Equal(0))
			Expect(gog.Zero[string]()).To(Equal(""))
			Expect(gog.Zero[bool]()).To(Equal(false))
		})
	})

	Describe("Must", func() {
		var fn func() (string, error)

		When("no error is returned", func() {
			BeforeEach(func() {
				fn = func() (string, error) {
					return "value", nil
				}
			})

			It("returns the value", func() {
				Expect(gog.Must(fn())).To(Equal("value"))
			})
		})

		When("an error is returned", func() {
			BeforeEach(func() {
				fn = func() (string, error) {
					return "", errors.New("stubbed to fail")
				}
			})

			It("panics", func() {
				Expect(func() { gog.Must(fn()) }).To(Panic())
			})
		})
	})

	Describe("Ternary", func() {
		It("returns the true value when the condition is true", func() {
			Expect(gog.Ternary(true, "true", "false")).To(Equal("true"))
		})

		It("returns the false value when the condition is false", func() {
			Expect(gog.Ternary(false, "true", "false")).To(Equal("false"))
		})
	})
})
