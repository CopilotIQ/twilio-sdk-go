package main

import (
	"log"
	"os"

	"github.com/CopilotIQ/twilio-sdk-go"
	v1 "github.com/CopilotIQ/twilio-sdk-go/service/conversations/v1"
	"github.com/CopilotIQ/twilio-sdk-go/service/conversations/v1/service/conversation/participants"
	"github.com/CopilotIQ/twilio-sdk-go/session/credentials"
	"github.com/CopilotIQ/twilio-sdk-go/utils"
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
	resp, err := conversationClient.
		Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Conversation("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Participants.
		Create(&participants.CreateParticipantInput{
			Identity: utils.String("test-participant"),
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}
