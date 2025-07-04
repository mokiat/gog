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

	It("is possible to unwrap an optional", func() {
		v := opt.V("hello")
		str, ok := v.Unwrap()
		Expect(ok).To(BeTrue())
		Expect(str).To(Equal("hello"))
	})

	It("is possible to wrap to an optional", func() {
		fn := func() (string, bool) {
			return "hello", true
		}
		v := opt.Wrap(fn())
		Expect(v.Value).To(Equal("hello"))
		Expect(v.Specified).To(BeTrue())
	})

	It("is possible to create an unspecified value from pointer", func() {
		v := opt.FromPtr[int](nil)
		Expect(v.Specified).To(BeFalse())
	})

	It("is possible to create a specified value from pointer", func() {
		actual := 10
		v := opt.FromPtr(&actual)
		Expect(v.Specified).To(BeTrue())
		Expect(v.Value).To(Equal(actual))
	})

	It("is possible to get a fallback value", func() {
		v := opt.Unspecified[string]()
		Expect(v.ValueOrDefault("fallback")).To(Equal("fallback"))
		v = opt.V("hello")
		Expect(v.ValueOrDefault("fallback")).To(Equal("hello"))
	})

	It("is possible to get a pointer representation of an unspecified value", func() {
		v := opt.Unspecified[string]()
		ptr := v.ToPtr()
		Expect(ptr).To(BeNil())
	})

	It("is possible to get a pointer representation of a specified value", func() {
		v := opt.V("hello")
		ptr := v.ToPtr()
		Expect(ptr).ToNot(BeNil())
		Expect(*ptr).To(Equal("hello"))
	})

})
