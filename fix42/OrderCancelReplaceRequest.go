package fix42

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//OrderCancelReplaceRequest msg type = G.
type OrderCancelReplaceRequest struct {
	message.Message
}

//OrderCancelReplaceRequestBuilder builds OrderCancelReplaceRequest messages.
type OrderCancelReplaceRequestBuilder struct {
	message.MessageBuilder
}

//CreateOrderCancelReplaceRequestBuilder returns an initialized OrderCancelReplaceRequestBuilder with specified required fields.
func CreateOrderCancelReplaceRequestBuilder(
	origclordid *field.OrigClOrdIDField,
	clordid *field.ClOrdIDField,
	handlinst *field.HandlInstField,
	symbol *field.SymbolField,
	side *field.SideField,
	transacttime *field.TransactTimeField,
	ordtype *field.OrdTypeField) OrderCancelReplaceRequestBuilder {
	var builder OrderCancelReplaceRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX42))
	builder.Header.Set(field.NewMsgType("G"))
	builder.Body.Set(origclordid)
	builder.Body.Set(clordid)
	builder.Body.Set(handlinst)
	builder.Body.Set(symbol)
	builder.Body.Set(side)
	builder.Body.Set(transacttime)
	builder.Body.Set(ordtype)
	return builder
}

//OrderID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrderID() (*field.OrderIDField, errors.MessageRejectError) {
	f := &field.OrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrderID(f *field.OrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClientID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ClientID() (*field.ClientIDField, errors.MessageRejectError) {
	f := &field.ClientIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClientID reads a ClientID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetClientID(f *field.ClientIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecBroker is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ExecBroker() (*field.ExecBrokerField, errors.MessageRejectError) {
	f := &field.ExecBrokerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecBroker reads a ExecBroker from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetExecBroker(f *field.ExecBrokerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigClOrdID is a required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrigClOrdID() (*field.OrigClOrdIDField, errors.MessageRejectError) {
	f := &field.OrigClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigClOrdID reads a OrigClOrdID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrigClOrdID(f *field.OrigClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ClOrdID() (*field.ClOrdIDField, errors.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetClOrdID(f *field.ClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ListID() (*field.ListIDField, errors.MessageRejectError) {
	f := &field.ListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetListID(f *field.ListIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Account() (*field.AccountField, errors.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetAccount(f *field.AccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAllocs is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NoAllocs() (*field.NoAllocsField, errors.MessageRejectError) {
	f := &field.NoAllocsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoAllocs reads a NoAllocs from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNoAllocs(f *field.NoAllocsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlmntTyp is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SettlmntTyp() (*field.SettlmntTypField, errors.MessageRejectError) {
	f := &field.SettlmntTypField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlmntTyp reads a SettlmntTyp from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSettlmntTyp(f *field.SettlmntTypField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) FutSettDate() (*field.FutSettDateField, errors.MessageRejectError) {
	f := &field.FutSettDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate reads a FutSettDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetFutSettDate(f *field.FutSettDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HandlInst is a required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) HandlInst() (*field.HandlInstField, errors.MessageRejectError) {
	f := &field.HandlInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHandlInst reads a HandlInst from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetHandlInst(f *field.HandlInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecInst is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ExecInst() (*field.ExecInstField, errors.MessageRejectError) {
	f := &field.ExecInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecInst reads a ExecInst from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetExecInst(f *field.ExecInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinQty is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MinQty() (*field.MinQtyField, errors.MessageRejectError) {
	f := &field.MinQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinQty reads a MinQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMinQty(f *field.MinQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxFloor is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MaxFloor() (*field.MaxFloorField, errors.MessageRejectError) {
	f := &field.MaxFloorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxFloor reads a MaxFloor from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMaxFloor(f *field.MaxFloorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestination is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ExDestination() (*field.ExDestinationField, errors.MessageRejectError) {
	f := &field.ExDestinationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExDestination reads a ExDestination from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetExDestination(f *field.ExDestinationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NoTradingSessions() (*field.NoTradingSessionsField, errors.MessageRejectError) {
	f := &field.NoTradingSessionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNoTradingSessions(f *field.NoTradingSessionsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IDSource is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) IDSource() (*field.IDSourceField, errors.MessageRejectError) {
	f := &field.IDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIDSource reads a IDSource from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetIDSource(f *field.IDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDay is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MaturityDay() (*field.MaturityDayField, errors.MessageRejectError) {
	f := &field.MaturityDayField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDay reads a MaturityDay from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMaturityDay(f *field.MaturityDayField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PutOrCall() (*field.PutOrCallField, errors.MessageRejectError) {
	f := &field.PutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPutOrCall(f *field.PutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrderQty() (*field.OrderQtyField, errors.MessageRejectError) {
	f := &field.OrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrderQty(f *field.OrderQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashOrderQty is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CashOrderQty() (*field.CashOrderQtyField, errors.MessageRejectError) {
	f := &field.CashOrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashOrderQty reads a CashOrderQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCashOrderQty(f *field.CashOrderQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrdType() (*field.OrdTypeField, errors.MessageRejectError) {
	f := &field.OrdTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrdType(f *field.OrdTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Price() (*field.PriceField, errors.MessageRejectError) {
	f := &field.PriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPrice(f *field.PriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StopPx is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) StopPx() (*field.StopPxField, errors.MessageRejectError) {
	f := &field.StopPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStopPx reads a StopPx from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetStopPx(f *field.StopPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegDifference is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PegDifference() (*field.PegDifferenceField, errors.MessageRejectError) {
	f := &field.PegDifferenceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegDifference reads a PegDifference from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPegDifference(f *field.PegDifferenceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionInst is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DiscretionInst() (*field.DiscretionInstField, errors.MessageRejectError) {
	f := &field.DiscretionInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionInst reads a DiscretionInst from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDiscretionInst(f *field.DiscretionInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffset is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DiscretionOffset() (*field.DiscretionOffsetField, errors.MessageRejectError) {
	f := &field.DiscretionOffsetField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffset reads a DiscretionOffset from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDiscretionOffset(f *field.DiscretionOffsetField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ComplianceID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ComplianceID() (*field.ComplianceIDField, errors.MessageRejectError) {
	f := &field.ComplianceIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetComplianceID reads a ComplianceID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetComplianceID(f *field.ComplianceIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SolicitedFlag is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SolicitedFlag() (*field.SolicitedFlagField, errors.MessageRejectError) {
	f := &field.SolicitedFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSolicitedFlag reads a SolicitedFlag from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSolicitedFlag(f *field.SolicitedFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeInForce is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TimeInForce() (*field.TimeInForceField, errors.MessageRejectError) {
	f := &field.TimeInForceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeInForce reads a TimeInForce from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTimeInForce(f *field.TimeInForceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) EffectiveTime() (*field.EffectiveTimeField, errors.MessageRejectError) {
	f := &field.EffectiveTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetEffectiveTime(f *field.EffectiveTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ExpireDate() (*field.ExpireDateField, errors.MessageRejectError) {
	f := &field.ExpireDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireDate reads a ExpireDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetExpireDate(f *field.ExpireDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ExpireTime() (*field.ExpireTimeField, errors.MessageRejectError) {
	f := &field.ExpireTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetExpireTime(f *field.ExpireTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GTBookingInst is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GTBookingInst() (*field.GTBookingInstField, errors.MessageRejectError) {
	f := &field.GTBookingInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGTBookingInst reads a GTBookingInst from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetGTBookingInst(f *field.GTBookingInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Commission is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Commission() (*field.CommissionField, errors.MessageRejectError) {
	f := &field.CommissionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommission reads a Commission from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCommission(f *field.CommissionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CommType() (*field.CommTypeField, errors.MessageRejectError) {
	f := &field.CommTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommType reads a CommType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCommType(f *field.CommTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Rule80A is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Rule80A() (*field.Rule80AField, errors.MessageRejectError) {
	f := &field.Rule80AField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRule80A reads a Rule80A from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetRule80A(f *field.Rule80AField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ForexReq is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ForexReq() (*field.ForexReqField, errors.MessageRejectError) {
	f := &field.ForexReqField{}
	err := m.Body.Get(f)
	return f, err
}

//GetForexReq reads a ForexReq from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetForexReq(f *field.ForexReqField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SettlCurrency() (*field.SettlCurrencyField, errors.MessageRejectError) {
	f := &field.SettlCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSettlCurrency(f *field.SettlCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate2 is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) FutSettDate2() (*field.FutSettDate2Field, errors.MessageRejectError) {
	f := &field.FutSettDate2Field{}
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate2 reads a FutSettDate2 from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetFutSettDate2(f *field.FutSettDate2Field) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty2 is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrderQty2() (*field.OrderQty2Field, errors.MessageRejectError) {
	f := &field.OrderQty2Field{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty2 reads a OrderQty2 from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrderQty2(f *field.OrderQty2Field) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OpenClose is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OpenClose() (*field.OpenCloseField, errors.MessageRejectError) {
	f := &field.OpenCloseField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOpenClose reads a OpenClose from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOpenClose(f *field.OpenCloseField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CoveredOrUncovered is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CoveredOrUncovered() (*field.CoveredOrUncoveredField, errors.MessageRejectError) {
	f := &field.CoveredOrUncoveredField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCoveredOrUncovered reads a CoveredOrUncovered from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCoveredOrUncovered(f *field.CoveredOrUncoveredField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustomerOrFirm is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CustomerOrFirm() (*field.CustomerOrFirmField, errors.MessageRejectError) {
	f := &field.CustomerOrFirmField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCustomerOrFirm reads a CustomerOrFirm from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCustomerOrFirm(f *field.CustomerOrFirmField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxShow is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MaxShow() (*field.MaxShowField, errors.MessageRejectError) {
	f := &field.MaxShowField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxShow reads a MaxShow from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMaxShow(f *field.MaxShowField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocateReqd is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) LocateReqd() (*field.LocateReqdField, errors.MessageRejectError) {
	f := &field.LocateReqdField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocateReqd reads a LocateReqd from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetLocateReqd(f *field.LocateReqdField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingFirm is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ClearingFirm() (*field.ClearingFirmField, errors.MessageRejectError) {
	f := &field.ClearingFirmField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingFirm reads a ClearingFirm from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetClearingFirm(f *field.ClearingFirmField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingAccount is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ClearingAccount() (*field.ClearingAccountField, errors.MessageRejectError) {
	f := &field.ClearingAccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingAccount reads a ClearingAccount from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetClearingAccount(f *field.ClearingAccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}
