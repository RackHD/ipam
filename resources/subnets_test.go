package resources_test

import (
	"github.com/RackHD/ipam/models"
	. "github.com/RackHD/ipam/resources"
	"gopkg.in/mgo.v2/bson"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SubnetsCreator", func() {
	It("should return a SubnetsV1 resource by default", func() {
		resource, err := SubnetsCreator("")
		Expect(err).ToNot(HaveOccurred())
		Expect(resource).To(BeAssignableToTypeOf(&SubnetsV1{}))
	})
})

var _ = Describe("SubnetsV1", func() {
	var (
		resource = SubnetsV1{}
		subnets  = []models.Subnet{
			{
				ID:       bson.NewObjectId(),
				Name:     "SubnetV1 Name",
				Tags:     []string{"SubnetV1"},
				Metadata: "SubnetV1 Metadata",
			},
		}
	)

	Describe("Type", func() {
		It("should return the correct resource type", func() {
			Expect(resource.Type()).To(Equal(SubnetsResourceType))
		})
	})

	Describe("Version", func() {
		It("should return the correct resource version", func() {
			Expect(resource.Version()).To(Equal(SubnetsResourceVersionV1))
		})
	})

	Describe("Marshal", func() {
		It("should copy the []models.Subnet to itself", func() {
			err := resource.Marshal(subnets)
			Expect(err).ToNot(HaveOccurred())

			Expect(len(resource.Subnets)).To(Equal(1))

			Expect(resource.Subnets[0].ID).To(Equal(subnets[0].ID.Hex()))
			Expect(resource.Subnets[0].Name).To(Equal(subnets[0].Name))
			Expect(resource.Subnets[0].Tags).To(Equal(subnets[0].Tags))
			Expect(resource.Subnets[0].Metadata).To(Equal(subnets[0].Metadata))
		})

		It("should return an error if a []model.Subnet is not provided", func() {
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
