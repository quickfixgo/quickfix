package fix42

import (
	"github.com/quickfixgo/quickfix/errors"
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
	builder.Header.Set(field.BuildMsgType("d"))
	builder.Body.Set(securityreqid)
	builder.Body.Set(securityresponseid)
	builder.Body.Set(totalnumsecurities)
	return builder
}

//SecurityReqID is a required field for SecurityDefinition.
func (m SecurityDefinition) SecurityReqID() (*field.SecurityReqID, errors.MessageRejectError) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReqID reads a SecurityReqID from SecurityDefinition.
func (m SecurityDefinition) GetSecurityReqID(f *field.SecurityReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityResponseID is a required field for SecurityDefinition.
func (m SecurityDefinition) SecurityResponseID() (*field.SecurityResponseID, errors.MessageRejectError) {
	f := new(field.SecurityResponseID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityResponseID reads a SecurityResponseID from SecurityDefinition.
func (m SecurityDefinition) GetSecurityResponseID(f *field.SecurityResponseID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityResponseType is a non-required field for SecurityDefinition.
func (m SecurityDefinition) SecurityResponseType() (*field.SecurityResponseType, errors.MessageRejectError) {
	f := new(field.SecurityResponseType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityResponseType reads a SecurityResponseType from SecurityDefinition.
func (m SecurityDefinition) GetSecurityResponseType(f *field.SecurityResponseType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalNumSecurities is a required field for SecurityDefinition.
func (m SecurityDefinition) TotalNumSecurities() (*field.TotalNumSecurities, errors.MessageRejectError) {
	f := new(field.TotalNumSecurities)
	err := m.Body.Get(f)
	return f, err
}

//GetTotalNumSecurities reads a TotalNumSecurities from SecurityDefinition.
func (m SecurityDefinition) GetTotalNumSecurities(f *field.TotalNumSecurities) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for SecurityDefinition.
func (m SecurityDefinition) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from SecurityDefinition.
func (m SecurityDefinition) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for SecurityDefinition.
func (m SecurityDefinition) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from SecurityDefinition.
func (m SecurityDefinition) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for SecurityDefinition.
func (m SecurityDefinition) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from SecurityDefinition.
func (m SecurityDefinition) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IDSource is a non-required field for SecurityDefinition.
func (m SecurityDefinition) IDSource() (*field.IDSource, errors.MessageRejectError) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetIDSource reads a IDSource from SecurityDefinition.
func (m SecurityDefinition) GetIDSource(f *field.IDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for SecurityDefinition.
func (m SecurityDefinition) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from SecurityDefinition.
func (m SecurityDefinition) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for SecurityDefinition.
func (m SecurityDefinition) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from SecurityDefinition.
func (m SecurityDefinition) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDay is a non-required field for SecurityDefinition.
func (m SecurityDefinition) MaturityDay() (*field.MaturityDay, errors.MessageRejectError) {
	f := new(field.MaturityDay)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDay reads a MaturityDay from SecurityDefinition.
func (m SecurityDefinition) GetMaturityDay(f *field.MaturityDay) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for SecurityDefinition.
func (m SecurityDefinition) PutOrCall() (*field.PutOrCall, errors.MessageRejectError) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from SecurityDefinition.
func (m SecurityDefinition) GetPutOrCall(f *field.PutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for SecurityDefinition.
func (m SecurityDefinition) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from SecurityDefinition.
func (m SecurityDefinition) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for SecurityDefinition.
func (m SecurityDefinition) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from SecurityDefinition.
func (m SecurityDefinition) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for SecurityDefinition.
func (m SecurityDefinition) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from SecurityDefinition.
func (m SecurityDefinition) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for SecurityDefinition.
func (m SecurityDefinition) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from SecurityDefinition.
func (m SecurityDefinition) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for SecurityDefinition.
func (m SecurityDefinition) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from SecurityDefinition.
func (m SecurityDefinition) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for SecurityDefinition.
func (m SecurityDefinition) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from SecurityDefinition.
func (m SecurityDefinition) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for SecurityDefinition.
func (m SecurityDefinition) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from SecurityDefinition.
func (m SecurityDefinition) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for SecurityDefinition.
func (m SecurityDefinition) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from SecurityDefinition.
func (m SecurityDefinition) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for SecurityDefinition.
func (m SecurityDefinition) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from SecurityDefinition.
func (m SecurityDefinition) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for SecurityDefinition.
func (m SecurityDefinition) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from SecurityDefinition.
func (m SecurityDefinition) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for SecurityDefinition.
func (m SecurityDefinition) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from SecurityDefinition.
func (m SecurityDefinition) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for SecurityDefinition.
func (m SecurityDefinition) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from SecurityDefinition.
func (m SecurityDefinition) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for SecurityDefinition.
func (m SecurityDefinition) TradingSessionID() (*field.TradingSessionID, errors.MessageRejectError) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from SecurityDefinition.
func (m SecurityDefinition) GetTradingSessionID(f *field.TradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for SecurityDefinition.
func (m SecurityDefinition) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from SecurityDefinition.
func (m SecurityDefinition) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for SecurityDefinition.
func (m SecurityDefinition) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from SecurityDefinition.
func (m SecurityDefinition) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for SecurityDefinition.
func (m SecurityDefinition) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from SecurityDefinition.
func (m SecurityDefinition) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a non-required field for SecurityDefinition.
func (m SecurityDefinition) NoRelatedSym() (*field.NoRelatedSym, errors.MessageRejectError) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from SecurityDefinition.
func (m SecurityDefinition) GetNoRelatedSym(f *field.NoRelatedSym) errors.MessageRejectError {
	return m.Body.Get(f)
}
