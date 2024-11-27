package main

import (
	"log"
	"os"

	"github.com/CopilotIQ/twilio-sdk-go"
	v1 "github.com/CopilotIQ/twilio-sdk-go/service/notify/v1"
	"github.com/CopilotIQ/twilio-sdk-go/service/notify/v1/credentials"
	sessionCredentials "github.com/CopilotIQ/twilio-sdk-go/session/credentials"
)

var notifyClient *v1.Notify

func init() {
	creds, err := sessionCredentials.New(sessionCredentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	notifyClient = twilio.NewWithCredentials(creds).Notify.V1
}

func main() {
	resp, err := notifyClient.
		Credentials.
		Page(&credentials.CredentialsPageOptions{})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("%v credential(s) found on page", len(resp.Credentials))
}
