package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/fax/v1"
	"github.com/RJPearson94/twilio-sdk-go/service/fax/v1/faxes"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var faxSession *v1.Fax

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	faxSession = twilio.NewWithCredentials(creds).Fax.V1
}

func main() {
	resp, err := faxSession.
		Faxes.
		Create(&faxes.CreateFaxInput{
			To:       "+1987654321",
			MediaURL: "http://localhost/media",
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}