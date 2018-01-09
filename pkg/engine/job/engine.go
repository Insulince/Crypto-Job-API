package job

import (
	"time"
	"crypto-jobs/pkg/models"
	"crypto-jobs/pkg/database"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

func StartEngine(config models.Config) () {
	for {
		jobs := database.FindStartedJobs()
		fmt.Println(jobs)
		for _, job := range jobs {
			if waitDurationSurpassed(job) {
				executeJob(job)
			}
		}
		time.Sleep(time.Millisecond * time.Duration(config.SleepTimeInMilliseconds))
	}
}

func waitDurationSurpassed(job models.Job) (bool) {
	fmt.Println(time.Now().Unix())
	return time.Now().Unix()-job.LastExecutionTime > job.WaitDuration
}

func executeJob(job models.Job) () {
	fmt.Println("Executing job \"" + job.Id.Hex() + "\"")
	database.UpdateJob(job.Id.Hex(), bson.M{"last-execution-time": time.Now().Unix()})
}
