// This is an autogenerated file. DO NOT MODIFY
package tasks

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type CreateTaskInput struct {
	Actions      *string `form:"Actions,omitempty"`
	ActionsURL   *string `form:"ActionsUrl,omitempty"`
	FriendlyName *string `form:"FriendlyName,omitempty"`
	UniqueName   string  `validate:"required" form:"UniqueName"`
}

type CreateTaskResponse struct {
	AccountSid   string     `json:"account_sid"`
	ActionsURL   string     `json:"actions_url"`
	AssistantSid string     `json:"assistant_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
	UniqueName   string     `json:"unique_name"`
}

func (c Client) Create(input *CreateTaskInput) (*CreateTaskResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

func (c Client) CreateWithContext(context context.Context, input *CreateTaskInput) (*CreateTaskResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{assistantSid}/Tasks",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
		},
	}

	response := &CreateTaskResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
