// This is an autogenerated file. DO NOT MODIFY
package user

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/user/binding"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/user/channel"
)

type Client struct {
	client *client.Client

	serviceSid string
	sid        string

	// Sub client to manage binding resources
	Binding func(string) *binding.Client
	// Sub client to manage channel resources
	Channel func(string) *channel.Client
}

// The properties required to manage the user resources
type ClientProperties struct {
	ServiceSid string
	Sid        string
}

// Create a new instance of the client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,

		Binding: func(bindingSid string) *binding.Client {
			return binding.New(client, binding.ClientProperties{
				ServiceSid: properties.ServiceSid,
				Sid:        bindingSid,
				UserSid:    properties.Sid,
			})
		},
		Channel: func(channelSid string) *channel.Client {
			return channel.New(client, channel.ClientProperties{
				ServiceSid: properties.ServiceSid,
				Sid:        channelSid,
				UserSid:    properties.Sid,
			})
		},
	}
}
