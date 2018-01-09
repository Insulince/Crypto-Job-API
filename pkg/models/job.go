package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Job struct {
	Id                bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	State             string        `json:"state" bson:"state"`
	WaitDuration      int64         `json:"wait-duration" bson:"wait-duration"`
	LastExecutionTime int64         `json:"last-execution-time" bson:"last-execution-time"`
}
