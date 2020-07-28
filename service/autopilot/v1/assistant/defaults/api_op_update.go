// This is an autogenerated file. DO NOT MODIFY
package defaults

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateDefaultInput defines input parameters/ properties for updating defaults
type UpdateDefaultInput struct {
	// Default object as a JSON string
	Defaults *string `form:"Defaults,omitempty"`
}

// UpdateDefaultResponse resource/ response properties for the updated defaults
type UpdateDefaultResponse struct {
	// The SID of the account which the resource is associated with
	AccountSid string `json:"account_sid"`
	// The SID of the assistant which the default is associated with
	AssistantSid string `json:"assistant_sid"`
	// The default JSON
	Data map[string]interface{} `json:"data"`
	// The URL for the resource
	URL string `json:"url"`
}

// Update modifies a defaults resource
// See https://www.twilio.com/docs/autopilot/api/assistant/defaults#update-a-defaults-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateDefaultInput) (*UpdateDefaultResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a defaults resource
// See https://www.twilio.com/docs/autopilot/api/assistant/defaults#update-a-defaults-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateDefaultInput) (*UpdateDefaultResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{assistantSid}/Defaults",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
		},
	}

	response := &UpdateDefaultResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
