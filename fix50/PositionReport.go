package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//PositionReport msg type = AP.
type PositionReport struct {
	message.Message
}

//PositionReportBuilder builds PositionReport messages.
type PositionReportBuilder struct {
	message.MessageBuilder
}

//CreatePositionReportBuilder returns an initialized PositionReportBuilder with specified required fields.
func CreatePositionReportBuilder(
	posmaintrptid field.PosMaintRptID,
	clearingbusinessdate field.ClearingBusinessDate) PositionReportBuilder {
	var builder PositionReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("AP"))
	builder.Body.Set(posmaintrptid)
	builder.Body.Set(clearingbusinessdate)
	return builder
}

//PosMaintRptID is a required field for PositionReport.
func (m PositionReport) PosMaintRptID() (*field.PosMaintRptID, errors.MessageRejectError) {
	f := new(field.PosMaintRptID)
	err := m.Body.Get(f)
	return f, err
}

//GetPosMaintRptID reads a PosMaintRptID from PositionReport.
func (m PositionReport) GetPosMaintRptID(f *field.PosMaintRptID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosReqID is a non-required field for PositionReport.
func (m PositionReport) PosReqID() (*field.PosReqID, errors.MessageRejectError) {
	f := new(field.PosReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetPosReqID reads a PosReqID from PositionReport.
func (m PositionReport) GetPosReqID(f *field.PosReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosReqType is a non-required field for PositionReport.
func (m PositionReport) PosReqType() (*field.PosReqType, errors.MessageRejectError) {
	f := new(field.PosReqType)
	err := m.Body.Get(f)
	return f, err
}

//GetPosReqType reads a PosReqType from PositionReport.
func (m PositionReport) GetPosReqType(f *field.PosReqType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a non-required field for PositionReport.
func (m PositionReport) SubscriptionRequestType() (*field.SubscriptionRequestType, errors.MessageRejectError) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from PositionReport.
func (m PositionReport) GetSubscriptionRequestType(f *field.SubscriptionRequestType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalNumPosReports is a non-required field for PositionReport.
func (m PositionReport) TotalNumPosReports() (*field.TotalNumPosReports, errors.MessageRejectError) {
	f := new(field.TotalNumPosReports)
	err := m.Body.Get(f)
	return f, err
}

//GetTotalNumPosReports reads a TotalNumPosReports from PositionReport.
func (m PositionReport) GetTotalNumPosReports(f *field.TotalNumPosReports) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnsolicitedIndicator is a non-required field for PositionReport.
func (m PositionReport) UnsolicitedIndicator() (*field.UnsolicitedIndicator, errors.MessageRejectError) {
	f := new(field.UnsolicitedIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetUnsolicitedIndicator reads a UnsolicitedIndicator from PositionReport.
func (m PositionReport) GetUnsolicitedIndicator(f *field.UnsolicitedIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosReqResult is a non-required field for PositionReport.
func (m PositionReport) PosReqResult() (*field.PosReqResult, errors.MessageRejectError) {
	f := new(field.PosReqResult)
	err := m.Body.Get(f)
	return f, err
}

//GetPosReqResult reads a PosReqResult from PositionReport.
func (m PositionReport) GetPosReqResult(f *field.PosReqResult) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a required field for PositionReport.
func (m PositionReport) ClearingBusinessDate() (*field.ClearingBusinessDate, errors.MessageRejectError) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from PositionReport.
func (m PositionReport) GetClearingBusinessDate(f *field.ClearingBusinessDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessID is a non-required field for PositionReport.
func (m PositionReport) SettlSessID() (*field.SettlSessID, errors.MessageRejectError) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessID reads a SettlSessID from PositionReport.
func (m PositionReport) GetSettlSessID(f *field.SettlSessID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessSubID is a non-required field for PositionReport.
func (m PositionReport) SettlSessSubID() (*field.SettlSessSubID, errors.MessageRejectError) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessSubID reads a SettlSessSubID from PositionReport.
func (m PositionReport) GetSettlSessSubID(f *field.SettlSessSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for PositionReport.
func (m PositionReport) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from PositionReport.
func (m PositionReport) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for PositionReport.
func (m PositionReport) Account() (*field.Account, errors.MessageRejectError) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from PositionReport.
func (m PositionReport) GetAccount(f *field.Account) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for PositionReport.
func (m PositionReport) AcctIDSource() (*field.AcctIDSource, errors.MessageRejectError) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from PositionReport.
func (m PositionReport) GetAcctIDSource(f *field.AcctIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for PositionReport.
func (m PositionReport) AccountType() (*field.AccountType, errors.MessageRejectError) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from PositionReport.
func (m PositionReport) GetAccountType(f *field.AccountType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for PositionReport.
func (m PositionReport) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from PositionReport.
func (m PositionReport) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for PositionReport.
func (m PositionReport) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from PositionReport.
func (m PositionReport) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for PositionReport.
func (m PositionReport) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from PositionReport.
func (m PositionReport) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for PositionReport.
func (m PositionReport) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from PositionReport.
func (m PositionReport) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for PositionReport.
func (m PositionReport) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from PositionReport.
func (m PositionReport) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for PositionReport.
func (m PositionReport) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from PositionReport.
func (m PositionReport) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for PositionReport.
func (m PositionReport) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from PositionReport.
func (m PositionReport) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for PositionReport.
func (m PositionReport) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from PositionReport.
func (m PositionReport) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for PositionReport.
func (m PositionReport) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from PositionReport.
func (m PositionReport) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for PositionReport.
func (m PositionReport) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from PositionReport.
func (m PositionReport) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for PositionReport.
func (m PositionReport) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from PositionReport.
func (m PositionReport) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for PositionReport.
func (m PositionReport) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from PositionReport.
func (m PositionReport) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for PositionReport.
func (m PositionReport) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from PositionReport.
func (m PositionReport) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for PositionReport.
func (m PositionReport) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from PositionReport.
func (m PositionReport) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for PositionReport.
func (m PositionReport) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from PositionReport.
func (m PositionReport) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for PositionReport.
func (m PositionReport) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from PositionReport.
func (m PositionReport) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for PositionReport.
func (m PositionReport) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from PositionReport.
func (m PositionReport) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for PositionReport.
func (m PositionReport) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from PositionReport.
func (m PositionReport) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for PositionReport.
func (m PositionReport) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from PositionReport.
func (m PositionReport) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for PositionReport.
func (m PositionReport) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from PositionReport.
func (m PositionReport) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for PositionReport.
func (m PositionReport) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from PositionReport.
func (m PositionReport) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for PositionReport.
func (m PositionReport) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from PositionReport.
func (m PositionReport) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for PositionReport.
func (m PositionReport) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from PositionReport.
func (m PositionReport) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for PositionReport.
func (m PositionReport) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from PositionReport.
func (m PositionReport) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for PositionReport.
func (m PositionReport) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from PositionReport.
func (m PositionReport) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for PositionReport.
func (m PositionReport) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from PositionReport.
func (m PositionReport) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for PositionReport.
func (m PositionReport) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from PositionReport.
func (m PositionReport) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for PositionReport.
func (m PositionReport) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from PositionReport.
func (m PositionReport) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for PositionReport.
func (m PositionReport) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from PositionReport.
func (m PositionReport) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for PositionReport.
func (m PositionReport) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from PositionReport.
func (m PositionReport) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for PositionReport.
func (m PositionReport) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from PositionReport.
func (m PositionReport) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for PositionReport.
func (m PositionReport) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from PositionReport.
func (m PositionReport) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for PositionReport.
func (m PositionReport) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from PositionReport.
func (m PositionReport) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for PositionReport.
func (m PositionReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from PositionReport.
func (m PositionReport) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for PositionReport.
func (m PositionReport) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from PositionReport.
func (m PositionReport) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for PositionReport.
func (m PositionReport) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from PositionReport.
func (m PositionReport) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for PositionReport.
func (m PositionReport) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from PositionReport.
func (m PositionReport) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for PositionReport.
func (m PositionReport) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from PositionReport.
func (m PositionReport) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for PositionReport.
func (m PositionReport) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from PositionReport.
func (m PositionReport) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for PositionReport.
func (m PositionReport) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from PositionReport.
func (m PositionReport) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for PositionReport.
func (m PositionReport) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from PositionReport.
func (m PositionReport) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for PositionReport.
func (m PositionReport) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from PositionReport.
func (m PositionReport) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for PositionReport.
func (m PositionReport) SecurityStatus() (*field.SecurityStatus, errors.MessageRejectError) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from PositionReport.
func (m PositionReport) GetSecurityStatus(f *field.SecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for PositionReport.
func (m PositionReport) SettleOnOpenFlag() (*field.SettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from PositionReport.
func (m PositionReport) GetSettleOnOpenFlag(f *field.SettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for PositionReport.
func (m PositionReport) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from PositionReport.
func (m PositionReport) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for PositionReport.
func (m PositionReport) StrikeMultiplier() (*field.StrikeMultiplier, errors.MessageRejectError) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from PositionReport.
func (m PositionReport) GetStrikeMultiplier(f *field.StrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for PositionReport.
func (m PositionReport) StrikeValue() (*field.StrikeValue, errors.MessageRejectError) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from PositionReport.
func (m PositionReport) GetStrikeValue(f *field.StrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for PositionReport.
func (m PositionReport) MinPriceIncrement() (*field.MinPriceIncrement, errors.MessageRejectError) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from PositionReport.
func (m PositionReport) GetMinPriceIncrement(f *field.MinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for PositionReport.
func (m PositionReport) PositionLimit() (*field.PositionLimit, errors.MessageRejectError) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from PositionReport.
func (m PositionReport) GetPositionLimit(f *field.PositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for PositionReport.
func (m PositionReport) NTPositionLimit() (*field.NTPositionLimit, errors.MessageRejectError) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from PositionReport.
func (m PositionReport) GetNTPositionLimit(f *field.NTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for PositionReport.
func (m PositionReport) NoInstrumentParties() (*field.NoInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from PositionReport.
func (m PositionReport) GetNoInstrumentParties(f *field.NoInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for PositionReport.
func (m PositionReport) UnitOfMeasure() (*field.UnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from PositionReport.
func (m PositionReport) GetUnitOfMeasure(f *field.UnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for PositionReport.
func (m PositionReport) TimeUnit() (*field.TimeUnit, errors.MessageRejectError) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from PositionReport.
func (m PositionReport) GetTimeUnit(f *field.TimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for PositionReport.
func (m PositionReport) MaturityTime() (*field.MaturityTime, errors.MessageRejectError) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from PositionReport.
func (m PositionReport) GetMaturityTime(f *field.MaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for PositionReport.
func (m PositionReport) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from PositionReport.
func (m PositionReport) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlPrice is a non-required field for PositionReport.
func (m PositionReport) SettlPrice() (*field.SettlPrice, errors.MessageRejectError) {
	f := new(field.SettlPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlPrice reads a SettlPrice from PositionReport.
func (m PositionReport) GetSettlPrice(f *field.SettlPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlPriceType is a non-required field for PositionReport.
func (m PositionReport) SettlPriceType() (*field.SettlPriceType, errors.MessageRejectError) {
	f := new(field.SettlPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlPriceType reads a SettlPriceType from PositionReport.
func (m PositionReport) GetSettlPriceType(f *field.SettlPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriorSettlPrice is a non-required field for PositionReport.
func (m PositionReport) PriorSettlPrice() (*field.PriorSettlPrice, errors.MessageRejectError) {
	f := new(field.PriorSettlPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetPriorSettlPrice reads a PriorSettlPrice from PositionReport.
func (m PositionReport) GetPriorSettlPrice(f *field.PriorSettlPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for PositionReport.
func (m PositionReport) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from PositionReport.
func (m PositionReport) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for PositionReport.
func (m PositionReport) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from PositionReport.
func (m PositionReport) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPositions is a non-required field for PositionReport.
func (m PositionReport) NoPositions() (*field.NoPositions, errors.MessageRejectError) {
	f := new(field.NoPositions)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPositions reads a NoPositions from PositionReport.
func (m PositionReport) GetNoPositions(f *field.NoPositions) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPosAmt is a non-required field for PositionReport.
func (m PositionReport) NoPosAmt() (*field.NoPosAmt, errors.MessageRejectError) {
	f := new(field.NoPosAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPosAmt reads a NoPosAmt from PositionReport.
func (m PositionReport) GetNoPosAmt(f *field.NoPosAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistStatus is a non-required field for PositionReport.
func (m PositionReport) RegistStatus() (*field.RegistStatus, errors.MessageRejectError) {
	f := new(field.RegistStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetRegistStatus reads a RegistStatus from PositionReport.
func (m PositionReport) GetRegistStatus(f *field.RegistStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryDate is a non-required field for PositionReport.
func (m PositionReport) DeliveryDate() (*field.DeliveryDate, errors.MessageRejectError) {
	f := new(field.DeliveryDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryDate reads a DeliveryDate from PositionReport.
func (m PositionReport) GetDeliveryDate(f *field.DeliveryDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for PositionReport.
func (m PositionReport) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from PositionReport.
func (m PositionReport) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for PositionReport.
func (m PositionReport) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from PositionReport.
func (m PositionReport) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for PositionReport.
func (m PositionReport) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from PositionReport.
func (m PositionReport) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchStatus is a non-required field for PositionReport.
func (m PositionReport) MatchStatus() (*field.MatchStatus, errors.MessageRejectError) {
	f := new(field.MatchStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetMatchStatus reads a MatchStatus from PositionReport.
func (m PositionReport) GetMatchStatus(f *field.MatchStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for PositionReport.
func (m PositionReport) PriceType() (*field.PriceType, errors.MessageRejectError) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from PositionReport.
func (m PositionReport) GetPriceType(f *field.PriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for PositionReport.
func (m PositionReport) SettlCurrency() (*field.SettlCurrency, errors.MessageRejectError) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from PositionReport.
func (m PositionReport) GetSettlCurrency(f *field.SettlCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageEventSource is a non-required field for PositionReport.
func (m PositionReport) MessageEventSource() (*field.MessageEventSource, errors.MessageRejectError) {
	f := new(field.MessageEventSource)
	err := m.Body.Get(f)
	return f, err
}

//GetMessageEventSource reads a MessageEventSource from PositionReport.
func (m PositionReport) GetMessageEventSource(f *field.MessageEventSource) errors.MessageRejectError {
	return m.Body.Get(f)
}
