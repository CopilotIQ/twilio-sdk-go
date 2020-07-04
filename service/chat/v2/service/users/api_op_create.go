// This is an autogenerated file. DO NOT MODIFY
package users

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type CreateUserInput struct {
	Identity     string  `validate:"required" form:"Identity"`
	RoleSid      *string `form:"RoleSid,omitempty"`
	Attributes   *string `form:"Attributes,omitempty"`
	FriendlyName *string `form:"FriendlyName,omitempty"`
}

type CreateUserOutput struct {
	Sid                 string     `json:"sid"`
	AccountSid          string     `json:"account_sid"`
	ServiceSid          string     `json:"service_sid"`
	RoleSid             string     `json:"role_sid"`
	Identity            string     `json:"identity"`
	Attributes          *string    `json:"attributes,omitempty"`
	FriendlyName        *string    `json:"friendly_name,omitempty"`
	IsOnline            *bool      `json:"is_online,omitempty"`
	IsNotifiable        *bool      `json:"is_notifiable,omitempty"`
	JoinedChannelsCount *int       `json:"joined_channels_count,omitempty"`
	DateCreated         time.Time  `json:"date_created"`
	DateUpdated         *time.Time `json:"date_updated,omitempty"`
	URL                 string     `json:"url"`
}

func (c Client) Create(input *CreateUserInput) (*CreateUserOutput, error) {
	return c.CreateWithContext(context.Background(), input)
}

func (c Client) CreateWithContext(context context.Context, input *CreateUserInput) (*CreateUserOutput, error) {
	op := client.Operation{
		HTTPMethod:  http.MethodPost,
		HTTPPath:    "/Services/{serviceSid}/Users",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	output := &CreateUserOutput{}
	if err := c.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}