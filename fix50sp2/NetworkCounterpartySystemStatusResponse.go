package fix50sp2

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type NetworkCounterpartySystemStatusResponse struct {
	quickfixgo.Message
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
