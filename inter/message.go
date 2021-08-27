package inter

import (
	"time"

	"github.com/choria-io/go-choria/protocol"
)

// Message is a message that is transportable over the Choria Protocol
type Message interface {
	Agent() string
	Base64Payload() string
	CacheTransport()
	CallerID() string
	Collective() string
	CustomTarget() string
	DiscoveredHosts() []string
	ExpectedMessageID() string
	Filter() *protocol.Filter
	IsCachedTransport() bool
	NotifyPublish()
	OnPublish(func())
	Payload() string
	ProtocolVersion() string
	ReplyTo() string
	Request() Message
	RequestID() string
	SenderID() string
	SetBase64Payload(payload string) error
	SetCollective(string) error
	SetCustomTarget(string)
	SetDiscoveredHosts(hosts []string)
	SetExpectedMsgID(id string) error
	SetFilter(*protocol.Filter)
	SetPayload(string)
	SetProtocolVersion(string)
	SetReplyTo(string) error
	SetTTL(int)
	SetType(string) error
	String() string
	TTL() int
	TimeStamp() time.Time
	Transport() (protocol.TransportMessage, error)
	Type() string
	Validate() (bool, error)
	ValidateTTL() bool
}