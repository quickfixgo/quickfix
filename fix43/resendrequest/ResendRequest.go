//Package resendrequest msg type = 2.
package resendrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix/enum"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a ResendRequest wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//BeginSeqNo is a required field for ResendRequest.
func (m Message) BeginSeqNo() (*field.BeginSeqNoField, quickfix.MessageRejectError) {
	f := &field.BeginSeqNoField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBeginSeqNo reads a BeginSeqNo from ResendRequest.
func (m Message) GetBeginSeqNo(f *field.BeginSeqNoField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EndSeqNo is a required field for ResendRequest.
func (m Message) EndSeqNo() (*field.EndSeqNoField, quickfix.MessageRejectError) {
	f := &field.EndSeqNoField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEndSeqNo reads a EndSeqNo from ResendRequest.
func (m Message) GetEndSeqNo(f *field.EndSeqNoField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for ResendRequest.
func New(
	beginseqno *field.BeginSeqNoField,
	endseqno *field.EndSeqNoField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIX43))
	builder.Header.Set(field.NewMsgType("2"))
	builder.Body.Set(beginseqno)
	builder.Body.Set(endseqno)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.BeginStringFIX43, "2", r
}
