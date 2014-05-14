//Package testrequest msg type = 1.
package testrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Message is a TestRequest wrapper for the generic Message type
type Message struct {
	message.Message
}

//TestReqID is a required field for TestRequest.
func (m Message) TestReqID() (*field.TestReqIDField, errors.MessageRejectError) {
	f := &field.TestReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTestReqID reads a TestReqID from TestRequest.
func (m Message) GetTestReqID(f *field.TestReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds TestRequest messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for TestRequest.
func Builder(
	testreqid *field.TestReqIDField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewMsgType("1"))
	builder.Body().Set(testreqid)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIXT11, "1", r
}
