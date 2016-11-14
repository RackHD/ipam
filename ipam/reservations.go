package ipam

import (
	"github.com/RackHD/ipam/models"
	"gopkg.in/mgo.v2/bson"
)

// IpamCollectionReservations is the name of the Mongo collection which stores Reservations.
const IpamCollectionReservations string = "reservations"

// GetReservations returns a list of Reservations.
func (ipam *Ipam) GetReservations(id string) ([]models.Reservation, error) {
	session := ipam.session.Copy()
	defer session.Close()

	var reservations []models.Reservation

	session.DB(IpamDatabase).C(IpamCollectionReservations).Find(bson.M{"subnet": bson.ObjectIdHex(id)}).All(&reservations)

	return reservations, nil
}

// GetReservation returns the requested Reservation.
func (ipam *Ipam) GetReservation(id string) (models.Reservation, error) {
	session := ipam.session.Copy()
	defer session.Close()

	var reservation models.Reservation

	return reservation, session.DB(IpamDatabase).C(IpamCollectionReservations).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&reservation)
}

// CreateReservation creates a Reservation.
func (ipam *Ipam) CreateReservation(reservation models.Reservation) error {
	session := ipam.session.Copy()
	defer session.Close()

	err := session.DB(IpamDatabase).C(IpamCollectionReservations).Insert(reservation)
	if err != nil {
		return err
	}

	return session.DB(IpamDatabase).C(IpamCollectionLeases).Update(
		bson.M{"reservation": nil, "subnet": reservation.Subnet},
		bson.M{"$set": bson.M{"reservation": reservation.ID}},
	)
}

// UpdateReservation updates a Reservation.
func (ipam *Ipam) UpdateReservation(reservation models.Reservation) error {
	session := ipam.session.Copy()
	defer session.Close()

	return session.DB(IpamDatabase).C(IpamCollectionReservations).UpdateId(reservation.ID, reservation)
}

// DeleteReservation removes a Reservation.
func (ipam *Ipam) DeleteReservation(id string) error {
	session := ipam.session.Copy()
	defer session.Close()

	err := session.DB(IpamDatabase).C(IpamCollectionLeases).Update(
		bson.M{"reservation": bson.ObjectIdHex(id)},
		bson.M{"$set": bson.M{"reservation": nil}},
	)
	if err != nil {
		return err
	}

	return session.DB(IpamDatabase).C(IpamCollectionReservations).RemoveId(bson.ObjectIdHex(id))
}

// DeleteReservations remove all reservations in a subnet
func (ipam *Ipam) DeleteReservations(id string) error {
	reservations, err := ipam.GetReservations(id)
	if err != nil {
		return err
	}

	for _, reservation := range reservations {
		err := ipam.DeleteReservation(reservation.ID.Hex())
		if err != nil {
			return err
		}
	}
	return nil
}
