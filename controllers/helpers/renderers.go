package helpers

import (
	"encoding/json"
	"fmt"
	"mime"
	"net/http"

	"github.com/RackHD/ipam/resources/factory"
)

// HeaderContentType represents the HTTP Content-Type header.
const HeaderContentType string = "Content-Type"

// HeaderAccept represents the HTTP Accept header.
const HeaderAccept string = "Accept"

// HeaderLocation represents the HTTP Location header.
const HeaderLocation string = "Location"

// RenderError writes the error and associated error HTTP status code to the response.
func RenderError(w http.ResponseWriter, err error) {
	if e, ok := err.(HTTPStatusError); ok {
		http.Error(w, e.Error(), e.HTTPStatus())
	} else {
		http.Error(w, err.Error(), ErrorToHTTPStatus(err))
	}
}

// RenderLocation writes the location and associated status code to the response.
func RenderLocation(w http.ResponseWriter, status int, location string) error {
	w.Header().Add(HeaderLocation, location)
	w.WriteHeader(status)
	w.Write([]byte{})

	return nil
}

// AcceptResource ...
func AcceptResource(r *http.Request, expected string) (interface{}, error) {
	mediaType, err := NewMediaType(r.Header.Get(HeaderContentType))
	if err != nil {
		return nil, NewHTTPStatusError(
			http.StatusUnsupportedMediaType,
			"Invalid Resource: %s", err,
		)
	}

	if mediaType.Type != expected {
		return nil, NewHTTPStatusError(
			http.StatusUnsupportedMediaType,
			"Unsupported Resource Type: %s != %s", mediaType.Type, expected,
		)
	}

	resource, err := factory.Require(mediaType.Type, mediaType.Version)
	if err != nil {
		return nil, NewHTTPStatusError(
			http.StatusUnsupportedMediaType,
			"Unsupported Resource Version: %s", err,
		)
	}

	err = json.NewDecoder(r.Body).Decode(&resource)
	if err != nil {
		return nil, err
	}

	return resource.Unmarshal()
}

// RenderResource accepts HTTP request/response objects as well as an intended HTTP status code,
// expected HTTP Content-Type, and a resource object to render.
func RenderResource(w http.ResponseWriter, r *http.Request, expected string, status int, object interface{}) error {
	// Ignore parsing errors and set the expected default.
	mediaType, _ := NewMediaType(r.Header.Get(HeaderAccept))
	if mediaType.Type != expected {
		mediaType.Type = expected
	}

	resource, err := factory.Request(mediaType.Type, mediaType.Version)
	if err != nil {
		return err
	}

	err = resource.Marshal(object)
	if err != nil {
		return err
	}

	data, err := json.Marshal(resource)
	if err != nil {
		return err
	}

	w.Header().Set(
		HeaderContentType,
		mime.FormatMediaType(
			fmt.Sprintf("%s+%s", resource.Type(), "json"),
			map[string]string{"version": resource.Version()},
		),
	)

	w.WriteHeader(status)
	w.Write(data)

	return nil
}
