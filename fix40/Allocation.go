package fix40

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
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
	allocid *field.AllocIDField,
	alloctranstype *field.AllocTransTypeField,
	noorders *field.NoOrdersField,
	side *field.SideField,
	symbol *field.SymbolField,
	shares *field.SharesField,
	avgpx *field.AvgPxField,
	tradedate *field.TradeDateField,
	noallocs *field.NoAllocsField) AllocationBuilder {
	var builder AllocationBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX40))
	builder.Header().Set(field.NewMsgType("J"))
	builder.Body().Set(allocid)
	builder.Body().Set(alloctranstype)
	builder.Body().Set(noorders)
	builder.Body().Set(side)
	builder.Body().Set(symbol)
	builder.Body().Set(shares)
	builder.Body().Set(avgpx)
	builder.Body().Set(tradedate)
	builder.Body().Set(noallocs)
	return builder
}

//AllocID is a required field for Allocation.
func (m Allocation) AllocID() (*field.AllocIDField, errors.MessageRejectError) {
	f := &field.AllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from Allocation.
func (m Allocation) GetAllocID(f *field.AllocIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocTransType is a required field for Allocation.
func (m Allocation) AllocTransType() (*field.AllocTransTypeField, errors.MessageRejectError) {
	f := &field.AllocTransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocTransType reads a AllocTransType from Allocation.
func (m Allocation) GetAllocTransType(f *field.AllocTransTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefAllocID is a non-required field for Allocation.
func (m Allocation) RefAllocID() (*field.RefAllocIDField, errors.MessageRejectError) {
	f := &field.RefAllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRefAllocID reads a RefAllocID from Allocation.
func (m Allocation) GetRefAllocID(f *field.RefAllocIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoOrders is a required field for Allocation.
func (m Allocation) NoOrders() (*field.NoOrdersField, errors.MessageRejectError) {
	f := &field.NoOrdersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoOrders reads a NoOrders from Allocation.
func (m Allocation) GetNoOrders(f *field.NoOrdersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoExecs is a non-required field for Allocation.
func (m Allocation) NoExecs() (*field.NoExecsField, errors.MessageRejectError) {
	f := &field.NoExecsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoExecs reads a NoExecs from Allocation.
func (m Allocation) GetNoExecs(f *field.NoExecsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for Allocation.
func (m Allocation) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from Allocation.
func (m Allocation) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a required field for Allocation.
func (m Allocation) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from Allocation.
func (m Allocation) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for Allocation.
func (m Allocation) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from Allocation.
func (m Allocation) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for Allocation.
func (m Allocation) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from Allocation.
func (m Allocation) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IDSource is a non-required field for Allocation.
func (m Allocation) IDSource() (*field.IDSourceField, errors.MessageRejectError) {
	f := &field.IDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIDSource reads a IDSource from Allocation.
func (m Allocation) GetIDSource(f *field.IDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for Allocation.
func (m Allocation) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from Allocation.
func (m Allocation) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for Allocation.
func (m Allocation) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from Allocation.
func (m Allocation) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Shares is a required field for Allocation.
func (m Allocation) Shares() (*field.SharesField, errors.MessageRejectError) {
	f := &field.SharesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetShares reads a Shares from Allocation.
func (m Allocation) GetShares(f *field.SharesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPx is a required field for Allocation.
func (m Allocation) AvgPx() (*field.AvgPxField, errors.MessageRejectError) {
	f := &field.AvgPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPx reads a AvgPx from Allocation.
func (m Allocation) GetAvgPx(f *field.AvgPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for Allocation.
func (m Allocation) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from Allocation.
func (m Allocation) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPrxPrecision is a non-required field for Allocation.
func (m Allocation) AvgPrxPrecision() (*field.AvgPrxPrecisionField, errors.MessageRejectError) {
	f := &field.AvgPrxPrecisionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPrxPrecision reads a AvgPrxPrecision from Allocation.
func (m Allocation) GetAvgPrxPrecision(f *field.AvgPrxPrecisionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a required field for Allocation.
func (m Allocation) TradeDate() (*field.TradeDateField, errors.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from Allocation.
func (m Allocation) GetTradeDate(f *field.TradeDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for Allocation.
func (m Allocation) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from Allocation.
func (m Allocation) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlmntTyp is a non-required field for Allocation.
func (m Allocation) SettlmntTyp() (*field.SettlmntTypField, errors.MessageRejectError) {
	f := &field.SettlmntTypField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlmntTyp reads a SettlmntTyp from Allocation.
func (m Allocation) GetSettlmntTyp(f *field.SettlmntTypField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate is a non-required field for Allocation.
func (m Allocation) FutSettDate() (*field.FutSettDateField, errors.MessageRejectError) {
	f := &field.FutSettDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate reads a FutSettDate from Allocation.
func (m Allocation) GetFutSettDate(f *field.FutSettDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NetMoney is a non-required field for Allocation.
func (m Allocation) NetMoney() (*field.NetMoneyField, errors.MessageRejectError) {
	f := &field.NetMoneyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNetMoney reads a NetMoney from Allocation.
func (m Allocation) GetNetMoney(f *field.NetMoneyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoMiscFees is a non-required field for Allocation.
func (m Allocation) NoMiscFees() (*field.NoMiscFeesField, errors.MessageRejectError) {
	f := &field.NoMiscFeesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoMiscFees reads a NoMiscFees from Allocation.
func (m Allocation) GetNoMiscFees(f *field.NoMiscFeesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrAmt is a non-required field for Allocation.
func (m Allocation) SettlCurrAmt() (*field.SettlCurrAmtField, errors.MessageRejectError) {
	f := &field.SettlCurrAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrAmt reads a SettlCurrAmt from Allocation.
func (m Allocation) GetSettlCurrAmt(f *field.SettlCurrAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for Allocation.
func (m Allocation) SettlCurrency() (*field.SettlCurrencyField, errors.MessageRejectError) {
	f := &field.SettlCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from Allocation.
func (m Allocation) GetSettlCurrency(f *field.SettlCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OpenClose is a non-required field for Allocation.
func (m Allocation) OpenClose() (*field.OpenCloseField, errors.MessageRejectError) {
	f := &field.OpenCloseField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOpenClose reads a OpenClose from Allocation.
func (m Allocation) GetOpenClose(f *field.OpenCloseField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for Allocation.
func (m Allocation) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from Allocation.
func (m Allocation) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAllocs is a required field for Allocation.
func (m Allocation) NoAllocs() (*field.NoAllocsField, errors.MessageRejectError) {
	f := &field.NoAllocsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoAllocs reads a NoAllocs from Allocation.
func (m Allocation) GetNoAllocs(f *field.NoAllocsField) errors.MessageRejectError {
	return m.Body.Get(f)
}
