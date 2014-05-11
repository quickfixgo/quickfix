package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//OrderMassStatusRequest msg type = AF.
type OrderMassStatusRequest struct {
	message.Message
}

//OrderMassStatusRequestBuilder builds OrderMassStatusRequest messages.
type OrderMassStatusRequestBuilder struct {
	message.MessageBuilder
}

//CreateOrderMassStatusRequestBuilder returns an initialized OrderMassStatusRequestBuilder with specified required fields.
func CreateOrderMassStatusRequestBuilder(
	massstatusreqid *field.MassStatusReqIDField,
	massstatusreqtype *field.MassStatusReqTypeField) OrderMassStatusRequestBuilder {
	var builder OrderMassStatusRequestBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header().Set(field.NewMsgType("AF"))
	builder.Body().Set(massstatusreqid)
	builder.Body().Set(massstatusreqtype)
	return builder
}

//MassStatusReqID is a required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) MassStatusReqID() (*field.MassStatusReqIDField, errors.MessageRejectError) {
	f := &field.MassStatusReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMassStatusReqID reads a MassStatusReqID from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetMassStatusReqID(f *field.MassStatusReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MassStatusReqType is a required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) MassStatusReqType() (*field.MassStatusReqTypeField, errors.MessageRejectError) {
	f := &field.MassStatusReqTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMassStatusReqType reads a MassStatusReqType from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetMassStatusReqType(f *field.MassStatusReqTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) Account() (*field.AccountField, errors.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetAccount(f *field.AccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) AcctIDSource() (*field.AcctIDSourceField, errors.MessageRejectError) {
	f := &field.AcctIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetAcctIDSource(f *field.AcctIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) TradingSessionID() (*field.TradingSessionIDField, errors.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetTradingSessionID(f *field.TradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) TradingSessionSubID() (*field.TradingSessionSubIDField, errors.MessageRejectError) {
	f := &field.TradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetTradingSessionSubID(f *field.TradingSessionSubIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SecurityIDSource() (*field.SecurityIDSourceField, errors.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetSecurityIDSource(f *field.SecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) NoSecurityAltID() (*field.NoSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetNoSecurityAltID(f *field.NoSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SecuritySubType() (*field.SecuritySubTypeField, errors.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetSecuritySubType(f *field.SecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) MaturityDate() (*field.MaturityDateField, errors.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetMaturityDate(f *field.MaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) CouponPaymentDate() (*field.CouponPaymentDateField, errors.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetCouponPaymentDate(f *field.CouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) IssueDate() (*field.IssueDateField, errors.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetIssueDate(f *field.IssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) RepurchaseTerm() (*field.RepurchaseTermField, errors.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetRepurchaseTerm(f *field.RepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) RepurchaseRate() (*field.RepurchaseRateField, errors.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetRepurchaseRate(f *field.RepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) Factor() (*field.FactorField, errors.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetFactor(f *field.FactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) CreditRating() (*field.CreditRatingField, errors.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetCreditRating(f *field.CreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) InstrRegistry() (*field.InstrRegistryField, errors.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetInstrRegistry(f *field.InstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) CountryOfIssue() (*field.CountryOfIssueField, errors.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetCountryOfIssue(f *field.CountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) LocaleOfIssue() (*field.LocaleOfIssueField, errors.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetLocaleOfIssue(f *field.LocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) RedemptionDate() (*field.RedemptionDateField, errors.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetRedemptionDate(f *field.RedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) StrikeCurrency() (*field.StrikeCurrencyField, errors.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetStrikeCurrency(f *field.StrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) Pool() (*field.PoolField, errors.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetPool(f *field.PoolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) ContractSettlMonth() (*field.ContractSettlMonthField, errors.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetContractSettlMonth(f *field.ContractSettlMonthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) CPProgram() (*field.CPProgramField, errors.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetCPProgram(f *field.CPProgramField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) CPRegType() (*field.CPRegTypeField, errors.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetCPRegType(f *field.CPRegTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) NoEvents() (*field.NoEventsField, errors.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetNoEvents(f *field.NoEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) DatedDate() (*field.DatedDateField, errors.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetDatedDate(f *field.DatedDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) InterestAccrualDate() (*field.InterestAccrualDateField, errors.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetInterestAccrualDate(f *field.InterestAccrualDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SecurityStatus() (*field.SecurityStatusField, errors.MessageRejectError) {
	f := &field.SecurityStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetSecurityStatus(f *field.SecurityStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SettleOnOpenFlag() (*field.SettleOnOpenFlagField, errors.MessageRejectError) {
	f := &field.SettleOnOpenFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetSettleOnOpenFlag(f *field.SettleOnOpenFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethodField, errors.MessageRejectError) {
	f := &field.InstrmtAssignmentMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) StrikeMultiplier() (*field.StrikeMultiplierField, errors.MessageRejectError) {
	f := &field.StrikeMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetStrikeMultiplier(f *field.StrikeMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) StrikeValue() (*field.StrikeValueField, errors.MessageRejectError) {
	f := &field.StrikeValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetStrikeValue(f *field.StrikeValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) MinPriceIncrement() (*field.MinPriceIncrementField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetMinPriceIncrement(f *field.MinPriceIncrementField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) PositionLimit() (*field.PositionLimitField, errors.MessageRejectError) {
	f := &field.PositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetPositionLimit(f *field.PositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) NTPositionLimit() (*field.NTPositionLimitField, errors.MessageRejectError) {
	f := &field.NTPositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetNTPositionLimit(f *field.NTPositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) NoInstrumentParties() (*field.NoInstrumentPartiesField, errors.MessageRejectError) {
	f := &field.NoInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetNoInstrumentParties(f *field.NoInstrumentPartiesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnitOfMeasure() (*field.UnitOfMeasureField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnitOfMeasure(f *field.UnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) TimeUnit() (*field.TimeUnitField, errors.MessageRejectError) {
	f := &field.TimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetTimeUnit(f *field.TimeUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) MaturityTime() (*field.MaturityTimeField, errors.MessageRejectError) {
	f := &field.MaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetMaturityTime(f *field.MaturityTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SecurityGroup() (*field.SecurityGroupField, errors.MessageRejectError) {
	f := &field.SecurityGroupField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetSecurityGroup(f *field.SecurityGroupField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) MinPriceIncrementAmount() (*field.MinPriceIncrementAmountField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnitOfMeasureQty() (*field.UnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnitOfMeasureQty(f *field.UnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SecurityXMLLen() (*field.SecurityXMLLenField, errors.MessageRejectError) {
	f := &field.SecurityXMLLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetSecurityXMLLen(f *field.SecurityXMLLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SecurityXML() (*field.SecurityXMLField, errors.MessageRejectError) {
	f := &field.SecurityXMLField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetSecurityXML(f *field.SecurityXMLField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SecurityXMLSchema() (*field.SecurityXMLSchemaField, errors.MessageRejectError) {
	f := &field.SecurityXMLSchemaField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetSecurityXMLSchema(f *field.SecurityXMLSchemaField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) ProductComplex() (*field.ProductComplexField, errors.MessageRejectError) {
	f := &field.ProductComplexField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetProductComplex(f *field.ProductComplexField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) PriceUnitOfMeasure() (*field.PriceUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.PriceUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.PriceUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SettlMethod() (*field.SettlMethodField, errors.MessageRejectError) {
	f := &field.SettlMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetSettlMethod(f *field.SettlMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) ExerciseStyle() (*field.ExerciseStyleField, errors.MessageRejectError) {
	f := &field.ExerciseStyleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetExerciseStyle(f *field.ExerciseStyleField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutAmount is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) OptPayoutAmount() (*field.OptPayoutAmountField, errors.MessageRejectError) {
	f := &field.OptPayoutAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutAmount reads a OptPayoutAmount from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetOptPayoutAmount(f *field.OptPayoutAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) PriceQuoteMethod() (*field.PriceQuoteMethodField, errors.MessageRejectError) {
	f := &field.PriceQuoteMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetPriceQuoteMethod(f *field.PriceQuoteMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) ListMethod() (*field.ListMethodField, errors.MessageRejectError) {
	f := &field.ListMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetListMethod(f *field.ListMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) CapPrice() (*field.CapPriceField, errors.MessageRejectError) {
	f := &field.CapPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetCapPrice(f *field.CapPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) FloorPrice() (*field.FloorPriceField, errors.MessageRejectError) {
	f := &field.FloorPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetFloorPrice(f *field.FloorPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) PutOrCall() (*field.PutOrCallField, errors.MessageRejectError) {
	f := &field.PutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetPutOrCall(f *field.PutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) FlexibleIndicator() (*field.FlexibleIndicatorField, errors.MessageRejectError) {
	f := &field.FlexibleIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetFlexibleIndicator(f *field.FlexibleIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicatorField, errors.MessageRejectError) {
	f := &field.FlexProductEligibilityIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ValuationMethod is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) ValuationMethod() (*field.ValuationMethodField, errors.MessageRejectError) {
	f := &field.ValuationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetValuationMethod reads a ValuationMethod from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetValuationMethod(f *field.ValuationMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplierUnit is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) ContractMultiplierUnit() (*field.ContractMultiplierUnitField, errors.MessageRejectError) {
	f := &field.ContractMultiplierUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplierUnit reads a ContractMultiplierUnit from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetContractMultiplierUnit(f *field.ContractMultiplierUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlowScheduleType is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) FlowScheduleType() (*field.FlowScheduleTypeField, errors.MessageRejectError) {
	f := &field.FlowScheduleTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlowScheduleType reads a FlowScheduleType from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetFlowScheduleType(f *field.FlowScheduleTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RestructuringType is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) RestructuringType() (*field.RestructuringTypeField, errors.MessageRejectError) {
	f := &field.RestructuringTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRestructuringType reads a RestructuringType from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetRestructuringType(f *field.RestructuringTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Seniority is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) Seniority() (*field.SeniorityField, errors.MessageRejectError) {
	f := &field.SeniorityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSeniority reads a Seniority from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetSeniority(f *field.SeniorityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NotionalPercentageOutstanding is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstandingField, errors.MessageRejectError) {
	f := &field.NotionalPercentageOutstandingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNotionalPercentageOutstanding reads a NotionalPercentageOutstanding from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetNotionalPercentageOutstanding(f *field.NotionalPercentageOutstandingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OriginalNotionalPercentageOutstanding is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstandingField, errors.MessageRejectError) {
	f := &field.OriginalNotionalPercentageOutstandingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOriginalNotionalPercentageOutstanding reads a OriginalNotionalPercentageOutstanding from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetOriginalNotionalPercentageOutstanding(f *field.OriginalNotionalPercentageOutstandingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AttachmentPoint is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) AttachmentPoint() (*field.AttachmentPointField, errors.MessageRejectError) {
	f := &field.AttachmentPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAttachmentPoint reads a AttachmentPoint from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetAttachmentPoint(f *field.AttachmentPointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DetachmentPoint is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) DetachmentPoint() (*field.DetachmentPointField, errors.MessageRejectError) {
	f := &field.DetachmentPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDetachmentPoint reads a DetachmentPoint from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetDetachmentPoint(f *field.DetachmentPointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceDeterminationMethod is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethodField, errors.MessageRejectError) {
	f := &field.StrikePriceDeterminationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceDeterminationMethod reads a StrikePriceDeterminationMethod from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetStrikePriceDeterminationMethod(f *field.StrikePriceDeterminationMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryMethod is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethodField, errors.MessageRejectError) {
	f := &field.StrikePriceBoundaryMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryMethod reads a StrikePriceBoundaryMethod from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetStrikePriceBoundaryMethod(f *field.StrikePriceBoundaryMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryPrecision is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecisionField, errors.MessageRejectError) {
	f := &field.StrikePriceBoundaryPrecisionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryPrecision reads a StrikePriceBoundaryPrecision from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetStrikePriceBoundaryPrecision(f *field.StrikePriceBoundaryPrecisionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceDeterminationMethod is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethodField, errors.MessageRejectError) {
	f := &field.UnderlyingPriceDeterminationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceDeterminationMethod reads a UnderlyingPriceDeterminationMethod from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingPriceDeterminationMethod(f *field.UnderlyingPriceDeterminationMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutType is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) OptPayoutType() (*field.OptPayoutTypeField, errors.MessageRejectError) {
	f := &field.OptPayoutTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutType reads a OptPayoutType from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetOptPayoutType(f *field.OptPayoutTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoComplexEvents is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) NoComplexEvents() (*field.NoComplexEventsField, errors.MessageRejectError) {
	f := &field.NoComplexEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoComplexEvents reads a NoComplexEvents from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetNoComplexEvents(f *field.NoComplexEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSymbol is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingSymbol() (*field.UnderlyingSymbolField, errors.MessageRejectError) {
	f := &field.UnderlyingSymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSymbol reads a UnderlyingSymbol from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingSymbol(f *field.UnderlyingSymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSymbolSfx is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingSymbolSfx() (*field.UnderlyingSymbolSfxField, errors.MessageRejectError) {
	f := &field.UnderlyingSymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSymbolSfx reads a UnderlyingSymbolSfx from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingSymbolSfx(f *field.UnderlyingSymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityID is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingSecurityID() (*field.UnderlyingSecurityIDField, errors.MessageRejectError) {
	f := &field.UnderlyingSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityID reads a UnderlyingSecurityID from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingSecurityID(f *field.UnderlyingSecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityIDSource is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingSecurityIDSource() (*field.UnderlyingSecurityIDSourceField, errors.MessageRejectError) {
	f := &field.UnderlyingSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityIDSource reads a UnderlyingSecurityIDSource from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingSecurityIDSource(f *field.UnderlyingSecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyingSecurityAltID is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) NoUnderlyingSecurityAltID() (*field.NoUnderlyingSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoUnderlyingSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyingSecurityAltID reads a NoUnderlyingSecurityAltID from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetNoUnderlyingSecurityAltID(f *field.NoUnderlyingSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingProduct is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingProduct() (*field.UnderlyingProductField, errors.MessageRejectError) {
	f := &field.UnderlyingProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingProduct reads a UnderlyingProduct from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingProduct(f *field.UnderlyingProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCFICode is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingCFICode() (*field.UnderlyingCFICodeField, errors.MessageRejectError) {
	f := &field.UnderlyingCFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCFICode reads a UnderlyingCFICode from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingCFICode(f *field.UnderlyingCFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityType is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingSecurityType() (*field.UnderlyingSecurityTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityType reads a UnderlyingSecurityType from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingSecurityType(f *field.UnderlyingSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecuritySubType is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingSecuritySubType() (*field.UnderlyingSecuritySubTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingSecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecuritySubType reads a UnderlyingSecuritySubType from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingSecuritySubType(f *field.UnderlyingSecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityMonthYear is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingMaturityMonthYear() (*field.UnderlyingMaturityMonthYearField, errors.MessageRejectError) {
	f := &field.UnderlyingMaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityMonthYear reads a UnderlyingMaturityMonthYear from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingMaturityMonthYear(f *field.UnderlyingMaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityDate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingMaturityDate() (*field.UnderlyingMaturityDateField, errors.MessageRejectError) {
	f := &field.UnderlyingMaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityDate reads a UnderlyingMaturityDate from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingMaturityDate(f *field.UnderlyingMaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCouponPaymentDate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingCouponPaymentDate() (*field.UnderlyingCouponPaymentDateField, errors.MessageRejectError) {
	f := &field.UnderlyingCouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCouponPaymentDate reads a UnderlyingCouponPaymentDate from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingCouponPaymentDate(f *field.UnderlyingCouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingIssueDate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingIssueDate() (*field.UnderlyingIssueDateField, errors.MessageRejectError) {
	f := &field.UnderlyingIssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingIssueDate reads a UnderlyingIssueDate from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingIssueDate(f *field.UnderlyingIssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepoCollateralSecurityType is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingRepoCollateralSecurityType() (*field.UnderlyingRepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingRepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepoCollateralSecurityType reads a UnderlyingRepoCollateralSecurityType from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingRepoCollateralSecurityType(f *field.UnderlyingRepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepurchaseTerm is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingRepurchaseTerm() (*field.UnderlyingRepurchaseTermField, errors.MessageRejectError) {
	f := &field.UnderlyingRepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepurchaseTerm reads a UnderlyingRepurchaseTerm from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingRepurchaseTerm(f *field.UnderlyingRepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepurchaseRate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingRepurchaseRate() (*field.UnderlyingRepurchaseRateField, errors.MessageRejectError) {
	f := &field.UnderlyingRepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepurchaseRate reads a UnderlyingRepurchaseRate from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingRepurchaseRate(f *field.UnderlyingRepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFactor is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingFactor() (*field.UnderlyingFactorField, errors.MessageRejectError) {
	f := &field.UnderlyingFactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFactor reads a UnderlyingFactor from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingFactor(f *field.UnderlyingFactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCreditRating is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingCreditRating() (*field.UnderlyingCreditRatingField, errors.MessageRejectError) {
	f := &field.UnderlyingCreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCreditRating reads a UnderlyingCreditRating from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingCreditRating(f *field.UnderlyingCreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingInstrRegistry is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingInstrRegistry() (*field.UnderlyingInstrRegistryField, errors.MessageRejectError) {
	f := &field.UnderlyingInstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingInstrRegistry reads a UnderlyingInstrRegistry from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingInstrRegistry(f *field.UnderlyingInstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCountryOfIssue is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingCountryOfIssue() (*field.UnderlyingCountryOfIssueField, errors.MessageRejectError) {
	f := &field.UnderlyingCountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCountryOfIssue reads a UnderlyingCountryOfIssue from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingCountryOfIssue(f *field.UnderlyingCountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStateOrProvinceOfIssue is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingStateOrProvinceOfIssue() (*field.UnderlyingStateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.UnderlyingStateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStateOrProvinceOfIssue reads a UnderlyingStateOrProvinceOfIssue from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingStateOrProvinceOfIssue(f *field.UnderlyingStateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingLocaleOfIssue is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingLocaleOfIssue() (*field.UnderlyingLocaleOfIssueField, errors.MessageRejectError) {
	f := &field.UnderlyingLocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingLocaleOfIssue reads a UnderlyingLocaleOfIssue from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingLocaleOfIssue(f *field.UnderlyingLocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRedemptionDate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingRedemptionDate() (*field.UnderlyingRedemptionDateField, errors.MessageRejectError) {
	f := &field.UnderlyingRedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRedemptionDate reads a UnderlyingRedemptionDate from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingRedemptionDate(f *field.UnderlyingRedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStrikePrice is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingStrikePrice() (*field.UnderlyingStrikePriceField, errors.MessageRejectError) {
	f := &field.UnderlyingStrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStrikePrice reads a UnderlyingStrikePrice from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingStrikePrice(f *field.UnderlyingStrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStrikeCurrency is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingStrikeCurrency() (*field.UnderlyingStrikeCurrencyField, errors.MessageRejectError) {
	f := &field.UnderlyingStrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStrikeCurrency reads a UnderlyingStrikeCurrency from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingStrikeCurrency(f *field.UnderlyingStrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingOptAttribute is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingOptAttribute() (*field.UnderlyingOptAttributeField, errors.MessageRejectError) {
	f := &field.UnderlyingOptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingOptAttribute reads a UnderlyingOptAttribute from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingOptAttribute(f *field.UnderlyingOptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingContractMultiplier is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingContractMultiplier() (*field.UnderlyingContractMultiplierField, errors.MessageRejectError) {
	f := &field.UnderlyingContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingContractMultiplier reads a UnderlyingContractMultiplier from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingContractMultiplier(f *field.UnderlyingContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCouponRate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingCouponRate() (*field.UnderlyingCouponRateField, errors.MessageRejectError) {
	f := &field.UnderlyingCouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCouponRate reads a UnderlyingCouponRate from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingCouponRate(f *field.UnderlyingCouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityExchange is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingSecurityExchange() (*field.UnderlyingSecurityExchangeField, errors.MessageRejectError) {
	f := &field.UnderlyingSecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityExchange reads a UnderlyingSecurityExchange from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingSecurityExchange(f *field.UnderlyingSecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingIssuer is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingIssuer() (*field.UnderlyingIssuerField, errors.MessageRejectError) {
	f := &field.UnderlyingIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingIssuer reads a UnderlyingIssuer from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingIssuer(f *field.UnderlyingIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingIssuerLen is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) EncodedUnderlyingIssuerLen() (*field.EncodedUnderlyingIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedUnderlyingIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingIssuerLen reads a EncodedUnderlyingIssuerLen from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetEncodedUnderlyingIssuerLen(f *field.EncodedUnderlyingIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingIssuer is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) EncodedUnderlyingIssuer() (*field.EncodedUnderlyingIssuerField, errors.MessageRejectError) {
	f := &field.EncodedUnderlyingIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingIssuer reads a EncodedUnderlyingIssuer from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetEncodedUnderlyingIssuer(f *field.EncodedUnderlyingIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityDesc is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingSecurityDesc() (*field.UnderlyingSecurityDescField, errors.MessageRejectError) {
	f := &field.UnderlyingSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityDesc reads a UnderlyingSecurityDesc from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingSecurityDesc(f *field.UnderlyingSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingSecurityDescLen is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) EncodedUnderlyingSecurityDescLen() (*field.EncodedUnderlyingSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedUnderlyingSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingSecurityDescLen reads a EncodedUnderlyingSecurityDescLen from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetEncodedUnderlyingSecurityDescLen(f *field.EncodedUnderlyingSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingSecurityDesc is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) EncodedUnderlyingSecurityDesc() (*field.EncodedUnderlyingSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedUnderlyingSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingSecurityDesc reads a EncodedUnderlyingSecurityDesc from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetEncodedUnderlyingSecurityDesc(f *field.EncodedUnderlyingSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCPProgram is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingCPProgram() (*field.UnderlyingCPProgramField, errors.MessageRejectError) {
	f := &field.UnderlyingCPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCPProgram reads a UnderlyingCPProgram from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingCPProgram(f *field.UnderlyingCPProgramField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCPRegType is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingCPRegType() (*field.UnderlyingCPRegTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingCPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCPRegType reads a UnderlyingCPRegType from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingCPRegType(f *field.UnderlyingCPRegTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCurrency is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingCurrency() (*field.UnderlyingCurrencyField, errors.MessageRejectError) {
	f := &field.UnderlyingCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCurrency reads a UnderlyingCurrency from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingCurrency(f *field.UnderlyingCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingQty is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingQty() (*field.UnderlyingQtyField, errors.MessageRejectError) {
	f := &field.UnderlyingQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingQty reads a UnderlyingQty from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingQty(f *field.UnderlyingQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPx is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingPx() (*field.UnderlyingPxField, errors.MessageRejectError) {
	f := &field.UnderlyingPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPx reads a UnderlyingPx from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingPx(f *field.UnderlyingPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingDirtyPrice is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingDirtyPrice() (*field.UnderlyingDirtyPriceField, errors.MessageRejectError) {
	f := &field.UnderlyingDirtyPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingDirtyPrice reads a UnderlyingDirtyPrice from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingDirtyPrice(f *field.UnderlyingDirtyPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingEndPrice is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingEndPrice() (*field.UnderlyingEndPriceField, errors.MessageRejectError) {
	f := &field.UnderlyingEndPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingEndPrice reads a UnderlyingEndPrice from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingEndPrice(f *field.UnderlyingEndPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStartValue is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingStartValue() (*field.UnderlyingStartValueField, errors.MessageRejectError) {
	f := &field.UnderlyingStartValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStartValue reads a UnderlyingStartValue from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingStartValue(f *field.UnderlyingStartValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCurrentValue is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingCurrentValue() (*field.UnderlyingCurrentValueField, errors.MessageRejectError) {
	f := &field.UnderlyingCurrentValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCurrentValue reads a UnderlyingCurrentValue from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingCurrentValue(f *field.UnderlyingCurrentValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingEndValue is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingEndValue() (*field.UnderlyingEndValueField, errors.MessageRejectError) {
	f := &field.UnderlyingEndValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingEndValue reads a UnderlyingEndValue from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingEndValue(f *field.UnderlyingEndValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyingStips is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) NoUnderlyingStips() (*field.NoUnderlyingStipsField, errors.MessageRejectError) {
	f := &field.NoUnderlyingStipsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyingStips reads a NoUnderlyingStips from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetNoUnderlyingStips(f *field.NoUnderlyingStipsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingAllocationPercent is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingAllocationPercent() (*field.UnderlyingAllocationPercentField, errors.MessageRejectError) {
	f := &field.UnderlyingAllocationPercentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingAllocationPercent reads a UnderlyingAllocationPercent from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingAllocationPercent(f *field.UnderlyingAllocationPercentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSettlementType is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingSettlementType() (*field.UnderlyingSettlementTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingSettlementTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSettlementType reads a UnderlyingSettlementType from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingSettlementType(f *field.UnderlyingSettlementTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCashAmount is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingCashAmount() (*field.UnderlyingCashAmountField, errors.MessageRejectError) {
	f := &field.UnderlyingCashAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCashAmount reads a UnderlyingCashAmount from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingCashAmount(f *field.UnderlyingCashAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCashType is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingCashType() (*field.UnderlyingCashTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingCashTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCashType reads a UnderlyingCashType from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingCashType(f *field.UnderlyingCashTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingUnitOfMeasure is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingUnitOfMeasure() (*field.UnderlyingUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.UnderlyingUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingUnitOfMeasure reads a UnderlyingUnitOfMeasure from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingUnitOfMeasure(f *field.UnderlyingUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingTimeUnit is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingTimeUnit() (*field.UnderlyingTimeUnitField, errors.MessageRejectError) {
	f := &field.UnderlyingTimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingTimeUnit reads a UnderlyingTimeUnit from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingTimeUnit(f *field.UnderlyingTimeUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCapValue is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingCapValue() (*field.UnderlyingCapValueField, errors.MessageRejectError) {
	f := &field.UnderlyingCapValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCapValue reads a UnderlyingCapValue from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingCapValue(f *field.UnderlyingCapValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUndlyInstrumentParties is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) NoUndlyInstrumentParties() (*field.NoUndlyInstrumentPartiesField, errors.MessageRejectError) {
	f := &field.NoUndlyInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUndlyInstrumentParties reads a NoUndlyInstrumentParties from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetNoUndlyInstrumentParties(f *field.NoUndlyInstrumentPartiesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSettlMethod is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingSettlMethod() (*field.UnderlyingSettlMethodField, errors.MessageRejectError) {
	f := &field.UnderlyingSettlMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSettlMethod reads a UnderlyingSettlMethod from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingSettlMethod(f *field.UnderlyingSettlMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingAdjustedQuantity is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingAdjustedQuantity() (*field.UnderlyingAdjustedQuantityField, errors.MessageRejectError) {
	f := &field.UnderlyingAdjustedQuantityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingAdjustedQuantity reads a UnderlyingAdjustedQuantity from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingAdjustedQuantity(f *field.UnderlyingAdjustedQuantityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFXRate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingFXRate() (*field.UnderlyingFXRateField, errors.MessageRejectError) {
	f := &field.UnderlyingFXRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFXRate reads a UnderlyingFXRate from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingFXRate(f *field.UnderlyingFXRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFXRateCalc is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingFXRateCalc() (*field.UnderlyingFXRateCalcField, errors.MessageRejectError) {
	f := &field.UnderlyingFXRateCalcField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFXRateCalc reads a UnderlyingFXRateCalc from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingFXRateCalc(f *field.UnderlyingFXRateCalcField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityTime is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingMaturityTime() (*field.UnderlyingMaturityTimeField, errors.MessageRejectError) {
	f := &field.UnderlyingMaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityTime reads a UnderlyingMaturityTime from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingMaturityTime(f *field.UnderlyingMaturityTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPutOrCall is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingPutOrCall() (*field.UnderlyingPutOrCallField, errors.MessageRejectError) {
	f := &field.UnderlyingPutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPutOrCall reads a UnderlyingPutOrCall from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingPutOrCall(f *field.UnderlyingPutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingExerciseStyle is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingExerciseStyle() (*field.UnderlyingExerciseStyleField, errors.MessageRejectError) {
	f := &field.UnderlyingExerciseStyleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingExerciseStyle reads a UnderlyingExerciseStyle from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingExerciseStyle(f *field.UnderlyingExerciseStyleField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingUnitOfMeasureQty is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingUnitOfMeasureQty() (*field.UnderlyingUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.UnderlyingUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingUnitOfMeasureQty reads a UnderlyingUnitOfMeasureQty from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingUnitOfMeasureQty(f *field.UnderlyingUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceUnitOfMeasure is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingPriceUnitOfMeasure() (*field.UnderlyingPriceUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.UnderlyingPriceUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceUnitOfMeasure reads a UnderlyingPriceUnitOfMeasure from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingPriceUnitOfMeasure(f *field.UnderlyingPriceUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceUnitOfMeasureQty is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingPriceUnitOfMeasureQty() (*field.UnderlyingPriceUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.UnderlyingPriceUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceUnitOfMeasureQty reads a UnderlyingPriceUnitOfMeasureQty from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingPriceUnitOfMeasureQty(f *field.UnderlyingPriceUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingContractMultiplierUnit is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingContractMultiplierUnit() (*field.UnderlyingContractMultiplierUnitField, errors.MessageRejectError) {
	f := &field.UnderlyingContractMultiplierUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingContractMultiplierUnit reads a UnderlyingContractMultiplierUnit from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingContractMultiplierUnit(f *field.UnderlyingContractMultiplierUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFlowScheduleType is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingFlowScheduleType() (*field.UnderlyingFlowScheduleTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingFlowScheduleTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFlowScheduleType reads a UnderlyingFlowScheduleType from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingFlowScheduleType(f *field.UnderlyingFlowScheduleTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRestructuringType is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingRestructuringType() (*field.UnderlyingRestructuringTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingRestructuringTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRestructuringType reads a UnderlyingRestructuringType from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingRestructuringType(f *field.UnderlyingRestructuringTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSeniority is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingSeniority() (*field.UnderlyingSeniorityField, errors.MessageRejectError) {
	f := &field.UnderlyingSeniorityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSeniority reads a UnderlyingSeniority from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingSeniority(f *field.UnderlyingSeniorityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingNotionalPercentageOutstanding is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingNotionalPercentageOutstanding() (*field.UnderlyingNotionalPercentageOutstandingField, errors.MessageRejectError) {
	f := &field.UnderlyingNotionalPercentageOutstandingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingNotionalPercentageOutstanding reads a UnderlyingNotionalPercentageOutstanding from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingNotionalPercentageOutstanding(f *field.UnderlyingNotionalPercentageOutstandingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingOriginalNotionalPercentageOutstanding is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingOriginalNotionalPercentageOutstanding() (*field.UnderlyingOriginalNotionalPercentageOutstandingField, errors.MessageRejectError) {
	f := &field.UnderlyingOriginalNotionalPercentageOutstandingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingOriginalNotionalPercentageOutstanding reads a UnderlyingOriginalNotionalPercentageOutstanding from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingOriginalNotionalPercentageOutstanding(f *field.UnderlyingOriginalNotionalPercentageOutstandingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingAttachmentPoint is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingAttachmentPoint() (*field.UnderlyingAttachmentPointField, errors.MessageRejectError) {
	f := &field.UnderlyingAttachmentPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingAttachmentPoint reads a UnderlyingAttachmentPoint from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingAttachmentPoint(f *field.UnderlyingAttachmentPointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingDetachmentPoint is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingDetachmentPoint() (*field.UnderlyingDetachmentPointField, errors.MessageRejectError) {
	f := &field.UnderlyingDetachmentPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingDetachmentPoint reads a UnderlyingDetachmentPoint from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetUnderlyingDetachmentPoint(f *field.UnderlyingDetachmentPointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTargetPartyIDs is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) NoTargetPartyIDs() (*field.NoTargetPartyIDsField, errors.MessageRejectError) {
	f := &field.NoTargetPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTargetPartyIDs reads a NoTargetPartyIDs from OrderMassStatusRequest.
func (m OrderMassStatusRequest) GetNoTargetPartyIDs(f *field.NoTargetPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}
