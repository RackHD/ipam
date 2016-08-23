package controllers_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/RackHD/ipam/controllers"
	"github.com/RackHD/ipam/controllers/helpers"
	"github.com/RackHD/ipam/resources"
	"github.com/RackHD/ipam/resources/factory"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reservations Controller", func() {
	var (
		router *mux.Router
		mock   *MockIpam
		server *httptest.Server
	)

	BeforeEach(func() {
		router = mux.NewRouter().StrictSlash(true)
		mock = NewMockIpam()

		_, err := controllers.NewReservationsController(router, mock)
		Expect(err).NotTo(HaveOccurred())

		server = httptest.NewServer(router)
	})

	Describe("Index", func() {
		It("should return a 200 status code", func() {
			req := NewRequest(
				"GET",
				server.URL+"/subnets/578af30bbc63780007d99195/reservations",
				nil,
			)

			res := Do(req)

			Expect(res.StatusCode).To(Equal(http.StatusOK))
		})

		It("should return a list of reservations", func() {
			req := NewRequest(
				"GET",
				server.URL+"/subnets/578af30bbc63780007d99195/reservations",
				nil,
			)

			res := Do(req)

			Expect(res.StatusCode).To(Equal(http.StatusOK))

			resource, err := factory.Request(resources.ReservationsResourceType, "1.0.0")
			Expect(err).ToNot(HaveOccurred())

			err = resource.Marshal(mock.Reservations)
			Expect(err).ToNot(HaveOccurred())

			json, err := json.Marshal(resource)
			Expect(err).ToNot(HaveOccurred())

			body, _ := ioutil.ReadAll(res.Body)
			Expect(err).ToNot(HaveOccurred())

			Expect(json).To(Equal(body))
		})

		It("should return a 500 if an error occurs", func() {
			mock.Err = fmt.Errorf("Fail")

			req := NewRequest(
				"GET",
				server.URL+"/subnets/578af30bbc63780007d99195/reservations",
				nil,
			)

			res := Do(req)

			Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
		})
	})

	Describe("Show", func() {
		It("should return a 200 status code", func() {
			req := NewRequest(
				"GET",
				server.URL+"/reservations/578af30bbc63780007d99195",
				nil,
			)

			res := Do(req)

			Expect(res.StatusCode).To(Equal(http.StatusOK))
		})

		It("should return the requested subnet", func() {
			req := NewRequest(
				"GET",
				server.URL+"/reservations/578af30bbc63780007d99195",
				nil,
			)

			res := Do(req)
			defer res.Body.Close()

			Expect(res.StatusCode).To(Equal(http.StatusOK))

			resource, err := factory.Request(resources.ReservationResourceType, "1.0.0")
			Expect(err).ToNot(HaveOccurred())

			err = resource.Marshal(mock.Reservations[0])
			Expect(err).ToNot(HaveOccurred())

			json, err := json.Marshal(resource)
			Expect(err).ToNot(HaveOccurred())

			body, err := ioutil.ReadAll(res.Body)
			Expect(err).ToNot(HaveOccurred())

			Expect(json).To(Equal(body))
		})

		It("should return a 500 if an error occurs", func() {
			mock.Err = fmt.Errorf("Fail")

			req := NewRequest(
				"GET",
				server.URL+"/reservations/578af30bbc63780007d99195",
				nil,
			)

			res := Do(req)

			Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
		})

		It("should return a 404 if the resource is not found", func() {
			mock.Err = fmt.Errorf("not found")

			req := NewRequest(
				"GET",
				server.URL+"/reservations/578af30bbc63780007d99195",
				nil,
			)

			res := Do(req)

			Expect(res.StatusCode).To(Equal(http.StatusNotFound))
		})
	})

	Describe("Create", func() {
		It("should return a 201 status code", func() {
			req := NewRequest(
				"POST",
				server.URL+"/subnets/578af30bbc63780007d99195/reservations",
				strings.NewReader(`{
					"name": "New Reservation",
					"tags": ["New Reservation Tag"],
					"metadata": {
						"one": 1
					}
				}`),
			)

			req.Header.Set(helpers.HeaderContentType, resources.ReservationResourceType+";version=1.0.0")

			res := Do(req)

			Expect(res.StatusCode).To(Equal(http.StatusCreated))
		})

		It("should add the new model with the corresponding fields", func() {
			req := NewRequest(
				"POST",
				server.URL+"/subnets/578af30bbc63780007d99195/reservations",
				strings.NewReader(`{
					"name": "New Reservation",
					"tags": ["New Reservation Tag"],
					"metadata": {
						"one": 1
					}
				}`),
			)

			req.Header.Set(helpers.HeaderContentType, resources.ReservationResourceType+";version=1.0.0")

			res := Do(req)

			Expect(res.StatusCode).To(Equal(http.StatusCreated))

			Expect(mock.ReservationCreated).NotTo(BeZero())
			Expect(mock.ReservationCreated.Name).To(Equal("New Reservation"))
			Expect(mock.ReservationCreated.Tags).To(Equal([]string{"New Reservation Tag"}))
		})

		It("should return a 415 status code if no resource type and version are specified", func() {
			req := NewRequest(
				"POST",
				server.URL+"/subnets/578af30bbc63780007d99195/reservations",
				strings.NewReader(`{
					"name": "New Reservation",
					"tags": ["New Reservation Tag"],
					"metadata": {
						"one": 1
					}
				}`),
			)

			res := Do(req)

			Expect(res.StatusCode).To(Equal(http.StatusUnsupportedMediaType))
		})

		It("should return a 415 status code if the resource version is not available", func() {
			req := NewRequest(
				"POST",
				server.URL+"/subnets/578af30bbc63780007d99195/reservations",
				strings.NewReader(`{
					"name": "New Reservation",
					"tags": ["New Reservation Tag"],
					"metadata": {
						"one": 1
					}
				}`),
			)

			req.Header.Set(helpers.HeaderContentType, resources.ReservationResourceType+";version=0.0.7")

			res := Do(req)

			Expect(res.StatusCode).To(Equal(http.StatusUnsupportedMediaType))
		})

		It("should return a 415 status code if the resource specified is the wrong type", func() {
			req := NewRequest(
				"POST",
				server.URL+"/subnets/578af30bbc63780007d99195/reservations",
				strings.NewReader(`{
					"name": "New Reservation",
					"tags": ["New Reservation Tag"],
					"metadata": {
						"one": 1
					}
				}`),
			)

			req.Header.Set(helpers.HeaderContentType, "application/vnd.ipam.lol;version=1.0.0")

			res := Do(req)

			Expect(res.StatusCode).To(Equal(http.StatusUnsupportedMediaType))
		})

		It("should return a 500 if an error occurs", func() {
			mock.Err = fmt.Errorf("Fail")

			req := NewRequest(
				"POST",
				server.URL+"/subnets/578af30bbc63780007d99195/reservations",
				strings.NewReader(`{
					"name": "New Reservation",
					"tags": ["New Reservation Tag"],
					"metadata": {
						"one": 1
					}
				}`),
			)

			req.Header.Set(helpers.HeaderContentType, resources.ReservationResourceType+";version=1.0.0")

			res := Do(req)

			Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
		})

	})

	Describe("Update", func() {
		It("should return a 204 status code", func() {
			req := NewRequest(
				"PUT",
				server.URL+"/reservations/578af30bbc63780007d99195",
				strings.NewReader(`{
					"name": "Updated Reservation",
					"tags": ["Updated Reservation Tag"],
					"metadata": {
						"one": "one"
					}
				}`),
			)

			req.Header.Set(helpers.HeaderContentType, resources.ReservationResourceType+";version=1.0.0")

			res := Do(req)

			Expect(res.StatusCode).To(Equal(http.StatusNoContent))
		})

		It("should update the model with the corresponding fields", func() {
			req := NewRequest(
				"PUT",
				server.URL+"/reservations/578af30bbc63780007d99195",
				strings.NewReader(`{
					"name": "Updated Reservation",
					"tags": ["Updated Reservation Tag"],
					"metadata": {
						"one": "one"
					}
				}`),
			)

			req.Header.Set(helpers.HeaderContentType, resources.ReservationResourceType+";version=1.0.0")

			res := Do(req)

			Expect(res.StatusCode).To(Equal(http.StatusNoContent))

			Expect(mock.ReservationUpdated).NotTo(BeZero())
			Expect(mock.ReservationUpdated.Name).To(Equal("Updated Reservation"))
			Expect(mock.ReservationUpdated.Tags).To(Equal([]string{"Updated Reservation Tag"}))
		})

		It("should return a 415 status code if no resource type and version are specified", func() {
			req := NewRequest(
				"PUT",
				server.URL+"/reservations/578af30bbc63780007d99195",
				strings.NewReader(`{
					"name": "Updated Reservation",
					"tags": ["Updated Reservation Tag"],
					"metadata": {
						"one": "one"
					}
				}`),
			)

			res := Do(req)

			Expect(res.StatusCode).To(Equal(http.StatusUnsupportedMediaType))
		})

		It("should return a 415 status code if the resource version is not available", func() {
			req := NewRequest(
				"PUT",
				server.URL+"/reservations/578af30bbc63780007d99195",
				strings.NewReader(`{
					"name": "Updated Reservation",
					"tags": ["Updated Reservation Tag"],
					"metadata": {
						"one": "one"
					}
				}`),
			)

			req.Header.Set(helpers.HeaderContentType, resources.ReservationResourceType+";version=0.0.7")

			res := Do(req)

			Expect(res.StatusCode).To(Equal(http.StatusUnsupportedMediaType))
		})

		It("should return a 415 status code if the resource specified is the wrong type", func() {
			req := NewRequest(
				"PUT",
				server.URL+"/reservations/578af30bbc63780007d99195",
				strings.NewReader(`{
					"name": "Updated Reservation",
					"tags": ["Updated Reservation Tag"],
					"metadata": {
						"one": "one"
					}
				}`),
			)

			req.Header.Set(helpers.HeaderContentType, "application/vnd.ipam.lol;version=1.0.0")

			res := Do(req)

			Expect(res.StatusCode).To(Equal(http.StatusUnsupportedMediaType))
		})

		It("should return a 500 if an error occurs", func() {
			mock.Err = fmt.Errorf("Fail")

			req := NewRequest(
				"PUT",
				server.URL+"/reservations/578af30bbc63780007d99195",
				strings.NewReader(`{
					"name": "Updated Reservation",
					"tags": ["Updated Reservation Tag"],
					"metadata": {
						"one": "one"
					}
				}`),
			)

			req.Header.Set(helpers.HeaderContentType, resources.ReservationResourceType+";version=1.0.0")

			res := Do(req)

			Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
		})

		It("should return a 404 if the resource is not found", func() {
			mock.Err = fmt.Errorf("not found")

			req := NewRequest(
				"PUT",
				server.URL+"/reservations/578af30bbc63780007d99195",
				strings.NewReader(`{
					"name": "Updated Reservation",
					"tags": ["Updated Reservation Tag"],
					"metadata": {
						"one": "one"
					}
				}`),
			)

			req.Header.Set(helpers.HeaderContentType, resources.ReservationResourceType+";version=1.0.0")

			res := Do(req)

			Expect(res.StatusCode).To(Equal(http.StatusNotFound))
		})
	})

	Describe("Delete", func() {
		It("should return a 200 status code", func() {
			req := NewRequest(
				"DELETE",
				server.URL+"/reservations/578af30bbc63780007d99195",
				nil,
			)

			res := Do(req)

			Expect(res.StatusCode).To(Equal(http.StatusOK))
		})

		It("should delete the model with the corresponding id", func() {
			req := NewRequest(
				"DELETE",
				server.URL+"/reservations/578af30bbc63780007d99195",
				nil,
			)

			res := Do(req)

			Expect(res.StatusCode).To(Equal(http.StatusOK))

			Expect(mock.ReservationDeleted).To(Equal("578af30bbc63780007d99195"))
		})

		It("should return a 500 if an error occurs", func() {
			mock.Err = fmt.Errorf("Fail")

			req := NewRequest(
				"DELETE",
				server.URL+"/reservations/578af30bbc63780007d99195",
				nil,
			)

			res := Do(req)

			Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
		})

		It("should return a 404 if the resource is not found", func() {
			mock.Err = fmt.Errorf("not found")

			req := NewRequest(
				"DELETE",
				server.URL+"/reservations/578af30bbc63780007d99195",
				nil,
			)

			res := Do(req)

			Expect(res.StatusCode).To(Equal(http.StatusNotFound))
		})
	})
})
