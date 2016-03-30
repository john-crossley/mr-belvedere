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
	Status    string        `json:"status" form:"status" binding:"required" bson:"status"`
	Author    string        `json:"author" form:"author" binding:"required" bson:"author"`
	CreatedOn int64         `json:"created_on" bson:"created_on"`
	UpdatedOn int64         `json:"updated_on" bson:"updated_on"`
}