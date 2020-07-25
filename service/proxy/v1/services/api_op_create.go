// This is an autogenerated file. DO NOT MODIFY
package services

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type CreateServiceInput struct {
	CallbackUrl             *string `form:"CallbackUrl,omitempty"`
	ChatInstanceSid         *string `form:"ChatInstanceSid,omitempty"`
	DefaultTtl              *int    `form:"DefaultTtl,omitempty"`
	GeoMatchLevel           *string `form:"GeoMatchLevel,omitempty"`
	InterceptCallbackUrl    *string `form:"InterceptCallbackUrl,omitempty"`
	NumberSelectionBehavior *string `form:"NumberSelectionBehavior,omitempty"`
	OutOfSessionCallbackUrl *string `form:"OutOfSessionCallbackUrl,omitempty"`
	UniqueName              string  `validate:"required" form:"UniqueName"`
}

type CreateServiceResponse struct {
	AccountSid              string     `json:"account_sid"`
	CallbackUrl             *string    `json:"callback_url,omitempty"`
	ChatInstanceSid         *string    `json:"chat_instance_sid,omitempty"`
	ChatServiceSid          string     `json:"chat_service_sid"`
	DateCreated             time.Time  `json:"date_created"`
	DateUpdated             *time.Time `json:"date_updated,omitempty"`
	DefaultTtl              *int       `json:"default_ttl,omitempty"`
	GeoMatchLevel           *string    `json:"geo_match_level,omitempty"`
	InterceptCallbackUrl    *string    `json:"intercept_callback_url,omitempty"`
	NumberSelectionBehavior *string    `json:"number_selection_behavior,omitempty"`
	OutOfSessionCallbackUrl *string    `json:"out_of_session_callback_url,omitempty"`
	Sid                     string     `json:"sid"`
	URL                     string     `json:"url"`
	UniqueName              string     `json:"unique_name"`
}

func (c Client) Create(input *CreateServiceInput) (*CreateServiceResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

func (c Client) CreateWithContext(context context.Context, input *CreateServiceInput) (*CreateServiceResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services",
		ContentType: client.URLEncoded,
	}

	response := &CreateServiceResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
