package models

import (
	"fmt"
	"os"
	"crypto-jobs/pkg/services"
)

type ActionTypeIndex int
type ActionType string

const (
	emailAction  ActionTypeIndex = iota
	twilioAction ActionTypeIndex = iota
)

var actionTypes = map[ActionTypeIndex]ActionType{
	emailAction:  "email",
	twilioAction: "twilio",
}

type Action struct {
	Name       string     `json:"name" bson:"name"`
	Type       ActionType `json:"type" bson:"type"`
	Value      string     `json:"value" bson:"value"`
	Currencies []string   `json:"currencies" bson:"currencies,omitempty"`
}

func (action *Action) Execute() () {
	switch action.Type {
	case actionTypes[emailAction]:
		fmt.Printf("Sending an email to \"%v\".\n", action.Value)
		services.SendEmailTo(action.Value, "crypto is whack", "ayyyyy this golang, wat up")
	case actionTypes[twilioAction]:
		fmt.Printf("Sending a text to \"%v\".\n", action.Value)
		services.SendTextTo(action.Value, "yo this is golang, wat up")
	default:
		fmt.Fprintln(os.Stderr, "Unrecognized action type \"%v\" encountered.", action.Type)
	}
}
