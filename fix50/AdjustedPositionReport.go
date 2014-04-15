package fix50

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//AdjustedPositionReport msg type = BL.
type AdjustedPositionReport struct {
	message.Message
}

//AdjustedPositionReportBuilder builds AdjustedPositionReport messages.
type AdjustedPositionReportBuilder struct {
	message.MessageBuilder
}

//CreateAdjustedPositionReportBuilder returns an initialized AdjustedPositionReportBuilder with specified required fields.
func CreateAdjustedPositionReportBuilder(
	posmaintrptid field.PosMaintRptID,
	clearingbusinessdate field.ClearingBusinessDate) AdjustedPositionReportBuilder {
	var builder AdjustedPositionReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(posmaintrptid)
	builder.Body.Set(clearingbusinessdate)
	return builder
}

//PosMaintRptID is a required field for AdjustedPositionReport.
func (m AdjustedPositionReport) PosMaintRptID() (field.PosMaintRptID, error) {
	var f field.PosMaintRptID
	err := m.Body.Get(&f)
	return f, err
}

//PosReqType is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) PosReqType() (field.PosReqType, error) {
	var f field.PosReqType
	err := m.Body.Get(&f)
	return f, err
}

//ClearingBusinessDate is a required field for AdjustedPositionReport.
func (m AdjustedPositionReport) ClearingBusinessDate() (field.ClearingBusinessDate, error) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessID is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) SettlSessID() (field.SettlSessID, error) {
	var f field.SettlSessID
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//NoPositions is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) NoPositions() (field.NoPositions, error) {
	var f field.NoPositions
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) SecurityStatus() (field.SecurityStatus, error) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) SettleOnOpenFlag() (field.SettleOnOpenFlag, error) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, error) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) StrikeMultiplier() (field.StrikeMultiplier, error) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) StrikeValue() (field.StrikeValue, error) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) MinPriceIncrement() (field.MinPriceIncrement, error) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) PositionLimit() (field.PositionLimit, error) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) NTPositionLimit() (field.NTPositionLimit, error) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) NoInstrumentParties() (field.NoInstrumentParties, error) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) UnitOfMeasure() (field.UnitOfMeasure, error) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) TimeUnit() (field.TimeUnit, error) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) MaturityTime() (field.MaturityTime, error) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//SettlPrice is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) SettlPrice() (field.SettlPrice, error) {
	var f field.SettlPrice
	err := m.Body.Get(&f)
	return f, err
}

//PriorSettlPrice is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) PriorSettlPrice() (field.PriorSettlPrice, error) {
	var f field.PriorSettlPrice
	err := m.Body.Get(&f)
	return f, err
}

//PosMaintRptRefID is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) PosMaintRptRefID() (field.PosMaintRptRefID, error) {
	var f field.PosMaintRptRefID
	err := m.Body.Get(&f)
	return f, err
}
