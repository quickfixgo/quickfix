package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
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
	posmaintrptid *field.PosMaintRptIDField,
	posreqresult *field.PosReqResultField,
	posreqstatus *field.PosReqStatusField) RequestForPositionsAckBuilder {
	var builder RequestForPositionsAckBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header.Set(field.NewMsgType("AO"))
	builder.Body.Set(posmaintrptid)
	builder.Body.Set(posreqresult)
	builder.Body.Set(posreqstatus)
	return builder
}

//PosMaintRptID is a required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PosMaintRptID() (*field.PosMaintRptIDField, errors.MessageRejectError) {
	f := &field.PosMaintRptIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPosMaintRptID reads a PosMaintRptID from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPosMaintRptID(f *field.PosMaintRptIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosReqID is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PosReqID() (*field.PosReqIDField, errors.MessageRejectError) {
	f := &field.PosReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPosReqID reads a PosReqID from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPosReqID(f *field.PosReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalNumPosReports is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) TotalNumPosReports() (*field.TotalNumPosReportsField, errors.MessageRejectError) {
	f := &field.TotalNumPosReportsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotalNumPosReports reads a TotalNumPosReports from RequestForPositionsAck.
func (m RequestForPositionsAck) GetTotalNumPosReports(f *field.TotalNumPosReportsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnsolicitedIndicator is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) UnsolicitedIndicator() (*field.UnsolicitedIndicatorField, errors.MessageRejectError) {
	f := &field.UnsolicitedIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnsolicitedIndicator reads a UnsolicitedIndicator from RequestForPositionsAck.
func (m RequestForPositionsAck) GetUnsolicitedIndicator(f *field.UnsolicitedIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosReqResult is a required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PosReqResult() (*field.PosReqResultField, errors.MessageRejectError) {
	f := &field.PosReqResultField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPosReqResult reads a PosReqResult from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPosReqResult(f *field.PosReqResultField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosReqStatus is a required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PosReqStatus() (*field.PosReqStatusField, errors.MessageRejectError) {
	f := &field.PosReqStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPosReqStatus reads a PosReqStatus from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPosReqStatus(f *field.PosReqStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from RequestForPositionsAck.
func (m RequestForPositionsAck) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Account() (*field.AccountField, errors.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from RequestForPositionsAck.
func (m RequestForPositionsAck) GetAccount(f *field.AccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) AcctIDSource() (*field.AcctIDSourceField, errors.MessageRejectError) {
	f := &field.AcctIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from RequestForPositionsAck.
func (m RequestForPositionsAck) GetAcctIDSource(f *field.AcctIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) AccountType() (*field.AccountTypeField, errors.MessageRejectError) {
	f := &field.AccountTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from RequestForPositionsAck.
func (m RequestForPositionsAck) GetAccountType(f *field.AccountTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityIDSource() (*field.SecurityIDSourceField, errors.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecurityIDSource(f *field.SecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NoSecurityAltID() (*field.NoSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from RequestForPositionsAck.
func (m RequestForPositionsAck) GetNoSecurityAltID(f *field.NoSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from RequestForPositionsAck.
func (m RequestForPositionsAck) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from RequestForPositionsAck.
func (m RequestForPositionsAck) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecuritySubType() (*field.SecuritySubTypeField, errors.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecuritySubType(f *field.SecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from RequestForPositionsAck.
func (m RequestForPositionsAck) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) MaturityDate() (*field.MaturityDateField, errors.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from RequestForPositionsAck.
func (m RequestForPositionsAck) GetMaturityDate(f *field.MaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CouponPaymentDate() (*field.CouponPaymentDateField, errors.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from RequestForPositionsAck.
func (m RequestForPositionsAck) GetCouponPaymentDate(f *field.CouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) IssueDate() (*field.IssueDateField, errors.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from RequestForPositionsAck.
func (m RequestForPositionsAck) GetIssueDate(f *field.IssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from RequestForPositionsAck.
func (m RequestForPositionsAck) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) RepurchaseTerm() (*field.RepurchaseTermField, errors.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from RequestForPositionsAck.
func (m RequestForPositionsAck) GetRepurchaseTerm(f *field.RepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) RepurchaseRate() (*field.RepurchaseRateField, errors.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from RequestForPositionsAck.
func (m RequestForPositionsAck) GetRepurchaseRate(f *field.RepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Factor() (*field.FactorField, errors.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from RequestForPositionsAck.
func (m RequestForPositionsAck) GetFactor(f *field.FactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CreditRating() (*field.CreditRatingField, errors.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from RequestForPositionsAck.
func (m RequestForPositionsAck) GetCreditRating(f *field.CreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) InstrRegistry() (*field.InstrRegistryField, errors.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from RequestForPositionsAck.
func (m RequestForPositionsAck) GetInstrRegistry(f *field.InstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CountryOfIssue() (*field.CountryOfIssueField, errors.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from RequestForPositionsAck.
func (m RequestForPositionsAck) GetCountryOfIssue(f *field.CountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from RequestForPositionsAck.
func (m RequestForPositionsAck) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) LocaleOfIssue() (*field.LocaleOfIssueField, errors.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from RequestForPositionsAck.
func (m RequestForPositionsAck) GetLocaleOfIssue(f *field.LocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) RedemptionDate() (*field.RedemptionDateField, errors.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from RequestForPositionsAck.
func (m RequestForPositionsAck) GetRedemptionDate(f *field.RedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from RequestForPositionsAck.
func (m RequestForPositionsAck) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) StrikeCurrency() (*field.StrikeCurrencyField, errors.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from RequestForPositionsAck.
func (m RequestForPositionsAck) GetStrikeCurrency(f *field.StrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from RequestForPositionsAck.
func (m RequestForPositionsAck) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from RequestForPositionsAck.
func (m RequestForPositionsAck) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from RequestForPositionsAck.
func (m RequestForPositionsAck) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from RequestForPositionsAck.
func (m RequestForPositionsAck) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from RequestForPositionsAck.
func (m RequestForPositionsAck) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from RequestForPositionsAck.
func (m RequestForPositionsAck) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from RequestForPositionsAck.
func (m RequestForPositionsAck) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from RequestForPositionsAck.
func (m RequestForPositionsAck) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Pool() (*field.PoolField, errors.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPool(f *field.PoolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ContractSettlMonth() (*field.ContractSettlMonthField, errors.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from RequestForPositionsAck.
func (m RequestForPositionsAck) GetContractSettlMonth(f *field.ContractSettlMonthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CPProgram() (*field.CPProgramField, errors.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from RequestForPositionsAck.
func (m RequestForPositionsAck) GetCPProgram(f *field.CPProgramField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CPRegType() (*field.CPRegTypeField, errors.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from RequestForPositionsAck.
func (m RequestForPositionsAck) GetCPRegType(f *field.CPRegTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NoEvents() (*field.NoEventsField, errors.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from RequestForPositionsAck.
func (m RequestForPositionsAck) GetNoEvents(f *field.NoEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) DatedDate() (*field.DatedDateField, errors.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from RequestForPositionsAck.
func (m RequestForPositionsAck) GetDatedDate(f *field.DatedDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) InterestAccrualDate() (*field.InterestAccrualDateField, errors.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from RequestForPositionsAck.
func (m RequestForPositionsAck) GetInterestAccrualDate(f *field.InterestAccrualDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityStatus() (*field.SecurityStatusField, errors.MessageRejectError) {
	f := &field.SecurityStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecurityStatus(f *field.SecurityStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SettleOnOpenFlag() (*field.SettleOnOpenFlagField, errors.MessageRejectError) {
	f := &field.SettleOnOpenFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSettleOnOpenFlag(f *field.SettleOnOpenFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethodField, errors.MessageRejectError) {
	f := &field.InstrmtAssignmentMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from RequestForPositionsAck.
func (m RequestForPositionsAck) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) StrikeMultiplier() (*field.StrikeMultiplierField, errors.MessageRejectError) {
	f := &field.StrikeMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from RequestForPositionsAck.
func (m RequestForPositionsAck) GetStrikeMultiplier(f *field.StrikeMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) StrikeValue() (*field.StrikeValueField, errors.MessageRejectError) {
	f := &field.StrikeValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from RequestForPositionsAck.
func (m RequestForPositionsAck) GetStrikeValue(f *field.StrikeValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) MinPriceIncrement() (*field.MinPriceIncrementField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from RequestForPositionsAck.
func (m RequestForPositionsAck) GetMinPriceIncrement(f *field.MinPriceIncrementField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PositionLimit() (*field.PositionLimitField, errors.MessageRejectError) {
	f := &field.PositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPositionLimit(f *field.PositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NTPositionLimit() (*field.NTPositionLimitField, errors.MessageRejectError) {
	f := &field.NTPositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from RequestForPositionsAck.
func (m RequestForPositionsAck) GetNTPositionLimit(f *field.NTPositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NoInstrumentParties() (*field.NoInstrumentPartiesField, errors.MessageRejectError) {
	f := &field.NoInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from RequestForPositionsAck.
func (m RequestForPositionsAck) GetNoInstrumentParties(f *field.NoInstrumentPartiesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) UnitOfMeasure() (*field.UnitOfMeasureField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from RequestForPositionsAck.
func (m RequestForPositionsAck) GetUnitOfMeasure(f *field.UnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) TimeUnit() (*field.TimeUnitField, errors.MessageRejectError) {
	f := &field.TimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from RequestForPositionsAck.
func (m RequestForPositionsAck) GetTimeUnit(f *field.TimeUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) MaturityTime() (*field.MaturityTimeField, errors.MessageRejectError) {
	f := &field.MaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from RequestForPositionsAck.
func (m RequestForPositionsAck) GetMaturityTime(f *field.MaturityTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityGroup() (*field.SecurityGroupField, errors.MessageRejectError) {
	f := &field.SecurityGroupField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecurityGroup(f *field.SecurityGroupField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) MinPriceIncrementAmount() (*field.MinPriceIncrementAmountField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from RequestForPositionsAck.
func (m RequestForPositionsAck) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) UnitOfMeasureQty() (*field.UnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from RequestForPositionsAck.
func (m RequestForPositionsAck) GetUnitOfMeasureQty(f *field.UnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityXMLLen() (*field.SecurityXMLLenField, errors.MessageRejectError) {
	f := &field.SecurityXMLLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecurityXMLLen(f *field.SecurityXMLLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityXML() (*field.SecurityXMLField, errors.MessageRejectError) {
	f := &field.SecurityXMLField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecurityXML(f *field.SecurityXMLField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SecurityXMLSchema() (*field.SecurityXMLSchemaField, errors.MessageRejectError) {
	f := &field.SecurityXMLSchemaField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSecurityXMLSchema(f *field.SecurityXMLSchemaField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ProductComplex() (*field.ProductComplexField, errors.MessageRejectError) {
	f := &field.ProductComplexField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from RequestForPositionsAck.
func (m RequestForPositionsAck) GetProductComplex(f *field.ProductComplexField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PriceUnitOfMeasure() (*field.PriceUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.PriceUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.PriceUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SettlMethod() (*field.SettlMethodField, errors.MessageRejectError) {
	f := &field.SettlMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSettlMethod(f *field.SettlMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ExerciseStyle() (*field.ExerciseStyleField, errors.MessageRejectError) {
	f := &field.ExerciseStyleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from RequestForPositionsAck.
func (m RequestForPositionsAck) GetExerciseStyle(f *field.ExerciseStyleField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayAmount is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) OptPayAmount() (*field.OptPayAmountField, errors.MessageRejectError) {
	f := &field.OptPayAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayAmount reads a OptPayAmount from RequestForPositionsAck.
func (m RequestForPositionsAck) GetOptPayAmount(f *field.OptPayAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PriceQuoteMethod() (*field.PriceQuoteMethodField, errors.MessageRejectError) {
	f := &field.PriceQuoteMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPriceQuoteMethod(f *field.PriceQuoteMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ListMethod() (*field.ListMethodField, errors.MessageRejectError) {
	f := &field.ListMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from RequestForPositionsAck.
func (m RequestForPositionsAck) GetListMethod(f *field.ListMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) CapPrice() (*field.CapPriceField, errors.MessageRejectError) {
	f := &field.CapPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from RequestForPositionsAck.
func (m RequestForPositionsAck) GetCapPrice(f *field.CapPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) FloorPrice() (*field.FloorPriceField, errors.MessageRejectError) {
	f := &field.FloorPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from RequestForPositionsAck.
func (m RequestForPositionsAck) GetFloorPrice(f *field.FloorPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PutOrCall() (*field.PutOrCallField, errors.MessageRejectError) {
	f := &field.PutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPutOrCall(f *field.PutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) FlexibleIndicator() (*field.FlexibleIndicatorField, errors.MessageRejectError) {
	f := &field.FlexibleIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from RequestForPositionsAck.
func (m RequestForPositionsAck) GetFlexibleIndicator(f *field.FlexibleIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicatorField, errors.MessageRejectError) {
	f := &field.FlexProductEligibilityIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from RequestForPositionsAck.
func (m RequestForPositionsAck) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FuturesValuationMethod is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) FuturesValuationMethod() (*field.FuturesValuationMethodField, errors.MessageRejectError) {
	f := &field.FuturesValuationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFuturesValuationMethod reads a FuturesValuationMethod from RequestForPositionsAck.
func (m RequestForPositionsAck) GetFuturesValuationMethod(f *field.FuturesValuationMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from RequestForPositionsAck.
func (m RequestForPositionsAck) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NoLegs() (*field.NoLegsField, errors.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from RequestForPositionsAck.
func (m RequestForPositionsAck) GetNoLegs(f *field.NoLegsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) NoUnderlyings() (*field.NoUnderlyingsField, errors.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from RequestForPositionsAck.
func (m RequestForPositionsAck) GetNoUnderlyings(f *field.NoUnderlyingsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ResponseTransportType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ResponseTransportType() (*field.ResponseTransportTypeField, errors.MessageRejectError) {
	f := &field.ResponseTransportTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetResponseTransportType reads a ResponseTransportType from RequestForPositionsAck.
func (m RequestForPositionsAck) GetResponseTransportType(f *field.ResponseTransportTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ResponseDestination is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ResponseDestination() (*field.ResponseDestinationField, errors.MessageRejectError) {
	f := &field.ResponseDestinationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetResponseDestination reads a ResponseDestination from RequestForPositionsAck.
func (m RequestForPositionsAck) GetResponseDestination(f *field.ResponseDestinationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from RequestForPositionsAck.
func (m RequestForPositionsAck) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from RequestForPositionsAck.
func (m RequestForPositionsAck) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from RequestForPositionsAck.
func (m RequestForPositionsAck) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosReqType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) PosReqType() (*field.PosReqTypeField, errors.MessageRejectError) {
	f := &field.PosReqTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPosReqType reads a PosReqType from RequestForPositionsAck.
func (m RequestForPositionsAck) GetPosReqType(f *field.PosReqTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchStatus is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) MatchStatus() (*field.MatchStatusField, errors.MessageRejectError) {
	f := &field.MatchStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMatchStatus reads a MatchStatus from RequestForPositionsAck.
func (m RequestForPositionsAck) GetMatchStatus(f *field.MatchStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) ClearingBusinessDate() (*field.ClearingBusinessDateField, errors.MessageRejectError) {
	f := &field.ClearingBusinessDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from RequestForPositionsAck.
func (m RequestForPositionsAck) GetClearingBusinessDate(f *field.ClearingBusinessDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SubscriptionRequestType() (*field.SubscriptionRequestTypeField, errors.MessageRejectError) {
	f := &field.SubscriptionRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSubscriptionRequestType(f *field.SubscriptionRequestTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessID is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SettlSessID() (*field.SettlSessIDField, errors.MessageRejectError) {
	f := &field.SettlSessIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessID reads a SettlSessID from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSettlSessID(f *field.SettlSessIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessSubID is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SettlSessSubID() (*field.SettlSessSubIDField, errors.MessageRejectError) {
	f := &field.SettlSessSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessSubID reads a SettlSessSubID from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSettlSessSubID(f *field.SettlSessSubIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for RequestForPositionsAck.
func (m RequestForPositionsAck) SettlCurrency() (*field.SettlCurrencyField, errors.MessageRejectError) {
	f := &field.SettlCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from RequestForPositionsAck.
func (m RequestForPositionsAck) GetSettlCurrency(f *field.SettlCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}
