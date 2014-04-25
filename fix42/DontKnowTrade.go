package fix42

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//DontKnowTrade msg type = Q.
type DontKnowTrade struct {
	message.Message
}

//DontKnowTradeBuilder builds DontKnowTrade messages.
type DontKnowTradeBuilder struct {
	message.MessageBuilder
}

//CreateDontKnowTradeBuilder returns an initialized DontKnowTradeBuilder with specified required fields.
func CreateDontKnowTradeBuilder(
	orderid field.OrderID,
	execid field.ExecID,
	dkreason field.DKReason,
	symbol field.Symbol,
	side field.Side) DontKnowTradeBuilder {
	var builder DontKnowTradeBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("Q"))
	builder.Body.Set(orderid)
	builder.Body.Set(execid)
	builder.Body.Set(dkreason)
	builder.Body.Set(symbol)
	builder.Body.Set(side)
	return builder
}

//OrderID is a required field for DontKnowTrade.
func (m DontKnowTrade) OrderID() (*field.OrderID, errors.MessageRejectError) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from DontKnowTrade.
func (m DontKnowTrade) GetOrderID(f *field.OrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecID is a required field for DontKnowTrade.
func (m DontKnowTrade) ExecID() (*field.ExecID, errors.MessageRejectError) {
	f := new(field.ExecID)
	err := m.Body.Get(f)
	return f, err
}

//GetExecID reads a ExecID from DontKnowTrade.
func (m DontKnowTrade) GetExecID(f *field.ExecID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DKReason is a required field for DontKnowTrade.
func (m DontKnowTrade) DKReason() (*field.DKReason, errors.MessageRejectError) {
	f := new(field.DKReason)
	err := m.Body.Get(f)
	return f, err
}

//GetDKReason reads a DKReason from DontKnowTrade.
func (m DontKnowTrade) GetDKReason(f *field.DKReason) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a required field for DontKnowTrade.
func (m DontKnowTrade) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from DontKnowTrade.
func (m DontKnowTrade) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for DontKnowTrade.
func (m DontKnowTrade) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from DontKnowTrade.
func (m DontKnowTrade) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for DontKnowTrade.
func (m DontKnowTrade) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from DontKnowTrade.
func (m DontKnowTrade) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IDSource is a non-required field for DontKnowTrade.
func (m DontKnowTrade) IDSource() (*field.IDSource, errors.MessageRejectError) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetIDSource reads a IDSource from DontKnowTrade.
func (m DontKnowTrade) GetIDSource(f *field.IDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for DontKnowTrade.
func (m DontKnowTrade) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from DontKnowTrade.
func (m DontKnowTrade) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for DontKnowTrade.
func (m DontKnowTrade) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from DontKnowTrade.
func (m DontKnowTrade) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDay is a non-required field for DontKnowTrade.
func (m DontKnowTrade) MaturityDay() (*field.MaturityDay, errors.MessageRejectError) {
	f := new(field.MaturityDay)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDay reads a MaturityDay from DontKnowTrade.
func (m DontKnowTrade) GetMaturityDay(f *field.MaturityDay) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for DontKnowTrade.
func (m DontKnowTrade) PutOrCall() (*field.PutOrCall, errors.MessageRejectError) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from DontKnowTrade.
func (m DontKnowTrade) GetPutOrCall(f *field.PutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for DontKnowTrade.
func (m DontKnowTrade) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from DontKnowTrade.
func (m DontKnowTrade) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for DontKnowTrade.
func (m DontKnowTrade) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from DontKnowTrade.
func (m DontKnowTrade) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for DontKnowTrade.
func (m DontKnowTrade) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from DontKnowTrade.
func (m DontKnowTrade) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for DontKnowTrade.
func (m DontKnowTrade) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from DontKnowTrade.
func (m DontKnowTrade) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for DontKnowTrade.
func (m DontKnowTrade) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from DontKnowTrade.
func (m DontKnowTrade) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for DontKnowTrade.
func (m DontKnowTrade) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from DontKnowTrade.
func (m DontKnowTrade) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for DontKnowTrade.
func (m DontKnowTrade) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from DontKnowTrade.
func (m DontKnowTrade) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for DontKnowTrade.
func (m DontKnowTrade) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from DontKnowTrade.
func (m DontKnowTrade) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for DontKnowTrade.
func (m DontKnowTrade) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from DontKnowTrade.
func (m DontKnowTrade) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for DontKnowTrade.
func (m DontKnowTrade) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from DontKnowTrade.
func (m DontKnowTrade) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for DontKnowTrade.
func (m DontKnowTrade) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from DontKnowTrade.
func (m DontKnowTrade) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for DontKnowTrade.
func (m DontKnowTrade) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from DontKnowTrade.
func (m DontKnowTrade) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for DontKnowTrade.
func (m DontKnowTrade) OrderQty() (*field.OrderQty, errors.MessageRejectError) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from DontKnowTrade.
func (m DontKnowTrade) GetOrderQty(f *field.OrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashOrderQty is a non-required field for DontKnowTrade.
func (m DontKnowTrade) CashOrderQty() (*field.CashOrderQty, errors.MessageRejectError) {
	f := new(field.CashOrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetCashOrderQty reads a CashOrderQty from DontKnowTrade.
func (m DontKnowTrade) GetCashOrderQty(f *field.CashOrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastShares is a non-required field for DontKnowTrade.
func (m DontKnowTrade) LastShares() (*field.LastShares, errors.MessageRejectError) {
	f := new(field.LastShares)
	err := m.Body.Get(f)
	return f, err
}

//GetLastShares reads a LastShares from DontKnowTrade.
func (m DontKnowTrade) GetLastShares(f *field.LastShares) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastPx is a non-required field for DontKnowTrade.
func (m DontKnowTrade) LastPx() (*field.LastPx, errors.MessageRejectError) {
	f := new(field.LastPx)
	err := m.Body.Get(f)
	return f, err
}

//GetLastPx reads a LastPx from DontKnowTrade.
func (m DontKnowTrade) GetLastPx(f *field.LastPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for DontKnowTrade.
func (m DontKnowTrade) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from DontKnowTrade.
func (m DontKnowTrade) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for DontKnowTrade.
func (m DontKnowTrade) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from DontKnowTrade.
func (m DontKnowTrade) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for DontKnowTrade.
func (m DontKnowTrade) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from DontKnowTrade.
func (m DontKnowTrade) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}
