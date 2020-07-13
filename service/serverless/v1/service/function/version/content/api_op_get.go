// This is an autogenerated file. DO NOT MODIFY
package content

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type GetContentOutput struct {
	Sid         string `json:"sid"`
	AccountSid  string `json:"account_sid"`
	ServiceSid  string `json:"service_sid"`
	FunctionSid string `json:"function_sid"`
	Content     string `json:"content"`
	URL         string `json:"url"`
}

func (c Client) Get() (*GetContentOutput, error) {
	return c.GetWithContext(context.Background())
}

func (c Client) GetWithContext(context context.Context) (*GetContentOutput, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Functions/{functionSid}/Versions/{versionSid}/Content",
		PathParams: map[string]string{
			"serviceSid":  c.serviceSid,
			"functionSid": c.functionSid,
			"versionSid":  c.versionSid,
		},
	}

	output := &GetContentOutput{}
	if err := c.client.Send(context, op, nil, output); err != nil {
		return nil, err
	}
	return output, nil
}
