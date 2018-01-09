package handlers

import (
	"net/http"
	"crypto-jobs/pkg/models"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) () {
	CallReceived(r)

	type Response models.Message
	Respond(w, Response{Message: "OK"})
}
