//Package heartbeat msg type = 0.
package heartbeat

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
)

//Message is a Heartbeat wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//TestReqID is a non-required field for Heartbeat.
func (m Message) TestReqID() (*field.TestReqIDField, quickfix.MessageRejectError) {
	f := &field.TestReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTestReqID reads a TestReqID from Heartbeat.
func (m Message) GetTestReqID(f *field.TestReqIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for Heartbeat.
func New() Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIX44))
	builder.Header.Set(field.NewMsgType("0"))
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.BeginStringFIX44, "0", r
}
