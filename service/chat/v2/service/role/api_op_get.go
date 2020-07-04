// This is an autogenerated file. DO NOT MODIFY
package role

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type GetRoleOutput struct {
	Sid          string     `json:"sid"`
	AccountSid   string     `json:"account_sid"`
	ServiceSid   string     `json:"service_sid"`
	FriendlyName string     `json:"friendly_name"`
	Type         string     `json:"type"`
	Permissions  []string   `json:"permissions"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	URL          string     `json:"url"`
}

func (c Client) Get() (*GetRoleOutput, error) {
	return c.GetWithContext(context.Background())
}

func (c Client) GetWithContext(context context.Context) (*GetRoleOutput, error) {
	op := client.Operation{
		HTTPMethod: http.MethodGet,
		HTTPPath:   "/Services/{serviceSid}/Roles/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	output := &GetRoleOutput{}
	if err := c.client.Send(context, op, nil, output); err != nil {
		return nil, err
	}
	return output, nil
}