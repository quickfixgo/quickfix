package fix44

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//NetworkCounterpartySystemStatusResponse msg type = BD.
type NetworkCounterpartySystemStatusResponse struct {
	message.Message
}

//NetworkCounterpartySystemStatusResponseBuilder builds NetworkCounterpartySystemStatusResponse messages.
type NetworkCounterpartySystemStatusResponseBuilder struct {
	message.MessageBuilder
}

//CreateNetworkCounterpartySystemStatusResponseBuilder returns an initialized NetworkCounterpartySystemStatusResponseBuilder with specified required fields.
func CreateNetworkCounterpartySystemStatusResponseBuilder(
	networkstatusresponsetype *field.NetworkStatusResponseTypeField,
	networkresponseid *field.NetworkResponseIDField,
	nocompids *field.NoCompIDsField) NetworkCounterpartySystemStatusResponseBuilder {
	var builder NetworkCounterpartySystemStatusResponseBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header().Set(field.NewMsgType("BD"))
	builder.Body().Set(networkstatusresponsetype)
	builder.Body().Set(networkresponseid)
	builder.Body().Set(nocompids)
	return builder
}

//NetworkStatusResponseType is a required field for NetworkCounterpartySystemStatusResponse.
func (m NetworkCounterpartySystemStatusResponse) NetworkStatusResponseType() (*field.NetworkStatusResponseTypeField, errors.MessageRejectError) {
	f := &field.NetworkStatusResponseTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNetworkStatusResponseType reads a NetworkStatusResponseType from NetworkCounterpartySystemStatusResponse.
func (m NetworkCounterpartySystemStatusResponse) GetNetworkStatusResponseType(f *field.NetworkStatusResponseTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NetworkRequestID is a non-required field for NetworkCounterpartySystemStatusResponse.
func (m NetworkCounterpartySystemStatusResponse) NetworkRequestID() (*field.NetworkRequestIDField, errors.MessageRejectError) {
	f := &field.NetworkRequestIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNetworkRequestID reads a NetworkRequestID from NetworkCounterpartySystemStatusResponse.
func (m NetworkCounterpartySystemStatusResponse) GetNetworkRequestID(f *field.NetworkRequestIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NetworkResponseID is a required field for NetworkCounterpartySystemStatusResponse.
func (m NetworkCounterpartySystemStatusResponse) NetworkResponseID() (*field.NetworkResponseIDField, errors.MessageRejectError) {
	f := &field.NetworkResponseIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNetworkResponseID reads a NetworkResponseID from NetworkCounterpartySystemStatusResponse.
func (m NetworkCounterpartySystemStatusResponse) GetNetworkResponseID(f *field.NetworkResponseIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastNetworkResponseID is a non-required field for NetworkCounterpartySystemStatusResponse.
func (m NetworkCounterpartySystemStatusResponse) LastNetworkResponseID() (*field.LastNetworkResponseIDField, errors.MessageRejectError) {
	f := &field.LastNetworkResponseIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastNetworkResponseID reads a LastNetworkResponseID from NetworkCounterpartySystemStatusResponse.
func (m NetworkCounterpartySystemStatusResponse) GetLastNetworkResponseID(f *field.LastNetworkResponseIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoCompIDs is a required field for NetworkCounterpartySystemStatusResponse.
func (m NetworkCounterpartySystemStatusResponse) NoCompIDs() (*field.NoCompIDsField, errors.MessageRejectError) {
	f := &field.NoCompIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoCompIDs reads a NoCompIDs from NetworkCounterpartySystemStatusResponse.
func (m NetworkCounterpartySystemStatusResponse) GetNoCompIDs(f *field.NoCompIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}
