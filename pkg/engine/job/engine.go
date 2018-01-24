package job

import (
	"time"
	"crypto-jobs/pkg/models"
	"crypto-jobs/pkg/database"
	"fmt"
	"os"
	"crypto-jobs/pkg/models/configuration"
)

func RunEngine() () {
	config := configuration.GetConfiguration()

	for {
		startedJobs, err := database.FindStartedJobs()
		if err != nil {
			fmt.Fprintln(os.Stderr, "could not find started jobs: \"%v\"", err)
		}

		fmt.Println(startedJobs)

		for _, startedJob := range startedJobs {
			if startedJob.WaitDurationSurpassed() {
				go ExecuteWrapper(startedJob)
			}
		}

		time.Sleep(time.Millisecond * time.Duration(config.JobEngineSleepTimeInMilliseconds))
	}
}

func ExecuteWrapper(job models.Job) () {
	err := database.ClaimJob(job.Id.Hex())
	if err != nil {
		fmt.Printf("Error claiming job: \"%v\"\n", err)
	}

	job.Execute()

	err = database.JobExecuted(job.Id.Hex(), job.QuantityOfExecutions)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error updating executed job: \"%v\"", err)
	}

	err = database.UnclaimJob(job.Id.Hex())
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error unclaiming job: \"%v\"", err)
	}
}
