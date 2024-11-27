package main

import (
	"log"
	"os"

	"github.com/CopilotIQ/twilio-sdk-go"
	v2 "github.com/CopilotIQ/twilio-sdk-go/service/studio/v2"
	"github.com/CopilotIQ/twilio-sdk-go/service/studio/v2/flows"
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
		Flows.
		Create(&flows.CreateFlowInput{
			FriendlyName: "test",
			Status:       "draft",
			Definition: `
{
	"description": "A New Flow",
	"flags": {
		"allow_concurrent_calls": true
	},
	"initial_state": "Trigger",
	"states": [
		{
			"name": "Trigger",
			"properties": {
				"offset": {
					"x": 0,
					"y": 0
				}
			},
			"transitions": [],
			"type": "trigger"
		}
	]
}`,
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}
