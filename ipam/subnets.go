package ipam

import (
	"github.com/RackHD/ipam/models"
	"gopkg.in/mgo.v2/bson"
)

// IpamCollectionSubnets is the name of the Mongo collection which stores Subnets.
const IpamCollectionSubnets string = "subnets"

// GetSubnets returns a list of Subnets.
func (ipam *Ipam) GetSubnets(id string) ([]models.Subnet, error) {
	session := ipam.session.Copy()
	defer session.Close()

	var subnets []models.Subnet

	session.DB(IpamDatabase).C(IpamCollectionSubnets).Find(bson.M{"subnet": bson.ObjectIdHex(id)}).All(&subnets)

	return subnets, nil
}

// GetSubnet returns the requested Subnet.
func (ipam *Ipam) GetSubnet(id string) (models.Subnet, error) {
	session := ipam.session.Copy()
	defer session.Close()

	var subnet models.Subnet

	return subnet, session.DB(IpamDatabase).C(IpamCollectionSubnets).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&subnet)
}

// CreateSubnet creates a Subnet.
func (ipam *Ipam) CreateSubnet(subnet models.Subnet) error {
	session := ipam.session.Copy()
	defer session.Close()

	return session.DB(IpamDatabase).C(IpamCollectionSubnets).Insert(subnet)
}

// UpdateSubnet updates a Subnet.
func (ipam *Ipam) UpdateSubnet(subnet models.Subnet) error {
	session := ipam.session.Copy()
	defer session.Close()

	return session.DB(IpamDatabase).C(IpamCollectionSubnets).UpdateId(subnet.ID, subnet)
}

// DeleteSubnet removes a Subnet.
func (ipam *Ipam) DeleteSubnet(id string) error {
	session := ipam.session.Copy()
	defer session.Close()

	return session.DB(IpamDatabase).C(IpamCollectionSubnets).RemoveId(id)
}
