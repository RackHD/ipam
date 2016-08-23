package interfaces

// Ipam interface is a collection of interfaces defined to support each
// area of business logic.  As new interfaces are defined they should be
// added to the Ipam interface to extend it's capabilities.
type Ipam interface {
	Pools
	Subnets
	Reservations
}
