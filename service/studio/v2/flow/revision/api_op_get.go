// This is an autogenerated file. DO NOT MODIFY
package revision

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type GetRevisionOutput struct {
	Sid           string         `json:"sid"`
	AccountSid    string         `json:"account_sid"`
	FriendlyName  string         `json:"friendly_name"`
	Definition    interface{}    `json:"definition"`
	Status        string         `json:"status"`
	Revision      *int           `json:"revision,string"`
	CommitMessage *string        `json:"commit_message,omitempty"`
	Valid         bool           `json:"valid"`
	Errors        *[]interface{} `json:"errors,omitempty"`
	Warnings      *[]interface{} `json:"warnings,omitempty"`
	DateCreated   time.Time      `json:"date_created"`
	DateUpdated   *time.Time     `json:"date_updated,omitempty"`
	WebhookURL    string         `json:"webhook_url"`
	URL           string         `json:"url"`
}

func (c Client) Get() (*GetRevisionOutput, error) {
	return c.GetWithContext(context.Background())
}

func (c Client) GetWithContext(context context.Context) (*GetRevisionOutput, error) {
	op := client.Operation{
		HTTPMethod: http.MethodGet,
		HTTPPath:   "/Flows/{flowSid}/Revisions/{revisionNumber}",
		PathParams: map[string]string{
			"flowSid":        c.flowSid,
			"revisionNumber": strconv.Itoa(c.revisionNumber),
		},
	}

	output := &GetRevisionOutput{}
	if err := c.client.Send(context, op, nil, output); err != nil {
		return nil, err
	}
	return output, nil
}
