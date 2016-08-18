package interfaces

import "github.com/RackHD/ipam/models"

// Pools interface defines the methods for implementing Pool related business logic.
type Pools interface {
	GetPools() ([]models.Pool, error)
	GetPool(string) (models.Pool, error)
	CreatePool(models.Pool) error
	UpdatePool(models.Pool) error
	DeletePool(string) error
}
