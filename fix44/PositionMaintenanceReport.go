package fix44

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//PositionMaintenanceReport msg type = AM.
type PositionMaintenanceReport struct {
	message.Message
}

//PositionMaintenanceReportBuilder builds PositionMaintenanceReport messages.
type PositionMaintenanceReportBuilder struct {
	message.MessageBuilder
}

//CreatePositionMaintenanceReportBuilder returns an initialized PositionMaintenanceReportBuilder with specified required fields.
func CreatePositionMaintenanceReportBuilder(
	posmaintrptid field.PosMaintRptID,
	postranstype field.PosTransType,
	posmaintaction field.PosMaintAction,
	origposreqrefid field.OrigPosReqRefID,
	posmaintstatus field.PosMaintStatus,
	clearingbusinessdate field.ClearingBusinessDate,
	account field.Account,
	accounttype field.AccountType,
	transacttime field.TransactTime) PositionMaintenanceReportBuilder {
	var builder PositionMaintenanceReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(posmaintrptid)
	builder.Body.Set(postranstype)
	builder.Body.Set(posmaintaction)
	builder.Body.Set(origposreqrefid)
	builder.Body.Set(posmaintstatus)
	builder.Body.Set(clearingbusinessdate)
	builder.Body.Set(account)
	builder.Body.Set(accounttype)
	builder.Body.Set(transacttime)
	return builder
}

//PosMaintRptID is a required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) PosMaintRptID() (*field.PosMaintRptID, errors.MessageRejectError) {
	f := new(field.PosMaintRptID)
	err := m.Body.Get(f)
	return f, err
}

//GetPosMaintRptID reads a PosMaintRptID from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetPosMaintRptID(f *field.PosMaintRptID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosTransType is a required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) PosTransType() (*field.PosTransType, errors.MessageRejectError) {
	f := new(field.PosTransType)
	err := m.Body.Get(f)
	return f, err
}

//GetPosTransType reads a PosTransType from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetPosTransType(f *field.PosTransType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosReqID is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) PosReqID() (*field.PosReqID, errors.MessageRejectError) {
	f := new(field.PosReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetPosReqID reads a PosReqID from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetPosReqID(f *field.PosReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosMaintAction is a required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) PosMaintAction() (*field.PosMaintAction, errors.MessageRejectError) {
	f := new(field.PosMaintAction)
	err := m.Body.Get(f)
	return f, err
}

//GetPosMaintAction reads a PosMaintAction from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetPosMaintAction(f *field.PosMaintAction) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigPosReqRefID is a required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) OrigPosReqRefID() (*field.OrigPosReqRefID, errors.MessageRejectError) {
	f := new(field.OrigPosReqRefID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrigPosReqRefID reads a OrigPosReqRefID from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetOrigPosReqRefID(f *field.OrigPosReqRefID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosMaintStatus is a required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) PosMaintStatus() (*field.PosMaintStatus, errors.MessageRejectError) {
	f := new(field.PosMaintStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetPosMaintStatus reads a PosMaintStatus from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetPosMaintStatus(f *field.PosMaintStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosMaintResult is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) PosMaintResult() (*field.PosMaintResult, errors.MessageRejectError) {
	f := new(field.PosMaintResult)
	err := m.Body.Get(f)
	return f, err
}

//GetPosMaintResult reads a PosMaintResult from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetPosMaintResult(f *field.PosMaintResult) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) ClearingBusinessDate() (*field.ClearingBusinessDate, errors.MessageRejectError) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetClearingBusinessDate(f *field.ClearingBusinessDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessID is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) SettlSessID() (*field.SettlSessID, errors.MessageRejectError) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessID reads a SettlSessID from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetSettlSessID(f *field.SettlSessID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessSubID is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) SettlSessSubID() (*field.SettlSessSubID, errors.MessageRejectError) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessSubID reads a SettlSessSubID from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetSettlSessSubID(f *field.SettlSessSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) Account() (*field.Account, errors.MessageRejectError) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetAccount(f *field.Account) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) AcctIDSource() (*field.AcctIDSource, errors.MessageRejectError) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetAcctIDSource(f *field.AcctIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) AccountType() (*field.AccountType, errors.MessageRejectError) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetAccountType(f *field.AccountType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) NoTradingSessions() (*field.NoTradingSessions, errors.MessageRejectError) {
	f := new(field.NoTradingSessions)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetNoTradingSessions(f *field.NoTradingSessions) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPositions is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) NoPositions() (*field.NoPositions, errors.MessageRejectError) {
	f := new(field.NoPositions)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPositions reads a NoPositions from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetNoPositions(f *field.NoPositions) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPosAmt is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) NoPosAmt() (*field.NoPosAmt, errors.MessageRejectError) {
	f := new(field.NoPosAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPosAmt reads a NoPosAmt from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetNoPosAmt(f *field.NoPosAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AdjustmentType is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) AdjustmentType() (*field.AdjustmentType, errors.MessageRejectError) {
	f := new(field.AdjustmentType)
	err := m.Body.Get(f)
	return f, err
}

//GetAdjustmentType reads a AdjustmentType from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetAdjustmentType(f *field.AdjustmentType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ThresholdAmount is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) ThresholdAmount() (*field.ThresholdAmount, errors.MessageRejectError) {
	f := new(field.ThresholdAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetThresholdAmount reads a ThresholdAmount from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetThresholdAmount(f *field.ThresholdAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}
