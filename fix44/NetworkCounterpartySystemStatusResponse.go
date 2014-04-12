package fix44

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

type NetworkCounterpartySystemStatusResponse struct {
	message.Message
}

func (m *NetworkCounterpartySystemStatusResponse) NetworkStatusResponseType() (*field.NetworkStatusResponseType, error) {
	f := new(field.NetworkStatusResponseType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NetworkCounterpartySystemStatusResponse) NetworkRequestID() (*field.NetworkRequestID, error) {
	f := new(field.NetworkRequestID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NetworkCounterpartySystemStatusResponse) NetworkResponseID() (*field.NetworkResponseID, error) {
	f := new(field.NetworkResponseID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NetworkCounterpartySystemStatusResponse) LastNetworkResponseID() (*field.LastNetworkResponseID, error) {
	f := new(field.LastNetworkResponseID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NetworkCounterpartySystemStatusResponse) NoCompIDs() (*field.NoCompIDs, error) {
	f := new(field.NoCompIDs)
	err := m.Body.Get(f)
	return f, err
}
