// Package service contains auto-generated files. DO NOT MODIFY
package service

import (
	"context"
	"net/http"

	"github.com/CopilotIQ/twilio-sdk-go/client"
)

// Delete removes a service resource from the account
// See https://www.twilio.com/docs/verify/api/service#delete-a-service for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a service resource from the account
// See https://www.twilio.com/docs/verify/api/service#delete-a-service for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Services/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
