package fix43

import (
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

//NewExecutionReportBuilder returns an initialized ExecutionReportBuilder with specified required fields.
func NewExecutionReportBuilder(
	orderid field.OrderID,
	execid field.ExecID,
	exectype field.ExecType,
	ordstatus field.OrdStatus,
	side field.Side,
	leavesqty field.LeavesQty,
	cumqty field.CumQty,
	avgpx field.AvgPx) *ExecutionReportBuilder {
	builder := new(ExecutionReportBuilder)
	builder.Body.Set(orderid)
	builder.Body.Set(execid)
	builder.Body.Set(exectype)
	builder.Body.Set(ordstatus)
	builder.Body.Set(side)
	builder.Body.Set(leavesqty)
	builder.Body.Set(cumqty)
	builder.Body.Set(avgpx)
	return builder
}

//OrderID is a required field for ExecutionReport.
func (m *ExecutionReport) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryOrderID is a non-required field for ExecutionReport.
func (m *ExecutionReport) SecondaryOrderID() (*field.SecondaryOrderID, error) {
	f := new(field.SecondaryOrderID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryClOrdID is a non-required field for ExecutionReport.
func (m *ExecutionReport) SecondaryClOrdID() (*field.SecondaryClOrdID, error) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryExecID is a non-required field for ExecutionReport.
func (m *ExecutionReport) SecondaryExecID() (*field.SecondaryExecID, error) {
	f := new(field.SecondaryExecID)
	err := m.Body.Get(f)
	return f, err
}

//ClOrdID is a non-required field for ExecutionReport.
func (m *ExecutionReport) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//OrigClOrdID is a non-required field for ExecutionReport.
func (m *ExecutionReport) OrigClOrdID() (*field.OrigClOrdID, error) {
	f := new(field.OrigClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//ClOrdLinkID is a non-required field for ExecutionReport.
func (m *ExecutionReport) ClOrdLinkID() (*field.ClOrdLinkID, error) {
	f := new(field.ClOrdLinkID)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyIDs is a non-required field for ExecutionReport.
func (m *ExecutionReport) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//TradeOriginationDate is a non-required field for ExecutionReport.
func (m *ExecutionReport) TradeOriginationDate() (*field.TradeOriginationDate, error) {
	f := new(field.TradeOriginationDate)
	err := m.Body.Get(f)
	return f, err
}

//NoContraBrokers is a non-required field for ExecutionReport.
func (m *ExecutionReport) NoContraBrokers() (*field.NoContraBrokers, error) {
	f := new(field.NoContraBrokers)
	err := m.Body.Get(f)
	return f, err
}

//ListID is a non-required field for ExecutionReport.
func (m *ExecutionReport) ListID() (*field.ListID, error) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//CrossID is a non-required field for ExecutionReport.
func (m *ExecutionReport) CrossID() (*field.CrossID, error) {
	f := new(field.CrossID)
	err := m.Body.Get(f)
	return f, err
}

//OrigCrossID is a non-required field for ExecutionReport.
func (m *ExecutionReport) OrigCrossID() (*field.OrigCrossID, error) {
	f := new(field.OrigCrossID)
	err := m.Body.Get(f)
	return f, err
}

//CrossType is a non-required field for ExecutionReport.
func (m *ExecutionReport) CrossType() (*field.CrossType, error) {
	f := new(field.CrossType)
	err := m.Body.Get(f)
	return f, err
}

//ExecID is a required field for ExecutionReport.
func (m *ExecutionReport) ExecID() (*field.ExecID, error) {
	f := new(field.ExecID)
	err := m.Body.Get(f)
	return f, err
}

//ExecRefID is a non-required field for ExecutionReport.
func (m *ExecutionReport) ExecRefID() (*field.ExecRefID, error) {
	f := new(field.ExecRefID)
	err := m.Body.Get(f)
	return f, err
}

//ExecType is a required field for ExecutionReport.
func (m *ExecutionReport) ExecType() (*field.ExecType, error) {
	f := new(field.ExecType)
	err := m.Body.Get(f)
	return f, err
}

//OrdStatus is a required field for ExecutionReport.
func (m *ExecutionReport) OrdStatus() (*field.OrdStatus, error) {
	f := new(field.OrdStatus)
	err := m.Body.Get(f)
	return f, err
}

//WorkingIndicator is a non-required field for ExecutionReport.
func (m *ExecutionReport) WorkingIndicator() (*field.WorkingIndicator, error) {
	f := new(field.WorkingIndicator)
	err := m.Body.Get(f)
	return f, err
}

//OrdRejReason is a non-required field for ExecutionReport.
func (m *ExecutionReport) OrdRejReason() (*field.OrdRejReason, error) {
	f := new(field.OrdRejReason)
	err := m.Body.Get(f)
	return f, err
}

//ExecRestatementReason is a non-required field for ExecutionReport.
func (m *ExecutionReport) ExecRestatementReason() (*field.ExecRestatementReason, error) {
	f := new(field.ExecRestatementReason)
	err := m.Body.Get(f)
	return f, err
}

//Account is a non-required field for ExecutionReport.
func (m *ExecutionReport) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//AccountType is a non-required field for ExecutionReport.
func (m *ExecutionReport) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//DayBookingInst is a non-required field for ExecutionReport.
func (m *ExecutionReport) DayBookingInst() (*field.DayBookingInst, error) {
	f := new(field.DayBookingInst)
	err := m.Body.Get(f)
	return f, err
}

//BookingUnit is a non-required field for ExecutionReport.
func (m *ExecutionReport) BookingUnit() (*field.BookingUnit, error) {
	f := new(field.BookingUnit)
	err := m.Body.Get(f)
	return f, err
}

//PreallocMethod is a non-required field for ExecutionReport.
func (m *ExecutionReport) PreallocMethod() (*field.PreallocMethod, error) {
	f := new(field.PreallocMethod)
	err := m.Body.Get(f)
	return f, err
}

//SettlmntTyp is a non-required field for ExecutionReport.
func (m *ExecutionReport) SettlmntTyp() (*field.SettlmntTyp, error) {
	f := new(field.SettlmntTyp)
	err := m.Body.Get(f)
	return f, err
}

//FutSettDate is a non-required field for ExecutionReport.
func (m *ExecutionReport) FutSettDate() (*field.FutSettDate, error) {
	f := new(field.FutSettDate)
	err := m.Body.Get(f)
	return f, err
}

//CashMargin is a non-required field for ExecutionReport.
func (m *ExecutionReport) CashMargin() (*field.CashMargin, error) {
	f := new(field.CashMargin)
	err := m.Body.Get(f)
	return f, err
}

//ClearingFeeIndicator is a non-required field for ExecutionReport.
func (m *ExecutionReport) ClearingFeeIndicator() (*field.ClearingFeeIndicator, error) {
	f := new(field.ClearingFeeIndicator)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for ExecutionReport.
func (m *ExecutionReport) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for ExecutionReport.
func (m *ExecutionReport) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for ExecutionReport.
func (m *ExecutionReport) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for ExecutionReport.
func (m *ExecutionReport) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for ExecutionReport.
func (m *ExecutionReport) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for ExecutionReport.
func (m *ExecutionReport) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for ExecutionReport.
func (m *ExecutionReport) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for ExecutionReport.
func (m *ExecutionReport) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for ExecutionReport.
func (m *ExecutionReport) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for ExecutionReport.
func (m *ExecutionReport) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for ExecutionReport.
func (m *ExecutionReport) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for ExecutionReport.
func (m *ExecutionReport) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for ExecutionReport.
func (m *ExecutionReport) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for ExecutionReport.
func (m *ExecutionReport) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for ExecutionReport.
func (m *ExecutionReport) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for ExecutionReport.
func (m *ExecutionReport) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for ExecutionReport.
func (m *ExecutionReport) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for ExecutionReport.
func (m *ExecutionReport) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for ExecutionReport.
func (m *ExecutionReport) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for ExecutionReport.
func (m *ExecutionReport) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for ExecutionReport.
func (m *ExecutionReport) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for ExecutionReport.
func (m *ExecutionReport) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for ExecutionReport.
func (m *ExecutionReport) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for ExecutionReport.
func (m *ExecutionReport) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for ExecutionReport.
func (m *ExecutionReport) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for ExecutionReport.
func (m *ExecutionReport) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for ExecutionReport.
func (m *ExecutionReport) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for ExecutionReport.
func (m *ExecutionReport) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for ExecutionReport.
func (m *ExecutionReport) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for ExecutionReport.
func (m *ExecutionReport) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for ExecutionReport.
func (m *ExecutionReport) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for ExecutionReport.
func (m *ExecutionReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for ExecutionReport.
func (m *ExecutionReport) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Side is a required field for ExecutionReport.
func (m *ExecutionReport) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//NoStipulations is a non-required field for ExecutionReport.
func (m *ExecutionReport) NoStipulations() (*field.NoStipulations, error) {
	f := new(field.NoStipulations)
	err := m.Body.Get(f)
	return f, err
}

//QuantityType is a non-required field for ExecutionReport.
func (m *ExecutionReport) QuantityType() (*field.QuantityType, error) {
	f := new(field.QuantityType)
	err := m.Body.Get(f)
	return f, err
}

//OrderQty is a non-required field for ExecutionReport.
func (m *ExecutionReport) OrderQty() (*field.OrderQty, error) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//CashOrderQty is a non-required field for ExecutionReport.
func (m *ExecutionReport) CashOrderQty() (*field.CashOrderQty, error) {
	f := new(field.CashOrderQty)
	err := m.Body.Get(f)
	return f, err
}

//OrderPercent is a non-required field for ExecutionReport.
func (m *ExecutionReport) OrderPercent() (*field.OrderPercent, error) {
	f := new(field.OrderPercent)
	err := m.Body.Get(f)
	return f, err
}

//RoundingDirection is a non-required field for ExecutionReport.
func (m *ExecutionReport) RoundingDirection() (*field.RoundingDirection, error) {
	f := new(field.RoundingDirection)
	err := m.Body.Get(f)
	return f, err
}

//RoundingModulus is a non-required field for ExecutionReport.
func (m *ExecutionReport) RoundingModulus() (*field.RoundingModulus, error) {
	f := new(field.RoundingModulus)
	err := m.Body.Get(f)
	return f, err
}

//OrdType is a non-required field for ExecutionReport.
func (m *ExecutionReport) OrdType() (*field.OrdType, error) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}

//PriceType is a non-required field for ExecutionReport.
func (m *ExecutionReport) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//Price is a non-required field for ExecutionReport.
func (m *ExecutionReport) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//StopPx is a non-required field for ExecutionReport.
func (m *ExecutionReport) StopPx() (*field.StopPx, error) {
	f := new(field.StopPx)
	err := m.Body.Get(f)
	return f, err
}

//PegDifference is a non-required field for ExecutionReport.
func (m *ExecutionReport) PegDifference() (*field.PegDifference, error) {
	f := new(field.PegDifference)
	err := m.Body.Get(f)
	return f, err
}

//DiscretionInst is a non-required field for ExecutionReport.
func (m *ExecutionReport) DiscretionInst() (*field.DiscretionInst, error) {
	f := new(field.DiscretionInst)
	err := m.Body.Get(f)
	return f, err
}

//DiscretionOffset is a non-required field for ExecutionReport.
func (m *ExecutionReport) DiscretionOffset() (*field.DiscretionOffset, error) {
	f := new(field.DiscretionOffset)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for ExecutionReport.
func (m *ExecutionReport) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//ComplianceID is a non-required field for ExecutionReport.
func (m *ExecutionReport) ComplianceID() (*field.ComplianceID, error) {
	f := new(field.ComplianceID)
	err := m.Body.Get(f)
	return f, err
}

//SolicitedFlag is a non-required field for ExecutionReport.
func (m *ExecutionReport) SolicitedFlag() (*field.SolicitedFlag, error) {
	f := new(field.SolicitedFlag)
	err := m.Body.Get(f)
	return f, err
}

//TimeInForce is a non-required field for ExecutionReport.
func (m *ExecutionReport) TimeInForce() (*field.TimeInForce, error) {
	f := new(field.TimeInForce)
	err := m.Body.Get(f)
	return f, err
}

//EffectiveTime is a non-required field for ExecutionReport.
func (m *ExecutionReport) EffectiveTime() (*field.EffectiveTime, error) {
	f := new(field.EffectiveTime)
	err := m.Body.Get(f)
	return f, err
}

//ExpireDate is a non-required field for ExecutionReport.
func (m *ExecutionReport) ExpireDate() (*field.ExpireDate, error) {
	f := new(field.ExpireDate)
	err := m.Body.Get(f)
	return f, err
}

//ExpireTime is a non-required field for ExecutionReport.
func (m *ExecutionReport) ExpireTime() (*field.ExpireTime, error) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}

//ExecInst is a non-required field for ExecutionReport.
func (m *ExecutionReport) ExecInst() (*field.ExecInst, error) {
	f := new(field.ExecInst)
	err := m.Body.Get(f)
	return f, err
}

//OrderCapacity is a non-required field for ExecutionReport.
func (m *ExecutionReport) OrderCapacity() (*field.OrderCapacity, error) {
	f := new(field.OrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//OrderRestrictions is a non-required field for ExecutionReport.
func (m *ExecutionReport) OrderRestrictions() (*field.OrderRestrictions, error) {
	f := new(field.OrderRestrictions)
	err := m.Body.Get(f)
	return f, err
}

//CustOrderCapacity is a non-required field for ExecutionReport.
func (m *ExecutionReport) CustOrderCapacity() (*field.CustOrderCapacity, error) {
	f := new(field.CustOrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//Rule80A is a non-required field for ExecutionReport.
func (m *ExecutionReport) Rule80A() (*field.Rule80A, error) {
	f := new(field.Rule80A)
	err := m.Body.Get(f)
	return f, err
}

//LastQty is a non-required field for ExecutionReport.
func (m *ExecutionReport) LastQty() (*field.LastQty, error) {
	f := new(field.LastQty)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingLastQty is a non-required field for ExecutionReport.
func (m *ExecutionReport) UnderlyingLastQty() (*field.UnderlyingLastQty, error) {
	f := new(field.UnderlyingLastQty)
	err := m.Body.Get(f)
	return f, err
}

//LastPx is a non-required field for ExecutionReport.
func (m *ExecutionReport) LastPx() (*field.LastPx, error) {
	f := new(field.LastPx)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingLastPx is a non-required field for ExecutionReport.
func (m *ExecutionReport) UnderlyingLastPx() (*field.UnderlyingLastPx, error) {
	f := new(field.UnderlyingLastPx)
	err := m.Body.Get(f)
	return f, err
}

//LastSpotRate is a non-required field for ExecutionReport.
func (m *ExecutionReport) LastSpotRate() (*field.LastSpotRate, error) {
	f := new(field.LastSpotRate)
	err := m.Body.Get(f)
	return f, err
}

//LastForwardPoints is a non-required field for ExecutionReport.
func (m *ExecutionReport) LastForwardPoints() (*field.LastForwardPoints, error) {
	f := new(field.LastForwardPoints)
	err := m.Body.Get(f)
	return f, err
}

//LastMkt is a non-required field for ExecutionReport.
func (m *ExecutionReport) LastMkt() (*field.LastMkt, error) {
	f := new(field.LastMkt)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionID is a non-required field for ExecutionReport.
func (m *ExecutionReport) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionSubID is a non-required field for ExecutionReport.
func (m *ExecutionReport) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//LastCapacity is a non-required field for ExecutionReport.
func (m *ExecutionReport) LastCapacity() (*field.LastCapacity, error) {
	f := new(field.LastCapacity)
	err := m.Body.Get(f)
	return f, err
}

//LeavesQty is a required field for ExecutionReport.
func (m *ExecutionReport) LeavesQty() (*field.LeavesQty, error) {
	f := new(field.LeavesQty)
	err := m.Body.Get(f)
	return f, err
}

//CumQty is a required field for ExecutionReport.
func (m *ExecutionReport) CumQty() (*field.CumQty, error) {
	f := new(field.CumQty)
	err := m.Body.Get(f)
	return f, err
}

//AvgPx is a required field for ExecutionReport.
func (m *ExecutionReport) AvgPx() (*field.AvgPx, error) {
	f := new(field.AvgPx)
	err := m.Body.Get(f)
	return f, err
}

//DayOrderQty is a non-required field for ExecutionReport.
func (m *ExecutionReport) DayOrderQty() (*field.DayOrderQty, error) {
	f := new(field.DayOrderQty)
	err := m.Body.Get(f)
	return f, err
}

//DayCumQty is a non-required field for ExecutionReport.
func (m *ExecutionReport) DayCumQty() (*field.DayCumQty, error) {
	f := new(field.DayCumQty)
	err := m.Body.Get(f)
	return f, err
}

//DayAvgPx is a non-required field for ExecutionReport.
func (m *ExecutionReport) DayAvgPx() (*field.DayAvgPx, error) {
	f := new(field.DayAvgPx)
	err := m.Body.Get(f)
	return f, err
}

//GTBookingInst is a non-required field for ExecutionReport.
func (m *ExecutionReport) GTBookingInst() (*field.GTBookingInst, error) {
	f := new(field.GTBookingInst)
	err := m.Body.Get(f)
	return f, err
}

//TradeDate is a non-required field for ExecutionReport.
func (m *ExecutionReport) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a non-required field for ExecutionReport.
func (m *ExecutionReport) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//ReportToExch is a non-required field for ExecutionReport.
func (m *ExecutionReport) ReportToExch() (*field.ReportToExch, error) {
	f := new(field.ReportToExch)
	err := m.Body.Get(f)
	return f, err
}

//Commission is a non-required field for ExecutionReport.
func (m *ExecutionReport) Commission() (*field.Commission, error) {
	f := new(field.Commission)
	err := m.Body.Get(f)
	return f, err
}

//CommType is a non-required field for ExecutionReport.
func (m *ExecutionReport) CommType() (*field.CommType, error) {
	f := new(field.CommType)
	err := m.Body.Get(f)
	return f, err
}

//CommCurrency is a non-required field for ExecutionReport.
func (m *ExecutionReport) CommCurrency() (*field.CommCurrency, error) {
	f := new(field.CommCurrency)
	err := m.Body.Get(f)
	return f, err
}

//FundRenewWaiv is a non-required field for ExecutionReport.
func (m *ExecutionReport) FundRenewWaiv() (*field.FundRenewWaiv, error) {
	f := new(field.FundRenewWaiv)
	err := m.Body.Get(f)
	return f, err
}

//Spread is a non-required field for ExecutionReport.
func (m *ExecutionReport) Spread() (*field.Spread, error) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for ExecutionReport.
func (m *ExecutionReport) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, error) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurveName is a non-required field for ExecutionReport.
func (m *ExecutionReport) BenchmarkCurveName() (*field.BenchmarkCurveName, error) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for ExecutionReport.
func (m *ExecutionReport) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, error) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}

//YieldType is a non-required field for ExecutionReport.
func (m *ExecutionReport) YieldType() (*field.YieldType, error) {
	f := new(field.YieldType)
	err := m.Body.Get(f)
	return f, err
}

//Yield is a non-required field for ExecutionReport.
func (m *ExecutionReport) Yield() (*field.Yield, error) {
	f := new(field.Yield)
	err := m.Body.Get(f)
	return f, err
}

//GrossTradeAmt is a non-required field for ExecutionReport.
func (m *ExecutionReport) GrossTradeAmt() (*field.GrossTradeAmt, error) {
	f := new(field.GrossTradeAmt)
	err := m.Body.Get(f)
	return f, err
}

//NumDaysInterest is a non-required field for ExecutionReport.
func (m *ExecutionReport) NumDaysInterest() (*field.NumDaysInterest, error) {
	f := new(field.NumDaysInterest)
	err := m.Body.Get(f)
	return f, err
}

//ExDate is a non-required field for ExecutionReport.
func (m *ExecutionReport) ExDate() (*field.ExDate, error) {
	f := new(field.ExDate)
	err := m.Body.Get(f)
	return f, err
}

//AccruedInterestRate is a non-required field for ExecutionReport.
func (m *ExecutionReport) AccruedInterestRate() (*field.AccruedInterestRate, error) {
	f := new(field.AccruedInterestRate)
	err := m.Body.Get(f)
	return f, err
}

//AccruedInterestAmt is a non-required field for ExecutionReport.
func (m *ExecutionReport) AccruedInterestAmt() (*field.AccruedInterestAmt, error) {
	f := new(field.AccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//TradedFlatSwitch is a non-required field for ExecutionReport.
func (m *ExecutionReport) TradedFlatSwitch() (*field.TradedFlatSwitch, error) {
	f := new(field.TradedFlatSwitch)
	err := m.Body.Get(f)
	return f, err
}

//BasisFeatureDate is a non-required field for ExecutionReport.
func (m *ExecutionReport) BasisFeatureDate() (*field.BasisFeatureDate, error) {
	f := new(field.BasisFeatureDate)
	err := m.Body.Get(f)
	return f, err
}

//BasisFeaturePrice is a non-required field for ExecutionReport.
func (m *ExecutionReport) BasisFeaturePrice() (*field.BasisFeaturePrice, error) {
	f := new(field.BasisFeaturePrice)
	err := m.Body.Get(f)
	return f, err
}

//Concession is a non-required field for ExecutionReport.
func (m *ExecutionReport) Concession() (*field.Concession, error) {
	f := new(field.Concession)
	err := m.Body.Get(f)
	return f, err
}

//TotalTakedown is a non-required field for ExecutionReport.
func (m *ExecutionReport) TotalTakedown() (*field.TotalTakedown, error) {
	f := new(field.TotalTakedown)
	err := m.Body.Get(f)
	return f, err
}

//NetMoney is a non-required field for ExecutionReport.
func (m *ExecutionReport) NetMoney() (*field.NetMoney, error) {
	f := new(field.NetMoney)
	err := m.Body.Get(f)
	return f, err
}

//SettlCurrAmt is a non-required field for ExecutionReport.
func (m *ExecutionReport) SettlCurrAmt() (*field.SettlCurrAmt, error) {
	f := new(field.SettlCurrAmt)
	err := m.Body.Get(f)
	return f, err
}

//SettlCurrency is a non-required field for ExecutionReport.
func (m *ExecutionReport) SettlCurrency() (*field.SettlCurrency, error) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}

//SettlCurrFxRate is a non-required field for ExecutionReport.
func (m *ExecutionReport) SettlCurrFxRate() (*field.SettlCurrFxRate, error) {
	f := new(field.SettlCurrFxRate)
	err := m.Body.Get(f)
	return f, err
}

//SettlCurrFxRateCalc is a non-required field for ExecutionReport.
func (m *ExecutionReport) SettlCurrFxRateCalc() (*field.SettlCurrFxRateCalc, error) {
	f := new(field.SettlCurrFxRateCalc)
	err := m.Body.Get(f)
	return f, err
}

//HandlInst is a non-required field for ExecutionReport.
func (m *ExecutionReport) HandlInst() (*field.HandlInst, error) {
	f := new(field.HandlInst)
	err := m.Body.Get(f)
	return f, err
}

//MinQty is a non-required field for ExecutionReport.
func (m *ExecutionReport) MinQty() (*field.MinQty, error) {
	f := new(field.MinQty)
	err := m.Body.Get(f)
	return f, err
}

//MaxFloor is a non-required field for ExecutionReport.
func (m *ExecutionReport) MaxFloor() (*field.MaxFloor, error) {
	f := new(field.MaxFloor)
	err := m.Body.Get(f)
	return f, err
}

//PositionEffect is a non-required field for ExecutionReport.
func (m *ExecutionReport) PositionEffect() (*field.PositionEffect, error) {
	f := new(field.PositionEffect)
	err := m.Body.Get(f)
	return f, err
}

//MaxShow is a non-required field for ExecutionReport.
func (m *ExecutionReport) MaxShow() (*field.MaxShow, error) {
	f := new(field.MaxShow)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for ExecutionReport.
func (m *ExecutionReport) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for ExecutionReport.
func (m *ExecutionReport) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for ExecutionReport.
func (m *ExecutionReport) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//FutSettDate2 is a non-required field for ExecutionReport.
func (m *ExecutionReport) FutSettDate2() (*field.FutSettDate2, error) {
	f := new(field.FutSettDate2)
	err := m.Body.Get(f)
	return f, err
}

//OrderQty2 is a non-required field for ExecutionReport.
func (m *ExecutionReport) OrderQty2() (*field.OrderQty2, error) {
	f := new(field.OrderQty2)
	err := m.Body.Get(f)
	return f, err
}

//LastForwardPoints2 is a non-required field for ExecutionReport.
func (m *ExecutionReport) LastForwardPoints2() (*field.LastForwardPoints2, error) {
	f := new(field.LastForwardPoints2)
	err := m.Body.Get(f)
	return f, err
}

//MultiLegReportingType is a non-required field for ExecutionReport.
func (m *ExecutionReport) MultiLegReportingType() (*field.MultiLegReportingType, error) {
	f := new(field.MultiLegReportingType)
	err := m.Body.Get(f)
	return f, err
}

//CancellationRights is a non-required field for ExecutionReport.
func (m *ExecutionReport) CancellationRights() (*field.CancellationRights, error) {
	f := new(field.CancellationRights)
	err := m.Body.Get(f)
	return f, err
}

//MoneyLaunderingStatus is a non-required field for ExecutionReport.
func (m *ExecutionReport) MoneyLaunderingStatus() (*field.MoneyLaunderingStatus, error) {
	f := new(field.MoneyLaunderingStatus)
	err := m.Body.Get(f)
	return f, err
}

//RegistID is a non-required field for ExecutionReport.
func (m *ExecutionReport) RegistID() (*field.RegistID, error) {
	f := new(field.RegistID)
	err := m.Body.Get(f)
	return f, err
}

//Designation is a non-required field for ExecutionReport.
func (m *ExecutionReport) Designation() (*field.Designation, error) {
	f := new(field.Designation)
	err := m.Body.Get(f)
	return f, err
}

//TransBkdTime is a non-required field for ExecutionReport.
func (m *ExecutionReport) TransBkdTime() (*field.TransBkdTime, error) {
	f := new(field.TransBkdTime)
	err := m.Body.Get(f)
	return f, err
}

//ExecValuationPoint is a non-required field for ExecutionReport.
func (m *ExecutionReport) ExecValuationPoint() (*field.ExecValuationPoint, error) {
	f := new(field.ExecValuationPoint)
	err := m.Body.Get(f)
	return f, err
}

//ExecPriceType is a non-required field for ExecutionReport.
func (m *ExecutionReport) ExecPriceType() (*field.ExecPriceType, error) {
	f := new(field.ExecPriceType)
	err := m.Body.Get(f)
	return f, err
}

//ExecPriceAdjustment is a non-required field for ExecutionReport.
func (m *ExecutionReport) ExecPriceAdjustment() (*field.ExecPriceAdjustment, error) {
	f := new(field.ExecPriceAdjustment)
	err := m.Body.Get(f)
	return f, err
}

//PriorityIndicator is a non-required field for ExecutionReport.
func (m *ExecutionReport) PriorityIndicator() (*field.PriorityIndicator, error) {
	f := new(field.PriorityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//PriceImprovement is a non-required field for ExecutionReport.
func (m *ExecutionReport) PriceImprovement() (*field.PriceImprovement, error) {
	f := new(field.PriceImprovement)
	err := m.Body.Get(f)
	return f, err
}

//NoContAmts is a non-required field for ExecutionReport.
func (m *ExecutionReport) NoContAmts() (*field.NoContAmts, error) {
	f := new(field.NoContAmts)
	err := m.Body.Get(f)
	return f, err
}

//NoLegs is a non-required field for ExecutionReport.
func (m *ExecutionReport) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}
