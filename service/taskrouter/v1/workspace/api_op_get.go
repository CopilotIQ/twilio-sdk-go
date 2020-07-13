// This is an autogenerated file. DO NOT MODIFY
package workspace

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type GetWorkspaceResponse struct {
	Sid                  string     `json:"sid"`
	AccountSid           string     `json:"account_sid"`
	FriendlyName         string     `json:"friendly_name"`
	EventCallbackURL     *string    `json:"event_callback_url,omitempty"`
	EventsFilter         *string    `json:"events_filter,omitempty"`
	DefaultActivityName  string     `json:"default_activity_name"`
	DefaultActivitySid   string     `json:"default_activity_sid"`
	MultiTaskEnabled     bool       `json:"multi_task_enabled"`
	PrioritizeQueueOrder string     `json:"prioritize_queue_order"`
	TimeoutActivityName  string     `json:"timeout_activity_name"`
	TimeoutActivitySid   string     `json:"timeout_activity_sid"`
	DateCreated          time.Time  `json:"date_created"`
	DateUpdated          *time.Time `json:"date_updated,omitempty"`
	URL                  string     `json:"url"`
}

func (c Client) Get() (*GetWorkspaceResponse, error) {
	return c.GetWithContext(context.Background())
}

func (c Client) GetWithContext(context context.Context) (*GetWorkspaceResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Workspaces/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &GetWorkspaceResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
