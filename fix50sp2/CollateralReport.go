package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//CollateralReport msg type = BA.
type CollateralReport struct {
	message.Message
}

//CollateralReportBuilder builds CollateralReport messages.
type CollateralReportBuilder struct {
	message.MessageBuilder
}

//CreateCollateralReportBuilder returns an initialized CollateralReportBuilder with specified required fields.
func CreateCollateralReportBuilder(
	collrptid field.CollRptID,
	collstatus field.CollStatus) CollateralReportBuilder {
	var builder CollateralReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("BA"))
	builder.Body.Set(collrptid)
	builder.Body.Set(collstatus)
	return builder
}

//CollRptID is a required field for CollateralReport.
func (m CollateralReport) CollRptID() (*field.CollRptID, errors.MessageRejectError) {
	f := new(field.CollRptID)
	err := m.Body.Get(f)
	return f, err
}

//GetCollRptID reads a CollRptID from CollateralReport.
func (m CollateralReport) GetCollRptID(f *field.CollRptID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CollInquiryID is a non-required field for CollateralReport.
func (m CollateralReport) CollInquiryID() (*field.CollInquiryID, errors.MessageRejectError) {
	f := new(field.CollInquiryID)
	err := m.Body.Get(f)
	return f, err
}

//GetCollInquiryID reads a CollInquiryID from CollateralReport.
func (m CollateralReport) GetCollInquiryID(f *field.CollInquiryID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CollStatus is a required field for CollateralReport.
func (m CollateralReport) CollStatus() (*field.CollStatus, errors.MessageRejectError) {
	f := new(field.CollStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetCollStatus reads a CollStatus from CollateralReport.
func (m CollateralReport) GetCollStatus(f *field.CollStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNumReports is a non-required field for CollateralReport.
func (m CollateralReport) TotNumReports() (*field.TotNumReports, errors.MessageRejectError) {
	f := new(field.TotNumReports)
	err := m.Body.Get(f)
	return f, err
}

//GetTotNumReports reads a TotNumReports from CollateralReport.
func (m CollateralReport) GetTotNumReports(f *field.TotNumReports) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastRptRequested is a non-required field for CollateralReport.
func (m CollateralReport) LastRptRequested() (*field.LastRptRequested, errors.MessageRejectError) {
	f := new(field.LastRptRequested)
	err := m.Body.Get(f)
	return f, err
}

//GetLastRptRequested reads a LastRptRequested from CollateralReport.
func (m CollateralReport) GetLastRptRequested(f *field.LastRptRequested) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for CollateralReport.
func (m CollateralReport) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from CollateralReport.
func (m CollateralReport) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for CollateralReport.
func (m CollateralReport) Account() (*field.Account, errors.MessageRejectError) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from CollateralReport.
func (m CollateralReport) GetAccount(f *field.Account) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for CollateralReport.
func (m CollateralReport) AccountType() (*field.AccountType, errors.MessageRejectError) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from CollateralReport.
func (m CollateralReport) GetAccountType(f *field.AccountType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for CollateralReport.
func (m CollateralReport) ClOrdID() (*field.ClOrdID, errors.MessageRejectError) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from CollateralReport.
func (m CollateralReport) GetClOrdID(f *field.ClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderID is a non-required field for CollateralReport.
func (m CollateralReport) OrderID() (*field.OrderID, errors.MessageRejectError) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from CollateralReport.
func (m CollateralReport) GetOrderID(f *field.OrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryOrderID is a non-required field for CollateralReport.
func (m CollateralReport) SecondaryOrderID() (*field.SecondaryOrderID, errors.MessageRejectError) {
	f := new(field.SecondaryOrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryOrderID reads a SecondaryOrderID from CollateralReport.
func (m CollateralReport) GetSecondaryOrderID(f *field.SecondaryOrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for CollateralReport.
func (m CollateralReport) SecondaryClOrdID() (*field.SecondaryClOrdID, errors.MessageRejectError) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from CollateralReport.
func (m CollateralReport) GetSecondaryClOrdID(f *field.SecondaryClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoExecs is a non-required field for CollateralReport.
func (m CollateralReport) NoExecs() (*field.NoExecs, errors.MessageRejectError) {
	f := new(field.NoExecs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoExecs reads a NoExecs from CollateralReport.
func (m CollateralReport) GetNoExecs(f *field.NoExecs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTrades is a non-required field for CollateralReport.
func (m CollateralReport) NoTrades() (*field.NoTrades, errors.MessageRejectError) {
	f := new(field.NoTrades)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTrades reads a NoTrades from CollateralReport.
func (m CollateralReport) GetNoTrades(f *field.NoTrades) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for CollateralReport.
func (m CollateralReport) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from CollateralReport.
func (m CollateralReport) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for CollateralReport.
func (m CollateralReport) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from CollateralReport.
func (m CollateralReport) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for CollateralReport.
func (m CollateralReport) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from CollateralReport.
func (m CollateralReport) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for CollateralReport.
func (m CollateralReport) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from CollateralReport.
func (m CollateralReport) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for CollateralReport.
func (m CollateralReport) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from CollateralReport.
func (m CollateralReport) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for CollateralReport.
func (m CollateralReport) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from CollateralReport.
func (m CollateralReport) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for CollateralReport.
func (m CollateralReport) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from CollateralReport.
func (m CollateralReport) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for CollateralReport.
func (m CollateralReport) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from CollateralReport.
func (m CollateralReport) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for CollateralReport.
func (m CollateralReport) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from CollateralReport.
func (m CollateralReport) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for CollateralReport.
func (m CollateralReport) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from CollateralReport.
func (m CollateralReport) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for CollateralReport.
func (m CollateralReport) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from CollateralReport.
func (m CollateralReport) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for CollateralReport.
func (m CollateralReport) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from CollateralReport.
func (m CollateralReport) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for CollateralReport.
func (m CollateralReport) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from CollateralReport.
func (m CollateralReport) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for CollateralReport.
func (m CollateralReport) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from CollateralReport.
func (m CollateralReport) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for CollateralReport.
func (m CollateralReport) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from CollateralReport.
func (m CollateralReport) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for CollateralReport.
func (m CollateralReport) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from CollateralReport.
func (m CollateralReport) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for CollateralReport.
func (m CollateralReport) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from CollateralReport.
func (m CollateralReport) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for CollateralReport.
func (m CollateralReport) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from CollateralReport.
func (m CollateralReport) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for CollateralReport.
func (m CollateralReport) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from CollateralReport.
func (m CollateralReport) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for CollateralReport.
func (m CollateralReport) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from CollateralReport.
func (m CollateralReport) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for CollateralReport.
func (m CollateralReport) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from CollateralReport.
func (m CollateralReport) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for CollateralReport.
func (m CollateralReport) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from CollateralReport.
func (m CollateralReport) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for CollateralReport.
func (m CollateralReport) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from CollateralReport.
func (m CollateralReport) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for CollateralReport.
func (m CollateralReport) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from CollateralReport.
func (m CollateralReport) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for CollateralReport.
func (m CollateralReport) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from CollateralReport.
func (m CollateralReport) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for CollateralReport.
func (m CollateralReport) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from CollateralReport.
func (m CollateralReport) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for CollateralReport.
func (m CollateralReport) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from CollateralReport.
func (m CollateralReport) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for CollateralReport.
func (m CollateralReport) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from CollateralReport.
func (m CollateralReport) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for CollateralReport.
func (m CollateralReport) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from CollateralReport.
func (m CollateralReport) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for CollateralReport.
func (m CollateralReport) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from CollateralReport.
func (m CollateralReport) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for CollateralReport.
func (m CollateralReport) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from CollateralReport.
func (m CollateralReport) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for CollateralReport.
func (m CollateralReport) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from CollateralReport.
func (m CollateralReport) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for CollateralReport.
func (m CollateralReport) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from CollateralReport.
func (m CollateralReport) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for CollateralReport.
func (m CollateralReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from CollateralReport.
func (m CollateralReport) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for CollateralReport.
func (m CollateralReport) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from CollateralReport.
func (m CollateralReport) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for CollateralReport.
func (m CollateralReport) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from CollateralReport.
func (m CollateralReport) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for CollateralReport.
func (m CollateralReport) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from CollateralReport.
func (m CollateralReport) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for CollateralReport.
func (m CollateralReport) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from CollateralReport.
func (m CollateralReport) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for CollateralReport.
func (m CollateralReport) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from CollateralReport.
func (m CollateralReport) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for CollateralReport.
func (m CollateralReport) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from CollateralReport.
func (m CollateralReport) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for CollateralReport.
func (m CollateralReport) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from CollateralReport.
func (m CollateralReport) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for CollateralReport.
func (m CollateralReport) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from CollateralReport.
func (m CollateralReport) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for CollateralReport.
func (m CollateralReport) SecurityStatus() (*field.SecurityStatus, errors.MessageRejectError) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from CollateralReport.
func (m CollateralReport) GetSecurityStatus(f *field.SecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for CollateralReport.
func (m CollateralReport) SettleOnOpenFlag() (*field.SettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from CollateralReport.
func (m CollateralReport) GetSettleOnOpenFlag(f *field.SettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for CollateralReport.
func (m CollateralReport) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from CollateralReport.
func (m CollateralReport) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for CollateralReport.
func (m CollateralReport) StrikeMultiplier() (*field.StrikeMultiplier, errors.MessageRejectError) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from CollateralReport.
func (m CollateralReport) GetStrikeMultiplier(f *field.StrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for CollateralReport.
func (m CollateralReport) StrikeValue() (*field.StrikeValue, errors.MessageRejectError) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from CollateralReport.
func (m CollateralReport) GetStrikeValue(f *field.StrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for CollateralReport.
func (m CollateralReport) MinPriceIncrement() (*field.MinPriceIncrement, errors.MessageRejectError) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from CollateralReport.
func (m CollateralReport) GetMinPriceIncrement(f *field.MinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for CollateralReport.
func (m CollateralReport) PositionLimit() (*field.PositionLimit, errors.MessageRejectError) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from CollateralReport.
func (m CollateralReport) GetPositionLimit(f *field.PositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for CollateralReport.
func (m CollateralReport) NTPositionLimit() (*field.NTPositionLimit, errors.MessageRejectError) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from CollateralReport.
func (m CollateralReport) GetNTPositionLimit(f *field.NTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for CollateralReport.
func (m CollateralReport) NoInstrumentParties() (*field.NoInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from CollateralReport.
func (m CollateralReport) GetNoInstrumentParties(f *field.NoInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for CollateralReport.
func (m CollateralReport) UnitOfMeasure() (*field.UnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from CollateralReport.
func (m CollateralReport) GetUnitOfMeasure(f *field.UnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for CollateralReport.
func (m CollateralReport) TimeUnit() (*field.TimeUnit, errors.MessageRejectError) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from CollateralReport.
func (m CollateralReport) GetTimeUnit(f *field.TimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for CollateralReport.
func (m CollateralReport) MaturityTime() (*field.MaturityTime, errors.MessageRejectError) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from CollateralReport.
func (m CollateralReport) GetMaturityTime(f *field.MaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for CollateralReport.
func (m CollateralReport) SecurityGroup() (*field.SecurityGroup, errors.MessageRejectError) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from CollateralReport.
func (m CollateralReport) GetSecurityGroup(f *field.SecurityGroup) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for CollateralReport.
func (m CollateralReport) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, errors.MessageRejectError) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from CollateralReport.
func (m CollateralReport) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for CollateralReport.
func (m CollateralReport) UnitOfMeasureQty() (*field.UnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from CollateralReport.
func (m CollateralReport) GetUnitOfMeasureQty(f *field.UnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for CollateralReport.
func (m CollateralReport) SecurityXMLLen() (*field.SecurityXMLLen, errors.MessageRejectError) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from CollateralReport.
func (m CollateralReport) GetSecurityXMLLen(f *field.SecurityXMLLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for CollateralReport.
func (m CollateralReport) SecurityXML() (*field.SecurityXML, errors.MessageRejectError) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from CollateralReport.
func (m CollateralReport) GetSecurityXML(f *field.SecurityXML) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for CollateralReport.
func (m CollateralReport) SecurityXMLSchema() (*field.SecurityXMLSchema, errors.MessageRejectError) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from CollateralReport.
func (m CollateralReport) GetSecurityXMLSchema(f *field.SecurityXMLSchema) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for CollateralReport.
func (m CollateralReport) ProductComplex() (*field.ProductComplex, errors.MessageRejectError) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from CollateralReport.
func (m CollateralReport) GetProductComplex(f *field.ProductComplex) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for CollateralReport.
func (m CollateralReport) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from CollateralReport.
func (m CollateralReport) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for CollateralReport.
func (m CollateralReport) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from CollateralReport.
func (m CollateralReport) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for CollateralReport.
func (m CollateralReport) SettlMethod() (*field.SettlMethod, errors.MessageRejectError) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from CollateralReport.
func (m CollateralReport) GetSettlMethod(f *field.SettlMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for CollateralReport.
func (m CollateralReport) ExerciseStyle() (*field.ExerciseStyle, errors.MessageRejectError) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from CollateralReport.
func (m CollateralReport) GetExerciseStyle(f *field.ExerciseStyle) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutAmount is a non-required field for CollateralReport.
func (m CollateralReport) OptPayoutAmount() (*field.OptPayoutAmount, errors.MessageRejectError) {
	f := new(field.OptPayoutAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutAmount reads a OptPayoutAmount from CollateralReport.
func (m CollateralReport) GetOptPayoutAmount(f *field.OptPayoutAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for CollateralReport.
func (m CollateralReport) PriceQuoteMethod() (*field.PriceQuoteMethod, errors.MessageRejectError) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from CollateralReport.
func (m CollateralReport) GetPriceQuoteMethod(f *field.PriceQuoteMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for CollateralReport.
func (m CollateralReport) ListMethod() (*field.ListMethod, errors.MessageRejectError) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from CollateralReport.
func (m CollateralReport) GetListMethod(f *field.ListMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for CollateralReport.
func (m CollateralReport) CapPrice() (*field.CapPrice, errors.MessageRejectError) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from CollateralReport.
func (m CollateralReport) GetCapPrice(f *field.CapPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for CollateralReport.
func (m CollateralReport) FloorPrice() (*field.FloorPrice, errors.MessageRejectError) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from CollateralReport.
func (m CollateralReport) GetFloorPrice(f *field.FloorPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for CollateralReport.
func (m CollateralReport) PutOrCall() (*field.PutOrCall, errors.MessageRejectError) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from CollateralReport.
func (m CollateralReport) GetPutOrCall(f *field.PutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for CollateralReport.
func (m CollateralReport) FlexibleIndicator() (*field.FlexibleIndicator, errors.MessageRejectError) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from CollateralReport.
func (m CollateralReport) GetFlexibleIndicator(f *field.FlexibleIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for CollateralReport.
func (m CollateralReport) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, errors.MessageRejectError) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from CollateralReport.
func (m CollateralReport) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ValuationMethod is a non-required field for CollateralReport.
func (m CollateralReport) ValuationMethod() (*field.ValuationMethod, errors.MessageRejectError) {
	f := new(field.ValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetValuationMethod reads a ValuationMethod from CollateralReport.
func (m CollateralReport) GetValuationMethod(f *field.ValuationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplierUnit is a non-required field for CollateralReport.
func (m CollateralReport) ContractMultiplierUnit() (*field.ContractMultiplierUnit, errors.MessageRejectError) {
	f := new(field.ContractMultiplierUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplierUnit reads a ContractMultiplierUnit from CollateralReport.
func (m CollateralReport) GetContractMultiplierUnit(f *field.ContractMultiplierUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlowScheduleType is a non-required field for CollateralReport.
func (m CollateralReport) FlowScheduleType() (*field.FlowScheduleType, errors.MessageRejectError) {
	f := new(field.FlowScheduleType)
	err := m.Body.Get(f)
	return f, err
}

//GetFlowScheduleType reads a FlowScheduleType from CollateralReport.
func (m CollateralReport) GetFlowScheduleType(f *field.FlowScheduleType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RestructuringType is a non-required field for CollateralReport.
func (m CollateralReport) RestructuringType() (*field.RestructuringType, errors.MessageRejectError) {
	f := new(field.RestructuringType)
	err := m.Body.Get(f)
	return f, err
}

//GetRestructuringType reads a RestructuringType from CollateralReport.
func (m CollateralReport) GetRestructuringType(f *field.RestructuringType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Seniority is a non-required field for CollateralReport.
func (m CollateralReport) Seniority() (*field.Seniority, errors.MessageRejectError) {
	f := new(field.Seniority)
	err := m.Body.Get(f)
	return f, err
}

//GetSeniority reads a Seniority from CollateralReport.
func (m CollateralReport) GetSeniority(f *field.Seniority) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NotionalPercentageOutstanding is a non-required field for CollateralReport.
func (m CollateralReport) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstanding, errors.MessageRejectError) {
	f := new(field.NotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//GetNotionalPercentageOutstanding reads a NotionalPercentageOutstanding from CollateralReport.
func (m CollateralReport) GetNotionalPercentageOutstanding(f *field.NotionalPercentageOutstanding) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OriginalNotionalPercentageOutstanding is a non-required field for CollateralReport.
func (m CollateralReport) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstanding, errors.MessageRejectError) {
	f := new(field.OriginalNotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//GetOriginalNotionalPercentageOutstanding reads a OriginalNotionalPercentageOutstanding from CollateralReport.
func (m CollateralReport) GetOriginalNotionalPercentageOutstanding(f *field.OriginalNotionalPercentageOutstanding) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AttachmentPoint is a non-required field for CollateralReport.
func (m CollateralReport) AttachmentPoint() (*field.AttachmentPoint, errors.MessageRejectError) {
	f := new(field.AttachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//GetAttachmentPoint reads a AttachmentPoint from CollateralReport.
func (m CollateralReport) GetAttachmentPoint(f *field.AttachmentPoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DetachmentPoint is a non-required field for CollateralReport.
func (m CollateralReport) DetachmentPoint() (*field.DetachmentPoint, errors.MessageRejectError) {
	f := new(field.DetachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//GetDetachmentPoint reads a DetachmentPoint from CollateralReport.
func (m CollateralReport) GetDetachmentPoint(f *field.DetachmentPoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceDeterminationMethod is a non-required field for CollateralReport.
func (m CollateralReport) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethod, errors.MessageRejectError) {
	f := new(field.StrikePriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceDeterminationMethod reads a StrikePriceDeterminationMethod from CollateralReport.
func (m CollateralReport) GetStrikePriceDeterminationMethod(f *field.StrikePriceDeterminationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryMethod is a non-required field for CollateralReport.
func (m CollateralReport) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethod, errors.MessageRejectError) {
	f := new(field.StrikePriceBoundaryMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryMethod reads a StrikePriceBoundaryMethod from CollateralReport.
func (m CollateralReport) GetStrikePriceBoundaryMethod(f *field.StrikePriceBoundaryMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryPrecision is a non-required field for CollateralReport.
func (m CollateralReport) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecision, errors.MessageRejectError) {
	f := new(field.StrikePriceBoundaryPrecision)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryPrecision reads a StrikePriceBoundaryPrecision from CollateralReport.
func (m CollateralReport) GetStrikePriceBoundaryPrecision(f *field.StrikePriceBoundaryPrecision) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceDeterminationMethod is a non-required field for CollateralReport.
func (m CollateralReport) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethod, errors.MessageRejectError) {
	f := new(field.UnderlyingPriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceDeterminationMethod reads a UnderlyingPriceDeterminationMethod from CollateralReport.
func (m CollateralReport) GetUnderlyingPriceDeterminationMethod(f *field.UnderlyingPriceDeterminationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutType is a non-required field for CollateralReport.
func (m CollateralReport) OptPayoutType() (*field.OptPayoutType, errors.MessageRejectError) {
	f := new(field.OptPayoutType)
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutType reads a OptPayoutType from CollateralReport.
func (m CollateralReport) GetOptPayoutType(f *field.OptPayoutType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoComplexEvents is a non-required field for CollateralReport.
func (m CollateralReport) NoComplexEvents() (*field.NoComplexEvents, errors.MessageRejectError) {
	f := new(field.NoComplexEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoComplexEvents reads a NoComplexEvents from CollateralReport.
func (m CollateralReport) GetNoComplexEvents(f *field.NoComplexEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDesc is a non-required field for CollateralReport.
func (m CollateralReport) AgreementDesc() (*field.AgreementDesc, errors.MessageRejectError) {
	f := new(field.AgreementDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDesc reads a AgreementDesc from CollateralReport.
func (m CollateralReport) GetAgreementDesc(f *field.AgreementDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementID is a non-required field for CollateralReport.
func (m CollateralReport) AgreementID() (*field.AgreementID, errors.MessageRejectError) {
	f := new(field.AgreementID)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementID reads a AgreementID from CollateralReport.
func (m CollateralReport) GetAgreementID(f *field.AgreementID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDate is a non-required field for CollateralReport.
func (m CollateralReport) AgreementDate() (*field.AgreementDate, errors.MessageRejectError) {
	f := new(field.AgreementDate)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDate reads a AgreementDate from CollateralReport.
func (m CollateralReport) GetAgreementDate(f *field.AgreementDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementCurrency is a non-required field for CollateralReport.
func (m CollateralReport) AgreementCurrency() (*field.AgreementCurrency, errors.MessageRejectError) {
	f := new(field.AgreementCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementCurrency reads a AgreementCurrency from CollateralReport.
func (m CollateralReport) GetAgreementCurrency(f *field.AgreementCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TerminationType is a non-required field for CollateralReport.
func (m CollateralReport) TerminationType() (*field.TerminationType, errors.MessageRejectError) {
	f := new(field.TerminationType)
	err := m.Body.Get(f)
	return f, err
}

//GetTerminationType reads a TerminationType from CollateralReport.
func (m CollateralReport) GetTerminationType(f *field.TerminationType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartDate is a non-required field for CollateralReport.
func (m CollateralReport) StartDate() (*field.StartDate, errors.MessageRejectError) {
	f := new(field.StartDate)
	err := m.Body.Get(f)
	return f, err
}

//GetStartDate reads a StartDate from CollateralReport.
func (m CollateralReport) GetStartDate(f *field.StartDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndDate is a non-required field for CollateralReport.
func (m CollateralReport) EndDate() (*field.EndDate, errors.MessageRejectError) {
	f := new(field.EndDate)
	err := m.Body.Get(f)
	return f, err
}

//GetEndDate reads a EndDate from CollateralReport.
func (m CollateralReport) GetEndDate(f *field.EndDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryType is a non-required field for CollateralReport.
func (m CollateralReport) DeliveryType() (*field.DeliveryType, errors.MessageRejectError) {
	f := new(field.DeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryType reads a DeliveryType from CollateralReport.
func (m CollateralReport) GetDeliveryType(f *field.DeliveryType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarginRatio is a non-required field for CollateralReport.
func (m CollateralReport) MarginRatio() (*field.MarginRatio, errors.MessageRejectError) {
	f := new(field.MarginRatio)
	err := m.Body.Get(f)
	return f, err
}

//GetMarginRatio reads a MarginRatio from CollateralReport.
func (m CollateralReport) GetMarginRatio(f *field.MarginRatio) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for CollateralReport.
func (m CollateralReport) SettlDate() (*field.SettlDate, errors.MessageRejectError) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from CollateralReport.
func (m CollateralReport) GetSettlDate(f *field.SettlDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Quantity is a non-required field for CollateralReport.
func (m CollateralReport) Quantity() (*field.Quantity, errors.MessageRejectError) {
	f := new(field.Quantity)
	err := m.Body.Get(f)
	return f, err
}

//GetQuantity reads a Quantity from CollateralReport.
func (m CollateralReport) GetQuantity(f *field.Quantity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QtyType is a non-required field for CollateralReport.
func (m CollateralReport) QtyType() (*field.QtyType, errors.MessageRejectError) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}

//GetQtyType reads a QtyType from CollateralReport.
func (m CollateralReport) GetQtyType(f *field.QtyType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for CollateralReport.
func (m CollateralReport) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from CollateralReport.
func (m CollateralReport) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for CollateralReport.
func (m CollateralReport) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from CollateralReport.
func (m CollateralReport) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for CollateralReport.
func (m CollateralReport) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from CollateralReport.
func (m CollateralReport) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarginExcess is a non-required field for CollateralReport.
func (m CollateralReport) MarginExcess() (*field.MarginExcess, errors.MessageRejectError) {
	f := new(field.MarginExcess)
	err := m.Body.Get(f)
	return f, err
}

//GetMarginExcess reads a MarginExcess from CollateralReport.
func (m CollateralReport) GetMarginExcess(f *field.MarginExcess) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalNetValue is a non-required field for CollateralReport.
func (m CollateralReport) TotalNetValue() (*field.TotalNetValue, errors.MessageRejectError) {
	f := new(field.TotalNetValue)
	err := m.Body.Get(f)
	return f, err
}

//GetTotalNetValue reads a TotalNetValue from CollateralReport.
func (m CollateralReport) GetTotalNetValue(f *field.TotalNetValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashOutstanding is a non-required field for CollateralReport.
func (m CollateralReport) CashOutstanding() (*field.CashOutstanding, errors.MessageRejectError) {
	f := new(field.CashOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//GetCashOutstanding reads a CashOutstanding from CollateralReport.
func (m CollateralReport) GetCashOutstanding(f *field.CashOutstanding) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTrdRegTimestamps is a non-required field for CollateralReport.
func (m CollateralReport) NoTrdRegTimestamps() (*field.NoTrdRegTimestamps, errors.MessageRejectError) {
	f := new(field.NoTrdRegTimestamps)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTrdRegTimestamps reads a NoTrdRegTimestamps from CollateralReport.
func (m CollateralReport) GetNoTrdRegTimestamps(f *field.NoTrdRegTimestamps) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a non-required field for CollateralReport.
func (m CollateralReport) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from CollateralReport.
func (m CollateralReport) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoMiscFees is a non-required field for CollateralReport.
func (m CollateralReport) NoMiscFees() (*field.NoMiscFees, errors.MessageRejectError) {
	f := new(field.NoMiscFees)
	err := m.Body.Get(f)
	return f, err
}

//GetNoMiscFees reads a NoMiscFees from CollateralReport.
func (m CollateralReport) GetNoMiscFees(f *field.NoMiscFees) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for CollateralReport.
func (m CollateralReport) Price() (*field.Price, errors.MessageRejectError) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from CollateralReport.
func (m CollateralReport) GetPrice(f *field.Price) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for CollateralReport.
func (m CollateralReport) PriceType() (*field.PriceType, errors.MessageRejectError) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from CollateralReport.
func (m CollateralReport) GetPriceType(f *field.PriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestAmt is a non-required field for CollateralReport.
func (m CollateralReport) AccruedInterestAmt() (*field.AccruedInterestAmt, errors.MessageRejectError) {
	f := new(field.AccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestAmt reads a AccruedInterestAmt from CollateralReport.
func (m CollateralReport) GetAccruedInterestAmt(f *field.AccruedInterestAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndAccruedInterestAmt is a non-required field for CollateralReport.
func (m CollateralReport) EndAccruedInterestAmt() (*field.EndAccruedInterestAmt, errors.MessageRejectError) {
	f := new(field.EndAccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetEndAccruedInterestAmt reads a EndAccruedInterestAmt from CollateralReport.
func (m CollateralReport) GetEndAccruedInterestAmt(f *field.EndAccruedInterestAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartCash is a non-required field for CollateralReport.
func (m CollateralReport) StartCash() (*field.StartCash, errors.MessageRejectError) {
	f := new(field.StartCash)
	err := m.Body.Get(f)
	return f, err
}

//GetStartCash reads a StartCash from CollateralReport.
func (m CollateralReport) GetStartCash(f *field.StartCash) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndCash is a non-required field for CollateralReport.
func (m CollateralReport) EndCash() (*field.EndCash, errors.MessageRejectError) {
	f := new(field.EndCash)
	err := m.Body.Get(f)
	return f, err
}

//GetEndCash reads a EndCash from CollateralReport.
func (m CollateralReport) GetEndCash(f *field.EndCash) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for CollateralReport.
func (m CollateralReport) Spread() (*field.Spread, errors.MessageRejectError) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from CollateralReport.
func (m CollateralReport) GetSpread(f *field.Spread) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for CollateralReport.
func (m CollateralReport) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, errors.MessageRejectError) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from CollateralReport.
func (m CollateralReport) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for CollateralReport.
func (m CollateralReport) BenchmarkCurveName() (*field.BenchmarkCurveName, errors.MessageRejectError) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from CollateralReport.
func (m CollateralReport) GetBenchmarkCurveName(f *field.BenchmarkCurveName) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for CollateralReport.
func (m CollateralReport) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, errors.MessageRejectError) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from CollateralReport.
func (m CollateralReport) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPrice is a non-required field for CollateralReport.
func (m CollateralReport) BenchmarkPrice() (*field.BenchmarkPrice, errors.MessageRejectError) {
	f := new(field.BenchmarkPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPrice reads a BenchmarkPrice from CollateralReport.
func (m CollateralReport) GetBenchmarkPrice(f *field.BenchmarkPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPriceType is a non-required field for CollateralReport.
func (m CollateralReport) BenchmarkPriceType() (*field.BenchmarkPriceType, errors.MessageRejectError) {
	f := new(field.BenchmarkPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPriceType reads a BenchmarkPriceType from CollateralReport.
func (m CollateralReport) GetBenchmarkPriceType(f *field.BenchmarkPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityID is a non-required field for CollateralReport.
func (m CollateralReport) BenchmarkSecurityID() (*field.BenchmarkSecurityID, errors.MessageRejectError) {
	f := new(field.BenchmarkSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityID reads a BenchmarkSecurityID from CollateralReport.
func (m CollateralReport) GetBenchmarkSecurityID(f *field.BenchmarkSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityIDSource is a non-required field for CollateralReport.
func (m CollateralReport) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSource, errors.MessageRejectError) {
	f := new(field.BenchmarkSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityIDSource reads a BenchmarkSecurityIDSource from CollateralReport.
func (m CollateralReport) GetBenchmarkSecurityIDSource(f *field.BenchmarkSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStipulations is a non-required field for CollateralReport.
func (m CollateralReport) NoStipulations() (*field.NoStipulations, errors.MessageRejectError) {
	f := new(field.NoStipulations)
	err := m.Body.Get(f)
	return f, err
}

//GetNoStipulations reads a NoStipulations from CollateralReport.
func (m CollateralReport) GetNoStipulations(f *field.NoStipulations) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDeliveryType is a non-required field for CollateralReport.
func (m CollateralReport) SettlDeliveryType() (*field.SettlDeliveryType, errors.MessageRejectError) {
	f := new(field.SettlDeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDeliveryType reads a SettlDeliveryType from CollateralReport.
func (m CollateralReport) GetSettlDeliveryType(f *field.SettlDeliveryType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbType is a non-required field for CollateralReport.
func (m CollateralReport) StandInstDbType() (*field.StandInstDbType, errors.MessageRejectError) {
	f := new(field.StandInstDbType)
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbType reads a StandInstDbType from CollateralReport.
func (m CollateralReport) GetStandInstDbType(f *field.StandInstDbType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbName is a non-required field for CollateralReport.
func (m CollateralReport) StandInstDbName() (*field.StandInstDbName, errors.MessageRejectError) {
	f := new(field.StandInstDbName)
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbName reads a StandInstDbName from CollateralReport.
func (m CollateralReport) GetStandInstDbName(f *field.StandInstDbName) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbID is a non-required field for CollateralReport.
func (m CollateralReport) StandInstDbID() (*field.StandInstDbID, errors.MessageRejectError) {
	f := new(field.StandInstDbID)
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbID reads a StandInstDbID from CollateralReport.
func (m CollateralReport) GetStandInstDbID(f *field.StandInstDbID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoDlvyInst is a non-required field for CollateralReport.
func (m CollateralReport) NoDlvyInst() (*field.NoDlvyInst, errors.MessageRejectError) {
	f := new(field.NoDlvyInst)
	err := m.Body.Get(f)
	return f, err
}

//GetNoDlvyInst reads a NoDlvyInst from CollateralReport.
func (m CollateralReport) GetNoDlvyInst(f *field.NoDlvyInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for CollateralReport.
func (m CollateralReport) TradingSessionID() (*field.TradingSessionID, errors.MessageRejectError) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from CollateralReport.
func (m CollateralReport) GetTradingSessionID(f *field.TradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for CollateralReport.
func (m CollateralReport) TradingSessionSubID() (*field.TradingSessionSubID, errors.MessageRejectError) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from CollateralReport.
func (m CollateralReport) GetTradingSessionSubID(f *field.TradingSessionSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessID is a non-required field for CollateralReport.
func (m CollateralReport) SettlSessID() (*field.SettlSessID, errors.MessageRejectError) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessID reads a SettlSessID from CollateralReport.
func (m CollateralReport) GetSettlSessID(f *field.SettlSessID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessSubID is a non-required field for CollateralReport.
func (m CollateralReport) SettlSessSubID() (*field.SettlSessSubID, errors.MessageRejectError) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessSubID reads a SettlSessSubID from CollateralReport.
func (m CollateralReport) GetSettlSessSubID(f *field.SettlSessSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a non-required field for CollateralReport.
func (m CollateralReport) ClearingBusinessDate() (*field.ClearingBusinessDate, errors.MessageRejectError) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from CollateralReport.
func (m CollateralReport) GetClearingBusinessDate(f *field.ClearingBusinessDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for CollateralReport.
func (m CollateralReport) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from CollateralReport.
func (m CollateralReport) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for CollateralReport.
func (m CollateralReport) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from CollateralReport.
func (m CollateralReport) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for CollateralReport.
func (m CollateralReport) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from CollateralReport.
func (m CollateralReport) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for CollateralReport.
func (m CollateralReport) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from CollateralReport.
func (m CollateralReport) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CollApplType is a non-required field for CollateralReport.
func (m CollateralReport) CollApplType() (*field.CollApplType, errors.MessageRejectError) {
	f := new(field.CollApplType)
	err := m.Body.Get(f)
	return f, err
}

//GetCollApplType reads a CollApplType from CollateralReport.
func (m CollateralReport) GetCollApplType(f *field.CollApplType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FinancialStatus is a non-required field for CollateralReport.
func (m CollateralReport) FinancialStatus() (*field.FinancialStatus, errors.MessageRejectError) {
	f := new(field.FinancialStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetFinancialStatus reads a FinancialStatus from CollateralReport.
func (m CollateralReport) GetFinancialStatus(f *field.FinancialStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}
