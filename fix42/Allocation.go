package fix42

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Allocation msg type = J.
type Allocation struct {
	message.Message
}

//AllocationBuilder builds Allocation messages.
type AllocationBuilder struct {
	message.MessageBuilder
}

//CreateAllocationBuilder returns an initialized AllocationBuilder with specified required fields.
func CreateAllocationBuilder(
	allocid field.AllocID,
	alloctranstype field.AllocTransType,
	side field.Side,
	symbol field.Symbol,
	shares field.Shares,
	avgpx field.AvgPx,
	tradedate field.TradeDate) AllocationBuilder {
	var builder AllocationBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(allocid)
	builder.Body.Set(alloctranstype)
	builder.Body.Set(side)
	builder.Body.Set(symbol)
	builder.Body.Set(shares)
	builder.Body.Set(avgpx)
	builder.Body.Set(tradedate)
	return builder
}

//AllocID is a required field for Allocation.
func (m Allocation) AllocID() (field.AllocID, errors.MessageRejectError) {
	var f field.AllocID
	err := m.Body.Get(&f)
	return f, err
}

//AllocTransType is a required field for Allocation.
func (m Allocation) AllocTransType() (field.AllocTransType, errors.MessageRejectError) {
	var f field.AllocTransType
	err := m.Body.Get(&f)
	return f, err
}

//RefAllocID is a non-required field for Allocation.
func (m Allocation) RefAllocID() (field.RefAllocID, errors.MessageRejectError) {
	var f field.RefAllocID
	err := m.Body.Get(&f)
	return f, err
}

//AllocLinkID is a non-required field for Allocation.
func (m Allocation) AllocLinkID() (field.AllocLinkID, errors.MessageRejectError) {
	var f field.AllocLinkID
	err := m.Body.Get(&f)
	return f, err
}

//AllocLinkType is a non-required field for Allocation.
func (m Allocation) AllocLinkType() (field.AllocLinkType, errors.MessageRejectError) {
	var f field.AllocLinkType
	err := m.Body.Get(&f)
	return f, err
}

//NoOrders is a non-required field for Allocation.
func (m Allocation) NoOrders() (field.NoOrders, errors.MessageRejectError) {
	var f field.NoOrders
	err := m.Body.Get(&f)
	return f, err
}

//NoExecs is a non-required field for Allocation.
func (m Allocation) NoExecs() (field.NoExecs, errors.MessageRejectError) {
	var f field.NoExecs
	err := m.Body.Get(&f)
	return f, err
}

//Side is a required field for Allocation.
func (m Allocation) Side() (field.Side, errors.MessageRejectError) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a required field for Allocation.
func (m Allocation) Symbol() (field.Symbol, errors.MessageRejectError) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for Allocation.
func (m Allocation) SymbolSfx() (field.SymbolSfx, errors.MessageRejectError) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for Allocation.
func (m Allocation) SecurityID() (field.SecurityID, errors.MessageRejectError) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//IDSource is a non-required field for Allocation.
func (m Allocation) IDSource() (field.IDSource, errors.MessageRejectError) {
	var f field.IDSource
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for Allocation.
func (m Allocation) SecurityType() (field.SecurityType, errors.MessageRejectError) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for Allocation.
func (m Allocation) MaturityMonthYear() (field.MaturityMonthYear, errors.MessageRejectError) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDay is a non-required field for Allocation.
func (m Allocation) MaturityDay() (field.MaturityDay, errors.MessageRejectError) {
	var f field.MaturityDay
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for Allocation.
func (m Allocation) PutOrCall() (field.PutOrCall, errors.MessageRejectError) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for Allocation.
func (m Allocation) StrikePrice() (field.StrikePrice, errors.MessageRejectError) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for Allocation.
func (m Allocation) OptAttribute() (field.OptAttribute, errors.MessageRejectError) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for Allocation.
func (m Allocation) ContractMultiplier() (field.ContractMultiplier, errors.MessageRejectError) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for Allocation.
func (m Allocation) CouponRate() (field.CouponRate, errors.MessageRejectError) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for Allocation.
func (m Allocation) SecurityExchange() (field.SecurityExchange, errors.MessageRejectError) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for Allocation.
func (m Allocation) Issuer() (field.Issuer, errors.MessageRejectError) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for Allocation.
func (m Allocation) EncodedIssuerLen() (field.EncodedIssuerLen, errors.MessageRejectError) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for Allocation.
func (m Allocation) EncodedIssuer() (field.EncodedIssuer, errors.MessageRejectError) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for Allocation.
func (m Allocation) SecurityDesc() (field.SecurityDesc, errors.MessageRejectError) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for Allocation.
func (m Allocation) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, errors.MessageRejectError) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for Allocation.
func (m Allocation) EncodedSecurityDesc() (field.EncodedSecurityDesc, errors.MessageRejectError) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Shares is a required field for Allocation.
func (m Allocation) Shares() (field.Shares, errors.MessageRejectError) {
	var f field.Shares
	err := m.Body.Get(&f)
	return f, err
}

//LastMkt is a non-required field for Allocation.
func (m Allocation) LastMkt() (field.LastMkt, errors.MessageRejectError) {
	var f field.LastMkt
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for Allocation.
func (m Allocation) TradingSessionID() (field.TradingSessionID, errors.MessageRejectError) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//AvgPx is a required field for Allocation.
func (m Allocation) AvgPx() (field.AvgPx, errors.MessageRejectError) {
	var f field.AvgPx
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for Allocation.
func (m Allocation) Currency() (field.Currency, errors.MessageRejectError) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//AvgPrxPrecision is a non-required field for Allocation.
func (m Allocation) AvgPrxPrecision() (field.AvgPrxPrecision, errors.MessageRejectError) {
	var f field.AvgPrxPrecision
	err := m.Body.Get(&f)
	return f, err
}

//TradeDate is a required field for Allocation.
func (m Allocation) TradeDate() (field.TradeDate, errors.MessageRejectError) {
	var f field.TradeDate
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for Allocation.
func (m Allocation) TransactTime() (field.TransactTime, errors.MessageRejectError) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//SettlmntTyp is a non-required field for Allocation.
func (m Allocation) SettlmntTyp() (field.SettlmntTyp, errors.MessageRejectError) {
	var f field.SettlmntTyp
	err := m.Body.Get(&f)
	return f, err
}

//FutSettDate is a non-required field for Allocation.
func (m Allocation) FutSettDate() (field.FutSettDate, errors.MessageRejectError) {
	var f field.FutSettDate
	err := m.Body.Get(&f)
	return f, err
}

//GrossTradeAmt is a non-required field for Allocation.
func (m Allocation) GrossTradeAmt() (field.GrossTradeAmt, errors.MessageRejectError) {
	var f field.GrossTradeAmt
	err := m.Body.Get(&f)
	return f, err
}

//NetMoney is a non-required field for Allocation.
func (m Allocation) NetMoney() (field.NetMoney, errors.MessageRejectError) {
	var f field.NetMoney
	err := m.Body.Get(&f)
	return f, err
}

//OpenClose is a non-required field for Allocation.
func (m Allocation) OpenClose() (field.OpenClose, errors.MessageRejectError) {
	var f field.OpenClose
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for Allocation.
func (m Allocation) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for Allocation.
func (m Allocation) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for Allocation.
func (m Allocation) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//NumDaysInterest is a non-required field for Allocation.
func (m Allocation) NumDaysInterest() (field.NumDaysInterest, errors.MessageRejectError) {
	var f field.NumDaysInterest
	err := m.Body.Get(&f)
	return f, err
}

//AccruedInterestRate is a non-required field for Allocation.
func (m Allocation) AccruedInterestRate() (field.AccruedInterestRate, errors.MessageRejectError) {
	var f field.AccruedInterestRate
	err := m.Body.Get(&f)
	return f, err
}

//NoAllocs is a non-required field for Allocation.
func (m Allocation) NoAllocs() (field.NoAllocs, errors.MessageRejectError) {
	var f field.NoAllocs
	err := m.Body.Get(&f)
	return f, err
}
