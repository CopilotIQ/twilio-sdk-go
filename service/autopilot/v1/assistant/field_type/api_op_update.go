// This is an autogenerated file. DO NOT MODIFY
package field_type

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type UpdateFieldTypeInput struct {
	UniqueName   *string `form:"UniqueName,omitempty"`
	FriendlyName *string `form:"FriendlyName,omitempty"`
}

type UpdateFieldTypeResponse struct {
	Sid          string     `json:"sid"`
	AccountSid   string     `json:"account_sid"`
	AssistantSid string     `json:"assistant_sid"`
	UniqueName   string     `json:"unique_name"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	URL          string     `json:"url"`
}

func (c Client) Update(input *UpdateFieldTypeInput) (*UpdateFieldTypeResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

func (c Client) UpdateWithContext(context context.Context, input *UpdateFieldTypeInput) (*UpdateFieldTypeResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{assistantSid}/FieldTypes/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"sid":          c.sid,
		},
	}

	response := &UpdateFieldTypeResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}