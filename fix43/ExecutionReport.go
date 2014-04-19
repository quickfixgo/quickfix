package fix43

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
func (m ExecutionReport) OrderID() (field.OrderID, errors.MessageRejectError) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryOrderID is a non-required field for ExecutionReport.
func (m ExecutionReport) SecondaryOrderID() (field.SecondaryOrderID, errors.MessageRejectError) {
	var f field.SecondaryOrderID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryClOrdID is a non-required field for ExecutionReport.
func (m ExecutionReport) SecondaryClOrdID() (field.SecondaryClOrdID, errors.MessageRejectError) {
	var f field.SecondaryClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryExecID is a non-required field for ExecutionReport.
func (m ExecutionReport) SecondaryExecID() (field.SecondaryExecID, errors.MessageRejectError) {
	var f field.SecondaryExecID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a non-required field for ExecutionReport.
func (m ExecutionReport) ClOrdID() (field.ClOrdID, errors.MessageRejectError) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//OrigClOrdID is a non-required field for ExecutionReport.
func (m ExecutionReport) OrigClOrdID() (field.OrigClOrdID, errors.MessageRejectError) {
	var f field.OrigClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdLinkID is a non-required field for ExecutionReport.
func (m ExecutionReport) ClOrdLinkID() (field.ClOrdLinkID, errors.MessageRejectError) {
	var f field.ClOrdLinkID
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for ExecutionReport.
func (m ExecutionReport) NoPartyIDs() (field.NoPartyIDs, errors.MessageRejectError) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//TradeOriginationDate is a non-required field for ExecutionReport.
func (m ExecutionReport) TradeOriginationDate() (field.TradeOriginationDate, errors.MessageRejectError) {
	var f field.TradeOriginationDate
	err := m.Body.Get(&f)
	return f, err
}

//NoContraBrokers is a non-required field for ExecutionReport.
func (m ExecutionReport) NoContraBrokers() (field.NoContraBrokers, errors.MessageRejectError) {
	var f field.NoContraBrokers
	err := m.Body.Get(&f)
	return f, err
}

//ListID is a non-required field for ExecutionReport.
func (m ExecutionReport) ListID() (field.ListID, errors.MessageRejectError) {
	var f field.ListID
	err := m.Body.Get(&f)
	return f, err
}

//CrossID is a non-required field for ExecutionReport.
func (m ExecutionReport) CrossID() (field.CrossID, errors.MessageRejectError) {
	var f field.CrossID
	err := m.Body.Get(&f)
	return f, err
}

//OrigCrossID is a non-required field for ExecutionReport.
func (m ExecutionReport) OrigCrossID() (field.OrigCrossID, errors.MessageRejectError) {
	var f field.OrigCrossID
	err := m.Body.Get(&f)
	return f, err
}

//CrossType is a non-required field for ExecutionReport.
func (m ExecutionReport) CrossType() (field.CrossType, errors.MessageRejectError) {
	var f field.CrossType
	err := m.Body.Get(&f)
	return f, err
}

//ExecID is a required field for ExecutionReport.
func (m ExecutionReport) ExecID() (field.ExecID, errors.MessageRejectError) {
	var f field.ExecID
	err := m.Body.Get(&f)
	return f, err
}

//ExecRefID is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecRefID() (field.ExecRefID, errors.MessageRejectError) {
	var f field.ExecRefID
	err := m.Body.Get(&f)
	return f, err
}

//ExecType is a required field for ExecutionReport.
func (m ExecutionReport) ExecType() (field.ExecType, errors.MessageRejectError) {
	var f field.ExecType
	err := m.Body.Get(&f)
	return f, err
}

//OrdStatus is a required field for ExecutionReport.
func (m ExecutionReport) OrdStatus() (field.OrdStatus, errors.MessageRejectError) {
	var f field.OrdStatus
	err := m.Body.Get(&f)
	return f, err
}

//WorkingIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) WorkingIndicator() (field.WorkingIndicator, errors.MessageRejectError) {
	var f field.WorkingIndicator
	err := m.Body.Get(&f)
	return f, err
}

//OrdRejReason is a non-required field for ExecutionReport.
func (m ExecutionReport) OrdRejReason() (field.OrdRejReason, errors.MessageRejectError) {
	var f field.OrdRejReason
	err := m.Body.Get(&f)
	return f, err
}

//ExecRestatementReason is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecRestatementReason() (field.ExecRestatementReason, errors.MessageRejectError) {
	var f field.ExecRestatementReason
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for ExecutionReport.
func (m ExecutionReport) Account() (field.Account, errors.MessageRejectError) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a non-required field for ExecutionReport.
func (m ExecutionReport) AccountType() (field.AccountType, errors.MessageRejectError) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//DayBookingInst is a non-required field for ExecutionReport.
func (m ExecutionReport) DayBookingInst() (field.DayBookingInst, errors.MessageRejectError) {
	var f field.DayBookingInst
	err := m.Body.Get(&f)
	return f, err
}

//BookingUnit is a non-required field for ExecutionReport.
func (m ExecutionReport) BookingUnit() (field.BookingUnit, errors.MessageRejectError) {
	var f field.BookingUnit
	err := m.Body.Get(&f)
	return f, err
}

//PreallocMethod is a non-required field for ExecutionReport.
func (m ExecutionReport) PreallocMethod() (field.PreallocMethod, errors.MessageRejectError) {
	var f field.PreallocMethod
	err := m.Body.Get(&f)
	return f, err
}

//SettlmntTyp is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlmntTyp() (field.SettlmntTyp, errors.MessageRejectError) {
	var f field.SettlmntTyp
	err := m.Body.Get(&f)
	return f, err
}

//FutSettDate is a non-required field for ExecutionReport.
func (m ExecutionReport) FutSettDate() (field.FutSettDate, errors.MessageRejectError) {
	var f field.FutSettDate
	err := m.Body.Get(&f)
	return f, err
}

//CashMargin is a non-required field for ExecutionReport.
func (m ExecutionReport) CashMargin() (field.CashMargin, errors.MessageRejectError) {
	var f field.CashMargin
	err := m.Body.Get(&f)
	return f, err
}

//ClearingFeeIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) ClearingFeeIndicator() (field.ClearingFeeIndicator, errors.MessageRejectError) {
	var f field.ClearingFeeIndicator
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for ExecutionReport.
func (m ExecutionReport) Symbol() (field.Symbol, errors.MessageRejectError) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for ExecutionReport.
func (m ExecutionReport) SymbolSfx() (field.SymbolSfx, errors.MessageRejectError) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityID() (field.SecurityID, errors.MessageRejectError) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityIDSource() (field.SecurityIDSource, errors.MessageRejectError) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for ExecutionReport.
func (m ExecutionReport) NoSecurityAltID() (field.NoSecurityAltID, errors.MessageRejectError) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for ExecutionReport.
func (m ExecutionReport) Product() (field.Product, errors.MessageRejectError) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for ExecutionReport.
func (m ExecutionReport) CFICode() (field.CFICode, errors.MessageRejectError) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityType() (field.SecurityType, errors.MessageRejectError) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for ExecutionReport.
func (m ExecutionReport) MaturityMonthYear() (field.MaturityMonthYear, errors.MessageRejectError) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for ExecutionReport.
func (m ExecutionReport) MaturityDate() (field.MaturityDate, errors.MessageRejectError) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for ExecutionReport.
func (m ExecutionReport) CouponPaymentDate() (field.CouponPaymentDate, errors.MessageRejectError) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for ExecutionReport.
func (m ExecutionReport) IssueDate() (field.IssueDate, errors.MessageRejectError) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for ExecutionReport.
func (m ExecutionReport) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, errors.MessageRejectError) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for ExecutionReport.
func (m ExecutionReport) RepurchaseTerm() (field.RepurchaseTerm, errors.MessageRejectError) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for ExecutionReport.
func (m ExecutionReport) RepurchaseRate() (field.RepurchaseRate, errors.MessageRejectError) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for ExecutionReport.
func (m ExecutionReport) Factor() (field.Factor, errors.MessageRejectError) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for ExecutionReport.
func (m ExecutionReport) CreditRating() (field.CreditRating, errors.MessageRejectError) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for ExecutionReport.
func (m ExecutionReport) InstrRegistry() (field.InstrRegistry, errors.MessageRejectError) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for ExecutionReport.
func (m ExecutionReport) CountryOfIssue() (field.CountryOfIssue, errors.MessageRejectError) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for ExecutionReport.
func (m ExecutionReport) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for ExecutionReport.
func (m ExecutionReport) LocaleOfIssue() (field.LocaleOfIssue, errors.MessageRejectError) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for ExecutionReport.
func (m ExecutionReport) RedemptionDate() (field.RedemptionDate, errors.MessageRejectError) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for ExecutionReport.
func (m ExecutionReport) StrikePrice() (field.StrikePrice, errors.MessageRejectError) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for ExecutionReport.
func (m ExecutionReport) OptAttribute() (field.OptAttribute, errors.MessageRejectError) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for ExecutionReport.
func (m ExecutionReport) ContractMultiplier() (field.ContractMultiplier, errors.MessageRejectError) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for ExecutionReport.
func (m ExecutionReport) CouponRate() (field.CouponRate, errors.MessageRejectError) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityExchange() (field.SecurityExchange, errors.MessageRejectError) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for ExecutionReport.
func (m ExecutionReport) Issuer() (field.Issuer, errors.MessageRejectError) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedIssuerLen() (field.EncodedIssuerLen, errors.MessageRejectError) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedIssuer() (field.EncodedIssuer, errors.MessageRejectError) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityDesc() (field.SecurityDesc, errors.MessageRejectError) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, errors.MessageRejectError) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedSecurityDesc() (field.EncodedSecurityDesc, errors.MessageRejectError) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Side is a required field for ExecutionReport.
func (m ExecutionReport) Side() (field.Side, errors.MessageRejectError) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//NoStipulations is a non-required field for ExecutionReport.
func (m ExecutionReport) NoStipulations() (field.NoStipulations, errors.MessageRejectError) {
	var f field.NoStipulations
	err := m.Body.Get(&f)
	return f, err
}

//QuantityType is a non-required field for ExecutionReport.
func (m ExecutionReport) QuantityType() (field.QuantityType, errors.MessageRejectError) {
	var f field.QuantityType
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderQty() (field.OrderQty, errors.MessageRejectError) {
	var f field.OrderQty
	err := m.Body.Get(&f)
	return f, err
}

//CashOrderQty is a non-required field for ExecutionReport.
func (m ExecutionReport) CashOrderQty() (field.CashOrderQty, errors.MessageRejectError) {
	var f field.CashOrderQty
	err := m.Body.Get(&f)
	return f, err
}

//OrderPercent is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderPercent() (field.OrderPercent, errors.MessageRejectError) {
	var f field.OrderPercent
	err := m.Body.Get(&f)
	return f, err
}

//RoundingDirection is a non-required field for ExecutionReport.
func (m ExecutionReport) RoundingDirection() (field.RoundingDirection, errors.MessageRejectError) {
	var f field.RoundingDirection
	err := m.Body.Get(&f)
	return f, err
}

//RoundingModulus is a non-required field for ExecutionReport.
func (m ExecutionReport) RoundingModulus() (field.RoundingModulus, errors.MessageRejectError) {
	var f field.RoundingModulus
	err := m.Body.Get(&f)
	return f, err
}

//OrdType is a non-required field for ExecutionReport.
func (m ExecutionReport) OrdType() (field.OrdType, errors.MessageRejectError) {
	var f field.OrdType
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for ExecutionReport.
func (m ExecutionReport) PriceType() (field.PriceType, errors.MessageRejectError) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//Price is a non-required field for ExecutionReport.
func (m ExecutionReport) Price() (field.Price, errors.MessageRejectError) {
	var f field.Price
	err := m.Body.Get(&f)
	return f, err
}

//StopPx is a non-required field for ExecutionReport.
func (m ExecutionReport) StopPx() (field.StopPx, errors.MessageRejectError) {
	var f field.StopPx
	err := m.Body.Get(&f)
	return f, err
}

//PegDifference is a non-required field for ExecutionReport.
func (m ExecutionReport) PegDifference() (field.PegDifference, errors.MessageRejectError) {
	var f field.PegDifference
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionInst is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionInst() (field.DiscretionInst, errors.MessageRejectError) {
	var f field.DiscretionInst
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionOffset is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionOffset() (field.DiscretionOffset, errors.MessageRejectError) {
	var f field.DiscretionOffset
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for ExecutionReport.
func (m ExecutionReport) Currency() (field.Currency, errors.MessageRejectError) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//ComplianceID is a non-required field for ExecutionReport.
func (m ExecutionReport) ComplianceID() (field.ComplianceID, errors.MessageRejectError) {
	var f field.ComplianceID
	err := m.Body.Get(&f)
	return f, err
}

//SolicitedFlag is a non-required field for ExecutionReport.
func (m ExecutionReport) SolicitedFlag() (field.SolicitedFlag, errors.MessageRejectError) {
	var f field.SolicitedFlag
	err := m.Body.Get(&f)
	return f, err
}

//TimeInForce is a non-required field for ExecutionReport.
func (m ExecutionReport) TimeInForce() (field.TimeInForce, errors.MessageRejectError) {
	var f field.TimeInForce
	err := m.Body.Get(&f)
	return f, err
}

//EffectiveTime is a non-required field for ExecutionReport.
func (m ExecutionReport) EffectiveTime() (field.EffectiveTime, errors.MessageRejectError) {
	var f field.EffectiveTime
	err := m.Body.Get(&f)
	return f, err
}

//ExpireDate is a non-required field for ExecutionReport.
func (m ExecutionReport) ExpireDate() (field.ExpireDate, errors.MessageRejectError) {
	var f field.ExpireDate
	err := m.Body.Get(&f)
	return f, err
}

//ExpireTime is a non-required field for ExecutionReport.
func (m ExecutionReport) ExpireTime() (field.ExpireTime, errors.MessageRejectError) {
	var f field.ExpireTime
	err := m.Body.Get(&f)
	return f, err
}

//ExecInst is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecInst() (field.ExecInst, errors.MessageRejectError) {
	var f field.ExecInst
	err := m.Body.Get(&f)
	return f, err
}

//OrderCapacity is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderCapacity() (field.OrderCapacity, errors.MessageRejectError) {
	var f field.OrderCapacity
	err := m.Body.Get(&f)
	return f, err
}

//OrderRestrictions is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderRestrictions() (field.OrderRestrictions, errors.MessageRejectError) {
	var f field.OrderRestrictions
	err := m.Body.Get(&f)
	return f, err
}

//CustOrderCapacity is a non-required field for ExecutionReport.
func (m ExecutionReport) CustOrderCapacity() (field.CustOrderCapacity, errors.MessageRejectError) {
	var f field.CustOrderCapacity
	err := m.Body.Get(&f)
	return f, err
}

//Rule80A is a non-required field for ExecutionReport.
func (m ExecutionReport) Rule80A() (field.Rule80A, errors.MessageRejectError) {
	var f field.Rule80A
	err := m.Body.Get(&f)
	return f, err
}

//LastQty is a non-required field for ExecutionReport.
func (m ExecutionReport) LastQty() (field.LastQty, errors.MessageRejectError) {
	var f field.LastQty
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingLastQty is a non-required field for ExecutionReport.
func (m ExecutionReport) UnderlyingLastQty() (field.UnderlyingLastQty, errors.MessageRejectError) {
	var f field.UnderlyingLastQty
	err := m.Body.Get(&f)
	return f, err
}

//LastPx is a non-required field for ExecutionReport.
func (m ExecutionReport) LastPx() (field.LastPx, errors.MessageRejectError) {
	var f field.LastPx
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingLastPx is a non-required field for ExecutionReport.
func (m ExecutionReport) UnderlyingLastPx() (field.UnderlyingLastPx, errors.MessageRejectError) {
	var f field.UnderlyingLastPx
	err := m.Body.Get(&f)
	return f, err
}

//LastSpotRate is a non-required field for ExecutionReport.
func (m ExecutionReport) LastSpotRate() (field.LastSpotRate, errors.MessageRejectError) {
	var f field.LastSpotRate
	err := m.Body.Get(&f)
	return f, err
}

//LastForwardPoints is a non-required field for ExecutionReport.
func (m ExecutionReport) LastForwardPoints() (field.LastForwardPoints, errors.MessageRejectError) {
	var f field.LastForwardPoints
	err := m.Body.Get(&f)
	return f, err
}

//LastMkt is a non-required field for ExecutionReport.
func (m ExecutionReport) LastMkt() (field.LastMkt, errors.MessageRejectError) {
	var f field.LastMkt
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for ExecutionReport.
func (m ExecutionReport) TradingSessionID() (field.TradingSessionID, errors.MessageRejectError) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for ExecutionReport.
func (m ExecutionReport) TradingSessionSubID() (field.TradingSessionSubID, errors.MessageRejectError) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//LastCapacity is a non-required field for ExecutionReport.
func (m ExecutionReport) LastCapacity() (field.LastCapacity, errors.MessageRejectError) {
	var f field.LastCapacity
	err := m.Body.Get(&f)
	return f, err
}

//LeavesQty is a required field for ExecutionReport.
func (m ExecutionReport) LeavesQty() (field.LeavesQty, errors.MessageRejectError) {
	var f field.LeavesQty
	err := m.Body.Get(&f)
	return f, err
}

//CumQty is a required field for ExecutionReport.
func (m ExecutionReport) CumQty() (field.CumQty, errors.MessageRejectError) {
	var f field.CumQty
	err := m.Body.Get(&f)
	return f, err
}

//AvgPx is a required field for ExecutionReport.
func (m ExecutionReport) AvgPx() (field.AvgPx, errors.MessageRejectError) {
	var f field.AvgPx
	err := m.Body.Get(&f)
	return f, err
}

//DayOrderQty is a non-required field for ExecutionReport.
func (m ExecutionReport) DayOrderQty() (field.DayOrderQty, errors.MessageRejectError) {
	var f field.DayOrderQty
	err := m.Body.Get(&f)
	return f, err
}

//DayCumQty is a non-required field for ExecutionReport.
func (m ExecutionReport) DayCumQty() (field.DayCumQty, errors.MessageRejectError) {
	var f field.DayCumQty
	err := m.Body.Get(&f)
	return f, err
}

//DayAvgPx is a non-required field for ExecutionReport.
func (m ExecutionReport) DayAvgPx() (field.DayAvgPx, errors.MessageRejectError) {
	var f field.DayAvgPx
	err := m.Body.Get(&f)
	return f, err
}

//GTBookingInst is a non-required field for ExecutionReport.
func (m ExecutionReport) GTBookingInst() (field.GTBookingInst, errors.MessageRejectError) {
	var f field.GTBookingInst
	err := m.Body.Get(&f)
	return f, err
}

//TradeDate is a non-required field for ExecutionReport.
func (m ExecutionReport) TradeDate() (field.TradeDate, errors.MessageRejectError) {
	var f field.TradeDate
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for ExecutionReport.
func (m ExecutionReport) TransactTime() (field.TransactTime, errors.MessageRejectError) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//ReportToExch is a non-required field for ExecutionReport.
func (m ExecutionReport) ReportToExch() (field.ReportToExch, errors.MessageRejectError) {
	var f field.ReportToExch
	err := m.Body.Get(&f)
	return f, err
}

//Commission is a non-required field for ExecutionReport.
func (m ExecutionReport) Commission() (field.Commission, errors.MessageRejectError) {
	var f field.Commission
	err := m.Body.Get(&f)
	return f, err
}

//CommType is a non-required field for ExecutionReport.
func (m ExecutionReport) CommType() (field.CommType, errors.MessageRejectError) {
	var f field.CommType
	err := m.Body.Get(&f)
	return f, err
}

//CommCurrency is a non-required field for ExecutionReport.
func (m ExecutionReport) CommCurrency() (field.CommCurrency, errors.MessageRejectError) {
	var f field.CommCurrency
	err := m.Body.Get(&f)
	return f, err
}

//FundRenewWaiv is a non-required field for ExecutionReport.
func (m ExecutionReport) FundRenewWaiv() (field.FundRenewWaiv, errors.MessageRejectError) {
	var f field.FundRenewWaiv
	err := m.Body.Get(&f)
	return f, err
}

//Spread is a non-required field for ExecutionReport.
func (m ExecutionReport) Spread() (field.Spread, errors.MessageRejectError) {
	var f field.Spread
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for ExecutionReport.
func (m ExecutionReport) BenchmarkCurveCurrency() (field.BenchmarkCurveCurrency, errors.MessageRejectError) {
	var f field.BenchmarkCurveCurrency
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveName is a non-required field for ExecutionReport.
func (m ExecutionReport) BenchmarkCurveName() (field.BenchmarkCurveName, errors.MessageRejectError) {
	var f field.BenchmarkCurveName
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for ExecutionReport.
func (m ExecutionReport) BenchmarkCurvePoint() (field.BenchmarkCurvePoint, errors.MessageRejectError) {
	var f field.BenchmarkCurvePoint
	err := m.Body.Get(&f)
	return f, err
}

//YieldType is a non-required field for ExecutionReport.
func (m ExecutionReport) YieldType() (field.YieldType, errors.MessageRejectError) {
	var f field.YieldType
	err := m.Body.Get(&f)
	return f, err
}

//Yield is a non-required field for ExecutionReport.
func (m ExecutionReport) Yield() (field.Yield, errors.MessageRejectError) {
	var f field.Yield
	err := m.Body.Get(&f)
	return f, err
}

//GrossTradeAmt is a non-required field for ExecutionReport.
func (m ExecutionReport) GrossTradeAmt() (field.GrossTradeAmt, errors.MessageRejectError) {
	var f field.GrossTradeAmt
	err := m.Body.Get(&f)
	return f, err
}

//NumDaysInterest is a non-required field for ExecutionReport.
func (m ExecutionReport) NumDaysInterest() (field.NumDaysInterest, errors.MessageRejectError) {
	var f field.NumDaysInterest
	err := m.Body.Get(&f)
	return f, err
}

//ExDate is a non-required field for ExecutionReport.
func (m ExecutionReport) ExDate() (field.ExDate, errors.MessageRejectError) {
	var f field.ExDate
	err := m.Body.Get(&f)
	return f, err
}

//AccruedInterestRate is a non-required field for ExecutionReport.
func (m ExecutionReport) AccruedInterestRate() (field.AccruedInterestRate, errors.MessageRejectError) {
	var f field.AccruedInterestRate
	err := m.Body.Get(&f)
	return f, err
}

//AccruedInterestAmt is a non-required field for ExecutionReport.
func (m ExecutionReport) AccruedInterestAmt() (field.AccruedInterestAmt, errors.MessageRejectError) {
	var f field.AccruedInterestAmt
	err := m.Body.Get(&f)
	return f, err
}

//TradedFlatSwitch is a non-required field for ExecutionReport.
func (m ExecutionReport) TradedFlatSwitch() (field.TradedFlatSwitch, errors.MessageRejectError) {
	var f field.TradedFlatSwitch
	err := m.Body.Get(&f)
	return f, err
}

//BasisFeatureDate is a non-required field for ExecutionReport.
func (m ExecutionReport) BasisFeatureDate() (field.BasisFeatureDate, errors.MessageRejectError) {
	var f field.BasisFeatureDate
	err := m.Body.Get(&f)
	return f, err
}

//BasisFeaturePrice is a non-required field for ExecutionReport.
func (m ExecutionReport) BasisFeaturePrice() (field.BasisFeaturePrice, errors.MessageRejectError) {
	var f field.BasisFeaturePrice
	err := m.Body.Get(&f)
	return f, err
}

//Concession is a non-required field for ExecutionReport.
func (m ExecutionReport) Concession() (field.Concession, errors.MessageRejectError) {
	var f field.Concession
	err := m.Body.Get(&f)
	return f, err
}

//TotalTakedown is a non-required field for ExecutionReport.
func (m ExecutionReport) TotalTakedown() (field.TotalTakedown, errors.MessageRejectError) {
	var f field.TotalTakedown
	err := m.Body.Get(&f)
	return f, err
}

//NetMoney is a non-required field for ExecutionReport.
func (m ExecutionReport) NetMoney() (field.NetMoney, errors.MessageRejectError) {
	var f field.NetMoney
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrAmt is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrAmt() (field.SettlCurrAmt, errors.MessageRejectError) {
	var f field.SettlCurrAmt
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrency is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrency() (field.SettlCurrency, errors.MessageRejectError) {
	var f field.SettlCurrency
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrFxRate is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrFxRate() (field.SettlCurrFxRate, errors.MessageRejectError) {
	var f field.SettlCurrFxRate
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrFxRateCalc is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrFxRateCalc() (field.SettlCurrFxRateCalc, errors.MessageRejectError) {
	var f field.SettlCurrFxRateCalc
	err := m.Body.Get(&f)
	return f, err
}

//HandlInst is a non-required field for ExecutionReport.
func (m ExecutionReport) HandlInst() (field.HandlInst, errors.MessageRejectError) {
	var f field.HandlInst
	err := m.Body.Get(&f)
	return f, err
}

//MinQty is a non-required field for ExecutionReport.
func (m ExecutionReport) MinQty() (field.MinQty, errors.MessageRejectError) {
	var f field.MinQty
	err := m.Body.Get(&f)
	return f, err
}

//MaxFloor is a non-required field for ExecutionReport.
func (m ExecutionReport) MaxFloor() (field.MaxFloor, errors.MessageRejectError) {
	var f field.MaxFloor
	err := m.Body.Get(&f)
	return f, err
}

//PositionEffect is a non-required field for ExecutionReport.
func (m ExecutionReport) PositionEffect() (field.PositionEffect, errors.MessageRejectError) {
	var f field.PositionEffect
	err := m.Body.Get(&f)
	return f, err
}

//MaxShow is a non-required field for ExecutionReport.
func (m ExecutionReport) MaxShow() (field.MaxShow, errors.MessageRejectError) {
	var f field.MaxShow
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for ExecutionReport.
func (m ExecutionReport) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//FutSettDate2 is a non-required field for ExecutionReport.
func (m ExecutionReport) FutSettDate2() (field.FutSettDate2, errors.MessageRejectError) {
	var f field.FutSettDate2
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty2 is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderQty2() (field.OrderQty2, errors.MessageRejectError) {
	var f field.OrderQty2
	err := m.Body.Get(&f)
	return f, err
}

//LastForwardPoints2 is a non-required field for ExecutionReport.
func (m ExecutionReport) LastForwardPoints2() (field.LastForwardPoints2, errors.MessageRejectError) {
	var f field.LastForwardPoints2
	err := m.Body.Get(&f)
	return f, err
}

//MultiLegReportingType is a non-required field for ExecutionReport.
func (m ExecutionReport) MultiLegReportingType() (field.MultiLegReportingType, errors.MessageRejectError) {
	var f field.MultiLegReportingType
	err := m.Body.Get(&f)
	return f, err
}

//CancellationRights is a non-required field for ExecutionReport.
func (m ExecutionReport) CancellationRights() (field.CancellationRights, errors.MessageRejectError) {
	var f field.CancellationRights
	err := m.Body.Get(&f)
	return f, err
}

//MoneyLaunderingStatus is a non-required field for ExecutionReport.
func (m ExecutionReport) MoneyLaunderingStatus() (field.MoneyLaunderingStatus, errors.MessageRejectError) {
	var f field.MoneyLaunderingStatus
	err := m.Body.Get(&f)
	return f, err
}

//RegistID is a non-required field for ExecutionReport.
func (m ExecutionReport) RegistID() (field.RegistID, errors.MessageRejectError) {
	var f field.RegistID
	err := m.Body.Get(&f)
	return f, err
}

//Designation is a non-required field for ExecutionReport.
func (m ExecutionReport) Designation() (field.Designation, errors.MessageRejectError) {
	var f field.Designation
	err := m.Body.Get(&f)
	return f, err
}

//TransBkdTime is a non-required field for ExecutionReport.
func (m ExecutionReport) TransBkdTime() (field.TransBkdTime, errors.MessageRejectError) {
	var f field.TransBkdTime
	err := m.Body.Get(&f)
	return f, err
}

//ExecValuationPoint is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecValuationPoint() (field.ExecValuationPoint, errors.MessageRejectError) {
	var f field.ExecValuationPoint
	err := m.Body.Get(&f)
	return f, err
}

//ExecPriceType is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecPriceType() (field.ExecPriceType, errors.MessageRejectError) {
	var f field.ExecPriceType
	err := m.Body.Get(&f)
	return f, err
}

//ExecPriceAdjustment is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecPriceAdjustment() (field.ExecPriceAdjustment, errors.MessageRejectError) {
	var f field.ExecPriceAdjustment
	err := m.Body.Get(&f)
	return f, err
}

//PriorityIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) PriorityIndicator() (field.PriorityIndicator, errors.MessageRejectError) {
	var f field.PriorityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//PriceImprovement is a non-required field for ExecutionReport.
func (m ExecutionReport) PriceImprovement() (field.PriceImprovement, errors.MessageRejectError) {
	var f field.PriceImprovement
	err := m.Body.Get(&f)
	return f, err
}

//NoContAmts is a non-required field for ExecutionReport.
func (m ExecutionReport) NoContAmts() (field.NoContAmts, errors.MessageRejectError) {
	var f field.NoContAmts
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for ExecutionReport.
func (m ExecutionReport) NoLegs() (field.NoLegs, errors.MessageRejectError) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}
