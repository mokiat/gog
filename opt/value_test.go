package opt_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mokiat/gog/opt"
)

var _ = Describe("Optional", func() {

	It("is possible to create an unspecified value", func() {
		v := opt.Unspecified[string]()
		Expect(v.Specified).To(BeFalse())
	})

	It("is possible to create a specified value", func() {
		v := opt.V("hello")
		Expect(v.Specified).To(BeTrue())
		Expect(v.Value).To(Equal("hello"))
	})

})
