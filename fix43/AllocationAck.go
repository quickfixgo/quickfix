package fix43

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type AllocationAck struct {
	quickfixgo.Message
}

func (m *AllocationAck) AllocID() (*field.AllocID, error) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationAck) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationAck) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationAck) LegalConfirm() (*field.LegalConfirm, error) {
	f := new(field.LegalConfirm)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationAck) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationAck) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationAck) AllocStatus() (*field.AllocStatus, error) {
	f := new(field.AllocStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationAck) AllocRejCode() (*field.AllocRejCode, error) {
	f := new(field.AllocRejCode)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationAck) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationAck) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
