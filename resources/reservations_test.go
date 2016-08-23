package resources_test

import (
	"github.com/RackHD/ipam/models"
	. "github.com/RackHD/ipam/resources"
	"gopkg.in/mgo.v2/bson"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ReservationsCreator", func() {
	It("should return a ReservationsV1 resource by default", func() {
		resource, err := ReservationsCreator("")
		Expect(err).ToNot(HaveOccurred())
		Expect(resource).To(BeAssignableToTypeOf(&ReservationsV1{}))
	})
})

var _ = Describe("ReservationsV1", func() {
	var (
		resource     = ReservationsV1{}
		reservations = []models.Reservation{
			{
				ID:       bson.NewObjectId(),
				Name:     "ReservationV1 Name",
				Tags:     []string{"ReservationV1"},
				Metadata: "ReservationV1 Metadata",
			},
		}
	)

	Describe("Type", func() {
		It("should return the correct resource type", func() {
			Expect(resource.Type()).To(Equal(ReservationsResourceType))
		})
	})

	Describe("Version", func() {
		It("should return the correct resource version", func() {
			Expect(resource.Version()).To(Equal(ReservationsResourceVersionV1))
		})
	})

	Describe("Marshal", func() {
		It("should copy the []models.Reservation to itself", func() {
			err := resource.Marshal(reservations)
			Expect(err).ToNot(HaveOccurred())

			Expect(len(resource.Reservations)).To(Equal(1))

			Expect(resource.Reservations[0].ID).To(Equal(reservations[0].ID.Hex()))
			Expect(resource.Reservations[0].Name).To(Equal(reservations[0].Name))
			Expect(resource.Reservations[0].Tags).To(Equal(reservations[0].Tags))
			Expect(resource.Reservations[0].Metadata).To(Equal(reservations[0].Metadata))
		})

		It("should return an error if a []model.Reservation is not provided", func() {
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
