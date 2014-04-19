package fix40

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
	noorders field.NoOrders,
	side field.Side,
	symbol field.Symbol,
	shares field.Shares,
	avgpx field.AvgPx,
	tradedate field.TradeDate,
	noallocs field.NoAllocs) AllocationBuilder {
	var builder AllocationBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(allocid)
	builder.Body.Set(alloctranstype)
	builder.Body.Set(noorders)
	builder.Body.Set(side)
	builder.Body.Set(symbol)
	builder.Body.Set(shares)
	builder.Body.Set(avgpx)
	builder.Body.Set(tradedate)
	builder.Body.Set(noallocs)
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

//NoOrders is a required field for Allocation.
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

//Issuer is a non-required field for Allocation.
func (m Allocation) Issuer() (field.Issuer, errors.MessageRejectError) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for Allocation.
func (m Allocation) SecurityDesc() (field.SecurityDesc, errors.MessageRejectError) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Shares is a required field for Allocation.
func (m Allocation) Shares() (field.Shares, errors.MessageRejectError) {
	var f field.Shares
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

//NetMoney is a non-required field for Allocation.
func (m Allocation) NetMoney() (field.NetMoney, errors.MessageRejectError) {
	var f field.NetMoney
	err := m.Body.Get(&f)
	return f, err
}

//NoMiscFees is a non-required field for Allocation.
func (m Allocation) NoMiscFees() (field.NoMiscFees, errors.MessageRejectError) {
	var f field.NoMiscFees
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrAmt is a non-required field for Allocation.
func (m Allocation) SettlCurrAmt() (field.SettlCurrAmt, errors.MessageRejectError) {
	var f field.SettlCurrAmt
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrency is a non-required field for Allocation.
func (m Allocation) SettlCurrency() (field.SettlCurrency, errors.MessageRejectError) {
	var f field.SettlCurrency
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

//NoAllocs is a required field for Allocation.
func (m Allocation) NoAllocs() (field.NoAllocs, errors.MessageRejectError) {
	var f field.NoAllocs
	err := m.Body.Get(&f)
	return f, err
}
