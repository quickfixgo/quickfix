package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
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

//CreatePartyDetailsListRequestBuilder returns an initialized PartyDetailsListRequestBuilder with specified required fields.
func CreatePartyDetailsListRequestBuilder(
	partydetailslistrequestid field.PartyDetailsListRequestID,
	nopartylistresponsetypes field.NoPartyListResponseTypes) PartyDetailsListRequestBuilder {
	var builder PartyDetailsListRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(partydetailslistrequestid)
	builder.Body.Set(nopartylistresponsetypes)
	return builder
}

//PartyDetailsListRequestID is a required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) PartyDetailsListRequestID() (field.PartyDetailsListRequestID, errors.MessageRejectError) {
	var f field.PartyDetailsListRequestID
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyListResponseTypes is a required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) NoPartyListResponseTypes() (field.NoPartyListResponseTypes, errors.MessageRejectError) {
	var f field.NoPartyListResponseTypes
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) NoPartyIDs() (field.NoPartyIDs, errors.MessageRejectError) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//NoRequestedPartyRoles is a non-required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) NoRequestedPartyRoles() (field.NoRequestedPartyRoles, errors.MessageRejectError) {
	var f field.NoRequestedPartyRoles
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyRelationships is a non-required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) NoPartyRelationships() (field.NoPartyRelationships, errors.MessageRejectError) {
	var f field.NoPartyRelationships
	err := m.Body.Get(&f)
	return f, err
}

//SubscriptionRequestType is a non-required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) SubscriptionRequestType() (field.SubscriptionRequestType, errors.MessageRejectError) {
	var f field.SubscriptionRequestType
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
