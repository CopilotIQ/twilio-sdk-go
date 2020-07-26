// This is an autogenerated file. DO NOT MODIFY
package service

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type UpdateServiceInput struct {
	AclEnabled                    *bool   `form:"AclEnabled,omitempty"`
	FriendlyName                  *string `form:"FriendlyName,omitempty"`
	ReachabilityDebouncingEnabled *bool   `form:"ReachabilityDebouncingEnabled,omitempty"`
	ReachabilityDebouncingWindow  *int    `form:"ReachabilityDebouncingWindow,omitempty"`
	ReachabilityWebhooksEnabled   *bool   `form:"ReachabilityWebhooksEnabled,omitempty"`
	WebhookURL                    *string `form:"WebhookUrl,omitempty"`
	WebhooksFromRestEnabled       *bool   `form:"WebhooksFromRestEnabled,omitempty"`
}

type UpdateServiceResponse struct {
	AccountSid                    string     `json:"account_sid"`
	AclEnabled                    bool       `json:"acl_enabled"`
	DateCreated                   time.Time  `json:"date_created"`
	DateUpdated                   *time.Time `json:"date_updated,omitempty"`
	FriendlyName                  *string    `json:"friendly_name,omitempty"`
	ReachabilityDebouncingEnabled bool       `json:"reachability_debouncing_enabled"`
	ReachabilityDebouncingWindow  int        `json:"reachability_debouncing_window"`
	ReachabilityWebhooksEnabled   bool       `json:"reachability_webhooks_enabled"`
	Sid                           string     `json:"sid"`
	URL                           string     `json:"url"`
	UniqueName                    *string    `json:"unique_name,omitempty"`
	WebhookURL                    *string    `json:"webhook_url,omitempty"`
	WebhooksFromRestEnabled       bool       `json:"webhooks_from_rest_enabled"`
}

func (c Client) Update(input *UpdateServiceInput) (*UpdateServiceResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

func (c Client) UpdateWithContext(context context.Context, input *UpdateServiceInput) (*UpdateServiceResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &UpdateServiceResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
