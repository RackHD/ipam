package resources_test

import (
	"github.com/RackHD/ipam/models"
	. "github.com/RackHD/ipam/resources"
	"gopkg.in/mgo.v2/bson"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SubnetCreator", func() {
	It("should return a SubnetV1 resource by default", func() {
		resource, err := SubnetCreator("")
		Expect(err).ToNot(HaveOccurred())
		Expect(resource).To(BeAssignableToTypeOf(&SubnetV1{}))
	})
})

var _ = Describe("SubnetV1", func() {
	var (
		resource = SubnetV1{
			ID:       bson.NewObjectId().Hex(),
			Name:     "SubnetV1 Name",
			Tags:     []string{"SubnetV1"},
			Metadata: "SubnetV1 Metadata",
		}
		model = models.Subnet{
			ID:       bson.NewObjectId(),
			Name:     "Subnet Name",
			Tags:     []string{"Subnet"},
			Metadata: "Subnet Metadata",
		}
	)

	Describe("Type", func() {
		It("should return the correct resource type", func() {
			Expect(resource.Type()).To(Equal(SubnetResourceType))
		})
	})

	Describe("Version", func() {
		It("should return the correct resource version", func() {
			Expect(resource.Version()).To(Equal(SubnetResourceVersionV1))
		})
	})

	Describe("Marshal", func() {
		It("should copy the models.Subnet to itself", func() {
			err := resource.Marshal(model)
			Expect(err).ToNot(HaveOccurred())

			Expect(resource.ID).To(Equal(model.ID.Hex()))
			Expect(resource.Name).To(Equal(model.Name))
			Expect(resource.Tags).To(Equal(model.Tags))
			Expect(resource.Metadata).To(Equal(model.Metadata))
		})

		It("should return an error if a model.Subnet is not provided", func() {
			err := resource.Marshal("invalid")
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("Unmarshal", func() {
		It("should copy itself to a models.Subnet", func() {
			m, err := resource.Unmarshal()
			Expect(err).ToNot(HaveOccurred())
			Expect(m).To(BeAssignableToTypeOf(models.Subnet{}))

			if result, ok := m.(models.Subnet); ok {
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
			Expect(m).To(BeAssignableToTypeOf(models.Subnet{}))

			if result, ok := m.(models.Subnet); ok {
				Expect(result.ID.Hex()).To(Equal(resource.ID))
				Expect(result.Name).To(Equal(resource.Name))
				Expect(result.Tags).To(Equal(resource.Tags))
				Expect(result.Metadata).To(Equal(resource.Metadata))
			}
		})
	})
})
