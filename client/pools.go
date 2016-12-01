package client

import (
	"errors"

	"github.com/RackHD/ipam/resources"
)

// IndexPools returns a list of Pools.
func (c *Client) IndexPools() (resources.PoolsV1, error) {
	pools, err := c.ReceiveResource("GET", "/pools", "", "")
	if err != nil {
		return resources.PoolsV1{}, err
	}

	if newPools, ok := pools.(*resources.PoolsV1); ok {
		return *newPools, nil
	}
	return resources.PoolsV1{}, errors.New("Pool Index call error.")
}

// CreatePool a pool and returns the location.
func (c *Client) CreatePool(poolToCreate resources.PoolV1) (string, error) {

	poolLocation, err := c.SendResource("POST", "/pools", &poolToCreate)
	if err != nil {
		return "", err
	}
	return poolLocation, nil
}

// CreateShowPool creates a pool and then returns that pool.
func (c *Client) CreateShowPool(poolToCreate resources.PoolV1) (resources.PoolV1, error) {
	receivedPool, err := c.SendReceiveResource("POST", "GET", "/pools", &poolToCreate)
	if err != nil {
		return resources.PoolV1{}, err
	}
	if pool, ok := receivedPool.(*resources.PoolV1); ok {
		return *pool, nil
	}
	return resources.PoolV1{}, errors.New("CreateShowPool call error.")
}

// ShowPool returns the requested Pool.
func (c *Client) ShowPool(poolID string, poolToShow resources.PoolV1) (resources.PoolV1, error) {
	receivedPool, err := c.ReceiveResource("GET", "/pools/"+poolID, poolToShow.Type(), poolToShow.Version())
	if err != nil {
		return resources.PoolV1{}, err
	}
	if pool, ok := receivedPool.(*resources.PoolV1); ok {
		return *pool, nil
	}
	return resources.PoolV1{}, errors.New("Pools Show call error.")
}

// UpdatePool updates the requested Pool and returns its location.
func (c *Client) UpdatePool(poolID string, poolToUpdate resources.PoolV1) (string, error) {
	location, err := c.SendResource("PATCH", "/pools/"+poolID, &poolToUpdate)
	if err != nil {
		return "", err
	}
	return location, nil
}

// UpdateShowPool updates a pool and then returns that pool.
func (c *Client) UpdateShowPool(poolID string, poolToUpdate resources.PoolV1) (resources.PoolV1, error) {
	receivedPool, err := c.SendReceiveResource("PATCH", "GET", "/pools/"+poolID, &poolToUpdate)
	if err != nil {
		return resources.PoolV1{}, err
	}
	if pools, ok := receivedPool.(*resources.PoolV1); ok {
		return *pools, nil
	}
	return resources.PoolV1{}, errors.New("UpdateShowPool call error.")
}

// DeletePool removes the requested Pool and returns the location.
func (c *Client) DeletePool(poolID string, poolToDelete resources.PoolV1) (string, error) {
	location, err := c.SendResource("DELETE", "/pools/"+poolID, &poolToDelete)
	if err != nil {
		return "", err
	}
	return location, nil
}
