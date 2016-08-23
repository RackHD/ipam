package interfaces

import "github.com/RackHD/ipam/models"

// Reservations interface defines the methods for implementing Reservation related business logic.
type Reservations interface {
	GetReservations(string) ([]models.Reservation, error)
	GetReservation(string) (models.Reservation, error)
	CreateReservation(models.Reservation) error
	UpdateReservation(models.Reservation) error
	DeleteReservation(string) error
}
