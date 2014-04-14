package fix50sp2

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

//NewNetworkCounterpartySystemStatusResponseBuilder returns an initialized NetworkCounterpartySystemStatusResponseBuilder with specified required fields.
func NewNetworkCounterpartySystemStatusResponseBuilder(
	networkstatusresponsetype field.NetworkStatusResponseType,
	networkresponseid field.NetworkResponseID,
	nocompids field.NoCompIDs) *NetworkCounterpartySystemStatusResponseBuilder {
	builder := new(NetworkCounterpartySystemStatusResponseBuilder)
	builder.Body.Set(networkstatusresponsetype)
	builder.Body.Set(networkresponseid)
	builder.Body.Set(nocompids)
	return builder
}

//NetworkStatusResponseType is a required field for NetworkCounterpartySystemStatusResponse.
func (m *NetworkCounterpartySystemStatusResponse) NetworkStatusResponseType() (*field.NetworkStatusResponseType, error) {
	f := new(field.NetworkStatusResponseType)
	err := m.Body.Get(f)
	return f, err
}

//NetworkRequestID is a non-required field for NetworkCounterpartySystemStatusResponse.
func (m *NetworkCounterpartySystemStatusResponse) NetworkRequestID() (*field.NetworkRequestID, error) {
	f := new(field.NetworkRequestID)
	err := m.Body.Get(f)
	return f, err
}

//NetworkResponseID is a required field for NetworkCounterpartySystemStatusResponse.
func (m *NetworkCounterpartySystemStatusResponse) NetworkResponseID() (*field.NetworkResponseID, error) {
	f := new(field.NetworkResponseID)
	err := m.Body.Get(f)
	return f, err
}

//LastNetworkResponseID is a non-required field for NetworkCounterpartySystemStatusResponse.
func (m *NetworkCounterpartySystemStatusResponse) LastNetworkResponseID() (*field.LastNetworkResponseID, error) {
	f := new(field.LastNetworkResponseID)
	err := m.Body.Get(f)
	return f, err
}

//NoCompIDs is a required field for NetworkCounterpartySystemStatusResponse.
func (m *NetworkCounterpartySystemStatusResponse) NoCompIDs() (*field.NoCompIDs, error) {
	f := new(field.NoCompIDs)
	err := m.Body.Get(f)
	return f, err
}
