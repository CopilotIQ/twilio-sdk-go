package main

import (
	"log"
	"os"

	"github.com/CopilotIQ/twilio-sdk-go"
	v2 "github.com/CopilotIQ/twilio-sdk-go/service/studio/v2"
	"github.com/CopilotIQ/twilio-sdk-go/service/studio/v2/flow/executions"
	"github.com/CopilotIQ/twilio-sdk-go/session/credentials"
	"github.com/CopilotIQ/twilio-sdk-go/utils"
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
		Executions.
		Create(&executions.CreateExecutionInput{
			To:         "+18001234567",
			From:       "+15017122661",
			Parameters: utils.String("{\"name\": \"CopilotIQ\"}"),
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}
