//Package networkcounterpartysystemstatusresponse msg type = BD.
package networkcounterpartysystemstatusresponse

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a NetworkCounterpartySystemStatusResponse wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//NetworkStatusResponseType is a required field for NetworkCounterpartySystemStatusResponse.
func (m Message) NetworkStatusResponseType() (*field.NetworkStatusResponseTypeField, errors.MessageRejectError) {
	f := &field.NetworkStatusResponseTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNetworkStatusResponseType reads a NetworkStatusResponseType from NetworkCounterpartySystemStatusResponse.
func (m Message) GetNetworkStatusResponseType(f *field.NetworkStatusResponseTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NetworkRequestID is a non-required field for NetworkCounterpartySystemStatusResponse.
func (m Message) NetworkRequestID() (*field.NetworkRequestIDField, errors.MessageRejectError) {
	f := &field.NetworkRequestIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNetworkRequestID reads a NetworkRequestID from NetworkCounterpartySystemStatusResponse.
func (m Message) GetNetworkRequestID(f *field.NetworkRequestIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NetworkResponseID is a required field for NetworkCounterpartySystemStatusResponse.
func (m Message) NetworkResponseID() (*field.NetworkResponseIDField, errors.MessageRejectError) {
	f := &field.NetworkResponseIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNetworkResponseID reads a NetworkResponseID from NetworkCounterpartySystemStatusResponse.
func (m Message) GetNetworkResponseID(f *field.NetworkResponseIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastNetworkResponseID is a non-required field for NetworkCounterpartySystemStatusResponse.
func (m Message) LastNetworkResponseID() (*field.LastNetworkResponseIDField, errors.MessageRejectError) {
	f := &field.LastNetworkResponseIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastNetworkResponseID reads a LastNetworkResponseID from NetworkCounterpartySystemStatusResponse.
func (m Message) GetLastNetworkResponseID(f *field.LastNetworkResponseIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoCompIDs is a required field for NetworkCounterpartySystemStatusResponse.
func (m Message) NoCompIDs() (*field.NoCompIDsField, errors.MessageRejectError) {
	f := &field.NoCompIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoCompIDs reads a NoCompIDs from NetworkCounterpartySystemStatusResponse.
func (m Message) GetNoCompIDs(f *field.NoCompIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds NetworkCounterpartySystemStatusResponse messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for NetworkCounterpartySystemStatusResponse.
func Builder(
	networkstatusresponsetype *field.NetworkStatusResponseTypeField,
	networkresponseid *field.NetworkResponseIDField,
	nocompids *field.NoCompIDsField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = quickfix.NewMessageBuilder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50))
	builder.Header().Set(field.NewMsgType("BD"))
	builder.Body().Set(networkstatusresponsetype)
	builder.Body().Set(networkresponseid)
	builder.Body().Set(nocompids)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX50, "BD", r
}
