package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

type PartyDetailsListRequest struct {
	message.Message
}

func (m *PartyDetailsListRequest) PartyDetailsListRequestID() (*field.PartyDetailsListRequestID, error) {
	f := new(field.PartyDetailsListRequestID)
	err := m.Body.Get(f)
	return f, err
}
func (m *PartyDetailsListRequest) NoPartyListResponseTypes() (*field.NoPartyListResponseTypes, error) {
	f := new(field.NoPartyListResponseTypes)
	err := m.Body.Get(f)
	return f, err
}
func (m *PartyDetailsListRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *PartyDetailsListRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *PartyDetailsListRequest) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}
func (m *PartyDetailsListRequest) NoRequestedPartyRoles() (*field.NoRequestedPartyRoles, error) {
	f := new(field.NoRequestedPartyRoles)
	err := m.Body.Get(f)
	return f, err
}
func (m *PartyDetailsListRequest) NoPartyRelationships() (*field.NoPartyRelationships, error) {
	f := new(field.NoPartyRelationships)
	err := m.Body.Get(f)
	return f, err
}
func (m *PartyDetailsListRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *PartyDetailsListRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
