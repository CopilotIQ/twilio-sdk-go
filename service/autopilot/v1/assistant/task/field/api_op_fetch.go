// This is an autogenerated file. DO NOT MODIFY
package field

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchFieldResponse resource/ response properties for the retrieved task field
type FetchFieldResponse struct {
	// The SID of the account which the resource is associated with
	AccountSid string `json:"account_sid"`
	// The SID of the assistant which the field is associated with
	AssistantSid string `json:"assistant_sid"`
	// The date and time (in RFC3339 format) when the resource was created
	DateCreated time.Time `json:"date_created"`
	// The date and time (in RFC3339 format) when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The type of field
	FieldType string `json:"field_type"`
	// The unique alphanumeric string for the resource
	Sid string `json:"sid"`
	// The SID of the task which the field is associated with
	TaskSid string `json:"task_sid"`
	// The URL for the resource
	URL string `json:"url"`
	// The unique human readable name of the field
	UniqueName string `json:"unique_name"`
}

// Fetch retrieves a task field resource
// See https://www.twilio.com/docs/autopilot/api/task-field#fetch-a-field-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchFieldResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a task field resource
// See https://www.twilio.com/docs/autopilot/api/task-field#fetch-a-field-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchFieldResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/Tasks/{taskSid}/Fields/{sid}",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"taskSid":      c.taskSid,
			"sid":          c.sid,
		},
	}

	response := &FetchFieldResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
