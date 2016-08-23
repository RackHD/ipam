package resources

import (
	"fmt"

	"github.com/RackHD/ipam/interfaces"
	"github.com/RackHD/ipam/models"
	"github.com/RackHD/ipam/resources/factory"
	"gopkg.in/mgo.v2/bson"
)

// ReservationResourceType is the media type assigned to a Reservation resource.
const ReservationResourceType string = "application/vnd.ipam.reservation"

// ReservationResourceVersionV1 is the semantic version identifier for the Subnet resource.
const ReservationResourceVersionV1 string = "1.0.0"

func init() {
	factory.Register(ReservationResourceType, ReservationCreator)
}

// ReservationCreator is a factory function for turning a version string into a Reservation resource.
func ReservationCreator(version string) (interfaces.Resource, error) {
	return &ReservationV1{}, nil
}

// ReservationV1 represents the v1.0.0 version of the Reservation resource.
type ReservationV1 struct {
	ID       string      `json:"id"`
	Name     string      `json:"name"`
	Tags     []string    `json:"tags"`
	Metadata interface{} `json:"metadata"`
	Subnet   string      `json:"subnet"`
}

// Type returns the resource type for use in rendering HTTP response headers.
func (s *ReservationV1) Type() string {
	return ReservationResourceType
}

// Version returns the resource version for use in rendering HTTP response headers.
func (s *ReservationV1) Version() string {
	return ReservationResourceVersionV1
}

// Marshal converts a models.Reservation object into this version of the resource.
func (s *ReservationV1) Marshal(object interface{}) error {
	if target, ok := object.(models.Reservation); ok {
		s.ID = target.ID.Hex()
		s.Name = target.Name
		s.Tags = target.Tags
		s.Metadata = target.Metadata
		s.Subnet = target.Subnet.Hex()

		return nil
	}

	return fmt.Errorf("Invalid Object Type: %+v", object)
}

// Unmarshal converts the resource into a models.Reservation object.
func (s *ReservationV1) Unmarshal() (interface{}, error) {
	if s.ID == "" {
		s.ID = bson.NewObjectId().Hex()
	}

	return models.Reservation{
		ID:       bson.ObjectIdHex(s.ID),
		Name:     s.Name,
		Tags:     s.Tags,
		Metadata: s.Metadata,
	}, nil
}
