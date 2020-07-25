// This is an autogenerated file. DO NOT MODIFY
package field_type

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/field_type/field_value"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/field_type/field_values"
)

type Client struct {
	client *client.Client

	assistantSid string
	sid          string

	FieldValue  func(string) *field_value.Client
	FieldValues *field_values.Client
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

		FieldValue: func(fieldValueSid string) *field_value.Client {
			return field_value.New(client, field_value.ClientProperties{
				AssistantSid: properties.AssistantSid,
				FieldTypeSid: properties.Sid,
				Sid:          fieldValueSid,
			})
		},
		FieldValues: field_values.New(client, field_values.ClientProperties{
			AssistantSid: properties.AssistantSid,
			FieldTypeSid: properties.Sid,
		}),
	}
}
