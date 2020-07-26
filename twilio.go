package twilio

import (
	"github.com/RJPearson94/twilio-sdk-go/service/chat"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations"
	"github.com/RJPearson94/twilio-sdk-go/service/fax"
	"github.com/RJPearson94/twilio-sdk-go/service/flex"
	"github.com/RJPearson94/twilio-sdk-go/service/messaging"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless"
	"github.com/RJPearson94/twilio-sdk-go/service/studio"
	"github.com/RJPearson94/twilio-sdk-go/service/sync"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/twiml"
)

type Twilio struct {
	Chat          *chat.Chat
	Conversations *conversations.Conversations
	Fax           *fax.Fax
	Flex          *flex.Flex
	Messaging     *messaging.Messaging
	Proxy         *proxy.Proxy
	Serverless    *serverless.Serverless
	Studio        *studio.Studio
	Sync          *sync.Sync
	TaskRouter    *taskrouter.TaskRouter
	TwiML         *twiml.TwiML
}

func New(sess *session.Session) *Twilio {
	c := &Twilio{}
	c.Chat = chat.New(sess)
	c.Conversations = conversations.New(sess)
	c.Fax = fax.New(sess)
	c.Flex = flex.New(sess)
	c.Messaging = messaging.New(sess)
	c.Proxy = proxy.New(sess)
	c.Serverless = serverless.New(sess)
	c.Studio = studio.New(sess)
	c.Sync = sync.New(sess)
	c.TaskRouter = taskrouter.New(sess)
	c.TwiML = twiml.New()
	return c
}

func NewWithCredentials(creds *credentials.Credentials) *Twilio {
	return New(session.New(creds))
}
