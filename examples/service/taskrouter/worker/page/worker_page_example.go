package main

import (
	"log"
	"os"

	"github.com/CopilotIQ/twilio-sdk-go"
	v1 "github.com/CopilotIQ/twilio-sdk-go/service/taskrouter/v1"
	"github.com/CopilotIQ/twilio-sdk-go/service/taskrouter/v1/workspace/workers"
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
	resp, err := taskrouterClient.
		Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Workers.
		Page(&workers.WorkersPageOptions{})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("%v worker(s) found on page", len(resp.Workers))
}
