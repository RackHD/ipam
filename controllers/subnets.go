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

// SubnetsController provides methods for handling requests to the Subnets API.
type SubnetsController struct {
	ipam interfaces.Ipam
}

// NewSubnetsController returns a newly configured SubnetsController.
func NewSubnetsController(router *mux.Router, ipam interfaces.Ipam) (*SubnetsController, error) {
	c := SubnetsController{
		ipam: ipam,
	}

	router.Handle("/pools/{id}/subnets", helpers.ErrorHandler(c.Index)).Methods(http.MethodGet)
	router.Handle("/pools/{id}/subnets", helpers.ErrorHandler(c.Create)).Methods(http.MethodPost)
	router.Handle("/subnets/{id}", helpers.ErrorHandler(c.Show)).Methods(http.MethodGet)
	router.Handle("/subnets/{id}", helpers.ErrorHandler(c.Update)).Methods(http.MethodPut, http.MethodPatch)
	router.Handle("/subnets/{id}", helpers.ErrorHandler(c.Delete)).Methods(http.MethodDelete)

	return &c, nil
}

// Index returns a list of Subnets.
func (c *SubnetsController) Index(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	subnets, err := c.ipam.GetSubnets(vars["id"])
	if err != nil {
		return err
	}

	return helpers.RenderResource(w, r, resources.SubnetsResourceType, http.StatusOK, subnets)
}

// Create creates a Subnet.
func (c *SubnetsController) Create(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	resource, err := helpers.AcceptResource(r, resources.SubnetResourceType)
	if err != nil {
		return err
	}

	if subnet, ok := resource.(models.Subnet); ok {
		subnet.Pool = bson.ObjectIdHex(vars["id"])

		err = c.ipam.CreateSubnet(subnet)
		if err != nil {
			return err
		}

		return helpers.RenderLocation(w, r, http.StatusCreated, fmt.Sprintf("/subnets/%s", subnet.ID.Hex()))
	}

	return fmt.Errorf("Invalid Resource Type")
}

// Show returns the requested Subnet.
func (c *SubnetsController) Show(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	subnet, err := c.ipam.GetSubnet(vars["id"])
	if err != nil {
		return err
	}

	return helpers.RenderResource(w, r, resources.SubnetResourceType, http.StatusOK, subnet)
}

// Update updates the requested Subnet.
func (c *SubnetsController) Update(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	resource, err := helpers.AcceptResource(r, resources.SubnetResourceType)
	if err != nil {
		return err
	}

	if subnet, ok := resource.(models.Subnet); ok {
		subnet.ID = bson.ObjectIdHex(vars["id"])

		err = c.ipam.UpdateSubnet(subnet)
		if err != nil {
			return err
		}

		return helpers.RenderLocation(w, r, http.StatusNoContent, fmt.Sprintf("/subnets/%s", subnet.ID.Hex()))
	}

	return fmt.Errorf("Invalid Resource Type")
}

// Delete removes the requested Subnet.
func (c *SubnetsController) Delete(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	err := c.ipam.DeleteSubnet(vars["id"])
	if err != nil {
		return err
	}

	return helpers.RenderLocation(w, r, http.StatusOK, fmt.Sprintf("/subnets"))
}
