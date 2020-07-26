// This is an autogenerated file. DO NOT MODIFY
package flex_flow

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type FetchFlexFlowResponseIntegration struct {
	Channel           *string `json:"channel,omitempty"`
	CreationOnMessage *string `json:"creation_on_message,omitempty"`
	FlowSid           *string `json:"flow_sid,omitempty"`
	Priority          *int    `json:"priority,omitempty"`
	RetryCount        *int    `json:"retry_count,omitempty"`
	Timeout           *int    `json:"timeout,omitempty"`
	URL               *string `json:"url,omitempty"`
	WorkspaceSid      *string `json:"workspace_sid,omitempty"`
}

type FetchFlexFlowResponse struct {
	AccountSid      string                            `json:"account_sid"`
	ChannelType     string                            `json:"channel_type"`
	ChatServiceSid  string                            `json:"chat_service_sid"`
	ContactIdentity *string                           `json:"contact_identity,omitempty"`
	DateCreated     time.Time                         `json:"date_created"`
	DateUpdated     *time.Time                        `json:"date_updated,omitempty"`
	Enabled         bool                              `json:"enabled"`
	FriendlyName    string                            `json:"friendly_name"`
	Integration     *FetchFlexFlowResponseIntegration `json:"integration,omitempty"`
	IntegrationType *string                           `json:"integration_type,omitempty"`
	JanitorEnabled  *bool                             `json:"janitor_enabled,omitempty"`
	LongLived       *bool                             `json:"long_lived,omitempty"`
	Sid             string                            `json:"sid"`
	URL             string                            `json:"url"`
}

func (c Client) Fetch() (*FetchFlexFlowResponse, error) {
	return c.FetchWithContext(context.Background())
}

func (c Client) FetchWithContext(context context.Context) (*FetchFlexFlowResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/FlexFlows/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchFlexFlowResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
