// This is an autogenerated file. DO NOT MODIFY
package field

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client *client.Client

	taskSid      string
	assistantSid string
	sid          string
}

type ClientProperties struct {
	TaskSid      string
	AssistantSid string
	Sid          string
}

func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		taskSid:      properties.TaskSid,
		assistantSid: properties.AssistantSid,
		sid:          properties.Sid,
	}
}