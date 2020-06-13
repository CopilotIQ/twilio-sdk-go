// This is an autogenerated file. DO NOT MODIFY
package builds

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type CreateBuildInput struct {
	AssetVersions    *[]string `form:"AssetVersions,omitempty"`
	FunctionVersions *[]string `form:"FunctionVersions,omitempty"`
	Dependencies     *[]string `form:"Dependencies,omitempty"`
}

type FunctionVersion struct {
	Sid         string    `json:"sid"`
	AccountSid  string    `json:"account_sid"`
	ServiceSid  string    `json:"service_sid"`
	FunctionSid string    `json:"function_sid"`
	DateCreated time.Time `json:"date_created"`
	Path        string    `json:"path"`
	Visibility  string    `json:"visibility"`
}

type AssetVersion struct {
	Sid         string    `json:"sid"`
	AccountSid  string    `json:"account_sid"`
	ServiceSid  string    `json:"service_sid"`
	AssetSid    string    `json:"asset_sid"`
	DateCreated time.Time `json:"date_created"`
	Path        string    `json:"path"`
	Visibility  string    `json:"visibility"`
}

type Dependency struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type CreateBuildOutput struct {
	Sid              string             `json:"sid"`
	AccountSid       string             `json:"account_sid"`
	ServiceSid       string             `json:"service_sid"`
	AssetVersions    *[]AssetVersion    `json:"asset_versions,omitempty"`
	FunctionVersions *[]FunctionVersion `json:"function_versions,omitempty"`
	Dependencies     *[]Dependency      `json:"dependencies,omitempty"`
	Status           string             `json:"status"`
	DateCreated      time.Time          `json:"date_created"`
	DateUpdated      *time.Time         `json:"date_updated,omitempty"`
	URL              string             `json:"url"`
}

func (c Client) Create(input *CreateBuildInput) (*CreateBuildOutput, error) {
	return c.CreateWithContext(context.Background(), input)
}

func (c Client) CreateWithContext(context context.Context, input *CreateBuildInput) (*CreateBuildOutput, error) {
	op := client.Operation{
		HTTPMethod:  http.MethodPost,
		HTTPPath:    "/Services/{serviceSid}/Builds",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	output := &CreateBuildOutput{}
	if err := c.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}
