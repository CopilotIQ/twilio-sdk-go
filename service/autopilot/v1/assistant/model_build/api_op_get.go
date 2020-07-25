// This is an autogenerated file. DO NOT MODIFY
package model_build

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type GetModelBuildResponse struct {
	AccountSid    string     `json:"account_sid"`
	AssistantSid  string     `json:"assistant_sid"`
	BuildDuration *int       `json:"build_duration,omitempty"`
	DateCreated   time.Time  `json:"date_created"`
	DateUpdated   *time.Time `json:"date_updated,omitempty"`
	ErrorCode     *int       `json:"error_code,omitempty"`
	Sid           string     `json:"sid"`
	Status        string     `json:"status"`
	URL           string     `json:"url"`
	UniqueName    string     `json:"unique_name"`
}

func (c Client) Get() (*GetModelBuildResponse, error) {
	return c.GetWithContext(context.Background())
}

func (c Client) GetWithContext(context context.Context) (*GetModelBuildResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/ModelBuilds/{sid}",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"sid":          c.sid,
		},
	}

	response := &GetModelBuildResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
