//Package indicationofinterest msg type = 6.
package indicationofinterest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Message is a IndicationofInterest wrapper for the generic Message type
type Message struct {
	message.Message
}

//IOIid is a required field for IndicationofInterest.
func (m Message) IOIid() (*field.IOIidField, errors.MessageRejectError) {
	f := &field.IOIidField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIid reads a IOIid from IndicationofInterest.
func (m Message) GetIOIid(f *field.IOIidField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOITransType is a required field for IndicationofInterest.
func (m Message) IOITransType() (*field.IOITransTypeField, errors.MessageRejectError) {
	f := &field.IOITransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOITransType reads a IOITransType from IndicationofInterest.
func (m Message) GetIOITransType(f *field.IOITransTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIRefID is a non-required field for IndicationofInterest.
func (m Message) IOIRefID() (*field.IOIRefIDField, errors.MessageRejectError) {
	f := &field.IOIRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIRefID reads a IOIRefID from IndicationofInterest.
func (m Message) GetIOIRefID(f *field.IOIRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a required field for IndicationofInterest.
func (m Message) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from IndicationofInterest.
func (m Message) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for IndicationofInterest.
func (m Message) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from IndicationofInterest.
func (m Message) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for IndicationofInterest.
func (m Message) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from IndicationofInterest.
func (m Message) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IDSource is a non-required field for IndicationofInterest.
func (m Message) IDSource() (*field.IDSourceField, errors.MessageRejectError) {
	f := &field.IDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIDSource reads a IDSource from IndicationofInterest.
func (m Message) GetIDSource(f *field.IDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for IndicationofInterest.
func (m Message) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from IndicationofInterest.
func (m Message) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for IndicationofInterest.
func (m Message) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from IndicationofInterest.
func (m Message) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDay is a non-required field for IndicationofInterest.
func (m Message) MaturityDay() (*field.MaturityDayField, errors.MessageRejectError) {
	f := &field.MaturityDayField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDay reads a MaturityDay from IndicationofInterest.
func (m Message) GetMaturityDay(f *field.MaturityDayField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for IndicationofInterest.
func (m Message) PutOrCall() (*field.PutOrCallField, errors.MessageRejectError) {
	f := &field.PutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from IndicationofInterest.
func (m Message) GetPutOrCall(f *field.PutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for IndicationofInterest.
func (m Message) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from IndicationofInterest.
func (m Message) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for IndicationofInterest.
func (m Message) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from IndicationofInterest.
func (m Message) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for IndicationofInterest.
func (m Message) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from IndicationofInterest.
func (m Message) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for IndicationofInterest.
func (m Message) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from IndicationofInterest.
func (m Message) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for IndicationofInterest.
func (m Message) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from IndicationofInterest.
func (m Message) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for IndicationofInterest.
func (m Message) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from IndicationofInterest.
func (m Message) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for IndicationofInterest.
func (m Message) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from IndicationofInterest.
func (m Message) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for IndicationofInterest.
func (m Message) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from IndicationofInterest.
func (m Message) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for IndicationofInterest.
func (m Message) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from IndicationofInterest.
func (m Message) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for IndicationofInterest.
func (m Message) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from IndicationofInterest.
func (m Message) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for IndicationofInterest.
func (m Message) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from IndicationofInterest.
func (m Message) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for IndicationofInterest.
func (m Message) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from IndicationofInterest.
func (m Message) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIShares is a required field for IndicationofInterest.
func (m Message) IOIShares() (*field.IOISharesField, errors.MessageRejectError) {
	f := &field.IOISharesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIShares reads a IOIShares from IndicationofInterest.
func (m Message) GetIOIShares(f *field.IOISharesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for IndicationofInterest.
func (m Message) Price() (*field.PriceField, errors.MessageRejectError) {
	f := &field.PriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from IndicationofInterest.
func (m Message) GetPrice(f *field.PriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for IndicationofInterest.
func (m Message) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from IndicationofInterest.
func (m Message) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ValidUntilTime is a non-required field for IndicationofInterest.
func (m Message) ValidUntilTime() (*field.ValidUntilTimeField, errors.MessageRejectError) {
	f := &field.ValidUntilTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetValidUntilTime reads a ValidUntilTime from IndicationofInterest.
func (m Message) GetValidUntilTime(f *field.ValidUntilTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIQltyInd is a non-required field for IndicationofInterest.
func (m Message) IOIQltyInd() (*field.IOIQltyIndField, errors.MessageRejectError) {
	f := &field.IOIQltyIndField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIQltyInd reads a IOIQltyInd from IndicationofInterest.
func (m Message) GetIOIQltyInd(f *field.IOIQltyIndField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOINaturalFlag is a non-required field for IndicationofInterest.
func (m Message) IOINaturalFlag() (*field.IOINaturalFlagField, errors.MessageRejectError) {
	f := &field.IOINaturalFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOINaturalFlag reads a IOINaturalFlag from IndicationofInterest.
func (m Message) GetIOINaturalFlag(f *field.IOINaturalFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoIOIQualifiers is a non-required field for IndicationofInterest.
func (m Message) NoIOIQualifiers() (*field.NoIOIQualifiersField, errors.MessageRejectError) {
	f := &field.NoIOIQualifiersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoIOIQualifiers reads a NoIOIQualifiers from IndicationofInterest.
func (m Message) GetNoIOIQualifiers(f *field.NoIOIQualifiersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for IndicationofInterest.
func (m Message) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from IndicationofInterest.
func (m Message) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for IndicationofInterest.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from IndicationofInterest.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for IndicationofInterest.
func (m Message) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from IndicationofInterest.
func (m Message) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for IndicationofInterest.
func (m Message) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from IndicationofInterest.
func (m Message) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//URLLink is a non-required field for IndicationofInterest.
func (m Message) URLLink() (*field.URLLinkField, errors.MessageRejectError) {
	f := &field.URLLinkField{}
	err := m.Body.Get(f)
	return f, err
}

//GetURLLink reads a URLLink from IndicationofInterest.
func (m Message) GetURLLink(f *field.URLLinkField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRoutingIDs is a non-required field for IndicationofInterest.
func (m Message) NoRoutingIDs() (*field.NoRoutingIDsField, errors.MessageRejectError) {
	f := &field.NoRoutingIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRoutingIDs reads a NoRoutingIDs from IndicationofInterest.
func (m Message) GetNoRoutingIDs(f *field.NoRoutingIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SpreadToBenchmark is a non-required field for IndicationofInterest.
func (m Message) SpreadToBenchmark() (*field.SpreadToBenchmarkField, errors.MessageRejectError) {
	f := &field.SpreadToBenchmarkField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSpreadToBenchmark reads a SpreadToBenchmark from IndicationofInterest.
func (m Message) GetSpreadToBenchmark(f *field.SpreadToBenchmarkField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Benchmark is a non-required field for IndicationofInterest.
func (m Message) Benchmark() (*field.BenchmarkField, errors.MessageRejectError) {
	f := &field.BenchmarkField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmark reads a Benchmark from IndicationofInterest.
func (m Message) GetBenchmark(f *field.BenchmarkField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds IndicationofInterest messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for IndicationofInterest.
func Builder(
	ioiid *field.IOIidField,
	ioitranstype *field.IOITransTypeField,
	symbol *field.SymbolField,
	side *field.SideField,
	ioishares *field.IOISharesField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX42))
	builder.Header().Set(field.NewMsgType("6"))
	builder.Body().Set(ioiid)
	builder.Body().Set(ioitranstype)
	builder.Body().Set(symbol)
	builder.Body().Set(side)
	builder.Body().Set(ioishares)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX42, "6", r
}
