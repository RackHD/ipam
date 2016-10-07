package resources

import (
	"fmt"

	"github.com/RackHD/ipam/interfaces"
	"github.com/RackHD/ipam/models"
	"github.com/RackHD/ipam/resources/factory"
)

// LeasesResourceType is the media type assigned to a collection of Lease resources.
const LeasesResourceType string = "application/vnd.ipam.leases"

// LeasesResourceVersionV1 is the semantic version identifier for the Pool resource.
const LeasesResourceVersionV1 string = "1.0.0"

func init() {
	factory.Register(LeasesResourceType, LeasesCreator)
}

// LeasesCreator is a factory function for turning a version string into a Leases resource.
func LeasesCreator(version string) (interfaces.Resource, error) {
	return &LeasesV1{}, nil
}

// LeasesV1 represents the v1.0.0 version of the Leases resource.
type LeasesV1 struct {
	Leases []LeaseV1 `json:"leases"`
}

// Type returns the resource type for use in rendering HTTP response headers.
func (p *LeasesV1) Type() string {
	return LeasesResourceType
}

// Version returns the resource version for use in rendering HTTP response headers.
func (p *LeasesV1) Version() string {
	return LeasesResourceVersionV1
}

// Marshal converts an array of models.Lease objects into this version of the resource.
func (p *LeasesV1) Marshal(object interface{}) error {
	if subnets, ok := object.([]models.Lease); ok {
		p.Leases = make([]LeaseV1, len(subnets))

		for i := range p.Leases {
			p.Leases[i].Marshal(subnets[i])
		}

		return nil
	}

	return fmt.Errorf("Invalid Object Type.")
}

// Unmarshal converts the resource into an array of models.Lease objects.
func (p *LeasesV1) Unmarshal() (interface{}, error) {
	return nil, fmt.Errorf("Invalid Action for Resource.")
}
