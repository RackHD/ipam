package controllers

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	"github.com/RackHD/ipam/controllers/helpers"
	"github.com/RackHD/ipam/interfaces"
	"github.com/RackHD/ipam/models"
	"github.com/RackHD/ipam/resources"
)

// PoolsController provides methods for handling requests to the Pools API.
type PoolsController struct {
	ipam interfaces.Ipam
}

// NewPoolsController returns a newly configured PoolsController.
func NewPoolsController(router *mux.Router, ipam interfaces.Ipam) (*PoolsController, error) {
	c := PoolsController{
		ipam: ipam,
	}

	router.Handle("/pools", helpers.ErrorHandler(c.Index)).Methods(http.MethodGet)
	router.Handle("/pools", helpers.ErrorHandler(c.Create)).Methods(http.MethodPost)
	router.Handle("/pools/{id}", helpers.ErrorHandler(c.Show)).Methods(http.MethodGet)
	router.Handle("/pools/{id}", helpers.ErrorHandler(c.Update)).Methods(http.MethodPut, http.MethodPatch)
	router.Handle("/pools/{id}", helpers.ErrorHandler(c.Delete)).Methods(http.MethodDelete)

	return &c, nil
}

// Index returns a list of Pools.
func (c *PoolsController) Index(w http.ResponseWriter, r *http.Request) error {
	pools, err := c.ipam.GetPools()
	if err != nil {
		return err
	}

	return helpers.RenderResource(w, r, resources.PoolsResourceType, http.StatusOK, pools)
}

// Create creates a Pool.
func (c *PoolsController) Create(w http.ResponseWriter, r *http.Request) error {
	resource, err := helpers.AcceptResource(r, resources.PoolResourceType)
	if err != nil {
		return err
	}

	if pool, ok := resource.(models.Pool); ok {
		err = c.ipam.CreatePool(pool)
		if err != nil {
			return err
		}

		return helpers.RenderLocation(w, r, http.StatusCreated, fmt.Sprintf("/pools/%s", pool.ID.Hex()))
	}

	return fmt.Errorf("Invalid Resource Type")
}

// Show returns the requested Pool.
func (c *PoolsController) Show(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	pool, err := c.ipam.GetPool(vars["id"])
	if err != nil {
		return err
	}

	return helpers.RenderResource(w, r, resources.PoolResourceType, http.StatusOK, pool)
}

// Update updates the requested Pool.
func (c *PoolsController) Update(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	resource, err := helpers.AcceptResource(r, resources.PoolResourceType)
	if err != nil {
		return err
	}

	if pool, ok := resource.(models.Pool); ok {
		pool.ID = bson.ObjectIdHex(vars["id"])

		err = c.ipam.UpdatePool(pool)
		if err != nil {
			return err
		}

		return helpers.RenderLocation(w, r, http.StatusNoContent, fmt.Sprintf("/pools/%s", pool.ID.Hex()))
	}

	return fmt.Errorf("Invalid Resource Type")
}

// Delete removes the requested Pool.
func (c *PoolsController) Delete(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	err := c.ipam.DeletePool(vars["id"])
	if err != nil {
		return err
	}

	return helpers.RenderLocation(w, r, http.StatusOK, fmt.Sprintf("/pools"))
}
