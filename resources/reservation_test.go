package resources_test

import (
	"github.com/RackHD/ipam/models"
	. "github.com/RackHD/ipam/resources"
	"gopkg.in/mgo.v2/bson"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ReservationCreator", func() {
	It("should return a ReservationV1 resource by default", func() {
		resource, err := ReservationCreator("")
		Expect(err).ToNot(HaveOccurred())
		Expect(resource).To(BeAssignableToTypeOf(&ReservationV1{}))
	})
})

var _ = Describe("ReservationV1", func() {
	var (
		resource = ReservationV1{
			ID:       bson.NewObjectId().Hex(),
			Name:     "ReservationV1 Name",
			Tags:     []string{"ReservationV1"},
			Metadata: "ReservationV1 Metadata",
		}
		model = models.Reservation{
			ID:       bson.NewObjectId(),
			Name:     "Reservation Name",
			Tags:     []string{"Reservation"},
			Metadata: "Reservation Metadata",
		}
	)

	Describe("Type", func() {
		It("should return the correct resource type", func() {
			Expect(resource.Type()).To(Equal(ReservationResourceType))
		})
	})

	Describe("Version", func() {
		It("should return the correct resource version", func() {
			Expect(resource.Version()).To(Equal(ReservationResourceVersionV1))
		})
	})

	Describe("Marshal", func() {
		It("should copy the models.Reservation to itself", func() {
			err := resource.Marshal(model)
			Expect(err).ToNot(HaveOccurred())

			Expect(resource.ID).To(Equal(model.ID.Hex()))
			Expect(resource.Name).To(Equal(model.Name))
			Expect(resource.Tags).To(Equal(model.Tags))
			Expect(resource.Metadata).To(Equal(model.Metadata))
		})

		It("should return an error if a model.Reservation is not provided", func() {
			err := resource.Marshal("invalid")
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("Unmarshal", func() {
		It("should copy itself to a models.Reservation", func() {
			m, err := resource.Unmarshal()
			Expect(err).ToNot(HaveOccurred())
			Expect(m).To(BeAssignableToTypeOf(models.Reservation{}))

			if result, ok := m.(models.Reservation); ok {
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
			Expect(m).To(BeAssignableToTypeOf(models.Reservation{}))

			if result, ok := m.(models.Reservation); ok {
				Expect(result.ID.Hex()).To(Equal(resource.ID))
				Expect(result.Name).To(Equal(resource.Name))
				Expect(result.Tags).To(Equal(resource.Tags))
				Expect(result.Metadata).To(Equal(resource.Metadata))
			}
		})
	})
})
