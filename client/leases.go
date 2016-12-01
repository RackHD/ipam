package client

import (
	"errors"

	"github.com/RackHD/ipam/resources"
)

// IndexLeases returns a list of Leases.
func (c *Client) IndexLeases(reservationID string) (resources.LeasesV1, error) {
	returnedLeases, err := c.ReceiveResource("GET", "/reservations/"+reservationID+"/leases", "", "")
	if err != nil {
		return resources.LeasesV1{}, err
	}
	if leases, ok := returnedLeases.(*resources.LeasesV1); ok {
		return *leases, nil
	}
	return resources.LeasesV1{}, errors.New("Lease Index call error.")
}

// ShowLease returns the requested Lease.
func (c *Client) ShowLease(leaseID string, leaseToShow resources.LeaseV1) (resources.LeaseV1, error) {
	returnedLease, err := c.ReceiveResource("GET", "/leases/"+leaseID, leaseToShow.Type(), leaseToShow.Version())
	if err != nil {
		return resources.LeaseV1{}, err
	}
	if lease, ok := returnedLease.(*resources.LeaseV1); ok {
		return *lease, nil
	}
	return resources.LeaseV1{}, errors.New("Lease Show call error.")
}

// UpdateLease updates the requested Lease and returns its location.
func (c *Client) UpdateLease(leaseID string, leaseToUpdate resources.LeaseV1) (string, error) {
	leaseLocation, err := c.SendResource("PATCH", "/leases/"+leaseID, &leaseToUpdate)
	if err != nil {
		return "", err
	}
	return leaseLocation, nil
}

// UpdateShowLease updates a Lease and then returns that Lease.
func (c *Client) UpdateShowLease(leaseID string, leaseToUpdate resources.LeaseV1) (resources.LeaseV1, error) {
	returnedLease, err := c.SendReceiveResource("PATCH", "GET", "/leases/"+leaseID, &leaseToUpdate)
	if err != nil {
		return resources.LeaseV1{}, err
	}
	if lease, ok := returnedLease.(*resources.LeaseV1); ok {
		return *lease, nil
	}
	return resources.LeaseV1{}, errors.New("UpdateShowLease call error.")
}
