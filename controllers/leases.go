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

// LeasesController provides methods for handling requests to the Leases API.
type LeasesController struct {
	ipam interfaces.Ipam
}

// NewLeasesController returns a newly configured LeasesController.
func NewLeasesController(router *mux.Router, ipam interfaces.Ipam) (*LeasesController, error) {
	c := LeasesController{
		ipam: ipam,
	}

	router.Handle("/reservations/{id}/leases", helpers.ErrorHandler(c.Index)).Methods(http.MethodGet)
	router.Handle("/leases/{id}", helpers.ErrorHandler(c.Show)).Methods(http.MethodGet)
	router.Handle("/leases/{id}", helpers.ErrorHandler(c.Update)).Methods(http.MethodPut, http.MethodPatch)

	return &c, nil
}

// Index returns a list of Leases.
func (c *LeasesController) Index(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	reservations, err := c.ipam.GetLeases(vars["id"])
	if err != nil {
		return err
	}

	return helpers.RenderResource(w, r, resources.LeasesResourceType, http.StatusOK, reservations)
}

// Show returns the requested Lease.
func (c *LeasesController) Show(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	reservation, err := c.ipam.GetLease(vars["id"])
	if err != nil {
		return err
	}

	return helpers.RenderResource(w, r, resources.LeaseResourceType, http.StatusOK, reservation)
}

// Update updates the requested Lease.
func (c *LeasesController) Update(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	resource, err := helpers.AcceptResource(r, resources.LeaseResourceType)
	if err != nil {
		return err
	}

	if reservation, ok := resource.(models.Lease); ok {
		reservation.ID = bson.ObjectIdHex(vars["id"])

		err = c.ipam.UpdateLease(reservation)
		if err != nil {
			return err
		}

		return helpers.RenderLocation(w, http.StatusNoContent, fmt.Sprintf("/reservations/%s", reservation.ID.Hex()))
	}

	return fmt.Errorf("Invalid Resource Type")
}
