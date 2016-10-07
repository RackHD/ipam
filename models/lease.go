package models

import "gopkg.in/mgo.v2/bson"

// Lease is a mgo model representing a collection of Subnet resources.
type Lease struct {
	ID          bson.ObjectId `bson:"_id"`
	Name        string        `bson:"name"`
	Tags        []string      `bson:"tags"`
	Metadata    interface{}   `bson:"metadata"`
	Subnet      bson.ObjectId `bson:"subnet,omitempty"`
	Reservation bson.ObjectId `bson:"reservation,omitempty"`
	Address     bson.Binary   `bson:"address"`
}
