package main

import (
	"log"
	"os"

	"github.com/CopilotIQ/twilio-sdk-go"
	v1 "github.com/CopilotIQ/twilio-sdk-go/service/proxy/v1"
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
	err := proxyClient.
		Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Session("KCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Delete()

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Println("Session deleted")
}
