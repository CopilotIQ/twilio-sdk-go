package main

import (
	"log"
	"os"

	"github.com/CopilotIQ/twilio-sdk-go"
	v1 "github.com/CopilotIQ/twilio-sdk-go/service/sync/v1"
	"github.com/CopilotIQ/twilio-sdk-go/session/credentials"
)

var syncClient *v1.Sync

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	syncClient = twilio.NewWithCredentials(creds).Sync.V1
}

func main() {
	err := syncClient.
		Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Delete()

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Println("Service Deleted")
}
