package main

import (
	"log"
	"os"

	"github.com/CopilotIQ/twilio-sdk-go"
	v2010 "github.com/CopilotIQ/twilio-sdk-go/service/api/v2010"
	"github.com/CopilotIQ/twilio-sdk-go/service/api/v2010/account/conference/participant"
	"github.com/CopilotIQ/twilio-sdk-go/session/credentials"
	"github.com/CopilotIQ/twilio-sdk-go/utils"
)

var apiClient *v2010.V2010

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	apiClient = twilio.NewWithCredentials(creds).API.V2010
}

func main() {
	resp, err := apiClient.
		Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Conference("CFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Participant("CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Update(&participant.UpdateParticipantInput{
			Muted: utils.Bool(true),
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("Call SID: %s", resp.CallSid)
}
