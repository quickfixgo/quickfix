package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
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
	clearingbusinessdate field.ClearingBusinessDate) AssignmentReportBuilder {
	var builder AssignmentReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("AW"))
	builder.Body.Set(asgnrptid)
	builder.Body.Set(clearingbusinessdate)
	return builder
}

//AsgnRptID is a required field for AssignmentReport.
func (m AssignmentReport) AsgnRptID() (*field.AsgnRptID, errors.MessageRejectError) {
	f := new(field.AsgnRptID)
	err := m.Body.Get(f)
	return f, err
}

//GetAsgnRptID reads a AsgnRptID from AssignmentReport.
func (m AssignmentReport) GetAsgnRptID(f *field.AsgnRptID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNumAssignmentReports is a non-required field for AssignmentReport.
func (m AssignmentReport) TotNumAssignmentReports() (*field.TotNumAssignmentReports, errors.MessageRejectError) {
	f := new(field.TotNumAssignmentReports)
	err := m.Body.Get(f)
	return f, err
}

//GetTotNumAssignmentReports reads a TotNumAssignmentReports from AssignmentReport.
func (m AssignmentReport) GetTotNumAssignmentReports(f *field.TotNumAssignmentReports) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastRptRequested is a non-required field for AssignmentReport.
func (m AssignmentReport) LastRptRequested() (*field.LastRptRequested, errors.MessageRejectError) {
	f := new(field.LastRptRequested)
	err := m.Body.Get(f)
	return f, err
}

//GetLastRptRequested reads a LastRptRequested from AssignmentReport.
func (m AssignmentReport) GetLastRptRequested(f *field.LastRptRequested) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for AssignmentReport.
func (m AssignmentReport) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from AssignmentReport.
func (m AssignmentReport) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for AssignmentReport.
func (m AssignmentReport) Account() (*field.Account, errors.MessageRejectError) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from AssignmentReport.
func (m AssignmentReport) GetAccount(f *field.Account) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for AssignmentReport.
func (m AssignmentReport) AccountType() (*field.AccountType, errors.MessageRejectError) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from AssignmentReport.
func (m AssignmentReport) GetAccountType(f *field.AccountType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for AssignmentReport.
func (m AssignmentReport) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from AssignmentReport.
func (m AssignmentReport) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for AssignmentReport.
func (m AssignmentReport) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from AssignmentReport.
func (m AssignmentReport) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for AssignmentReport.
func (m AssignmentReport) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from AssignmentReport.
func (m AssignmentReport) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for AssignmentReport.
func (m AssignmentReport) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from AssignmentReport.
func (m AssignmentReport) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for AssignmentReport.
func (m AssignmentReport) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from AssignmentReport.
func (m AssignmentReport) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for AssignmentReport.
func (m AssignmentReport) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from AssignmentReport.
func (m AssignmentReport) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for AssignmentReport.
func (m AssignmentReport) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from AssignmentReport.
func (m AssignmentReport) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for AssignmentReport.
func (m AssignmentReport) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from AssignmentReport.
func (m AssignmentReport) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for AssignmentReport.
func (m AssignmentReport) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from AssignmentReport.
func (m AssignmentReport) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for AssignmentReport.
func (m AssignmentReport) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from AssignmentReport.
func (m AssignmentReport) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for AssignmentReport.
func (m AssignmentReport) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from AssignmentReport.
func (m AssignmentReport) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for AssignmentReport.
func (m AssignmentReport) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from AssignmentReport.
func (m AssignmentReport) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for AssignmentReport.
func (m AssignmentReport) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from AssignmentReport.
func (m AssignmentReport) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for AssignmentReport.
func (m AssignmentReport) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from AssignmentReport.
func (m AssignmentReport) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for AssignmentReport.
func (m AssignmentReport) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from AssignmentReport.
func (m AssignmentReport) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for AssignmentReport.
func (m AssignmentReport) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from AssignmentReport.
func (m AssignmentReport) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for AssignmentReport.
func (m AssignmentReport) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from AssignmentReport.
func (m AssignmentReport) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for AssignmentReport.
func (m AssignmentReport) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from AssignmentReport.
func (m AssignmentReport) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for AssignmentReport.
func (m AssignmentReport) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from AssignmentReport.
func (m AssignmentReport) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for AssignmentReport.
func (m AssignmentReport) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from AssignmentReport.
func (m AssignmentReport) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for AssignmentReport.
func (m AssignmentReport) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from AssignmentReport.
func (m AssignmentReport) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for AssignmentReport.
func (m AssignmentReport) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from AssignmentReport.
func (m AssignmentReport) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for AssignmentReport.
func (m AssignmentReport) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from AssignmentReport.
func (m AssignmentReport) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for AssignmentReport.
func (m AssignmentReport) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from AssignmentReport.
func (m AssignmentReport) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for AssignmentReport.
func (m AssignmentReport) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from AssignmentReport.
func (m AssignmentReport) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for AssignmentReport.
func (m AssignmentReport) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from AssignmentReport.
func (m AssignmentReport) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for AssignmentReport.
func (m AssignmentReport) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from AssignmentReport.
func (m AssignmentReport) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for AssignmentReport.
func (m AssignmentReport) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from AssignmentReport.
func (m AssignmentReport) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for AssignmentReport.
func (m AssignmentReport) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from AssignmentReport.
func (m AssignmentReport) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for AssignmentReport.
func (m AssignmentReport) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from AssignmentReport.
func (m AssignmentReport) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for AssignmentReport.
func (m AssignmentReport) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from AssignmentReport.
func (m AssignmentReport) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for AssignmentReport.
func (m AssignmentReport) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from AssignmentReport.
func (m AssignmentReport) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for AssignmentReport.
func (m AssignmentReport) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from AssignmentReport.
func (m AssignmentReport) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for AssignmentReport.
func (m AssignmentReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from AssignmentReport.
func (m AssignmentReport) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for AssignmentReport.
func (m AssignmentReport) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from AssignmentReport.
func (m AssignmentReport) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for AssignmentReport.
func (m AssignmentReport) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from AssignmentReport.
func (m AssignmentReport) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for AssignmentReport.
func (m AssignmentReport) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from AssignmentReport.
func (m AssignmentReport) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for AssignmentReport.
func (m AssignmentReport) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from AssignmentReport.
func (m AssignmentReport) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for AssignmentReport.
func (m AssignmentReport) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from AssignmentReport.
func (m AssignmentReport) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for AssignmentReport.
func (m AssignmentReport) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from AssignmentReport.
func (m AssignmentReport) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for AssignmentReport.
func (m AssignmentReport) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from AssignmentReport.
func (m AssignmentReport) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for AssignmentReport.
func (m AssignmentReport) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from AssignmentReport.
func (m AssignmentReport) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for AssignmentReport.
func (m AssignmentReport) SecurityStatus() (*field.SecurityStatus, errors.MessageRejectError) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from AssignmentReport.
func (m AssignmentReport) GetSecurityStatus(f *field.SecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for AssignmentReport.
func (m AssignmentReport) SettleOnOpenFlag() (*field.SettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from AssignmentReport.
func (m AssignmentReport) GetSettleOnOpenFlag(f *field.SettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for AssignmentReport.
func (m AssignmentReport) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from AssignmentReport.
func (m AssignmentReport) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for AssignmentReport.
func (m AssignmentReport) StrikeMultiplier() (*field.StrikeMultiplier, errors.MessageRejectError) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from AssignmentReport.
func (m AssignmentReport) GetStrikeMultiplier(f *field.StrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for AssignmentReport.
func (m AssignmentReport) StrikeValue() (*field.StrikeValue, errors.MessageRejectError) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from AssignmentReport.
func (m AssignmentReport) GetStrikeValue(f *field.StrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for AssignmentReport.
func (m AssignmentReport) MinPriceIncrement() (*field.MinPriceIncrement, errors.MessageRejectError) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from AssignmentReport.
func (m AssignmentReport) GetMinPriceIncrement(f *field.MinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for AssignmentReport.
func (m AssignmentReport) PositionLimit() (*field.PositionLimit, errors.MessageRejectError) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from AssignmentReport.
func (m AssignmentReport) GetPositionLimit(f *field.PositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for AssignmentReport.
func (m AssignmentReport) NTPositionLimit() (*field.NTPositionLimit, errors.MessageRejectError) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from AssignmentReport.
func (m AssignmentReport) GetNTPositionLimit(f *field.NTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for AssignmentReport.
func (m AssignmentReport) NoInstrumentParties() (*field.NoInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from AssignmentReport.
func (m AssignmentReport) GetNoInstrumentParties(f *field.NoInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for AssignmentReport.
func (m AssignmentReport) UnitOfMeasure() (*field.UnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from AssignmentReport.
func (m AssignmentReport) GetUnitOfMeasure(f *field.UnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for AssignmentReport.
func (m AssignmentReport) TimeUnit() (*field.TimeUnit, errors.MessageRejectError) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from AssignmentReport.
func (m AssignmentReport) GetTimeUnit(f *field.TimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for AssignmentReport.
func (m AssignmentReport) MaturityTime() (*field.MaturityTime, errors.MessageRejectError) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from AssignmentReport.
func (m AssignmentReport) GetMaturityTime(f *field.MaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for AssignmentReport.
func (m AssignmentReport) SecurityGroup() (*field.SecurityGroup, errors.MessageRejectError) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from AssignmentReport.
func (m AssignmentReport) GetSecurityGroup(f *field.SecurityGroup) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for AssignmentReport.
func (m AssignmentReport) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, errors.MessageRejectError) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from AssignmentReport.
func (m AssignmentReport) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for AssignmentReport.
func (m AssignmentReport) UnitOfMeasureQty() (*field.UnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from AssignmentReport.
func (m AssignmentReport) GetUnitOfMeasureQty(f *field.UnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for AssignmentReport.
func (m AssignmentReport) SecurityXMLLen() (*field.SecurityXMLLen, errors.MessageRejectError) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from AssignmentReport.
func (m AssignmentReport) GetSecurityXMLLen(f *field.SecurityXMLLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for AssignmentReport.
func (m AssignmentReport) SecurityXML() (*field.SecurityXML, errors.MessageRejectError) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from AssignmentReport.
func (m AssignmentReport) GetSecurityXML(f *field.SecurityXML) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for AssignmentReport.
func (m AssignmentReport) SecurityXMLSchema() (*field.SecurityXMLSchema, errors.MessageRejectError) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from AssignmentReport.
func (m AssignmentReport) GetSecurityXMLSchema(f *field.SecurityXMLSchema) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for AssignmentReport.
func (m AssignmentReport) ProductComplex() (*field.ProductComplex, errors.MessageRejectError) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from AssignmentReport.
func (m AssignmentReport) GetProductComplex(f *field.ProductComplex) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for AssignmentReport.
func (m AssignmentReport) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from AssignmentReport.
func (m AssignmentReport) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for AssignmentReport.
func (m AssignmentReport) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from AssignmentReport.
func (m AssignmentReport) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for AssignmentReport.
func (m AssignmentReport) SettlMethod() (*field.SettlMethod, errors.MessageRejectError) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from AssignmentReport.
func (m AssignmentReport) GetSettlMethod(f *field.SettlMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for AssignmentReport.
func (m AssignmentReport) ExerciseStyle() (*field.ExerciseStyle, errors.MessageRejectError) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from AssignmentReport.
func (m AssignmentReport) GetExerciseStyle(f *field.ExerciseStyle) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutAmount is a non-required field for AssignmentReport.
func (m AssignmentReport) OptPayoutAmount() (*field.OptPayoutAmount, errors.MessageRejectError) {
	f := new(field.OptPayoutAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutAmount reads a OptPayoutAmount from AssignmentReport.
func (m AssignmentReport) GetOptPayoutAmount(f *field.OptPayoutAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for AssignmentReport.
func (m AssignmentReport) PriceQuoteMethod() (*field.PriceQuoteMethod, errors.MessageRejectError) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from AssignmentReport.
func (m AssignmentReport) GetPriceQuoteMethod(f *field.PriceQuoteMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for AssignmentReport.
func (m AssignmentReport) ListMethod() (*field.ListMethod, errors.MessageRejectError) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from AssignmentReport.
func (m AssignmentReport) GetListMethod(f *field.ListMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for AssignmentReport.
func (m AssignmentReport) CapPrice() (*field.CapPrice, errors.MessageRejectError) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from AssignmentReport.
func (m AssignmentReport) GetCapPrice(f *field.CapPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for AssignmentReport.
func (m AssignmentReport) FloorPrice() (*field.FloorPrice, errors.MessageRejectError) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from AssignmentReport.
func (m AssignmentReport) GetFloorPrice(f *field.FloorPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for AssignmentReport.
func (m AssignmentReport) PutOrCall() (*field.PutOrCall, errors.MessageRejectError) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from AssignmentReport.
func (m AssignmentReport) GetPutOrCall(f *field.PutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for AssignmentReport.
func (m AssignmentReport) FlexibleIndicator() (*field.FlexibleIndicator, errors.MessageRejectError) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from AssignmentReport.
func (m AssignmentReport) GetFlexibleIndicator(f *field.FlexibleIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for AssignmentReport.
func (m AssignmentReport) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, errors.MessageRejectError) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from AssignmentReport.
func (m AssignmentReport) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ValuationMethod is a non-required field for AssignmentReport.
func (m AssignmentReport) ValuationMethod() (*field.ValuationMethod, errors.MessageRejectError) {
	f := new(field.ValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetValuationMethod reads a ValuationMethod from AssignmentReport.
func (m AssignmentReport) GetValuationMethod(f *field.ValuationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplierUnit is a non-required field for AssignmentReport.
func (m AssignmentReport) ContractMultiplierUnit() (*field.ContractMultiplierUnit, errors.MessageRejectError) {
	f := new(field.ContractMultiplierUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplierUnit reads a ContractMultiplierUnit from AssignmentReport.
func (m AssignmentReport) GetContractMultiplierUnit(f *field.ContractMultiplierUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlowScheduleType is a non-required field for AssignmentReport.
func (m AssignmentReport) FlowScheduleType() (*field.FlowScheduleType, errors.MessageRejectError) {
	f := new(field.FlowScheduleType)
	err := m.Body.Get(f)
	return f, err
}

//GetFlowScheduleType reads a FlowScheduleType from AssignmentReport.
func (m AssignmentReport) GetFlowScheduleType(f *field.FlowScheduleType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RestructuringType is a non-required field for AssignmentReport.
func (m AssignmentReport) RestructuringType() (*field.RestructuringType, errors.MessageRejectError) {
	f := new(field.RestructuringType)
	err := m.Body.Get(f)
	return f, err
}

//GetRestructuringType reads a RestructuringType from AssignmentReport.
func (m AssignmentReport) GetRestructuringType(f *field.RestructuringType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Seniority is a non-required field for AssignmentReport.
func (m AssignmentReport) Seniority() (*field.Seniority, errors.MessageRejectError) {
	f := new(field.Seniority)
	err := m.Body.Get(f)
	return f, err
}

//GetSeniority reads a Seniority from AssignmentReport.
func (m AssignmentReport) GetSeniority(f *field.Seniority) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NotionalPercentageOutstanding is a non-required field for AssignmentReport.
func (m AssignmentReport) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstanding, errors.MessageRejectError) {
	f := new(field.NotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//GetNotionalPercentageOutstanding reads a NotionalPercentageOutstanding from AssignmentReport.
func (m AssignmentReport) GetNotionalPercentageOutstanding(f *field.NotionalPercentageOutstanding) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OriginalNotionalPercentageOutstanding is a non-required field for AssignmentReport.
func (m AssignmentReport) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstanding, errors.MessageRejectError) {
	f := new(field.OriginalNotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//GetOriginalNotionalPercentageOutstanding reads a OriginalNotionalPercentageOutstanding from AssignmentReport.
func (m AssignmentReport) GetOriginalNotionalPercentageOutstanding(f *field.OriginalNotionalPercentageOutstanding) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AttachmentPoint is a non-required field for AssignmentReport.
func (m AssignmentReport) AttachmentPoint() (*field.AttachmentPoint, errors.MessageRejectError) {
	f := new(field.AttachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//GetAttachmentPoint reads a AttachmentPoint from AssignmentReport.
func (m AssignmentReport) GetAttachmentPoint(f *field.AttachmentPoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DetachmentPoint is a non-required field for AssignmentReport.
func (m AssignmentReport) DetachmentPoint() (*field.DetachmentPoint, errors.MessageRejectError) {
	f := new(field.DetachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//GetDetachmentPoint reads a DetachmentPoint from AssignmentReport.
func (m AssignmentReport) GetDetachmentPoint(f *field.DetachmentPoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceDeterminationMethod is a non-required field for AssignmentReport.
func (m AssignmentReport) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethod, errors.MessageRejectError) {
	f := new(field.StrikePriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceDeterminationMethod reads a StrikePriceDeterminationMethod from AssignmentReport.
func (m AssignmentReport) GetStrikePriceDeterminationMethod(f *field.StrikePriceDeterminationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryMethod is a non-required field for AssignmentReport.
func (m AssignmentReport) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethod, errors.MessageRejectError) {
	f := new(field.StrikePriceBoundaryMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryMethod reads a StrikePriceBoundaryMethod from AssignmentReport.
func (m AssignmentReport) GetStrikePriceBoundaryMethod(f *field.StrikePriceBoundaryMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryPrecision is a non-required field for AssignmentReport.
func (m AssignmentReport) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecision, errors.MessageRejectError) {
	f := new(field.StrikePriceBoundaryPrecision)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryPrecision reads a StrikePriceBoundaryPrecision from AssignmentReport.
func (m AssignmentReport) GetStrikePriceBoundaryPrecision(f *field.StrikePriceBoundaryPrecision) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceDeterminationMethod is a non-required field for AssignmentReport.
func (m AssignmentReport) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethod, errors.MessageRejectError) {
	f := new(field.UnderlyingPriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceDeterminationMethod reads a UnderlyingPriceDeterminationMethod from AssignmentReport.
func (m AssignmentReport) GetUnderlyingPriceDeterminationMethod(f *field.UnderlyingPriceDeterminationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutType is a non-required field for AssignmentReport.
func (m AssignmentReport) OptPayoutType() (*field.OptPayoutType, errors.MessageRejectError) {
	f := new(field.OptPayoutType)
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutType reads a OptPayoutType from AssignmentReport.
func (m AssignmentReport) GetOptPayoutType(f *field.OptPayoutType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoComplexEvents is a non-required field for AssignmentReport.
func (m AssignmentReport) NoComplexEvents() (*field.NoComplexEvents, errors.MessageRejectError) {
	f := new(field.NoComplexEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoComplexEvents reads a NoComplexEvents from AssignmentReport.
func (m AssignmentReport) GetNoComplexEvents(f *field.NoComplexEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for AssignmentReport.
func (m AssignmentReport) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from AssignmentReport.
func (m AssignmentReport) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for AssignmentReport.
func (m AssignmentReport) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from AssignmentReport.
func (m AssignmentReport) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for AssignmentReport.
func (m AssignmentReport) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from AssignmentReport.
func (m AssignmentReport) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPositions is a non-required field for AssignmentReport.
func (m AssignmentReport) NoPositions() (*field.NoPositions, errors.MessageRejectError) {
	f := new(field.NoPositions)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPositions reads a NoPositions from AssignmentReport.
func (m AssignmentReport) GetNoPositions(f *field.NoPositions) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPosAmt is a non-required field for AssignmentReport.
func (m AssignmentReport) NoPosAmt() (*field.NoPosAmt, errors.MessageRejectError) {
	f := new(field.NoPosAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPosAmt reads a NoPosAmt from AssignmentReport.
func (m AssignmentReport) GetNoPosAmt(f *field.NoPosAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ThresholdAmount is a non-required field for AssignmentReport.
func (m AssignmentReport) ThresholdAmount() (*field.ThresholdAmount, errors.MessageRejectError) {
	f := new(field.ThresholdAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetThresholdAmount reads a ThresholdAmount from AssignmentReport.
func (m AssignmentReport) GetThresholdAmount(f *field.ThresholdAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlPrice is a non-required field for AssignmentReport.
func (m AssignmentReport) SettlPrice() (*field.SettlPrice, errors.MessageRejectError) {
	f := new(field.SettlPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlPrice reads a SettlPrice from AssignmentReport.
func (m AssignmentReport) GetSettlPrice(f *field.SettlPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlPriceType is a non-required field for AssignmentReport.
func (m AssignmentReport) SettlPriceType() (*field.SettlPriceType, errors.MessageRejectError) {
	f := new(field.SettlPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlPriceType reads a SettlPriceType from AssignmentReport.
func (m AssignmentReport) GetSettlPriceType(f *field.SettlPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSettlPrice is a non-required field for AssignmentReport.
func (m AssignmentReport) UnderlyingSettlPrice() (*field.UnderlyingSettlPrice, errors.MessageRejectError) {
	f := new(field.UnderlyingSettlPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSettlPrice reads a UnderlyingSettlPrice from AssignmentReport.
func (m AssignmentReport) GetUnderlyingSettlPrice(f *field.UnderlyingSettlPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireDate is a non-required field for AssignmentReport.
func (m AssignmentReport) ExpireDate() (*field.ExpireDate, errors.MessageRejectError) {
	f := new(field.ExpireDate)
	err := m.Body.Get(f)
	return f, err
}

//GetExpireDate reads a ExpireDate from AssignmentReport.
func (m AssignmentReport) GetExpireDate(f *field.ExpireDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AssignmentMethod is a non-required field for AssignmentReport.
func (m AssignmentReport) AssignmentMethod() (*field.AssignmentMethod, errors.MessageRejectError) {
	f := new(field.AssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetAssignmentMethod reads a AssignmentMethod from AssignmentReport.
func (m AssignmentReport) GetAssignmentMethod(f *field.AssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AssignmentUnit is a non-required field for AssignmentReport.
func (m AssignmentReport) AssignmentUnit() (*field.AssignmentUnit, errors.MessageRejectError) {
	f := new(field.AssignmentUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetAssignmentUnit reads a AssignmentUnit from AssignmentReport.
func (m AssignmentReport) GetAssignmentUnit(f *field.AssignmentUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OpenInterest is a non-required field for AssignmentReport.
func (m AssignmentReport) OpenInterest() (*field.OpenInterest, errors.MessageRejectError) {
	f := new(field.OpenInterest)
	err := m.Body.Get(f)
	return f, err
}

//GetOpenInterest reads a OpenInterest from AssignmentReport.
func (m AssignmentReport) GetOpenInterest(f *field.OpenInterest) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseMethod is a non-required field for AssignmentReport.
func (m AssignmentReport) ExerciseMethod() (*field.ExerciseMethod, errors.MessageRejectError) {
	f := new(field.ExerciseMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseMethod reads a ExerciseMethod from AssignmentReport.
func (m AssignmentReport) GetExerciseMethod(f *field.ExerciseMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessID is a non-required field for AssignmentReport.
func (m AssignmentReport) SettlSessID() (*field.SettlSessID, errors.MessageRejectError) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessID reads a SettlSessID from AssignmentReport.
func (m AssignmentReport) GetSettlSessID(f *field.SettlSessID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessSubID is a non-required field for AssignmentReport.
func (m AssignmentReport) SettlSessSubID() (*field.SettlSessSubID, errors.MessageRejectError) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessSubID reads a SettlSessSubID from AssignmentReport.
func (m AssignmentReport) GetSettlSessSubID(f *field.SettlSessSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a required field for AssignmentReport.
func (m AssignmentReport) ClearingBusinessDate() (*field.ClearingBusinessDate, errors.MessageRejectError) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from AssignmentReport.
func (m AssignmentReport) GetClearingBusinessDate(f *field.ClearingBusinessDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for AssignmentReport.
func (m AssignmentReport) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from AssignmentReport.
func (m AssignmentReport) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for AssignmentReport.
func (m AssignmentReport) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from AssignmentReport.
func (m AssignmentReport) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for AssignmentReport.
func (m AssignmentReport) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from AssignmentReport.
func (m AssignmentReport) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriorSettlPrice is a non-required field for AssignmentReport.
func (m AssignmentReport) PriorSettlPrice() (*field.PriorSettlPrice, errors.MessageRejectError) {
	f := new(field.PriorSettlPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetPriorSettlPrice reads a PriorSettlPrice from AssignmentReport.
func (m AssignmentReport) GetPriorSettlPrice(f *field.PriorSettlPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for AssignmentReport.
func (m AssignmentReport) ApplID() (*field.ApplID, errors.MessageRejectError) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from AssignmentReport.
func (m AssignmentReport) GetApplID(f *field.ApplID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for AssignmentReport.
func (m AssignmentReport) ApplSeqNum() (*field.ApplSeqNum, errors.MessageRejectError) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from AssignmentReport.
func (m AssignmentReport) GetApplSeqNum(f *field.ApplSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for AssignmentReport.
func (m AssignmentReport) ApplLastSeqNum() (*field.ApplLastSeqNum, errors.MessageRejectError) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from AssignmentReport.
func (m AssignmentReport) GetApplLastSeqNum(f *field.ApplLastSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for AssignmentReport.
func (m AssignmentReport) ApplResendFlag() (*field.ApplResendFlag, errors.MessageRejectError) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from AssignmentReport.
func (m AssignmentReport) GetApplResendFlag(f *field.ApplResendFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosReqID is a non-required field for AssignmentReport.
func (m AssignmentReport) PosReqID() (*field.PosReqID, errors.MessageRejectError) {
	f := new(field.PosReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetPosReqID reads a PosReqID from AssignmentReport.
func (m AssignmentReport) GetPosReqID(f *field.PosReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}
