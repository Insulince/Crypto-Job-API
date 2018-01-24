package services

import (
	"net/url"
	"net/http"
	"strings"
	"encoding/json"
	"fmt"
	"crypto-jobs/pkg/models/twilio"
	"crypto-jobs/pkg/models/configuration"
	"os"
)

func SendTextTo(receiverPhoneNumber string, message string) () {
	config := configuration.GetConfiguration()

	bodyData := url.Values{}
	bodyData.Set("To", receiverPhoneNumber)
	bodyData.Set("From", config.SenderPhoneNumber)
	bodyData.Set("Body", message)
	httpClient := &http.Client{}
	request, err := http.NewRequest("POST", config.TwilioAPIURL, strings.NewReader(bodyData.Encode()))
	if err != nil {
		panic(err)
	}
	request.SetBasicAuth(config.TwilioAccountSID, config.TwilioAuthToken)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Close = true

	response, err := httpClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		var twilioResponse twilio.Response
		err := json.NewDecoder(response.Body).Decode(&twilioResponse)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Twilio Response Error: \"%v\"\n", response.Status)
	}
}
