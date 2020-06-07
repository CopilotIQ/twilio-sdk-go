// This is an autogenerated file. DO NOT MODIFY
package workspaces

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type CreateWorkspaceInput struct {
	FriendlyName         string `validate:"required" form:"FriendlyName"`
	EventCallbackUrl     string `form:"EventCallbackUrl,omitempty"`
	EventsFilter         string `form:"EventsFilter,omitempty"`
	MultiTaskEnabled     bool   `form:"MultiTaskEnabled,omitempty"`
	Template             bool   `form:"Template,omitempty"`
	PrioritizeQueueOrder bool   `form:"PrioritizeQueueOrder,omitempty"`
}

type CreateWorkspaceOutput struct {
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

func (c Client) Create(input *CreateWorkspaceInput) (*CreateWorkspaceOutput, error) {
	return c.CreateWithContext(context.Background(), input)
}

func (c Client) CreateWithContext(context context.Context, input *CreateWorkspaceInput) (*CreateWorkspaceOutput, error) {
	op := client.Operation{
		HTTPMethod:  http.MethodPost,
		HTTPPath:    "/Workspaces",
		ContentType: client.URLEncoded,
	}

	output := &CreateWorkspaceOutput{}
	if err := c.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}