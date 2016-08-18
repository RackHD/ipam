package resources

import (
	"fmt"

	"github.com/RackHD/ipam/interfaces"
	"github.com/RackHD/ipam/models"
	"github.com/RackHD/ipam/resources/factory"
	"gopkg.in/mgo.v2/bson"
)

// SubnetResourceType is the media type assigned to a Subnet resource.
const SubnetResourceType string = "application/vnd.ipam.subnet"

// SubnetResourceVersionV1 is the semantic version identifier for the Pool resource.
const SubnetResourceVersionV1 string = "1.0.0"

func init() {
	factory.Register(SubnetResourceType, SubnetCreator)
}

// SubnetCreator is a factory function for turning a version string into a Subnet resource.
func SubnetCreator(version string) (interfaces.Resource, error) {
	return &SubnetV1{}, nil
}

// SubnetV1 represents the v1.0.0 version of the Subnet resource.
type SubnetV1 struct {
	ID       string      `json:"id"`
	Name     string      `json:"name"`
	Tags     []string    `json:"tags"`
	Metadata interface{} `json:"metadata"`
	Pool     string      `json:"pool"`
}

// Type returns the resource type for use in rendering HTTP response headers.
func (s *SubnetV1) Type() string {
	return SubnetResourceType
}

// Version returns the resource version for use in rendering HTTP response headers.
func (s *SubnetV1) Version() string {
	return SubnetResourceVersionV1
}

// Marshal converts a models.Subnet object into this version of the resource.
func (s *SubnetV1) Marshal(object interface{}) error {
	if target, ok := object.(models.Subnet); ok {
		s.ID = target.ID.Hex()
		s.Name = target.Name
		s.Tags = target.Tags
		s.Metadata = target.Metadata
		s.Pool = target.Pool.Hex()

		return nil
	}

	return fmt.Errorf("Invalid Object Type: %+v", object)
}

// Unmarshal converts the resource into a models.Subnet object.
func (s *SubnetV1) Unmarshal() (interface{}, error) {
	if s.ID == "" {
		s.ID = bson.NewObjectId().Hex()
	}

	return models.Subnet{
		ID:       bson.ObjectIdHex(s.ID),
		Name:     s.Name,
		Tags:     s.Tags,
		Metadata: s.Metadata,
	}, nil
}
