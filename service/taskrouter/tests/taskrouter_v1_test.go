package tests

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task_queue"

	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task_queues"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspaces"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Taskrouter V1", func() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	creds, err := credentials.New(credentials.Account{
		Sid:       "ACxxx",
		AuthToken: "Test",
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	taskrouterSession := taskrouter.NewWithCredentials(creds).V1

	Describe("Given the Workspace Client", func() {
		workspacesClient := taskrouterSession.Workspaces

		Describe("When the Task Queue is successfully created", func() {
			createInput := &workspaces.CreateWorkspaceInput{
				FriendlyName: "Test 2",
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workspaceResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := workspacesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create workspace response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DefaultActivityName).To(Equal("Offline"))
				Expect(resp.DefaultActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				eventCallbackURL := "https://ngrok.com"
				Expect(resp.EventCallbackURL).To(Equal(&eventCallbackURL))
				Expect(resp.EventsFilter).To(BeNil())
				Expect(resp.FriendlyName).To(Equal("NewWorkspace"))
				Expect(resp.MultiTaskEnabled).To(Equal(false))
				Expect(resp.PrioritizeQueueOrder).To(Equal("FIFO"))
				Expect(resp.TimeoutActivityName).To(Equal("Offline"))
				Expect(resp.TimeoutActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1"))
				Expect(resp.Sid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the Workspace request does not contain a friendly name", func() {
			createInput := &workspaces.CreateWorkspaceInput{}

			resp, err := workspacesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create workspace response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Workspaces API returns a 500 response", func() {
			createInput := &workspaces.CreateWorkspaceInput{
				FriendlyName: "Test 2",
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := workspacesClient.Create(createInput)
			It("Then an error should be returned", func() {
				Expect(err).ToNot(BeNil())
				twilioErr, ok := err.(*utils.TwilioError)
				Expect(ok).To(Equal(true))
				Expect(twilioErr.Code).To(BeNil())
				Expect(twilioErr.Message).To(Equal("An error occurred"))
				Expect(twilioErr.MoreInfo).To(BeNil())
				Expect(twilioErr.Status).To(Equal(500))
			})

			It("Then the create workspace response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a Workspace SID", func() {
		workspaceClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the Workspace is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/workspaceResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := workspaceClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get workspace response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DefaultActivityName).To(Equal("Offline"))
				Expect(resp.DefaultActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				eventCallbackURL := "https://ngrok.com"
				Expect(resp.EventCallbackURL).To(Equal(&eventCallbackURL))
				Expect(resp.EventsFilter).To(BeNil())
				Expect(resp.FriendlyName).To(Equal("NewWorkspace"))
				Expect(resp.MultiTaskEnabled).To(Equal(false))
				Expect(resp.PrioritizeQueueOrder).To(Equal("FIFO"))
				Expect(resp.TimeoutActivityName).To(Equal("Offline"))
				Expect(resp.TimeoutActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1"))
				Expect(resp.Sid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get workspace response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := taskrouterSession.Workspace("WS71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get workspace response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Workspace is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updatedWorkspaceResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &workspace.UpdateWorkspaceInput{}

			resp, err := workspaceClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update workspace response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DefaultActivityName).To(Equal("Offline"))
				Expect(resp.DefaultActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2016-08-01T23:10:40Z"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2016-08-01T22:10:40Z"))
				eventCallbackURL := "https://ngrok.com"
				Expect(resp.EventCallbackURL).To(Equal(&eventCallbackURL))
				Expect(resp.EventsFilter).To(BeNil())
				Expect(resp.FriendlyName).To(Equal("NewWorkspace"))
				Expect(resp.MultiTaskEnabled).To(Equal(false))
				Expect(resp.PrioritizeQueueOrder).To(Equal("FIFO"))
				Expect(resp.TimeoutActivityName).To(Equal("Offline"))
				Expect(resp.TimeoutActivitySid).To(Equal("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1"))
				Expect(resp.Sid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update flow response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &workspace.UpdateWorkspaceInput{
				FriendlyName: "Test Workspace",
			}

			resp, err := taskrouterSession.Workspace("WS71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update workspace response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the workspace is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := workspaceClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete workspace response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WS71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := taskrouterSession.Workspace("WS71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given the Task Queues Client", func() {
		taskQueuesClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskQueues

		Describe("When the Task Queue is successfully created", func() {
			createInput := &task_queues.CreateTaskQueueInput{
				FriendlyName: "Test 2",
			}

			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/taskQueueResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := taskQueuesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create task queue response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssignmentActivityName).To(Equal("817ca1c5-3a05-11e5-9292-98e0d9a1eb73"))
				Expect(resp.AssignmentActivitySid).To(Equal("WA21d51f4c72583766988f9860de3e130a"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-08-04T01:31:41Z"))
				Expect(resp.FriendlyName).To(Equal("English"))
				Expect(resp.MaxReservedWorkers).To(Equal(1))
				Expect(resp.ReservationActivityName).To(Equal("80fa2beb-3a05-11e5-8fc8-98e0d9a1eb74"))
				Expect(resp.ReservationActivitySid).To(Equal("WAea296a56ebce4bfbff0e99abadf16934"))
				Expect(resp.Sid).To(Equal("WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TargetWorkers).To(Equal("languages HAS \"english\""))
				Expect(resp.TaskOrder).To(Equal("FIFO"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})
	})

	Describe("Given I have a Task Queue SID", func() {
		taskQueueClient := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskQueue("WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the Task Queue is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/taskQueueResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := taskQueueClient.Get()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get task queue response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssignmentActivityName).To(Equal("817ca1c5-3a05-11e5-9292-98e0d9a1eb73"))
				Expect(resp.AssignmentActivitySid).To(Equal("WA21d51f4c72583766988f9860de3e130a"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-08-04T01:31:41Z"))
				Expect(resp.FriendlyName).To(Equal("English"))
				Expect(resp.MaxReservedWorkers).To(Equal(1))
				Expect(resp.ReservationActivityName).To(Equal("80fa2beb-3a05-11e5-8fc8-98e0d9a1eb74"))
				Expect(resp.ReservationActivitySid).To(Equal("WAea296a56ebce4bfbff0e99abadf16934"))
				Expect(resp.Sid).To(Equal("WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TargetWorkers).To(Equal("languages HAS \"english\""))
				Expect(resp.TaskOrder).To(Equal("FIFO"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the get workspace response returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQ71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskQueue("WQ71").Get()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get task queue response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the Workspace is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updatedTaskQueueResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &task_queue.UpdateTaskQueueInput{}

			resp, err := taskQueueClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update workspace response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AssignmentActivityName).To(Equal("817ca1c5-3a05-11e5-9292-98e0d9a1eb73"))
				Expect(resp.AssignmentActivitySid).To(Equal("WA21d51f4c72583766988f9860de3e130a"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2015-08-04T02:31:41Z"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2015-08-04T01:31:41Z"))
				Expect(resp.FriendlyName).To(Equal("English"))
				Expect(resp.MaxReservedWorkers).To(Equal(1))
				Expect(resp.ReservationActivityName).To(Equal("80fa2beb-3a05-11e5-8fc8-98e0d9a1eb74"))
				Expect(resp.ReservationActivitySid).To(Equal("WAea296a56ebce4bfbff0e99abadf16934"))
				Expect(resp.Sid).To(Equal("WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.TargetWorkers).To(Equal("languages HAS \"english\""))
				Expect(resp.TaskOrder).To(Equal("FIFO"))
				Expect(resp.URL).To(Equal("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.WorkspaceSid).To(Equal("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update flow response returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQ71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &task_queue.UpdateTaskQueueInput{
				FriendlyName: "Test Queue",
			}

			resp, err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskQueue("WQ71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update task queue response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the task is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := taskQueueClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the delete workspace response returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQ71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := taskrouterSession.Workspace("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").TaskQueue("WQ71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})
})

func ExpectInvalidInputError(err error) {
	ExpectErrorToNotBeATwilioError(err)
	Expect(err.Error()).To(Equal("Invalid input supplied"))
}

func ExpectNotFoundError(err error) {
	Expect(err).ToNot(BeNil())
	twilioErr, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(true))

	code := 20404
	Expect(twilioErr.Code).To(Equal(&code))
	Expect(twilioErr.Message).To(Equal("The requested resource /Flows/FW71 was not found"))

	moreInfo := "https://www.twilio.com/docs/errors/20404"
	Expect(twilioErr.MoreInfo).To(Equal(&moreInfo))
	Expect(twilioErr.Status).To(Equal(404))
}

func ExpectErrorToNotBeATwilioError(err error) {
	Expect(err).ToNot(BeNil())
	_, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(false))
}