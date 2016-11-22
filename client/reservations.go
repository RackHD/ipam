package ipamapi

import (
	"errors"

	"github.com/RackHD/ipam/resources"
)

// IndexReservations returns a list of Reservations.
func (c *Client) IndexReservations(subnetID string) (resources.ReservationsV1, error) {
	receivedReservations, err := c.ReceiveResource("GET", "/subnets/"+subnetID+"/reservations", "", "")
	if err != nil {
		return resources.ReservationsV1{}, err
	}
	if reservations, ok := receivedReservations.(*resources.ReservationsV1); ok {
		return *reservations, nil
	}
	return resources.ReservationsV1{}, errors.New("Reservation Index call error.")
}

// CreateReservation a Reservation and return the location.
func (c *Client) CreateReservation(subnetID string, reservationToCreate resources.ReservationV1) (string, error) {
	reservationLocation, err := c.SendResource("POST", "/subnets/"+subnetID+"/reservations", &reservationToCreate)
	if err != nil {
		return "", err
	}
	return reservationLocation, nil
}

// CreateShowReservation creates a Reservation and then returns that Reservation.
func (c *Client) CreateShowReservation(subnetID string, reservationToCreate resources.ReservationV1) (resources.ReservationV1, error) {
	receivedReservation, err := c.SendReceiveResource("POST", "GET", "/subnets/"+subnetID+"/reservations", &reservationToCreate)
	if err != nil {
		return resources.ReservationV1{}, err
	}
	if reservation, ok := receivedReservation.(*resources.ReservationV1); ok {
		return *reservation, nil
	}
	return resources.ReservationV1{}, errors.New("CreateShowReservation call error.")
}

// ShowReservation returns the requested Reservation.
func (c *Client) ShowReservation(reservationID string, reservationToShow resources.ReservationV1) (resources.ReservationV1, error) {
	receivedReservation, err := c.ReceiveResource("GET", "/reservations/"+reservationID, reservationToShow.Type(), reservationToShow.Version())
	if err != nil {
		return resources.ReservationV1{}, err
	}
	if reservation, ok := receivedReservation.(*resources.ReservationV1); ok {
		return *reservation, nil
	}
	return resources.ReservationV1{}, errors.New("Reservation Show call error.")
}

// UpdateReservation updates the requested Reservation and returns its location.
func (c *Client) UpdateReservation(reservationID string, reservationToUpdate resources.ReservationV1) (string, error) {
	reservationLocation, err := c.SendResource("PATCH", "/reservations/"+reservationID, &reservationToUpdate)
	if err != nil {
		return "", err
	}
	return reservationLocation, nil
}

// UpdateShowReservation updates a Reservation and then returns that Reservation.
func (c *Client) UpdateShowReservation(reservationID string, reservationToUpdate resources.ReservationV1) (resources.ReservationV1, error) {
	receivedReservation, err := c.SendReceiveResource("PATCH", "GET", "/reservations/"+reservationID, &reservationToUpdate)
	if err != nil {
		return resources.ReservationV1{}, err
	}
	if reservation, ok := receivedReservation.(*resources.ReservationV1); ok {
		return *reservation, nil
	}
	return resources.ReservationV1{}, errors.New("UpdateShowReservation call error.")
}

// DeleteReservation removed the requested Reservation and returns the location.
func (c *Client) DeleteReservation(reservationID string, reservationToDelete resources.ReservationV1) (string, error) {
	reservationLocation, err := c.SendResource("DELETE", "/reservations/"+reservationID, &reservationToDelete)
	if err != nil {
		return "", err
	}
	return reservationLocation, nil
}
