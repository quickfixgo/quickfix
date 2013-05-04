package fix50sp2

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type PartyDetailsListRequest struct {
	quickfixgo.Message
}

func (m *PartyDetailsListRequest) PartyDetailsListRequestID() (*field.PartyDetailsListRequestID, error) {
	f := new(field.PartyDetailsListRequestID)
	err := m.Body.Get(f)
	return f, err
}
func (m *PartyDetailsListRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *PartyDetailsListRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *PartyDetailsListRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *PartyDetailsListRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
