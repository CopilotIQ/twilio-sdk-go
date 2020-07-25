// This is an autogenerated file. DO NOT MODIFY
package actions

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type GetActionResponse struct {
	AccountSid   string      `json:"account_sid"`
	AssistantSid string      `json:"assistant_sid"`
	Data         interface{} `json:"data"`
	TaskSid      string      `json:"task_sid"`
	URL          string      `json:"url"`
}

func (c Client) Get() (*GetActionResponse, error) {
	return c.GetWithContext(context.Background())
}

func (c Client) GetWithContext(context context.Context) (*GetActionResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/Tasks/{taskSid}/Actions",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"taskSid":      c.taskSid,
		},
	}

	response := &GetActionResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
