// This is an autogenerated file. DO NOT MODIFY
package channels

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type CreateChannelInput struct {
	FlexFlowSid          string `validate:"required" form:"FlexFlowSid"`
	Identity             string `validate:"required" form:"Identity"`
	ChatUserFriendlyName string `validate:"required" form:"ChatUserFriendlyName"`
	ChatFriendlyName     string `validate:"required" form:"ChatFriendlyName"`
	Target               string `form:"Target,omitempty"`
	ChatUniqueName       string `form:"ChatUniqueName,omitempty"`
	PreEngagementData    string `form:"PreEngagementData,omitempty"`
	TaskSid              string `form:"TaskSid,omitempty"`
	TaskAttributes       string `form:"TaskAttributes,omitempty"`
	LongLived            *bool  `form:"LongLived,omitempty"`
}

type CreateChannelOutput struct {
	Sid         string     `json:"sid"`
	AccountSid  string     `json:"account_sid"`
	FlexFlowSid string     `json:"flex_flow_sid"`
	TaskSid     *string    `json:"task_sid,omitempty"`
	UserSid     string     `json:"user_sid"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	URL         string     `json:"url"`
}

func (c Client) Create(input *CreateChannelInput) (*CreateChannelOutput, error) {
	return c.CreateWithContext(context.Background(), input)
}

func (c Client) CreateWithContext(context context.Context, input *CreateChannelInput) (*CreateChannelOutput, error) {
	op := client.Operation{
		HTTPMethod:  http.MethodPost,
		HTTPPath:    "/Channels",
		ContentType: client.URLEncoded,
	}

	output := &CreateChannelOutput{}
	if err := c.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}
