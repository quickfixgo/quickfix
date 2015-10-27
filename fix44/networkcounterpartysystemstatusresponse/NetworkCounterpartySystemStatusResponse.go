//Package networkcounterpartysystemstatusresponse msg type = BD.
package networkcounterpartysystemstatusresponse

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
)

//Message is a NetworkCounterpartySystemStatusResponse wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//NetworkStatusResponseType is a required field for NetworkCounterpartySystemStatusResponse.
func (m Message) NetworkStatusResponseType() (*field.NetworkStatusResponseTypeField, quickfix.MessageRejectError) {
	f := &field.NetworkStatusResponseTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNetworkStatusResponseType reads a NetworkStatusResponseType from NetworkCounterpartySystemStatusResponse.
func (m Message) GetNetworkStatusResponseType(f *field.NetworkStatusResponseTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NetworkRequestID is a non-required field for NetworkCounterpartySystemStatusResponse.
func (m Message) NetworkRequestID() (*field.NetworkRequestIDField, quickfix.MessageRejectError) {
	f := &field.NetworkRequestIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNetworkRequestID reads a NetworkRequestID from NetworkCounterpartySystemStatusResponse.
func (m Message) GetNetworkRequestID(f *field.NetworkRequestIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NetworkResponseID is a required field for NetworkCounterpartySystemStatusResponse.
func (m Message) NetworkResponseID() (*field.NetworkResponseIDField, quickfix.MessageRejectError) {
	f := &field.NetworkResponseIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNetworkResponseID reads a NetworkResponseID from NetworkCounterpartySystemStatusResponse.
func (m Message) GetNetworkResponseID(f *field.NetworkResponseIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastNetworkResponseID is a non-required field for NetworkCounterpartySystemStatusResponse.
func (m Message) LastNetworkResponseID() (*field.LastNetworkResponseIDField, quickfix.MessageRejectError) {
	f := &field.LastNetworkResponseIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastNetworkResponseID reads a LastNetworkResponseID from NetworkCounterpartySystemStatusResponse.
func (m Message) GetLastNetworkResponseID(f *field.LastNetworkResponseIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoCompIDs is a required field for NetworkCounterpartySystemStatusResponse.
func (m Message) NoCompIDs() (*field.NoCompIDsField, quickfix.MessageRejectError) {
	f := &field.NoCompIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoCompIDs reads a NoCompIDs from NetworkCounterpartySystemStatusResponse.
func (m Message) GetNoCompIDs(f *field.NoCompIDsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for NetworkCounterpartySystemStatusResponse.
func New(
	networkstatusresponsetype *field.NetworkStatusResponseTypeField,
	networkresponseid *field.NetworkResponseIDField,
	nocompids *field.NoCompIDsField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIX44))
	builder.Header.Set(field.NewMsgType("BD"))
	builder.Body.Set(networkstatusresponsetype)
	builder.Body.Set(networkresponseid)
	builder.Body.Set(nocompids)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.BeginStringFIX44, "BD", r
}
