//Package networkcounterpartysystemstatusrequest msg type = BC.
package networkcounterpartysystemstatusrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a NetworkCounterpartySystemStatusRequest wrapper for the generic Message type
type Message struct {
	message.Message
}

//NetworkRequestType is a required field for NetworkCounterpartySystemStatusRequest.
func (m Message) NetworkRequestType() (*field.NetworkRequestTypeField, errors.MessageRejectError) {
	f := &field.NetworkRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNetworkRequestType reads a NetworkRequestType from NetworkCounterpartySystemStatusRequest.
func (m Message) GetNetworkRequestType(f *field.NetworkRequestTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NetworkRequestID is a required field for NetworkCounterpartySystemStatusRequest.
func (m Message) NetworkRequestID() (*field.NetworkRequestIDField, errors.MessageRejectError) {
	f := &field.NetworkRequestIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNetworkRequestID reads a NetworkRequestID from NetworkCounterpartySystemStatusRequest.
func (m Message) GetNetworkRequestID(f *field.NetworkRequestIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoCompIDs is a non-required field for NetworkCounterpartySystemStatusRequest.
func (m Message) NoCompIDs() (*field.NoCompIDsField, errors.MessageRejectError) {
	f := &field.NoCompIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoCompIDs reads a NoCompIDs from NetworkCounterpartySystemStatusRequest.
func (m Message) GetNoCompIDs(f *field.NoCompIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds NetworkCounterpartySystemStatusRequest messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for NetworkCounterpartySystemStatusRequest.
func Builder(
	networkrequesttype *field.NetworkRequestTypeField,
	networkrequestid *field.NetworkRequestIDField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header().Set(field.NewMsgType("BC"))
	builder.Body().Set(networkrequesttype)
	builder.Body().Set(networkrequestid)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP1, "BC", r
}
