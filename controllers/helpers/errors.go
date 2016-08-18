package helpers

import (
	"fmt"
	"net/http"
)

// ErrorHandler provides an error handling capable HTTP handler.
type ErrorHandler func(http.ResponseWriter, *http.Request) error

// ServeHTTP implements the standard HTTP handler interface and handles ErrorHandler
// based handler functions.
func (handler ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := handler(w, r); err != nil {
		RenderError(w, err)
	}
}

// StatusError provides an additional interface method on errors to query an
// error for a HTTP status code.
type StatusError interface {
	error
	HTTPStatus() int
}

// HTTPStatusError implements the StatusError interface to provide HTTP status
// aware errors.
type HTTPStatusError struct {
	status  int
	message string
}

// NewHTTPStatusError returns a properly configured HTTPStatusError object.
func NewHTTPStatusError(status int, format string, a ...interface{}) HTTPStatusError {
	return HTTPStatusError{
		status:  status,
		message: fmt.Sprintf(format, a...),
	}
}

// Error is the implementation of the standard library error interface.
func (err HTTPStatusError) Error() string {
	return err.message
}

// HTTPStatus returns the configured status code for use in HTTP responses.
func (err HTTPStatusError) HTTPStatus() int {
	return err.status
}

// ErrorToHTTPStatus converts an error into a proper HTTP status code for use in our routers.
// It is a fallback for errors produced in libraries we don't  control which utilize error
// strings to denote their type of behavior.  Our primary case at this point is the mgo 'not found' error.
func ErrorToHTTPStatus(err error) int {
	switch err.Error() {
	case "not found":
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
