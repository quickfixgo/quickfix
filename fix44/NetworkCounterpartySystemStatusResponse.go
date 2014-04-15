package fix44

import (
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
	networkstatusresponsetype field.NetworkStatusResponseType,
	networkresponseid field.NetworkResponseID,
	nocompids field.NoCompIDs) NetworkCounterpartySystemStatusResponseBuilder {
	var builder NetworkCounterpartySystemStatusResponseBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(networkstatusresponsetype)
	builder.Body.Set(networkresponseid)
	builder.Body.Set(nocompids)
	return builder
}

//NetworkStatusResponseType is a required field for NetworkCounterpartySystemStatusResponse.
func (m NetworkCounterpartySystemStatusResponse) NetworkStatusResponseType() (field.NetworkStatusResponseType, error) {
	var f field.NetworkStatusResponseType
	err := m.Body.Get(&f)
	return f, err
}

//NetworkRequestID is a non-required field for NetworkCounterpartySystemStatusResponse.
func (m NetworkCounterpartySystemStatusResponse) NetworkRequestID() (field.NetworkRequestID, error) {
	var f field.NetworkRequestID
	err := m.Body.Get(&f)
	return f, err
}

//NetworkResponseID is a required field for NetworkCounterpartySystemStatusResponse.
func (m NetworkCounterpartySystemStatusResponse) NetworkResponseID() (field.NetworkResponseID, error) {
	var f field.NetworkResponseID
	err := m.Body.Get(&f)
	return f, err
}

//LastNetworkResponseID is a non-required field for NetworkCounterpartySystemStatusResponse.
func (m NetworkCounterpartySystemStatusResponse) LastNetworkResponseID() (field.LastNetworkResponseID, error) {
	var f field.LastNetworkResponseID
	err := m.Body.Get(&f)
	return f, err
}

//NoCompIDs is a required field for NetworkCounterpartySystemStatusResponse.
func (m NetworkCounterpartySystemStatusResponse) NoCompIDs() (field.NoCompIDs, error) {
	var f field.NoCompIDs
	err := m.Body.Get(&f)
	return f, err
}
