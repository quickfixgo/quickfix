package fix43

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type MarketDataRequestReject struct {
	quickfixgo.Message
}

func (m *MarketDataRequestReject) MDReqID() (*field.MDReqID, error) {
	f := new(field.MDReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataRequestReject) MDReqRejReason() (*field.MDReqRejReason, error) {
	f := new(field.MDReqRejReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataRequestReject) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataRequestReject) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataRequestReject) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
