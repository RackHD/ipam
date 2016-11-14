package ipam

import (
	"github.com/RackHD/ipam/models"
	"gopkg.in/mgo.v2/bson"
)

// IpamCollectionLeases is the name of the Mongo collection which stores Leases.
const IpamCollectionLeases string = "leases"

// GetLeases returns a list of Leases.
func (ipam *Ipam) GetLeases(id string) ([]models.Lease, error) {
	session := ipam.session.Copy()
	defer session.Close()

	var reservations []models.Lease

	session.DB(IpamDatabase).C(IpamCollectionLeases).Find(bson.M{"reservation": bson.ObjectIdHex(id)}).All(&reservations)

	return reservations, nil
}

// GetLease returns the requested Lease.
func (ipam *Ipam) GetLease(id string) (models.Lease, error) {
	session := ipam.session.Copy()
	defer session.Close()

	var reservation models.Lease

	return reservation, session.DB(IpamDatabase).C(IpamCollectionLeases).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&reservation)
}

// UpdateLease updates a Lease.
func (ipam *Ipam) UpdateLease(reservation models.Lease) error {
	session := ipam.session.Copy()
	defer session.Close()

	return session.DB(IpamDatabase).C(IpamCollectionLeases).UpdateId(reservation.ID, reservation)
}

// DeleteLeases removes all leases associated to a subnet
func (ipam *Ipam) DeleteLeases(id string) error {
	session := ipam.session.Copy()
	defer session.Close()

	_, err := session.DB(IpamDatabase).C(IpamCollectionLeases).RemoveAll(bson.M{"subnet": bson.ObjectIdHex(id)})
	return err
}
