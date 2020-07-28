// This is an autogenerated file. DO NOT MODIFY
package model_build

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateModelBuildInput defines input parameters/ properties for updating a model build
type UpdateModelBuildInput struct {
	// The unique human readable name of the model build
	UniqueName *string `form:"UniqueName,omitempty"`
}

// UpdateModelBuildResponse resource/ response properties for the updated model build
type UpdateModelBuildResponse struct {
	// The SID of the account which the resource is associated with
	AccountSid string `json:"account_sid"`
	// The SID of the assistant which the model build is associated with
	AssistantSid string `json:"assistant_sid"`
	// The time (in seconds) taken for the model to finish building
	BuildDuration *int `json:"build_duration,omitempty"`
	// The date and time (in RFC3339 format) when the resource was created
	DateCreated time.Time `json:"date_created"`
	// The date and time (in RFC3339 format) when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The Twilio error code for the issue that occurred
	ErrorCode *int `json:"error_code,omitempty"`
	// The unique alphanumeric string for the resource
	Sid string `json:"sid"`
	// The current status of the build
	Status string `json:"status"`
	// The URL for the resource
	URL string `json:"url"`
	// The unique human readable name of the model build
	UniqueName string `json:"unique_name"`
}

// Update modifies an model build resource
// See https://www.twilio.com/docs/autopilot/api/model-build#update-a-modelbuild-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateModelBuildInput) (*UpdateModelBuildResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies an model build resource
// See https://www.twilio.com/docs/autopilot/api/model-build#update-a-modelbuild-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateModelBuildInput) (*UpdateModelBuildResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{assistantSid}/ModelBuilds/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"sid":          c.sid,
		},
	}

	response := &UpdateModelBuildResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
