package main

import (
	"log"
	"os"

	"github.com/CopilotIQ/twilio-sdk-go"
	v1 "github.com/CopilotIQ/twilio-sdk-go/service/fax/v1"
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
	err := faxClient.
		Fax("FXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Delete()

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Println("Fax deleted")
}
