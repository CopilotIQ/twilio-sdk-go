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
	paginator := faxClient.
		Fax("FXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		MediaFiles.
		NewMediaPaginator()

	for paginator.Next() {
		currentPage := paginator.CurrentPage()
		log.Printf("%v media file(s) found on page %v", len(currentPage.Media), currentPage.Meta.Page)
	}

	if paginator.Error() != nil {
		log.Panicf("%s", paginator.Error())
	}

	log.Printf("Total number of media file(s) found: %v", len(paginator.Media))
}
