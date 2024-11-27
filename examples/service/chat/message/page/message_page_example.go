package main

import (
	"log"
	"os"

	"github.com/CopilotIQ/twilio-sdk-go"
	v2 "github.com/CopilotIQ/twilio-sdk-go/service/chat/v2"
	"github.com/CopilotIQ/twilio-sdk-go/service/chat/v2/service/channel/messages"
	"github.com/CopilotIQ/twilio-sdk-go/session/credentials"
)

var chatClient *v2.Chat

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	chatClient = twilio.NewWithCredentials(creds).Chat.V2
}

func main() {
	resp, err := chatClient.
		Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Channel("CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Messages.
		Page(&messages.ChannelMessagesPageOptions{})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("%v message(s) found on page", len(resp.Messages))
}
