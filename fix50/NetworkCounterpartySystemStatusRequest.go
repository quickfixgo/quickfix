package fix50

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//NetworkCounterpartySystemStatusRequest msg type = BC.
type NetworkCounterpartySystemStatusRequest struct {
	message.Message
}

//NetworkCounterpartySystemStatusRequestBuilder builds NetworkCounterpartySystemStatusRequest messages.
type NetworkCounterpartySystemStatusRequestBuilder struct {
	message.MessageBuilder
}

//NewNetworkCounterpartySystemStatusRequestBuilder returns an initialized NetworkCounterpartySystemStatusRequestBuilder with specified required fields.
func NewNetworkCounterpartySystemStatusRequestBuilder(
	networkrequesttype field.NetworkRequestType,
	networkrequestid field.NetworkRequestID) *NetworkCounterpartySystemStatusRequestBuilder {
	builder := new(NetworkCounterpartySystemStatusRequestBuilder)
	builder.Body.Set(networkrequesttype)
	builder.Body.Set(networkrequestid)
	return builder
}

//NetworkRequestType is a required field for NetworkCounterpartySystemStatusRequest.
func (m *NetworkCounterpartySystemStatusRequest) NetworkRequestType() (*field.NetworkRequestType, error) {
	f := new(field.NetworkRequestType)
	err := m.Body.Get(f)
	return f, err
}

//NetworkRequestID is a required field for NetworkCounterpartySystemStatusRequest.
func (m *NetworkCounterpartySystemStatusRequest) NetworkRequestID() (*field.NetworkRequestID, error) {
	f := new(field.NetworkRequestID)
	err := m.Body.Get(f)
	return f, err
}

//NoCompIDs is a non-required field for NetworkCounterpartySystemStatusRequest.
func (m *NetworkCounterpartySystemStatusRequest) NoCompIDs() (*field.NoCompIDs, error) {
	f := new(field.NoCompIDs)
	err := m.Body.Get(f)
	return f, err
}
