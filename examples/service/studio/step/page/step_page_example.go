package main

import (
	"log"
	"os"

	"github.com/CopilotIQ/twilio-sdk-go"
	v2 "github.com/CopilotIQ/twilio-sdk-go/service/studio/v2"
	"github.com/CopilotIQ/twilio-sdk-go/service/studio/v2/flow/execution/steps"
	"github.com/CopilotIQ/twilio-sdk-go/session/credentials"
)

var studioClient *v2.Studio

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	studioClient = twilio.NewWithCredentials(creds).Studio.V2
}

func main() {
	resp, err := studioClient.
		Flow("FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Execution("FNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Steps.
		Page(&steps.StepsPageOptions{})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("%v step(s) found on page", len(resp.Steps))
}
