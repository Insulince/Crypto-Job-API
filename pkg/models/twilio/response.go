package twilio

type Response struct {
	ApiVersion          string         `json:"api_version"`
	PriceUnit           string         `json:"price_unit"`
	ErrorCode           string         `json:"error_code"`
	DateCreated         string         `json:"date_created"`
	From                string         `json:"from"`
	MessagingServiceSID string         `json:"messaging_service_sid"`
	Direction           string         `json:"direction"`
	ErrorMessage        string         `json:"error_message"`
	DateUpdated         string         `json:"date_updated"`
	DateSent            string         `json:"date_sent"`
	Status              string         `json:"status"`
	QuantitySegments    string         `json:"num_segments"`
	URI                 string         `json:"uri"`
	SubresourcURIs      SubResourceURI `json:"subresource_uris"`
	SID                 string         `json:"sid"`
	AccountSID          string         `json:"account_sid"`
	To                  string         `json:"to"`
	Body                string         `json:"body"`
	QuantityMedia       string         `json:"num_media"`
	Price               string         `json:"price"`
}
