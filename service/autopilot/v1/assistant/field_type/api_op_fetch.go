// This is an autogenerated file. DO NOT MODIFY
package field_type

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchFieldTypeResponse resource/ response properties for the retrieved field type
type FetchFieldTypeResponse struct {
	// The SID of the account which the resource is associated with
	AccountSid string `json:"account_sid"`
	// The SID of the assistant which the field type is associated with
	AssistantSid string `json:"assistant_sid"`
	// The date and time (in RFC3339 format) when the resource was created
	DateCreated time.Time `json:"date_created"`
	// The date and time (in RFC3339 format) when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The human readable name of the field type
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The unique alphanumeric string for the resource
	Sid string `json:"sid"`
	// The URL for the resource
	URL string `json:"url"`
	// The unique human readable name of the field type
	UniqueName string `json:"unique_name"`
}

// Fetch retrieves a field type resource
// See https://www.twilio.com/docs/autopilot/api/field-type#fetch-a-fieldtype-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchFieldTypeResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a field type resource
// See https://www.twilio.com/docs/autopilot/api/field-type#fetch-a-fieldtype-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchFieldTypeResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/FieldTypes/{sid}",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"sid":          c.sid,
		},
	}

	response := &FetchFieldTypeResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
