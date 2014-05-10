package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//SecurityTypes msg type = w.
type SecurityTypes struct {
	message.Message
}

//SecurityTypesBuilder builds SecurityTypes messages.
type SecurityTypesBuilder struct {
	message.MessageBuilder
}

//CreateSecurityTypesBuilder returns an initialized SecurityTypesBuilder with specified required fields.
func CreateSecurityTypesBuilder(
	securityreqid *field.SecurityReqIDField,
	securityresponseid *field.SecurityResponseIDField,
	securityresponsetype *field.SecurityResponseTypeField) SecurityTypesBuilder {
	var builder SecurityTypesBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50))
	builder.Header.Set(field.NewMsgType("w"))
	builder.Body.Set(securityreqid)
	builder.Body.Set(securityresponseid)
	builder.Body.Set(securityresponsetype)
	return builder
}

//SecurityReqID is a required field for SecurityTypes.
func (m SecurityTypes) SecurityReqID() (*field.SecurityReqIDField, errors.MessageRejectError) {
	f := &field.SecurityReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReqID reads a SecurityReqID from SecurityTypes.
func (m SecurityTypes) GetSecurityReqID(f *field.SecurityReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityResponseID is a required field for SecurityTypes.
func (m SecurityTypes) SecurityResponseID() (*field.SecurityResponseIDField, errors.MessageRejectError) {
	f := &field.SecurityResponseIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityResponseID reads a SecurityResponseID from SecurityTypes.
func (m SecurityTypes) GetSecurityResponseID(f *field.SecurityResponseIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityResponseType is a required field for SecurityTypes.
func (m SecurityTypes) SecurityResponseType() (*field.SecurityResponseTypeField, errors.MessageRejectError) {
	f := &field.SecurityResponseTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityResponseType reads a SecurityResponseType from SecurityTypes.
func (m SecurityTypes) GetSecurityResponseType(f *field.SecurityResponseTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoSecurityTypes is a non-required field for SecurityTypes.
func (m SecurityTypes) TotNoSecurityTypes() (*field.TotNoSecurityTypesField, errors.MessageRejectError) {
	f := &field.TotNoSecurityTypesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoSecurityTypes reads a TotNoSecurityTypes from SecurityTypes.
func (m SecurityTypes) GetTotNoSecurityTypes(f *field.TotNoSecurityTypesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for SecurityTypes.
func (m SecurityTypes) LastFragment() (*field.LastFragmentField, errors.MessageRejectError) {
	f := &field.LastFragmentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from SecurityTypes.
func (m SecurityTypes) GetLastFragment(f *field.LastFragmentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityTypes is a non-required field for SecurityTypes.
func (m SecurityTypes) NoSecurityTypes() (*field.NoSecurityTypesField, errors.MessageRejectError) {
	f := &field.NoSecurityTypesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityTypes reads a NoSecurityTypes from SecurityTypes.
func (m SecurityTypes) GetNoSecurityTypes(f *field.NoSecurityTypesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for SecurityTypes.
func (m SecurityTypes) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from SecurityTypes.
func (m SecurityTypes) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for SecurityTypes.
func (m SecurityTypes) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from SecurityTypes.
func (m SecurityTypes) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for SecurityTypes.
func (m SecurityTypes) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from SecurityTypes.
func (m SecurityTypes) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for SecurityTypes.
func (m SecurityTypes) TradingSessionID() (*field.TradingSessionIDField, errors.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from SecurityTypes.
func (m SecurityTypes) GetTradingSessionID(f *field.TradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for SecurityTypes.
func (m SecurityTypes) TradingSessionSubID() (*field.TradingSessionSubIDField, errors.MessageRejectError) {
	f := &field.TradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from SecurityTypes.
func (m SecurityTypes) GetTradingSessionSubID(f *field.TradingSessionSubIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a non-required field for SecurityTypes.
func (m SecurityTypes) SubscriptionRequestType() (*field.SubscriptionRequestTypeField, errors.MessageRejectError) {
	f := &field.SubscriptionRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from SecurityTypes.
func (m SecurityTypes) GetSubscriptionRequestType(f *field.SubscriptionRequestTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}
