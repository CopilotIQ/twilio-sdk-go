// This is an autogenerated file. DO NOT MODIFY
package webhooks

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client *client.Client

	channelSid string
	serviceSid string
}

// The properties required to manage the webhooks resources
type ClientProperties struct {
	ChannelSid string
	ServiceSid string
}

// Create a new instance of the client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		channelSid: properties.ChannelSid,
		serviceSid: properties.ServiceSid,
	}
}
