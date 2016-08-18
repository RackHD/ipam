package resources

import (
	"fmt"

	"github.com/RackHD/ipam/interfaces"
	"github.com/RackHD/ipam/models"
	"github.com/RackHD/ipam/resources/factory"
)

// SubnetsResourceType is the media type assigned to a collection of Subnet resources.
const SubnetsResourceType string = "application/vnd.ipam.subnets"

// SubnetsResourceVersionV1 is the semantic version identifier for the Pool resource.
const SubnetsResourceVersionV1 string = "1.0.0"

func init() {
	factory.Register(SubnetsResourceType, SubnetsCreator)
}

// SubnetsCreator is a factory function for turning a version string into a Subnets resource.
func SubnetsCreator(version string) (interfaces.Resource, error) {
	return &SubnetsV1{}, nil
}

// SubnetsV1 represents the v1.0.0 version of the Subnets resource.
type SubnetsV1 struct {
	Subnets []SubnetV1 `json:"subnets"`
}

// Type returns the resource type for use in rendering HTTP response headers.
func (p *SubnetsV1) Type() string {
	return SubnetsResourceType
}

// Version returns the resource version for use in rendering HTTP response headers.
func (p *SubnetsV1) Version() string {
	return SubnetsResourceVersionV1
}

// Marshal converts an array of models.Subnet objects into this version of the resource.
func (p *SubnetsV1) Marshal(object interface{}) error {
	if subnets, ok := object.([]models.Subnet); ok {
		p.Subnets = make([]SubnetV1, len(subnets))

		for i := range p.Subnets {
			p.Subnets[i].Marshal(subnets[i])
		}

		return nil
	}

	return fmt.Errorf("Invalid Object Type.")
}

// Unmarshal converts the resource into an array of models.Subnet objects.
func (p *SubnetsV1) Unmarshal() (interface{}, error) {
	return nil, fmt.Errorf("Invalid Action for Resource.")
}
