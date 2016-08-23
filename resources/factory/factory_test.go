package factory_test

import (
	"github.com/RackHD/ipam/resources"
	. "github.com/RackHD/ipam/resources/factory"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Factory", func() {
	Describe("Request", func() {
		It("should return the requested resource", func() {
			resource, err := Request(resources.PoolResourceType, resources.PoolResourceVersionV1)
			Expect(err).ToNot(HaveOccurred())
			Expect(resource).To(BeAssignableToTypeOf(&resources.PoolV1{}))
		})

		It("should return an error if the requested resource is not registered", func() {
			_, err := Request("invalid", "1.0.0")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(HavePrefix("Request"))
		})
	})

	Describe("Require", func() {
		It("should return the requested resource", func() {
			resource, err := Require(resources.PoolResourceType, resources.PoolResourceVersionV1)
			Expect(err).ToNot(HaveOccurred())
			Expect(resource).To(BeAssignableToTypeOf(&resources.PoolV1{}))
		})

		It("should return an error if the requested resource is not registered", func() {
			_, err := Require("invalid", "1.0.0")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(HavePrefix("Request"))
		})

		It("should return an error if the requested version is not registered", func() {
			_, err := Require(resources.PoolResourceType, "invalid")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(HavePrefix("Require"))
		})
	})
})
