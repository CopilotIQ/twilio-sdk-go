// This is an autogenerated file. DO NOT MODIFY
package workers

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type CreateWorkerInput struct {
	FriendlyName string  `validate:"required" form:"FriendlyName"`
	ActivitySid  *string `form:"ActivitySid,omitempty"`
	Attributes   *string `form:"Attributes,omitempty"`
}

type CreateWorkerOutput struct {
	Sid              string      `json:"sid"`
	AccountSid       string      `json:"account_sid"`
	WorkspaceSid     string      `json:"workspace_sid"`
	ActivitySid      string      `json:"activity_sid"`
	FriendlyName     string      `json:"friendly_name"`
	ActivityName     string      `json:"activity_name"`
	Attributes       interface{} `json:"attributes"`
	Available        bool        `json:"available"`
	DateCreated      time.Time   `json:"date_created"`
	DateUpdated      *time.Time  `json:"date_updated,omitempty"`
	DateStatusChange *time.Time  `json:"date_status_changed,omitempty"`
	URL              string      `json:"url"`
}

func (c Client) Create(input *CreateWorkerInput) (*CreateWorkerOutput, error) {
	return c.CreateWithContext(context.Background(), input)
}

func (c Client) CreateWithContext(context context.Context, input *CreateWorkerInput) (*CreateWorkerOutput, error) {
	op := client.Operation{
		HTTPMethod:  http.MethodPost,
		HTTPPath:    "/Workspaces/{workspaceSid}/Workers",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
		},
	}

	output := &CreateWorkerOutput{}
	if err := c.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}
