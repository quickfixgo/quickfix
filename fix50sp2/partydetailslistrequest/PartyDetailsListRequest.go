//Package partydetailslistrequest msg type = CF.
package partydetailslistrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
)

//Message is a PartyDetailsListRequest wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//PartyDetailsListRequestID is a required field for PartyDetailsListRequest.
func (m Message) PartyDetailsListRequestID() (*field.PartyDetailsListRequestIDField, quickfix.MessageRejectError) {
	f := &field.PartyDetailsListRequestIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPartyDetailsListRequestID reads a PartyDetailsListRequestID from PartyDetailsListRequest.
func (m Message) GetPartyDetailsListRequestID(f *field.PartyDetailsListRequestIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyListResponseTypes is a required field for PartyDetailsListRequest.
func (m Message) NoPartyListResponseTypes() (*field.NoPartyListResponseTypesField, quickfix.MessageRejectError) {
	f := &field.NoPartyListResponseTypesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyListResponseTypes reads a NoPartyListResponseTypes from PartyDetailsListRequest.
func (m Message) GetNoPartyListResponseTypes(f *field.NoPartyListResponseTypesField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for PartyDetailsListRequest.
func (m Message) NoPartyIDs() (*field.NoPartyIDsField, quickfix.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from PartyDetailsListRequest.
func (m Message) GetNoPartyIDs(f *field.NoPartyIDsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoRequestedPartyRoles is a non-required field for PartyDetailsListRequest.
func (m Message) NoRequestedPartyRoles() (*field.NoRequestedPartyRolesField, quickfix.MessageRejectError) {
	f := &field.NoRequestedPartyRolesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRequestedPartyRoles reads a NoRequestedPartyRoles from PartyDetailsListRequest.
func (m Message) GetNoRequestedPartyRoles(f *field.NoRequestedPartyRolesField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyRelationships is a non-required field for PartyDetailsListRequest.
func (m Message) NoPartyRelationships() (*field.NoPartyRelationshipsField, quickfix.MessageRejectError) {
	f := &field.NoPartyRelationshipsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyRelationships reads a NoPartyRelationships from PartyDetailsListRequest.
func (m Message) GetNoPartyRelationships(f *field.NoPartyRelationshipsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a non-required field for PartyDetailsListRequest.
func (m Message) SubscriptionRequestType() (*field.SubscriptionRequestTypeField, quickfix.MessageRejectError) {
	f := &field.SubscriptionRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from PartyDetailsListRequest.
func (m Message) GetSubscriptionRequestType(f *field.SubscriptionRequestTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for PartyDetailsListRequest.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from PartyDetailsListRequest.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for PartyDetailsListRequest.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, quickfix.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from PartyDetailsListRequest.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for PartyDetailsListRequest.
func (m Message) EncodedText() (*field.EncodedTextField, quickfix.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from PartyDetailsListRequest.
func (m Message) GetEncodedText(f *field.EncodedTextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for PartyDetailsListRequest.
func New(
	partydetailslistrequestid *field.PartyDetailsListRequestIDField,
	nopartylistresponsetypes *field.NoPartyListResponseTypesField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("CF"))
	builder.Body.Set(partydetailslistrequestid)
	builder.Body.Set(nopartylistresponsetypes)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "CF", r
}
