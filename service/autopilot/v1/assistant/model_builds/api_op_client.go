// This is an autogenerated file. DO NOT MODIFY
package model_builds

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client *client.Client

	assistantSid string
}

// The properties required to manage the model builds resources
type ClientProperties struct {
	AssistantSid string
}

// Create a new instance of the client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		assistantSid: properties.AssistantSid,
	}
}
