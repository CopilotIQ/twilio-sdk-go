// This is an autogenerated file. DO NOT MODIFY
package role

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client     *client.Client
	serviceSid string
	sid        string
}

func New(client *client.Client, serviceSid string, sid string) *Client {
	return &Client{
		client:     client,
		serviceSid: serviceSid,
		sid:        sid,
	}
}
