// This is an autogenerated file. DO NOT MODIFY
package variable

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type GetVariableOutput struct {
	Sid            string     `json:"sid"`
	AccountSid     string     `json:"account_sid"`
	ServiceSid     string     `json:"service_sid"`
	EnvironmentSid string     `json:"environment_sid"`
	Key            string     `json:"key"`
	Value          string     `json:"value"`
	DateCreated    time.Time  `json:"date_created"`
	DateUpdated    *time.Time `json:"date_updated,omitempty"`
	URL            string     `json:"url"`
}

func (c Client) Get() (*GetVariableOutput, error) {
	return c.GetWithContext(context.Background())
}

func (c Client) GetWithContext(context context.Context) (*GetVariableOutput, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Environments/{environmentSid}/Variables/{sid}",
		PathParams: map[string]string{
			"serviceSid":     c.serviceSid,
			"environmentSid": c.environmentSid,
			"sid":            c.sid,
		},
	}

	output := &GetVariableOutput{}
	if err := c.client.Send(context, op, nil, output); err != nil {
		return nil, err
	}
	return output, nil
}
