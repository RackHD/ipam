package resources

import (
	"fmt"

	"github.com/RackHD/ipam/interfaces"
	"github.com/RackHD/ipam/models"
	"github.com/RackHD/ipam/resources/factory"
)

// ReservationsResourceType is the media type assigned to a collection of Reservation resources.
const ReservationsResourceType string = "application/vnd.ipam.reservations"

// ReservationsResourceVersionV1 is the semantic version identifier for the Pool resource.
const ReservationsResourceVersionV1 string = "1.0.0"

func init() {
	factory.Register(ReservationsResourceType, ReservationsCreator)
}

// ReservationsCreator is a factory function for turning a version string into a Reservations resource.
func ReservationsCreator(version string) (interfaces.Resource, error) {
	return &ReservationsV1{}, nil
}

// ReservationsV1 represents the v1.0.0 version of the Reservations resource.
type ReservationsV1 struct {
	Reservations []ReservationV1 `json:"reservations"`
}

// Type returns the resource type for use in rendering HTTP response headers.
func (p *ReservationsV1) Type() string {
	return ReservationsResourceType
}

// Version returns the resource version for use in rendering HTTP response headers.
func (p *ReservationsV1) Version() string {
	return ReservationsResourceVersionV1
}

// Marshal converts an array of models.Reservation objects into this version of the resource.
func (p *ReservationsV1) Marshal(object interface{}) error {
	if reservations, ok := object.([]models.Reservation); ok {
		p.Reservations = make([]ReservationV1, len(reservations))

		for i := range p.Reservations {
			p.Reservations[i].Marshal(reservations[i])
		}

		return nil
	}

	return fmt.Errorf("Invalid Object Type.")
}

// Unmarshal converts the resource into an array of models.Reservation objects.
func (p *ReservationsV1) Unmarshal() (interface{}, error) {
	return nil, fmt.Errorf("Invalid Action for Resource.")
}
