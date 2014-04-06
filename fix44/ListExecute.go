package fix44

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type ListExecute struct {
	quickfix.Message
}

func (m *ListExecute) ListID() (*field.ListID, error) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListExecute) ClientBidID() (*field.ClientBidID, error) {
	f := new(field.ClientBidID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListExecute) BidID() (*field.BidID, error) {
	f := new(field.BidID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListExecute) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListExecute) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListExecute) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListExecute) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
