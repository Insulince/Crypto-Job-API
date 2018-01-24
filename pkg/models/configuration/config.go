package configuration

import (
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	TwilioAccountSID                 string `json:"twilio-account-sid"`
	TwilioAuthToken                  string `json:"twilio-auth-token"`
	TwilioAPIURL                     string `json:"twilio-api-url"`
	SenderPhoneNumber                string `json:"sender-phone-number"`
	SenderEmailUsername              string `json:"sender-email-username"`
	SenderEmailPassword              string `json:"sender-email-password"`
	Port                             int    `json:"port"`
	CryptoPriceFetcherPort           int    `json:"crypto-price-fetcher-port"`
	JobEngineSleepTimeInMilliseconds int    `json:"job-engine-sleep-time-in-milliseconds"`
	MongoDBURL                       string `json:"mongo-db-url"`
}

func GetConfiguration() (config Config) {
	jsonFile, err := ioutil.ReadFile("../config/config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonFile, &config)
	if err != nil {
		panic(err)
	}

	return config
}
