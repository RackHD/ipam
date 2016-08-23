package interfaces

import "github.com/RackHD/ipam/models"

// Subnets interface defines the methods for implementing Subnet related business logic.
type Subnets interface {
	GetSubnets(string) ([]models.Subnet, error)
	GetSubnet(string) (models.Subnet, error)
	CreateSubnet(models.Subnet) error
	UpdateSubnet(models.Subnet) error
	DeleteSubnet(string) error
}
