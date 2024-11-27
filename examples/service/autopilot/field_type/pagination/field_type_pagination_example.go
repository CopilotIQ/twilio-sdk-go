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
	paginator := autopilotClient.
		Assistant("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		FieldTypes.
		NewFieldTypesPaginator()

	for paginator.Next() {
		currentPage := paginator.CurrentPage()
		log.Printf("%v field type(s) found on page %v", len(currentPage.FieldTypes), currentPage.Meta.Page)
	}

	if paginator.Error() != nil {
		log.Panicf("%s", paginator.Error())
	}

	log.Printf("Total number of field type(s) found: %v", len(paginator.FieldTypes))
}
