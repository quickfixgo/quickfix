package fix44

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
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
	posmaintrptid *field.PosMaintRptIDField,
	postranstype *field.PosTransTypeField,
	posmaintaction *field.PosMaintActionField,
	origposreqrefid *field.OrigPosReqRefIDField,
	posmaintstatus *field.PosMaintStatusField,
	clearingbusinessdate *field.ClearingBusinessDateField,
	account *field.AccountField,
	accounttype *field.AccountTypeField,
	transacttime *field.TransactTimeField) PositionMaintenanceReportBuilder {
	var builder PositionMaintenanceReportBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header().Set(field.NewMsgType("AM"))
	builder.Body().Set(posmaintrptid)
	builder.Body().Set(postranstype)
	builder.Body().Set(posmaintaction)
	builder.Body().Set(origposreqrefid)
	builder.Body().Set(posmaintstatus)
	builder.Body().Set(clearingbusinessdate)
	builder.Body().Set(account)
	builder.Body().Set(accounttype)
	builder.Body().Set(transacttime)
	return builder
}

//PosMaintRptID is a required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) PosMaintRptID() (*field.PosMaintRptIDField, errors.MessageRejectError) {
	f := &field.PosMaintRptIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPosMaintRptID reads a PosMaintRptID from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetPosMaintRptID(f *field.PosMaintRptIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosTransType is a required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) PosTransType() (*field.PosTransTypeField, errors.MessageRejectError) {
	f := &field.PosTransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPosTransType reads a PosTransType from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetPosTransType(f *field.PosTransTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosReqID is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) PosReqID() (*field.PosReqIDField, errors.MessageRejectError) {
	f := &field.PosReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPosReqID reads a PosReqID from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetPosReqID(f *field.PosReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosMaintAction is a required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) PosMaintAction() (*field.PosMaintActionField, errors.MessageRejectError) {
	f := &field.PosMaintActionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPosMaintAction reads a PosMaintAction from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetPosMaintAction(f *field.PosMaintActionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigPosReqRefID is a required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) OrigPosReqRefID() (*field.OrigPosReqRefIDField, errors.MessageRejectError) {
	f := &field.OrigPosReqRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigPosReqRefID reads a OrigPosReqRefID from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetOrigPosReqRefID(f *field.OrigPosReqRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosMaintStatus is a required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) PosMaintStatus() (*field.PosMaintStatusField, errors.MessageRejectError) {
	f := &field.PosMaintStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPosMaintStatus reads a PosMaintStatus from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetPosMaintStatus(f *field.PosMaintStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosMaintResult is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) PosMaintResult() (*field.PosMaintResultField, errors.MessageRejectError) {
	f := &field.PosMaintResultField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPosMaintResult reads a PosMaintResult from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetPosMaintResult(f *field.PosMaintResultField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) ClearingBusinessDate() (*field.ClearingBusinessDateField, errors.MessageRejectError) {
	f := &field.ClearingBusinessDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetClearingBusinessDate(f *field.ClearingBusinessDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessID is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) SettlSessID() (*field.SettlSessIDField, errors.MessageRejectError) {
	f := &field.SettlSessIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessID reads a SettlSessID from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetSettlSessID(f *field.SettlSessIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessSubID is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) SettlSessSubID() (*field.SettlSessSubIDField, errors.MessageRejectError) {
	f := &field.SettlSessSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessSubID reads a SettlSessSubID from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetSettlSessSubID(f *field.SettlSessSubIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) Account() (*field.AccountField, errors.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetAccount(f *field.AccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) AcctIDSource() (*field.AcctIDSourceField, errors.MessageRejectError) {
	f := &field.AcctIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetAcctIDSource(f *field.AcctIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) AccountType() (*field.AccountTypeField, errors.MessageRejectError) {
	f := &field.AccountTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetAccountType(f *field.AccountTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) SecurityIDSource() (*field.SecurityIDSourceField, errors.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetSecurityIDSource(f *field.SecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) NoSecurityAltID() (*field.NoSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetNoSecurityAltID(f *field.NoSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) SecuritySubType() (*field.SecuritySubTypeField, errors.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetSecuritySubType(f *field.SecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) MaturityDate() (*field.MaturityDateField, errors.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetMaturityDate(f *field.MaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) CouponPaymentDate() (*field.CouponPaymentDateField, errors.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetCouponPaymentDate(f *field.CouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) IssueDate() (*field.IssueDateField, errors.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetIssueDate(f *field.IssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) RepurchaseTerm() (*field.RepurchaseTermField, errors.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetRepurchaseTerm(f *field.RepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) RepurchaseRate() (*field.RepurchaseRateField, errors.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetRepurchaseRate(f *field.RepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) Factor() (*field.FactorField, errors.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetFactor(f *field.FactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) CreditRating() (*field.CreditRatingField, errors.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetCreditRating(f *field.CreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) InstrRegistry() (*field.InstrRegistryField, errors.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetInstrRegistry(f *field.InstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) CountryOfIssue() (*field.CountryOfIssueField, errors.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetCountryOfIssue(f *field.CountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) LocaleOfIssue() (*field.LocaleOfIssueField, errors.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetLocaleOfIssue(f *field.LocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) RedemptionDate() (*field.RedemptionDateField, errors.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetRedemptionDate(f *field.RedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) StrikeCurrency() (*field.StrikeCurrencyField, errors.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetStrikeCurrency(f *field.StrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) Pool() (*field.PoolField, errors.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetPool(f *field.PoolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) ContractSettlMonth() (*field.ContractSettlMonthField, errors.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetContractSettlMonth(f *field.ContractSettlMonthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) CPProgram() (*field.CPProgramField, errors.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetCPProgram(f *field.CPProgramField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) CPRegType() (*field.CPRegTypeField, errors.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetCPRegType(f *field.CPRegTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) NoEvents() (*field.NoEventsField, errors.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetNoEvents(f *field.NoEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) DatedDate() (*field.DatedDateField, errors.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetDatedDate(f *field.DatedDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) InterestAccrualDate() (*field.InterestAccrualDateField, errors.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetInterestAccrualDate(f *field.InterestAccrualDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) NoLegs() (*field.NoLegsField, errors.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetNoLegs(f *field.NoLegsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) NoUnderlyings() (*field.NoUnderlyingsField, errors.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetNoUnderlyings(f *field.NoUnderlyingsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) NoTradingSessions() (*field.NoTradingSessionsField, errors.MessageRejectError) {
	f := &field.NoTradingSessionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetNoTradingSessions(f *field.NoTradingSessionsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPositions is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) NoPositions() (*field.NoPositionsField, errors.MessageRejectError) {
	f := &field.NoPositionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPositions reads a NoPositions from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetNoPositions(f *field.NoPositionsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPosAmt is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) NoPosAmt() (*field.NoPosAmtField, errors.MessageRejectError) {
	f := &field.NoPosAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPosAmt reads a NoPosAmt from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetNoPosAmt(f *field.NoPosAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AdjustmentType is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) AdjustmentType() (*field.AdjustmentTypeField, errors.MessageRejectError) {
	f := &field.AdjustmentTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAdjustmentType reads a AdjustmentType from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetAdjustmentType(f *field.AdjustmentTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ThresholdAmount is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) ThresholdAmount() (*field.ThresholdAmountField, errors.MessageRejectError) {
	f := &field.ThresholdAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetThresholdAmount reads a ThresholdAmount from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetThresholdAmount(f *field.ThresholdAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for PositionMaintenanceReport.
func (m PositionMaintenanceReport) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from PositionMaintenanceReport.
func (m PositionMaintenanceReport) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
