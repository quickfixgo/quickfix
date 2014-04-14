package fix50sp1

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

//TotNoSecurityTypes is a non-required field for SecurityTypes.
func (m *SecurityTypes) TotNoSecurityTypes() (*field.TotNoSecurityTypes, error) {
	f := new(field.TotNoSecurityTypes)
	err := m.Body.Get(f)
	return f, err
}

//LastFragment is a non-required field for SecurityTypes.
func (m *SecurityTypes) LastFragment() (*field.LastFragment, error) {
	f := new(field.LastFragment)
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

//MarketID is a non-required field for SecurityTypes.
func (m *SecurityTypes) MarketID() (*field.MarketID, error) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}

//MarketSegmentID is a non-required field for SecurityTypes.
func (m *SecurityTypes) MarketSegmentID() (*field.MarketSegmentID, error) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}

//ApplID is a non-required field for SecurityTypes.
func (m *SecurityTypes) ApplID() (*field.ApplID, error) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}

//ApplSeqNum is a non-required field for SecurityTypes.
func (m *SecurityTypes) ApplSeqNum() (*field.ApplSeqNum, error) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//ApplLastSeqNum is a non-required field for SecurityTypes.
func (m *SecurityTypes) ApplLastSeqNum() (*field.ApplLastSeqNum, error) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//ApplResendFlag is a non-required field for SecurityTypes.
func (m *SecurityTypes) ApplResendFlag() (*field.ApplResendFlag, error) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}
