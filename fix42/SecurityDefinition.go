package fix42

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
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
	securityreqid *field.SecurityReqIDField,
	securityresponseid *field.SecurityResponseIDField,
	totalnumsecurities *field.TotalNumSecuritiesField) SecurityDefinitionBuilder {
	var builder SecurityDefinitionBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX42))
	builder.Header().Set(field.NewMsgType("d"))
	builder.Body().Set(securityreqid)
	builder.Body().Set(securityresponseid)
	builder.Body().Set(totalnumsecurities)
	return builder
}

//SecurityReqID is a required field for SecurityDefinition.
func (m SecurityDefinition) SecurityReqID() (*field.SecurityReqIDField, errors.MessageRejectError) {
	f := &field.SecurityReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReqID reads a SecurityReqID from SecurityDefinition.
func (m SecurityDefinition) GetSecurityReqID(f *field.SecurityReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityResponseID is a required field for SecurityDefinition.
func (m SecurityDefinition) SecurityResponseID() (*field.SecurityResponseIDField, errors.MessageRejectError) {
	f := &field.SecurityResponseIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityResponseID reads a SecurityResponseID from SecurityDefinition.
func (m SecurityDefinition) GetSecurityResponseID(f *field.SecurityResponseIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityResponseType is a non-required field for SecurityDefinition.
func (m SecurityDefinition) SecurityResponseType() (*field.SecurityResponseTypeField, errors.MessageRejectError) {
	f := &field.SecurityResponseTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityResponseType reads a SecurityResponseType from SecurityDefinition.
func (m SecurityDefinition) GetSecurityResponseType(f *field.SecurityResponseTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalNumSecurities is a required field for SecurityDefinition.
func (m SecurityDefinition) TotalNumSecurities() (*field.TotalNumSecuritiesField, errors.MessageRejectError) {
	f := &field.TotalNumSecuritiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotalNumSecurities reads a TotalNumSecurities from SecurityDefinition.
func (m SecurityDefinition) GetTotalNumSecurities(f *field.TotalNumSecuritiesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for SecurityDefinition.
func (m SecurityDefinition) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from SecurityDefinition.
func (m SecurityDefinition) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for SecurityDefinition.
func (m SecurityDefinition) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from SecurityDefinition.
func (m SecurityDefinition) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for SecurityDefinition.
func (m SecurityDefinition) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from SecurityDefinition.
func (m SecurityDefinition) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IDSource is a non-required field for SecurityDefinition.
func (m SecurityDefinition) IDSource() (*field.IDSourceField, errors.MessageRejectError) {
	f := &field.IDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIDSource reads a IDSource from SecurityDefinition.
func (m SecurityDefinition) GetIDSource(f *field.IDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for SecurityDefinition.
func (m SecurityDefinition) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from SecurityDefinition.
func (m SecurityDefinition) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for SecurityDefinition.
func (m SecurityDefinition) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from SecurityDefinition.
func (m SecurityDefinition) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDay is a non-required field for SecurityDefinition.
func (m SecurityDefinition) MaturityDay() (*field.MaturityDayField, errors.MessageRejectError) {
	f := &field.MaturityDayField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDay reads a MaturityDay from SecurityDefinition.
func (m SecurityDefinition) GetMaturityDay(f *field.MaturityDayField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for SecurityDefinition.
func (m SecurityDefinition) PutOrCall() (*field.PutOrCallField, errors.MessageRejectError) {
	f := &field.PutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from SecurityDefinition.
func (m SecurityDefinition) GetPutOrCall(f *field.PutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for SecurityDefinition.
func (m SecurityDefinition) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from SecurityDefinition.
func (m SecurityDefinition) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for SecurityDefinition.
func (m SecurityDefinition) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from SecurityDefinition.
func (m SecurityDefinition) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for SecurityDefinition.
func (m SecurityDefinition) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from SecurityDefinition.
func (m SecurityDefinition) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for SecurityDefinition.
func (m SecurityDefinition) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from SecurityDefinition.
func (m SecurityDefinition) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for SecurityDefinition.
func (m SecurityDefinition) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from SecurityDefinition.
func (m SecurityDefinition) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for SecurityDefinition.
func (m SecurityDefinition) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from SecurityDefinition.
func (m SecurityDefinition) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for SecurityDefinition.
func (m SecurityDefinition) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from SecurityDefinition.
func (m SecurityDefinition) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for SecurityDefinition.
func (m SecurityDefinition) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from SecurityDefinition.
func (m SecurityDefinition) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for SecurityDefinition.
func (m SecurityDefinition) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from SecurityDefinition.
func (m SecurityDefinition) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for SecurityDefinition.
func (m SecurityDefinition) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from SecurityDefinition.
func (m SecurityDefinition) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for SecurityDefinition.
func (m SecurityDefinition) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from SecurityDefinition.
func (m SecurityDefinition) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for SecurityDefinition.
func (m SecurityDefinition) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from SecurityDefinition.
func (m SecurityDefinition) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for SecurityDefinition.
func (m SecurityDefinition) TradingSessionID() (*field.TradingSessionIDField, errors.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from SecurityDefinition.
func (m SecurityDefinition) GetTradingSessionID(f *field.TradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for SecurityDefinition.
func (m SecurityDefinition) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from SecurityDefinition.
func (m SecurityDefinition) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for SecurityDefinition.
func (m SecurityDefinition) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from SecurityDefinition.
func (m SecurityDefinition) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for SecurityDefinition.
func (m SecurityDefinition) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from SecurityDefinition.
func (m SecurityDefinition) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a non-required field for SecurityDefinition.
func (m SecurityDefinition) NoRelatedSym() (*field.NoRelatedSymField, errors.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from SecurityDefinition.
func (m SecurityDefinition) GetNoRelatedSym(f *field.NoRelatedSymField) errors.MessageRejectError {
	return m.Body.Get(f)
}
