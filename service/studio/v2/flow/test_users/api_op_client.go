// This is an autogenerated file. DO NOT MODIFY
package test_users

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client  *client.Client
	flowSid string
}

func New(client *client.Client, flowSid string) *Client {
	return &Client{
		client:  client,
		flowSid: flowSid,
	}
}
