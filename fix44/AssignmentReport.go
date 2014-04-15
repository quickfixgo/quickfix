package fix44

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

//CreateAssignmentReportBuilder returns an initialized AssignmentReportBuilder with specified required fields.
func CreateAssignmentReportBuilder(
	asgnrptid field.AsgnRptID,
	accounttype field.AccountType,
	settlprice field.SettlPrice,
	settlpricetype field.SettlPriceType,
	underlyingsettlprice field.UnderlyingSettlPrice,
	assignmentmethod field.AssignmentMethod,
	openinterest field.OpenInterest,
	exercisemethod field.ExerciseMethod,
	settlsessid field.SettlSessID,
	settlsesssubid field.SettlSessSubID,
	clearingbusinessdate field.ClearingBusinessDate) AssignmentReportBuilder {
	var builder AssignmentReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(asgnrptid)
	builder.Body.Set(accounttype)
	builder.Body.Set(settlprice)
	builder.Body.Set(settlpricetype)
	builder.Body.Set(underlyingsettlprice)
	builder.Body.Set(assignmentmethod)
	builder.Body.Set(openinterest)
	builder.Body.Set(exercisemethod)
	builder.Body.Set(settlsessid)
	builder.Body.Set(settlsesssubid)
	builder.Body.Set(clearingbusinessdate)
	return builder
}

//AsgnRptID is a required field for AssignmentReport.
func (m AssignmentReport) AsgnRptID() (field.AsgnRptID, error) {
	var f field.AsgnRptID
	err := m.Body.Get(&f)
	return f, err
}

//TotNumAssignmentReports is a non-required field for AssignmentReport.
func (m AssignmentReport) TotNumAssignmentReports() (field.TotNumAssignmentReports, error) {
	var f field.TotNumAssignmentReports
	err := m.Body.Get(&f)
	return f, err
}

//LastRptRequested is a non-required field for AssignmentReport.
func (m AssignmentReport) LastRptRequested() (field.LastRptRequested, error) {
	var f field.LastRptRequested
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for AssignmentReport.
func (m AssignmentReport) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for AssignmentReport.
func (m AssignmentReport) Account() (field.Account, error) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a required field for AssignmentReport.
func (m AssignmentReport) AccountType() (field.AccountType, error) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for AssignmentReport.
func (m AssignmentReport) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for AssignmentReport.
func (m AssignmentReport) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for AssignmentReport.
func (m AssignmentReport) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for AssignmentReport.
func (m AssignmentReport) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for AssignmentReport.
func (m AssignmentReport) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for AssignmentReport.
func (m AssignmentReport) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for AssignmentReport.
func (m AssignmentReport) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for AssignmentReport.
func (m AssignmentReport) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for AssignmentReport.
func (m AssignmentReport) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for AssignmentReport.
func (m AssignmentReport) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for AssignmentReport.
func (m AssignmentReport) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for AssignmentReport.
func (m AssignmentReport) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for AssignmentReport.
func (m AssignmentReport) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for AssignmentReport.
func (m AssignmentReport) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for AssignmentReport.
func (m AssignmentReport) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for AssignmentReport.
func (m AssignmentReport) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for AssignmentReport.
func (m AssignmentReport) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for AssignmentReport.
func (m AssignmentReport) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for AssignmentReport.
func (m AssignmentReport) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for AssignmentReport.
func (m AssignmentReport) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for AssignmentReport.
func (m AssignmentReport) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for AssignmentReport.
func (m AssignmentReport) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for AssignmentReport.
func (m AssignmentReport) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for AssignmentReport.
func (m AssignmentReport) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for AssignmentReport.
func (m AssignmentReport) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for AssignmentReport.
func (m AssignmentReport) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for AssignmentReport.
func (m AssignmentReport) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for AssignmentReport.
func (m AssignmentReport) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for AssignmentReport.
func (m AssignmentReport) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for AssignmentReport.
func (m AssignmentReport) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for AssignmentReport.
func (m AssignmentReport) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for AssignmentReport.
func (m AssignmentReport) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for AssignmentReport.
func (m AssignmentReport) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for AssignmentReport.
func (m AssignmentReport) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for AssignmentReport.
func (m AssignmentReport) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for AssignmentReport.
func (m AssignmentReport) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for AssignmentReport.
func (m AssignmentReport) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for AssignmentReport.
func (m AssignmentReport) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for AssignmentReport.
func (m AssignmentReport) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for AssignmentReport.
func (m AssignmentReport) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for AssignmentReport.
func (m AssignmentReport) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for AssignmentReport.
func (m AssignmentReport) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for AssignmentReport.
func (m AssignmentReport) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for AssignmentReport.
func (m AssignmentReport) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for AssignmentReport.
func (m AssignmentReport) NoUnderlyings() (field.NoUnderlyings, error) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//NoPositions is a non-required field for AssignmentReport.
func (m AssignmentReport) NoPositions() (field.NoPositions, error) {
	var f field.NoPositions
	err := m.Body.Get(&f)
	return f, err
}

//NoPosAmt is a non-required field for AssignmentReport.
func (m AssignmentReport) NoPosAmt() (field.NoPosAmt, error) {
	var f field.NoPosAmt
	err := m.Body.Get(&f)
	return f, err
}

//ThresholdAmount is a non-required field for AssignmentReport.
func (m AssignmentReport) ThresholdAmount() (field.ThresholdAmount, error) {
	var f field.ThresholdAmount
	err := m.Body.Get(&f)
	return f, err
}

//SettlPrice is a required field for AssignmentReport.
func (m AssignmentReport) SettlPrice() (field.SettlPrice, error) {
	var f field.SettlPrice
	err := m.Body.Get(&f)
	return f, err
}

//SettlPriceType is a required field for AssignmentReport.
func (m AssignmentReport) SettlPriceType() (field.SettlPriceType, error) {
	var f field.SettlPriceType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSettlPrice is a required field for AssignmentReport.
func (m AssignmentReport) UnderlyingSettlPrice() (field.UnderlyingSettlPrice, error) {
	var f field.UnderlyingSettlPrice
	err := m.Body.Get(&f)
	return f, err
}

//ExpireDate is a non-required field for AssignmentReport.
func (m AssignmentReport) ExpireDate() (field.ExpireDate, error) {
	var f field.ExpireDate
	err := m.Body.Get(&f)
	return f, err
}

//AssignmentMethod is a required field for AssignmentReport.
func (m AssignmentReport) AssignmentMethod() (field.AssignmentMethod, error) {
	var f field.AssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//AssignmentUnit is a non-required field for AssignmentReport.
func (m AssignmentReport) AssignmentUnit() (field.AssignmentUnit, error) {
	var f field.AssignmentUnit
	err := m.Body.Get(&f)
	return f, err
}

//OpenInterest is a required field for AssignmentReport.
func (m AssignmentReport) OpenInterest() (field.OpenInterest, error) {
	var f field.OpenInterest
	err := m.Body.Get(&f)
	return f, err
}

//ExerciseMethod is a required field for AssignmentReport.
func (m AssignmentReport) ExerciseMethod() (field.ExerciseMethod, error) {
	var f field.ExerciseMethod
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessID is a required field for AssignmentReport.
func (m AssignmentReport) SettlSessID() (field.SettlSessID, error) {
	var f field.SettlSessID
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessSubID is a required field for AssignmentReport.
func (m AssignmentReport) SettlSessSubID() (field.SettlSessSubID, error) {
	var f field.SettlSessSubID
	err := m.Body.Get(&f)
	return f, err
}

//ClearingBusinessDate is a required field for AssignmentReport.
func (m AssignmentReport) ClearingBusinessDate() (field.ClearingBusinessDate, error) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for AssignmentReport.
func (m AssignmentReport) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for AssignmentReport.
func (m AssignmentReport) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for AssignmentReport.
func (m AssignmentReport) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
