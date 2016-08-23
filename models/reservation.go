package models

import "gopkg.in/mgo.v2/bson"

// Reservation is a mgo model representing a collection of Subnet resources.
type Reservation struct {
	ID       bson.ObjectId `bson:"_id"`
	Name     string        `bson:"name"`
	Tags     []string      `bson:"tags"`
	Metadata interface{}   `bson:"metadata"`
	Subnet   bson.ObjectId `bson:"subnet,omitempty"`
}
