package models

type Config struct {
	TwilioAccountSID        string `json:"twilio-account-sid"`
	TwilioAuthToken         string `json:"twilio-auth-token"`
	TwilioAPIURL            string `json:"twilio-api-url"`
	SenderPhoneNumber       string `json:"sender-phone-number"`
	ReceiverPhoneNumber     string `json:"receiver-phone-number"`
	Port                    int    `json:"port"`
	CryptoPriceFetcherPort  int    `json:"crypto-price-fetcher-port"`
	SleepTimeInMilliseconds int    `json:"sleep-time-in-milliseconds"`
	MongoDBURL              string `json:"mongo-db-url"`
}
