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
	networkstatusresponsetype field.NetworkStatusResponseType,
	networkresponseid field.NetworkResponseID,
	nocompids field.NoCompIDs) NetworkCounterpartySystemStatusResponseBuilder {
	var builder NetworkCounterpartySystemStatusResponseBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildDefaultApplVerID(enum.ApplVerID_FIX50))
	builder.Header.Set(field.BuildMsgType("BD"))
	builder.Body.Set(networkstatusresponsetype)
	builder.Body.Set(networkresponseid)
	builder.Body.Set(nocompids)
	return builder
}

//NetworkStatusResponseType is a required field for NetworkCounterpartySystemStatusResponse.
func (m NetworkCounterpartySystemStatusResponse) NetworkStatusResponseType() (*field.NetworkStatusResponseType, errors.MessageRejectError) {
	f := new(field.NetworkStatusResponseType)
	err := m.Body.Get(f)
	return f, err
}

//GetNetworkStatusResponseType reads a NetworkStatusResponseType from NetworkCounterpartySystemStatusResponse.
func (m NetworkCounterpartySystemStatusResponse) GetNetworkStatusResponseType(f *field.NetworkStatusResponseType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NetworkRequestID is a non-required field for NetworkCounterpartySystemStatusResponse.
func (m NetworkCounterpartySystemStatusResponse) NetworkRequestID() (*field.NetworkRequestID, errors.MessageRejectError) {
	f := new(field.NetworkRequestID)
	err := m.Body.Get(f)
	return f, err
}

//GetNetworkRequestID reads a NetworkRequestID from NetworkCounterpartySystemStatusResponse.
func (m NetworkCounterpartySystemStatusResponse) GetNetworkRequestID(f *field.NetworkRequestID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NetworkResponseID is a required field for NetworkCounterpartySystemStatusResponse.
func (m NetworkCounterpartySystemStatusResponse) NetworkResponseID() (*field.NetworkResponseID, errors.MessageRejectError) {
	f := new(field.NetworkResponseID)
	err := m.Body.Get(f)
	return f, err
}

//GetNetworkResponseID reads a NetworkResponseID from NetworkCounterpartySystemStatusResponse.
func (m NetworkCounterpartySystemStatusResponse) GetNetworkResponseID(f *field.NetworkResponseID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastNetworkResponseID is a non-required field for NetworkCounterpartySystemStatusResponse.
func (m NetworkCounterpartySystemStatusResponse) LastNetworkResponseID() (*field.LastNetworkResponseID, errors.MessageRejectError) {
	f := new(field.LastNetworkResponseID)
	err := m.Body.Get(f)
	return f, err
}

//GetLastNetworkResponseID reads a LastNetworkResponseID from NetworkCounterpartySystemStatusResponse.
func (m NetworkCounterpartySystemStatusResponse) GetLastNetworkResponseID(f *field.LastNetworkResponseID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoCompIDs is a required field for NetworkCounterpartySystemStatusResponse.
func (m NetworkCounterpartySystemStatusResponse) NoCompIDs() (*field.NoCompIDs, errors.MessageRejectError) {
	f := new(field.NoCompIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoCompIDs reads a NoCompIDs from NetworkCounterpartySystemStatusResponse.
func (m NetworkCounterpartySystemStatusResponse) GetNoCompIDs(f *field.NoCompIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}
