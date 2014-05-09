package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//NetworkCounterpartySystemStatusRequest msg type = BC.
type NetworkCounterpartySystemStatusRequest struct {
	message.Message
}

//NetworkCounterpartySystemStatusRequestBuilder builds NetworkCounterpartySystemStatusRequest messages.
type NetworkCounterpartySystemStatusRequestBuilder struct {
	message.MessageBuilder
}

//CreateNetworkCounterpartySystemStatusRequestBuilder returns an initialized NetworkCounterpartySystemStatusRequestBuilder with specified required fields.
func CreateNetworkCounterpartySystemStatusRequestBuilder(
	networkrequesttype *field.NetworkRequestTypeField,
	networkrequestid *field.NetworkRequestIDField) NetworkCounterpartySystemStatusRequestBuilder {
	var builder NetworkCounterpartySystemStatusRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("BC"))
	builder.Body.Set(networkrequesttype)
	builder.Body.Set(networkrequestid)
	return builder
}

//NetworkRequestType is a required field for NetworkCounterpartySystemStatusRequest.
func (m NetworkCounterpartySystemStatusRequest) NetworkRequestType() (*field.NetworkRequestTypeField, errors.MessageRejectError) {
	f := &field.NetworkRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNetworkRequestType reads a NetworkRequestType from NetworkCounterpartySystemStatusRequest.
func (m NetworkCounterpartySystemStatusRequest) GetNetworkRequestType(f *field.NetworkRequestTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NetworkRequestID is a required field for NetworkCounterpartySystemStatusRequest.
func (m NetworkCounterpartySystemStatusRequest) NetworkRequestID() (*field.NetworkRequestIDField, errors.MessageRejectError) {
	f := &field.NetworkRequestIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNetworkRequestID reads a NetworkRequestID from NetworkCounterpartySystemStatusRequest.
func (m NetworkCounterpartySystemStatusRequest) GetNetworkRequestID(f *field.NetworkRequestIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoCompIDs is a non-required field for NetworkCounterpartySystemStatusRequest.
func (m NetworkCounterpartySystemStatusRequest) NoCompIDs() (*field.NoCompIDsField, errors.MessageRejectError) {
	f := &field.NoCompIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoCompIDs reads a NoCompIDs from NetworkCounterpartySystemStatusRequest.
func (m NetworkCounterpartySystemStatusRequest) GetNoCompIDs(f *field.NoCompIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}
