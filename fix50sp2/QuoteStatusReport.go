package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//QuoteStatusReport msg type = AI.
type QuoteStatusReport struct {
	message.Message
}

//QuoteStatusReportBuilder builds QuoteStatusReport messages.
type QuoteStatusReportBuilder struct {
	message.MessageBuilder
}

//CreateQuoteStatusReportBuilder returns an initialized QuoteStatusReportBuilder with specified required fields.
func CreateQuoteStatusReportBuilder() QuoteStatusReportBuilder {
	var builder QuoteStatusReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	return builder
}

//QuoteStatusReqID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteStatusReqID() (*field.QuoteStatusReqID, errors.MessageRejectError) {
	f := new(field.QuoteStatusReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteStatusReqID reads a QuoteStatusReqID from QuoteStatusReport.
func (m QuoteStatusReport) GetQuoteStatusReqID(f *field.QuoteStatusReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteReqID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteReqID() (*field.QuoteReqID, errors.MessageRejectError) {
	f := new(field.QuoteReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteReqID reads a QuoteReqID from QuoteStatusReport.
func (m QuoteStatusReport) GetQuoteReqID(f *field.QuoteReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteID() (*field.QuoteID, errors.MessageRejectError) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from QuoteStatusReport.
func (m QuoteStatusReport) GetQuoteID(f *field.QuoteID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteRespID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteRespID() (*field.QuoteRespID, errors.MessageRejectError) {
	f := new(field.QuoteRespID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteRespID reads a QuoteRespID from QuoteStatusReport.
func (m QuoteStatusReport) GetQuoteRespID(f *field.QuoteRespID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteType() (*field.QuoteType, errors.MessageRejectError) {
	f := new(field.QuoteType)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteType reads a QuoteType from QuoteStatusReport.
func (m QuoteStatusReport) GetQuoteType(f *field.QuoteType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from QuoteStatusReport.
func (m QuoteStatusReport) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) TradingSessionID() (*field.TradingSessionID, errors.MessageRejectError) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from QuoteStatusReport.
func (m QuoteStatusReport) GetTradingSessionID(f *field.TradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) TradingSessionSubID() (*field.TradingSessionSubID, errors.MessageRejectError) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from QuoteStatusReport.
func (m QuoteStatusReport) GetTradingSessionSubID(f *field.TradingSessionSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from QuoteStatusReport.
func (m QuoteStatusReport) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from QuoteStatusReport.
func (m QuoteStatusReport) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from QuoteStatusReport.
func (m QuoteStatusReport) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from QuoteStatusReport.
func (m QuoteStatusReport) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from QuoteStatusReport.
func (m QuoteStatusReport) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from QuoteStatusReport.
func (m QuoteStatusReport) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from QuoteStatusReport.
func (m QuoteStatusReport) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from QuoteStatusReport.
func (m QuoteStatusReport) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from QuoteStatusReport.
func (m QuoteStatusReport) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from QuoteStatusReport.
func (m QuoteStatusReport) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from QuoteStatusReport.
func (m QuoteStatusReport) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from QuoteStatusReport.
func (m QuoteStatusReport) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from QuoteStatusReport.
func (m QuoteStatusReport) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from QuoteStatusReport.
func (m QuoteStatusReport) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from QuoteStatusReport.
func (m QuoteStatusReport) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from QuoteStatusReport.
func (m QuoteStatusReport) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from QuoteStatusReport.
func (m QuoteStatusReport) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from QuoteStatusReport.
func (m QuoteStatusReport) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from QuoteStatusReport.
func (m QuoteStatusReport) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from QuoteStatusReport.
func (m QuoteStatusReport) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from QuoteStatusReport.
func (m QuoteStatusReport) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from QuoteStatusReport.
func (m QuoteStatusReport) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from QuoteStatusReport.
func (m QuoteStatusReport) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from QuoteStatusReport.
func (m QuoteStatusReport) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from QuoteStatusReport.
func (m QuoteStatusReport) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from QuoteStatusReport.
func (m QuoteStatusReport) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from QuoteStatusReport.
func (m QuoteStatusReport) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from QuoteStatusReport.
func (m QuoteStatusReport) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from QuoteStatusReport.
func (m QuoteStatusReport) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from QuoteStatusReport.
func (m QuoteStatusReport) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from QuoteStatusReport.
func (m QuoteStatusReport) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from QuoteStatusReport.
func (m QuoteStatusReport) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from QuoteStatusReport.
func (m QuoteStatusReport) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from QuoteStatusReport.
func (m QuoteStatusReport) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from QuoteStatusReport.
func (m QuoteStatusReport) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from QuoteStatusReport.
func (m QuoteStatusReport) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from QuoteStatusReport.
func (m QuoteStatusReport) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from QuoteStatusReport.
func (m QuoteStatusReport) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from QuoteStatusReport.
func (m QuoteStatusReport) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from QuoteStatusReport.
func (m QuoteStatusReport) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from QuoteStatusReport.
func (m QuoteStatusReport) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from QuoteStatusReport.
func (m QuoteStatusReport) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityStatus() (*field.SecurityStatus, errors.MessageRejectError) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from QuoteStatusReport.
func (m QuoteStatusReport) GetSecurityStatus(f *field.SecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SettleOnOpenFlag() (*field.SettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from QuoteStatusReport.
func (m QuoteStatusReport) GetSettleOnOpenFlag(f *field.SettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from QuoteStatusReport.
func (m QuoteStatusReport) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) StrikeMultiplier() (*field.StrikeMultiplier, errors.MessageRejectError) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from QuoteStatusReport.
func (m QuoteStatusReport) GetStrikeMultiplier(f *field.StrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) StrikeValue() (*field.StrikeValue, errors.MessageRejectError) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from QuoteStatusReport.
func (m QuoteStatusReport) GetStrikeValue(f *field.StrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MinPriceIncrement() (*field.MinPriceIncrement, errors.MessageRejectError) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from QuoteStatusReport.
func (m QuoteStatusReport) GetMinPriceIncrement(f *field.MinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) PositionLimit() (*field.PositionLimit, errors.MessageRejectError) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from QuoteStatusReport.
func (m QuoteStatusReport) GetPositionLimit(f *field.PositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NTPositionLimit() (*field.NTPositionLimit, errors.MessageRejectError) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from QuoteStatusReport.
func (m QuoteStatusReport) GetNTPositionLimit(f *field.NTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NoInstrumentParties() (*field.NoInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from QuoteStatusReport.
func (m QuoteStatusReport) GetNoInstrumentParties(f *field.NoInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) UnitOfMeasure() (*field.UnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from QuoteStatusReport.
func (m QuoteStatusReport) GetUnitOfMeasure(f *field.UnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) TimeUnit() (*field.TimeUnit, errors.MessageRejectError) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from QuoteStatusReport.
func (m QuoteStatusReport) GetTimeUnit(f *field.TimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MaturityTime() (*field.MaturityTime, errors.MessageRejectError) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from QuoteStatusReport.
func (m QuoteStatusReport) GetMaturityTime(f *field.MaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityGroup() (*field.SecurityGroup, errors.MessageRejectError) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from QuoteStatusReport.
func (m QuoteStatusReport) GetSecurityGroup(f *field.SecurityGroup) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, errors.MessageRejectError) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from QuoteStatusReport.
func (m QuoteStatusReport) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) UnitOfMeasureQty() (*field.UnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from QuoteStatusReport.
func (m QuoteStatusReport) GetUnitOfMeasureQty(f *field.UnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityXMLLen() (*field.SecurityXMLLen, errors.MessageRejectError) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from QuoteStatusReport.
func (m QuoteStatusReport) GetSecurityXMLLen(f *field.SecurityXMLLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityXML() (*field.SecurityXML, errors.MessageRejectError) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from QuoteStatusReport.
func (m QuoteStatusReport) GetSecurityXML(f *field.SecurityXML) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityXMLSchema() (*field.SecurityXMLSchema, errors.MessageRejectError) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from QuoteStatusReport.
func (m QuoteStatusReport) GetSecurityXMLSchema(f *field.SecurityXMLSchema) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ProductComplex() (*field.ProductComplex, errors.MessageRejectError) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from QuoteStatusReport.
func (m QuoteStatusReport) GetProductComplex(f *field.ProductComplex) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from QuoteStatusReport.
func (m QuoteStatusReport) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from QuoteStatusReport.
func (m QuoteStatusReport) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SettlMethod() (*field.SettlMethod, errors.MessageRejectError) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from QuoteStatusReport.
func (m QuoteStatusReport) GetSettlMethod(f *field.SettlMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ExerciseStyle() (*field.ExerciseStyle, errors.MessageRejectError) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from QuoteStatusReport.
func (m QuoteStatusReport) GetExerciseStyle(f *field.ExerciseStyle) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutAmount is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OptPayoutAmount() (*field.OptPayoutAmount, errors.MessageRejectError) {
	f := new(field.OptPayoutAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutAmount reads a OptPayoutAmount from QuoteStatusReport.
func (m QuoteStatusReport) GetOptPayoutAmount(f *field.OptPayoutAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) PriceQuoteMethod() (*field.PriceQuoteMethod, errors.MessageRejectError) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from QuoteStatusReport.
func (m QuoteStatusReport) GetPriceQuoteMethod(f *field.PriceQuoteMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ListMethod() (*field.ListMethod, errors.MessageRejectError) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from QuoteStatusReport.
func (m QuoteStatusReport) GetListMethod(f *field.ListMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CapPrice() (*field.CapPrice, errors.MessageRejectError) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from QuoteStatusReport.
func (m QuoteStatusReport) GetCapPrice(f *field.CapPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) FloorPrice() (*field.FloorPrice, errors.MessageRejectError) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from QuoteStatusReport.
func (m QuoteStatusReport) GetFloorPrice(f *field.FloorPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) PutOrCall() (*field.PutOrCall, errors.MessageRejectError) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from QuoteStatusReport.
func (m QuoteStatusReport) GetPutOrCall(f *field.PutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) FlexibleIndicator() (*field.FlexibleIndicator, errors.MessageRejectError) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from QuoteStatusReport.
func (m QuoteStatusReport) GetFlexibleIndicator(f *field.FlexibleIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, errors.MessageRejectError) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from QuoteStatusReport.
func (m QuoteStatusReport) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ValuationMethod is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ValuationMethod() (*field.ValuationMethod, errors.MessageRejectError) {
	f := new(field.ValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetValuationMethod reads a ValuationMethod from QuoteStatusReport.
func (m QuoteStatusReport) GetValuationMethod(f *field.ValuationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplierUnit is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ContractMultiplierUnit() (*field.ContractMultiplierUnit, errors.MessageRejectError) {
	f := new(field.ContractMultiplierUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplierUnit reads a ContractMultiplierUnit from QuoteStatusReport.
func (m QuoteStatusReport) GetContractMultiplierUnit(f *field.ContractMultiplierUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlowScheduleType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) FlowScheduleType() (*field.FlowScheduleType, errors.MessageRejectError) {
	f := new(field.FlowScheduleType)
	err := m.Body.Get(f)
	return f, err
}

//GetFlowScheduleType reads a FlowScheduleType from QuoteStatusReport.
func (m QuoteStatusReport) GetFlowScheduleType(f *field.FlowScheduleType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RestructuringType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) RestructuringType() (*field.RestructuringType, errors.MessageRejectError) {
	f := new(field.RestructuringType)
	err := m.Body.Get(f)
	return f, err
}

//GetRestructuringType reads a RestructuringType from QuoteStatusReport.
func (m QuoteStatusReport) GetRestructuringType(f *field.RestructuringType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Seniority is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Seniority() (*field.Seniority, errors.MessageRejectError) {
	f := new(field.Seniority)
	err := m.Body.Get(f)
	return f, err
}

//GetSeniority reads a Seniority from QuoteStatusReport.
func (m QuoteStatusReport) GetSeniority(f *field.Seniority) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NotionalPercentageOutstanding is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstanding, errors.MessageRejectError) {
	f := new(field.NotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//GetNotionalPercentageOutstanding reads a NotionalPercentageOutstanding from QuoteStatusReport.
func (m QuoteStatusReport) GetNotionalPercentageOutstanding(f *field.NotionalPercentageOutstanding) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OriginalNotionalPercentageOutstanding is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstanding, errors.MessageRejectError) {
	f := new(field.OriginalNotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//GetOriginalNotionalPercentageOutstanding reads a OriginalNotionalPercentageOutstanding from QuoteStatusReport.
func (m QuoteStatusReport) GetOriginalNotionalPercentageOutstanding(f *field.OriginalNotionalPercentageOutstanding) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AttachmentPoint is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) AttachmentPoint() (*field.AttachmentPoint, errors.MessageRejectError) {
	f := new(field.AttachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//GetAttachmentPoint reads a AttachmentPoint from QuoteStatusReport.
func (m QuoteStatusReport) GetAttachmentPoint(f *field.AttachmentPoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DetachmentPoint is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) DetachmentPoint() (*field.DetachmentPoint, errors.MessageRejectError) {
	f := new(field.DetachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//GetDetachmentPoint reads a DetachmentPoint from QuoteStatusReport.
func (m QuoteStatusReport) GetDetachmentPoint(f *field.DetachmentPoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceDeterminationMethod is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethod, errors.MessageRejectError) {
	f := new(field.StrikePriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceDeterminationMethod reads a StrikePriceDeterminationMethod from QuoteStatusReport.
func (m QuoteStatusReport) GetStrikePriceDeterminationMethod(f *field.StrikePriceDeterminationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryMethod is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethod, errors.MessageRejectError) {
	f := new(field.StrikePriceBoundaryMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryMethod reads a StrikePriceBoundaryMethod from QuoteStatusReport.
func (m QuoteStatusReport) GetStrikePriceBoundaryMethod(f *field.StrikePriceBoundaryMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryPrecision is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecision, errors.MessageRejectError) {
	f := new(field.StrikePriceBoundaryPrecision)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryPrecision reads a StrikePriceBoundaryPrecision from QuoteStatusReport.
func (m QuoteStatusReport) GetStrikePriceBoundaryPrecision(f *field.StrikePriceBoundaryPrecision) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceDeterminationMethod is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethod, errors.MessageRejectError) {
	f := new(field.UnderlyingPriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceDeterminationMethod reads a UnderlyingPriceDeterminationMethod from QuoteStatusReport.
func (m QuoteStatusReport) GetUnderlyingPriceDeterminationMethod(f *field.UnderlyingPriceDeterminationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OptPayoutType() (*field.OptPayoutType, errors.MessageRejectError) {
	f := new(field.OptPayoutType)
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutType reads a OptPayoutType from QuoteStatusReport.
func (m QuoteStatusReport) GetOptPayoutType(f *field.OptPayoutType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoComplexEvents is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NoComplexEvents() (*field.NoComplexEvents, errors.MessageRejectError) {
	f := new(field.NoComplexEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoComplexEvents reads a NoComplexEvents from QuoteStatusReport.
func (m QuoteStatusReport) GetNoComplexEvents(f *field.NoComplexEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDesc is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) AgreementDesc() (*field.AgreementDesc, errors.MessageRejectError) {
	f := new(field.AgreementDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDesc reads a AgreementDesc from QuoteStatusReport.
func (m QuoteStatusReport) GetAgreementDesc(f *field.AgreementDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) AgreementID() (*field.AgreementID, errors.MessageRejectError) {
	f := new(field.AgreementID)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementID reads a AgreementID from QuoteStatusReport.
func (m QuoteStatusReport) GetAgreementID(f *field.AgreementID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) AgreementDate() (*field.AgreementDate, errors.MessageRejectError) {
	f := new(field.AgreementDate)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDate reads a AgreementDate from QuoteStatusReport.
func (m QuoteStatusReport) GetAgreementDate(f *field.AgreementDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementCurrency is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) AgreementCurrency() (*field.AgreementCurrency, errors.MessageRejectError) {
	f := new(field.AgreementCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementCurrency reads a AgreementCurrency from QuoteStatusReport.
func (m QuoteStatusReport) GetAgreementCurrency(f *field.AgreementCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TerminationType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) TerminationType() (*field.TerminationType, errors.MessageRejectError) {
	f := new(field.TerminationType)
	err := m.Body.Get(f)
	return f, err
}

//GetTerminationType reads a TerminationType from QuoteStatusReport.
func (m QuoteStatusReport) GetTerminationType(f *field.TerminationType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) StartDate() (*field.StartDate, errors.MessageRejectError) {
	f := new(field.StartDate)
	err := m.Body.Get(f)
	return f, err
}

//GetStartDate reads a StartDate from QuoteStatusReport.
func (m QuoteStatusReport) GetStartDate(f *field.StartDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) EndDate() (*field.EndDate, errors.MessageRejectError) {
	f := new(field.EndDate)
	err := m.Body.Get(f)
	return f, err
}

//GetEndDate reads a EndDate from QuoteStatusReport.
func (m QuoteStatusReport) GetEndDate(f *field.EndDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) DeliveryType() (*field.DeliveryType, errors.MessageRejectError) {
	f := new(field.DeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryType reads a DeliveryType from QuoteStatusReport.
func (m QuoteStatusReport) GetDeliveryType(f *field.DeliveryType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarginRatio is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MarginRatio() (*field.MarginRatio, errors.MessageRejectError) {
	f := new(field.MarginRatio)
	err := m.Body.Get(f)
	return f, err
}

//GetMarginRatio reads a MarginRatio from QuoteStatusReport.
func (m QuoteStatusReport) GetMarginRatio(f *field.MarginRatio) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from QuoteStatusReport.
func (m QuoteStatusReport) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from QuoteStatusReport.
func (m QuoteStatusReport) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OrderQty() (*field.OrderQty, errors.MessageRejectError) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from QuoteStatusReport.
func (m QuoteStatusReport) GetOrderQty(f *field.OrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashOrderQty is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CashOrderQty() (*field.CashOrderQty, errors.MessageRejectError) {
	f := new(field.CashOrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetCashOrderQty reads a CashOrderQty from QuoteStatusReport.
func (m QuoteStatusReport) GetCashOrderQty(f *field.CashOrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderPercent is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OrderPercent() (*field.OrderPercent, errors.MessageRejectError) {
	f := new(field.OrderPercent)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderPercent reads a OrderPercent from QuoteStatusReport.
func (m QuoteStatusReport) GetOrderPercent(f *field.OrderPercent) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingDirection is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) RoundingDirection() (*field.RoundingDirection, errors.MessageRejectError) {
	f := new(field.RoundingDirection)
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingDirection reads a RoundingDirection from QuoteStatusReport.
func (m QuoteStatusReport) GetRoundingDirection(f *field.RoundingDirection) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingModulus is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) RoundingModulus() (*field.RoundingModulus, errors.MessageRejectError) {
	f := new(field.RoundingModulus)
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingModulus reads a RoundingModulus from QuoteStatusReport.
func (m QuoteStatusReport) GetRoundingModulus(f *field.RoundingModulus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SettlType() (*field.SettlType, errors.MessageRejectError) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlType reads a SettlType from QuoteStatusReport.
func (m QuoteStatusReport) GetSettlType(f *field.SettlType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SettlDate() (*field.SettlDate, errors.MessageRejectError) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from QuoteStatusReport.
func (m QuoteStatusReport) GetSettlDate(f *field.SettlDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate2 is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SettlDate2() (*field.SettlDate2, errors.MessageRejectError) {
	f := new(field.SettlDate2)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate2 reads a SettlDate2 from QuoteStatusReport.
func (m QuoteStatusReport) GetSettlDate2(f *field.SettlDate2) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty2 is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OrderQty2() (*field.OrderQty2, errors.MessageRejectError) {
	f := new(field.OrderQty2)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty2 reads a OrderQty2 from QuoteStatusReport.
func (m QuoteStatusReport) GetOrderQty2(f *field.OrderQty2) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from QuoteStatusReport.
func (m QuoteStatusReport) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStipulations is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NoStipulations() (*field.NoStipulations, errors.MessageRejectError) {
	f := new(field.NoStipulations)
	err := m.Body.Get(f)
	return f, err
}

//GetNoStipulations reads a NoStipulations from QuoteStatusReport.
func (m QuoteStatusReport) GetNoStipulations(f *field.NoStipulations) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Account() (*field.Account, errors.MessageRejectError) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from QuoteStatusReport.
func (m QuoteStatusReport) GetAccount(f *field.Account) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) AcctIDSource() (*field.AcctIDSource, errors.MessageRejectError) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from QuoteStatusReport.
func (m QuoteStatusReport) GetAcctIDSource(f *field.AcctIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) AccountType() (*field.AccountType, errors.MessageRejectError) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from QuoteStatusReport.
func (m QuoteStatusReport) GetAccountType(f *field.AccountType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from QuoteStatusReport.
func (m QuoteStatusReport) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoQuoteQualifiers is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NoQuoteQualifiers() (*field.NoQuoteQualifiers, errors.MessageRejectError) {
	f := new(field.NoQuoteQualifiers)
	err := m.Body.Get(f)
	return f, err
}

//GetNoQuoteQualifiers reads a NoQuoteQualifiers from QuoteStatusReport.
func (m QuoteStatusReport) GetNoQuoteQualifiers(f *field.NoQuoteQualifiers) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ExpireTime() (*field.ExpireTime, errors.MessageRejectError) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from QuoteStatusReport.
func (m QuoteStatusReport) GetExpireTime(f *field.ExpireTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Price() (*field.Price, errors.MessageRejectError) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from QuoteStatusReport.
func (m QuoteStatusReport) GetPrice(f *field.Price) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) PriceType() (*field.PriceType, errors.MessageRejectError) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from QuoteStatusReport.
func (m QuoteStatusReport) GetPriceType(f *field.PriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Spread() (*field.Spread, errors.MessageRejectError) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from QuoteStatusReport.
func (m QuoteStatusReport) GetSpread(f *field.Spread) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, errors.MessageRejectError) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from QuoteStatusReport.
func (m QuoteStatusReport) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BenchmarkCurveName() (*field.BenchmarkCurveName, errors.MessageRejectError) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from QuoteStatusReport.
func (m QuoteStatusReport) GetBenchmarkCurveName(f *field.BenchmarkCurveName) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, errors.MessageRejectError) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from QuoteStatusReport.
func (m QuoteStatusReport) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPrice is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BenchmarkPrice() (*field.BenchmarkPrice, errors.MessageRejectError) {
	f := new(field.BenchmarkPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPrice reads a BenchmarkPrice from QuoteStatusReport.
func (m QuoteStatusReport) GetBenchmarkPrice(f *field.BenchmarkPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPriceType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BenchmarkPriceType() (*field.BenchmarkPriceType, errors.MessageRejectError) {
	f := new(field.BenchmarkPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPriceType reads a BenchmarkPriceType from QuoteStatusReport.
func (m QuoteStatusReport) GetBenchmarkPriceType(f *field.BenchmarkPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BenchmarkSecurityID() (*field.BenchmarkSecurityID, errors.MessageRejectError) {
	f := new(field.BenchmarkSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityID reads a BenchmarkSecurityID from QuoteStatusReport.
func (m QuoteStatusReport) GetBenchmarkSecurityID(f *field.BenchmarkSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityIDSource is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSource, errors.MessageRejectError) {
	f := new(field.BenchmarkSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityIDSource reads a BenchmarkSecurityIDSource from QuoteStatusReport.
func (m QuoteStatusReport) GetBenchmarkSecurityIDSource(f *field.BenchmarkSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) YieldType() (*field.YieldType, errors.MessageRejectError) {
	f := new(field.YieldType)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldType reads a YieldType from QuoteStatusReport.
func (m QuoteStatusReport) GetYieldType(f *field.YieldType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Yield is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Yield() (*field.Yield, errors.MessageRejectError) {
	f := new(field.Yield)
	err := m.Body.Get(f)
	return f, err
}

//GetYield reads a Yield from QuoteStatusReport.
func (m QuoteStatusReport) GetYield(f *field.Yield) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldCalcDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) YieldCalcDate() (*field.YieldCalcDate, errors.MessageRejectError) {
	f := new(field.YieldCalcDate)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldCalcDate reads a YieldCalcDate from QuoteStatusReport.
func (m QuoteStatusReport) GetYieldCalcDate(f *field.YieldCalcDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) YieldRedemptionDate() (*field.YieldRedemptionDate, errors.MessageRejectError) {
	f := new(field.YieldRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionDate reads a YieldRedemptionDate from QuoteStatusReport.
func (m QuoteStatusReport) GetYieldRedemptionDate(f *field.YieldRedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPrice is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) YieldRedemptionPrice() (*field.YieldRedemptionPrice, errors.MessageRejectError) {
	f := new(field.YieldRedemptionPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPrice reads a YieldRedemptionPrice from QuoteStatusReport.
func (m QuoteStatusReport) GetYieldRedemptionPrice(f *field.YieldRedemptionPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPriceType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) YieldRedemptionPriceType() (*field.YieldRedemptionPriceType, errors.MessageRejectError) {
	f := new(field.YieldRedemptionPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPriceType reads a YieldRedemptionPriceType from QuoteStatusReport.
func (m QuoteStatusReport) GetYieldRedemptionPriceType(f *field.YieldRedemptionPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidPx is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BidPx() (*field.BidPx, errors.MessageRejectError) {
	f := new(field.BidPx)
	err := m.Body.Get(f)
	return f, err
}

//GetBidPx reads a BidPx from QuoteStatusReport.
func (m QuoteStatusReport) GetBidPx(f *field.BidPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OfferPx is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OfferPx() (*field.OfferPx, errors.MessageRejectError) {
	f := new(field.OfferPx)
	err := m.Body.Get(f)
	return f, err
}

//GetOfferPx reads a OfferPx from QuoteStatusReport.
func (m QuoteStatusReport) GetOfferPx(f *field.OfferPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MktBidPx is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MktBidPx() (*field.MktBidPx, errors.MessageRejectError) {
	f := new(field.MktBidPx)
	err := m.Body.Get(f)
	return f, err
}

//GetMktBidPx reads a MktBidPx from QuoteStatusReport.
func (m QuoteStatusReport) GetMktBidPx(f *field.MktBidPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MktOfferPx is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MktOfferPx() (*field.MktOfferPx, errors.MessageRejectError) {
	f := new(field.MktOfferPx)
	err := m.Body.Get(f)
	return f, err
}

//GetMktOfferPx reads a MktOfferPx from QuoteStatusReport.
func (m QuoteStatusReport) GetMktOfferPx(f *field.MktOfferPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinBidSize is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MinBidSize() (*field.MinBidSize, errors.MessageRejectError) {
	f := new(field.MinBidSize)
	err := m.Body.Get(f)
	return f, err
}

//GetMinBidSize reads a MinBidSize from QuoteStatusReport.
func (m QuoteStatusReport) GetMinBidSize(f *field.MinBidSize) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidSize is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BidSize() (*field.BidSize, errors.MessageRejectError) {
	f := new(field.BidSize)
	err := m.Body.Get(f)
	return f, err
}

//GetBidSize reads a BidSize from QuoteStatusReport.
func (m QuoteStatusReport) GetBidSize(f *field.BidSize) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinOfferSize is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MinOfferSize() (*field.MinOfferSize, errors.MessageRejectError) {
	f := new(field.MinOfferSize)
	err := m.Body.Get(f)
	return f, err
}

//GetMinOfferSize reads a MinOfferSize from QuoteStatusReport.
func (m QuoteStatusReport) GetMinOfferSize(f *field.MinOfferSize) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OfferSize is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OfferSize() (*field.OfferSize, errors.MessageRejectError) {
	f := new(field.OfferSize)
	err := m.Body.Get(f)
	return f, err
}

//GetOfferSize reads a OfferSize from QuoteStatusReport.
func (m QuoteStatusReport) GetOfferSize(f *field.OfferSize) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ValidUntilTime is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ValidUntilTime() (*field.ValidUntilTime, errors.MessageRejectError) {
	f := new(field.ValidUntilTime)
	err := m.Body.Get(f)
	return f, err
}

//GetValidUntilTime reads a ValidUntilTime from QuoteStatusReport.
func (m QuoteStatusReport) GetValidUntilTime(f *field.ValidUntilTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidSpotRate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BidSpotRate() (*field.BidSpotRate, errors.MessageRejectError) {
	f := new(field.BidSpotRate)
	err := m.Body.Get(f)
	return f, err
}

//GetBidSpotRate reads a BidSpotRate from QuoteStatusReport.
func (m QuoteStatusReport) GetBidSpotRate(f *field.BidSpotRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OfferSpotRate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OfferSpotRate() (*field.OfferSpotRate, errors.MessageRejectError) {
	f := new(field.OfferSpotRate)
	err := m.Body.Get(f)
	return f, err
}

//GetOfferSpotRate reads a OfferSpotRate from QuoteStatusReport.
func (m QuoteStatusReport) GetOfferSpotRate(f *field.OfferSpotRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidForwardPoints is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BidForwardPoints() (*field.BidForwardPoints, errors.MessageRejectError) {
	f := new(field.BidForwardPoints)
	err := m.Body.Get(f)
	return f, err
}

//GetBidForwardPoints reads a BidForwardPoints from QuoteStatusReport.
func (m QuoteStatusReport) GetBidForwardPoints(f *field.BidForwardPoints) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OfferForwardPoints is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OfferForwardPoints() (*field.OfferForwardPoints, errors.MessageRejectError) {
	f := new(field.OfferForwardPoints)
	err := m.Body.Get(f)
	return f, err
}

//GetOfferForwardPoints reads a OfferForwardPoints from QuoteStatusReport.
func (m QuoteStatusReport) GetOfferForwardPoints(f *field.OfferForwardPoints) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MidPx is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MidPx() (*field.MidPx, errors.MessageRejectError) {
	f := new(field.MidPx)
	err := m.Body.Get(f)
	return f, err
}

//GetMidPx reads a MidPx from QuoteStatusReport.
func (m QuoteStatusReport) GetMidPx(f *field.MidPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidYield is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BidYield() (*field.BidYield, errors.MessageRejectError) {
	f := new(field.BidYield)
	err := m.Body.Get(f)
	return f, err
}

//GetBidYield reads a BidYield from QuoteStatusReport.
func (m QuoteStatusReport) GetBidYield(f *field.BidYield) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MidYield is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MidYield() (*field.MidYield, errors.MessageRejectError) {
	f := new(field.MidYield)
	err := m.Body.Get(f)
	return f, err
}

//GetMidYield reads a MidYield from QuoteStatusReport.
func (m QuoteStatusReport) GetMidYield(f *field.MidYield) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OfferYield is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OfferYield() (*field.OfferYield, errors.MessageRejectError) {
	f := new(field.OfferYield)
	err := m.Body.Get(f)
	return f, err
}

//GetOfferYield reads a OfferYield from QuoteStatusReport.
func (m QuoteStatusReport) GetOfferYield(f *field.OfferYield) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from QuoteStatusReport.
func (m QuoteStatusReport) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OrdType() (*field.OrdType, errors.MessageRejectError) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from QuoteStatusReport.
func (m QuoteStatusReport) GetOrdType(f *field.OrdType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidForwardPoints2 is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BidForwardPoints2() (*field.BidForwardPoints2, errors.MessageRejectError) {
	f := new(field.BidForwardPoints2)
	err := m.Body.Get(f)
	return f, err
}

//GetBidForwardPoints2 reads a BidForwardPoints2 from QuoteStatusReport.
func (m QuoteStatusReport) GetBidForwardPoints2(f *field.BidForwardPoints2) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OfferForwardPoints2 is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OfferForwardPoints2() (*field.OfferForwardPoints2, errors.MessageRejectError) {
	f := new(field.OfferForwardPoints2)
	err := m.Body.Get(f)
	return f, err
}

//GetOfferForwardPoints2 reads a OfferForwardPoints2 from QuoteStatusReport.
func (m QuoteStatusReport) GetOfferForwardPoints2(f *field.OfferForwardPoints2) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrBidFxRate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SettlCurrBidFxRate() (*field.SettlCurrBidFxRate, errors.MessageRejectError) {
	f := new(field.SettlCurrBidFxRate)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrBidFxRate reads a SettlCurrBidFxRate from QuoteStatusReport.
func (m QuoteStatusReport) GetSettlCurrBidFxRate(f *field.SettlCurrBidFxRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrOfferFxRate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SettlCurrOfferFxRate() (*field.SettlCurrOfferFxRate, errors.MessageRejectError) {
	f := new(field.SettlCurrOfferFxRate)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrOfferFxRate reads a SettlCurrOfferFxRate from QuoteStatusReport.
func (m QuoteStatusReport) GetSettlCurrOfferFxRate(f *field.SettlCurrOfferFxRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrFxRateCalc is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SettlCurrFxRateCalc() (*field.SettlCurrFxRateCalc, errors.MessageRejectError) {
	f := new(field.SettlCurrFxRateCalc)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrFxRateCalc reads a SettlCurrFxRateCalc from QuoteStatusReport.
func (m QuoteStatusReport) GetSettlCurrFxRateCalc(f *field.SettlCurrFxRateCalc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CommType() (*field.CommType, errors.MessageRejectError) {
	f := new(field.CommType)
	err := m.Body.Get(f)
	return f, err
}

//GetCommType reads a CommType from QuoteStatusReport.
func (m QuoteStatusReport) GetCommType(f *field.CommType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Commission is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Commission() (*field.Commission, errors.MessageRejectError) {
	f := new(field.Commission)
	err := m.Body.Get(f)
	return f, err
}

//GetCommission reads a Commission from QuoteStatusReport.
func (m QuoteStatusReport) GetCommission(f *field.Commission) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustOrderCapacity is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CustOrderCapacity() (*field.CustOrderCapacity, errors.MessageRejectError) {
	f := new(field.CustOrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//GetCustOrderCapacity reads a CustOrderCapacity from QuoteStatusReport.
func (m QuoteStatusReport) GetCustOrderCapacity(f *field.CustOrderCapacity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestination is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ExDestination() (*field.ExDestination, errors.MessageRejectError) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}

//GetExDestination reads a ExDestination from QuoteStatusReport.
func (m QuoteStatusReport) GetExDestination(f *field.ExDestination) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteStatus is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteStatus() (*field.QuoteStatus, errors.MessageRejectError) {
	f := new(field.QuoteStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteStatus reads a QuoteStatus from QuoteStatusReport.
func (m QuoteStatusReport) GetQuoteStatus(f *field.QuoteStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from QuoteStatusReport.
func (m QuoteStatusReport) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from QuoteStatusReport.
func (m QuoteStatusReport) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from QuoteStatusReport.
func (m QuoteStatusReport) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestinationIDSource is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ExDestinationIDSource() (*field.ExDestinationIDSource, errors.MessageRejectError) {
	f := new(field.ExDestinationIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetExDestinationIDSource reads a ExDestinationIDSource from QuoteStatusReport.
func (m QuoteStatusReport) GetExDestinationIDSource(f *field.ExDestinationIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteCancelType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteCancelType() (*field.QuoteCancelType, errors.MessageRejectError) {
	f := new(field.QuoteCancelType)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteCancelType reads a QuoteCancelType from QuoteStatusReport.
func (m QuoteStatusReport) GetQuoteCancelType(f *field.QuoteCancelType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteMsgID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteMsgID() (*field.QuoteMsgID, errors.MessageRejectError) {
	f := new(field.QuoteMsgID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteMsgID reads a QuoteMsgID from QuoteStatusReport.
func (m QuoteStatusReport) GetQuoteMsgID(f *field.QuoteMsgID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteRejectReason is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteRejectReason() (*field.QuoteRejectReason, errors.MessageRejectError) {
	f := new(field.QuoteRejectReason)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteRejectReason reads a QuoteRejectReason from QuoteStatusReport.
func (m QuoteStatusReport) GetQuoteRejectReason(f *field.QuoteRejectReason) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinQty is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MinQty() (*field.MinQty, errors.MessageRejectError) {
	f := new(field.MinQty)
	err := m.Body.Get(f)
	return f, err
}

//GetMinQty reads a MinQty from QuoteStatusReport.
func (m QuoteStatusReport) GetMinQty(f *field.MinQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BookingType() (*field.BookingType, errors.MessageRejectError) {
	f := new(field.BookingType)
	err := m.Body.Get(f)
	return f, err
}

//GetBookingType reads a BookingType from QuoteStatusReport.
func (m QuoteStatusReport) GetBookingType(f *field.BookingType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderCapacity is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OrderCapacity() (*field.OrderCapacity, errors.MessageRejectError) {
	f := new(field.OrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderCapacity reads a OrderCapacity from QuoteStatusReport.
func (m QuoteStatusReport) GetOrderCapacity(f *field.OrderCapacity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderRestrictions is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OrderRestrictions() (*field.OrderRestrictions, errors.MessageRejectError) {
	f := new(field.OrderRestrictions)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderRestrictions reads a OrderRestrictions from QuoteStatusReport.
func (m QuoteStatusReport) GetOrderRestrictions(f *field.OrderRestrictions) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTargetPartyIDs is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NoTargetPartyIDs() (*field.NoTargetPartyIDs, errors.MessageRejectError) {
	f := new(field.NoTargetPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTargetPartyIDs reads a NoTargetPartyIDs from QuoteStatusReport.
func (m QuoteStatusReport) GetNoTargetPartyIDs(f *field.NoTargetPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}
