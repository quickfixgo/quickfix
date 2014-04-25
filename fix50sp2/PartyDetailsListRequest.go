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
	builder.Header.Set(field.BuildMsgType("CF"))
	builder.Body.Set(partydetailslistrequestid)
	builder.Body.Set(nopartylistresponsetypes)
	return builder
}

//PartyDetailsListRequestID is a required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) PartyDetailsListRequestID() (*field.PartyDetailsListRequestID, errors.MessageRejectError) {
	f := new(field.PartyDetailsListRequestID)
	err := m.Body.Get(f)
	return f, err
}

//GetPartyDetailsListRequestID reads a PartyDetailsListRequestID from PartyDetailsListRequest.
func (m PartyDetailsListRequest) GetPartyDetailsListRequestID(f *field.PartyDetailsListRequestID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyListResponseTypes is a required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) NoPartyListResponseTypes() (*field.NoPartyListResponseTypes, errors.MessageRejectError) {
	f := new(field.NoPartyListResponseTypes)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyListResponseTypes reads a NoPartyListResponseTypes from PartyDetailsListRequest.
func (m PartyDetailsListRequest) GetNoPartyListResponseTypes(f *field.NoPartyListResponseTypes) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from PartyDetailsListRequest.
func (m PartyDetailsListRequest) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRequestedPartyRoles is a non-required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) NoRequestedPartyRoles() (*field.NoRequestedPartyRoles, errors.MessageRejectError) {
	f := new(field.NoRequestedPartyRoles)
	err := m.Body.Get(f)
	return f, err
}

//GetNoRequestedPartyRoles reads a NoRequestedPartyRoles from PartyDetailsListRequest.
func (m PartyDetailsListRequest) GetNoRequestedPartyRoles(f *field.NoRequestedPartyRoles) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyRelationships is a non-required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) NoPartyRelationships() (*field.NoPartyRelationships, errors.MessageRejectError) {
	f := new(field.NoPartyRelationships)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyRelationships reads a NoPartyRelationships from PartyDetailsListRequest.
func (m PartyDetailsListRequest) GetNoPartyRelationships(f *field.NoPartyRelationships) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a non-required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, errors.MessageRejectError) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from PartyDetailsListRequest.
func (m PartyDetailsListRequest) GetSubscriptionRequestType(f *field.SubscriptionRequestType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from PartyDetailsListRequest.
func (m PartyDetailsListRequest) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from PartyDetailsListRequest.
func (m PartyDetailsListRequest) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from PartyDetailsListRequest.
func (m PartyDetailsListRequest) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}
