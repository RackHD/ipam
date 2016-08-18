package ipam

import (
	"gopkg.in/mgo.v2"
)

// IpamDatabase is the name of the Mongo database used to store IPAM models.
const IpamDatabase string = "ipam"

// Ipam is an object which implements the IPAM business logic interface.
type Ipam struct {
	session *mgo.Session
}

// NewIpam returns a new Ipam object.
func NewIpam(session *mgo.Session) (*Ipam, error) {
	// This is starting us off with an object to play with by default.
	return &Ipam{
		session: session,
	}, nil
}
