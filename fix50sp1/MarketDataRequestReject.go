package fix50sp1

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type MarketDataRequestReject struct {
	quickfix.Message
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
func (m *MarketDataRequestReject) NoAltMDSource() (*field.NoAltMDSource, error) {
	f := new(field.NoAltMDSource)
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
func (m *MarketDataRequestReject) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}
