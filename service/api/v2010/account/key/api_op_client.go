// This is an autogenerated file. DO NOT MODIFY
package key

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client *client.Client

	accountSid string
	sid        string
}

func New(client *client.Client, accountSid string, sid string) *Client {
	return &Client{
		client: client,

		accountSid: accountSid,
		sid:        sid,
	}
}
