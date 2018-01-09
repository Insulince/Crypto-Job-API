package handlers

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"encoding/json"
)

func CallReceived(r *http.Request) (routeVariables map[string]string, queryParameters map[string][]string, requestBody []byte) {
	fmt.Printf("Call Received: \"" + r.Method + " " + r.URL.Path + "\"\n")
	return getRequestInformation(r)
}

func getRequestInformation(r *http.Request) (routeVariables map[string]string, queryParameters map[string][]string, requestBody []byte) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	return mux.Vars(r), r.URL.Query(), requestBody
}

func Respond(w http.ResponseWriter, response interface{}) () {
	w.Header().Set("Content-Type", "application/json")
	responseBody, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(responseBody))
}
