// This is an autogenerated file. DO NOT MODIFY
package webhook

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type UpdateConversationWebhookInput struct {
	ConfigurationUrl      *string   `form:"Configuration.Url,omitempty"`
	ConfigurationMethod   *string   `form:"Configuration.Method,omitempty"`
	ConfigurationFilters  *[]string `form:"Configuration.Filters,omitempty"`
	ConfigurationTriggers *[]string `form:"Configuration.Triggers,omitempty"`
	ConfigurationFlowSid  *string   `form:"Configuration.FlowSid,omitempty"`
}

type UpdateConversationWebhookOutputConfiguration struct {
	Url         *string   `json:"url,omitempty"`
	Method      *string   `json:"method,omitempty"`
	Filters     *[]string `json:"filters,omitempty"`
	Triggers    *[]string `json:"triggers,omitempty"`
	FlowSid     *string   `json:"flow_sid,omitempty"`
	ReplayAfter *int      `json:"replay_after,omitempty"`
}

type UpdateConversationWebhookOutput struct {
	Sid             string                                       `json:"sid"`
	AccountSid      string                                       `json:"account_sid"`
	ConversationSid string                                       `json:"conversation_sid"`
	Target          string                                       `json:"target"`
	Configuration   UpdateConversationWebhookOutputConfiguration `json:"configuration"`
	DateCreated     time.Time                                    `json:"date_created"`
	DateUpdated     *time.Time                                   `json:"date_updated,omitempty"`
	URL             string                                       `json:"url"`
}

func (c Client) Update(input *UpdateConversationWebhookInput) (*UpdateConversationWebhookOutput, error) {
	return c.UpdateWithContext(context.Background(), input)
}

func (c Client) UpdateWithContext(context context.Context, input *UpdateConversationWebhookInput) (*UpdateConversationWebhookOutput, error) {
	op := client.Operation{
		HTTPMethod:  http.MethodPost,
		HTTPPath:    "/Conversations/{conversationSid}/Webhooks/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"conversationSid": c.conversationSid,
			"sid":             c.sid,
		},
	}

	output := &UpdateConversationWebhookOutput{}
	if err := c.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}
