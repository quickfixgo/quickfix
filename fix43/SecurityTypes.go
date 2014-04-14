package fix43

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//SecurityTypes msg type = w.
type SecurityTypes struct {
	message.Message
}

//SecurityTypesBuilder builds SecurityTypes messages.
type SecurityTypesBuilder struct {
	message.MessageBuilder
}

//NewSecurityTypesBuilder returns an initialized SecurityTypesBuilder with specified required fields.
func NewSecurityTypesBuilder(
	securityreqid field.SecurityReqID,
	securityresponseid field.SecurityResponseID,
	securityresponsetype field.SecurityResponseType) *SecurityTypesBuilder {
	builder := new(SecurityTypesBuilder)
	builder.Body.Set(securityreqid)
	builder.Body.Set(securityresponseid)
	builder.Body.Set(securityresponsetype)
	return builder
}

//SecurityReqID is a required field for SecurityTypes.
func (m *SecurityTypes) SecurityReqID() (*field.SecurityReqID, error) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityResponseID is a required field for SecurityTypes.
func (m *SecurityTypes) SecurityResponseID() (*field.SecurityResponseID, error) {
	f := new(field.SecurityResponseID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityResponseType is a required field for SecurityTypes.
func (m *SecurityTypes) SecurityResponseType() (*field.SecurityResponseType, error) {
	f := new(field.SecurityResponseType)
	err := m.Body.Get(f)
	return f, err
}

//TotalNumSecurityTypes is a non-required field for SecurityTypes.
func (m *SecurityTypes) TotalNumSecurityTypes() (*field.TotalNumSecurityTypes, error) {
	f := new(field.TotalNumSecurityTypes)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityTypes is a non-required field for SecurityTypes.
func (m *SecurityTypes) NoSecurityTypes() (*field.NoSecurityTypes, error) {
	f := new(field.NoSecurityTypes)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for SecurityTypes.
func (m *SecurityTypes) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for SecurityTypes.
func (m *SecurityTypes) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for SecurityTypes.
func (m *SecurityTypes) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionID is a non-required field for SecurityTypes.
func (m *SecurityTypes) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionSubID is a non-required field for SecurityTypes.
func (m *SecurityTypes) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//SubscriptionRequestType is a non-required field for SecurityTypes.
func (m *SecurityTypes) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}
