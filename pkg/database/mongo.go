package database

import (
	"gopkg.in/mgo.v2"
	"crypto-jobs/pkg/models"
	"gopkg.in/mgo.v2/bson"
	"errors"
	"time"
)

var db *mgo.Database

func InitializeDatabase(config models.Config) () {
	session, err := mgo.Dial(config.MongoDBURL)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Strong, true)
	db = session.DB("crypto")
}

func Jobs() (*mgo.Collection) {
	return db.C("jobs")
}

func CreateJob(job models.Job) () {
	job.LastExecutionTime = time.Now().Unix()
	job.State = "stopped"

	err := Jobs().Insert(job)
	if err != nil {
		panic(err)
	}
}

func FindJobs() ([]models.Job) {
	var jobs []models.Job
	err := Jobs().Find(nil).All(&jobs)
	if err != nil {
		panic(err)
	}
	return jobs
}

func FindStartedJobs() ([]models.Job) {
	var jobs []models.Job
	err := Jobs().Find(bson.M{"state": "started"}).All(&jobs)
	if err != nil {
		panic(err)
	}
	return jobs
}

func FindJobByID(id string) (models.Job) {
	if !bson.IsObjectIdHex(id) {
		panic(errors.New("Provided ID \"" + id + "\" is not a valid MongoDB ID."))
	}
	var job models.Job
	err := Jobs().FindId(bson.ObjectIdHex(id)).One(&job)
	if err != nil {
		panic(err)
	}
	return job
}

func UpdateJob(id string, updates bson.M) () {
	if !bson.IsObjectIdHex(id) {
		panic(errors.New("Provided ID \"" + id + "\" is not a valid MongoDB ID."))
	}
	err := Jobs().UpdateId(bson.ObjectIdHex(id), bson.M{"$set": updates})
	if err != nil {
		panic(err)
	}
}

func DeleteJob(id string) () {
	if !bson.IsObjectIdHex(id) {
		panic(errors.New("Provided ID \"" + id + "\" is not a valid MongoDB ID."))
	}
	err := Jobs().RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		panic(err)
	}
}

func StartJob(id string) () {
	UpdateJob(id, bson.M{"state": "started"})
}

func StopJob(id string) () {
	UpdateJob(id, bson.M{"state": "stopped"})
}
