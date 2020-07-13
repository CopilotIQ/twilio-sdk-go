// This is an autogenerated file. DO NOT MODIFY
package task

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type GetTaskResponse struct {
	Sid                   string      `json:"sid"`
	AccountSid            string      `json:"account_sid"`
	WorkspaceSid          string      `json:"workspace_sid"`
	AssignmentStatus      string      `json:"assignment_status"`
	Age                   int         `json:"age"`
	Attributes            interface{} `json:"attributes"`
	DateCreated           time.Time   `json:"date_created"`
	DateUpdated           *time.Time  `json:"date_updated,omitempty"`
	TaskQueueEnteredDate  *time.Time  `json:"task_queue_entered_date,omitempty"`
	Priority              *int        `json:"priority,omitempty"`
	Reason                *string     `json:"reason,omitempty"`
	TaskQueueSid          *string     `json:"task_queue_sid,omitempty"`
	TaskQueueFriendlyName *string     `json:"task_queue_friendly_name,omitempty"`
	WorkflowSid           *string     `json:"workflow_sid,omitempty"`
	WorkflowFriendlyName  *string     `json:"workflow_friendly_name,omitempty"`
	TaskChannelSid        *string     `json:"task_channel_sid,omitempty"`
	TaskChannelUniqueName *string     `json:"task_channel_unique_name,omitempty"`
	Timeout               int         `json:"timeout"`
	URL                   string      `json:"url"`
}

func (c Client) Get() (*GetTaskResponse, error) {
	return c.GetWithContext(context.Background())
}

func (c Client) GetWithContext(context context.Context) (*GetTaskResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Workspaces/{workspaceSid}/Tasks/{sid}",
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"sid":          c.sid,
		},
	}

	response := &GetTaskResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
