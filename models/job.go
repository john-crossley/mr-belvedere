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
	Branch    string        `json:"branch" form:"branch" binding:"required" bson:"branch"`
	Version   string        `json:"version" form:"version" binding:"required" bson:"version"`
	Author    string        `json:"author" form:"author" binding:"required" bson:"author"`
	Status    string        `json:"status" form:"status" binding:"required" bson:"status"`
	BuildUrl  string `json:"build_url" form:"build_url" binding:"required" bson:"build_url"`
	BranchUrl string `json:"branch_url" form:"branch_url" binding:"required" bson:"branch_url"`
	CreatedOn int64         `json:"created_on" bson:"created_on"`
	UpdatedOn int64         `json:"updated_on" bson:"updated_on"`
}