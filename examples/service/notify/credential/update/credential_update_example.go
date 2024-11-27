package main

import (
	"log"
	"os"

	"github.com/google/uuid"

	"github.com/CopilotIQ/twilio-sdk-go"
	v1 "github.com/CopilotIQ/twilio-sdk-go/service/notify/v1"
	"github.com/CopilotIQ/twilio-sdk-go/service/notify/v1/credential"
	sessionCredentials "github.com/CopilotIQ/twilio-sdk-go/session/credentials"
	"github.com/CopilotIQ/twilio-sdk-go/utils"
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
		Credential("CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Update(&credential.UpdateCredentialInput{
			FriendlyName: utils.String(uuid.New().String()),
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}
