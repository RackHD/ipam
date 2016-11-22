package ipamapi

import (
	"errors"

	"github.com/RackHD/ipam/resources"
)

//Pools can be used to query the Pools routes
type Pools struct {
	client *Client
}

// Index returns a list of Pools.
func (p *Pools) Index() (resources.PoolsV1, error) {
	pools, err := p.client.ReceiveResource("GET", "/pools", "", "")
	if err != nil {
		return resources.PoolsV1{}, err
	}

	if newPools, ok := pools.(*resources.PoolsV1); ok {
		return *newPools, nil
	}
	return resources.PoolsV1{}, errors.New("Pool Index call error.")
}

// Create a pool and returns the location.
func (p *Pools) Create(poolToCreate resources.PoolV1) (string, error) {

	poolLocation, err := p.client.SendResource("POST", "/pools", &poolToCreate)
	if err != nil {
		return "", err
	}
	return poolLocation, nil
}

// CreateShowPool creates a pool and then returns that pool.
func (p *Pools) CreateShowPool(poolToCreate resources.PoolV1) (resources.PoolV1, error) {
	receivedPool, err := p.client.SendReceiveResource("POST", "GET", "/pools", &poolToCreate)
	if err != nil {
		return resources.PoolV1{}, err
	}
	if pool, ok := receivedPool.(*resources.PoolV1); ok {
		return *pool, nil
	}
	return resources.PoolV1{}, errors.New("CreateShowPool call error.")
}

// Show returns the requested Pool.
func (p *Pools) Show(poolID string, poolToShow resources.PoolV1) (resources.PoolV1, error) {
	receivedPool, err := p.client.ReceiveResource("GET", "/pools/"+poolID, poolToShow.Type(), poolToShow.Version())
	if err != nil {
		return resources.PoolV1{}, err
	}
	if pool, ok := receivedPool.(*resources.PoolV1); ok {
		return *pool, nil
	}
	return resources.PoolV1{}, errors.New("Pools Show call error.")
}

// Update updates the requested Pool and returns its location.
func (p *Pools) Update(poolID string, poolToUpdate resources.PoolV1) (string, error) {
	location, err := p.client.SendResource("PATCH", "/pools/"+poolID, &poolToUpdate)
	if err != nil {
		return "", err
	}
	return location, nil
}

// UpdateShowPool updates a pool and then returns that pool.
func (p *Pools) UpdateShowPool(poolID string, poolToUpdate resources.PoolV1) (resources.PoolV1, error) {
	receivedPool, err := p.client.SendReceiveResource("PATCH", "GET", "/pools/"+poolID, &poolToUpdate)
	if err != nil {
		return resources.PoolV1{}, err
	}
	if pools, ok := receivedPool.(*resources.PoolV1); ok {
		return *pools, nil
	}
	return resources.PoolV1{}, errors.New("UpdateShowPool call error.")
}

// Delete removes the requested Pool and returns the location.
func (p *Pools) Delete(poolID string, poolToDelete resources.PoolV1) (string, error) {
	location, err := p.client.SendResource("DELETE", "/pools/"+poolID, &poolToDelete)
	if err != nil {
		return "", err
	}
	return location, nil
}
