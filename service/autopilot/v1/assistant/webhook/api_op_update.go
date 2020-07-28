// This is an autogenerated file. DO NOT MODIFY
package webhook

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// UpdateWebhookInput defines input parameters/ properties for updating a webhook
type UpdateWebhookInput struct {
	// Space separated string of event which triggers the webhook
	Events *string `form:"Events,omitempty"`
	// The unique human readable name of the webhook
	UniqueName *string `form:"UniqueName,omitempty"`
	// The HTTP method used to call the webhook URL
	WebhookMethod *string `form:"WebhookMethod,omitempty"`
	// The URL of the webhook
	WebhookURL *string `form:"WebhookUrl,omitempty"`
}

// UpdateWebhookResponse resource/ response properties for the updated webhook
type UpdateWebhookResponse struct {
	// The SID of the account which the resource is associated with
	AccountSid string `json:"account_sid"`
	// The SID of the assistant which the webhook is associated with
	AssistantSid string `json:"assistant_sid"`
	// The date and time (in RFC3339 format) when the resource was created
	DateCreated time.Time `json:"date_created"`
	// The date and time (in RFC3339 format) when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Space separated string of event which triggers the webhook
	Events string `json:"events"`
	// The unique alphanumeric string for the resource
	Sid string `json:"sid"`
	// The URL for the resource
	URL string `json:"url"`
	// The unique human readable name of the webhook
	UniqueName string `json:"unique_name"`
	// The HTTP method used to call the webhook URL
	WebhookMethod string `json:"webhook_method"`
	// The URL of the webhook
	WebhookURL string `json:"webhook_url"`
}

// Update modifies an webhook resource
// See https://www.twilio.com/docs/autopilot/api/event-webhooks#update-a-webhook-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateWebhookInput) (*UpdateWebhookResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies an webhook resource
// See https://www.twilio.com/docs/autopilot/api/event-webhooks#update-a-webhook-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateWebhookInput) (*UpdateWebhookResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{assistantSid}/Webhooks/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"sid":          c.sid,
		},
	}

	response := &UpdateWebhookResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
