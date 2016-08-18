package interfaces

// ResourceCreator is a function signature used for requesting a versioned resource.
type ResourceCreator func(version string) (Resource, error)

// ResourceMarshaler marshals a given object to it's internal representation.
type ResourceMarshaler interface {
	Marshal(interface{}) error
}

// ResourceUnmarshaler marshals the internal resource representation to a given object (typically a model).
type ResourceUnmarshaler interface {
	Unmarshal() (interface{}, error)
}

// ResourceVersioner provides methods for obtaining information about a particular resource to be used in
// content negotiation in HTTP handlers.
type ResourceVersioner interface {
	Type() string
	Version() string
}

// Resource is a composition of interfaces required to represent a resource.
type Resource interface {
	ResourceMarshaler
	ResourceUnmarshaler
	ResourceVersioner
}
