package fix42

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//ExecutionReport msg type = 8.
type ExecutionReport struct {
	message.Message
}

//ExecutionReportBuilder builds ExecutionReport messages.
type ExecutionReportBuilder struct {
	message.MessageBuilder
}

//CreateExecutionReportBuilder returns an initialized ExecutionReportBuilder with specified required fields.
func CreateExecutionReportBuilder(
	orderid *field.OrderIDField,
	execid *field.ExecIDField,
	exectranstype *field.ExecTransTypeField,
	exectype *field.ExecTypeField,
	ordstatus *field.OrdStatusField,
	symbol *field.SymbolField,
	side *field.SideField,
	leavesqty *field.LeavesQtyField,
	cumqty *field.CumQtyField,
	avgpx *field.AvgPxField) ExecutionReportBuilder {
	var builder ExecutionReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX42))
	builder.Header.Set(field.NewMsgType("8"))
	builder.Body.Set(orderid)
	builder.Body.Set(execid)
	builder.Body.Set(exectranstype)
	builder.Body.Set(exectype)
	builder.Body.Set(ordstatus)
	builder.Body.Set(symbol)
	builder.Body.Set(side)
	builder.Body.Set(leavesqty)
	builder.Body.Set(cumqty)
	builder.Body.Set(avgpx)
	return builder
}

//OrderID is a required field for ExecutionReport.
func (m ExecutionReport) OrderID() (*field.OrderIDField, errors.MessageRejectError) {
	f := &field.OrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from ExecutionReport.
func (m ExecutionReport) GetOrderID(f *field.OrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryOrderID is a non-required field for ExecutionReport.
func (m ExecutionReport) SecondaryOrderID() (*field.SecondaryOrderIDField, errors.MessageRejectError) {
	f := &field.SecondaryOrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryOrderID reads a SecondaryOrderID from ExecutionReport.
func (m ExecutionReport) GetSecondaryOrderID(f *field.SecondaryOrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for ExecutionReport.
func (m ExecutionReport) ClOrdID() (*field.ClOrdIDField, errors.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from ExecutionReport.
func (m ExecutionReport) GetClOrdID(f *field.ClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigClOrdID is a non-required field for ExecutionReport.
func (m ExecutionReport) OrigClOrdID() (*field.OrigClOrdIDField, errors.MessageRejectError) {
	f := &field.OrigClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigClOrdID reads a OrigClOrdID from ExecutionReport.
func (m ExecutionReport) GetOrigClOrdID(f *field.OrigClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClientID is a non-required field for ExecutionReport.
func (m ExecutionReport) ClientID() (*field.ClientIDField, errors.MessageRejectError) {
	f := &field.ClientIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClientID reads a ClientID from ExecutionReport.
func (m ExecutionReport) GetClientID(f *field.ClientIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecBroker is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecBroker() (*field.ExecBrokerField, errors.MessageRejectError) {
	f := &field.ExecBrokerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecBroker reads a ExecBroker from ExecutionReport.
func (m ExecutionReport) GetExecBroker(f *field.ExecBrokerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoContraBrokers is a non-required field for ExecutionReport.
func (m ExecutionReport) NoContraBrokers() (*field.NoContraBrokersField, errors.MessageRejectError) {
	f := &field.NoContraBrokersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoContraBrokers reads a NoContraBrokers from ExecutionReport.
func (m ExecutionReport) GetNoContraBrokers(f *field.NoContraBrokersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListID is a non-required field for ExecutionReport.
func (m ExecutionReport) ListID() (*field.ListIDField, errors.MessageRejectError) {
	f := &field.ListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from ExecutionReport.
func (m ExecutionReport) GetListID(f *field.ListIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecID is a required field for ExecutionReport.
func (m ExecutionReport) ExecID() (*field.ExecIDField, errors.MessageRejectError) {
	f := &field.ExecIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecID reads a ExecID from ExecutionReport.
func (m ExecutionReport) GetExecID(f *field.ExecIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecTransType is a required field for ExecutionReport.
func (m ExecutionReport) ExecTransType() (*field.ExecTransTypeField, errors.MessageRejectError) {
	f := &field.ExecTransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecTransType reads a ExecTransType from ExecutionReport.
func (m ExecutionReport) GetExecTransType(f *field.ExecTransTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecRefID is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecRefID() (*field.ExecRefIDField, errors.MessageRejectError) {
	f := &field.ExecRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecRefID reads a ExecRefID from ExecutionReport.
func (m ExecutionReport) GetExecRefID(f *field.ExecRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecType is a required field for ExecutionReport.
func (m ExecutionReport) ExecType() (*field.ExecTypeField, errors.MessageRejectError) {
	f := &field.ExecTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecType reads a ExecType from ExecutionReport.
func (m ExecutionReport) GetExecType(f *field.ExecTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdStatus is a required field for ExecutionReport.
func (m ExecutionReport) OrdStatus() (*field.OrdStatusField, errors.MessageRejectError) {
	f := &field.OrdStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdStatus reads a OrdStatus from ExecutionReport.
func (m ExecutionReport) GetOrdStatus(f *field.OrdStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdRejReason is a non-required field for ExecutionReport.
func (m ExecutionReport) OrdRejReason() (*field.OrdRejReasonField, errors.MessageRejectError) {
	f := &field.OrdRejReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdRejReason reads a OrdRejReason from ExecutionReport.
func (m ExecutionReport) GetOrdRejReason(f *field.OrdRejReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecRestatementReason is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecRestatementReason() (*field.ExecRestatementReasonField, errors.MessageRejectError) {
	f := &field.ExecRestatementReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecRestatementReason reads a ExecRestatementReason from ExecutionReport.
func (m ExecutionReport) GetExecRestatementReason(f *field.ExecRestatementReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for ExecutionReport.
func (m ExecutionReport) Account() (*field.AccountField, errors.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from ExecutionReport.
func (m ExecutionReport) GetAccount(f *field.AccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlmntTyp is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlmntTyp() (*field.SettlmntTypField, errors.MessageRejectError) {
	f := &field.SettlmntTypField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlmntTyp reads a SettlmntTyp from ExecutionReport.
func (m ExecutionReport) GetSettlmntTyp(f *field.SettlmntTypField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate is a non-required field for ExecutionReport.
func (m ExecutionReport) FutSettDate() (*field.FutSettDateField, errors.MessageRejectError) {
	f := &field.FutSettDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate reads a FutSettDate from ExecutionReport.
func (m ExecutionReport) GetFutSettDate(f *field.FutSettDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a required field for ExecutionReport.
func (m ExecutionReport) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from ExecutionReport.
func (m ExecutionReport) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for ExecutionReport.
func (m ExecutionReport) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from ExecutionReport.
func (m ExecutionReport) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from ExecutionReport.
func (m ExecutionReport) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IDSource is a non-required field for ExecutionReport.
func (m ExecutionReport) IDSource() (*field.IDSourceField, errors.MessageRejectError) {
	f := &field.IDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIDSource reads a IDSource from ExecutionReport.
func (m ExecutionReport) GetIDSource(f *field.IDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from ExecutionReport.
func (m ExecutionReport) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for ExecutionReport.
func (m ExecutionReport) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from ExecutionReport.
func (m ExecutionReport) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDay is a non-required field for ExecutionReport.
func (m ExecutionReport) MaturityDay() (*field.MaturityDayField, errors.MessageRejectError) {
	f := &field.MaturityDayField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDay reads a MaturityDay from ExecutionReport.
func (m ExecutionReport) GetMaturityDay(f *field.MaturityDayField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for ExecutionReport.
func (m ExecutionReport) PutOrCall() (*field.PutOrCallField, errors.MessageRejectError) {
	f := &field.PutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from ExecutionReport.
func (m ExecutionReport) GetPutOrCall(f *field.PutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for ExecutionReport.
func (m ExecutionReport) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from ExecutionReport.
func (m ExecutionReport) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for ExecutionReport.
func (m ExecutionReport) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from ExecutionReport.
func (m ExecutionReport) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for ExecutionReport.
func (m ExecutionReport) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from ExecutionReport.
func (m ExecutionReport) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for ExecutionReport.
func (m ExecutionReport) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from ExecutionReport.
func (m ExecutionReport) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from ExecutionReport.
func (m ExecutionReport) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for ExecutionReport.
func (m ExecutionReport) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from ExecutionReport.
func (m ExecutionReport) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from ExecutionReport.
func (m ExecutionReport) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from ExecutionReport.
func (m ExecutionReport) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from ExecutionReport.
func (m ExecutionReport) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from ExecutionReport.
func (m ExecutionReport) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from ExecutionReport.
func (m ExecutionReport) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for ExecutionReport.
func (m ExecutionReport) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from ExecutionReport.
func (m ExecutionReport) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderQty() (*field.OrderQtyField, errors.MessageRejectError) {
	f := &field.OrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from ExecutionReport.
func (m ExecutionReport) GetOrderQty(f *field.OrderQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashOrderQty is a non-required field for ExecutionReport.
func (m ExecutionReport) CashOrderQty() (*field.CashOrderQtyField, errors.MessageRejectError) {
	f := &field.CashOrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashOrderQty reads a CashOrderQty from ExecutionReport.
func (m ExecutionReport) GetCashOrderQty(f *field.CashOrderQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a non-required field for ExecutionReport.
func (m ExecutionReport) OrdType() (*field.OrdTypeField, errors.MessageRejectError) {
	f := &field.OrdTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from ExecutionReport.
func (m ExecutionReport) GetOrdType(f *field.OrdTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for ExecutionReport.
func (m ExecutionReport) Price() (*field.PriceField, errors.MessageRejectError) {
	f := &field.PriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from ExecutionReport.
func (m ExecutionReport) GetPrice(f *field.PriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StopPx is a non-required field for ExecutionReport.
func (m ExecutionReport) StopPx() (*field.StopPxField, errors.MessageRejectError) {
	f := &field.StopPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStopPx reads a StopPx from ExecutionReport.
func (m ExecutionReport) GetStopPx(f *field.StopPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegDifference is a non-required field for ExecutionReport.
func (m ExecutionReport) PegDifference() (*field.PegDifferenceField, errors.MessageRejectError) {
	f := &field.PegDifferenceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegDifference reads a PegDifference from ExecutionReport.
func (m ExecutionReport) GetPegDifference(f *field.PegDifferenceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionInst is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionInst() (*field.DiscretionInstField, errors.MessageRejectError) {
	f := &field.DiscretionInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionInst reads a DiscretionInst from ExecutionReport.
func (m ExecutionReport) GetDiscretionInst(f *field.DiscretionInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffset is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionOffset() (*field.DiscretionOffsetField, errors.MessageRejectError) {
	f := &field.DiscretionOffsetField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffset reads a DiscretionOffset from ExecutionReport.
func (m ExecutionReport) GetDiscretionOffset(f *field.DiscretionOffsetField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for ExecutionReport.
func (m ExecutionReport) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from ExecutionReport.
func (m ExecutionReport) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ComplianceID is a non-required field for ExecutionReport.
func (m ExecutionReport) ComplianceID() (*field.ComplianceIDField, errors.MessageRejectError) {
	f := &field.ComplianceIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetComplianceID reads a ComplianceID from ExecutionReport.
func (m ExecutionReport) GetComplianceID(f *field.ComplianceIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SolicitedFlag is a non-required field for ExecutionReport.
func (m ExecutionReport) SolicitedFlag() (*field.SolicitedFlagField, errors.MessageRejectError) {
	f := &field.SolicitedFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSolicitedFlag reads a SolicitedFlag from ExecutionReport.
func (m ExecutionReport) GetSolicitedFlag(f *field.SolicitedFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeInForce is a non-required field for ExecutionReport.
func (m ExecutionReport) TimeInForce() (*field.TimeInForceField, errors.MessageRejectError) {
	f := &field.TimeInForceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeInForce reads a TimeInForce from ExecutionReport.
func (m ExecutionReport) GetTimeInForce(f *field.TimeInForceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for ExecutionReport.
func (m ExecutionReport) EffectiveTime() (*field.EffectiveTimeField, errors.MessageRejectError) {
	f := &field.EffectiveTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from ExecutionReport.
func (m ExecutionReport) GetEffectiveTime(f *field.EffectiveTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireDate is a non-required field for ExecutionReport.
func (m ExecutionReport) ExpireDate() (*field.ExpireDateField, errors.MessageRejectError) {
	f := &field.ExpireDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireDate reads a ExpireDate from ExecutionReport.
func (m ExecutionReport) GetExpireDate(f *field.ExpireDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for ExecutionReport.
func (m ExecutionReport) ExpireTime() (*field.ExpireTimeField, errors.MessageRejectError) {
	f := &field.ExpireTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from ExecutionReport.
func (m ExecutionReport) GetExpireTime(f *field.ExpireTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecInst is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecInst() (*field.ExecInstField, errors.MessageRejectError) {
	f := &field.ExecInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecInst reads a ExecInst from ExecutionReport.
func (m ExecutionReport) GetExecInst(f *field.ExecInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Rule80A is a non-required field for ExecutionReport.
func (m ExecutionReport) Rule80A() (*field.Rule80AField, errors.MessageRejectError) {
	f := &field.Rule80AField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRule80A reads a Rule80A from ExecutionReport.
func (m ExecutionReport) GetRule80A(f *field.Rule80AField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastShares is a non-required field for ExecutionReport.
func (m ExecutionReport) LastShares() (*field.LastSharesField, errors.MessageRejectError) {
	f := &field.LastSharesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastShares reads a LastShares from ExecutionReport.
func (m ExecutionReport) GetLastShares(f *field.LastSharesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastPx is a non-required field for ExecutionReport.
func (m ExecutionReport) LastPx() (*field.LastPxField, errors.MessageRejectError) {
	f := &field.LastPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastPx reads a LastPx from ExecutionReport.
func (m ExecutionReport) GetLastPx(f *field.LastPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastSpotRate is a non-required field for ExecutionReport.
func (m ExecutionReport) LastSpotRate() (*field.LastSpotRateField, errors.MessageRejectError) {
	f := &field.LastSpotRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastSpotRate reads a LastSpotRate from ExecutionReport.
func (m ExecutionReport) GetLastSpotRate(f *field.LastSpotRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastForwardPoints is a non-required field for ExecutionReport.
func (m ExecutionReport) LastForwardPoints() (*field.LastForwardPointsField, errors.MessageRejectError) {
	f := &field.LastForwardPointsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastForwardPoints reads a LastForwardPoints from ExecutionReport.
func (m ExecutionReport) GetLastForwardPoints(f *field.LastForwardPointsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastMkt is a non-required field for ExecutionReport.
func (m ExecutionReport) LastMkt() (*field.LastMktField, errors.MessageRejectError) {
	f := &field.LastMktField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastMkt reads a LastMkt from ExecutionReport.
func (m ExecutionReport) GetLastMkt(f *field.LastMktField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for ExecutionReport.
func (m ExecutionReport) TradingSessionID() (*field.TradingSessionIDField, errors.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from ExecutionReport.
func (m ExecutionReport) GetTradingSessionID(f *field.TradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastCapacity is a non-required field for ExecutionReport.
func (m ExecutionReport) LastCapacity() (*field.LastCapacityField, errors.MessageRejectError) {
	f := &field.LastCapacityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastCapacity reads a LastCapacity from ExecutionReport.
func (m ExecutionReport) GetLastCapacity(f *field.LastCapacityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LeavesQty is a required field for ExecutionReport.
func (m ExecutionReport) LeavesQty() (*field.LeavesQtyField, errors.MessageRejectError) {
	f := &field.LeavesQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLeavesQty reads a LeavesQty from ExecutionReport.
func (m ExecutionReport) GetLeavesQty(f *field.LeavesQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CumQty is a required field for ExecutionReport.
func (m ExecutionReport) CumQty() (*field.CumQtyField, errors.MessageRejectError) {
	f := &field.CumQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCumQty reads a CumQty from ExecutionReport.
func (m ExecutionReport) GetCumQty(f *field.CumQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPx is a required field for ExecutionReport.
func (m ExecutionReport) AvgPx() (*field.AvgPxField, errors.MessageRejectError) {
	f := &field.AvgPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPx reads a AvgPx from ExecutionReport.
func (m ExecutionReport) GetAvgPx(f *field.AvgPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DayOrderQty is a non-required field for ExecutionReport.
func (m ExecutionReport) DayOrderQty() (*field.DayOrderQtyField, errors.MessageRejectError) {
	f := &field.DayOrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDayOrderQty reads a DayOrderQty from ExecutionReport.
func (m ExecutionReport) GetDayOrderQty(f *field.DayOrderQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DayCumQty is a non-required field for ExecutionReport.
func (m ExecutionReport) DayCumQty() (*field.DayCumQtyField, errors.MessageRejectError) {
	f := &field.DayCumQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDayCumQty reads a DayCumQty from ExecutionReport.
func (m ExecutionReport) GetDayCumQty(f *field.DayCumQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DayAvgPx is a non-required field for ExecutionReport.
func (m ExecutionReport) DayAvgPx() (*field.DayAvgPxField, errors.MessageRejectError) {
	f := &field.DayAvgPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDayAvgPx reads a DayAvgPx from ExecutionReport.
func (m ExecutionReport) GetDayAvgPx(f *field.DayAvgPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GTBookingInst is a non-required field for ExecutionReport.
func (m ExecutionReport) GTBookingInst() (*field.GTBookingInstField, errors.MessageRejectError) {
	f := &field.GTBookingInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGTBookingInst reads a GTBookingInst from ExecutionReport.
func (m ExecutionReport) GetGTBookingInst(f *field.GTBookingInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for ExecutionReport.
func (m ExecutionReport) TradeDate() (*field.TradeDateField, errors.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from ExecutionReport.
func (m ExecutionReport) GetTradeDate(f *field.TradeDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for ExecutionReport.
func (m ExecutionReport) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from ExecutionReport.
func (m ExecutionReport) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ReportToExch is a non-required field for ExecutionReport.
func (m ExecutionReport) ReportToExch() (*field.ReportToExchField, errors.MessageRejectError) {
	f := &field.ReportToExchField{}
	err := m.Body.Get(f)
	return f, err
}

//GetReportToExch reads a ReportToExch from ExecutionReport.
func (m ExecutionReport) GetReportToExch(f *field.ReportToExchField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Commission is a non-required field for ExecutionReport.
func (m ExecutionReport) Commission() (*field.CommissionField, errors.MessageRejectError) {
	f := &field.CommissionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommission reads a Commission from ExecutionReport.
func (m ExecutionReport) GetCommission(f *field.CommissionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommType is a non-required field for ExecutionReport.
func (m ExecutionReport) CommType() (*field.CommTypeField, errors.MessageRejectError) {
	f := &field.CommTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommType reads a CommType from ExecutionReport.
func (m ExecutionReport) GetCommType(f *field.CommTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GrossTradeAmt is a non-required field for ExecutionReport.
func (m ExecutionReport) GrossTradeAmt() (*field.GrossTradeAmtField, errors.MessageRejectError) {
	f := &field.GrossTradeAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGrossTradeAmt reads a GrossTradeAmt from ExecutionReport.
func (m ExecutionReport) GetGrossTradeAmt(f *field.GrossTradeAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrAmt is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrAmt() (*field.SettlCurrAmtField, errors.MessageRejectError) {
	f := &field.SettlCurrAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrAmt reads a SettlCurrAmt from ExecutionReport.
func (m ExecutionReport) GetSettlCurrAmt(f *field.SettlCurrAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrency() (*field.SettlCurrencyField, errors.MessageRejectError) {
	f := &field.SettlCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from ExecutionReport.
func (m ExecutionReport) GetSettlCurrency(f *field.SettlCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrFxRate is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrFxRate() (*field.SettlCurrFxRateField, errors.MessageRejectError) {
	f := &field.SettlCurrFxRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrFxRate reads a SettlCurrFxRate from ExecutionReport.
func (m ExecutionReport) GetSettlCurrFxRate(f *field.SettlCurrFxRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrFxRateCalc is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrFxRateCalc() (*field.SettlCurrFxRateCalcField, errors.MessageRejectError) {
	f := &field.SettlCurrFxRateCalcField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrFxRateCalc reads a SettlCurrFxRateCalc from ExecutionReport.
func (m ExecutionReport) GetSettlCurrFxRateCalc(f *field.SettlCurrFxRateCalcField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HandlInst is a non-required field for ExecutionReport.
func (m ExecutionReport) HandlInst() (*field.HandlInstField, errors.MessageRejectError) {
	f := &field.HandlInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHandlInst reads a HandlInst from ExecutionReport.
func (m ExecutionReport) GetHandlInst(f *field.HandlInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinQty is a non-required field for ExecutionReport.
func (m ExecutionReport) MinQty() (*field.MinQtyField, errors.MessageRejectError) {
	f := &field.MinQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinQty reads a MinQty from ExecutionReport.
func (m ExecutionReport) GetMinQty(f *field.MinQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxFloor is a non-required field for ExecutionReport.
func (m ExecutionReport) MaxFloor() (*field.MaxFloorField, errors.MessageRejectError) {
	f := &field.MaxFloorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxFloor reads a MaxFloor from ExecutionReport.
func (m ExecutionReport) GetMaxFloor(f *field.MaxFloorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OpenClose is a non-required field for ExecutionReport.
func (m ExecutionReport) OpenClose() (*field.OpenCloseField, errors.MessageRejectError) {
	f := &field.OpenCloseField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOpenClose reads a OpenClose from ExecutionReport.
func (m ExecutionReport) GetOpenClose(f *field.OpenCloseField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxShow is a non-required field for ExecutionReport.
func (m ExecutionReport) MaxShow() (*field.MaxShowField, errors.MessageRejectError) {
	f := &field.MaxShowField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxShow reads a MaxShow from ExecutionReport.
func (m ExecutionReport) GetMaxShow(f *field.MaxShowField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ExecutionReport.
func (m ExecutionReport) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ExecutionReport.
func (m ExecutionReport) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from ExecutionReport.
func (m ExecutionReport) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from ExecutionReport.
func (m ExecutionReport) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate2 is a non-required field for ExecutionReport.
func (m ExecutionReport) FutSettDate2() (*field.FutSettDate2Field, errors.MessageRejectError) {
	f := &field.FutSettDate2Field{}
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate2 reads a FutSettDate2 from ExecutionReport.
func (m ExecutionReport) GetFutSettDate2(f *field.FutSettDate2Field) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty2 is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderQty2() (*field.OrderQty2Field, errors.MessageRejectError) {
	f := &field.OrderQty2Field{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty2 reads a OrderQty2 from ExecutionReport.
func (m ExecutionReport) GetOrderQty2(f *field.OrderQty2Field) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingFirm is a non-required field for ExecutionReport.
func (m ExecutionReport) ClearingFirm() (*field.ClearingFirmField, errors.MessageRejectError) {
	f := &field.ClearingFirmField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingFirm reads a ClearingFirm from ExecutionReport.
func (m ExecutionReport) GetClearingFirm(f *field.ClearingFirmField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingAccount is a non-required field for ExecutionReport.
func (m ExecutionReport) ClearingAccount() (*field.ClearingAccountField, errors.MessageRejectError) {
	f := &field.ClearingAccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingAccount reads a ClearingAccount from ExecutionReport.
func (m ExecutionReport) GetClearingAccount(f *field.ClearingAccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultiLegReportingType is a non-required field for ExecutionReport.
func (m ExecutionReport) MultiLegReportingType() (*field.MultiLegReportingTypeField, errors.MessageRejectError) {
	f := &field.MultiLegReportingTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMultiLegReportingType reads a MultiLegReportingType from ExecutionReport.
func (m ExecutionReport) GetMultiLegReportingType(f *field.MultiLegReportingTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}
