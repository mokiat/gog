package ds_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mokiat/gog/ds"
)

var _ = Describe("Pool", func() {
	type Item struct {
		Value string
	}

	var pool *ds.Pool[Item]

	BeforeEach(func() {
		pool = ds.NewPool[Item]()
	})

	It("is initially empty", func() {
		Expect(pool.IsEmpty()).To(BeTrue())
	})

	It("is possible to fetch an item", func() {
		item := pool.Fetch()
		Expect(item).ToNot(BeNil())
	})

	When("an item is fetched", func() {
		var item *Item

		BeforeEach(func() {
			item = pool.Fetch()
			item.Value = "Changed"
		})

		It("is possible to return the item", func() {
			pool.Restore(item)
		})

		When("an item is restored", func() {
			BeforeEach(func() {
				pool.Restore(item)
			})

			It("is no longer empty", func() {
				Expect(pool.IsEmpty()).To(BeFalse())
			})

			It("is possible to fetch a cached item", func() {
				newItem := pool.Fetch()
				Expect(newItem).ToNot(BeNil())
				Expect(newItem.Value).To(Equal("Changed"))
			})

			When("the pool is cleared", func() {
				BeforeEach(func() {
					pool.Clear()
				})

				It("is empty", func() {
					Expect(pool.IsEmpty()).To(BeTrue())
				})

				It("is not possible to fetch a cached item", func() {
					newItem := pool.Fetch()
					Expect(newItem).ToNot(BeNil())
					Expect(newItem.Value).To(BeEmpty())
				})
			})
		})
	})
})
