// This is an autogenerated file. DO NOT MODIFY
package query

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type FetchQueryResponseField struct {
	// The name of the query
	Name string `json:"name"`
	// The field type
	Type string `json:"type"`
	// The value of the query
	Value string `json:"value"`
}

type FetchQueryResponseResult struct {
	// The fields that were extracted from the query
	Fields []FetchQueryResponseField `json:"fields"`
	// The name of the recognised task
	Task string `json:"task"`
}

// FetchQueryResponse resource/ response properties for the retrieved query
type FetchQueryResponse struct {
	// The SID of the account which the resource is associated with
	AccountSid string `json:"account_sid"`
	// The SID of the assistant which the query is associated with
	AssistantSid string `json:"assistant_sid"`
	// The date and time (in RFC3339 format) when the resource was created
	DateCreated time.Time `json:"date_created"`
	// The date and time (in RFC3339 format) when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The SID of the dialogue which the query is associated with
	DialogueSid *string `json:"dialogue_sid,omitempty"`
	// The ISO language country code for the query
	Language string `json:"language"`
	// The SID of the model build which the query is associated with
	ModelBuildSid string `json:"model_build_sid"`
	// The query text
	Query string `json:"query"`
	// An object containing the recognised task and extracted field
	Results FetchQueryResponseResult `json:"results"`
	// The SID of the sample which the query is associated with
	SampleSid string `json:"sample_sid"`
	// The unique alphanumeric string for the resource
	Sid string `json:"sid"`
	// The channel the query was sent to
	SourceChannel string `json:"source_channel"`
	// The current status of the query
	Status string `json:"status"`
	// The URL for the resource
	URL string `json:"url"`
}

// Fetch retrieves a query resource
// See https://www.twilio.com/docs/autopilot/api/query#fetch-a-query-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchQueryResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a query resource
// See https://www.twilio.com/docs/autopilot/api/query#fetch-a-query-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchQueryResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/Queries/{sid}",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"sid":          c.sid,
		},
	}

	response := &FetchQueryResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
