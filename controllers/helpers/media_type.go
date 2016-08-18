package helpers

import (
	"mime"
	"regexp"
)

// MediaType represents a set of fields useful to our HTTP routers.
type MediaType struct {
	Type    string
	Version string
}

// NewMediaType parses an HTTP Accept/Content-Type header and returns a MediaType
// struct with key fields extracted for use in our HTTP routers.
func NewMediaType(requested string) (MediaType, error) {
	media, parameters, err := mime.ParseMediaType(requested)
	if err != nil {
		return MediaType{}, err
	}

	regex, err := regexp.Compile(`\+.+`)
	if err != nil {
		return MediaType{}, err
	}

	mediaType := MediaType{
		Type:    regex.ReplaceAllString(media, ""),
		Version: parameters["version"],
	}

	return mediaType, nil
}
