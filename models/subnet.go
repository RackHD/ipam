package models

import "gopkg.in/mgo.v2/bson"

// Subnet is a mgo model representing a collection of Subnet resources.
type Subnet struct {
	ID       bson.ObjectId `bson:"_id"`
	Name     string        `bson:"name"`
	Tags     []string      `bson:"tags"`
	Metadata interface{}   `bson:"metadata"`
	Pool     bson.ObjectId `bson:"pool,omitempty"`
}
