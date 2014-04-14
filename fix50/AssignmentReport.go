package fix50

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//AssignmentReport msg type = AW.
type AssignmentReport struct {
	message.Message
}

//AssignmentReportBuilder builds AssignmentReport messages.
type AssignmentReportBuilder struct {
	message.MessageBuilder
}

//NewAssignmentReportBuilder returns an initialized AssignmentReportBuilder with specified required fields.
func NewAssignmentReportBuilder(
	asgnrptid field.AsgnRptID,
	clearingbusinessdate field.ClearingBusinessDate) *AssignmentReportBuilder {
	builder := new(AssignmentReportBuilder)
	builder.Body.Set(asgnrptid)
	builder.Body.Set(clearingbusinessdate)
	return builder
}

//AsgnRptID is a required field for AssignmentReport.
func (m *AssignmentReport) AsgnRptID() (*field.AsgnRptID, error) {
	f := new(field.AsgnRptID)
	err := m.Body.Get(f)
	return f, err
}

//TotNumAssignmentReports is a non-required field for AssignmentReport.
func (m *AssignmentReport) TotNumAssignmentReports() (*field.TotNumAssignmentReports, error) {
	f := new(field.TotNumAssignmentReports)
	err := m.Body.Get(f)
	return f, err
}

//LastRptRequested is a non-required field for AssignmentReport.
func (m *AssignmentReport) LastRptRequested() (*field.LastRptRequested, error) {
	f := new(field.LastRptRequested)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyIDs is a non-required field for AssignmentReport.
func (m *AssignmentReport) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//Account is a non-required field for AssignmentReport.
func (m *AssignmentReport) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//AccountType is a non-required field for AssignmentReport.
func (m *AssignmentReport) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for AssignmentReport.
func (m *AssignmentReport) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for AssignmentReport.
func (m *AssignmentReport) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for AssignmentReport.
func (m *AssignmentReport) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for AssignmentReport.
func (m *AssignmentReport) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for AssignmentReport.
func (m *AssignmentReport) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for AssignmentReport.
func (m *AssignmentReport) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for AssignmentReport.
func (m *AssignmentReport) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for AssignmentReport.
func (m *AssignmentReport) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//SecuritySubType is a non-required field for AssignmentReport.
func (m *AssignmentReport) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for AssignmentReport.
func (m *AssignmentReport) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for AssignmentReport.
func (m *AssignmentReport) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for AssignmentReport.
func (m *AssignmentReport) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for AssignmentReport.
func (m *AssignmentReport) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for AssignmentReport.
func (m *AssignmentReport) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for AssignmentReport.
func (m *AssignmentReport) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for AssignmentReport.
func (m *AssignmentReport) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for AssignmentReport.
func (m *AssignmentReport) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for AssignmentReport.
func (m *AssignmentReport) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for AssignmentReport.
func (m *AssignmentReport) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for AssignmentReport.
func (m *AssignmentReport) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for AssignmentReport.
func (m *AssignmentReport) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for AssignmentReport.
func (m *AssignmentReport) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for AssignmentReport.
func (m *AssignmentReport) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for AssignmentReport.
func (m *AssignmentReport) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//StrikeCurrency is a non-required field for AssignmentReport.
func (m *AssignmentReport) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for AssignmentReport.
func (m *AssignmentReport) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for AssignmentReport.
func (m *AssignmentReport) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for AssignmentReport.
func (m *AssignmentReport) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for AssignmentReport.
func (m *AssignmentReport) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for AssignmentReport.
func (m *AssignmentReport) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for AssignmentReport.
func (m *AssignmentReport) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for AssignmentReport.
func (m *AssignmentReport) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for AssignmentReport.
func (m *AssignmentReport) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for AssignmentReport.
func (m *AssignmentReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for AssignmentReport.
func (m *AssignmentReport) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Pool is a non-required field for AssignmentReport.
func (m *AssignmentReport) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//ContractSettlMonth is a non-required field for AssignmentReport.
func (m *AssignmentReport) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//CPProgram is a non-required field for AssignmentReport.
func (m *AssignmentReport) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//CPRegType is a non-required field for AssignmentReport.
func (m *AssignmentReport) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//NoEvents is a non-required field for AssignmentReport.
func (m *AssignmentReport) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//DatedDate is a non-required field for AssignmentReport.
func (m *AssignmentReport) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//InterestAccrualDate is a non-required field for AssignmentReport.
func (m *AssignmentReport) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityStatus is a non-required field for AssignmentReport.
func (m *AssignmentReport) SecurityStatus() (*field.SecurityStatus, error) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for AssignmentReport.
func (m *AssignmentReport) SettleOnOpenFlag() (*field.SettleOnOpenFlag, error) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for AssignmentReport.
func (m *AssignmentReport) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, error) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//StrikeMultiplier is a non-required field for AssignmentReport.
func (m *AssignmentReport) StrikeMultiplier() (*field.StrikeMultiplier, error) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//StrikeValue is a non-required field for AssignmentReport.
func (m *AssignmentReport) StrikeValue() (*field.StrikeValue, error) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrement is a non-required field for AssignmentReport.
func (m *AssignmentReport) MinPriceIncrement() (*field.MinPriceIncrement, error) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//PositionLimit is a non-required field for AssignmentReport.
func (m *AssignmentReport) PositionLimit() (*field.PositionLimit, error) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NTPositionLimit is a non-required field for AssignmentReport.
func (m *AssignmentReport) NTPositionLimit() (*field.NTPositionLimit, error) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NoInstrumentParties is a non-required field for AssignmentReport.
func (m *AssignmentReport) NoInstrumentParties() (*field.NoInstrumentParties, error) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasure is a non-required field for AssignmentReport.
func (m *AssignmentReport) UnitOfMeasure() (*field.UnitOfMeasure, error) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//TimeUnit is a non-required field for AssignmentReport.
func (m *AssignmentReport) TimeUnit() (*field.TimeUnit, error) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//MaturityTime is a non-required field for AssignmentReport.
func (m *AssignmentReport) MaturityTime() (*field.MaturityTime, error) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for AssignmentReport.
func (m *AssignmentReport) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//NoLegs is a non-required field for AssignmentReport.
func (m *AssignmentReport) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyings is a non-required field for AssignmentReport.
func (m *AssignmentReport) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//NoPositions is a non-required field for AssignmentReport.
func (m *AssignmentReport) NoPositions() (*field.NoPositions, error) {
	f := new(field.NoPositions)
	err := m.Body.Get(f)
	return f, err
}

//NoPosAmt is a non-required field for AssignmentReport.
func (m *AssignmentReport) NoPosAmt() (*field.NoPosAmt, error) {
	f := new(field.NoPosAmt)
	err := m.Body.Get(f)
	return f, err
}

//ThresholdAmount is a non-required field for AssignmentReport.
func (m *AssignmentReport) ThresholdAmount() (*field.ThresholdAmount, error) {
	f := new(field.ThresholdAmount)
	err := m.Body.Get(f)
	return f, err
}

//SettlPrice is a non-required field for AssignmentReport.
func (m *AssignmentReport) SettlPrice() (*field.SettlPrice, error) {
	f := new(field.SettlPrice)
	err := m.Body.Get(f)
	return f, err
}

//SettlPriceType is a non-required field for AssignmentReport.
func (m *AssignmentReport) SettlPriceType() (*field.SettlPriceType, error) {
	f := new(field.SettlPriceType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSettlPrice is a non-required field for AssignmentReport.
func (m *AssignmentReport) UnderlyingSettlPrice() (*field.UnderlyingSettlPrice, error) {
	f := new(field.UnderlyingSettlPrice)
	err := m.Body.Get(f)
	return f, err
}

//ExpireDate is a non-required field for AssignmentReport.
func (m *AssignmentReport) ExpireDate() (*field.ExpireDate, error) {
	f := new(field.ExpireDate)
	err := m.Body.Get(f)
	return f, err
}

//AssignmentMethod is a non-required field for AssignmentReport.
func (m *AssignmentReport) AssignmentMethod() (*field.AssignmentMethod, error) {
	f := new(field.AssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//AssignmentUnit is a non-required field for AssignmentReport.
func (m *AssignmentReport) AssignmentUnit() (*field.AssignmentUnit, error) {
	f := new(field.AssignmentUnit)
	err := m.Body.Get(f)
	return f, err
}

//OpenInterest is a non-required field for AssignmentReport.
func (m *AssignmentReport) OpenInterest() (*field.OpenInterest, error) {
	f := new(field.OpenInterest)
	err := m.Body.Get(f)
	return f, err
}

//ExerciseMethod is a non-required field for AssignmentReport.
func (m *AssignmentReport) ExerciseMethod() (*field.ExerciseMethod, error) {
	f := new(field.ExerciseMethod)
	err := m.Body.Get(f)
	return f, err
}

//SettlSessID is a non-required field for AssignmentReport.
func (m *AssignmentReport) SettlSessID() (*field.SettlSessID, error) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}

//SettlSessSubID is a non-required field for AssignmentReport.
func (m *AssignmentReport) SettlSessSubID() (*field.SettlSessSubID, error) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}

//ClearingBusinessDate is a required field for AssignmentReport.
func (m *AssignmentReport) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for AssignmentReport.
func (m *AssignmentReport) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for AssignmentReport.
func (m *AssignmentReport) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for AssignmentReport.
func (m *AssignmentReport) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//PriorSettlPrice is a non-required field for AssignmentReport.
func (m *AssignmentReport) PriorSettlPrice() (*field.PriorSettlPrice, error) {
	f := new(field.PriorSettlPrice)
	err := m.Body.Get(f)
	return f, err
}
