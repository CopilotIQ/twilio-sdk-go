package main

import (
	"log"
	"os"

	"github.com/CopilotIQ/twilio-sdk-go"
	v2 "github.com/CopilotIQ/twilio-sdk-go/service/verify/v2"
	"github.com/CopilotIQ/twilio-sdk-go/session/credentials"
)

var verifyClient *v2.Verify

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	verifyClient = twilio.NewWithCredentials(creds).Verify.V2
}

func main() {
	paginator := verifyClient.
		Service("VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Entities.
		NewEntitiesPaginator()

	for paginator.Next() {
		currentPage := paginator.CurrentPage()
		log.Printf("%v entities found on page %v", len(currentPage.Entities), currentPage.Meta.Page)
	}

	if paginator.Error() != nil {
		log.Panicf("%s", paginator.Error())
	}

	log.Printf("Total number of entities found: %v", len(paginator.Entities))
}
