package fix50sp1

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

//QuoteRespID is a non-required field for ExecutionReport.
func (m ExecutionReport) QuoteRespID() (field.QuoteRespID, error) {
	var f field.QuoteRespID
	err := m.Body.Get(&f)
	return f, err
}

//OrdStatusReqID is a non-required field for ExecutionReport.
func (m ExecutionReport) OrdStatusReqID() (field.OrdStatusReqID, error) {
	var f field.OrdStatusReqID
	err := m.Body.Get(&f)
	return f, err
}

//MassStatusReqID is a non-required field for ExecutionReport.
func (m ExecutionReport) MassStatusReqID() (field.MassStatusReqID, error) {
	var f field.MassStatusReqID
	err := m.Body.Get(&f)
	return f, err
}

//TotNumReports is a non-required field for ExecutionReport.
func (m ExecutionReport) TotNumReports() (field.TotNumReports, error) {
	var f field.TotNumReports
	err := m.Body.Get(&f)
	return f, err
}

//LastRptRequested is a non-required field for ExecutionReport.
func (m ExecutionReport) LastRptRequested() (field.LastRptRequested, error) {
	var f field.LastRptRequested
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

//AcctIDSource is a non-required field for ExecutionReport.
func (m ExecutionReport) AcctIDSource() (field.AcctIDSource, error) {
	var f field.AcctIDSource
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

//SettlType is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlType() (field.SettlType, error) {
	var f field.SettlType
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlDate() (field.SettlDate, error) {
	var f field.SettlDate
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

//SecuritySubType is a non-required field for ExecutionReport.
func (m ExecutionReport) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
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

//StrikeCurrency is a non-required field for ExecutionReport.
func (m ExecutionReport) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
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

//Pool is a non-required field for ExecutionReport.
func (m ExecutionReport) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for ExecutionReport.
func (m ExecutionReport) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for ExecutionReport.
func (m ExecutionReport) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for ExecutionReport.
func (m ExecutionReport) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for ExecutionReport.
func (m ExecutionReport) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for ExecutionReport.
func (m ExecutionReport) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for ExecutionReport.
func (m ExecutionReport) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityStatus() (field.SecurityStatus, error) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for ExecutionReport.
func (m ExecutionReport) SettleOnOpenFlag() (field.SettleOnOpenFlag, error) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for ExecutionReport.
func (m ExecutionReport) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, error) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for ExecutionReport.
func (m ExecutionReport) StrikeMultiplier() (field.StrikeMultiplier, error) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for ExecutionReport.
func (m ExecutionReport) StrikeValue() (field.StrikeValue, error) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for ExecutionReport.
func (m ExecutionReport) MinPriceIncrement() (field.MinPriceIncrement, error) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for ExecutionReport.
func (m ExecutionReport) PositionLimit() (field.PositionLimit, error) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for ExecutionReport.
func (m ExecutionReport) NTPositionLimit() (field.NTPositionLimit, error) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for ExecutionReport.
func (m ExecutionReport) NoInstrumentParties() (field.NoInstrumentParties, error) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for ExecutionReport.
func (m ExecutionReport) UnitOfMeasure() (field.UnitOfMeasure, error) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for ExecutionReport.
func (m ExecutionReport) TimeUnit() (field.TimeUnit, error) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for ExecutionReport.
func (m ExecutionReport) MaturityTime() (field.MaturityTime, error) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//SecurityGroup is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityGroup() (field.SecurityGroup, error) {
	var f field.SecurityGroup
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for ExecutionReport.
func (m ExecutionReport) MinPriceIncrementAmount() (field.MinPriceIncrementAmount, error) {
	var f field.MinPriceIncrementAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for ExecutionReport.
func (m ExecutionReport) UnitOfMeasureQty() (field.UnitOfMeasureQty, error) {
	var f field.UnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLLen is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityXMLLen() (field.SecurityXMLLen, error) {
	var f field.SecurityXMLLen
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXML is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityXML() (field.SecurityXML, error) {
	var f field.SecurityXML
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLSchema is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityXMLSchema() (field.SecurityXMLSchema, error) {
	var f field.SecurityXMLSchema
	err := m.Body.Get(&f)
	return f, err
}

//ProductComplex is a non-required field for ExecutionReport.
func (m ExecutionReport) ProductComplex() (field.ProductComplex, error) {
	var f field.ProductComplex
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for ExecutionReport.
func (m ExecutionReport) PriceUnitOfMeasure() (field.PriceUnitOfMeasure, error) {
	var f field.PriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for ExecutionReport.
func (m ExecutionReport) PriceUnitOfMeasureQty() (field.PriceUnitOfMeasureQty, error) {
	var f field.PriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SettlMethod is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlMethod() (field.SettlMethod, error) {
	var f field.SettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//ExerciseStyle is a non-required field for ExecutionReport.
func (m ExecutionReport) ExerciseStyle() (field.ExerciseStyle, error) {
	var f field.ExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//OptPayAmount is a non-required field for ExecutionReport.
func (m ExecutionReport) OptPayAmount() (field.OptPayAmount, error) {
	var f field.OptPayAmount
	err := m.Body.Get(&f)
	return f, err
}

//PriceQuoteMethod is a non-required field for ExecutionReport.
func (m ExecutionReport) PriceQuoteMethod() (field.PriceQuoteMethod, error) {
	var f field.PriceQuoteMethod
	err := m.Body.Get(&f)
	return f, err
}

//ListMethod is a non-required field for ExecutionReport.
func (m ExecutionReport) ListMethod() (field.ListMethod, error) {
	var f field.ListMethod
	err := m.Body.Get(&f)
	return f, err
}

//CapPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) CapPrice() (field.CapPrice, error) {
	var f field.CapPrice
	err := m.Body.Get(&f)
	return f, err
}

//FloorPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) FloorPrice() (field.FloorPrice, error) {
	var f field.FloorPrice
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for ExecutionReport.
func (m ExecutionReport) PutOrCall() (field.PutOrCall, error) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//FlexibleIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) FlexibleIndicator() (field.FlexibleIndicator, error) {
	var f field.FlexibleIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) FlexProductEligibilityIndicator() (field.FlexProductEligibilityIndicator, error) {
	var f field.FlexProductEligibilityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FuturesValuationMethod is a non-required field for ExecutionReport.
func (m ExecutionReport) FuturesValuationMethod() (field.FuturesValuationMethod, error) {
	var f field.FuturesValuationMethod
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDesc is a non-required field for ExecutionReport.
func (m ExecutionReport) AgreementDesc() (field.AgreementDesc, error) {
	var f field.AgreementDesc
	err := m.Body.Get(&f)
	return f, err
}

//AgreementID is a non-required field for ExecutionReport.
func (m ExecutionReport) AgreementID() (field.AgreementID, error) {
	var f field.AgreementID
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDate is a non-required field for ExecutionReport.
func (m ExecutionReport) AgreementDate() (field.AgreementDate, error) {
	var f field.AgreementDate
	err := m.Body.Get(&f)
	return f, err
}

//AgreementCurrency is a non-required field for ExecutionReport.
func (m ExecutionReport) AgreementCurrency() (field.AgreementCurrency, error) {
	var f field.AgreementCurrency
	err := m.Body.Get(&f)
	return f, err
}

//TerminationType is a non-required field for ExecutionReport.
func (m ExecutionReport) TerminationType() (field.TerminationType, error) {
	var f field.TerminationType
	err := m.Body.Get(&f)
	return f, err
}

//StartDate is a non-required field for ExecutionReport.
func (m ExecutionReport) StartDate() (field.StartDate, error) {
	var f field.StartDate
	err := m.Body.Get(&f)
	return f, err
}

//EndDate is a non-required field for ExecutionReport.
func (m ExecutionReport) EndDate() (field.EndDate, error) {
	var f field.EndDate
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryType is a non-required field for ExecutionReport.
func (m ExecutionReport) DeliveryType() (field.DeliveryType, error) {
	var f field.DeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//MarginRatio is a non-required field for ExecutionReport.
func (m ExecutionReport) MarginRatio() (field.MarginRatio, error) {
	var f field.MarginRatio
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for ExecutionReport.
func (m ExecutionReport) NoUnderlyings() (field.NoUnderlyings, error) {
	var f field.NoUnderlyings
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

//QtyType is a non-required field for ExecutionReport.
func (m ExecutionReport) QtyType() (field.QtyType, error) {
	var f field.QtyType
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

//PegOffsetValue is a non-required field for ExecutionReport.
func (m ExecutionReport) PegOffsetValue() (field.PegOffsetValue, error) {
	var f field.PegOffsetValue
	err := m.Body.Get(&f)
	return f, err
}

//PegMoveType is a non-required field for ExecutionReport.
func (m ExecutionReport) PegMoveType() (field.PegMoveType, error) {
	var f field.PegMoveType
	err := m.Body.Get(&f)
	return f, err
}

//PegOffsetType is a non-required field for ExecutionReport.
func (m ExecutionReport) PegOffsetType() (field.PegOffsetType, error) {
	var f field.PegOffsetType
	err := m.Body.Get(&f)
	return f, err
}

//PegLimitType is a non-required field for ExecutionReport.
func (m ExecutionReport) PegLimitType() (field.PegLimitType, error) {
	var f field.PegLimitType
	err := m.Body.Get(&f)
	return f, err
}

//PegRoundDirection is a non-required field for ExecutionReport.
func (m ExecutionReport) PegRoundDirection() (field.PegRoundDirection, error) {
	var f field.PegRoundDirection
	err := m.Body.Get(&f)
	return f, err
}

//PegScope is a non-required field for ExecutionReport.
func (m ExecutionReport) PegScope() (field.PegScope, error) {
	var f field.PegScope
	err := m.Body.Get(&f)
	return f, err
}

//PegPriceType is a non-required field for ExecutionReport.
func (m ExecutionReport) PegPriceType() (field.PegPriceType, error) {
	var f field.PegPriceType
	err := m.Body.Get(&f)
	return f, err
}

//PegSecurityIDSource is a non-required field for ExecutionReport.
func (m ExecutionReport) PegSecurityIDSource() (field.PegSecurityIDSource, error) {
	var f field.PegSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//PegSecurityID is a non-required field for ExecutionReport.
func (m ExecutionReport) PegSecurityID() (field.PegSecurityID, error) {
	var f field.PegSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//PegSymbol is a non-required field for ExecutionReport.
func (m ExecutionReport) PegSymbol() (field.PegSymbol, error) {
	var f field.PegSymbol
	err := m.Body.Get(&f)
	return f, err
}

//PegSecurityDesc is a non-required field for ExecutionReport.
func (m ExecutionReport) PegSecurityDesc() (field.PegSecurityDesc, error) {
	var f field.PegSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionInst is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionInst() (field.DiscretionInst, error) {
	var f field.DiscretionInst
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionOffsetValue is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionOffsetValue() (field.DiscretionOffsetValue, error) {
	var f field.DiscretionOffsetValue
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionMoveType is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionMoveType() (field.DiscretionMoveType, error) {
	var f field.DiscretionMoveType
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionOffsetType is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionOffsetType() (field.DiscretionOffsetType, error) {
	var f field.DiscretionOffsetType
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionLimitType is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionLimitType() (field.DiscretionLimitType, error) {
	var f field.DiscretionLimitType
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionRoundDirection is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionRoundDirection() (field.DiscretionRoundDirection, error) {
	var f field.DiscretionRoundDirection
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionScope is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionScope() (field.DiscretionScope, error) {
	var f field.DiscretionScope
	err := m.Body.Get(&f)
	return f, err
}

//PeggedPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) PeggedPrice() (field.PeggedPrice, error) {
	var f field.PeggedPrice
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionPrice() (field.DiscretionPrice, error) {
	var f field.DiscretionPrice
	err := m.Body.Get(&f)
	return f, err
}

//TargetStrategy is a non-required field for ExecutionReport.
func (m ExecutionReport) TargetStrategy() (field.TargetStrategy, error) {
	var f field.TargetStrategy
	err := m.Body.Get(&f)
	return f, err
}

//TargetStrategyParameters is a non-required field for ExecutionReport.
func (m ExecutionReport) TargetStrategyParameters() (field.TargetStrategyParameters, error) {
	var f field.TargetStrategyParameters
	err := m.Body.Get(&f)
	return f, err
}

//ParticipationRate is a non-required field for ExecutionReport.
func (m ExecutionReport) ParticipationRate() (field.ParticipationRate, error) {
	var f field.ParticipationRate
	err := m.Body.Get(&f)
	return f, err
}

//TargetStrategyPerformance is a non-required field for ExecutionReport.
func (m ExecutionReport) TargetStrategyPerformance() (field.TargetStrategyPerformance, error) {
	var f field.TargetStrategyPerformance
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

//LastParPx is a non-required field for ExecutionReport.
func (m ExecutionReport) LastParPx() (field.LastParPx, error) {
	var f field.LastParPx
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

//TimeBracket is a non-required field for ExecutionReport.
func (m ExecutionReport) TimeBracket() (field.TimeBracket, error) {
	var f field.TimeBracket
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

//AvgPx is a non-required field for ExecutionReport.
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

//BenchmarkPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) BenchmarkPrice() (field.BenchmarkPrice, error) {
	var f field.BenchmarkPrice
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPriceType is a non-required field for ExecutionReport.
func (m ExecutionReport) BenchmarkPriceType() (field.BenchmarkPriceType, error) {
	var f field.BenchmarkPriceType
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for ExecutionReport.
func (m ExecutionReport) BenchmarkSecurityID() (field.BenchmarkSecurityID, error) {
	var f field.BenchmarkSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for ExecutionReport.
func (m ExecutionReport) BenchmarkSecurityIDSource() (field.BenchmarkSecurityIDSource, error) {
	var f field.BenchmarkSecurityIDSource
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

//YieldCalcDate is a non-required field for ExecutionReport.
func (m ExecutionReport) YieldCalcDate() (field.YieldCalcDate, error) {
	var f field.YieldCalcDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionDate is a non-required field for ExecutionReport.
func (m ExecutionReport) YieldRedemptionDate() (field.YieldRedemptionDate, error) {
	var f field.YieldRedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) YieldRedemptionPrice() (field.YieldRedemptionPrice, error) {
	var f field.YieldRedemptionPrice
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPriceType is a non-required field for ExecutionReport.
func (m ExecutionReport) YieldRedemptionPriceType() (field.YieldRedemptionPriceType, error) {
	var f field.YieldRedemptionPriceType
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

//InterestAtMaturity is a non-required field for ExecutionReport.
func (m ExecutionReport) InterestAtMaturity() (field.InterestAtMaturity, error) {
	var f field.InterestAtMaturity
	err := m.Body.Get(&f)
	return f, err
}

//EndAccruedInterestAmt is a non-required field for ExecutionReport.
func (m ExecutionReport) EndAccruedInterestAmt() (field.EndAccruedInterestAmt, error) {
	var f field.EndAccruedInterestAmt
	err := m.Body.Get(&f)
	return f, err
}

//StartCash is a non-required field for ExecutionReport.
func (m ExecutionReport) StartCash() (field.StartCash, error) {
	var f field.StartCash
	err := m.Body.Get(&f)
	return f, err
}

//EndCash is a non-required field for ExecutionReport.
func (m ExecutionReport) EndCash() (field.EndCash, error) {
	var f field.EndCash
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

//BookingType is a non-required field for ExecutionReport.
func (m ExecutionReport) BookingType() (field.BookingType, error) {
	var f field.BookingType
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

//SettlDate2 is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlDate2() (field.SettlDate2, error) {
	var f field.SettlDate2
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

//LastLiquidityInd is a non-required field for ExecutionReport.
func (m ExecutionReport) LastLiquidityInd() (field.LastLiquidityInd, error) {
	var f field.LastLiquidityInd
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

//CopyMsgIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) CopyMsgIndicator() (field.CopyMsgIndicator, error) {
	var f field.CopyMsgIndicator
	err := m.Body.Get(&f)
	return f, err
}

//NoMiscFees is a non-required field for ExecutionReport.
func (m ExecutionReport) NoMiscFees() (field.NoMiscFees, error) {
	var f field.NoMiscFees
	err := m.Body.Get(&f)
	return f, err
}

//NoStrategyParameters is a non-required field for ExecutionReport.
func (m ExecutionReport) NoStrategyParameters() (field.NoStrategyParameters, error) {
	var f field.NoStrategyParameters
	err := m.Body.Get(&f)
	return f, err
}

//HostCrossID is a non-required field for ExecutionReport.
func (m ExecutionReport) HostCrossID() (field.HostCrossID, error) {
	var f field.HostCrossID
	err := m.Body.Get(&f)
	return f, err
}

//ManualOrderIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) ManualOrderIndicator() (field.ManualOrderIndicator, error) {
	var f field.ManualOrderIndicator
	err := m.Body.Get(&f)
	return f, err
}

//CustDirectedOrder is a non-required field for ExecutionReport.
func (m ExecutionReport) CustDirectedOrder() (field.CustDirectedOrder, error) {
	var f field.CustDirectedOrder
	err := m.Body.Get(&f)
	return f, err
}

//ReceivedDeptID is a non-required field for ExecutionReport.
func (m ExecutionReport) ReceivedDeptID() (field.ReceivedDeptID, error) {
	var f field.ReceivedDeptID
	err := m.Body.Get(&f)
	return f, err
}

//CustOrderHandlingInst is a non-required field for ExecutionReport.
func (m ExecutionReport) CustOrderHandlingInst() (field.CustOrderHandlingInst, error) {
	var f field.CustOrderHandlingInst
	err := m.Body.Get(&f)
	return f, err
}

//OrderHandlingInstSource is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderHandlingInstSource() (field.OrderHandlingInstSource, error) {
	var f field.OrderHandlingInstSource
	err := m.Body.Get(&f)
	return f, err
}

//NoTrdRegTimestamps is a non-required field for ExecutionReport.
func (m ExecutionReport) NoTrdRegTimestamps() (field.NoTrdRegTimestamps, error) {
	var f field.NoTrdRegTimestamps
	err := m.Body.Get(&f)
	return f, err
}

//AggressorIndicator is a non-required field for ExecutionReport.
func (m ExecutionReport) AggressorIndicator() (field.AggressorIndicator, error) {
	var f field.AggressorIndicator
	err := m.Body.Get(&f)
	return f, err
}

//CalculatedCcyLastQty is a non-required field for ExecutionReport.
func (m ExecutionReport) CalculatedCcyLastQty() (field.CalculatedCcyLastQty, error) {
	var f field.CalculatedCcyLastQty
	err := m.Body.Get(&f)
	return f, err
}

//LastSwapPoints is a non-required field for ExecutionReport.
func (m ExecutionReport) LastSwapPoints() (field.LastSwapPoints, error) {
	var f field.LastSwapPoints
	err := m.Body.Get(&f)
	return f, err
}

//MatchType is a non-required field for ExecutionReport.
func (m ExecutionReport) MatchType() (field.MatchType, error) {
	var f field.MatchType
	err := m.Body.Get(&f)
	return f, err
}

//OrderCategory is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderCategory() (field.OrderCategory, error) {
	var f field.OrderCategory
	err := m.Body.Get(&f)
	return f, err
}

//LotType is a non-required field for ExecutionReport.
func (m ExecutionReport) LotType() (field.LotType, error) {
	var f field.LotType
	err := m.Body.Get(&f)
	return f, err
}

//PriceProtectionScope is a non-required field for ExecutionReport.
func (m ExecutionReport) PriceProtectionScope() (field.PriceProtectionScope, error) {
	var f field.PriceProtectionScope
	err := m.Body.Get(&f)
	return f, err
}

//TriggerType is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerType() (field.TriggerType, error) {
	var f field.TriggerType
	err := m.Body.Get(&f)
	return f, err
}

//TriggerAction is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerAction() (field.TriggerAction, error) {
	var f field.TriggerAction
	err := m.Body.Get(&f)
	return f, err
}

//TriggerPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerPrice() (field.TriggerPrice, error) {
	var f field.TriggerPrice
	err := m.Body.Get(&f)
	return f, err
}

//TriggerSymbol is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerSymbol() (field.TriggerSymbol, error) {
	var f field.TriggerSymbol
	err := m.Body.Get(&f)
	return f, err
}

//TriggerSecurityID is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerSecurityID() (field.TriggerSecurityID, error) {
	var f field.TriggerSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//TriggerSecurityIDSource is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerSecurityIDSource() (field.TriggerSecurityIDSource, error) {
	var f field.TriggerSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//TriggerSecurityDesc is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerSecurityDesc() (field.TriggerSecurityDesc, error) {
	var f field.TriggerSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//TriggerPriceType is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerPriceType() (field.TriggerPriceType, error) {
	var f field.TriggerPriceType
	err := m.Body.Get(&f)
	return f, err
}

//TriggerPriceTypeScope is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerPriceTypeScope() (field.TriggerPriceTypeScope, error) {
	var f field.TriggerPriceTypeScope
	err := m.Body.Get(&f)
	return f, err
}

//TriggerPriceDirection is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerPriceDirection() (field.TriggerPriceDirection, error) {
	var f field.TriggerPriceDirection
	err := m.Body.Get(&f)
	return f, err
}

//TriggerNewPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerNewPrice() (field.TriggerNewPrice, error) {
	var f field.TriggerNewPrice
	err := m.Body.Get(&f)
	return f, err
}

//TriggerOrderType is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerOrderType() (field.TriggerOrderType, error) {
	var f field.TriggerOrderType
	err := m.Body.Get(&f)
	return f, err
}

//TriggerNewQty is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerNewQty() (field.TriggerNewQty, error) {
	var f field.TriggerNewQty
	err := m.Body.Get(&f)
	return f, err
}

//TriggerTradingSessionID is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerTradingSessionID() (field.TriggerTradingSessionID, error) {
	var f field.TriggerTradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TriggerTradingSessionSubID is a non-required field for ExecutionReport.
func (m ExecutionReport) TriggerTradingSessionSubID() (field.TriggerTradingSessionSubID, error) {
	var f field.TriggerTradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//PeggedRefPrice is a non-required field for ExecutionReport.
func (m ExecutionReport) PeggedRefPrice() (field.PeggedRefPrice, error) {
	var f field.PeggedRefPrice
	err := m.Body.Get(&f)
	return f, err
}

//PreTradeAnonymity is a non-required field for ExecutionReport.
func (m ExecutionReport) PreTradeAnonymity() (field.PreTradeAnonymity, error) {
	var f field.PreTradeAnonymity
	err := m.Body.Get(&f)
	return f, err
}

//MatchIncrement is a non-required field for ExecutionReport.
func (m ExecutionReport) MatchIncrement() (field.MatchIncrement, error) {
	var f field.MatchIncrement
	err := m.Body.Get(&f)
	return f, err
}

//MaxPriceLevels is a non-required field for ExecutionReport.
func (m ExecutionReport) MaxPriceLevels() (field.MaxPriceLevels, error) {
	var f field.MaxPriceLevels
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryDisplayQty is a non-required field for ExecutionReport.
func (m ExecutionReport) SecondaryDisplayQty() (field.SecondaryDisplayQty, error) {
	var f field.SecondaryDisplayQty
	err := m.Body.Get(&f)
	return f, err
}

//DisplayWhen is a non-required field for ExecutionReport.
func (m ExecutionReport) DisplayWhen() (field.DisplayWhen, error) {
	var f field.DisplayWhen
	err := m.Body.Get(&f)
	return f, err
}

//DisplayMethod is a non-required field for ExecutionReport.
func (m ExecutionReport) DisplayMethod() (field.DisplayMethod, error) {
	var f field.DisplayMethod
	err := m.Body.Get(&f)
	return f, err
}

//DisplayLowQty is a non-required field for ExecutionReport.
func (m ExecutionReport) DisplayLowQty() (field.DisplayLowQty, error) {
	var f field.DisplayLowQty
	err := m.Body.Get(&f)
	return f, err
}

//DisplayHighQty is a non-required field for ExecutionReport.
func (m ExecutionReport) DisplayHighQty() (field.DisplayHighQty, error) {
	var f field.DisplayHighQty
	err := m.Body.Get(&f)
	return f, err
}

//DisplayMinIncr is a non-required field for ExecutionReport.
func (m ExecutionReport) DisplayMinIncr() (field.DisplayMinIncr, error) {
	var f field.DisplayMinIncr
	err := m.Body.Get(&f)
	return f, err
}

//RefreshQty is a non-required field for ExecutionReport.
func (m ExecutionReport) RefreshQty() (field.RefreshQty, error) {
	var f field.RefreshQty
	err := m.Body.Get(&f)
	return f, err
}

//DisplayQty is a non-required field for ExecutionReport.
func (m ExecutionReport) DisplayQty() (field.DisplayQty, error) {
	var f field.DisplayQty
	err := m.Body.Get(&f)
	return f, err
}

//Volatility is a non-required field for ExecutionReport.
func (m ExecutionReport) Volatility() (field.Volatility, error) {
	var f field.Volatility
	err := m.Body.Get(&f)
	return f, err
}

//TimeToExpiration is a non-required field for ExecutionReport.
func (m ExecutionReport) TimeToExpiration() (field.TimeToExpiration, error) {
	var f field.TimeToExpiration
	err := m.Body.Get(&f)
	return f, err
}

//RiskFreeRate is a non-required field for ExecutionReport.
func (m ExecutionReport) RiskFreeRate() (field.RiskFreeRate, error) {
	var f field.RiskFreeRate
	err := m.Body.Get(&f)
	return f, err
}

//PriceDelta is a non-required field for ExecutionReport.
func (m ExecutionReport) PriceDelta() (field.PriceDelta, error) {
	var f field.PriceDelta
	err := m.Body.Get(&f)
	return f, err
}

//TrdMatchID is a non-required field for ExecutionReport.
func (m ExecutionReport) TrdMatchID() (field.TrdMatchID, error) {
	var f field.TrdMatchID
	err := m.Body.Get(&f)
	return f, err
}

//AllocID is a non-required field for ExecutionReport.
func (m ExecutionReport) AllocID() (field.AllocID, error) {
	var f field.AllocID
	err := m.Body.Get(&f)
	return f, err
}

//NoAllocs is a non-required field for ExecutionReport.
func (m ExecutionReport) NoAllocs() (field.NoAllocs, error) {
	var f field.NoAllocs
	err := m.Body.Get(&f)
	return f, err
}

//TotNoFills is a non-required field for ExecutionReport.
func (m ExecutionReport) TotNoFills() (field.TotNoFills, error) {
	var f field.TotNoFills
	err := m.Body.Get(&f)
	return f, err
}

//LastFragment is a non-required field for ExecutionReport.
func (m ExecutionReport) LastFragment() (field.LastFragment, error) {
	var f field.LastFragment
	err := m.Body.Get(&f)
	return f, err
}

//NoFills is a non-required field for ExecutionReport.
func (m ExecutionReport) NoFills() (field.NoFills, error) {
	var f field.NoFills
	err := m.Body.Get(&f)
	return f, err
}

//DividendYield is a non-required field for ExecutionReport.
func (m ExecutionReport) DividendYield() (field.DividendYield, error) {
	var f field.DividendYield
	err := m.Body.Get(&f)
	return f, err
}

//ApplID is a non-required field for ExecutionReport.
func (m ExecutionReport) ApplID() (field.ApplID, error) {
	var f field.ApplID
	err := m.Body.Get(&f)
	return f, err
}

//ApplSeqNum is a non-required field for ExecutionReport.
func (m ExecutionReport) ApplSeqNum() (field.ApplSeqNum, error) {
	var f field.ApplSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplLastSeqNum is a non-required field for ExecutionReport.
func (m ExecutionReport) ApplLastSeqNum() (field.ApplLastSeqNum, error) {
	var f field.ApplLastSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplResendFlag is a non-required field for ExecutionReport.
func (m ExecutionReport) ApplResendFlag() (field.ApplResendFlag, error) {
	var f field.ApplResendFlag
	err := m.Body.Get(&f)
	return f, err
}
