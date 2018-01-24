package database

import (
	"gopkg.in/mgo.v2"
	"crypto-jobs/pkg/models"
	"gopkg.in/mgo.v2/bson"
	"errors"
	"time"
	"crypto-jobs/pkg/models/configuration"
)

var db *mgo.Database

func InitializeDatabase(config configuration.Config) (err error) {
	session, err := mgo.Dial(config.MongoDBURL)
	if err != nil {
		return err
	}
	session.SetMode(mgo.Strong, true)
	db = session.DB("crypto")
	return nil
}

func Jobs() (jobs *mgo.Collection) {
	return db.C("jobs")
}

func InsertJob(job models.Job, started bool) (err error) {
	job.LastExecutionTime = time.Now().Unix()
	if started {
		job.State = "started"
	} else {
		job.State = "stopped"
	}
	job.Claimed = false
	job.QuantityOfExecutions = 0

	return Jobs().Insert(job)
}

func FindJobs() (jobs []models.Job, err error) {
	err = Jobs().Find(nil).All(&jobs)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func FindStartedJobs() (jobs []models.Job, err error) {
	err = Jobs().Find(bson.M{"state": "started"}).All(&jobs)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func FindJobByID(id string) (job *models.Job, err error) {
	if !bson.IsObjectIdHex(id) {
		return nil, errors.New("Provided ID \"" + id + "\" is not a valid MongoDB ID.")
	}

	err = Jobs().FindId(bson.ObjectIdHex(id)).One(&job)
	if err != nil {
		return nil, err
	}
	return job, nil
}

func UpdateJob(id string, additionsOrChanges bson.M, removals bson.M) (err error) {
	if !bson.IsObjectIdHex(id) {
		return errors.New("Provided ID \"" + id + "\" is not a valid MongoDB ID.")
	}

	if len(additionsOrChanges) > 0 {
		if len(removals) > 0 {
			return Jobs().UpdateId(bson.ObjectIdHex(id), bson.M{"$set": additionsOrChanges, "$unset": removals})
		} else {
			return Jobs().UpdateId(bson.ObjectIdHex(id), bson.M{"$set": additionsOrChanges})
		}
	} else {
		if len(removals) > 0 {
			return Jobs().UpdateId(bson.ObjectIdHex(id), bson.M{"$unset": removals})
		} else {
			return errors.New("No changes provided to update job with!")
		}
	}
}

func DeleteJob(id string) (err error) {
	if !bson.IsObjectIdHex(id) {
		return errors.New("Provided ID \"" + id + "\" is not a valid MongoDB ID.")
	}

	return Jobs().RemoveId(bson.ObjectIdHex(id))
}

func StartJob(id string) (err error) {
	return UpdateJob(id, bson.M{"state": "started"}, bson.M{})
}

func StopJob(id string) (err error) {
	return UpdateJob(id, bson.M{"state": "stopped"}, bson.M{})
}

func JobExecuted(id string, quantityOfExecutions int64) (err error) {
	return UpdateJob(id, bson.M{"last-execution-time": time.Now().Unix(), "quantity-of-executions": quantityOfExecutions + 1}, bson.M{})
}

func ClaimJob(id string) (err error) {
	job, err := FindJobByID(id)
	if err != nil {
		return err
	}

	if job.Claimed != false {
		return errors.New("job is already claimed")
	}

	return UpdateJob(id, bson.M{"claimed": true}, bson.M{})
}

func UnclaimJob(id string) (err error) {
	job, err := FindJobByID(id)
	if err != nil {
		return err
	}

	if job.Claimed != true {
		return errors.New("job is already unclaimed")
	}

	return UpdateJob(id, bson.M{"claimed": false}, bson.M{})
}
