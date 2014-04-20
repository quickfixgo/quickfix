package fix42

import (
	"github.com/quickfixgo/quickfix/errors"
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
	orderid field.OrderID,
	execid field.ExecID,
	exectranstype field.ExecTransType,
	exectype field.ExecType,
	ordstatus field.OrdStatus,
	symbol field.Symbol,
	side field.Side,
	leavesqty field.LeavesQty,
	cumqty field.CumQty,
	avgpx field.AvgPx) ExecutionReportBuilder {
	var builder ExecutionReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
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
func (m ExecutionReport) OrderID() (*field.OrderID, errors.MessageRejectError) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from ExecutionReport.
func (m ExecutionReport) GetOrderID(f *field.OrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryOrderID is a non-required field for ExecutionReport.
func (m ExecutionReport) SecondaryOrderID() (*field.SecondaryOrderID, errors.MessageRejectError) {
	f := new(field.SecondaryOrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryOrderID reads a SecondaryOrderID from ExecutionReport.
func (m ExecutionReport) GetSecondaryOrderID(f *field.SecondaryOrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for ExecutionReport.
func (m ExecutionReport) ClOrdID() (*field.ClOrdID, errors.MessageRejectError) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from ExecutionReport.
func (m ExecutionReport) GetClOrdID(f *field.ClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigClOrdID is a non-required field for ExecutionReport.
func (m ExecutionReport) OrigClOrdID() (*field.OrigClOrdID, errors.MessageRejectError) {
	f := new(field.OrigClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrigClOrdID reads a OrigClOrdID from ExecutionReport.
func (m ExecutionReport) GetOrigClOrdID(f *field.OrigClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClientID is a non-required field for ExecutionReport.
func (m ExecutionReport) ClientID() (*field.ClientID, errors.MessageRejectError) {
	f := new(field.ClientID)
	err := m.Body.Get(f)
	return f, err
}

//GetClientID reads a ClientID from ExecutionReport.
func (m ExecutionReport) GetClientID(f *field.ClientID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecBroker is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecBroker() (*field.ExecBroker, errors.MessageRejectError) {
	f := new(field.ExecBroker)
	err := m.Body.Get(f)
	return f, err
}

//GetExecBroker reads a ExecBroker from ExecutionReport.
func (m ExecutionReport) GetExecBroker(f *field.ExecBroker) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoContraBrokers is a non-required field for ExecutionReport.
func (m ExecutionReport) NoContraBrokers() (*field.NoContraBrokers, errors.MessageRejectError) {
	f := new(field.NoContraBrokers)
	err := m.Body.Get(f)
	return f, err
}

//GetNoContraBrokers reads a NoContraBrokers from ExecutionReport.
func (m ExecutionReport) GetNoContraBrokers(f *field.NoContraBrokers) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListID is a non-required field for ExecutionReport.
func (m ExecutionReport) ListID() (*field.ListID, errors.MessageRejectError) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from ExecutionReport.
func (m ExecutionReport) GetListID(f *field.ListID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecID is a required field for ExecutionReport.
func (m ExecutionReport) ExecID() (*field.ExecID, errors.MessageRejectError) {
	f := new(field.ExecID)
	err := m.Body.Get(f)
	return f, err
}

//GetExecID reads a ExecID from ExecutionReport.
func (m ExecutionReport) GetExecID(f *field.ExecID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecTransType is a required field for ExecutionReport.
func (m ExecutionReport) ExecTransType() (*field.ExecTransType, errors.MessageRejectError) {
	f := new(field.ExecTransType)
	err := m.Body.Get(f)
	return f, err
}

//GetExecTransType reads a ExecTransType from ExecutionReport.
func (m ExecutionReport) GetExecTransType(f *field.ExecTransType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecRefID is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecRefID() (*field.ExecRefID, errors.MessageRejectError) {
	f := new(field.ExecRefID)
	err := m.Body.Get(f)
	return f, err
}

//GetExecRefID reads a ExecRefID from ExecutionReport.
func (m ExecutionReport) GetExecRefID(f *field.ExecRefID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecType is a required field for ExecutionReport.
func (m ExecutionReport) ExecType() (*field.ExecType, errors.MessageRejectError) {
	f := new(field.ExecType)
	err := m.Body.Get(f)
	return f, err
}

//GetExecType reads a ExecType from ExecutionReport.
func (m ExecutionReport) GetExecType(f *field.ExecType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdStatus is a required field for ExecutionReport.
func (m ExecutionReport) OrdStatus() (*field.OrdStatus, errors.MessageRejectError) {
	f := new(field.OrdStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetOrdStatus reads a OrdStatus from ExecutionReport.
func (m ExecutionReport) GetOrdStatus(f *field.OrdStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdRejReason is a non-required field for ExecutionReport.
func (m ExecutionReport) OrdRejReason() (*field.OrdRejReason, errors.MessageRejectError) {
	f := new(field.OrdRejReason)
	err := m.Body.Get(f)
	return f, err
}

//GetOrdRejReason reads a OrdRejReason from ExecutionReport.
func (m ExecutionReport) GetOrdRejReason(f *field.OrdRejReason) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecRestatementReason is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecRestatementReason() (*field.ExecRestatementReason, errors.MessageRejectError) {
	f := new(field.ExecRestatementReason)
	err := m.Body.Get(f)
	return f, err
}

//GetExecRestatementReason reads a ExecRestatementReason from ExecutionReport.
func (m ExecutionReport) GetExecRestatementReason(f *field.ExecRestatementReason) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for ExecutionReport.
func (m ExecutionReport) Account() (*field.Account, errors.MessageRejectError) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from ExecutionReport.
func (m ExecutionReport) GetAccount(f *field.Account) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlmntTyp is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlmntTyp() (*field.SettlmntTyp, errors.MessageRejectError) {
	f := new(field.SettlmntTyp)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlmntTyp reads a SettlmntTyp from ExecutionReport.
func (m ExecutionReport) GetSettlmntTyp(f *field.SettlmntTyp) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate is a non-required field for ExecutionReport.
func (m ExecutionReport) FutSettDate() (*field.FutSettDate, errors.MessageRejectError) {
	f := new(field.FutSettDate)
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate reads a FutSettDate from ExecutionReport.
func (m ExecutionReport) GetFutSettDate(f *field.FutSettDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a required field for ExecutionReport.
func (m ExecutionReport) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from ExecutionReport.
func (m ExecutionReport) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for ExecutionReport.
func (m ExecutionReport) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from ExecutionReport.
func (m ExecutionReport) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from ExecutionReport.
func (m ExecutionReport) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IDSource is a non-required field for ExecutionReport.
func (m ExecutionReport) IDSource() (*field.IDSource, errors.MessageRejectError) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetIDSource reads a IDSource from ExecutionReport.
func (m ExecutionReport) GetIDSource(f *field.IDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from ExecutionReport.
func (m ExecutionReport) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for ExecutionReport.
func (m ExecutionReport) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from ExecutionReport.
func (m ExecutionReport) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDay is a non-required field for ExecutionReport.
func (m ExecutionReport) MaturityDay() (*field.MaturityDay, errors.MessageRejectError) {
	f := new(field.MaturityDay)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDay reads a MaturityDay from ExecutionReport.
func (m ExecutionReport) GetMaturityDay(f *field.MaturityDay) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for ExecutionReport.
func (m ExecutionReport) PutOrCall() (*field.PutOrCall, errors.MessageRejectError) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from ExecutionReport.
func (m ExecutionReport) GetPutOrCall(f *field.PutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for ExecutionReport.
func (m ExecutionReport) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from ExecutionReport.
func (m ExecutionReport) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for ExecutionReport.
func (m ExecutionReport) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from ExecutionReport.
func (m ExecutionReport) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for ExecutionReport.
func (m ExecutionReport) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from ExecutionReport.
func (m ExecutionReport) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for ExecutionReport.
func (m ExecutionReport) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from ExecutionReport.
func (m ExecutionReport) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from ExecutionReport.
func (m ExecutionReport) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for ExecutionReport.
func (m ExecutionReport) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from ExecutionReport.
func (m ExecutionReport) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from ExecutionReport.
func (m ExecutionReport) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from ExecutionReport.
func (m ExecutionReport) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from ExecutionReport.
func (m ExecutionReport) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from ExecutionReport.
func (m ExecutionReport) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from ExecutionReport.
func (m ExecutionReport) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for ExecutionReport.
func (m ExecutionReport) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from ExecutionReport.
func (m ExecutionReport) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderQty() (*field.OrderQty, errors.MessageRejectError) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from ExecutionReport.
func (m ExecutionReport) GetOrderQty(f *field.OrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashOrderQty is a non-required field for ExecutionReport.
func (m ExecutionReport) CashOrderQty() (*field.CashOrderQty, errors.MessageRejectError) {
	f := new(field.CashOrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetCashOrderQty reads a CashOrderQty from ExecutionReport.
func (m ExecutionReport) GetCashOrderQty(f *field.CashOrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a non-required field for ExecutionReport.
func (m ExecutionReport) OrdType() (*field.OrdType, errors.MessageRejectError) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from ExecutionReport.
func (m ExecutionReport) GetOrdType(f *field.OrdType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for ExecutionReport.
func (m ExecutionReport) Price() (*field.Price, errors.MessageRejectError) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from ExecutionReport.
func (m ExecutionReport) GetPrice(f *field.Price) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StopPx is a non-required field for ExecutionReport.
func (m ExecutionReport) StopPx() (*field.StopPx, errors.MessageRejectError) {
	f := new(field.StopPx)
	err := m.Body.Get(f)
	return f, err
}

//GetStopPx reads a StopPx from ExecutionReport.
func (m ExecutionReport) GetStopPx(f *field.StopPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegDifference is a non-required field for ExecutionReport.
func (m ExecutionReport) PegDifference() (*field.PegDifference, errors.MessageRejectError) {
	f := new(field.PegDifference)
	err := m.Body.Get(f)
	return f, err
}

//GetPegDifference reads a PegDifference from ExecutionReport.
func (m ExecutionReport) GetPegDifference(f *field.PegDifference) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionInst is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionInst() (*field.DiscretionInst, errors.MessageRejectError) {
	f := new(field.DiscretionInst)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionInst reads a DiscretionInst from ExecutionReport.
func (m ExecutionReport) GetDiscretionInst(f *field.DiscretionInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffset is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionOffset() (*field.DiscretionOffset, errors.MessageRejectError) {
	f := new(field.DiscretionOffset)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffset reads a DiscretionOffset from ExecutionReport.
func (m ExecutionReport) GetDiscretionOffset(f *field.DiscretionOffset) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for ExecutionReport.
func (m ExecutionReport) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from ExecutionReport.
func (m ExecutionReport) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ComplianceID is a non-required field for ExecutionReport.
func (m ExecutionReport) ComplianceID() (*field.ComplianceID, errors.MessageRejectError) {
	f := new(field.ComplianceID)
	err := m.Body.Get(f)
	return f, err
}

//GetComplianceID reads a ComplianceID from ExecutionReport.
func (m ExecutionReport) GetComplianceID(f *field.ComplianceID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SolicitedFlag is a non-required field for ExecutionReport.
func (m ExecutionReport) SolicitedFlag() (*field.SolicitedFlag, errors.MessageRejectError) {
	f := new(field.SolicitedFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSolicitedFlag reads a SolicitedFlag from ExecutionReport.
func (m ExecutionReport) GetSolicitedFlag(f *field.SolicitedFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeInForce is a non-required field for ExecutionReport.
func (m ExecutionReport) TimeInForce() (*field.TimeInForce, errors.MessageRejectError) {
	f := new(field.TimeInForce)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeInForce reads a TimeInForce from ExecutionReport.
func (m ExecutionReport) GetTimeInForce(f *field.TimeInForce) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for ExecutionReport.
func (m ExecutionReport) EffectiveTime() (*field.EffectiveTime, errors.MessageRejectError) {
	f := new(field.EffectiveTime)
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from ExecutionReport.
func (m ExecutionReport) GetEffectiveTime(f *field.EffectiveTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireDate is a non-required field for ExecutionReport.
func (m ExecutionReport) ExpireDate() (*field.ExpireDate, errors.MessageRejectError) {
	f := new(field.ExpireDate)
	err := m.Body.Get(f)
	return f, err
}

//GetExpireDate reads a ExpireDate from ExecutionReport.
func (m ExecutionReport) GetExpireDate(f *field.ExpireDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for ExecutionReport.
func (m ExecutionReport) ExpireTime() (*field.ExpireTime, errors.MessageRejectError) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from ExecutionReport.
func (m ExecutionReport) GetExpireTime(f *field.ExpireTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecInst is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecInst() (*field.ExecInst, errors.MessageRejectError) {
	f := new(field.ExecInst)
	err := m.Body.Get(f)
	return f, err
}

//GetExecInst reads a ExecInst from ExecutionReport.
func (m ExecutionReport) GetExecInst(f *field.ExecInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Rule80A is a non-required field for ExecutionReport.
func (m ExecutionReport) Rule80A() (*field.Rule80A, errors.MessageRejectError) {
	f := new(field.Rule80A)
	err := m.Body.Get(f)
	return f, err
}

//GetRule80A reads a Rule80A from ExecutionReport.
func (m ExecutionReport) GetRule80A(f *field.Rule80A) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastShares is a non-required field for ExecutionReport.
func (m ExecutionReport) LastShares() (*field.LastShares, errors.MessageRejectError) {
	f := new(field.LastShares)
	err := m.Body.Get(f)
	return f, err
}

//GetLastShares reads a LastShares from ExecutionReport.
func (m ExecutionReport) GetLastShares(f *field.LastShares) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastPx is a non-required field for ExecutionReport.
func (m ExecutionReport) LastPx() (*field.LastPx, errors.MessageRejectError) {
	f := new(field.LastPx)
	err := m.Body.Get(f)
	return f, err
}

//GetLastPx reads a LastPx from ExecutionReport.
func (m ExecutionReport) GetLastPx(f *field.LastPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastSpotRate is a non-required field for ExecutionReport.
func (m ExecutionReport) LastSpotRate() (*field.LastSpotRate, errors.MessageRejectError) {
	f := new(field.LastSpotRate)
	err := m.Body.Get(f)
	return f, err
}

//GetLastSpotRate reads a LastSpotRate from ExecutionReport.
func (m ExecutionReport) GetLastSpotRate(f *field.LastSpotRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastForwardPoints is a non-required field for ExecutionReport.
func (m ExecutionReport) LastForwardPoints() (*field.LastForwardPoints, errors.MessageRejectError) {
	f := new(field.LastForwardPoints)
	err := m.Body.Get(f)
	return f, err
}

//GetLastForwardPoints reads a LastForwardPoints from ExecutionReport.
func (m ExecutionReport) GetLastForwardPoints(f *field.LastForwardPoints) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastMkt is a non-required field for ExecutionReport.
func (m ExecutionReport) LastMkt() (*field.LastMkt, errors.MessageRejectError) {
	f := new(field.LastMkt)
	err := m.Body.Get(f)
	return f, err
}

//GetLastMkt reads a LastMkt from ExecutionReport.
func (m ExecutionReport) GetLastMkt(f *field.LastMkt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for ExecutionReport.
func (m ExecutionReport) TradingSessionID() (*field.TradingSessionID, errors.MessageRejectError) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from ExecutionReport.
func (m ExecutionReport) GetTradingSessionID(f *field.TradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastCapacity is a non-required field for ExecutionReport.
func (m ExecutionReport) LastCapacity() (*field.LastCapacity, errors.MessageRejectError) {
	f := new(field.LastCapacity)
	err := m.Body.Get(f)
	return f, err
}

//GetLastCapacity reads a LastCapacity from ExecutionReport.
func (m ExecutionReport) GetLastCapacity(f *field.LastCapacity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LeavesQty is a required field for ExecutionReport.
func (m ExecutionReport) LeavesQty() (*field.LeavesQty, errors.MessageRejectError) {
	f := new(field.LeavesQty)
	err := m.Body.Get(f)
	return f, err
}

//GetLeavesQty reads a LeavesQty from ExecutionReport.
func (m ExecutionReport) GetLeavesQty(f *field.LeavesQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CumQty is a required field for ExecutionReport.
func (m ExecutionReport) CumQty() (*field.CumQty, errors.MessageRejectError) {
	f := new(field.CumQty)
	err := m.Body.Get(f)
	return f, err
}

//GetCumQty reads a CumQty from ExecutionReport.
func (m ExecutionReport) GetCumQty(f *field.CumQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPx is a required field for ExecutionReport.
func (m ExecutionReport) AvgPx() (*field.AvgPx, errors.MessageRejectError) {
	f := new(field.AvgPx)
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPx reads a AvgPx from ExecutionReport.
func (m ExecutionReport) GetAvgPx(f *field.AvgPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DayOrderQty is a non-required field for ExecutionReport.
func (m ExecutionReport) DayOrderQty() (*field.DayOrderQty, errors.MessageRejectError) {
	f := new(field.DayOrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetDayOrderQty reads a DayOrderQty from ExecutionReport.
func (m ExecutionReport) GetDayOrderQty(f *field.DayOrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DayCumQty is a non-required field for ExecutionReport.
func (m ExecutionReport) DayCumQty() (*field.DayCumQty, errors.MessageRejectError) {
	f := new(field.DayCumQty)
	err := m.Body.Get(f)
	return f, err
}

//GetDayCumQty reads a DayCumQty from ExecutionReport.
func (m ExecutionReport) GetDayCumQty(f *field.DayCumQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DayAvgPx is a non-required field for ExecutionReport.
func (m ExecutionReport) DayAvgPx() (*field.DayAvgPx, errors.MessageRejectError) {
	f := new(field.DayAvgPx)
	err := m.Body.Get(f)
	return f, err
}

//GetDayAvgPx reads a DayAvgPx from ExecutionReport.
func (m ExecutionReport) GetDayAvgPx(f *field.DayAvgPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GTBookingInst is a non-required field for ExecutionReport.
func (m ExecutionReport) GTBookingInst() (*field.GTBookingInst, errors.MessageRejectError) {
	f := new(field.GTBookingInst)
	err := m.Body.Get(f)
	return f, err
}

//GetGTBookingInst reads a GTBookingInst from ExecutionReport.
func (m ExecutionReport) GetGTBookingInst(f *field.GTBookingInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for ExecutionReport.
func (m ExecutionReport) TradeDate() (*field.TradeDate, errors.MessageRejectError) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from ExecutionReport.
func (m ExecutionReport) GetTradeDate(f *field.TradeDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for ExecutionReport.
func (m ExecutionReport) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from ExecutionReport.
func (m ExecutionReport) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ReportToExch is a non-required field for ExecutionReport.
func (m ExecutionReport) ReportToExch() (*field.ReportToExch, errors.MessageRejectError) {
	f := new(field.ReportToExch)
	err := m.Body.Get(f)
	return f, err
}

//GetReportToExch reads a ReportToExch from ExecutionReport.
func (m ExecutionReport) GetReportToExch(f *field.ReportToExch) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Commission is a non-required field for ExecutionReport.
func (m ExecutionReport) Commission() (*field.Commission, errors.MessageRejectError) {
	f := new(field.Commission)
	err := m.Body.Get(f)
	return f, err
}

//GetCommission reads a Commission from ExecutionReport.
func (m ExecutionReport) GetCommission(f *field.Commission) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommType is a non-required field for ExecutionReport.
func (m ExecutionReport) CommType() (*field.CommType, errors.MessageRejectError) {
	f := new(field.CommType)
	err := m.Body.Get(f)
	return f, err
}

//GetCommType reads a CommType from ExecutionReport.
func (m ExecutionReport) GetCommType(f *field.CommType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GrossTradeAmt is a non-required field for ExecutionReport.
func (m ExecutionReport) GrossTradeAmt() (*field.GrossTradeAmt, errors.MessageRejectError) {
	f := new(field.GrossTradeAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetGrossTradeAmt reads a GrossTradeAmt from ExecutionReport.
func (m ExecutionReport) GetGrossTradeAmt(f *field.GrossTradeAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrAmt is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrAmt() (*field.SettlCurrAmt, errors.MessageRejectError) {
	f := new(field.SettlCurrAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrAmt reads a SettlCurrAmt from ExecutionReport.
func (m ExecutionReport) GetSettlCurrAmt(f *field.SettlCurrAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrency() (*field.SettlCurrency, errors.MessageRejectError) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from ExecutionReport.
func (m ExecutionReport) GetSettlCurrency(f *field.SettlCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrFxRate is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrFxRate() (*field.SettlCurrFxRate, errors.MessageRejectError) {
	f := new(field.SettlCurrFxRate)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrFxRate reads a SettlCurrFxRate from ExecutionReport.
func (m ExecutionReport) GetSettlCurrFxRate(f *field.SettlCurrFxRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrFxRateCalc is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrFxRateCalc() (*field.SettlCurrFxRateCalc, errors.MessageRejectError) {
	f := new(field.SettlCurrFxRateCalc)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrFxRateCalc reads a SettlCurrFxRateCalc from ExecutionReport.
func (m ExecutionReport) GetSettlCurrFxRateCalc(f *field.SettlCurrFxRateCalc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HandlInst is a non-required field for ExecutionReport.
func (m ExecutionReport) HandlInst() (*field.HandlInst, errors.MessageRejectError) {
	f := new(field.HandlInst)
	err := m.Body.Get(f)
	return f, err
}

//GetHandlInst reads a HandlInst from ExecutionReport.
func (m ExecutionReport) GetHandlInst(f *field.HandlInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinQty is a non-required field for ExecutionReport.
func (m ExecutionReport) MinQty() (*field.MinQty, errors.MessageRejectError) {
	f := new(field.MinQty)
	err := m.Body.Get(f)
	return f, err
}

//GetMinQty reads a MinQty from ExecutionReport.
func (m ExecutionReport) GetMinQty(f *field.MinQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxFloor is a non-required field for ExecutionReport.
func (m ExecutionReport) MaxFloor() (*field.MaxFloor, errors.MessageRejectError) {
	f := new(field.MaxFloor)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxFloor reads a MaxFloor from ExecutionReport.
func (m ExecutionReport) GetMaxFloor(f *field.MaxFloor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OpenClose is a non-required field for ExecutionReport.
func (m ExecutionReport) OpenClose() (*field.OpenClose, errors.MessageRejectError) {
	f := new(field.OpenClose)
	err := m.Body.Get(f)
	return f, err
}

//GetOpenClose reads a OpenClose from ExecutionReport.
func (m ExecutionReport) GetOpenClose(f *field.OpenClose) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxShow is a non-required field for ExecutionReport.
func (m ExecutionReport) MaxShow() (*field.MaxShow, errors.MessageRejectError) {
	f := new(field.MaxShow)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxShow reads a MaxShow from ExecutionReport.
func (m ExecutionReport) GetMaxShow(f *field.MaxShow) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ExecutionReport.
func (m ExecutionReport) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ExecutionReport.
func (m ExecutionReport) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from ExecutionReport.
func (m ExecutionReport) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from ExecutionReport.
func (m ExecutionReport) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate2 is a non-required field for ExecutionReport.
func (m ExecutionReport) FutSettDate2() (*field.FutSettDate2, errors.MessageRejectError) {
	f := new(field.FutSettDate2)
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate2 reads a FutSettDate2 from ExecutionReport.
func (m ExecutionReport) GetFutSettDate2(f *field.FutSettDate2) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty2 is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderQty2() (*field.OrderQty2, errors.MessageRejectError) {
	f := new(field.OrderQty2)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty2 reads a OrderQty2 from ExecutionReport.
func (m ExecutionReport) GetOrderQty2(f *field.OrderQty2) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingFirm is a non-required field for ExecutionReport.
func (m ExecutionReport) ClearingFirm() (*field.ClearingFirm, errors.MessageRejectError) {
	f := new(field.ClearingFirm)
	err := m.Body.Get(f)
	return f, err
}

//GetClearingFirm reads a ClearingFirm from ExecutionReport.
func (m ExecutionReport) GetClearingFirm(f *field.ClearingFirm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingAccount is a non-required field for ExecutionReport.
func (m ExecutionReport) ClearingAccount() (*field.ClearingAccount, errors.MessageRejectError) {
	f := new(field.ClearingAccount)
	err := m.Body.Get(f)
	return f, err
}

//GetClearingAccount reads a ClearingAccount from ExecutionReport.
func (m ExecutionReport) GetClearingAccount(f *field.ClearingAccount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultiLegReportingType is a non-required field for ExecutionReport.
func (m ExecutionReport) MultiLegReportingType() (*field.MultiLegReportingType, errors.MessageRejectError) {
	f := new(field.MultiLegReportingType)
	err := m.Body.Get(f)
	return f, err
}

//GetMultiLegReportingType reads a MultiLegReportingType from ExecutionReport.
func (m ExecutionReport) GetMultiLegReportingType(f *field.MultiLegReportingType) errors.MessageRejectError {
	return m.Body.Get(f)
}
