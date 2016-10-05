package resources

import (
	"fmt"
	"net"

	"github.com/RackHD/ipam/interfaces"
	"github.com/RackHD/ipam/models"
	"github.com/RackHD/ipam/resources/factory"
	"gopkg.in/mgo.v2/bson"
)

// LeaseResourceType is the media type assigned to a Lease resource.
const LeaseResourceType string = "application/vnd.ipam.lease"

// LeaseResourceVersionV1 is the semantic version identifier for the Subnet resource.
const LeaseResourceVersionV1 string = "1.0.0"

func init() {
	factory.Register(LeaseResourceType, LeaseCreator)
}

// LeaseCreator is a factory function for turning a version string into a Lease resource.
func LeaseCreator(version string) (interfaces.Resource, error) {
	return &LeaseV1{}, nil
}

// LeaseV1 represents the v1.0.0 version of the Lease resource.
type LeaseV1 struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Tags        []string    `json:"tags"`
	Metadata    interface{} `json:"metadata"`
	Subnet      string      `json:"subnet"`
	Reservation string      `json:"reservation"`
	Address     string      `json:"address"`
}

// Type returns the resource type for use in rendering HTTP response headers.
func (s *LeaseV1) Type() string {
	return LeaseResourceType
}

// Version returns the resource version for use in rendering HTTP response headers.
func (s *LeaseV1) Version() string {
	return LeaseResourceVersionV1
}

// Marshal converts a models.Lease object into this version of the resource.
func (s *LeaseV1) Marshal(object interface{}) error {
	if target, ok := object.(models.Lease); ok {
		s.ID = target.ID.Hex()
		s.Name = target.Name
		s.Tags = target.Tags
		s.Metadata = target.Metadata
		s.Subnet = target.Subnet.Hex()
		s.Reservation = target.Reservation.Hex()
		s.Address = net.IP(target.Address.Data).String()

		return nil
	}

	return fmt.Errorf("Invalid Object Type: %+v", object)
}

// Unmarshal converts the resource into a models.Lease object.
func (s *LeaseV1) Unmarshal() (interface{}, error) {
	if s.ID == "" {
		s.ID = bson.NewObjectId().Hex()
	}

	return models.Lease{
		ID:       bson.ObjectIdHex(s.ID),
		Name:     s.Name,
		Tags:     s.Tags,
		Metadata: s.Metadata,
	}, nil
}
