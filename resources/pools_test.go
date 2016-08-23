package resources_test

import (
	"github.com/RackHD/ipam/models"
	. "github.com/RackHD/ipam/resources"
	"gopkg.in/mgo.v2/bson"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PoolsCreator", func() {
	It("should return a PoolsV1 resource by default", func() {
		resource, err := PoolsCreator("")
		Expect(err).ToNot(HaveOccurred())
		Expect(resource).To(BeAssignableToTypeOf(&PoolsV1{}))
	})
})

var _ = Describe("PoolsV1", func() {
	var (
		resource = PoolsV1{}
		pools    = []models.Pool{
			{
				ID:       bson.NewObjectId(),
				Name:     "PoolV1 Name",
				Tags:     []string{"PoolV1"},
				Metadata: "PoolV1 Metadata",
			},
		}
	)

	Describe("Type", func() {
		It("should return the correct resource type", func() {
			Expect(resource.Type()).To(Equal(PoolsResourceType))
		})
	})

	Describe("Version", func() {
		It("should return the correct resource version", func() {
			Expect(resource.Version()).To(Equal(PoolsResourceVersionV1))
		})
	})

	Describe("Marshal", func() {
		It("should copy the []models.Pool to itself", func() {
			err := resource.Marshal(pools)
			Expect(err).ToNot(HaveOccurred())

			Expect(len(resource.Pools)).To(Equal(1))

			Expect(resource.Pools[0].ID).To(Equal(pools[0].ID.Hex()))
			Expect(resource.Pools[0].Name).To(Equal(pools[0].Name))
			Expect(resource.Pools[0].Tags).To(Equal(pools[0].Tags))
			Expect(resource.Pools[0].Metadata).To(Equal(pools[0].Metadata))
		})

		It("should return an error if a []model.Pool is not provided", func() {
			err := resource.Marshal("invalid")
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("Unmarshal", func() {
		It("should return an error because the operation is not supported", func() {
			_, err := resource.Unmarshal()
			Expect(err).To(HaveOccurred())
		})
	})
})
