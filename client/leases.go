package ipamapi

import (
	"errors"

	"github.com/RackHD/ipam/resources"
)

//Leases can be used to query the Leases routes
type Leases struct {
	client *Client
}

// Index returns a list of Leases.
func (l *Leases) Index(reservationID string) (resources.LeasesV1, error) {
	returnedLeases, err := l.client.ReceiveResource("GET", "/reservations/"+reservationID+"/leases", "", "")
	if err != nil {
		return resources.LeasesV1{}, err
	}
	if leases, ok := returnedLeases.(*resources.LeasesV1); ok {
		return *leases, nil
	}
	return resources.LeasesV1{}, errors.New("Lease Index call error.")
}

// Show returns the requested Lease.
func (l *Leases) Show(leaseID string, leaseToShow resources.LeaseV1) (resources.LeaseV1, error) {
	returnedLease, err := l.client.ReceiveResource("GET", "/leases/"+leaseID, leaseToShow.Type(), leaseToShow.Version())
	if err != nil {
		return resources.LeaseV1{}, err
	}
	if lease, ok := returnedLease.(*resources.LeaseV1); ok {
		return *lease, nil
	}
	return resources.LeaseV1{}, errors.New("Lease Show call error.")
}

// Update updates the requested Lease and returns its location.
func (l *Leases) Update(leaseID string, leaseToUpdate resources.LeaseV1) (string, error) {
	leaseLocation, err := l.client.SendResource("PATCH", "/leases/"+leaseID, &leaseToUpdate)
	if err != nil {
		return "", err
	}
	return leaseLocation, nil
}

// UpdateShowLease updates a Lease and then returns that Lease.
func (l *Leases) UpdateShowLease(leaseID string, leaseToUpdate resources.LeaseV1) (resources.LeaseV1, error) {
	returnedLease, err := l.client.SendReceiveResource("PATCH", "GET", "/leases/"+leaseID, &leaseToUpdate)
	if err != nil {
		return resources.LeaseV1{}, err
	}
	if lease, ok := returnedLease.(*resources.LeaseV1); ok {
		return *lease, nil
	}
	return resources.LeaseV1{}, errors.New("UpdateShowLease call error.")
}
