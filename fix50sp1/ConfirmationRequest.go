package fix50sp1

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type ConfirmationRequest struct {
	quickfix.Message
}

func (m *ConfirmationRequest) ConfirmReqID() (*field.ConfirmReqID, error) {
	f := new(field.ConfirmReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ConfirmationRequest) NoOrders() (*field.NoOrders, error) {
	f := new(field.NoOrders)
	err := m.Body.Get(f)
	return f, err
}
func (m *ConfirmationRequest) SecondaryAllocID() (*field.SecondaryAllocID, error) {
	f := new(field.SecondaryAllocID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ConfirmationRequest) IndividualAllocID() (*field.IndividualAllocID, error) {
	f := new(field.IndividualAllocID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ConfirmationRequest) AllocAcctIDSource() (*field.AllocAcctIDSource, error) {
	f := new(field.AllocAcctIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *ConfirmationRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *ConfirmationRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *ConfirmationRequest) ConfirmType() (*field.ConfirmType, error) {
	f := new(field.ConfirmType)
	err := m.Body.Get(f)
	return f, err
}
func (m *ConfirmationRequest) AllocID() (*field.AllocID, error) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ConfirmationRequest) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *ConfirmationRequest) AllocAccount() (*field.AllocAccount, error) {
	f := new(field.AllocAccount)
	err := m.Body.Get(f)
	return f, err
}
func (m *ConfirmationRequest) AllocAccountType() (*field.AllocAccountType, error) {
	f := new(field.AllocAccountType)
	err := m.Body.Get(f)
	return f, err
}
func (m *ConfirmationRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
