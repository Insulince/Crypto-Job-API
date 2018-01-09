package handlers

import (
	"net/http"
	"crypto-jobs/pkg/database"
	"crypto-jobs/pkg/models"
)

func GetJobs(w http.ResponseWriter, r *http.Request) () {
	CallReceived(r)

	type Response struct {
	}
	Respond(w, Response{})
}

func CreateJob(w http.ResponseWriter, r *http.Request) () {
	CallReceived(r)

	database.CreateJob(models.Job{WaitDuration: 5})

	type Response struct {
	}
	Respond(w, Response{})
}

func GetJob(w http.ResponseWriter, r *http.Request) () {
	CallReceived(r)

	type Response struct {
	}
	Respond(w, Response{})
}

func UpdateJob(w http.ResponseWriter, r *http.Request) () {
	CallReceived(r)

	type Response struct {
	}
	Respond(w, Response{})
}

func DeleteJob(w http.ResponseWriter, r *http.Request) () {
	CallReceived(r)

	type Response struct {
	}
	Respond(w, Response{})
}

func StartProvidedJobs(w http.ResponseWriter, r *http.Request) () {
	CallReceived(r)

	type Response struct {
	}
	Respond(w, Response{})
}

func StopAllJobs(w http.ResponseWriter, r *http.Request) () {
	CallReceived(r)

	type Response struct {
	}
	Respond(w, Response{})
}

func StartJob(w http.ResponseWriter, r *http.Request) () {
	routeVariables, _, _ := CallReceived(r)
	id := routeVariables["id"]

	database.StartJob(id)

	type Response struct {
	}
	Respond(w, Response{})
}

func StopJob(w http.ResponseWriter, r *http.Request) () {
	routeVariables, _, _ := CallReceived(r)
	id := routeVariables["id"]

	database.StopJob(id)

	type Response struct {
	}
	Respond(w, Response{})
}
