// This is an autogenerated file. DO NOT MODIFY
package services

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client *client.Client
}

// Create a new instance of the client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}
