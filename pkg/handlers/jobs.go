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
		Jobs []models.Job `json:"jobs"`
	}
	Respond(w, Response{Jobs: jobs}, http.StatusOK)
}

func CreateJob(w http.ResponseWriter, r *http.Request) () {
	_, _, _, err := ProtectedCallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	err = database.InsertJob(models.Job{FrequencyInSeconds: 5})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Error{Message: "Could not insert new job."}, http.StatusInternalServerError)
		return
	}

	Respond(w, responses.Empty{}, http.StatusOK)
}

func GetJob(w http.ResponseWriter, r *http.Request) () {
	_, _, _, err := ProtectedCallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	Respond(w, responses.Empty{}, http.StatusOK)
}

func UpdateJob(w http.ResponseWriter, r *http.Request) () {
	_, _, _, err := ProtectedCallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	Respond(w, responses.Empty{}, http.StatusOK)
}

func DeleteJob(w http.ResponseWriter, r *http.Request) () {
	_, _, _, err := ProtectedCallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	Respond(w, responses.Empty{}, http.StatusOK)
}

func StartProvidedJobs(w http.ResponseWriter, r *http.Request) () {
	_, _, _, err := ProtectedCallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	Respond(w, responses.Empty{}, http.StatusOK)
}

func StopAllJobs(w http.ResponseWriter, r *http.Request) () {
	_, _, _, err := ProtectedCallReceived(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		Respond(w, responses.Message{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	Respond(w, responses.Empty{}, http.StatusOK)
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
