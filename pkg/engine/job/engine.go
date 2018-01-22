package job

import (
	"time"
	"crypto-jobs/pkg/models"
	"crypto-jobs/pkg/database"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"os"
)

func StartEngine(config models.Config) () {
	for {
		jobs, err := database.FindStartedJobs()
		if err != nil {
			fmt.Fprintln(os.Stderr, "could not find started jobs: \"%v\"", err)
		}

		fmt.Println(jobs)

		for _, job := range jobs {
			if waitDurationSurpassed(job) {
				go executeJob(job)
			}
		}

		time.Sleep(time.Millisecond * time.Duration(config.SleepTimeInMilliseconds))
	}
}

func waitDurationSurpassed(job models.Job) (bool) {
	fmt.Println(time.Now().Unix())
	return time.Now().Unix()-job.LastExecutionTime > job.FrequencyInSeconds
}

func executeJob(job models.Job) () {
	err := database.ClaimJob(job.Id.Hex())
	if err != nil {
		fmt.Printf("Error claiming job: \"%v\"\n", err)
	}

	fmt.Println("Executing job \"" + job.Id.Hex() + "\"")
	err = database.UpdateJob(job.Id.Hex(), bson.M{"last-execution-time": time.Now().Unix(), "quantity-of-executions": job.QuantityOfExecutions + 1}, bson.M{})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error executing job: \"%v\"", err)
	}

	err = database.UnclaimJob(job.Id.Hex())
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error unclaiming job: \"%v\"", err)
	}
}
