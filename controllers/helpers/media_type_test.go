package helpers_test

import (
	. "github.com/RackHD/ipam/controllers/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MediaType", func() {
	Describe("NewMediaType", func() {
		It("should return an error if parsing failed", func() {
			mt, err := NewMediaType("")

			Expect(err).To(HaveOccurred())
			Expect(mt).To(BeZero())
		})

		It("should properly parse out an ipam resource", func() {
			mt, err := NewMediaType("application/vnd.ipam.pool")

			Expect(err).ToNot(HaveOccurred())
			Expect(mt).To(Equal(MediaType{"application/vnd.ipam.pool", ""}))
		})

		It("should properly parse out an ipam resource and format", func() {
			mt, err := NewMediaType("application/vnd.ipam.pool+xml")

			Expect(err).ToNot(HaveOccurred())
			Expect(mt).To(Equal(MediaType{"application/vnd.ipam.pool", ""}))
		})

		It("should propery parse out an ipam resource, format, and version", func() {
			mt, err := NewMediaType("application/vnd.ipam.pool+json;version=1.0.0")

			Expect(err).ToNot(HaveOccurred())
			Expect(mt).To(Equal(MediaType{"application/vnd.ipam.pool", "1.0.0"}))
		})

		It("should ignore any extra parameters", func() {
			mt, err := NewMediaType("application/vnd.ipam.pool+json;extra=lol;version=1.0.0")

			Expect(err).ToNot(HaveOccurred())
			Expect(mt).To(Equal(MediaType{"application/vnd.ipam.pool", "1.0.0"}))
		})
	})
})
