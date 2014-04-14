package fixt11

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Heartbeat msg type = 0.
type Heartbeat struct {
	message.Message
}

//HeartbeatBuilder builds Heartbeat messages.
type HeartbeatBuilder struct {
	message.MessageBuilder
}

//NewHeartbeatBuilder returns an initialized HeartbeatBuilder with specified required fields.
func NewHeartbeatBuilder() *HeartbeatBuilder {
	builder := new(HeartbeatBuilder)
	return builder
}

//TestReqID is a non-required field for Heartbeat.
func (m *Heartbeat) TestReqID() (*field.TestReqID, error) {
	f := new(field.TestReqID)
	err := m.Body.Get(f)
	return f, err
}
