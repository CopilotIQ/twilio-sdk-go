// Package user contains auto-generated files. DO NOT MODIFY
package user

import (
	"context"
	"net/http"

	"github.com/CopilotIQ/twilio-sdk-go/client"
)

// Delete removes a user resource from the account
// See https://www.twilio.com/docs/chat/rest/user-resource#delete-a-user-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a user resource from the account
// See https://www.twilio.com/docs/chat/rest/user-resource#delete-a-user-resource for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Services/{serviceSid}/Users/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
