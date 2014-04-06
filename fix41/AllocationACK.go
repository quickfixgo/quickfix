package fix41

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type AllocationACK struct {
	quickfix.Message
}

func (m *AllocationACK) ClientID() (*field.ClientID, error) {
	f := new(field.ClientID)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationACK) ExecBroker() (*field.ExecBroker, error) {
	f := new(field.ExecBroker)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationACK) AllocID() (*field.AllocID, error) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationACK) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationACK) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationACK) AllocStatus() (*field.AllocStatus, error) {
	f := new(field.AllocStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationACK) AllocRejCode() (*field.AllocRejCode, error) {
	f := new(field.AllocRejCode)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationACK) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
