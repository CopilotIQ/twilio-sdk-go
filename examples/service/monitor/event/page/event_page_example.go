package main

import (
	"log"
	"os"

	"github.com/CopilotIQ/twilio-sdk-go"
	v1 "github.com/CopilotIQ/twilio-sdk-go/service/monitor/v1"
	"github.com/CopilotIQ/twilio-sdk-go/service/monitor/v1/events"
	"github.com/CopilotIQ/twilio-sdk-go/session/credentials"
)

var monitorClient *v1.Monitor

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	monitorClient = twilio.NewWithCredentials(creds).Monitor.V1
}

func main() {
	resp, err := monitorClient.
		Events.
		Page(&events.EventsPageOptions{})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("%v event(s) found on page", len(resp.Events))
}
