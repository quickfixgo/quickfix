package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

type NetworkCounterpartySystemStatusRequest struct {
	message.Message
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
