package resources

import (
	"fmt"

	"github.com/RackHD/ipam/interfaces"
	"github.com/RackHD/ipam/models"
	"github.com/RackHD/ipam/resources/factory"
)

// PoolsResourceType is the media type assigned to a collection of Pool resources.
const PoolsResourceType string = "application/vnd.ipam.pools"

// PoolsResourceVersionV1 is the semantic version identifier for the Pool resource.
const PoolsResourceVersionV1 string = "1.0.0"

func init() {
	factory.Register(PoolsResourceType, PoolsCreator)
}

// PoolsCreator is a factory function for turning a version string into a Pools resource.
func PoolsCreator(version string) (interfaces.Resource, error) {
	return &PoolsV1{}, nil
}

// PoolsV1 represents the v1.0.0 version of the Pools resource.
type PoolsV1 struct {
	Pools []PoolV1 `json:"pools"`
}

// Type returns the resource type for use in rendering HTTP response headers.
func (p *PoolsV1) Type() string {
	return PoolsResourceType
}

// Version returns the resource version for use in rendering HTTP response headers.
func (p *PoolsV1) Version() string {
	return PoolsResourceVersionV1
}

// Marshal converts an array of models.Pool objects into this version of the resource.
func (p *PoolsV1) Marshal(object interface{}) error {
	if pools, ok := object.([]models.Pool); ok {
		p.Pools = make([]PoolV1, len(pools))

		for i := range p.Pools {
			p.Pools[i].Marshal(pools[i])
		}

		return nil
	}

	return fmt.Errorf("Invalid Object Type.")
}

// Unmarshal converts the resource into an array of models.Pool objects.
func (p *PoolsV1) Unmarshal() (interface{}, error) {
	return nil, fmt.Errorf("Invalid Action for Resource.")
}
