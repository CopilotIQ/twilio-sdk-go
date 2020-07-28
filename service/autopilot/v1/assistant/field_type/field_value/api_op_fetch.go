// This is an autogenerated file. DO NOT MODIFY
package field_value

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// FetchFieldValueResponse resource/ response properties for the retrieved field value
type FetchFieldValueResponse struct {
	// The SID of the account which the resource is associated with
	AccountSid string `json:"account_sid"`
	// The SID of the assistant which the field value is associated with
	AssistantSid string `json:"assistant_sid"`
	// The date and time (in RFC3339 format) when the resource was created
	DateCreated time.Time `json:"date_created"`
	// The date and time (in RFC3339 format) when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The SID of the field type which the field value is associated with
	FieldTypeSid string `json:"field_type_sid"`
	// The ISO language country code for the field value
	Language string `json:"language"`
	// The unique alphanumeric string for the resource
	Sid string `json:"sid"`
	// The SID of the field value which this field value is a synonym of
	SynonymOf *string `json:"synonym_of,omitempty"`
	// The URL for the resource
	URL string `json:"url"`
	// The field value
	Value string `json:"value"`
}

// Fetch retrieves a field value resource
// See https://www.twilio.com/docs/autopilot/api/field-value#fetch-a-fieldvalue-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchFieldValueResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a field value resource
// See https://www.twilio.com/docs/autopilot/api/field-value#fetch-a-fieldvalue-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchFieldValueResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/FieldTypes/{fieldTypeSid}/FieldValues/{sid}",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"fieldTypeSid": c.fieldTypeSid,
			"sid":          c.sid,
		},
	}

	response := &FetchFieldValueResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
