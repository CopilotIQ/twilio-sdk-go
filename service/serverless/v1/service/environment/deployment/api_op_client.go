// This is an autogenerated file. DO NOT MODIFY
package deployment

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client *client.Client

	serviceSid     string
	sid            string
	environmentSid string
}

type ClientProperties struct {
	ServiceSid     string
	Sid            string
	EnvironmentSid string
}

func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid:     properties.ServiceSid,
		sid:            properties.Sid,
		environmentSid: properties.EnvironmentSid,
	}
}
