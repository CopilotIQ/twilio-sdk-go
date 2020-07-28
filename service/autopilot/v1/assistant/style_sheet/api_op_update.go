// This is an autogenerated file. DO NOT MODIFY
package style_sheet

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateStyleSheetInput defines input parameters/ properties for updating a stylesheet
type UpdateStyleSheetInput struct {
	// Style sheet object as a JSON string
	StyleSheet *string `form:"StyleSheet,omitempty"`
}

// UpdateStyleSheetResponse resource/ response properties for the updated style sheet
type UpdateStyleSheetResponse struct {
	// The SID of the account which the resource is associated with
	AccountSid string `json:"account_sid"`
	// The SID of the assistant which the style sheet is associated with
	AssistantSid string `json:"assistant_sid"`
	// The style sheet JSON
	Data map[string]interface{} `json:"data"`
	// The URL for the resource
	URL string `json:"url"`
}

// Update modifies a style sheet resource
// See https://www.twilio.com/docs/autopilot/api/assistant/stylesheet#update-a-stylesheet-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateStyleSheetInput) (*UpdateStyleSheetResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a style sheet resource
// See https://www.twilio.com/docs/autopilot/api/assistant/stylesheet#update-a-stylesheet-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateStyleSheetInput) (*UpdateStyleSheetResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{assistantSid}/StyleSheet",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
		},
	}

	response := &UpdateStyleSheetResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
