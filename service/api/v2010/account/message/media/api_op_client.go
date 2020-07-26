// This is an autogenerated file. DO NOT MODIFY
package media

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific message media resource
// See https://www.twilio.com/docs/sms/api/media-resource for more details
type Client struct {
	client *client.Client

	accountSid string
	messageSid string
	sid        string
}

// ClientProperties are the properties required to manage the media resources
type ClientProperties struct {
	AccountSid string
	MessageSid string
	Sid        string
}

// New creates a new instance of the media client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		messageSid: properties.MessageSid,
		sid:        properties.Sid,
	}
}
