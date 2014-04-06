package fix50

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type NetworkCounterpartySystemStatusRequest struct {
	quickfix.Message
}

func (m *NetworkCounterpartySystemStatusRequest) NetworkRequestType() (*field.NetworkRequestType, error) {
	f := new(field.NetworkRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NetworkCounterpartySystemStatusRequest) NetworkRequestID() (*field.NetworkRequestID, error) {
	f := new(field.NetworkRequestID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NetworkCounterpartySystemStatusRequest) NoCompIDs() (*field.NoCompIDs, error) {
	f := new(field.NoCompIDs)
	err := m.Body.Get(f)
	return f, err
}
