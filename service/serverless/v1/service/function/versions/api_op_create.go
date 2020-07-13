// This is an autogenerated file. DO NOT MODIFY
package versions

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

type ContentDetails struct {
	Body        io.ReadSeeker `validate:"required" mapstructure:"Body"`
	FileName    string        `validate:"required" mapstructure:"FileName"`
	ContentType string        `validate:"required" mapstructure:"ContentType"`
}

type CreateVersionInput struct {
	Content    ContentDetails `validate:"required" mapstructure:"Content"`
	Path       string         `validate:"required" mapstructure:"Path"`
	Visibility string         `validate:"required" mapstructure:"Visibility"`
}

type CreateVersionResponse struct {
	Sid         string    `json:"sid"`
	AccountSid  string    `json:"account_sid"`
	ServiceSid  string    `json:"service_sid"`
	FunctionSid string    `json:"function_sid"`
	Path        string    `json:"path"`
	Visibility  string    `json:"visibility"`
	DateCreated time.Time `json:"date_created"`
	URL         string    `json:"url"`
}

func (c Client) Create(input *CreateVersionInput) (*CreateVersionResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

func (c Client) CreateWithContext(context context.Context, input *CreateVersionInput) (*CreateVersionResponse, error) {
	op := client.Operation{
		OverrideBaseURL: utils.String(client.CreateBaseURL("serverless-upload", "v1")),
		Method:          http.MethodPost,
		URI:             "/Services/{serviceSid}/Functions/{functionSid}/Versions",
		ContentType:     client.FormData,
		PathParams: map[string]string{
			"serviceSid":  c.serviceSid,
			"functionSid": c.functionSid,
		},
	}

	response := &CreateVersionResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
