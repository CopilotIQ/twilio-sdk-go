package main

import (
	"log"
	"os"

	"github.com/CopilotIQ/twilio-sdk-go"
	v1 "github.com/CopilotIQ/twilio-sdk-go/service/conversations/v1"
	"github.com/CopilotIQ/twilio-sdk-go/session/credentials"
)

var conversationClient *v1.Conversations

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	conversationClient = twilio.NewWithCredentials(creds).Conversations.V1
}

func main() {
	paginator := conversationClient.
		Roles.
		NewRolesPaginator()

	for paginator.Next() {
		currentPage := paginator.CurrentPage()
		log.Printf("%v role(s) found on page %v", len(currentPage.Roles), currentPage.Meta.Page)
	}

	if paginator.Error() != nil {
		log.Panicf("%s", paginator.Error())
	}

	log.Printf("Total number of role(s) found: %v", len(paginator.Roles))
}
