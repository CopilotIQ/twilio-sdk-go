package main

import (
	"log"
	"os"

	"github.com/CopilotIQ/twilio-sdk-go"
	v1 "github.com/CopilotIQ/twilio-sdk-go/service/proxy/v1"
	"github.com/CopilotIQ/twilio-sdk-go/service/proxy/v1/service/short_codes"
	"github.com/CopilotIQ/twilio-sdk-go/session/credentials"
)

var proxyClient *v1.Proxy

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	proxyClient = twilio.NewWithCredentials(creds).Proxy.V1
}

func main() {
	resp, err := proxyClient.
		Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		ShortCodes.
		Page(&short_codes.ShortCodesPageOptions{})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("%v short code(s) found on page", len(resp.ShortCodes))
}
