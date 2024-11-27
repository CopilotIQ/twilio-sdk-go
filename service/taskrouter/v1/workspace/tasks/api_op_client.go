// Package tasks contains auto-generated files. DO NOT MODIFY
package tasks

import "github.com/CopilotIQ/twilio-sdk-go/client"

// Client for managing task resources
// See https://www.twilio.com/docs/taskrouter/api/task for more details
type Client struct {
	client *client.Client

	workspaceSid string
}

// ClientProperties are the properties required to manage the tasks resources
type ClientProperties struct {
	WorkspaceSid string
}

// New creates a new instance of the tasks client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		workspaceSid: properties.WorkspaceSid,
	}
}
