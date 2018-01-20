package handlers

import (
	"net/http"
	"crypto-jobs/pkg/database"
	"crypto-jobs/pkg/models"
	"fmt"
	"os"
	"crypto-jobs/pkg/models/responses"
)

func GetJobs(w http.ResponseWriter, r *http.Request) () {
	_, _, _, err := ProtectedCallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: err.Error()})
		return
	}

	jobs, err := database.FindJobs()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Error{Message: "Could not fetch jobs"})
		return
	}

	if len(jobs) == 0 {
		jobs = make([]models.Job, 0)
	}

	type Response struct {
		Jobs []models.Job `json:"jobs"`
	}
	Respond(w, Response{Jobs: jobs})
}

func CreateJob(w http.ResponseWriter, r *http.Request) () {
	_, _, _, err := ProtectedCallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: err.Error()})
		return
	}

	err = database.InsertJob(models.Job{FrequencyInSeconds: 5})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Error{Message: "Could not insert new job."})
		return
	}

	Respond(w, responses.Empty{})
}

func GetJob(w http.ResponseWriter, r *http.Request) () {
	_, _, _, err := ProtectedCallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: err.Error()})
		return
	}

	Respond(w, responses.Empty{})
}

func UpdateJob(w http.ResponseWriter, r *http.Request) () {
	_, _, _, err := ProtectedCallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: err.Error()})
		return
	}

	Respond(w, responses.Empty{})
}

func DeleteJob(w http.ResponseWriter, r *http.Request) () {
	_, _, _, err := ProtectedCallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: err.Error()})
		return
	}

	Respond(w, responses.Empty{})
}

func StartProvidedJobs(w http.ResponseWriter, r *http.Request) () {
	_, _, _, err := ProtectedCallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: err.Error()})
		return
	}

	Respond(w, responses.Empty{})
}

func StopAllJobs(w http.ResponseWriter, r *http.Request) () {
	_, _, _, err := ProtectedCallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: err.Error()})
		return
	}

	Respond(w, responses.Empty{})
}

func StartJob(w http.ResponseWriter, r *http.Request) () {
	routeVariables, _, _, err := ProtectedCallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: err.Error()})
		return
	}

	id := routeVariables["id"]

	err = database.StartJob(id)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Error{Message: "Could not start job."})
		return
	}

	Respond(w, responses.Message{Message: "Job started."})
}

func StopJob(w http.ResponseWriter, r *http.Request) () {
	routeVariables, _, _, err := ProtectedCallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: err.Error()})
		return
	}

	id := routeVariables["id"]

	err = database.StopJob(id)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Error{Message: "Could not stop job."})
		return
	}

	Respond(w, responses.Message{Message: "Job stopped."})
}
