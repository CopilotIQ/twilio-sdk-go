// This is an autogenerated file. DO NOT MODIFY
package context

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type FetchContextResponse struct {
	AccountSid   string      `json:"account_sid"`
	Context      interface{} `json:"context"`
	ExecutionSid string      `json:"execution_sid"`
	FlowSid      string      `json:"flow_sid"`
	URL          string      `json:"url"`
}

func (c Client) Fetch() (*FetchContextResponse, error) {
	return c.FetchWithContext(context.Background())
}

func (c Client) FetchWithContext(context context.Context) (*FetchContextResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Flows/{flowSid}/Executions/{executionSid}/Context",
		PathParams: map[string]string{
			"flowSid":      c.flowSid,
			"executionSid": c.executionSid,
		},
	}

	response := &FetchContextResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}