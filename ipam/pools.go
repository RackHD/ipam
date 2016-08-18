package ipam

import (
	"github.com/RackHD/ipam/models"
	"gopkg.in/mgo.v2/bson"
)

// IpamCollectionPools is the name of the Mongo collection which stores Pools.
const IpamCollectionPools string = "pools"

// GetPools returns a list of Pools.
func (ipam *Ipam) GetPools() ([]models.Pool, error) {
	session := ipam.session.Copy()
	defer session.Close()

	var pools []models.Pool

	session.DB(IpamDatabase).C(IpamCollectionPools).Find(nil).All(&pools)

	return pools, nil
}

// GetPool returns the requested Pool.
func (ipam *Ipam) GetPool(id string) (models.Pool, error) {
	session := ipam.session.Copy()
	defer session.Close()

	var pool models.Pool

	return pool, session.DB(IpamDatabase).C(IpamCollectionPools).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&pool)
}

// CreatePool creates a Pool.
func (ipam *Ipam) CreatePool(pool models.Pool) error {
	session := ipam.session.Copy()
	defer session.Close()

	return session.DB(IpamDatabase).C(IpamCollectionPools).Insert(pool)
}

// UpdatePool updates a Pool.
func (ipam *Ipam) UpdatePool(pool models.Pool) error {
	session := ipam.session.Copy()
	defer session.Close()

	return session.DB(IpamDatabase).C(IpamCollectionPools).UpdateId(pool.ID, pool)
}

// DeletePool removes a Pool.
func (ipam *Ipam) DeletePool(id string) error {
	session := ipam.session.Copy()
	defer session.Close()

	return session.DB(IpamDatabase).C(IpamCollectionPools).RemoveId(id)
}
