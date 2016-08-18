package helpers_test

import (
	"fmt"
	"net/http"

	. "github.com/RackHD/ipam/controllers/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ErrorToHTTPStatus", func() {
	It("should return http.StatusNotFound if the error specifies 'not found'", func() {
		err := fmt.Errorf("not found")

		Expect(ErrorToHTTPStatus(err)).To(Equal(http.StatusNotFound))
	})

	It("should return http.StatusInternalServerError for all other errors", func() {
		err := fmt.Errorf("any other error")

		Expect(ErrorToHTTPStatus(err)).To(Equal(http.StatusInternalServerError))
	})
})

var _ = Describe("HTTPStatusError", func() {
	It("should return the provided status via the HTTPStatus method", func() {
		err := NewHTTPStatusError(100, "hello")
		Expect(err.HTTPStatus()).To(Equal(100))
	})

	It("should return the provided message via the Error method", func() {
		err := NewHTTPStatusError(100, "hello")
		Expect(err.Error()).To(Equal("hello"))
	})

	It("should format the provided message using fmt.Sprintf", func() {
		err := NewHTTPStatusError(100, "hello %s", "world")
		Expect(err.Error()).To(Equal("hello world"))
	})
})
