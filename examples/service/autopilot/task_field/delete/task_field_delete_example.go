package main

import (
	"log"
	"os"

	"github.com/CopilotIQ/twilio-sdk-go"
	v1 "github.com/CopilotIQ/twilio-sdk-go/service/autopilot/v1"
	"github.com/CopilotIQ/twilio-sdk-go/session/credentials"
)

var autopilotClient *v1.Autopilot

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	autopilotClient = twilio.NewWithCredentials(creds).Autopilot.V1
}

func main() {
	err := autopilotClient.
		Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Task("UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Field("UEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Delete()

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("Task field successfully deleted")
}
