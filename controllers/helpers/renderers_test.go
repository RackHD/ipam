package helpers_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"

	. "github.com/RackHD/ipam/controllers/helpers"
	"github.com/RackHD/ipam/models"
	"github.com/RackHD/ipam/resources"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type FakeResponseWriter struct {
	headers http.Header
	status  int
	data    []byte
}

func NewFakeResponseWriter(status int, data []byte) FakeResponseWriter {
	return FakeResponseWriter{
		headers: http.Header{},
		status:  status,
		data:    data,
	}
}

func (w FakeResponseWriter) Header() http.Header {
	return w.headers
}

func (w FakeResponseWriter) Write(data []byte) (int, error) {
	left := strings.TrimSpace(string(data))
	right := strings.TrimSpace(string(w.data))
	Expect(left).To(Equal(right))
	return len(data), nil
}

func (w FakeResponseWriter) WriteHeader(status int) {
	Expect(status).To(Equal(w.status))
}

func NewRequest(method, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	Expect(err).NotTo(HaveOccurred())
	return req
}

var _ = Describe("Renderers", func() {
	Describe("RenderError", func() {
		It("should render HTTPStatusError status code and message", func() {
			err := NewHTTPStatusError(http.StatusExpectationFailed, "Fake Error")

			w := NewFakeResponseWriter(http.StatusExpectationFailed, ([]byte)(err.Error()))

			RenderError(w, err)
		})

		It("should render default error status code and message", func() {
			err := fmt.Errorf("Fake Error")

			w := NewFakeResponseWriter(http.StatusInternalServerError, ([]byte)(err.Error()))

			RenderError(w, err)
		})
	})

	Describe("RenderLocation", func() {
		It("should render the given location header, status, and empty body", func() {
			w := NewFakeResponseWriter(http.StatusOK, []byte{})

			RenderLocation(w, http.StatusOK, "/location")

			Expect(w.Header().Get("Location")).To(Equal("/location"))
		})
	})

	Describe("AcceptResource", func() {
		It("should render unsupported media type if the mime type was not parsed", func() {
			r := NewRequest(
				"POST",
				"http://fake/pools",
				strings.NewReader(`{
					"name": "New Pool",
					"tags": ["New Pool Tag"],
					"metadata": {
						"one": 1
					}
				}`),
			)

			r.Header.Set(HeaderContentType, ";invalid")

			_, err := AcceptResource(r, resources.PoolResourceType)

			Expect(err.Error()).To(Equal("Invalid Resource: mime: no media type"))
		})

		It("should render unsupported media type if the mime type was not expected", func() {
			r := NewRequest(
				"POST",
				"http://fake/pools",
				strings.NewReader(`{
					"name": "New Pool",
					"tags": ["New Pool Tag"],
					"metadata": {
						"one": 1
					}
				}`),
			)

			r.Header.Set(HeaderContentType, "vnd.ipam.reservation+json;version=1.0.0")

			_, err := AcceptResource(r, resources.PoolResourceType)

			Expect(err.Error()).To(Equal("Unsupported Resource Type: vnd.ipam.reservation != application/vnd.ipam.pool"))
		})

		It("should render unsupported media type if the resource version is not supported", func() {
			r := NewRequest(
				"POST",
				"http://fake/pools",
				strings.NewReader(`{
					"name": "New Pool",
					"tags": ["New Pool Tag"],
					"metadata": {
						"one": 1
					}
				}`),
			)

			r.Header.Set(HeaderContentType, resources.PoolResourceType+";version=0.0.0")

			_, err := AcceptResource(r, resources.PoolResourceType)

			Expect(err.Error()).To(Equal("Unsupported Resource Version: Require: Unable to locate resource application/vnd.ipam.pool, version 0.0.0."))
		})

		It("should render an error if the body of the request was not valid json", func() {
			r := NewRequest(
				"POST",
				"http://fake/pools",
				strings.NewReader(`invalid json`),
			)

			r.Header.Set(HeaderContentType, resources.PoolResourceType+";version=1.0.0")

			_, err := AcceptResource(r, resources.PoolResourceType)
			Expect(err.Error()).To(Equal("invalid character 'i' looking for beginning of value"))
		})

		It("should unmarshal a valid resource", func() {
			r := NewRequest(
				"POST",
				"http://fake/pools",
				strings.NewReader(`{
					"name": "New Pool",
					"tags": ["New Pool Tag"],
					"metadata": {
						"one": 1
					}
				}`),
			)

			r.Header.Set(HeaderContentType, resources.PoolResourceType+";version=1.0.0")

			_, err := AcceptResource(r, resources.PoolResourceType)

			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("RenderResource", func() {
		var (
			subnet = models.Subnet{
				ID:       bson.NewObjectId(),
				Name:     "Subnet Name",
				Tags:     []string{"Subnet Tag"},
				Metadata: "Subnet Metadata",
				Pool:     bson.NewObjectId(),
			}
			data = []byte{}
		)

		BeforeSuite(func() {
			resource := resources.SubnetV1{}

			err := resource.Marshal(subnet)
			Expect(err).ToNot(HaveOccurred())

			data, err = json.Marshal(resource)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should set the media type to the expected type if the media type is absent or incorrect", func() {
			r := NewRequest(
				"GET",
				"http://fake/pools",
				nil,
			)

			r.Header.Set(HeaderAccept, "unexpected")

			w := NewFakeResponseWriter(http.StatusOK, data)

			err := RenderResource(w, r, resources.SubnetResourceType, http.StatusOK, subnet)

			Expect(err).ToNot(HaveOccurred())
		})

		It("should return an error if the requested resource is not present", func() {
			r := NewRequest(
				"GET",
				"http://fake/pools",
				nil,
			)

			r.Header.Set(HeaderAccept, resources.SubnetResourceType+";version=1.0.0")

			w := NewFakeResponseWriter(http.StatusOK, data)

			err := RenderResource(w, r, "invalid", http.StatusOK, subnet)

			Expect(err).To(HaveOccurred())
		})

		It("should set the content type header to the correct resource value", func() {
			r := NewRequest(
				"GET",
				"http://fake/pools",
				nil,
			)

			r.Header.Set(HeaderAccept, resources.SubnetResourceType+";version=1.0.0")

			w := NewFakeResponseWriter(http.StatusOK, data)

			err := RenderResource(w, r, resources.SubnetResourceType, http.StatusOK, subnet)

			Expect(err).ToNot(HaveOccurred())

			Expect(w.Header().Get(HeaderContentType)).To(Equal("application/vnd.ipam.subnet+json; version=1.0.0"))
		})
	})
})
