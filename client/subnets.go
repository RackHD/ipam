package ipamapi

import (
	"errors"

	"github.com/RackHD/ipam/resources"
)

//Subnets can be used to query the Subnets routes
type Subnets struct {
	client *Client
}

// Index returns a list of Subnets.
func (s *Subnets) Index(poolID string) (resources.SubnetsV1, error) {
	receivedSubnets, err := s.client.ReceiveResource("GET", "/pools/"+poolID+"/subnets", "", "")
	if err != nil {
		return resources.SubnetsV1{}, err
	}
	if subnets, ok := receivedSubnets.(*resources.SubnetsV1); ok {
		return *subnets, nil
	}
	return resources.SubnetsV1{}, errors.New("Subnet Index call error.")
}

// Creates a subnet and return the location.
func (s *Subnets) Creates(poolID string, subnetToCreate resources.SubnetV1) (string, error) {
	subnetLocation, err := s.client.SendResource("POST", "/pools/"+poolID+"/subnets", &subnetToCreate)
	if err != nil {
		return "", err
	}
	return subnetLocation, nil
}

// CreateShowSubnet creates a subnet and then returns that subnet.
func (s *Subnets) CreateShowSubnet(poolID string, subnetToCreate resources.SubnetV1) (resources.SubnetV1, error) {
	receivedSubnet, err := s.client.SendReceiveResource("POST", "GET", "/pools/"+poolID+"/subnets", &subnetToCreate)
	if err != nil {
		return resources.SubnetV1{}, err
	}
	if subnet, ok := receivedSubnet.(*resources.SubnetV1); ok {
		return *subnet, nil
	}
	return resources.SubnetV1{}, errors.New("CreateShowSubnet call error.")
}

// Show returns the requested subnet.
func (s *Subnets) Show(subnetID string, subnetToGet resources.SubnetV1) (resources.SubnetV1, error) {
	receivedSubnet, err := s.client.ReceiveResource("GET", "/subnets/"+subnetID, subnetToGet.Type(), subnetToGet.Version())
	if err != nil {
		return resources.SubnetV1{}, err
	}
	if subnet, ok := receivedSubnet.(*resources.SubnetV1); ok {
		return *subnet, nil
	}
	return resources.SubnetV1{}, errors.New("Subnet Show call error.")
}

// Update updates the requested subnet and returns its location.
func (s *Subnets) Update(subnetID string, subnetToUpdate resources.SubnetV1) (string, error) {
	subnetLocation, err := s.client.SendResource("PATCH", "/subnets/"+subnetID, &subnetToUpdate)
	if err != nil {
		return "", err
	}
	return subnetLocation, nil
}

// UpdateShowSubnet updates a Subnet and then returns that Subnet.
func (s *Subnets) UpdateShowSubnet(subnetID string, subnetToUpdate resources.SubnetV1) (resources.SubnetV1, error) {
	receivedSubnet, err := s.client.SendReceiveResource("PATCH", "GET", "/subnets/"+subnetID, &subnetToUpdate)
	if err != nil {
		return resources.SubnetV1{}, err
	}
	if subnet, ok := receivedSubnet.(*resources.SubnetV1); ok {
		return *subnet, nil
	}
	return resources.SubnetV1{}, errors.New("UpdateShowSubnet call error.")
}

// Delete removed the requested subnet and returns the location.
func (s *Subnets) Delete(subnetID string, subnetToDelete resources.SubnetV1) (string, error) {
	subnetLocation, err := s.client.SendResource("DELETE", "/subnets/"+subnetID, &subnetToDelete)
	if err != nil {
		return "", err
	}
	return subnetLocation, nil
}
