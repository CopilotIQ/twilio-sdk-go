package main

import (
	"log"
	"os"

	"github.com/CopilotIQ/twilio-sdk-go"
	v1 "github.com/CopilotIQ/twilio-sdk-go/service/taskrouter/v1"
	"github.com/CopilotIQ/twilio-sdk-go/session/credentials"
)

var taskrouterClient *v1.TaskRouter

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	taskrouterClient = twilio.NewWithCredentials(creds).TaskRouter.V1
}

func main() {
	paginator := taskrouterClient.
		Workspaces.
		NewWorkspacesPaginator()

	for paginator.Next() {
		currentPage := paginator.CurrentPage()
		log.Printf("%v workspace(s) found on page %v", len(currentPage.Workspaces), currentPage.Meta.Page)
	}

	if paginator.Error() != nil {
		log.Panicf("%s", paginator.Error())
	}

	log.Printf("Total number of workspace(s) found: %v", len(paginator.Workspaces))
}
