package fix43

import (
	"github.com/quickfixgo/quickfix/errors"
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

//CreateHeartbeatBuilder returns an initialized HeartbeatBuilder with specified required fields.
func CreateHeartbeatBuilder() HeartbeatBuilder {
	var builder HeartbeatBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	return builder
}

//TestReqID is a non-required field for Heartbeat.
func (m Heartbeat) TestReqID() (field.TestReqID, errors.MessageRejectError) {
	var f field.TestReqID
	err := m.Body.Get(&f)
	return f, err
}
