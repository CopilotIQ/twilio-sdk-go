// This is an autogenerated file. DO NOT MODIFY
package actions

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type UpdateActionInput struct {
	Actions *string `form:"Actions,omitempty"`
}

type UpdateActionResponse struct {
	AccountSid   string      `json:"account_sid"`
	AssistantSid string      `json:"assistant_sid"`
	TaskSid      string      `json:"task_sid"`
	Data         interface{} `json:"data"`
	URL          string      `json:"url"`
}

func (c Client) Update(input *UpdateActionInput) (*UpdateActionResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

func (c Client) UpdateWithContext(context context.Context, input *UpdateActionInput) (*UpdateActionResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{assistantSid}/Tasks/{taskSid}/Actions",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"taskSid":      c.taskSid,
		},
	}

	response := &UpdateActionResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
