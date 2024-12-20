package main

import (
	"log"
	"os"

	"github.com/CopilotIQ/twilio-sdk-go"
	v2010 "github.com/CopilotIQ/twilio-sdk-go/service/api/v2010"
	sipCredentials "github.com/CopilotIQ/twilio-sdk-go/service/api/v2010/account/sip/credential_list/credentials"
	"github.com/CopilotIQ/twilio-sdk-go/session/credentials"
	"github.com/google/uuid"
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
		Sip.
		CredentialList("CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Credentials.
		Create(&sipCredentials.CreateCredentialInput{
			Username: uuid.New().String()[0:32],
			Password: "Test" + uuid.New().String(),
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}
