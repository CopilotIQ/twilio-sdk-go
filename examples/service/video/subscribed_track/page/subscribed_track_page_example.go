package main

import (
	"log"
	"os"

	"github.com/CopilotIQ/twilio-sdk-go"
	v1 "github.com/CopilotIQ/twilio-sdk-go/service/video/v1"
	"github.com/CopilotIQ/twilio-sdk-go/service/video/v1/room/participant/subscribed_tracks"
	"github.com/CopilotIQ/twilio-sdk-go/session/credentials"
)

var videoClient *v1.Video

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	videoClient = twilio.NewWithCredentials(creds).Video.V1
}

func main() {
	resp, err := videoClient.
		Room("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Participant("PAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		SubscribedTracks.
		Page(&subscribed_tracks.SubscribedTracksPageOptions{})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("%v subscribed track(s) found on page", len(resp.SubscribedTracks))
}
