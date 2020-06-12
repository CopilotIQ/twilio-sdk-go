// This is an autogenerated file. DO NOT MODIFY
package environment

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/environment/variable"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/environment/variables"
)

type Client struct {
	client     *client.Client
	serviceSid string
	sid        string
	Variables  *variables.Client
	Variable   func(string) *variable.Client
}

func New(client *client.Client, serviceSid string, sid string) *Client {
	return &Client{
		client:     client,
		serviceSid: serviceSid,
		sid:        sid,
		Variables:  variables.New(client, sid, serviceSid),
		Variable:   func(variableSid string) *variable.Client { return variable.New(client, sid, serviceSid, variableSid) },
	}
}
