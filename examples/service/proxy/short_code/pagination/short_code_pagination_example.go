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
	paginator := proxyClient.
		Service("KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		ShortCodes.
		NewShortCodesPaginator()

	for paginator.Next() {
		currentPage := paginator.CurrentPage()
		log.Printf("%v short code(s) found on page %v", len(currentPage.ShortCodes), currentPage.Meta.Page)
	}

	if paginator.Error() != nil {
		log.Panicf("%s", paginator.Error())
	}

	log.Printf("Total number of short code(s) found: %v", len(paginator.ShortCodes))
}
