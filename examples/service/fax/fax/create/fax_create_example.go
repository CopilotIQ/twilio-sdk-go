package main

import (
	"log"
	"os"

	"github.com/CopilotIQ/twilio-sdk-go"
	v1 "github.com/CopilotIQ/twilio-sdk-go/service/fax/v1"
	"github.com/CopilotIQ/twilio-sdk-go/service/fax/v1/faxes"
	"github.com/CopilotIQ/twilio-sdk-go/session/credentials"
)

var faxClient *v1.Fax

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	faxClient = twilio.NewWithCredentials(creds).Fax.V1
}

func main() {
	resp, err := faxClient.
		Faxes.
		Create(&faxes.CreateFaxInput{
			To:       "+1987654321",
			MediaURL: "http://localhost/media",
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}
