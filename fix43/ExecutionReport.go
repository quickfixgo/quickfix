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

//CreateExecutionReportBuilder returns an initialized ExecutionReportBuilder with specified required fields.
func CreateExecutionReportBuilder(
	orderid field.OrderID,
	execid field.ExecID,
	exectype field.ExecType,
	ordstatus field.OrdStatus,
	side field.Side,
	leavesqty field.LeavesQty,
	cumqty field.CumQty,
	avgpx field.AvgPx) ExecutionReportBuilder {
	var builder ExecutionReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
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
func (m ExecutionReport) OrderID() (field.OrderID, error) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryOrderID is a non-required field for ExecutionReport.
func (m ExecutionReport) SecondaryOrderID() (field.SecondaryOrderID, error) {
	var f field.SecondaryOrderID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryClOrdID is a non-required field for ExecutionReport.
func (m ExecutionReport) SecondaryClOrdID() (field.SecondaryClOrdID, error) {
	var f field.SecondaryClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryExecID is a non-required field for ExecutionReport.
func (m ExecutionReport) SecondaryExecID() (field.SecondaryExecID, error) {
	var f field.SecondaryExecID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a non-required field for ExecutionReport.
func (m ExecutionReport) ClOrdID() (field.ClOrdID, error) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//OrigClOrdID is a non-required field for ExecutionReport.
func (m ExecutionReport) OrigClOrdID() (field.OrigClOrdID, error) {
	var f field.OrigClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdLinkID is a non-required field for ExecutionReport.
func (m ExecutionReport) ClOrdLinkID() (field.ClOrdLinkID, error) {
	var f field.ClOrdLinkID
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for ExecutionReport.
func (m ExecutionReport) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//TradeOriginationDate is a non-required field for ExecutionReport.
func (m ExecutionReport) TradeOriginationDate() (field.TradeOriginationDate, error) {
	var f field.TradeOriginationDate
	err := m.Body.Get(&f)
	return f, err
}

//NoContraBrokers is a non-required field for ExecutionReport.
func (m ExecutionReport) NoContraBrokers() (field.NoContraBrokers, error) {
	var f field.NoContraBrokers
	err := m.Body.Get(&f)
	return f, err
}

//ListID is a non-required field for ExecutionReport.
func (m ExecutionReport) ListID() (field.ListID, error) {
	var f field.ListID
	err := m.Body.Get(&f)
	return f, err
}

//CrossID is a non-required field for ExecutionReport.
func (m ExecutionReport) CrossID() (field.CrossID, error) {
	var f field.CrossID
	err := m.Body.Get(&f)
	return f, err
}

//OrigCrossID is a non-required field for ExecutionReport.
func (m ExecutionReport) OrigCrossID() (field.OrigCrossID, error) {
	var f field.OrigCrossID
	err := m.Body.Get(&f)
	return f, err
}

//CrossType is a non-required field for ExecutionReport.
func (m ExecutionReport) CrossType() (field.CrossType, error) {
	var f field.CrossType
	err := m.Body.Get(&f)
	return f, err
}

//ExecID is a required field for ExecutionReport.
func (m ExecutionReport) ExecID() (field.ExecID, error) {
	var f field.ExecID
	err := m.Body.Get(&f)
	return f, err
}

//ExecRefID is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecRefID() (field.ExecRefID, error) {
	var f field.ExecRefID
	err := m.Body.Get(&f)
	return f, err
}

//ExecType is a required field for ExecutionReport.
func (m ExecutionReport) ExecType() (field.ExecType, error) {
	var f field.ExecType
	err := m.Body.Get(&f)
	return f, err
}

//OrdStatus is a required field for ExecutionReport.
func (m ExecutionReport) OrdStatus() (field.OrdStatus, error) {
	var f field.OrdStatus
	err := m.Body.Get(&f)
	return f, err
}

//WorkingIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) WorkingIndicator() (field.WorkingIndicator, error) {
	var f field.WorkingIndicator
	err := m.Body.Get(&f)
	return f, err
}

//OrdRejReason is a non-required field for ExecutionReport.
func (m ExecutionReport) OrdRejReason() (field.OrdRejReason, error) {
	var f field.OrdRejReason
	err := m.Body.Get(&f)
	return f, err
}

//ExecRestatementReason is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecRestatementReason() (field.ExecRestatementReason, error) {
	var f field.ExecRestatementReason
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for ExecutionReport.
func (m ExecutionReport) Account() (field.Account, error) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a non-required field for ExecutionReport.
func (m ExecutionReport) AccountType() (field.AccountType, error) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//DayBookingInst is a non-required field for ExecutionReport.
func (m ExecutionReport) DayBookingInst() (field.DayBookingInst, error) {
	var f field.DayBookingInst
	err := m.Body.Get(&f)
	return f, err
}

//BookingUnit is a non-required field for ExecutionReport.
func (m ExecutionReport) BookingUnit() (field.BookingUnit, error) {
	var f field.BookingUnit
	err := m.Body.Get(&f)
	return f, err
}

//PreallocMethod is a non-required field for ExecutionReport.
func (m ExecutionReport) PreallocMethod() (field.PreallocMethod, error) {
	var f field.PreallocMethod
	err := m.Body.Get(&f)
	return f, err
}

//SettlmntTyp is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlmntTyp() (field.SettlmntTyp, error) {
	var f field.SettlmntTyp
	err := m.Body.Get(&f)
	return f, err
}

//FutSettDate is a non-required field for ExecutionReport.
func (m ExecutionReport) FutSettDate() (field.FutSettDate, error) {
	var f field.FutSettDate
	err := m.Body.Get(&f)
	return f, err
}

//CashMargin is a non-required field for ExecutionReport.
func (m ExecutionReport) CashMargin() (field.CashMargin, error) {
	var f field.CashMargin
	err := m.Body.Get(&f)
	return f, err
}

//ClearingFeeIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) ClearingFeeIndicator() (field.ClearingFeeIndicator, error) {
	var f field.ClearingFeeIndicator
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for ExecutionReport.
func (m ExecutionReport) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for ExecutionReport.
func (m ExecutionReport) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for ExecutionReport.
func (m ExecutionReport) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for ExecutionReport.
func (m ExecutionReport) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for ExecutionReport.
func (m ExecutionReport) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for ExecutionReport.
func (m ExecutionReport) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for ExecutionReport.
func (m ExecutionReport) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for ExecutionReport.
func (m ExecutionReport) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for ExecutionReport.
func (m ExecutionReport) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for ExecutionReport.
func (m ExecutionReport) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for ExecutionReport.
func (m ExecutionReport) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for ExecutionReport.
func (m ExecutionReport) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for ExecutionReport.
func (m ExecutionReport) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for ExecutionReport.
func (m ExecutionReport) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for ExecutionReport.
func (m ExecutionReport) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for ExecutionReport.
func (m ExecutionReport) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for ExecutionReport.
func (m ExecutionReport) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for ExecutionReport.
func (m ExecutionReport) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for ExecutionReport.
func (m ExecutionReport) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for ExecutionReport.
func (m ExecutionReport) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for ExecutionReport.
func (m ExecutionReport) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for ExecutionReport.
func (m ExecutionReport) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for ExecutionReport.
func (m ExecutionReport) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for ExecutionReport.
func (m ExecutionReport) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Side is a required field for ExecutionReport.
func (m ExecutionReport) Side() (field.Side, error) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//NoStipulations is a non-required field for ExecutionReport.
func (m ExecutionReport) NoStipulations() (field.NoStipulations, error) {
	var f field.NoStipulations
	err := m.Body.Get(&f)
	return f, err
}

//QuantityType is a non-required field for ExecutionReport.
func (m ExecutionReport) QuantityType() (field.QuantityType, error) {
	var f field.QuantityType
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderQty() (field.OrderQty, error) {
	var f field.OrderQty
	err := m.Body.Get(&f)
	return f, err
}

//CashOrderQty is a non-required field for ExecutionReport.
func (m ExecutionReport) CashOrderQty() (field.CashOrderQty, error) {
	var f field.CashOrderQty
	err := m.Body.Get(&f)
	return f, err
}

//OrderPercent is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderPercent() (field.OrderPercent, error) {
	var f field.OrderPercent
	err := m.Body.Get(&f)
	return f, err
}

//RoundingDirection is a non-required field for ExecutionReport.
func (m ExecutionReport) RoundingDirection() (field.RoundingDirection, error) {
	var f field.RoundingDirection
	err := m.Body.Get(&f)
	return f, err
}

//RoundingModulus is a non-required field for ExecutionReport.
func (m ExecutionReport) RoundingModulus() (field.RoundingModulus, error) {
	var f field.RoundingModulus
	err := m.Body.Get(&f)
	return f, err
}

//OrdType is a non-required field for ExecutionReport.
func (m ExecutionReport) OrdType() (field.OrdType, error) {
	var f field.OrdType
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for ExecutionReport.
func (m ExecutionReport) PriceType() (field.PriceType, error) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//Price is a non-required field for ExecutionReport.
func (m ExecutionReport) Price() (field.Price, error) {
	var f field.Price
	err := m.Body.Get(&f)
	return f, err
}

//StopPx is a non-required field for ExecutionReport.
func (m ExecutionReport) StopPx() (field.StopPx, error) {
	var f field.StopPx
	err := m.Body.Get(&f)
	return f, err
}

//PegDifference is a non-required field for ExecutionReport.
func (m ExecutionReport) PegDifference() (field.PegDifference, error) {
	var f field.PegDifference
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionInst is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionInst() (field.DiscretionInst, error) {
	var f field.DiscretionInst
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionOffset is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionOffset() (field.DiscretionOffset, error) {
	var f field.DiscretionOffset
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for ExecutionReport.
func (m ExecutionReport) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//ComplianceID is a non-required field for ExecutionReport.
func (m ExecutionReport) ComplianceID() (field.ComplianceID, error) {
	var f field.ComplianceID
	err := m.Body.Get(&f)
	return f, err
}

//SolicitedFlag is a non-required field for ExecutionReport.
func (m ExecutionReport) SolicitedFlag() (field.SolicitedFlag, error) {
	var f field.SolicitedFlag
	err := m.Body.Get(&f)
	return f, err
}

//TimeInForce is a non-required field for ExecutionReport.
func (m ExecutionReport) TimeInForce() (field.TimeInForce, error) {
	var f field.TimeInForce
	err := m.Body.Get(&f)
	return f, err
}

//EffectiveTime is a non-required field for ExecutionReport.
func (m ExecutionReport) EffectiveTime() (field.EffectiveTime, error) {
	var f field.EffectiveTime
	err := m.Body.Get(&f)
	return f, err
}

//ExpireDate is a non-required field for ExecutionReport.
func (m ExecutionReport) ExpireDate() (field.ExpireDate, error) {
	var f field.ExpireDate
	err := m.Body.Get(&f)
	return f, err
}

//ExpireTime is a non-required field for ExecutionReport.
func (m ExecutionReport) ExpireTime() (field.ExpireTime, error) {
	var f field.ExpireTime
	err := m.Body.Get(&f)
	return f, err
}

//ExecInst is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecInst() (field.ExecInst, error) {
	var f field.ExecInst
	err := m.Body.Get(&f)
	return f, err
}

//OrderCapacity is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderCapacity() (field.OrderCapacity, error) {
	var f field.OrderCapacity
	err := m.Body.Get(&f)
	return f, err
}

//OrderRestrictions is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderRestrictions() (field.OrderRestrictions, error) {
	var f field.OrderRestrictions
	err := m.Body.Get(&f)
	return f, err
}

//CustOrderCapacity is a non-required field for ExecutionReport.
func (m ExecutionReport) CustOrderCapacity() (field.CustOrderCapacity, error) {
	var f field.CustOrderCapacity
	err := m.Body.Get(&f)
	return f, err
}

//Rule80A is a non-required field for ExecutionReport.
func (m ExecutionReport) Rule80A() (field.Rule80A, error) {
	var f field.Rule80A
	err := m.Body.Get(&f)
	return f, err
}

//LastQty is a non-required field for ExecutionReport.
func (m ExecutionReport) LastQty() (field.LastQty, error) {
	var f field.LastQty
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingLastQty is a non-required field for ExecutionReport.
func (m ExecutionReport) UnderlyingLastQty() (field.UnderlyingLastQty, error) {
	var f field.UnderlyingLastQty
	err := m.Body.Get(&f)
	return f, err
}

//LastPx is a non-required field for ExecutionReport.
func (m ExecutionReport) LastPx() (field.LastPx, error) {
	var f field.LastPx
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingLastPx is a non-required field for ExecutionReport.
func (m ExecutionReport) UnderlyingLastPx() (field.UnderlyingLastPx, error) {
	var f field.UnderlyingLastPx
	err := m.Body.Get(&f)
	return f, err
}

//LastSpotRate is a non-required field for ExecutionReport.
func (m ExecutionReport) LastSpotRate() (field.LastSpotRate, error) {
	var f field.LastSpotRate
	err := m.Body.Get(&f)
	return f, err
}

//LastForwardPoints is a non-required field for ExecutionReport.
func (m ExecutionReport) LastForwardPoints() (field.LastForwardPoints, error) {
	var f field.LastForwardPoints
	err := m.Body.Get(&f)
	return f, err
}

//LastMkt is a non-required field for ExecutionReport.
func (m ExecutionReport) LastMkt() (field.LastMkt, error) {
	var f field.LastMkt
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for ExecutionReport.
func (m ExecutionReport) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for ExecutionReport.
func (m ExecutionReport) TradingSessionSubID() (field.TradingSessionSubID, error) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//LastCapacity is a non-required field for ExecutionReport.
func (m ExecutionReport) LastCapacity() (field.LastCapacity, error) {
	var f field.LastCapacity
	err := m.Body.Get(&f)
	return f, err
}

//LeavesQty is a required field for ExecutionReport.
func (m ExecutionReport) LeavesQty() (field.LeavesQty, error) {
	var f field.LeavesQty
	err := m.Body.Get(&f)
	return f, err
}

//CumQty is a required field for ExecutionReport.
func (m ExecutionReport) CumQty() (field.CumQty, error) {
	var f field.CumQty
	err := m.Body.Get(&f)
	return f, err
}

//AvgPx is a required field for ExecutionReport.
func (m ExecutionReport) AvgPx() (field.AvgPx, error) {
	var f field.AvgPx
	err := m.Body.Get(&f)
	return f, err
}

//DayOrderQty is a non-required field for ExecutionReport.
func (m ExecutionReport) DayOrderQty() (field.DayOrderQty, error) {
	var f field.DayOrderQty
	err := m.Body.Get(&f)
	return f, err
}

//DayCumQty is a non-required field for ExecutionReport.
func (m ExecutionReport) DayCumQty() (field.DayCumQty, error) {
	var f field.DayCumQty
	err := m.Body.Get(&f)
	return f, err
}

//DayAvgPx is a non-required field for ExecutionReport.
func (m ExecutionReport) DayAvgPx() (field.DayAvgPx, error) {
	var f field.DayAvgPx
	err := m.Body.Get(&f)
	return f, err
}

//GTBookingInst is a non-required field for ExecutionReport.
func (m ExecutionReport) GTBookingInst() (field.GTBookingInst, error) {
	var f field.GTBookingInst
	err := m.Body.Get(&f)
	return f, err
}

//TradeDate is a non-required field for ExecutionReport.
func (m ExecutionReport) TradeDate() (field.TradeDate, error) {
	var f field.TradeDate
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for ExecutionReport.
func (m ExecutionReport) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//ReportToExch is a non-required field for ExecutionReport.
func (m ExecutionReport) ReportToExch() (field.ReportToExch, error) {
	var f field.ReportToExch
	err := m.Body.Get(&f)
	return f, err
}

//Commission is a non-required field for ExecutionReport.
func (m ExecutionReport) Commission() (field.Commission, error) {
	var f field.Commission
	err := m.Body.Get(&f)
	return f, err
}

//CommType is a non-required field for ExecutionReport.
func (m ExecutionReport) CommType() (field.CommType, error) {
	var f field.CommType
	err := m.Body.Get(&f)
	return f, err
}

//CommCurrency is a non-required field for ExecutionReport.
func (m ExecutionReport) CommCurrency() (field.CommCurrency, error) {
	var f field.CommCurrency
	err := m.Body.Get(&f)
	return f, err
}

//FundRenewWaiv is a non-required field for ExecutionReport.
func (m ExecutionReport) FundRenewWaiv() (field.FundRenewWaiv, error) {
	var f field.FundRenewWaiv
	err := m.Body.Get(&f)
	return f, err
}

//Spread is a non-required field for ExecutionReport.
func (m ExecutionReport) Spread() (field.Spread, error) {
	var f field.Spread
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for ExecutionReport.
func (m ExecutionReport) BenchmarkCurveCurrency() (field.BenchmarkCurveCurrency, error) {
	var f field.BenchmarkCurveCurrency
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveName is a non-required field for ExecutionReport.
func (m ExecutionReport) BenchmarkCurveName() (field.BenchmarkCurveName, error) {
	var f field.BenchmarkCurveName
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for ExecutionReport.
func (m ExecutionReport) BenchmarkCurvePoint() (field.BenchmarkCurvePoint, error) {
	var f field.BenchmarkCurvePoint
	err := m.Body.Get(&f)
	return f, err
}

//YieldType is a non-required field for ExecutionReport.
func (m ExecutionReport) YieldType() (field.YieldType, error) {
	var f field.YieldType
	err := m.Body.Get(&f)
	return f, err
}

//Yield is a non-required field for ExecutionReport.
func (m ExecutionReport) Yield() (field.Yield, error) {
	var f field.Yield
	err := m.Body.Get(&f)
	return f, err
}

//GrossTradeAmt is a non-required field for ExecutionReport.
func (m ExecutionReport) GrossTradeAmt() (field.GrossTradeAmt, error) {
	var f field.GrossTradeAmt
	err := m.Body.Get(&f)
	return f, err
}

//NumDaysInterest is a non-required field for ExecutionReport.
func (m ExecutionReport) NumDaysInterest() (field.NumDaysInterest, error) {
	var f field.NumDaysInterest
	err := m.Body.Get(&f)
	return f, err
}

//ExDate is a non-required field for ExecutionReport.
func (m ExecutionReport) ExDate() (field.ExDate, error) {
	var f field.ExDate
	err := m.Body.Get(&f)
	return f, err
}

//AccruedInterestRate is a non-required field for ExecutionReport.
func (m ExecutionReport) AccruedInterestRate() (field.AccruedInterestRate, error) {
	var f field.AccruedInterestRate
	err := m.Body.Get(&f)
	return f, err
}

//AccruedInterestAmt is a non-required field for ExecutionReport.
func (m ExecutionReport) AccruedInterestAmt() (field.AccruedInterestAmt, error) {
	var f field.AccruedInterestAmt
	err := m.Body.Get(&f)
	return f, err
}

//TradedFlatSwitch is a non-required field for ExecutionReport.
func (m ExecutionReport) TradedFlatSwitch() (field.TradedFlatSwitch, error) {
	var f field.TradedFlatSwitch
	err := m.Body.Get(&f)
	return f, err
}

//BasisFeatureDate is a non-required field for ExecutionReport.
func (m ExecutionReport) BasisFeatureDate() (field.BasisFeatureDate, error) {
	var f field.BasisFeatureDate
	err := m.Body.Get(&f)
	return f, err
}

//BasisFeaturePrice is a non-required field for ExecutionReport.
func (m ExecutionReport) BasisFeaturePrice() (field.BasisFeaturePrice, error) {
	var f field.BasisFeaturePrice
	err := m.Body.Get(&f)
	return f, err
}

//Concession is a non-required field for ExecutionReport.
func (m ExecutionReport) Concession() (field.Concession, error) {
	var f field.Concession
	err := m.Body.Get(&f)
	return f, err
}

//TotalTakedown is a non-required field for ExecutionReport.
func (m ExecutionReport) TotalTakedown() (field.TotalTakedown, error) {
	var f field.TotalTakedown
	err := m.Body.Get(&f)
	return f, err
}

//NetMoney is a non-required field for ExecutionReport.
func (m ExecutionReport) NetMoney() (field.NetMoney, error) {
	var f field.NetMoney
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrAmt is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrAmt() (field.SettlCurrAmt, error) {
	var f field.SettlCurrAmt
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrency is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrency() (field.SettlCurrency, error) {
	var f field.SettlCurrency
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrFxRate is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrFxRate() (field.SettlCurrFxRate, error) {
	var f field.SettlCurrFxRate
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrFxRateCalc is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrFxRateCalc() (field.SettlCurrFxRateCalc, error) {
	var f field.SettlCurrFxRateCalc
	err := m.Body.Get(&f)
	return f, err
}

//HandlInst is a non-required field for ExecutionReport.
func (m ExecutionReport) HandlInst() (field.HandlInst, error) {
	var f field.HandlInst
	err := m.Body.Get(&f)
	return f, err
}

//MinQty is a non-required field for ExecutionReport.
func (m ExecutionReport) MinQty() (field.MinQty, error) {
	var f field.MinQty
	err := m.Body.Get(&f)
	return f, err
}

//MaxFloor is a non-required field for ExecutionReport.
func (m ExecutionReport) MaxFloor() (field.MaxFloor, error) {
	var f field.MaxFloor
	err := m.Body.Get(&f)
	return f, err
}

//PositionEffect is a non-required field for ExecutionReport.
func (m ExecutionReport) PositionEffect() (field.PositionEffect, error) {
	var f field.PositionEffect
	err := m.Body.Get(&f)
	return f, err
}

//MaxShow is a non-required field for ExecutionReport.
func (m ExecutionReport) MaxShow() (field.MaxShow, error) {
	var f field.MaxShow
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for ExecutionReport.
func (m ExecutionReport) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//FutSettDate2 is a non-required field for ExecutionReport.
func (m ExecutionReport) FutSettDate2() (field.FutSettDate2, error) {
	var f field.FutSettDate2
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty2 is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderQty2() (field.OrderQty2, error) {
	var f field.OrderQty2
	err := m.Body.Get(&f)
	return f, err
}

//LastForwardPoints2 is a non-required field for ExecutionReport.
func (m ExecutionReport) LastForwardPoints2() (field.LastForwardPoints2, error) {
	var f field.LastForwardPoints2
	err := m.Body.Get(&f)
	return f, err
}

//MultiLegReportingType is a non-required field for ExecutionReport.
func (m ExecutionReport) MultiLegReportingType() (field.MultiLegReportingType, error) {
	var f field.MultiLegReportingType
	err := m.Body.Get(&f)
	return f, err
}

//CancellationRights is a non-required field for ExecutionReport.
func (m ExecutionReport) CancellationRights() (field.CancellationRights, error) {
	var f field.CancellationRights
	err := m.Body.Get(&f)
	return f, err
}

//MoneyLaunderingStatus is a non-required field for ExecutionReport.
func (m ExecutionReport) MoneyLaunderingStatus() (field.MoneyLaunderingStatus, error) {
	var f field.MoneyLaunderingStatus
	err := m.Body.Get(&f)
	return f, err
}

//RegistID is a non-required field for ExecutionReport.
func (m ExecutionReport) RegistID() (field.RegistID, error) {
	var f field.RegistID
	err := m.Body.Get(&f)
	return f, err
}

//Designation is a non-required field for ExecutionReport.
func (m ExecutionReport) Designation() (field.Designation, error) {
	var f field.Designation
	err := m.Body.Get(&f)
	return f, err
}

//TransBkdTime is a non-required field for ExecutionReport.
func (m ExecutionReport) TransBkdTime() (field.TransBkdTime, error) {
	var f field.TransBkdTime
	err := m.Body.Get(&f)
	return f, err
}

//ExecValuationPoint is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecValuationPoint() (field.ExecValuationPoint, error) {
	var f field.ExecValuationPoint
	err := m.Body.Get(&f)
	return f, err
}

//ExecPriceType is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecPriceType() (field.ExecPriceType, error) {
	var f field.ExecPriceType
	err := m.Body.Get(&f)
	return f, err
}

//ExecPriceAdjustment is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecPriceAdjustment() (field.ExecPriceAdjustment, error) {
	var f field.ExecPriceAdjustment
	err := m.Body.Get(&f)
	return f, err
}

//PriorityIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) PriorityIndicator() (field.PriorityIndicator, error) {
	var f field.PriorityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//PriceImprovement is a non-required field for ExecutionReport.
func (m ExecutionReport) PriceImprovement() (field.PriceImprovement, error) {
	var f field.PriceImprovement
	err := m.Body.Get(&f)
	return f, err
}

//NoContAmts is a non-required field for ExecutionReport.
func (m ExecutionReport) NoContAmts() (field.NoContAmts, error) {
	var f field.NoContAmts
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for ExecutionReport.
func (m ExecutionReport) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}
