package fix41

import (
	"github.com/quickfixgo/quickfix/errors"
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
	advid field.AdvId,
	advtranstype field.AdvTransType,
	symbol field.Symbol,
	advside field.AdvSide,
	shares field.Shares) AdvertisementBuilder {
	var builder AdvertisementBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("7"))
	builder.Body.Set(advid)
	builder.Body.Set(advtranstype)
	builder.Body.Set(symbol)
	builder.Body.Set(advside)
	builder.Body.Set(shares)
	return builder
}

//AdvId is a required field for Advertisement.
func (m Advertisement) AdvId() (*field.AdvId, errors.MessageRejectError) {
	f := new(field.AdvId)
	err := m.Body.Get(f)
	return f, err
}

//GetAdvId reads a AdvId from Advertisement.
func (m Advertisement) GetAdvId(f *field.AdvId) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AdvTransType is a required field for Advertisement.
func (m Advertisement) AdvTransType() (*field.AdvTransType, errors.MessageRejectError) {
	f := new(field.AdvTransType)
	err := m.Body.Get(f)
	return f, err
}

//GetAdvTransType reads a AdvTransType from Advertisement.
func (m Advertisement) GetAdvTransType(f *field.AdvTransType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AdvRefID is a non-required field for Advertisement.
func (m Advertisement) AdvRefID() (*field.AdvRefID, errors.MessageRejectError) {
	f := new(field.AdvRefID)
	err := m.Body.Get(f)
	return f, err
}

//GetAdvRefID reads a AdvRefID from Advertisement.
func (m Advertisement) GetAdvRefID(f *field.AdvRefID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a required field for Advertisement.
func (m Advertisement) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from Advertisement.
func (m Advertisement) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for Advertisement.
func (m Advertisement) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from Advertisement.
func (m Advertisement) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for Advertisement.
func (m Advertisement) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from Advertisement.
func (m Advertisement) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IDSource is a non-required field for Advertisement.
func (m Advertisement) IDSource() (*field.IDSource, errors.MessageRejectError) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetIDSource reads a IDSource from Advertisement.
func (m Advertisement) GetIDSource(f *field.IDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for Advertisement.
func (m Advertisement) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from Advertisement.
func (m Advertisement) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for Advertisement.
func (m Advertisement) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from Advertisement.
func (m Advertisement) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDay is a non-required field for Advertisement.
func (m Advertisement) MaturityDay() (*field.MaturityDay, errors.MessageRejectError) {
	f := new(field.MaturityDay)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDay reads a MaturityDay from Advertisement.
func (m Advertisement) GetMaturityDay(f *field.MaturityDay) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for Advertisement.
func (m Advertisement) PutOrCall() (*field.PutOrCall, errors.MessageRejectError) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from Advertisement.
func (m Advertisement) GetPutOrCall(f *field.PutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for Advertisement.
func (m Advertisement) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from Advertisement.
func (m Advertisement) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for Advertisement.
func (m Advertisement) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from Advertisement.
func (m Advertisement) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for Advertisement.
func (m Advertisement) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from Advertisement.
func (m Advertisement) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for Advertisement.
func (m Advertisement) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from Advertisement.
func (m Advertisement) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for Advertisement.
func (m Advertisement) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from Advertisement.
func (m Advertisement) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AdvSide is a required field for Advertisement.
func (m Advertisement) AdvSide() (*field.AdvSide, errors.MessageRejectError) {
	f := new(field.AdvSide)
	err := m.Body.Get(f)
	return f, err
}

//GetAdvSide reads a AdvSide from Advertisement.
func (m Advertisement) GetAdvSide(f *field.AdvSide) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Shares is a required field for Advertisement.
func (m Advertisement) Shares() (*field.Shares, errors.MessageRejectError) {
	f := new(field.Shares)
	err := m.Body.Get(f)
	return f, err
}

//GetShares reads a Shares from Advertisement.
func (m Advertisement) GetShares(f *field.Shares) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for Advertisement.
func (m Advertisement) Price() (*field.Price, errors.MessageRejectError) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from Advertisement.
func (m Advertisement) GetPrice(f *field.Price) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for Advertisement.
func (m Advertisement) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from Advertisement.
func (m Advertisement) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for Advertisement.
func (m Advertisement) TradeDate() (*field.TradeDate, errors.MessageRejectError) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from Advertisement.
func (m Advertisement) GetTradeDate(f *field.TradeDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for Advertisement.
func (m Advertisement) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from Advertisement.
func (m Advertisement) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for Advertisement.
func (m Advertisement) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from Advertisement.
func (m Advertisement) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//URLLink is a non-required field for Advertisement.
func (m Advertisement) URLLink() (*field.URLLink, errors.MessageRejectError) {
	f := new(field.URLLink)
	err := m.Body.Get(f)
	return f, err
}

//GetURLLink reads a URLLink from Advertisement.
func (m Advertisement) GetURLLink(f *field.URLLink) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastMkt is a non-required field for Advertisement.
func (m Advertisement) LastMkt() (*field.LastMkt, errors.MessageRejectError) {
	f := new(field.LastMkt)
	err := m.Body.Get(f)
	return f, err
}

//GetLastMkt reads a LastMkt from Advertisement.
func (m Advertisement) GetLastMkt(f *field.LastMkt) errors.MessageRejectError {
	return m.Body.Get(f)
}
