package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v2010 "github.com/RJPearson94/twilio-sdk-go/service/api/v2010"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/address"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var apiSession *v2010.V2010

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	apiSession = twilio.NewWithCredentials(creds).API.V2010
}

func main() {
	resp, err := apiSession.
		Account("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Address("ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Update(&address.UpdateAddressInput{
			PostalCode: utils.String("Fake Postal Code"),
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}