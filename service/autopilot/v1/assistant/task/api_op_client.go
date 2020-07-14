// This is an autogenerated file. DO NOT MODIFY
package task

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/task/actions"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/task/statistics"
)

type Client struct {
	client *client.Client

	assistantSid string
	sid          string

	Actions    func() *actions.Client
	Statistics func() *statistics.Client
}

type ClientProperties struct {
	AssistantSid string
	Sid          string
}

func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		assistantSid: properties.AssistantSid,
		sid:          properties.Sid,

		Actions: func() *actions.Client {
			return actions.New(client, actions.ClientProperties{
				TaskSid:      properties.Sid,
				AssistantSid: properties.AssistantSid,
			})
		},
		Statistics: func() *statistics.Client {
			return statistics.New(client, statistics.ClientProperties{
				AssistantSid: properties.AssistantSid,
				TaskSid:      properties.Sid,
			})
		},
	}
}
