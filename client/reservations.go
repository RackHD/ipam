package ipamapi

import (
	"errors"

	"github.com/RackHD/ipam/resources"
)

//Reservations can be used to query the Reservations routes
type Reservations struct {
	client *Client
}

// Index returns a list of Reservations.
func (r *Reservations) Index(subnetID string) (resources.ReservationsV1, error) {
	receivedReservations, err := r.client.ReceiveResource("GET", "/subnets/"+subnetID+"/reservations", "", "")
	if err != nil {
		return resources.ReservationsV1{}, err
	}
	if reservations, ok := receivedReservations.(*resources.ReservationsV1); ok {
		return *reservations, nil
	}
	return resources.ReservationsV1{}, errors.New("Reservation Index call error.")
}

// Create a Reservation and return the location.
func (r *Reservations) Create(subnetID string, reservationToCreate resources.ReservationV1) (string, error) {
	reservationLocation, err := r.client.SendResource("POST", "/subnets/"+subnetID+"/reservations", &reservationToCreate)
	if err != nil {
		return "", err
	}
	return reservationLocation, nil
}

// CreateShowReservation creates a Reservation and then returns that Reservation.
func (r *Reservations) CreateShowReservation(subnetID string, reservationToCreate resources.ReservationV1) (resources.ReservationV1, error) {
	receivedReservation, err := r.client.SendReceiveResource("POST", "GET", "/subnets/"+subnetID+"/reservations", &reservationToCreate)
	if err != nil {
		return resources.ReservationV1{}, err
	}
	if reservation, ok := receivedReservation.(*resources.ReservationV1); ok {
		return *reservation, nil
	}
	return resources.ReservationV1{}, errors.New("CreateShowReservation call error.")
}

// Show returns the requested Reservation.
func (r *Reservations) Show(reservationID string, reservationToShow resources.ReservationV1) (resources.ReservationV1, error) {
	receivedReservation, err := r.client.ReceiveResource("GET", "/reservations/"+reservationID, reservationToShow.Type(), reservationToShow.Version())
	if err != nil {
		return resources.ReservationV1{}, err
	}
	if reservation, ok := receivedReservation.(*resources.ReservationV1); ok {
		return *reservation, nil
	}
	return resources.ReservationV1{}, errors.New("Reservation Show call error.")
}

// Update updates the requested Reservation and returns its location.
func (r *Reservations) Update(reservationID string, reservationToUpdate resources.ReservationV1) (string, error) {
	reservationLocation, err := r.client.SendResource("PATCH", "/reservations/"+reservationID, &reservationToUpdate)
	if err != nil {
		return "", err
	}
	return reservationLocation, nil
}

// UpdateShowReservation updates a Reservation and then returns that Reservation.
func (r *Reservations) UpdateShowReservation(reservationID string, reservationToUpdate resources.ReservationV1) (resources.ReservationV1, error) {
	receivedReservation, err := r.client.SendReceiveResource("PATCH", "GET", "/reservations/"+reservationID, &reservationToUpdate)
	if err != nil {
		return resources.ReservationV1{}, err
	}
	if reservation, ok := receivedReservation.(*resources.ReservationV1); ok {
		return *reservation, nil
	}
	return resources.ReservationV1{}, errors.New("UpdateShowReservation call error.")
}

// Delete removed the requested Reservation and returns the location.
func (r *Reservations) Delete(reservationID string, reservationToDelete resources.ReservationV1) (string, error) {
	reservationLocation, err := r.client.SendResource("DELETE", "/reservations/"+reservationID, &reservationToDelete)
	if err != nil {
		return "", err
	}
	return reservationLocation, nil
}
