package ipamapi

import (
	"errors"

	"github.com/RackHD/ipam/resources"
)

// IndexSubnets returns a list of Subnets.
func (c *Client) IndexSubnets(poolID string) (resources.SubnetsV1, error) {
	receivedSubnets, err := c.ReceiveResource("GET", "/pools/"+poolID+"/subnets", "", "")
	if err != nil {
		return resources.SubnetsV1{}, err
	}
	if subnets, ok := receivedSubnets.(*resources.SubnetsV1); ok {
		return *subnets, nil
	}
	return resources.SubnetsV1{}, errors.New("Subnet Index call error.")
}

// CreateSubnet a subnet and return the location.
func (c *Client) CreateSubnet(poolID string, subnetToCreate resources.SubnetV1) (string, error) {
	subnetLocation, err := c.SendResource("POST", "/pools/"+poolID+"/subnets", &subnetToCreate)
	if err != nil {
		return "", err
	}
	return subnetLocation, nil
}

// CreateShowSubnet creates a subnet and then returns that subnet.
func (c *Client) CreateShowSubnet(poolID string, subnetToCreate resources.SubnetV1) (resources.SubnetV1, error) {
	receivedSubnet, err := c.SendReceiveResource("POST", "GET", "/pools/"+poolID+"/subnets", &subnetToCreate)
	if err != nil {
		return resources.SubnetV1{}, err
	}
	if subnet, ok := receivedSubnet.(*resources.SubnetV1); ok {
		return *subnet, nil
	}
	return resources.SubnetV1{}, errors.New("CreateShowSubnet call error.")
}

// ShowSubnet returns the requested subnet.
func (c *Client) ShowSubnet(subnetID string, subnetToGet resources.SubnetV1) (resources.SubnetV1, error) {
	receivedSubnet, err := c.ReceiveResource("GET", "/subnets/"+subnetID, subnetToGet.Type(), subnetToGet.Version())
	if err != nil {
		return resources.SubnetV1{}, err
	}
	if subnet, ok := receivedSubnet.(*resources.SubnetV1); ok {
		return *subnet, nil
	}
	return resources.SubnetV1{}, errors.New("Subnet Show call error.")
}

// UpdateSubnet updates the requested subnet and returns its location.
func (c *Client) UpdateSubnet(subnetID string, subnetToUpdate resources.SubnetV1) (string, error) {
	subnetLocation, err := c.SendResource("PATCH", "/subnets/"+subnetID, &subnetToUpdate)
	if err != nil {
		return "", err
	}
	return subnetLocation, nil
}

// UpdateShowSubnet updates a Subnet and then returns that Subnet.
func (c *Client) UpdateShowSubnet(subnetID string, subnetToUpdate resources.SubnetV1) (resources.SubnetV1, error) {
	receivedSubnet, err := c.SendReceiveResource("PATCH", "GET", "/subnets/"+subnetID, &subnetToUpdate)
	if err != nil {
		return resources.SubnetV1{}, err
	}
	if subnet, ok := receivedSubnet.(*resources.SubnetV1); ok {
		return *subnet, nil
	}
	return resources.SubnetV1{}, errors.New("UpdateShowSubnet call error.")
}

// DeleteSubnet removed the requested subnet and returns the location.
func (c *Client) DeleteSubnet(subnetID string, subnetToDelete resources.SubnetV1) (string, error) {
	subnetLocation, err := c.SendResource("DELETE", "/subnets/"+subnetID, &subnetToDelete)
	if err != nil {
		return "", err
	}
	return subnetLocation, nil
}
