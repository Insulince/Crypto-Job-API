package handlers

import (
	"net/http"
	"crypto-jobs/pkg/models"
)

func Home(w http.ResponseWriter, r *http.Request) () {
	CallReceived(r)

	type Response models.Message
	Respond(w, Response{Message: "Welcome!"})
}
