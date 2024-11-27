// Package short_code contains auto-generated files. DO NOT MODIFY
package short_code

import (
	"context"
	"net/http"
	"time"

	"github.com/CopilotIQ/twilio-sdk-go/client"
)

// FetchShortCodeResponse defines the response fields for the retrieved short code
type FetchShortCodeResponse struct {
	AccountSid   string     `json:"account_sid"`
	Capabilities []string   `json:"capabilities"`
	CountryCode  string     `json:"country_code"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	ServiceSid   string     `json:"service_sid"`
	ShortCode    string     `json:"short_code"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
}

// Fetch retrieves a short code resource
// See https://www.twilio.com/docs/sms/services/api/shortcode-resource#fetch-a-shortcode-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchShortCodeResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a short code resource
// See https://www.twilio.com/docs/sms/services/api/shortcode-resource#fetch-a-shortcode-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchShortCodeResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/ShortCodes/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	response := &FetchShortCodeResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
