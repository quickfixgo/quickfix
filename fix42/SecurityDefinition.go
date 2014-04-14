package fix42

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//SecurityDefinition msg type = d.
type SecurityDefinition struct {
	message.Message
}

//SecurityDefinitionBuilder builds SecurityDefinition messages.
type SecurityDefinitionBuilder struct {
	message.MessageBuilder
}

//NewSecurityDefinitionBuilder returns an initialized SecurityDefinitionBuilder with specified required fields.
func NewSecurityDefinitionBuilder(
	securityreqid field.SecurityReqID,
	securityresponseid field.SecurityResponseID,
	totalnumsecurities field.TotalNumSecurities) *SecurityDefinitionBuilder {
	builder := new(SecurityDefinitionBuilder)
	builder.Body.Set(securityreqid)
	builder.Body.Set(securityresponseid)
	builder.Body.Set(totalnumsecurities)
	return builder
}

//SecurityReqID is a required field for SecurityDefinition.
func (m *SecurityDefinition) SecurityReqID() (*field.SecurityReqID, error) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityResponseID is a required field for SecurityDefinition.
func (m *SecurityDefinition) SecurityResponseID() (*field.SecurityResponseID, error) {
	f := new(field.SecurityResponseID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityResponseType is a non-required field for SecurityDefinition.
func (m *SecurityDefinition) SecurityResponseType() (*field.SecurityResponseType, error) {
	f := new(field.SecurityResponseType)
	err := m.Body.Get(f)
	return f, err
}

//TotalNumSecurities is a required field for SecurityDefinition.
func (m *SecurityDefinition) TotalNumSecurities() (*field.TotalNumSecurities, error) {
	f := new(field.TotalNumSecurities)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for SecurityDefinition.
func (m *SecurityDefinition) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for SecurityDefinition.
func (m *SecurityDefinition) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for SecurityDefinition.
func (m *SecurityDefinition) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//IDSource is a non-required field for SecurityDefinition.
func (m *SecurityDefinition) IDSource() (*field.IDSource, error) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for SecurityDefinition.
func (m *SecurityDefinition) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for SecurityDefinition.
func (m *SecurityDefinition) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDay is a non-required field for SecurityDefinition.
func (m *SecurityDefinition) MaturityDay() (*field.MaturityDay, error) {
	f := new(field.MaturityDay)
	err := m.Body.Get(f)
	return f, err
}

//PutOrCall is a non-required field for SecurityDefinition.
func (m *SecurityDefinition) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for SecurityDefinition.
func (m *SecurityDefinition) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for SecurityDefinition.
func (m *SecurityDefinition) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for SecurityDefinition.
func (m *SecurityDefinition) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for SecurityDefinition.
func (m *SecurityDefinition) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for SecurityDefinition.
func (m *SecurityDefinition) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for SecurityDefinition.
func (m *SecurityDefinition) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for SecurityDefinition.
func (m *SecurityDefinition) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for SecurityDefinition.
func (m *SecurityDefinition) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for SecurityDefinition.
func (m *SecurityDefinition) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for SecurityDefinition.
func (m *SecurityDefinition) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for SecurityDefinition.
func (m *SecurityDefinition) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for SecurityDefinition.
func (m *SecurityDefinition) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionID is a non-required field for SecurityDefinition.
func (m *SecurityDefinition) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for SecurityDefinition.
func (m *SecurityDefinition) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for SecurityDefinition.
func (m *SecurityDefinition) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for SecurityDefinition.
func (m *SecurityDefinition) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//NoRelatedSym is a non-required field for SecurityDefinition.
func (m *SecurityDefinition) NoRelatedSym() (*field.NoRelatedSym, error) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}
