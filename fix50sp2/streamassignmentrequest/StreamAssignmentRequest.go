//Package streamassignmentrequest msg type = CC.
package streamassignmentrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a StreamAssignmentRequest wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//StreamAsgnReqID is a required field for StreamAssignmentRequest.
func (m Message) StreamAsgnReqID() (*field.StreamAsgnReqIDField, quickfix.MessageRejectError) {
	f := &field.StreamAsgnReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnReqID reads a StreamAsgnReqID from StreamAssignmentRequest.
func (m Message) GetStreamAsgnReqID(f *field.StreamAsgnReqIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StreamAsgnReqType is a required field for StreamAssignmentRequest.
func (m Message) StreamAsgnReqType() (*field.StreamAsgnReqTypeField, quickfix.MessageRejectError) {
	f := &field.StreamAsgnReqTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnReqType reads a StreamAsgnReqType from StreamAssignmentRequest.
func (m Message) GetStreamAsgnReqType(f *field.StreamAsgnReqTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoAsgnReqs is a non-required field for StreamAssignmentRequest.
func (m Message) NoAsgnReqs() (*field.NoAsgnReqsField, quickfix.MessageRejectError) {
	f := &field.NoAsgnReqsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoAsgnReqs reads a NoAsgnReqs from StreamAssignmentRequest.
func (m Message) GetNoAsgnReqs(f *field.NoAsgnReqsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for StreamAssignmentRequest.
func New(
	streamasgnreqid *field.StreamAsgnReqIDField,
	streamasgnreqtype *field.StreamAsgnReqTypeField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("CC"))
	builder.Body.Set(streamasgnreqid)
	builder.Body.Set(streamasgnreqtype)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "CC", r
}
