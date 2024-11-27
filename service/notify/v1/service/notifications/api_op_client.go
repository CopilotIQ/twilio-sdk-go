// Package notifications contains auto-generated files. DO NOT MODIFY
package notifications

import "github.com/CopilotIQ/twilio-sdk-go/client"

// Client for managing service notification resources
// See https://www.twilio.com/docs/notify/api/notification-resource for more details
type Client struct {
	client *client.Client

	serviceSid string
}

// ClientProperties are the properties required to manage the notifications resources
type ClientProperties struct {
	ServiceSid string
}

// New creates a new instance of the notifications client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
	}
}
