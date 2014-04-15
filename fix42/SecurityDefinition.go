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

//CreateSecurityDefinitionBuilder returns an initialized SecurityDefinitionBuilder with specified required fields.
func CreateSecurityDefinitionBuilder(
	securityreqid field.SecurityReqID,
	securityresponseid field.SecurityResponseID,
	totalnumsecurities field.TotalNumSecurities) SecurityDefinitionBuilder {
	var builder SecurityDefinitionBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(securityreqid)
	builder.Body.Set(securityresponseid)
	builder.Body.Set(totalnumsecurities)
	return builder
}

//SecurityReqID is a required field for SecurityDefinition.
func (m SecurityDefinition) SecurityReqID() (field.SecurityReqID, error) {
	var f field.SecurityReqID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityResponseID is a required field for SecurityDefinition.
func (m SecurityDefinition) SecurityResponseID() (field.SecurityResponseID, error) {
	var f field.SecurityResponseID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityResponseType is a non-required field for SecurityDefinition.
func (m SecurityDefinition) SecurityResponseType() (field.SecurityResponseType, error) {
	var f field.SecurityResponseType
	err := m.Body.Get(&f)
	return f, err
}

//TotalNumSecurities is a required field for SecurityDefinition.
func (m SecurityDefinition) TotalNumSecurities() (field.TotalNumSecurities, error) {
	var f field.TotalNumSecurities
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for SecurityDefinition.
func (m SecurityDefinition) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for SecurityDefinition.
func (m SecurityDefinition) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for SecurityDefinition.
func (m SecurityDefinition) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//IDSource is a non-required field for SecurityDefinition.
func (m SecurityDefinition) IDSource() (field.IDSource, error) {
	var f field.IDSource
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for SecurityDefinition.
func (m SecurityDefinition) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for SecurityDefinition.
func (m SecurityDefinition) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDay is a non-required field for SecurityDefinition.
func (m SecurityDefinition) MaturityDay() (field.MaturityDay, error) {
	var f field.MaturityDay
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for SecurityDefinition.
func (m SecurityDefinition) PutOrCall() (field.PutOrCall, error) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for SecurityDefinition.
func (m SecurityDefinition) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for SecurityDefinition.
func (m SecurityDefinition) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for SecurityDefinition.
func (m SecurityDefinition) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for SecurityDefinition.
func (m SecurityDefinition) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for SecurityDefinition.
func (m SecurityDefinition) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for SecurityDefinition.
func (m SecurityDefinition) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for SecurityDefinition.
func (m SecurityDefinition) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for SecurityDefinition.
func (m SecurityDefinition) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for SecurityDefinition.
func (m SecurityDefinition) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for SecurityDefinition.
func (m SecurityDefinition) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for SecurityDefinition.
func (m SecurityDefinition) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for SecurityDefinition.
func (m SecurityDefinition) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for SecurityDefinition.
func (m SecurityDefinition) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for SecurityDefinition.
func (m SecurityDefinition) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for SecurityDefinition.
func (m SecurityDefinition) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for SecurityDefinition.
func (m SecurityDefinition) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//NoRelatedSym is a non-required field for SecurityDefinition.
func (m SecurityDefinition) NoRelatedSym() (field.NoRelatedSym, error) {
	var f field.NoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}
