package ipam

import (
	"encoding/binary"
	"fmt"

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

	session.DB(IpamDatabase).C(IpamCollectionSubnets).Find(bson.M{"pool": bson.ObjectIdHex(id)}).All(&subnets)

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

	err := session.DB(IpamDatabase).C(IpamCollectionSubnets).Insert(subnet)
	if err != nil {
		return err
	}

	// Convert byte arrays to integers (IPv4 only).
	start := binary.BigEndian.Uint32(subnet.Start.Data[len(subnet.Start.Data)-4:])
	end := binary.BigEndian.Uint32(subnet.End.Data[len(subnet.End.Data)-4:])

	// Iterate through the range of IP's and insert a record for each.
	for ; start < end; start++ {
		// IP's are stored as 16 byte arrays and we're only doing IPv4 so prepend
		// the net.IP prefix that denotes an IPv4 address.
		prefix := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff}
		address := make([]byte, 4)

		binary.BigEndian.PutUint32(address, start)

		// Create the lease record, tie it to the subnet.
		lease := models.Lease{
			ID:     bson.NewObjectId(),
			Subnet: subnet.ID,
			Address: bson.Binary{
				Kind: 0,
				Data: append(prefix, address...),
			},
		}

		// Insert, though if we fail one we will have a partially populated subnet pool.
		err := session.DB(IpamDatabase).C(IpamCollectionLeases).Insert(lease)
		if err != nil {
			return err
		}
	}

	return nil
}

// UpdateSubnet updates a Subnet.
func (ipam *Ipam) UpdateSubnet(subnet models.Subnet) error {
	return fmt.Errorf("UpdateSubnet Temporarily Disabled.")

	// session := ipam.session.Copy()
	// defer session.Close()
	//
	// return session.DB(IpamDatabase).C(IpamCollectionSubnets).UpdateId(subnet.ID, subnet)
}

// DeleteSubnet removes a Subnet.
func (ipam *Ipam) DeleteSubnet(id string) error {
	session := ipam.session.Copy()
	defer session.Close()

	return session.DB(IpamDatabase).C(IpamCollectionSubnets).RemoveId(bson.ObjectIdHex(id))
}
