package main

import (
	"log"
	"os"
	"strings"

	"github.com/CopilotIQ/twilio-sdk-go"
	v1 "github.com/CopilotIQ/twilio-sdk-go/service/serverless/v1"
	"github.com/CopilotIQ/twilio-sdk-go/service/serverless/v1/service/function/versions"
	"github.com/CopilotIQ/twilio-sdk-go/session/credentials"
)

var serverlessClient *v1.Serverless

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	serverlessClient = twilio.NewWithCredentials(creds).Serverless.V1
}

func main() {
	resp, err := serverlessClient.
		Service("ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Function("ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Versions.
		Create(&versions.CreateVersionInput{
			Content: versions.CreateContentDetails{
				Body:        strings.NewReader(`exports.handler = function (context, event, callback) { callback(null, "Hello World"); };`),
				ContentType: "application/javascript",
				FileName:    "test.js",
			},
			Path:       "/test",
			Visibility: "private",
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}
