package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Job struct {
	Id                   bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	State                string        `json:"state" bson:"state"`
	Claimed              bool          `json:"claimed" bson:"claimed"`
	Actions              []Action      `json:"actions" bson:"actions"`
	FrequencyInSeconds   int64         `json:"frequency-in-seconds" bson:"frequency-in-seconds"`
	LastExecutionTime    int64         `json:"last-execution-time" bson:"last-execution-time"`
	CreationTimestamp    int64         `json:"creation-timestamp" bson:"creation-timestamp"`
	CreatedBy            string        `json:"created-by" bson:"created-by"`
	QuantityOfExecutions int64         `json:"quantity-of-executions" bson:"quantity-of-executions"`
}

type Action struct {
	Type       string `json:"type" bson:"type"`
	For        string `json:"for" bson:"for"`
	Currencies string `json:"currencies" bson:"currencies,omitempty"`
	Via        string `json:"via" bson:"via"`
}
