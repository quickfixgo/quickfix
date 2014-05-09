package fix42

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Advertisement msg type = 7.
type Advertisement struct {
	message.Message
}

//AdvertisementBuilder builds Advertisement messages.
type AdvertisementBuilder struct {
	message.MessageBuilder
}

//CreateAdvertisementBuilder returns an initialized AdvertisementBuilder with specified required fields.
func CreateAdvertisementBuilder(
	advid *field.AdvIdField,
	advtranstype *field.AdvTransTypeField,
	symbol *field.SymbolField,
	advside *field.AdvSideField,
	shares *field.SharesField) AdvertisementBuilder {
	var builder AdvertisementBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX42))
	builder.Header.Set(field.NewMsgType("7"))
	builder.Body.Set(advid)
	builder.Body.Set(advtranstype)
	builder.Body.Set(symbol)
	builder.Body.Set(advside)
	builder.Body.Set(shares)
	return builder
}

//AdvId is a required field for Advertisement.
func (m Advertisement) AdvId() (*field.AdvIdField, errors.MessageRejectError) {
	f := &field.AdvIdField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAdvId reads a AdvId from Advertisement.
func (m Advertisement) GetAdvId(f *field.AdvIdField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AdvTransType is a required field for Advertisement.
func (m Advertisement) AdvTransType() (*field.AdvTransTypeField, errors.MessageRejectError) {
	f := &field.AdvTransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAdvTransType reads a AdvTransType from Advertisement.
func (m Advertisement) GetAdvTransType(f *field.AdvTransTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AdvRefID is a non-required field for Advertisement.
func (m Advertisement) AdvRefID() (*field.AdvRefIDField, errors.MessageRejectError) {
	f := &field.AdvRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAdvRefID reads a AdvRefID from Advertisement.
func (m Advertisement) GetAdvRefID(f *field.AdvRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a required field for Advertisement.
func (m Advertisement) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from Advertisement.
func (m Advertisement) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for Advertisement.
func (m Advertisement) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from Advertisement.
func (m Advertisement) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for Advertisement.
func (m Advertisement) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from Advertisement.
func (m Advertisement) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IDSource is a non-required field for Advertisement.
func (m Advertisement) IDSource() (*field.IDSourceField, errors.MessageRejectError) {
	f := &field.IDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIDSource reads a IDSource from Advertisement.
func (m Advertisement) GetIDSource(f *field.IDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for Advertisement.
func (m Advertisement) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from Advertisement.
func (m Advertisement) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for Advertisement.
func (m Advertisement) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from Advertisement.
func (m Advertisement) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDay is a non-required field for Advertisement.
func (m Advertisement) MaturityDay() (*field.MaturityDayField, errors.MessageRejectError) {
	f := &field.MaturityDayField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDay reads a MaturityDay from Advertisement.
func (m Advertisement) GetMaturityDay(f *field.MaturityDayField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for Advertisement.
func (m Advertisement) PutOrCall() (*field.PutOrCallField, errors.MessageRejectError) {
	f := &field.PutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from Advertisement.
func (m Advertisement) GetPutOrCall(f *field.PutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for Advertisement.
func (m Advertisement) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from Advertisement.
func (m Advertisement) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for Advertisement.
func (m Advertisement) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from Advertisement.
func (m Advertisement) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for Advertisement.
func (m Advertisement) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from Advertisement.
func (m Advertisement) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for Advertisement.
func (m Advertisement) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from Advertisement.
func (m Advertisement) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for Advertisement.
func (m Advertisement) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from Advertisement.
func (m Advertisement) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for Advertisement.
func (m Advertisement) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from Advertisement.
func (m Advertisement) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for Advertisement.
func (m Advertisement) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from Advertisement.
func (m Advertisement) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for Advertisement.
func (m Advertisement) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from Advertisement.
func (m Advertisement) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for Advertisement.
func (m Advertisement) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from Advertisement.
func (m Advertisement) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for Advertisement.
func (m Advertisement) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from Advertisement.
func (m Advertisement) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for Advertisement.
func (m Advertisement) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from Advertisement.
func (m Advertisement) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AdvSide is a required field for Advertisement.
func (m Advertisement) AdvSide() (*field.AdvSideField, errors.MessageRejectError) {
	f := &field.AdvSideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAdvSide reads a AdvSide from Advertisement.
func (m Advertisement) GetAdvSide(f *field.AdvSideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Shares is a required field for Advertisement.
func (m Advertisement) Shares() (*field.SharesField, errors.MessageRejectError) {
	f := &field.SharesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetShares reads a Shares from Advertisement.
func (m Advertisement) GetShares(f *field.SharesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for Advertisement.
func (m Advertisement) Price() (*field.PriceField, errors.MessageRejectError) {
	f := &field.PriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from Advertisement.
func (m Advertisement) GetPrice(f *field.PriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for Advertisement.
func (m Advertisement) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from Advertisement.
func (m Advertisement) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for Advertisement.
func (m Advertisement) TradeDate() (*field.TradeDateField, errors.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from Advertisement.
func (m Advertisement) GetTradeDate(f *field.TradeDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for Advertisement.
func (m Advertisement) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from Advertisement.
func (m Advertisement) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for Advertisement.
func (m Advertisement) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from Advertisement.
func (m Advertisement) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for Advertisement.
func (m Advertisement) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from Advertisement.
func (m Advertisement) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for Advertisement.
func (m Advertisement) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from Advertisement.
func (m Advertisement) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//URLLink is a non-required field for Advertisement.
func (m Advertisement) URLLink() (*field.URLLinkField, errors.MessageRejectError) {
	f := &field.URLLinkField{}
	err := m.Body.Get(f)
	return f, err
}

//GetURLLink reads a URLLink from Advertisement.
func (m Advertisement) GetURLLink(f *field.URLLinkField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastMkt is a non-required field for Advertisement.
func (m Advertisement) LastMkt() (*field.LastMktField, errors.MessageRejectError) {
	f := &field.LastMktField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastMkt reads a LastMkt from Advertisement.
func (m Advertisement) GetLastMkt(f *field.LastMktField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for Advertisement.
func (m Advertisement) TradingSessionID() (*field.TradingSessionIDField, errors.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from Advertisement.
func (m Advertisement) GetTradingSessionID(f *field.TradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}
