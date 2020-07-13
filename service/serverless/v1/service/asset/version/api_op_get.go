// This is an autogenerated file. DO NOT MODIFY
package version

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type GetVersionOutput struct {
	Sid         string    `json:"sid"`
	AccountSid  string    `json:"account_sid"`
	ServiceSid  string    `json:"service_sid"`
	AssetSid    string    `json:"asset_sid"`
	Path        string    `json:"path"`
	Visibility  string    `json:"visibility"`
	DateCreated time.Time `json:"date_created"`
	URL         string    `json:"url"`
}

func (c Client) Get() (*GetVersionOutput, error) {
	return c.GetWithContext(context.Background())
}

func (c Client) GetWithContext(context context.Context) (*GetVersionOutput, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Assets/{assetSid}/Versions/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"assetSid":   c.assetSid,
			"sid":        c.sid,
		},
	}

	output := &GetVersionOutput{}
	if err := c.client.Send(context, op, nil, output); err != nil {
		return nil, err
	}
	return output, nil
}
