package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"fmt"
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

func (job *Job) WaitDurationSurpassed() (bool) {
	return time.Now().Unix()-job.LastExecutionTime > job.FrequencyInSeconds
}

func (job *Job) Execute() () {
	fmt.Println("Executing job \"" + job.Id.Hex() + "\"")

	for _, action := range job.Actions {
		go func(action Action) {
			action.Execute()
		}(action)
	}
}
