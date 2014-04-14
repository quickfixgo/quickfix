package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//PartyDetailsListRequest msg type = CF.
type PartyDetailsListRequest struct {
	message.Message
}

//PartyDetailsListRequestBuilder builds PartyDetailsListRequest messages.
type PartyDetailsListRequestBuilder struct {
	message.MessageBuilder
}

//NewPartyDetailsListRequestBuilder returns an initialized PartyDetailsListRequestBuilder with specified required fields.
func NewPartyDetailsListRequestBuilder(
	partydetailslistrequestid field.PartyDetailsListRequestID,
	nopartylistresponsetypes field.NoPartyListResponseTypes) *PartyDetailsListRequestBuilder {
	builder := new(PartyDetailsListRequestBuilder)
	builder.Body.Set(partydetailslistrequestid)
	builder.Body.Set(nopartylistresponsetypes)
	return builder
}

//PartyDetailsListRequestID is a required field for PartyDetailsListRequest.
func (m *PartyDetailsListRequest) PartyDetailsListRequestID() (*field.PartyDetailsListRequestID, error) {
	f := new(field.PartyDetailsListRequestID)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyListResponseTypes is a required field for PartyDetailsListRequest.
func (m *PartyDetailsListRequest) NoPartyListResponseTypes() (*field.NoPartyListResponseTypes, error) {
	f := new(field.NoPartyListResponseTypes)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyIDs is a non-required field for PartyDetailsListRequest.
func (m *PartyDetailsListRequest) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//NoRequestedPartyRoles is a non-required field for PartyDetailsListRequest.
func (m *PartyDetailsListRequest) NoRequestedPartyRoles() (*field.NoRequestedPartyRoles, error) {
	f := new(field.NoRequestedPartyRoles)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyRelationships is a non-required field for PartyDetailsListRequest.
func (m *PartyDetailsListRequest) NoPartyRelationships() (*field.NoPartyRelationships, error) {
	f := new(field.NoPartyRelationships)
	err := m.Body.Get(f)
	return f, err
}

//SubscriptionRequestType is a non-required field for PartyDetailsListRequest.
func (m *PartyDetailsListRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for PartyDetailsListRequest.
func (m *PartyDetailsListRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for PartyDetailsListRequest.
func (m *PartyDetailsListRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for PartyDetailsListRequest.
func (m *PartyDetailsListRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
