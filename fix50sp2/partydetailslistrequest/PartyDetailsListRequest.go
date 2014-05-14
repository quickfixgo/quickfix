//Package partydetailslistrequest msg type = CF.
package partydetailslistrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a PartyDetailsListRequest wrapper for the generic Message type
type Message struct {
	message.Message
}

//PartyDetailsListRequestID is a required field for PartyDetailsListRequest.
func (m Message) PartyDetailsListRequestID() (*field.PartyDetailsListRequestIDField, errors.MessageRejectError) {
	f := &field.PartyDetailsListRequestIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPartyDetailsListRequestID reads a PartyDetailsListRequestID from PartyDetailsListRequest.
func (m Message) GetPartyDetailsListRequestID(f *field.PartyDetailsListRequestIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyListResponseTypes is a required field for PartyDetailsListRequest.
func (m Message) NoPartyListResponseTypes() (*field.NoPartyListResponseTypesField, errors.MessageRejectError) {
	f := &field.NoPartyListResponseTypesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyListResponseTypes reads a NoPartyListResponseTypes from PartyDetailsListRequest.
func (m Message) GetNoPartyListResponseTypes(f *field.NoPartyListResponseTypesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for PartyDetailsListRequest.
func (m Message) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from PartyDetailsListRequest.
func (m Message) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRequestedPartyRoles is a non-required field for PartyDetailsListRequest.
func (m Message) NoRequestedPartyRoles() (*field.NoRequestedPartyRolesField, errors.MessageRejectError) {
	f := &field.NoRequestedPartyRolesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRequestedPartyRoles reads a NoRequestedPartyRoles from PartyDetailsListRequest.
func (m Message) GetNoRequestedPartyRoles(f *field.NoRequestedPartyRolesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyRelationships is a non-required field for PartyDetailsListRequest.
func (m Message) NoPartyRelationships() (*field.NoPartyRelationshipsField, errors.MessageRejectError) {
	f := &field.NoPartyRelationshipsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyRelationships reads a NoPartyRelationships from PartyDetailsListRequest.
func (m Message) GetNoPartyRelationships(f *field.NoPartyRelationshipsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a non-required field for PartyDetailsListRequest.
func (m Message) SubscriptionRequestType() (*field.SubscriptionRequestTypeField, errors.MessageRejectError) {
	f := &field.SubscriptionRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from PartyDetailsListRequest.
func (m Message) GetSubscriptionRequestType(f *field.SubscriptionRequestTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for PartyDetailsListRequest.
func (m Message) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from PartyDetailsListRequest.
func (m Message) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for PartyDetailsListRequest.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from PartyDetailsListRequest.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for PartyDetailsListRequest.
func (m Message) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from PartyDetailsListRequest.
func (m Message) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds PartyDetailsListRequest messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for PartyDetailsListRequest.
func Builder(
	partydetailslistrequestid *field.PartyDetailsListRequestIDField,
	nopartylistresponsetypes *field.NoPartyListResponseTypesField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header().Set(field.NewMsgType("CF"))
	builder.Body().Set(partydetailslistrequestid)
	builder.Body().Set(nopartylistresponsetypes)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "CF", r
}
