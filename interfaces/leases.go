package interfaces

import "github.com/RackHD/ipam/models"

// Leases interface defines the methods for implementing Lease related business logic.
type Leases interface {
	GetLeases(string) ([]models.Lease, error)
	GetLease(string) (models.Lease, error)
	UpdateLease(models.Lease) error
}
