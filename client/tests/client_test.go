package tests

import (
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/CopilotIQ/twilio-sdk-go/client"
	"github.com/CopilotIQ/twilio-sdk-go/session"
	"github.com/CopilotIQ/twilio-sdk-go/session/credentials"
	"github.com/CopilotIQ/twilio-sdk-go/utils"
)

var _ = Describe("Client", func() {
	Describe("Given the client", func() {
		apiClientConfig := &client.APIClientConfig{
			Beta:         false,
			DebugEnabled: false,
			RetryConfig: client.RetryConfig{
				Attempts: 0,
				WaitTime: 0,
			},
			SubDomain:  "test",
			APIVersion: "v1",
		}

		sess := session.New(&credentials.Credentials{
			Username: "Test Username",
			Password: "Test Password",
		})

		twilioClient := client.New(sess, apiClientConfig)

		httpmock.ActivateNonDefault(twilioClient.GetRestyClient().GetClient())
		defer httpmock.DeactivateAndReset()

		Describe("When a POST request is made with path params, input & output interfaces", func() {
			httpmock.RegisterResponder("POST", "https://test.twilio.com/v1/test/1234",
				func(req *http.Request) (*http.Response, error) {
					resp := map[string]string{
						"id":   "1234",
						"name": "test",
					}
					return httpmock.NewJsonResponse(201, &resp)
				},
			)

			op := client.Operation{
				Method:      http.MethodPost,
				URI:         "/test/{id}",
				ContentType: client.URLEncoded,
				PathParams: map[string]string{
					"id": "1234",
				},
			}

			testInput := &TestStructInput{
				Name: "test",
			}
			testOutput := &TestStructResponse{}
			err := twilioClient.Send(context.Background(), op, testInput, testOutput)

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the output should contain a id and name", func() {
				Expect(testOutput.ID).To(Equal("1234"))
				Expect(testOutput.Name).To(Equal("test"))
			})
		})

		Describe("When a POST request is made without name property", func() {
			httpmock.RegisterResponder("POST", "https://test.twilio.com/v1/test/1234",
				func(req *http.Request) (*http.Response, error) {
					resp := map[string]string{
						"id":   "1234",
						"name": "test",
					}
					return httpmock.NewJsonResponse(201, &resp)
				},
			)

			op := client.Operation{
				Method:      http.MethodPost,
				URI:         "/test/{id}",
				ContentType: client.URLEncoded,
				PathParams: map[string]string{
					"id": "1234",
				},
			}

			testInput := &TestStructInput{}
			err := twilioClient.Send(context.Background(), op, testInput, nil)

			It("Then input validation error should be returned", func() {
				Expect(err.Error()).To(Equal("Invalid input supplied"))
			})
		})

		Describe("When a POST request is made with a unsupported content type", func() {
			op := client.Operation{
				Method:      http.MethodPost,
				URI:         "/test/{id}",
				ContentType: "application/pdf",
				PathParams: map[string]string{
					"id": "1234",
				},
			}

			testInput := &TestStructInput{
				Name: "Test",
			}
			err := twilioClient.Send(context.Background(), op, testInput, nil)

			It("Then unsupported content type error should be returned", func() {
				Expect(err.Error()).To(Equal("application/pdf is not a supported content type"))
			})
		})

		Describe("When a POST request is made with a form data", func() {
			httpmock.RegisterResponder("POST", "https://test.twilio.com/v1/test/1234",
				func(req *http.Request) (*http.Response, error) {
					resp := map[string]string{
						"id":   "1234",
						"name": "test",
					}
					return httpmock.NewJsonResponse(201, &resp)
				},
			)

			op := client.Operation{
				Method:      http.MethodPost,
				URI:         "/test/{id}",
				ContentType: client.FormData,
				PathParams: map[string]string{
					"id": "1234",
				},
			}

			testInput := &TestStructInput{
				Name: "Test",
			}
			testOutput := &TestStructResponse{}
			err := twilioClient.Send(context.Background(), op, testInput, testOutput)

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the output should contain a id and name", func() {
				Expect(testOutput.ID).To(Equal("1234"))
				Expect(testOutput.Name).To(Equal("test"))
			})
		})

		Describe("When a POST request is made with a form data and content deatils", func() {
			httpmock.RegisterResponder("POST", "https://test.twilio.com/v1/test/1234",
				func(req *http.Request) (*http.Response, error) {
					resp := map[string]string{
						"id":   "1234",
						"name": "test",
					}
					return httpmock.NewJsonResponse(201, &resp)
				},
			)

			op := client.Operation{
				Method:      http.MethodPost,
				URI:         "/test/{id}",
				ContentType: client.FormData,
				PathParams: map[string]string{
					"id": "1234",
				},
			}

			testInput := &TestContentStructInput{
				Name: "Test",
				Content: ContentDetails{
					Body:        strings.NewReader("{}"),
					ContentType: "application/json",
					FileName:    "test.json",
				},
			}
			testOutput := &TestStructResponse{}
			err := twilioClient.Send(context.Background(), op, testInput, testOutput)

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the output should contain a id and name", func() {
				Expect(testOutput.ID).To(Equal("1234"))
				Expect(testOutput.Name).To(Equal("test"))
			})
		})

		Describe("When a DELETE request is made with no path params, input & output interfaces", func() {
			httpmock.RegisterResponder("DELETE", "https://test.twilio.com/v1/test", httpmock.NewStringResponder(200, ""))

			op := client.Operation{
				Method: http.MethodDelete,
				URI:    "/test",
			}

			err := twilioClient.Send(context.Background(), op, nil, nil)

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When a DELETE request is made with an overridden base URI", func() {
			httpmock.RegisterResponder("DELETE", "https://test-2.twilio.com/v2/test", httpmock.NewStringResponder(200, ""))

			op := client.Operation{
				OverrideBaseURL: utils.String(client.CreateBaseURL("test-2", "v2", nil, nil)),
				Method:          http.MethodDelete,
				URI:             "/test",
			}

			err := twilioClient.Send(context.Background(), op, nil, nil)

			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("When a DELETE request is made with region and edge configuration", func() {
		httpmock.RegisterResponder("DELETE", "https://test.dublin.ie1.twilio.com/v1/request", httpmock.NewStringResponder(200, ""))

		apiClientConfig := &client.APIClientConfig{
			Beta:         false,
			DebugEnabled: false,
			RetryConfig: client.RetryConfig{
				Attempts: 0,
				WaitTime: 0,
			},
			SubDomain:  "test",
			APIVersion: "v1",
			Edge:       utils.String("dublin"),
			Region:     utils.String("ie1"),
		}

		sess := session.New(&credentials.Credentials{
			Username: "Test Username",
			Password: "Test Password",
		})

		twilioClient := client.New(sess, apiClientConfig)

		httpmock.ActivateNonDefault(twilioClient.GetRestyClient().GetClient())
		defer httpmock.DeactivateAndReset()

		op := client.Operation{
			Method: http.MethodDelete,
			URI:    "/request",
		}

		err := twilioClient.Send(context.Background(), op, nil, nil)

		It("Then no error should be returned", func() {
			Expect(err).To(BeNil())
		})
	})

	Describe("When a DELETE request is made with region and default edge configuration", func() {
		httpmock.RegisterResponder("DELETE", "https://test.dublin.us1.twilio.com/v1/request", httpmock.NewStringResponder(200, ""))

		apiClientConfig := &client.APIClientConfig{
			Beta:         false,
			DebugEnabled: false,
			RetryConfig: client.RetryConfig{
				Attempts: 0,
				WaitTime: 0,
			},
			SubDomain:  "test",
			APIVersion: "v1",
			Edge:       utils.String("dublin"),
		}

		sess := session.New(&credentials.Credentials{
			Username: "Test Username",
			Password: "Test Password",
		})

		twilioClient := client.New(sess, apiClientConfig)

		httpmock.ActivateNonDefault(twilioClient.GetRestyClient().GetClient())
		defer httpmock.DeactivateAndReset()

		op := client.Operation{
			Method: http.MethodDelete,
			URI:    "/request",
		}

		err := twilioClient.Send(context.Background(), op, nil, nil)

		It("Then no error should be returned", func() {
			Expect(err).To(BeNil())
		})
	})

	Describe("When a DELETE request is made with edge configuration", func() {
		httpmock.RegisterResponder("DELETE", "https://test.au1.twilio.com/v1/request", httpmock.NewStringResponder(200, ""))

		apiClientConfig := &client.APIClientConfig{
			Beta:         false,
			DebugEnabled: false,
			RetryConfig: client.RetryConfig{
				Attempts: 0,
				WaitTime: 0,
			},
			SubDomain:  "test",
			APIVersion: "v1",
			Region:     utils.String("au1"),
		}

		sess := session.New(&credentials.Credentials{
			Username: "Test Username",
			Password: "Test Password",
		})

		twilioClient := client.New(sess, apiClientConfig)

		httpmock.ActivateNonDefault(twilioClient.GetRestyClient().GetClient())
		defer httpmock.DeactivateAndReset()

		op := client.Operation{
			Method: http.MethodDelete,
			URI:    "/request",
		}

		err := twilioClient.Send(context.Background(), op, nil, nil)

		It("Then no error should be returned", func() {
			Expect(err).To(BeNil())
		})
	})

})

type TestStructInput struct {
	Name string `validate:"required"`
}

type ContentDetails struct {
	Body        io.ReadSeeker `validate:"required"`
	ContentType string        `validate:"required"`
	FileName    string        `validate:"required"`
}

type TestContentStructInput struct {
	Name    string         `validate:"required"`
	Content ContentDetails `validate:"required"`
}

type TestStructResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
