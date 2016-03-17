package models

import (
	"gopkg.in/mgo.v2/bson"
)

const (
// CollectionJob holds the name of the jobs collection
	CollectionJob = "jobs"
)

type Job struct {
	Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string        `json:"name" form:"name" binding:"required" bson:"name"`
	Success   bool          `json:"success" form:"success" binding:"required" bson:"success"`
	CreatedOn int64         `json:"created_on" bson:"created_on"`
	UpdatedOn int64         `json:"updated_on" bson:"updated_on"`
}