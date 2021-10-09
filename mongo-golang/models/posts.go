package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// The model
type Posts struct {
	User_Id         bson.ObjectId `json:"userid" bson:"_userid"`
	Id              bson.ObjectId `json:"id" bson:"_id"`
	Caption         string        `json:"caption" bson:"caption"`
	ImageURL        string        `json:"imageurl" bson:"imageurl"`
	PostedTimestamp time.Time     `json:"postedtimestamp" bson:"postedtimestamp"`
}
