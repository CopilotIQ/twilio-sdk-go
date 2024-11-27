// Package balance contains auto-generated files. DO NOT MODIFY
package balance

import "github.com/CopilotIQ/twilio-sdk-go/client"

// Client for managing a specific account balance resource
type Client struct {
	client *client.Client

	accountSid string
}

// ClientProperties are the properties required to manage the balance resources
type ClientProperties struct {
	AccountSid string
}

// New creates a new instance of the balance client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
	}
}
