package main

import (
	"log"
	"os"

	"github.com/CopilotIQ/twilio-sdk-go"
	v1 "github.com/CopilotIQ/twilio-sdk-go/service/flex/v1"
	"github.com/CopilotIQ/twilio-sdk-go/service/flex/v1/flex_flows"
	"github.com/CopilotIQ/twilio-sdk-go/session/credentials"
	"github.com/CopilotIQ/twilio-sdk-go/utils"
	"github.com/google/uuid"
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
		FlexFlows.
		Create(&flex_flows.CreateFlexFlowInput{
			FriendlyName:    uuid.New().String(),
			ChatServiceSid:  "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			ChannelType:     "web",
			IntegrationType: utils.String("external"),
			Integration: &flex_flows.CreateFlexFlowIntegrationInput{
				URL: utils.String("https://test.com/external"),
			},
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}
