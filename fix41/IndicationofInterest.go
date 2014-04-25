package fix41

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//IndicationofInterest msg type = 6.
type IndicationofInterest struct {
	message.Message
}

//IndicationofInterestBuilder builds IndicationofInterest messages.
type IndicationofInterestBuilder struct {
	message.MessageBuilder
}

//CreateIndicationofInterestBuilder returns an initialized IndicationofInterestBuilder with specified required fields.
func CreateIndicationofInterestBuilder(
	ioiid field.IOIid,
	ioitranstype field.IOITransType,
	symbol field.Symbol,
	side field.Side,
	ioishares field.IOIShares) IndicationofInterestBuilder {
	var builder IndicationofInterestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("6"))
	builder.Body.Set(ioiid)
	builder.Body.Set(ioitranstype)
	builder.Body.Set(symbol)
	builder.Body.Set(side)
	builder.Body.Set(ioishares)
	return builder
}

//IOIid is a required field for IndicationofInterest.
func (m IndicationofInterest) IOIid() (*field.IOIid, errors.MessageRejectError) {
	f := new(field.IOIid)
	err := m.Body.Get(f)
	return f, err
}

//GetIOIid reads a IOIid from IndicationofInterest.
func (m IndicationofInterest) GetIOIid(f *field.IOIid) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOITransType is a required field for IndicationofInterest.
func (m IndicationofInterest) IOITransType() (*field.IOITransType, errors.MessageRejectError) {
	f := new(field.IOITransType)
	err := m.Body.Get(f)
	return f, err
}

//GetIOITransType reads a IOITransType from IndicationofInterest.
func (m IndicationofInterest) GetIOITransType(f *field.IOITransType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIRefID is a non-required field for IndicationofInterest.
func (m IndicationofInterest) IOIRefID() (*field.IOIRefID, errors.MessageRejectError) {
	f := new(field.IOIRefID)
	err := m.Body.Get(f)
	return f, err
}

//GetIOIRefID reads a IOIRefID from IndicationofInterest.
func (m IndicationofInterest) GetIOIRefID(f *field.IOIRefID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a required field for IndicationofInterest.
func (m IndicationofInterest) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from IndicationofInterest.
func (m IndicationofInterest) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for IndicationofInterest.
func (m IndicationofInterest) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from IndicationofInterest.
func (m IndicationofInterest) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for IndicationofInterest.
func (m IndicationofInterest) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from IndicationofInterest.
func (m IndicationofInterest) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IDSource is a non-required field for IndicationofInterest.
func (m IndicationofInterest) IDSource() (*field.IDSource, errors.MessageRejectError) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetIDSource reads a IDSource from IndicationofInterest.
func (m IndicationofInterest) GetIDSource(f *field.IDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for IndicationofInterest.
func (m IndicationofInterest) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from IndicationofInterest.
func (m IndicationofInterest) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for IndicationofInterest.
func (m IndicationofInterest) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from IndicationofInterest.
func (m IndicationofInterest) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDay is a non-required field for IndicationofInterest.
func (m IndicationofInterest) MaturityDay() (*field.MaturityDay, errors.MessageRejectError) {
	f := new(field.MaturityDay)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDay reads a MaturityDay from IndicationofInterest.
func (m IndicationofInterest) GetMaturityDay(f *field.MaturityDay) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for IndicationofInterest.
func (m IndicationofInterest) PutOrCall() (*field.PutOrCall, errors.MessageRejectError) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from IndicationofInterest.
func (m IndicationofInterest) GetPutOrCall(f *field.PutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for IndicationofInterest.
func (m IndicationofInterest) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from IndicationofInterest.
func (m IndicationofInterest) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for IndicationofInterest.
func (m IndicationofInterest) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from IndicationofInterest.
func (m IndicationofInterest) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for IndicationofInterest.
func (m IndicationofInterest) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from IndicationofInterest.
func (m IndicationofInterest) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for IndicationofInterest.
func (m IndicationofInterest) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from IndicationofInterest.
func (m IndicationofInterest) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for IndicationofInterest.
func (m IndicationofInterest) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from IndicationofInterest.
func (m IndicationofInterest) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for IndicationofInterest.
func (m IndicationofInterest) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from IndicationofInterest.
func (m IndicationofInterest) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIShares is a required field for IndicationofInterest.
func (m IndicationofInterest) IOIShares() (*field.IOIShares, errors.MessageRejectError) {
	f := new(field.IOIShares)
	err := m.Body.Get(f)
	return f, err
}

//GetIOIShares reads a IOIShares from IndicationofInterest.
func (m IndicationofInterest) GetIOIShares(f *field.IOIShares) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for IndicationofInterest.
func (m IndicationofInterest) Price() (*field.Price, errors.MessageRejectError) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from IndicationofInterest.
func (m IndicationofInterest) GetPrice(f *field.Price) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for IndicationofInterest.
func (m IndicationofInterest) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from IndicationofInterest.
func (m IndicationofInterest) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ValidUntilTime is a non-required field for IndicationofInterest.
func (m IndicationofInterest) ValidUntilTime() (*field.ValidUntilTime, errors.MessageRejectError) {
	f := new(field.ValidUntilTime)
	err := m.Body.Get(f)
	return f, err
}

//GetValidUntilTime reads a ValidUntilTime from IndicationofInterest.
func (m IndicationofInterest) GetValidUntilTime(f *field.ValidUntilTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIQltyInd is a non-required field for IndicationofInterest.
func (m IndicationofInterest) IOIQltyInd() (*field.IOIQltyInd, errors.MessageRejectError) {
	f := new(field.IOIQltyInd)
	err := m.Body.Get(f)
	return f, err
}

//GetIOIQltyInd reads a IOIQltyInd from IndicationofInterest.
func (m IndicationofInterest) GetIOIQltyInd(f *field.IOIQltyInd) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIOthSvc is a non-required field for IndicationofInterest.
func (m IndicationofInterest) IOIOthSvc() (*field.IOIOthSvc, errors.MessageRejectError) {
	f := new(field.IOIOthSvc)
	err := m.Body.Get(f)
	return f, err
}

//GetIOIOthSvc reads a IOIOthSvc from IndicationofInterest.
func (m IndicationofInterest) GetIOIOthSvc(f *field.IOIOthSvc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOINaturalFlag is a non-required field for IndicationofInterest.
func (m IndicationofInterest) IOINaturalFlag() (*field.IOINaturalFlag, errors.MessageRejectError) {
	f := new(field.IOINaturalFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetIOINaturalFlag reads a IOINaturalFlag from IndicationofInterest.
func (m IndicationofInterest) GetIOINaturalFlag(f *field.IOINaturalFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoIOIQualifiers is a non-required field for IndicationofInterest.
func (m IndicationofInterest) NoIOIQualifiers() (*field.NoIOIQualifiers, errors.MessageRejectError) {
	f := new(field.NoIOIQualifiers)
	err := m.Body.Get(f)
	return f, err
}

//GetNoIOIQualifiers reads a NoIOIQualifiers from IndicationofInterest.
func (m IndicationofInterest) GetNoIOIQualifiers(f *field.NoIOIQualifiers) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for IndicationofInterest.
func (m IndicationofInterest) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from IndicationofInterest.
func (m IndicationofInterest) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for IndicationofInterest.
func (m IndicationofInterest) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from IndicationofInterest.
func (m IndicationofInterest) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//URLLink is a non-required field for IndicationofInterest.
func (m IndicationofInterest) URLLink() (*field.URLLink, errors.MessageRejectError) {
	f := new(field.URLLink)
	err := m.Body.Get(f)
	return f, err
}

//GetURLLink reads a URLLink from IndicationofInterest.
func (m IndicationofInterest) GetURLLink(f *field.URLLink) errors.MessageRejectError {
	return m.Body.Get(f)
}
