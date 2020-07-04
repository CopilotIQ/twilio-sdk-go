// This is an autogenerated file. DO NOT MODIFY
package webhook

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type UpdateWebhookInput struct {
	Method         *string   `form:"Method,omitempty"`
	Filters        *[]string `form:"Filters,omitempty"`
	PreWebhookUrl  *string   `form:"PreWebhookUrl,omitempty"`
	PostWebhookUrl *string   `form:"PostWebhookUrl,omitempty"`
	Target         *string   `form:"Target,omitempty"`
}

type UpdateWebhookOutput struct {
	AccountSid     string   `json:"account_sid"`
	Method         string   `json:"method"`
	Target         string   `json:"target"`
	Filters        []string `json:"filters"`
	PreWebhookUrl  *string  `json:"pre_webhook_url,omitempty"`
	PostWebhookUrl *string  `json:"post_webhook_url,omitempty"`
	URL            string   `json:"url"`
}

func (c Client) Update(input *UpdateWebhookInput) (*UpdateWebhookOutput, error) {
	return c.UpdateWithContext(context.Background(), input)
}

func (c Client) UpdateWithContext(context context.Context, input *UpdateWebhookInput) (*UpdateWebhookOutput, error) {
	op := client.Operation{
		HTTPMethod:  http.MethodPost,
		HTTPPath:    "/Conversations/Webhooks",
		ContentType: client.URLEncoded,
	}

	output := &UpdateWebhookOutput{}
	if err := c.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}
