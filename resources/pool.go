package resources

import (
	"fmt"

	"github.com/RackHD/ipam/interfaces"
	"github.com/RackHD/ipam/models"
	"github.com/RackHD/ipam/resources/factory"
	"gopkg.in/mgo.v2/bson"
)

// PoolResourceType is the media type assigned to a Pool resource.
const PoolResourceType string = "application/vnd.ipam.pool"

// PoolResourceVersionV1 is the semantic version identifier for the Pool resource.
const PoolResourceVersionV1 string = "1.0.0"

func init() {
	factory.Register(PoolResourceType, PoolCreator)
}

// PoolCreator is a factory function for turning a version string into a Pool resource.
func PoolCreator(version string) (interfaces.Resource, error) {
	return &PoolV1{}, nil
}

// PoolV1 represents the v1.0.0 version of the Pool resource.
type PoolV1 struct {
	ID       string      `json:"id"`
	Name     string      `json:"name"`
	Tags     []string    `json:"tags"`
	Metadata interface{} `json:"metadata"`
}

// Type returns the resource type for use in rendering HTTP response headers.
func (p *PoolV1) Type() string {
	return PoolResourceType
}

// Version returns the resource version for use in rendering HTTP response headers.
func (p *PoolV1) Version() string {
	return PoolResourceVersionV1
}

// Marshal converts a models.Pool object into this version of the resource.
func (p *PoolV1) Marshal(object interface{}) error {
	if target, ok := object.(models.Pool); ok {
		p.ID = target.ID.Hex()
		p.Name = target.Name
		p.Tags = target.Tags
		p.Metadata = target.Metadata

		return nil
	}

	return fmt.Errorf("Invalid Object Type: %+v", object)
}

// Unmarshal converts the resource into a models.Pool object.
func (p *PoolV1) Unmarshal() (interface{}, error) {
	if p.ID == "" {
		p.ID = bson.NewObjectId().Hex()
	}

	return models.Pool{
		ID:       bson.ObjectIdHex(p.ID),
		Name:     p.Name,
		Tags:     p.Tags,
		Metadata: p.Metadata,
	}, nil
}
