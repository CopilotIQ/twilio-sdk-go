package main

import (
	"log"
	"os"

	"github.com/CopilotIQ/twilio-sdk-go"
	v1 "github.com/CopilotIQ/twilio-sdk-go/service/flex/v1"
	"github.com/CopilotIQ/twilio-sdk-go/service/flex/v1/channels"
	"github.com/CopilotIQ/twilio-sdk-go/session/credentials"
)

var flexClient *v1.Flex

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	flexClient = twilio.NewWithCredentials(creds).Flex.V1
}

func main() {
	resp, err := flexClient.
		Channels.
		Page(&channels.ChannelsPageOptions{})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("%v channel(s) found on page", len(resp.Channels))
}
