// This is an autogenerated file. DO NOT MODIFY
package step

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow/execution/step/context"
)

type Client struct {
	client *client.Client

	executionSid string
	flowSid      string
	sid          string

	Context func() *context.Client
}

type ClientProperties struct {
	ExecutionSid string
	FlowSid      string
	Sid          string
}

func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		executionSid: properties.ExecutionSid,
		flowSid:      properties.FlowSid,
		sid:          properties.Sid,

		Context: func() *context.Client {
			return context.New(client, context.ClientProperties{
				ExecutionSid: properties.ExecutionSid,
				FlowSid:      properties.FlowSid,
				StepSid:      properties.Sid,
			})
		},
	}
}
