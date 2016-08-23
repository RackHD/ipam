package resources_test

import (
	"github.com/RackHD/ipam/models"
	. "github.com/RackHD/ipam/resources"
	"gopkg.in/mgo.v2/bson"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PoolCreator", func() {
	It("should return a PoolV1 resource by default", func() {
		resource, err := PoolCreator("")
		Expect(err).ToNot(HaveOccurred())
		Expect(resource).To(BeAssignableToTypeOf(&PoolV1{}))
	})
})

var _ = Describe("PoolV1", func() {
	var (
		resource = PoolV1{
			ID:       bson.NewObjectId().Hex(),
			Name:     "PoolV1 Name",
			Tags:     []string{"PoolV1"},
			Metadata: "PoolV1 Metadata",
		}
		model = models.Pool{
			ID:       bson.NewObjectId(),
			Name:     "Pool Name",
			Tags:     []string{"Pool"},
			Metadata: "Pool Metadata",
		}
	)

	Describe("Type", func() {
		It("should return the correct resource type", func() {
			Expect(resource.Type()).To(Equal(PoolResourceType))
		})
	})

	Describe("Version", func() {
		It("should return the correct resource version", func() {
			Expect(resource.Version()).To(Equal(PoolResourceVersionV1))
		})
	})

	Describe("Marshal", func() {
		It("should copy the models.Pool to itself", func() {
			err := resource.Marshal(model)
			Expect(err).ToNot(HaveOccurred())

			Expect(resource.ID).To(Equal(model.ID.Hex()))
			Expect(resource.Name).To(Equal(model.Name))
			Expect(resource.Tags).To(Equal(model.Tags))
			Expect(resource.Metadata).To(Equal(model.Metadata))
		})

		It("should return an error if a model.Pool is not provided", func() {
			err := resource.Marshal("invalid")
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("Unmarshal", func() {
		It("should copy itself to a models.Pool", func() {
			m, err := resource.Unmarshal()
			Expect(err).ToNot(HaveOccurred())
			Expect(m).To(BeAssignableToTypeOf(models.Pool{}))

			if result, ok := m.(models.Pool); ok {
				Expect(result.ID.Hex()).To(Equal(resource.ID))
				Expect(result.Name).To(Equal(resource.Name))
				Expect(result.Tags).To(Equal(resource.Tags))
				Expect(result.Metadata).To(Equal(resource.Metadata))
			}
		})

		It("should generate a new object ID if one is not present", func() {
			resource.ID = ""

			m, err := resource.Unmarshal()
			Expect(err).ToNot(HaveOccurred())
			Expect(m).To(BeAssignableToTypeOf(models.Pool{}))

			if result, ok := m.(models.Pool); ok {
				Expect(result.ID.Hex()).To(Equal(resource.ID))
				Expect(result.Name).To(Equal(resource.Name))
				Expect(result.Tags).To(Equal(resource.Tags))
				Expect(result.Metadata).To(Equal(resource.Metadata))
			}
		})
	})
})
