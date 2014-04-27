package fix50

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
	networkrequesttype field.NetworkRequestType,
	networkrequestid field.NetworkRequestID) NetworkCounterpartySystemStatusRequestBuilder {
	var builder NetworkCounterpartySystemStatusRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildDefaultApplVerID(enum.ApplVerID_FIX50))
	builder.Header.Set(field.BuildMsgType("BC"))
	builder.Body.Set(networkrequesttype)
	builder.Body.Set(networkrequestid)
	return builder
}

//NetworkRequestType is a required field for NetworkCounterpartySystemStatusRequest.
func (m NetworkCounterpartySystemStatusRequest) NetworkRequestType() (*field.NetworkRequestType, errors.MessageRejectError) {
	f := new(field.NetworkRequestType)
	err := m.Body.Get(f)
	return f, err
}

//GetNetworkRequestType reads a NetworkRequestType from NetworkCounterpartySystemStatusRequest.
func (m NetworkCounterpartySystemStatusRequest) GetNetworkRequestType(f *field.NetworkRequestType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NetworkRequestID is a required field for NetworkCounterpartySystemStatusRequest.
func (m NetworkCounterpartySystemStatusRequest) NetworkRequestID() (*field.NetworkRequestID, errors.MessageRejectError) {
	f := new(field.NetworkRequestID)
	err := m.Body.Get(f)
	return f, err
}

//GetNetworkRequestID reads a NetworkRequestID from NetworkCounterpartySystemStatusRequest.
func (m NetworkCounterpartySystemStatusRequest) GetNetworkRequestID(f *field.NetworkRequestID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoCompIDs is a non-required field for NetworkCounterpartySystemStatusRequest.
func (m NetworkCounterpartySystemStatusRequest) NoCompIDs() (*field.NoCompIDs, errors.MessageRejectError) {
	f := new(field.NoCompIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoCompIDs reads a NoCompIDs from NetworkCounterpartySystemStatusRequest.
func (m NetworkCounterpartySystemStatusRequest) GetNoCompIDs(f *field.NoCompIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}
