package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
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
	partydetailslistrequestid *field.PartyDetailsListRequestIDField,
	nopartylistresponsetypes *field.NoPartyListResponseTypesField) PartyDetailsListRequestBuilder {
	var builder PartyDetailsListRequestBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("CF"))
	builder.Body.Set(partydetailslistrequestid)
	builder.Body.Set(nopartylistresponsetypes)
	return builder
}

//PartyDetailsListRequestID is a required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) PartyDetailsListRequestID() (*field.PartyDetailsListRequestIDField, errors.MessageRejectError) {
	f := &field.PartyDetailsListRequestIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPartyDetailsListRequestID reads a PartyDetailsListRequestID from PartyDetailsListRequest.
func (m PartyDetailsListRequest) GetPartyDetailsListRequestID(f *field.PartyDetailsListRequestIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyListResponseTypes is a required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) NoPartyListResponseTypes() (*field.NoPartyListResponseTypesField, errors.MessageRejectError) {
	f := &field.NoPartyListResponseTypesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyListResponseTypes reads a NoPartyListResponseTypes from PartyDetailsListRequest.
func (m PartyDetailsListRequest) GetNoPartyListResponseTypes(f *field.NoPartyListResponseTypesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from PartyDetailsListRequest.
func (m PartyDetailsListRequest) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRequestedPartyRoles is a non-required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) NoRequestedPartyRoles() (*field.NoRequestedPartyRolesField, errors.MessageRejectError) {
	f := &field.NoRequestedPartyRolesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRequestedPartyRoles reads a NoRequestedPartyRoles from PartyDetailsListRequest.
func (m PartyDetailsListRequest) GetNoRequestedPartyRoles(f *field.NoRequestedPartyRolesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyRelationships is a non-required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) NoPartyRelationships() (*field.NoPartyRelationshipsField, errors.MessageRejectError) {
	f := &field.NoPartyRelationshipsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyRelationships reads a NoPartyRelationships from PartyDetailsListRequest.
func (m PartyDetailsListRequest) GetNoPartyRelationships(f *field.NoPartyRelationshipsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a non-required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) SubscriptionRequestType() (*field.SubscriptionRequestTypeField, errors.MessageRejectError) {
	f := &field.SubscriptionRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from PartyDetailsListRequest.
func (m PartyDetailsListRequest) GetSubscriptionRequestType(f *field.SubscriptionRequestTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from PartyDetailsListRequest.
func (m PartyDetailsListRequest) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from PartyDetailsListRequest.
func (m PartyDetailsListRequest) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for PartyDetailsListRequest.
func (m PartyDetailsListRequest) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from PartyDetailsListRequest.
func (m PartyDetailsListRequest) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
