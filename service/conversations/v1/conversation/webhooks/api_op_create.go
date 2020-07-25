// This is an autogenerated file. DO NOT MODIFY
package webhooks

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type CreateConversationWebhookInput struct {
	ConfigurationFilters     *[]string `form:"Configuration.Filters,omitempty"`
	ConfigurationFlowSid     *string   `form:"Configuration.FlowSid,omitempty"`
	ConfigurationMethod      *string   `form:"Configuration.Method,omitempty"`
	ConfigurationReplayAfter *int      `form:"Configuration.ReplayAfter,omitempty"`
	ConfigurationTriggers    *[]string `form:"Configuration.Triggers,omitempty"`
	ConfigurationUrl         *string   `form:"Configuration.Url,omitempty"`
	Target                   string    `validate:"required" form:"target"`
}

type CreateConversationWebhookResponseConfiguration struct {
	Filters     *[]string `json:"filters,omitempty"`
	FlowSid     *string   `json:"flow_sid,omitempty"`
	Method      *string   `json:"method,omitempty"`
	ReplayAfter *int      `json:"replay_after,omitempty"`
	Triggers    *[]string `json:"triggers,omitempty"`
	URL         *string   `json:"url,omitempty"`
}

type CreateConversationWebhookResponse struct {
	AccountSid      string                                         `json:"account_sid"`
	Configuration   CreateConversationWebhookResponseConfiguration `json:"configuration"`
	ConversationSid string                                         `json:"conversation_sid"`
	DateCreated     time.Time                                      `json:"date_created"`
	DateUpdated     *time.Time                                     `json:"date_updated,omitempty"`
	Sid             string                                         `json:"sid"`
	Target          string                                         `json:"target"`
	URL             string                                         `json:"url"`
}

func (c Client) Create(input *CreateConversationWebhookInput) (*CreateConversationWebhookResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

func (c Client) CreateWithContext(context context.Context, input *CreateConversationWebhookInput) (*CreateConversationWebhookResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Conversations/{conversationSid}/Webhooks",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"conversationSid": c.conversationSid,
		},
	}

	response := &CreateConversationWebhookResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
