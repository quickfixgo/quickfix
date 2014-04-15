package fix44

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

//CreateNetworkCounterpartySystemStatusRequestBuilder returns an initialized NetworkCounterpartySystemStatusRequestBuilder with specified required fields.
func CreateNetworkCounterpartySystemStatusRequestBuilder(
	networkrequesttype field.NetworkRequestType,
	networkrequestid field.NetworkRequestID) NetworkCounterpartySystemStatusRequestBuilder {
	var builder NetworkCounterpartySystemStatusRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(networkrequesttype)
	builder.Body.Set(networkrequestid)
	return builder
}

//NetworkRequestType is a required field for NetworkCounterpartySystemStatusRequest.
func (m NetworkCounterpartySystemStatusRequest) NetworkRequestType() (field.NetworkRequestType, error) {
	var f field.NetworkRequestType
	err := m.Body.Get(&f)
	return f, err
}

//NetworkRequestID is a required field for NetworkCounterpartySystemStatusRequest.
func (m NetworkCounterpartySystemStatusRequest) NetworkRequestID() (field.NetworkRequestID, error) {
	var f field.NetworkRequestID
	err := m.Body.Get(&f)
	return f, err
}

//NoCompIDs is a non-required field for NetworkCounterpartySystemStatusRequest.
func (m NetworkCounterpartySystemStatusRequest) NoCompIDs() (field.NoCompIDs, error) {
	var f field.NoCompIDs
	err := m.Body.Get(&f)
	return f, err
}
