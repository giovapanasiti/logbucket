package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type LogPortale struct {
	ID      	bson.ObjectId 	`bson:"_id" json:"id"`
	Payload    	interface{}     `bson:"payload" json:"payload"`
	CreatedAt	time.Time		`bson:"createdAt" json:"createdAt"`
}