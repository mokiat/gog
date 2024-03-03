package gog_test

import (
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

})
