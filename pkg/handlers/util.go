package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"encoding/json"
	cryptoUsersResponses "crypto-users/pkg/models/responses"
	"fmt"
	"errors"
	"os"
)

func CallReceived(r *http.Request) (routeVariables map[string]string, queryParameters map[string][]string, requestBody []byte, err error) {
	fmt.Printf("Call Received: \"" + r.Method + " " + r.URL.Path + "\"\n")
	return getRequestInformation(r)
}

func ProtectedCallReceived(r *http.Request) (routeVariables map[string]string, queryParameters map[string][]string, requestBody []byte, err error) {
	fmt.Printf("Call Received: \"" + r.Method + " " + r.URL.Path + "\"\n")
	routeVariables, queryParameters, requestBody, err = getRequestInformation(r)
	if err != nil {
		return nil, nil, nil, err
	}

	if len(queryParameters["token-id"]) != 1 {
		return nil, nil, nil, errors.New("No \"token-id\" query parameter provided for authentication, access denied.")
	}

	if verifyToken(queryParameters["token-id"][0]) != true {
		return nil, nil, nil, errors.New("Invalid token id.")
	}

	return routeVariables, queryParameters, requestBody, nil
}

func getRequestInformation(r *http.Request) (routeVariables map[string]string, queryParameters map[string][]string, requestBody []byte, err error) {
	requestBody, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, nil, nil, errors.New("Could not read request body.")
	}
	return mux.Vars(r), r.URL.Query(), requestBody, nil
}

func verifyToken(tokenId string) (valid bool) {
	response, err := http.Get("http://localhost:2576/token/verify?token-id=" + tokenId)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Response error: \"%v\"", err)
		return false
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Response error: \"%v\"", err)
		return false
	}

	var statusMessage cryptoUsersResponses.StatusMessage
	err = json.Unmarshal(responseBody, &statusMessage)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Response error: \"%v\"", err)
		return false
	}

	return statusMessage.Status == "Valid"
}

func Respond(w http.ResponseWriter, response interface{}) () {
	w.Header().Set("Content-Type", "application/json")
	responseBody, err := json.Marshal(response)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintf(w, "{\n\t\"message\": \"Could not process response body.\"\n}")
		return
	}
	fmt.Fprintf(w, string(responseBody))
}
