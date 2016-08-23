package controllers

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/RackHD/ipam/controllers/helpers"
	"github.com/RackHD/ipam/interfaces"
	"github.com/RackHD/ipam/models"
	"github.com/RackHD/ipam/resources"
	"github.com/gorilla/mux"
)

// ReservationsController provides methods for handling requests to the Reservations API.
type ReservationsController struct {
	ipam interfaces.Ipam
}

// NewReservationsController returns a newly configured ReservationsController.
func NewReservationsController(router *mux.Router, ipam interfaces.Ipam) (*ReservationsController, error) {
	c := ReservationsController{
		ipam: ipam,
	}

	router.Handle("/subnets/{id}/reservations", helpers.ErrorHandler(c.Index)).Methods(http.MethodGet)
	router.Handle("/subnets/{id}/reservations", helpers.ErrorHandler(c.Create)).Methods(http.MethodPost)
	router.Handle("/reservations/{id}", helpers.ErrorHandler(c.Show)).Methods(http.MethodGet)
	router.Handle("/reservations/{id}", helpers.ErrorHandler(c.Update)).Methods(http.MethodPut, http.MethodPatch)
	router.Handle("/reservations/{id}", helpers.ErrorHandler(c.Delete)).Methods(http.MethodDelete)

	return &c, nil
}

// Index returns a list of Reservations.
func (c *ReservationsController) Index(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	reservations, err := c.ipam.GetReservations(vars["id"])
	if err != nil {
		return err
	}

	return helpers.RenderResource(w, r, resources.ReservationsResourceType, http.StatusOK, reservations)
}

// Create creates a Reservation.
func (c *ReservationsController) Create(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	resource, err := helpers.AcceptResource(r, resources.ReservationResourceType)
	if err != nil {
		return err
	}

	if reservation, ok := resource.(models.Reservation); ok {
		reservation.Subnet = bson.ObjectIdHex(vars["id"])

		err = c.ipam.CreateReservation(reservation)
		if err != nil {
			return err
		}

		return helpers.RenderLocation(w, r, http.StatusCreated, fmt.Sprintf("/reservations/%s", reservation.ID.Hex()))
	}

	return fmt.Errorf("Invalid Resource Type")
}

// Show returns the requested Reservation.
func (c *ReservationsController) Show(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	reservation, err := c.ipam.GetReservation(vars["id"])
	if err != nil {
		return err
	}

	return helpers.RenderResource(w, r, resources.ReservationResourceType, http.StatusOK, reservation)
}

// Update updates the requested Reservation.
func (c *ReservationsController) Update(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	resource, err := helpers.AcceptResource(r, resources.ReservationResourceType)
	if err != nil {
		return err
	}

	if reservation, ok := resource.(models.Reservation); ok {
		reservation.ID = bson.ObjectIdHex(vars["id"])

		err = c.ipam.UpdateReservation(reservation)
		if err != nil {
			return err
		}

		return helpers.RenderLocation(w, r, http.StatusNoContent, fmt.Sprintf("/reservations/%s", reservation.ID.Hex()))
	}

	return fmt.Errorf("Invalid Resource Type")
}

// Delete removes the requested Reservation.
func (c *ReservationsController) Delete(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	err := c.ipam.DeleteReservation(vars["id"])
	if err != nil {
		return err
	}

	return helpers.RenderLocation(w, r, http.StatusOK, fmt.Sprintf("/reservations"))
}
