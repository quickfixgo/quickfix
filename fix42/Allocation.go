package fix42

import (
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

//NewAllocationBuilder returns an initialized AllocationBuilder with specified required fields.
func NewAllocationBuilder(
	allocid field.AllocID,
	alloctranstype field.AllocTransType,
	side field.Side,
	symbol field.Symbol,
	shares field.Shares,
	avgpx field.AvgPx,
	tradedate field.TradeDate) *AllocationBuilder {
	builder := new(AllocationBuilder)
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
func (m *Allocation) AllocID() (*field.AllocID, error) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}

//AllocTransType is a required field for Allocation.
func (m *Allocation) AllocTransType() (*field.AllocTransType, error) {
	f := new(field.AllocTransType)
	err := m.Body.Get(f)
	return f, err
}

//RefAllocID is a non-required field for Allocation.
func (m *Allocation) RefAllocID() (*field.RefAllocID, error) {
	f := new(field.RefAllocID)
	err := m.Body.Get(f)
	return f, err
}

//AllocLinkID is a non-required field for Allocation.
func (m *Allocation) AllocLinkID() (*field.AllocLinkID, error) {
	f := new(field.AllocLinkID)
	err := m.Body.Get(f)
	return f, err
}

//AllocLinkType is a non-required field for Allocation.
func (m *Allocation) AllocLinkType() (*field.AllocLinkType, error) {
	f := new(field.AllocLinkType)
	err := m.Body.Get(f)
	return f, err
}

//NoOrders is a non-required field for Allocation.
func (m *Allocation) NoOrders() (*field.NoOrders, error) {
	f := new(field.NoOrders)
	err := m.Body.Get(f)
	return f, err
}

//NoExecs is a non-required field for Allocation.
func (m *Allocation) NoExecs() (*field.NoExecs, error) {
	f := new(field.NoExecs)
	err := m.Body.Get(f)
	return f, err
}

//Side is a required field for Allocation.
func (m *Allocation) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a required field for Allocation.
func (m *Allocation) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for Allocation.
func (m *Allocation) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for Allocation.
func (m *Allocation) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//IDSource is a non-required field for Allocation.
func (m *Allocation) IDSource() (*field.IDSource, error) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for Allocation.
func (m *Allocation) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for Allocation.
func (m *Allocation) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDay is a non-required field for Allocation.
func (m *Allocation) MaturityDay() (*field.MaturityDay, error) {
	f := new(field.MaturityDay)
	err := m.Body.Get(f)
	return f, err
}

//PutOrCall is a non-required field for Allocation.
func (m *Allocation) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for Allocation.
func (m *Allocation) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for Allocation.
func (m *Allocation) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for Allocation.
func (m *Allocation) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for Allocation.
func (m *Allocation) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for Allocation.
func (m *Allocation) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for Allocation.
func (m *Allocation) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for Allocation.
func (m *Allocation) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for Allocation.
func (m *Allocation) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for Allocation.
func (m *Allocation) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for Allocation.
func (m *Allocation) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for Allocation.
func (m *Allocation) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Shares is a required field for Allocation.
func (m *Allocation) Shares() (*field.Shares, error) {
	f := new(field.Shares)
	err := m.Body.Get(f)
	return f, err
}

//LastMkt is a non-required field for Allocation.
func (m *Allocation) LastMkt() (*field.LastMkt, error) {
	f := new(field.LastMkt)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionID is a non-required field for Allocation.
func (m *Allocation) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//AvgPx is a required field for Allocation.
func (m *Allocation) AvgPx() (*field.AvgPx, error) {
	f := new(field.AvgPx)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for Allocation.
func (m *Allocation) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//AvgPrxPrecision is a non-required field for Allocation.
func (m *Allocation) AvgPrxPrecision() (*field.AvgPrxPrecision, error) {
	f := new(field.AvgPrxPrecision)
	err := m.Body.Get(f)
	return f, err
}

//TradeDate is a required field for Allocation.
func (m *Allocation) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a non-required field for Allocation.
func (m *Allocation) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//SettlmntTyp is a non-required field for Allocation.
func (m *Allocation) SettlmntTyp() (*field.SettlmntTyp, error) {
	f := new(field.SettlmntTyp)
	err := m.Body.Get(f)
	return f, err
}

//FutSettDate is a non-required field for Allocation.
func (m *Allocation) FutSettDate() (*field.FutSettDate, error) {
	f := new(field.FutSettDate)
	err := m.Body.Get(f)
	return f, err
}

//GrossTradeAmt is a non-required field for Allocation.
func (m *Allocation) GrossTradeAmt() (*field.GrossTradeAmt, error) {
	f := new(field.GrossTradeAmt)
	err := m.Body.Get(f)
	return f, err
}

//NetMoney is a non-required field for Allocation.
func (m *Allocation) NetMoney() (*field.NetMoney, error) {
	f := new(field.NetMoney)
	err := m.Body.Get(f)
	return f, err
}

//OpenClose is a non-required field for Allocation.
func (m *Allocation) OpenClose() (*field.OpenClose, error) {
	f := new(field.OpenClose)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for Allocation.
func (m *Allocation) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for Allocation.
func (m *Allocation) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for Allocation.
func (m *Allocation) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//NumDaysInterest is a non-required field for Allocation.
func (m *Allocation) NumDaysInterest() (*field.NumDaysInterest, error) {
	f := new(field.NumDaysInterest)
	err := m.Body.Get(f)
	return f, err
}

//AccruedInterestRate is a non-required field for Allocation.
func (m *Allocation) AccruedInterestRate() (*field.AccruedInterestRate, error) {
	f := new(field.AccruedInterestRate)
	err := m.Body.Get(f)
	return f, err
}

//NoAllocs is a non-required field for Allocation.
func (m *Allocation) NoAllocs() (*field.NoAllocs, error) {
	f := new(field.NoAllocs)
	err := m.Body.Get(f)
	return f, err
}
