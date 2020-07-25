// This is an autogenerated file. DO NOT MODIFY
package execution

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type GetExecutionResponse struct {
	AccountSid            string      `json:"account_sid"`
	ContactChannelAddress string      `json:"contact_channel_address"`
	Context               interface{} `json:"context"`
	DateCreated           time.Time   `json:"date_created"`
	DateUpdated           *time.Time  `json:"date_updated,omitempty"`
	FlowSid               string      `json:"flow_sid"`
	Sid                   string      `json:"sid"`
	Status                string      `json:"status"`
	URL                   string      `json:"url"`
}

func (c Client) Get() (*GetExecutionResponse, error) {
	return c.GetWithContext(context.Background())
}

func (c Client) GetWithContext(context context.Context) (*GetExecutionResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Flows/{flowSid}/Executions/{sid}",
		PathParams: map[string]string{
			"flowSid": c.flowSid,
			"sid":     c.sid,
		},
	}

	response := &GetExecutionResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
