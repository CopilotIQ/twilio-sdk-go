package main

import (
	"log"

	"github.com/CopilotIQ/twilio-sdk-go/twiml"
)

var twimlClient *twiml.TwiML

func init() {
	twimlClient = twiml.New()
}

func main() {
	response := twimlClient.VoiceResponse()
	connect := response.Connect()
	connect.Autopilot("UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	twiML, err := response.ToTwiML()

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("TwiML: %s", *twiML)
}
