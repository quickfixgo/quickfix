package fix50sp2

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
	cumqty field.CumQty) ExecutionReportBuilder {
	var builder ExecutionReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(orderid)
	builder.Body.Set(execid)
	builder.Body.Set(exectype)
	builder.Body.Set(ordstatus)
	builder.Body.Set(side)
	builder.Body.Set(leavesqty)
	builder.Body.Set(cumqty)
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

//QuoteRespID is a non-required field for ExecutionReport.
func (m ExecutionReport) QuoteRespID() (field.QuoteRespID, errors.MessageRejectError) {
	var f field.QuoteRespID
	err := m.Body.Get(&f)
	return f, err
}

//OrdStatusReqID is a non-required field for ExecutionReport.
func (m ExecutionReport) OrdStatusReqID() (field.OrdStatusReqID, errors.MessageRejectError) {
	var f field.OrdStatusReqID
	err := m.Body.Get(&f)
	return f, err
}

//MassStatusReqID is a non-required field for ExecutionReport.
func (m ExecutionReport) MassStatusReqID() (field.MassStatusReqID, errors.MessageRejectError) {
	var f field.MassStatusReqID
	err := m.Body.Get(&f)
	return f, err
}

//TotNumReports is a non-required field for ExecutionReport.
func (m ExecutionReport) TotNumReports() (field.TotNumReports, errors.MessageRejectError) {
	var f field.TotNumReports
	err := m.Body.Get(&f)
	return f, err
}

//LastRptRequested is a non-required field for ExecutionReport.
func (m ExecutionReport) LastRptRequested() (field.LastRptRequested, errors.MessageRejectError) {
	var f field.LastRptRequested
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

//AcctIDSource is a non-required field for ExecutionReport.
func (m ExecutionReport) AcctIDSource() (field.AcctIDSource, errors.MessageRejectError) {
	var f field.AcctIDSource
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

//SettlType is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlType() (field.SettlType, errors.MessageRejectError) {
	var f field.SettlType
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlDate() (field.SettlDate, errors.MessageRejectError) {
	var f field.SettlDate
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

//SecuritySubType is a non-required field for ExecutionReport.
func (m ExecutionReport) SecuritySubType() (field.SecuritySubType, errors.MessageRejectError) {
	var f field.SecuritySubType
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

//StrikeCurrency is a non-required field for ExecutionReport.
func (m ExecutionReport) StrikeCurrency() (field.StrikeCurrency, errors.MessageRejectError) {
	var f field.StrikeCurrency
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

//Pool is a non-required field for ExecutionReport.
func (m ExecutionReport) Pool() (field.Pool, errors.MessageRejectError) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for ExecutionReport.
func (m ExecutionReport) ContractSettlMonth() (field.ContractSettlMonth, errors.MessageRejectError) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for ExecutionReport.
func (m ExecutionReport) CPProgram() (field.CPProgram, errors.MessageRejectError) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for ExecutionReport.
func (m ExecutionReport) CPRegType() (field.CPRegType, errors.MessageRejectError) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for ExecutionReport.
func (m ExecutionReport) NoEvents() (field.NoEvents, errors.MessageRejectError) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for ExecutionReport.
func (m ExecutionReport) DatedDate() (field.DatedDate, errors.MessageRejectError) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for ExecutionReport.
func (m ExecutionReport) InterestAccrualDate() (field.InterestAccrualDate, errors.MessageRejectError) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityStatus() (field.SecurityStatus, errors.MessageRejectError) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for ExecutionReport.
func (m ExecutionReport) SettleOnOpenFlag() (field.SettleOnOpenFlag, errors.MessageRejectError) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for ExecutionReport.
func (m ExecutionReport) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for ExecutionReport.
func (m ExecutionReport) StrikeMultiplier() (field.StrikeMultiplier, errors.MessageRejectError) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for ExecutionReport.
func (m ExecutionReport) StrikeValue() (field.StrikeValue, errors.MessageRejectError) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for ExecutionReport.
func (m ExecutionReport) MinPriceIncrement() (field.MinPriceIncrement, errors.MessageRejectError) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for ExecutionReport.
func (m ExecutionReport) PositionLimit() (field.PositionLimit, errors.MessageRejectError) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for ExecutionReport.
func (m ExecutionReport) NTPositionLimit() (field.NTPositionLimit, errors.MessageRejectError) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for ExecutionReport.
func (m ExecutionReport) NoInstrumentParties() (field.NoInstrumentParties, errors.MessageRejectError) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for ExecutionReport.
func (m ExecutionReport) UnitOfMeasure() (field.UnitOfMeasure, errors.MessageRejectError) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for ExecutionReport.
func (m ExecutionReport) TimeUnit() (field.TimeUnit, errors.MessageRejectError) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for ExecutionReport.
func (m ExecutionReport) MaturityTime() (field.MaturityTime, errors.MessageRejectError) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//SecurityGroup is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityGroup() (field.SecurityGroup, errors.MessageRejectError) {
	var f field.SecurityGroup
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for ExecutionReport.
func (m ExecutionReport) MinPriceIncrementAmount() (field.MinPriceIncrementAmount, errors.MessageRejectError) {
	var f field.MinPriceIncrementAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for ExecutionReport.
func (m ExecutionReport) UnitOfMeasureQty() (field.UnitOfMeasureQty, errors.MessageRejectError) {
	var f field.UnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLLen is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityXMLLen() (field.SecurityXMLLen, errors.MessageRejectError) {
	var f field.SecurityXMLLen
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXML is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityXML() (field.SecurityXML, errors.MessageRejectError) {
	var f field.SecurityXML
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLSchema is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityXMLSchema() (field.SecurityXMLSchema, errors.MessageRejectError) {
	var f field.SecurityXMLSchema
	err := m.Body.Get(&f)
	return f, err
}

//ProductComplex is a non-required field for ExecutionReport.
func (m ExecutionReport) ProductComplex() (field.ProductComplex, errors.MessageRejectError) {
	var f field.ProductComplex
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for ExecutionReport.
func (m ExecutionReport) PriceUnitOfMeasure() (field.PriceUnitOfMeasure, errors.MessageRejectError) {
	var f field.PriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for ExecutionReport.
func (m ExecutionReport) PriceUnitOfMeasureQty() (field.PriceUnitOfMeasureQty, errors.MessageRejectError) {
	var f field.PriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SettlMethod is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlMethod() (field.SettlMethod, errors.MessageRejectError) {
	var f field.SettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//ExerciseStyle is a non-required field for ExecutionReport.
func (m ExecutionReport) ExerciseStyle() (field.ExerciseStyle, errors.MessageRejectError) {
	var f field.ExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutAmount is a non-required field for ExecutionReport.
func (m ExecutionReport) OptPayoutAmount() (field.OptPayoutAmount, errors.MessageRejectError) {
	var f field.OptPayoutAmount
	err := m.Body.Get(&f)
	return f, err
}

//PriceQuoteMethod is a non-required field for ExecutionReport.
func (m ExecutionReport) PriceQuoteMethod() (field.PriceQuoteMethod, errors.MessageRejectError) {
	var f field.PriceQuoteMethod
	err := m.Body.Get(&f)
	return f, err
}

//ListMethod is a non-required field for ExecutionReport.
func (m ExecutionReport) ListMethod() (field.ListMethod, errors.MessageRejectError) {
	var f field.ListMethod
	err := m.Body.Get(&f)
	return f, err
}

//CapPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) CapPrice() (field.CapPrice, errors.MessageRejectError) {
	var f field.CapPrice
	err := m.Body.Get(&f)
	return f, err
}

//FloorPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) FloorPrice() (field.FloorPrice, errors.MessageRejectError) {
	var f field.FloorPrice
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for ExecutionReport.
func (m ExecutionReport) PutOrCall() (field.PutOrCall, errors.MessageRejectError) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//FlexibleIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) FlexibleIndicator() (field.FlexibleIndicator, errors.MessageRejectError) {
	var f field.FlexibleIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) FlexProductEligibilityIndicator() (field.FlexProductEligibilityIndicator, errors.MessageRejectError) {
	var f field.FlexProductEligibilityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//ValuationMethod is a non-required field for ExecutionReport.
func (m ExecutionReport) ValuationMethod() (field.ValuationMethod, errors.MessageRejectError) {
	var f field.ValuationMethod
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplierUnit is a non-required field for ExecutionReport.
func (m ExecutionReport) ContractMultiplierUnit() (field.ContractMultiplierUnit, errors.MessageRejectError) {
	var f field.ContractMultiplierUnit
	err := m.Body.Get(&f)
	return f, err
}

//FlowScheduleType is a non-required field for ExecutionReport.
func (m ExecutionReport) FlowScheduleType() (field.FlowScheduleType, errors.MessageRejectError) {
	var f field.FlowScheduleType
	err := m.Body.Get(&f)
	return f, err
}

//RestructuringType is a non-required field for ExecutionReport.
func (m ExecutionReport) RestructuringType() (field.RestructuringType, errors.MessageRejectError) {
	var f field.RestructuringType
	err := m.Body.Get(&f)
	return f, err
}

//Seniority is a non-required field for ExecutionReport.
func (m ExecutionReport) Seniority() (field.Seniority, errors.MessageRejectError) {
	var f field.Seniority
	err := m.Body.Get(&f)
	return f, err
}

//NotionalPercentageOutstanding is a non-required field for ExecutionReport.
func (m ExecutionReport) NotionalPercentageOutstanding() (field.NotionalPercentageOutstanding, errors.MessageRejectError) {
	var f field.NotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//OriginalNotionalPercentageOutstanding is a non-required field for ExecutionReport.
func (m ExecutionReport) OriginalNotionalPercentageOutstanding() (field.OriginalNotionalPercentageOutstanding, errors.MessageRejectError) {
	var f field.OriginalNotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//AttachmentPoint is a non-required field for ExecutionReport.
func (m ExecutionReport) AttachmentPoint() (field.AttachmentPoint, errors.MessageRejectError) {
	var f field.AttachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//DetachmentPoint is a non-required field for ExecutionReport.
func (m ExecutionReport) DetachmentPoint() (field.DetachmentPoint, errors.MessageRejectError) {
	var f field.DetachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceDeterminationMethod is a non-required field for ExecutionReport.
func (m ExecutionReport) StrikePriceDeterminationMethod() (field.StrikePriceDeterminationMethod, errors.MessageRejectError) {
	var f field.StrikePriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryMethod is a non-required field for ExecutionReport.
func (m ExecutionReport) StrikePriceBoundaryMethod() (field.StrikePriceBoundaryMethod, errors.MessageRejectError) {
	var f field.StrikePriceBoundaryMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryPrecision is a non-required field for ExecutionReport.
func (m ExecutionReport) StrikePriceBoundaryPrecision() (field.StrikePriceBoundaryPrecision, errors.MessageRejectError) {
	var f field.StrikePriceBoundaryPrecision
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceDeterminationMethod is a non-required field for ExecutionReport.
func (m ExecutionReport) UnderlyingPriceDeterminationMethod() (field.UnderlyingPriceDeterminationMethod, errors.MessageRejectError) {
	var f field.UnderlyingPriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutType is a non-required field for ExecutionReport.
func (m ExecutionReport) OptPayoutType() (field.OptPayoutType, errors.MessageRejectError) {
	var f field.OptPayoutType
	err := m.Body.Get(&f)
	return f, err
}

//NoComplexEvents is a non-required field for ExecutionReport.
func (m ExecutionReport) NoComplexEvents() (field.NoComplexEvents, errors.MessageRejectError) {
	var f field.NoComplexEvents
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDesc is a non-required field for ExecutionReport.
func (m ExecutionReport) AgreementDesc() (field.AgreementDesc, errors.MessageRejectError) {
	var f field.AgreementDesc
	err := m.Body.Get(&f)
	return f, err
}

//AgreementID is a non-required field for ExecutionReport.
func (m ExecutionReport) AgreementID() (field.AgreementID, errors.MessageRejectError) {
	var f field.AgreementID
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDate is a non-required field for ExecutionReport.
func (m ExecutionReport) AgreementDate() (field.AgreementDate, errors.MessageRejectError) {
	var f field.AgreementDate
	err := m.Body.Get(&f)
	return f, err
}

//AgreementCurrency is a non-required field for ExecutionReport.
func (m ExecutionReport) AgreementCurrency() (field.AgreementCurrency, errors.MessageRejectError) {
	var f field.AgreementCurrency
	err := m.Body.Get(&f)
	return f, err
}

//TerminationType is a non-required field for ExecutionReport.
func (m ExecutionReport) TerminationType() (field.TerminationType, errors.MessageRejectError) {
	var f field.TerminationType
	err := m.Body.Get(&f)
	return f, err
}

//StartDate is a non-required field for ExecutionReport.
func (m ExecutionReport) StartDate() (field.StartDate, errors.MessageRejectError) {
	var f field.StartDate
	err := m.Body.Get(&f)
	return f, err
}

//EndDate is a non-required field for ExecutionReport.
func (m ExecutionReport) EndDate() (field.EndDate, errors.MessageRejectError) {
	var f field.EndDate
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryType is a non-required field for ExecutionReport.
func (m ExecutionReport) DeliveryType() (field.DeliveryType, errors.MessageRejectError) {
	var f field.DeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//MarginRatio is a non-required field for ExecutionReport.
func (m ExecutionReport) MarginRatio() (field.MarginRatio, errors.MessageRejectError) {
	var f field.MarginRatio
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for ExecutionReport.
func (m ExecutionReport) NoUnderlyings() (field.NoUnderlyings, errors.MessageRejectError) {
	var f field.NoUnderlyings
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

//QtyType is a non-required field for ExecutionReport.
func (m ExecutionReport) QtyType() (field.QtyType, errors.MessageRejectError) {
	var f field.QtyType
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

//PegOffsetValue is a non-required field for ExecutionReport.
func (m ExecutionReport) PegOffsetValue() (field.PegOffsetValue, errors.MessageRejectError) {
	var f field.PegOffsetValue
	err := m.Body.Get(&f)
	return f, err
}

//PegMoveType is a non-required field for ExecutionReport.
func (m ExecutionReport) PegMoveType() (field.PegMoveType, errors.MessageRejectError) {
	var f field.PegMoveType
	err := m.Body.Get(&f)
	return f, err
}

//PegOffsetType is a non-required field for ExecutionReport.
func (m ExecutionReport) PegOffsetType() (field.PegOffsetType, errors.MessageRejectError) {
	var f field.PegOffsetType
	err := m.Body.Get(&f)
	return f, err
}

//PegLimitType is a non-required field for ExecutionReport.
func (m ExecutionReport) PegLimitType() (field.PegLimitType, errors.MessageRejectError) {
	var f field.PegLimitType
	err := m.Body.Get(&f)
	return f, err
}

//PegRoundDirection is a non-required field for ExecutionReport.
func (m ExecutionReport) PegRoundDirection() (field.PegRoundDirection, errors.MessageRejectError) {
	var f field.PegRoundDirection
	err := m.Body.Get(&f)
	return f, err
}

//PegScope is a non-required field for ExecutionReport.
func (m ExecutionReport) PegScope() (field.PegScope, errors.MessageRejectError) {
	var f field.PegScope
	err := m.Body.Get(&f)
	return f, err
}

//PegPriceType is a non-required field for ExecutionReport.
func (m ExecutionReport) PegPriceType() (field.PegPriceType, errors.MessageRejectError) {
	var f field.PegPriceType
	err := m.Body.Get(&f)
	return f, err
}

//PegSecurityIDSource is a non-required field for ExecutionReport.
func (m ExecutionReport) PegSecurityIDSource() (field.PegSecurityIDSource, errors.MessageRejectError) {
	var f field.PegSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//PegSecurityID is a non-required field for ExecutionReport.
func (m ExecutionReport) PegSecurityID() (field.PegSecurityID, errors.MessageRejectError) {
	var f field.PegSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//PegSymbol is a non-required field for ExecutionReport.
func (m ExecutionReport) PegSymbol() (field.PegSymbol, errors.MessageRejectError) {
	var f field.PegSymbol
	err := m.Body.Get(&f)
	return f, err
}

//PegSecurityDesc is a non-required field for ExecutionReport.
func (m ExecutionReport) PegSecurityDesc() (field.PegSecurityDesc, errors.MessageRejectError) {
	var f field.PegSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionInst is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionInst() (field.DiscretionInst, errors.MessageRejectError) {
	var f field.DiscretionInst
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionOffsetValue is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionOffsetValue() (field.DiscretionOffsetValue, errors.MessageRejectError) {
	var f field.DiscretionOffsetValue
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionMoveType is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionMoveType() (field.DiscretionMoveType, errors.MessageRejectError) {
	var f field.DiscretionMoveType
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionOffsetType is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionOffsetType() (field.DiscretionOffsetType, errors.MessageRejectError) {
	var f field.DiscretionOffsetType
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionLimitType is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionLimitType() (field.DiscretionLimitType, errors.MessageRejectError) {
	var f field.DiscretionLimitType
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionRoundDirection is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionRoundDirection() (field.DiscretionRoundDirection, errors.MessageRejectError) {
	var f field.DiscretionRoundDirection
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionScope is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionScope() (field.DiscretionScope, errors.MessageRejectError) {
	var f field.DiscretionScope
	err := m.Body.Get(&f)
	return f, err
}

//PeggedPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) PeggedPrice() (field.PeggedPrice, errors.MessageRejectError) {
	var f field.PeggedPrice
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionPrice() (field.DiscretionPrice, errors.MessageRejectError) {
	var f field.DiscretionPrice
	err := m.Body.Get(&f)
	return f, err
}

//TargetStrategy is a non-required field for ExecutionReport.
func (m ExecutionReport) TargetStrategy() (field.TargetStrategy, errors.MessageRejectError) {
	var f field.TargetStrategy
	err := m.Body.Get(&f)
	return f, err
}

//TargetStrategyParameters is a non-required field for ExecutionReport.
func (m ExecutionReport) TargetStrategyParameters() (field.TargetStrategyParameters, errors.MessageRejectError) {
	var f field.TargetStrategyParameters
	err := m.Body.Get(&f)
	return f, err
}

//ParticipationRate is a non-required field for ExecutionReport.
func (m ExecutionReport) ParticipationRate() (field.ParticipationRate, errors.MessageRejectError) {
	var f field.ParticipationRate
	err := m.Body.Get(&f)
	return f, err
}

//TargetStrategyPerformance is a non-required field for ExecutionReport.
func (m ExecutionReport) TargetStrategyPerformance() (field.TargetStrategyPerformance, errors.MessageRejectError) {
	var f field.TargetStrategyPerformance
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

//LastParPx is a non-required field for ExecutionReport.
func (m ExecutionReport) LastParPx() (field.LastParPx, errors.MessageRejectError) {
	var f field.LastParPx
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

//TimeBracket is a non-required field for ExecutionReport.
func (m ExecutionReport) TimeBracket() (field.TimeBracket, errors.MessageRejectError) {
	var f field.TimeBracket
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

//AvgPx is a non-required field for ExecutionReport.
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

//BenchmarkPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) BenchmarkPrice() (field.BenchmarkPrice, errors.MessageRejectError) {
	var f field.BenchmarkPrice
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPriceType is a non-required field for ExecutionReport.
func (m ExecutionReport) BenchmarkPriceType() (field.BenchmarkPriceType, errors.MessageRejectError) {
	var f field.BenchmarkPriceType
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for ExecutionReport.
func (m ExecutionReport) BenchmarkSecurityID() (field.BenchmarkSecurityID, errors.MessageRejectError) {
	var f field.BenchmarkSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for ExecutionReport.
func (m ExecutionReport) BenchmarkSecurityIDSource() (field.BenchmarkSecurityIDSource, errors.MessageRejectError) {
	var f field.BenchmarkSecurityIDSource
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

//YieldCalcDate is a non-required field for ExecutionReport.
func (m ExecutionReport) YieldCalcDate() (field.YieldCalcDate, errors.MessageRejectError) {
	var f field.YieldCalcDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionDate is a non-required field for ExecutionReport.
func (m ExecutionReport) YieldRedemptionDate() (field.YieldRedemptionDate, errors.MessageRejectError) {
	var f field.YieldRedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) YieldRedemptionPrice() (field.YieldRedemptionPrice, errors.MessageRejectError) {
	var f field.YieldRedemptionPrice
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPriceType is a non-required field for ExecutionReport.
func (m ExecutionReport) YieldRedemptionPriceType() (field.YieldRedemptionPriceType, errors.MessageRejectError) {
	var f field.YieldRedemptionPriceType
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

//InterestAtMaturity is a non-required field for ExecutionReport.
func (m ExecutionReport) InterestAtMaturity() (field.InterestAtMaturity, errors.MessageRejectError) {
	var f field.InterestAtMaturity
	err := m.Body.Get(&f)
	return f, err
}

//EndAccruedInterestAmt is a non-required field for ExecutionReport.
func (m ExecutionReport) EndAccruedInterestAmt() (field.EndAccruedInterestAmt, errors.MessageRejectError) {
	var f field.EndAccruedInterestAmt
	err := m.Body.Get(&f)
	return f, err
}

//StartCash is a non-required field for ExecutionReport.
func (m ExecutionReport) StartCash() (field.StartCash, errors.MessageRejectError) {
	var f field.StartCash
	err := m.Body.Get(&f)
	return f, err
}

//EndCash is a non-required field for ExecutionReport.
func (m ExecutionReport) EndCash() (field.EndCash, errors.MessageRejectError) {
	var f field.EndCash
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

//BookingType is a non-required field for ExecutionReport.
func (m ExecutionReport) BookingType() (field.BookingType, errors.MessageRejectError) {
	var f field.BookingType
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

//SettlDate2 is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlDate2() (field.SettlDate2, errors.MessageRejectError) {
	var f field.SettlDate2
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

//LastLiquidityInd is a non-required field for ExecutionReport.
func (m ExecutionReport) LastLiquidityInd() (field.LastLiquidityInd, errors.MessageRejectError) {
	var f field.LastLiquidityInd
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

//CopyMsgIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) CopyMsgIndicator() (field.CopyMsgIndicator, errors.MessageRejectError) {
	var f field.CopyMsgIndicator
	err := m.Body.Get(&f)
	return f, err
}

//NoMiscFees is a non-required field for ExecutionReport.
func (m ExecutionReport) NoMiscFees() (field.NoMiscFees, errors.MessageRejectError) {
	var f field.NoMiscFees
	err := m.Body.Get(&f)
	return f, err
}

//NoStrategyParameters is a non-required field for ExecutionReport.
func (m ExecutionReport) NoStrategyParameters() (field.NoStrategyParameters, errors.MessageRejectError) {
	var f field.NoStrategyParameters
	err := m.Body.Get(&f)
	return f, err
}

//HostCrossID is a non-required field for ExecutionReport.
func (m ExecutionReport) HostCrossID() (field.HostCrossID, errors.MessageRejectError) {
	var f field.HostCrossID
	err := m.Body.Get(&f)
	return f, err
}

//ManualOrderIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) ManualOrderIndicator() (field.ManualOrderIndicator, errors.MessageRejectError) {
	var f field.ManualOrderIndicator
	err := m.Body.Get(&f)
	return f, err
}

//CustDirectedOrder is a non-required field for ExecutionReport.
func (m ExecutionReport) CustDirectedOrder() (field.CustDirectedOrder, errors.MessageRejectError) {
	var f field.CustDirectedOrder
	err := m.Body.Get(&f)
	return f, err
}

//ReceivedDeptID is a non-required field for ExecutionReport.
func (m ExecutionReport) ReceivedDeptID() (field.ReceivedDeptID, errors.MessageRejectError) {
	var f field.ReceivedDeptID
	err := m.Body.Get(&f)
	return f, err
}

//CustOrderHandlingInst is a non-required field for ExecutionReport.
func (m ExecutionReport) CustOrderHandlingInst() (field.CustOrderHandlingInst, errors.MessageRejectError) {
	var f field.CustOrderHandlingInst
	err := m.Body.Get(&f)
	return f, err
}

//OrderHandlingInstSource is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderHandlingInstSource() (field.OrderHandlingInstSource, errors.MessageRejectError) {
	var f field.OrderHandlingInstSource
	err := m.Body.Get(&f)
	return f, err
}

//NoTrdRegTimestamps is a non-required field for ExecutionReport.
func (m ExecutionReport) NoTrdRegTimestamps() (field.NoTrdRegTimestamps, errors.MessageRejectError) {
	var f field.NoTrdRegTimestamps
	err := m.Body.Get(&f)
	return f, err
}

//AggressorIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) AggressorIndicator() (field.AggressorIndicator, errors.MessageRejectError) {
	var f field.AggressorIndicator
	err := m.Body.Get(&f)
	return f, err
}

//CalculatedCcyLastQty is a non-required field for ExecutionReport.
func (m ExecutionReport) CalculatedCcyLastQty() (field.CalculatedCcyLastQty, errors.MessageRejectError) {
	var f field.CalculatedCcyLastQty
	err := m.Body.Get(&f)
	return f, err
}

//LastSwapPoints is a non-required field for ExecutionReport.
func (m ExecutionReport) LastSwapPoints() (field.LastSwapPoints, errors.MessageRejectError) {
	var f field.LastSwapPoints
	err := m.Body.Get(&f)
	return f, err
}

//MatchType is a non-required field for ExecutionReport.
func (m ExecutionReport) MatchType() (field.MatchType, errors.MessageRejectError) {
	var f field.MatchType
	err := m.Body.Get(&f)
	return f, err
}

//OrderCategory is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderCategory() (field.OrderCategory, errors.MessageRejectError) {
	var f field.OrderCategory
	err := m.Body.Get(&f)
	return f, err
}

//LotType is a non-required field for ExecutionReport.
func (m ExecutionReport) LotType() (field.LotType, errors.MessageRejectError) {
	var f field.LotType
	err := m.Body.Get(&f)
	return f, err
}

//PriceProtectionScope is a non-required field for ExecutionReport.
func (m ExecutionReport) PriceProtectionScope() (field.PriceProtectionScope, errors.MessageRejectError) {
	var f field.PriceProtectionScope
	err := m.Body.Get(&f)
	return f, err
}

//TriggerType is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerType() (field.TriggerType, errors.MessageRejectError) {
	var f field.TriggerType
	err := m.Body.Get(&f)
	return f, err
}

//TriggerAction is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerAction() (field.TriggerAction, errors.MessageRejectError) {
	var f field.TriggerAction
	err := m.Body.Get(&f)
	return f, err
}

//TriggerPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerPrice() (field.TriggerPrice, errors.MessageRejectError) {
	var f field.TriggerPrice
	err := m.Body.Get(&f)
	return f, err
}

//TriggerSymbol is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerSymbol() (field.TriggerSymbol, errors.MessageRejectError) {
	var f field.TriggerSymbol
	err := m.Body.Get(&f)
	return f, err
}

//TriggerSecurityID is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerSecurityID() (field.TriggerSecurityID, errors.MessageRejectError) {
	var f field.TriggerSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//TriggerSecurityIDSource is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerSecurityIDSource() (field.TriggerSecurityIDSource, errors.MessageRejectError) {
	var f field.TriggerSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//TriggerSecurityDesc is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerSecurityDesc() (field.TriggerSecurityDesc, errors.MessageRejectError) {
	var f field.TriggerSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//TriggerPriceType is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerPriceType() (field.TriggerPriceType, errors.MessageRejectError) {
	var f field.TriggerPriceType
	err := m.Body.Get(&f)
	return f, err
}

//TriggerPriceTypeScope is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerPriceTypeScope() (field.TriggerPriceTypeScope, errors.MessageRejectError) {
	var f field.TriggerPriceTypeScope
	err := m.Body.Get(&f)
	return f, err
}

//TriggerPriceDirection is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerPriceDirection() (field.TriggerPriceDirection, errors.MessageRejectError) {
	var f field.TriggerPriceDirection
	err := m.Body.Get(&f)
	return f, err
}

//TriggerNewPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerNewPrice() (field.TriggerNewPrice, errors.MessageRejectError) {
	var f field.TriggerNewPrice
	err := m.Body.Get(&f)
	return f, err
}

//TriggerOrderType is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerOrderType() (field.TriggerOrderType, errors.MessageRejectError) {
	var f field.TriggerOrderType
	err := m.Body.Get(&f)
	return f, err
}

//TriggerNewQty is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerNewQty() (field.TriggerNewQty, errors.MessageRejectError) {
	var f field.TriggerNewQty
	err := m.Body.Get(&f)
	return f, err
}

//TriggerTradingSessionID is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerTradingSessionID() (field.TriggerTradingSessionID, errors.MessageRejectError) {
	var f field.TriggerTradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TriggerTradingSessionSubID is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerTradingSessionSubID() (field.TriggerTradingSessionSubID, errors.MessageRejectError) {
	var f field.TriggerTradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//PeggedRefPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) PeggedRefPrice() (field.PeggedRefPrice, errors.MessageRejectError) {
	var f field.PeggedRefPrice
	err := m.Body.Get(&f)
	return f, err
}

//PreTradeAnonymity is a non-required field for ExecutionReport.
func (m ExecutionReport) PreTradeAnonymity() (field.PreTradeAnonymity, errors.MessageRejectError) {
	var f field.PreTradeAnonymity
	err := m.Body.Get(&f)
	return f, err
}

//MatchIncrement is a non-required field for ExecutionReport.
func (m ExecutionReport) MatchIncrement() (field.MatchIncrement, errors.MessageRejectError) {
	var f field.MatchIncrement
	err := m.Body.Get(&f)
	return f, err
}

//MaxPriceLevels is a non-required field for ExecutionReport.
func (m ExecutionReport) MaxPriceLevels() (field.MaxPriceLevels, errors.MessageRejectError) {
	var f field.MaxPriceLevels
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryDisplayQty is a non-required field for ExecutionReport.
func (m ExecutionReport) SecondaryDisplayQty() (field.SecondaryDisplayQty, errors.MessageRejectError) {
	var f field.SecondaryDisplayQty
	err := m.Body.Get(&f)
	return f, err
}

//DisplayWhen is a non-required field for ExecutionReport.
func (m ExecutionReport) DisplayWhen() (field.DisplayWhen, errors.MessageRejectError) {
	var f field.DisplayWhen
	err := m.Body.Get(&f)
	return f, err
}

//DisplayMethod is a non-required field for ExecutionReport.
func (m ExecutionReport) DisplayMethod() (field.DisplayMethod, errors.MessageRejectError) {
	var f field.DisplayMethod
	err := m.Body.Get(&f)
	return f, err
}

//DisplayLowQty is a non-required field for ExecutionReport.
func (m ExecutionReport) DisplayLowQty() (field.DisplayLowQty, errors.MessageRejectError) {
	var f field.DisplayLowQty
	err := m.Body.Get(&f)
	return f, err
}

//DisplayHighQty is a non-required field for ExecutionReport.
func (m ExecutionReport) DisplayHighQty() (field.DisplayHighQty, errors.MessageRejectError) {
	var f field.DisplayHighQty
	err := m.Body.Get(&f)
	return f, err
}

//DisplayMinIncr is a non-required field for ExecutionReport.
func (m ExecutionReport) DisplayMinIncr() (field.DisplayMinIncr, errors.MessageRejectError) {
	var f field.DisplayMinIncr
	err := m.Body.Get(&f)
	return f, err
}

//RefreshQty is a non-required field for ExecutionReport.
func (m ExecutionReport) RefreshQty() (field.RefreshQty, errors.MessageRejectError) {
	var f field.RefreshQty
	err := m.Body.Get(&f)
	return f, err
}

//DisplayQty is a non-required field for ExecutionReport.
func (m ExecutionReport) DisplayQty() (field.DisplayQty, errors.MessageRejectError) {
	var f field.DisplayQty
	err := m.Body.Get(&f)
	return f, err
}

//Volatility is a non-required field for ExecutionReport.
func (m ExecutionReport) Volatility() (field.Volatility, errors.MessageRejectError) {
	var f field.Volatility
	err := m.Body.Get(&f)
	return f, err
}

//TimeToExpiration is a non-required field for ExecutionReport.
func (m ExecutionReport) TimeToExpiration() (field.TimeToExpiration, errors.MessageRejectError) {
	var f field.TimeToExpiration
	err := m.Body.Get(&f)
	return f, err
}

//RiskFreeRate is a non-required field for ExecutionReport.
func (m ExecutionReport) RiskFreeRate() (field.RiskFreeRate, errors.MessageRejectError) {
	var f field.RiskFreeRate
	err := m.Body.Get(&f)
	return f, err
}

//PriceDelta is a non-required field for ExecutionReport.
func (m ExecutionReport) PriceDelta() (field.PriceDelta, errors.MessageRejectError) {
	var f field.PriceDelta
	err := m.Body.Get(&f)
	return f, err
}

//TrdMatchID is a non-required field for ExecutionReport.
func (m ExecutionReport) TrdMatchID() (field.TrdMatchID, errors.MessageRejectError) {
	var f field.TrdMatchID
	err := m.Body.Get(&f)
	return f, err
}

//AllocID is a non-required field for ExecutionReport.
func (m ExecutionReport) AllocID() (field.AllocID, errors.MessageRejectError) {
	var f field.AllocID
	err := m.Body.Get(&f)
	return f, err
}

//NoAllocs is a non-required field for ExecutionReport.
func (m ExecutionReport) NoAllocs() (field.NoAllocs, errors.MessageRejectError) {
	var f field.NoAllocs
	err := m.Body.Get(&f)
	return f, err
}

//TotNoFills is a non-required field for ExecutionReport.
func (m ExecutionReport) TotNoFills() (field.TotNoFills, errors.MessageRejectError) {
	var f field.TotNoFills
	err := m.Body.Get(&f)
	return f, err
}

//LastFragment is a non-required field for ExecutionReport.
func (m ExecutionReport) LastFragment() (field.LastFragment, errors.MessageRejectError) {
	var f field.LastFragment
	err := m.Body.Get(&f)
	return f, err
}

//NoFills is a non-required field for ExecutionReport.
func (m ExecutionReport) NoFills() (field.NoFills, errors.MessageRejectError) {
	var f field.NoFills
	err := m.Body.Get(&f)
	return f, err
}

//DividendYield is a non-required field for ExecutionReport.
func (m ExecutionReport) DividendYield() (field.DividendYield, errors.MessageRejectError) {
	var f field.DividendYield
	err := m.Body.Get(&f)
	return f, err
}

//ApplID is a non-required field for ExecutionReport.
func (m ExecutionReport) ApplID() (field.ApplID, errors.MessageRejectError) {
	var f field.ApplID
	err := m.Body.Get(&f)
	return f, err
}

//ApplSeqNum is a non-required field for ExecutionReport.
func (m ExecutionReport) ApplSeqNum() (field.ApplSeqNum, errors.MessageRejectError) {
	var f field.ApplSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplLastSeqNum is a non-required field for ExecutionReport.
func (m ExecutionReport) ApplLastSeqNum() (field.ApplLastSeqNum, errors.MessageRejectError) {
	var f field.ApplLastSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplResendFlag is a non-required field for ExecutionReport.
func (m ExecutionReport) ApplResendFlag() (field.ApplResendFlag, errors.MessageRejectError) {
	var f field.ApplResendFlag
	err := m.Body.Get(&f)
	return f, err
}

//NoRateSources is a non-required field for ExecutionReport.
func (m ExecutionReport) NoRateSources() (field.NoRateSources, errors.MessageRejectError) {
	var f field.NoRateSources
	err := m.Body.Get(&f)
	return f, err
}
