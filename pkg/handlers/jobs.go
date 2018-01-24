package handlers

import (
	"net/http"
	"crypto-jobs/pkg/database"
	"crypto-jobs/pkg/models"
	"fmt"
	"os"
	"crypto-jobs/pkg/models/responses"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
)

func GetJobs(w http.ResponseWriter, r *http.Request) () {
	_, _, _, err := ProtectedCallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	jobs, err := database.FindJobs()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Error{Message: "Could not fetch jobs"}, http.StatusInternalServerError)
		return
	}

	if len(jobs) == 0 {
		jobs = make([]models.Job, 0)
	}

	type Response struct {
		Message string       `json:"message"`
		Jobs    []models.Job `json:"jobs"`
	}
	Respond(w, Response{Message: "Success.", Jobs: jobs}, http.StatusOK)
}

func CreateJob(w http.ResponseWriter, r *http.Request) () {
	_, _, rawPostBody, err := ProtectedCallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	type PostBody struct {
		Actions            []models.Action `json:"actions"`
		FrequencyInSeconds int64           `json:"frequency-in-seconds"`
		Started            bool            `json:"started"`
	}
	var postBody PostBody
	err = json.Unmarshal(rawPostBody, &postBody)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Error{Message: "Could not read request body."}, http.StatusBadRequest)
		return
	}

	job := models.Job{Actions: postBody.Actions, FrequencyInSeconds: postBody.FrequencyInSeconds}

	err = database.InsertJob(job, postBody.Started)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Error{Message: "Could not insert new job."}, http.StatusInternalServerError)
		return
	}

	Respond(w, responses.Message{Message: "Success."}, http.StatusOK)
}

func StartProvidedJobs(w http.ResponseWriter, r *http.Request) () {
	_, _, rawPostBody, err := ProtectedCallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	type PostBody struct {
		Ids []string `json:"ids"`
	}
	var postBody PostBody
	err = json.Unmarshal(rawPostBody, &postBody)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Error{Message: "Could not read request body."}, http.StatusBadRequest)
		return
	}

	for _, id := range postBody.Ids {
		err = database.StartJob(id)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			Respond(w, responses.Error{Message: "Could not start all jobs."}, http.StatusInternalServerError)
			return
		}
	}

	Respond(w, responses.Message{Message: "Success"}, http.StatusOK)
}

func StopAllJobs(w http.ResponseWriter, r *http.Request) () {
	_, _, _, err := ProtectedCallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	startedJobs, err := database.FindStartedJobs()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: "Could not fetch jobs."}, http.StatusInternalServerError)
		return
	}
	for _, startedJob := range startedJobs {
		err = database.StopJob(startedJob.Id.Hex())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			Respond(w, responses.Error{Message: "Could not stop all jobs."}, http.StatusInternalServerError)
			return
		}
	}

	Respond(w, responses.Message{Message: "Success."}, http.StatusOK)
}

func GetJob(w http.ResponseWriter, r *http.Request) () {
	routeVariables, _, _, err := ProtectedCallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	id := routeVariables["id"]

	job, err := database.FindJobByID(id)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Error{Message: "Could not find job."}, http.StatusInternalServerError)
		return
	}

	Respond(w, *job, http.StatusOK)
}

func UpdateJob(w http.ResponseWriter, r *http.Request) () {
	routeVariables, _, rawPutBody, err := ProtectedCallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	id := routeVariables["id"]

	job, err := database.FindJobByID(id)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Error{Message: "Could not find job to update."}, http.StatusBadRequest)
		return
	}

	type PutBody struct {
		Actions struct {
			AddToOrChange []models.Action `json:"add-to-or-change"`
			Remove        []string        `json:"remove"`
		} `json:"actions"`
		FrequencyInSeconds int64 `json:"frequency-in-seconds"`
	}
	var putBody PutBody
	err = json.Unmarshal(rawPutBody, &putBody)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Error{Message: "Could not read request body."}, http.StatusBadRequest)
		return
	}

	if len(putBody.Actions.AddToOrChange) == 0 && len(putBody.Actions.Remove) == 0 && putBody.FrequencyInSeconds == 0 {
		fmt.Fprintln(os.Stderr, "No updates provided.")
		Respond(w, responses.Error{Message: "No updates provided."}, http.StatusBadRequest)
		return
	}

	var removalIndices []int
	for _, removeActionName := range putBody.Actions.Remove {
		found := false

		for i, action := range job.Actions {
			if action.Name == removeActionName {
				removalIndices = append(removalIndices, i)
				found = true
				break
			}
		}

		if !found {
			fmt.Fprintf(os.Stderr, "Action \"%v\" was not found in requested job.\n", removeActionName)
			Respond(w, responses.Error{Message: "Could not remove action \"" + removeActionName + "\" from job, it is not in this job to begin with."}, http.StatusBadRequest)
			return
		}
	}

	for i := len(removalIndices) - 1; i > -1; i-- {
		job.Actions = append(job.Actions[:removalIndices[i]], job.Actions[removalIndices[i]+1:]...)
	}

	for _, addAction := range putBody.Actions.AddToOrChange {
		for _, action := range job.Actions {
			if action.Name == addAction.Name {
				fmt.Fprintf(os.Stderr, "Action \"%v\" already in requested job.\n", addAction.Name)
				Respond(w, responses.Error{Message: "Could not add action \"" + addAction.Name + "\" to job, it is already in this job. If you are trying to change this job you will need to remove it first then add it (this can be done at the same time)."}, http.StatusBadRequest)
				return
			}
		}
	}

	job.Actions = append(job.Actions, putBody.Actions.AddToOrChange...)

	changes := bson.M{}
	changes["actions"] = job.Actions

	if putBody.FrequencyInSeconds > 0 {
		changes["frequency-in-seconds"] = putBody.FrequencyInSeconds
	} else if putBody.FrequencyInSeconds == 0 {
		changes["frequency-in-seconds"] = job.FrequencyInSeconds
	} else {
		fmt.Fprintln(os.Stderr, "Frequency must be greater than 0")
		Respond(w, responses.Error{Message: "Frequency must be greater than 0"}, http.StatusBadRequest)
		return
	}

	err = database.UpdateJob(job.Id.Hex(), changes, bson.M{})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Error{Message: "Could not update job."}, http.StatusInternalServerError)
		return
	}

	Respond(w, responses.Message{Message: "Success."}, http.StatusOK)
}

func DeleteJob(w http.ResponseWriter, r *http.Request) () {
	routeVariables, _, _, err := ProtectedCallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	id := routeVariables["id"]

	err = database.DeleteJob(id)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Error{Message: "Could not find job to delete."}, http.StatusBadRequest)
		return
	}

	Respond(w, responses.Message{Message: "Success."}, http.StatusOK)
}

func StartJob(w http.ResponseWriter, r *http.Request) () {
	routeVariables, _, _, err := ProtectedCallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	id := routeVariables["id"]

	err = database.StartJob(id)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Error{Message: "Could not start job."}, http.StatusInternalServerError)
		return
	}

	Respond(w, responses.Message{Message: "Job started."}, http.StatusOK)
}

func StopJob(w http.ResponseWriter, r *http.Request) () {
	routeVariables, _, _, err := ProtectedCallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	id := routeVariables["id"]

	err = database.StopJob(id)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Error{Message: "Could not stop job."}, http.StatusInternalServerError)
		return
	}

	Respond(w, responses.Message{Message: "Job stopped."}, http.StatusOK)
}
