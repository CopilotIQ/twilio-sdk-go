// This is an autogenerated file. DO NOT MODIFY
package flow_validation

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type ValidateFlowInput struct {
	FriendlyName  string  `validate:"required" form:"FriendlyName"`
	Status        string  `validate:"required" form:"Status"`
	Definition    string  `validate:"required" form:"Definition"`
	CommitMessage *string `form:"CommitMessage,omitempty"`
}

type ValidateFlowOutput struct {
	Valid bool `json:"valid"`
}

func (c Client) Validate(input *ValidateFlowInput) (*ValidateFlowOutput, error) {
	return c.ValidateWithContext(context.Background(), input)
}

func (c Client) ValidateWithContext(context context.Context, input *ValidateFlowInput) (*ValidateFlowOutput, error) {
	op := client.Operation{
		HTTPMethod:  http.MethodPost,
		HTTPPath:    "/Flows/Validate",
		ContentType: client.URLEncoded,
	}

	output := &ValidateFlowOutput{}
	if err := c.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}
