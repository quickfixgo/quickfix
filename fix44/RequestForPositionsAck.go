package fix44

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//RequestForPositionsAck msg type = AO.
type RequestForPositionsAck struct {
	message.Message
}

//RequestForPositionsAckBuilder builds RequestForPositionsAck messages.
type RequestForPositionsAckBuilder struct {
	message.MessageBuilder
}

//CreateRequestForPositionsAckBuilder returns an initialized RequestForPositionsAckBuilder with specified required fields.
func CreateRequestForPositionsAckBuilder(
	posmaintrptid field.PosMaintRptID,
	posreqresult field.PosReqResult,
	posreqstatus field.PosReqStatus,
	account field.Account,
	accounttype field.AccountType) RequestForPositionsAckBuilder {
	var builder RequestForPositionsAckBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("AO"))
	builder.Body.Set(posmaintrptid)
	builder.Body.Set(posreqresult)
	builder.Body.Set(posreqstatus)
	builder.Body.Set(account)
	builder.Body.Set(accounttype)
	return builder
}

//PosMaintRptID is a required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PosMaintRptID() (*field.PosMaintRptID, errors.MessageRejectError) {
	f := new(field.PosMaintRptID)
	err := m.Body.Get(f)
	return f, err
}

//GetPosMaintRptID reads a PosMaintRptID from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPosMaintRptID(f *field.PosMaintRptID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosReqID is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PosReqID() (*field.PosReqID, errors.MessageRejectError) {
	f := new(field.PosReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetPosReqID reads a PosReqID from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPosReqID(f *field.PosReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalNumPosReports is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) TotalNumPosReports() (*field.TotalNumPosReports, errors.MessageRejectError) {
	f := new(field.TotalNumPosReports)
	err := m.Body.Get(f)
	return f, err
}

//GetTotalNumPosReports reads a TotalNumPosReports from RequestForPositionsAck.
func (m RequestForPositionsAck) GetTotalNumPosReports(f *field.TotalNumPosReports) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnsolicitedIndicator is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) UnsolicitedIndicator() (*field.UnsolicitedIndicator, errors.MessageRejectError) {
	f := new(field.UnsolicitedIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetUnsolicitedIndicator reads a UnsolicitedIndicator from RequestForPositionsAck.
func (m RequestForPositionsAck) GetUnsolicitedIndicator(f *field.UnsolicitedIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosReqResult is a required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PosReqResult() (*field.PosReqResult, errors.MessageRejectError) {
	f := new(field.PosReqResult)
	err := m.Body.Get(f)
	return f, err
}

//GetPosReqResult reads a PosReqResult from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPosReqResult(f *field.PosReqResult) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosReqStatus is a required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PosReqStatus() (*field.PosReqStatus, errors.MessageRejectError) {
	f := new(field.PosReqStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetPosReqStatus reads a PosReqStatus from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPosReqStatus(f *field.PosReqStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from RequestForPositionsAck.
func (m RequestForPositionsAck) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Account() (*field.Account, errors.MessageRejectError) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from RequestForPositionsAck.
func (m RequestForPositionsAck) GetAccount(f *field.Account) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) AcctIDSource() (*field.AcctIDSource, errors.MessageRejectError) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from RequestForPositionsAck.
func (m RequestForPositionsAck) GetAcctIDSource(f *field.AcctIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a required field for RequestForPositionsAck.
func (m RequestForPositionsAck) AccountType() (*field.AccountType, errors.MessageRejectError) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from RequestForPositionsAck.
func (m RequestForPositionsAck) GetAccountType(f *field.AccountType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from RequestForPositionsAck.
func (m RequestForPositionsAck) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from RequestForPositionsAck.
func (m RequestForPositionsAck) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from RequestForPositionsAck.
func (m RequestForPositionsAck) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from RequestForPositionsAck.
func (m RequestForPositionsAck) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from RequestForPositionsAck.
func (m RequestForPositionsAck) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from RequestForPositionsAck.
func (m RequestForPositionsAck) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from RequestForPositionsAck.
func (m RequestForPositionsAck) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from RequestForPositionsAck.
func (m RequestForPositionsAck) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from RequestForPositionsAck.
func (m RequestForPositionsAck) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from RequestForPositionsAck.
func (m RequestForPositionsAck) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from RequestForPositionsAck.
func (m RequestForPositionsAck) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from RequestForPositionsAck.
func (m RequestForPositionsAck) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from RequestForPositionsAck.
func (m RequestForPositionsAck) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from RequestForPositionsAck.
func (m RequestForPositionsAck) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from RequestForPositionsAck.
func (m RequestForPositionsAck) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from RequestForPositionsAck.
func (m RequestForPositionsAck) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from RequestForPositionsAck.
func (m RequestForPositionsAck) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from RequestForPositionsAck.
func (m RequestForPositionsAck) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from RequestForPositionsAck.
func (m RequestForPositionsAck) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from RequestForPositionsAck.
func (m RequestForPositionsAck) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from RequestForPositionsAck.
func (m RequestForPositionsAck) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from RequestForPositionsAck.
func (m RequestForPositionsAck) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from RequestForPositionsAck.
func (m RequestForPositionsAck) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from RequestForPositionsAck.
func (m RequestForPositionsAck) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from RequestForPositionsAck.
func (m RequestForPositionsAck) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from RequestForPositionsAck.
func (m RequestForPositionsAck) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from RequestForPositionsAck.
func (m RequestForPositionsAck) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from RequestForPositionsAck.
func (m RequestForPositionsAck) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from RequestForPositionsAck.
func (m RequestForPositionsAck) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from RequestForPositionsAck.
func (m RequestForPositionsAck) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from RequestForPositionsAck.
func (m RequestForPositionsAck) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from RequestForPositionsAck.
func (m RequestForPositionsAck) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from RequestForPositionsAck.
func (m RequestForPositionsAck) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from RequestForPositionsAck.
func (m RequestForPositionsAck) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from RequestForPositionsAck.
func (m RequestForPositionsAck) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from RequestForPositionsAck.
func (m RequestForPositionsAck) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ResponseTransportType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ResponseTransportType() (*field.ResponseTransportType, errors.MessageRejectError) {
	f := new(field.ResponseTransportType)
	err := m.Body.Get(f)
	return f, err
}

//GetResponseTransportType reads a ResponseTransportType from RequestForPositionsAck.
func (m RequestForPositionsAck) GetResponseTransportType(f *field.ResponseTransportType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ResponseDestination is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ResponseDestination() (*field.ResponseDestination, errors.MessageRejectError) {
	f := new(field.ResponseDestination)
	err := m.Body.Get(f)
	return f, err
}

//GetResponseDestination reads a ResponseDestination from RequestForPositionsAck.
func (m RequestForPositionsAck) GetResponseDestination(f *field.ResponseDestination) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from RequestForPositionsAck.
func (m RequestForPositionsAck) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from RequestForPositionsAck.
func (m RequestForPositionsAck) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from RequestForPositionsAck.
func (m RequestForPositionsAck) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}
