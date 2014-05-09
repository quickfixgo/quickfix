package fix41

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
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
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX41))
	builder.Header.Set(field.NewMsgType("0"))
	return builder
}

//TestReqID is a non-required field for Heartbeat.
func (m Heartbeat) TestReqID() (*field.TestReqIDField, errors.MessageRejectError) {
	f := &field.TestReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTestReqID reads a TestReqID from Heartbeat.
func (m Heartbeat) GetTestReqID(f *field.TestReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}
