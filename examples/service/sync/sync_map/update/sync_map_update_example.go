package main

import (
	"log"
	"os"

	"github.com/CopilotIQ/twilio-sdk-go"
	v1 "github.com/CopilotIQ/twilio-sdk-go/service/sync/v1"
	"github.com/CopilotIQ/twilio-sdk-go/service/sync/v1/service/sync_map"
	"github.com/CopilotIQ/twilio-sdk-go/session/credentials"
	"github.com/CopilotIQ/twilio-sdk-go/utils"
)

var syncClient *v1.Sync

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	syncClient = twilio.NewWithCredentials(creds).Sync.V1
}

func main() {
	resp, err := syncClient.
		Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		SyncMap("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Update(&sync_map.UpdateSyncMapInput{
			CollectionTtl: utils.Int(1),
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}
