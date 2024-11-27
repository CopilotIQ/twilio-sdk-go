package main

import (
	"log"

	"github.com/CopilotIQ/twilio-sdk-go/twiml"
	"github.com/CopilotIQ/twilio-sdk-go/twiml/fax/verbs"
	"github.com/CopilotIQ/twilio-sdk-go/utils"
)

var twimlClient *twiml.TwiML

func init() {
	twimlClient = twiml.New()
}

func main() {
	response := twimlClient.FaxResponse()
	response.ReceiveWithAttributes(verbs.ReceiveAttributes{
		Action: utils.String("http://localhost/action"),
	})
	twiML, err := response.ToTwiML()

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("TwiML: %s", *twiML)
}
