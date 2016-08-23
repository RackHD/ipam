package factory

import (
	"fmt"

	"github.com/RackHD/ipam/interfaces"
)

// factory is a storage map for resource creator functions registered via init.
var factory = make(map[string]interfaces.ResourceCreator)

// Register associates the resource identifier with a resource creator function.
func Register(resource string, creator interfaces.ResourceCreator) {
	factory[resource] = creator
}

// Request finds a resource creator function by the resource identifier and calls
// the creator returning the result.  The resulting resource may be a default version
// if the reqeusted version is not present.
func Request(resource string, version string) (interfaces.Resource, error) {
	if creator, ok := factory[resource]; ok {
		return creator(version)
	}

	return nil, fmt.Errorf("Request: Unable to locate resource %s.", resource)
}

// Require finds a resource creator function by the resource identifier and verifies
// the created resource matches the requested version.  If not, an error will be
// returned.
func Require(resource string, version string) (interfaces.Resource, error) {
	provided, err := Request(resource, version)
	if err != nil {
		return nil, err
	}

	if provided.Version() != version {
		return nil, fmt.Errorf("Require: Unable to locate resource %s, version %s.", resource, version)
	}

	return provided, nil
}
