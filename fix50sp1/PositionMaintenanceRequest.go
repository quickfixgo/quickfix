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

//PositionMaintenanceRequest msg type = AL.
type PositionMaintenanceRequest struct {
	message.Message
}

//PositionMaintenanceRequestBuilder builds PositionMaintenanceRequest messages.
type PositionMaintenanceRequestBuilder struct {
	message.MessageBuilder
}

//CreatePositionMaintenanceRequestBuilder returns an initialized PositionMaintenanceRequestBuilder with specified required fields.
func CreatePositionMaintenanceRequestBuilder(
	postranstype *field.PosTransTypeField,
	posmaintaction *field.PosMaintActionField,
	clearingbusinessdate *field.ClearingBusinessDateField) PositionMaintenanceRequestBuilder {
	var builder PositionMaintenanceRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header.Set(field.NewMsgType("AL"))
	builder.Body.Set(postranstype)
	builder.Body.Set(posmaintaction)
	builder.Body.Set(clearingbusinessdate)
	return builder
}

//PosReqID is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) PosReqID() (*field.PosReqIDField, errors.MessageRejectError) {
	f := &field.PosReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPosReqID reads a PosReqID from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetPosReqID(f *field.PosReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosTransType is a required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) PosTransType() (*field.PosTransTypeField, errors.MessageRejectError) {
	f := &field.PosTransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPosTransType reads a PosTransType from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetPosTransType(f *field.PosTransTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosMaintAction is a required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) PosMaintAction() (*field.PosMaintActionField, errors.MessageRejectError) {
	f := &field.PosMaintActionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPosMaintAction reads a PosMaintAction from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetPosMaintAction(f *field.PosMaintActionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigPosReqRefID is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) OrigPosReqRefID() (*field.OrigPosReqRefIDField, errors.MessageRejectError) {
	f := &field.OrigPosReqRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigPosReqRefID reads a OrigPosReqRefID from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetOrigPosReqRefID(f *field.OrigPosReqRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosMaintRptRefID is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) PosMaintRptRefID() (*field.PosMaintRptRefIDField, errors.MessageRejectError) {
	f := &field.PosMaintRptRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPosMaintRptRefID reads a PosMaintRptRefID from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetPosMaintRptRefID(f *field.PosMaintRptRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) ClearingBusinessDate() (*field.ClearingBusinessDateField, errors.MessageRejectError) {
	f := &field.ClearingBusinessDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetClearingBusinessDate(f *field.ClearingBusinessDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessID is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) SettlSessID() (*field.SettlSessIDField, errors.MessageRejectError) {
	f := &field.SettlSessIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessID reads a SettlSessID from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetSettlSessID(f *field.SettlSessIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessSubID is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) SettlSessSubID() (*field.SettlSessSubIDField, errors.MessageRejectError) {
	f := &field.SettlSessSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessSubID reads a SettlSessSubID from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetSettlSessSubID(f *field.SettlSessSubIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) Account() (*field.AccountField, errors.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetAccount(f *field.AccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) AcctIDSource() (*field.AcctIDSourceField, errors.MessageRejectError) {
	f := &field.AcctIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetAcctIDSource(f *field.AcctIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) AccountType() (*field.AccountTypeField, errors.MessageRejectError) {
	f := &field.AccountTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetAccountType(f *field.AccountTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) SecurityIDSource() (*field.SecurityIDSourceField, errors.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetSecurityIDSource(f *field.SecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) NoSecurityAltID() (*field.NoSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetNoSecurityAltID(f *field.NoSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) SecuritySubType() (*field.SecuritySubTypeField, errors.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetSecuritySubType(f *field.SecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) MaturityDate() (*field.MaturityDateField, errors.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetMaturityDate(f *field.MaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) CouponPaymentDate() (*field.CouponPaymentDateField, errors.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetCouponPaymentDate(f *field.CouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) IssueDate() (*field.IssueDateField, errors.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetIssueDate(f *field.IssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) RepurchaseTerm() (*field.RepurchaseTermField, errors.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetRepurchaseTerm(f *field.RepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) RepurchaseRate() (*field.RepurchaseRateField, errors.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetRepurchaseRate(f *field.RepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) Factor() (*field.FactorField, errors.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetFactor(f *field.FactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) CreditRating() (*field.CreditRatingField, errors.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetCreditRating(f *field.CreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) InstrRegistry() (*field.InstrRegistryField, errors.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetInstrRegistry(f *field.InstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) CountryOfIssue() (*field.CountryOfIssueField, errors.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetCountryOfIssue(f *field.CountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) LocaleOfIssue() (*field.LocaleOfIssueField, errors.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetLocaleOfIssue(f *field.LocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) RedemptionDate() (*field.RedemptionDateField, errors.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetRedemptionDate(f *field.RedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) StrikeCurrency() (*field.StrikeCurrencyField, errors.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetStrikeCurrency(f *field.StrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) Pool() (*field.PoolField, errors.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetPool(f *field.PoolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) ContractSettlMonth() (*field.ContractSettlMonthField, errors.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetContractSettlMonth(f *field.ContractSettlMonthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) CPProgram() (*field.CPProgramField, errors.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetCPProgram(f *field.CPProgramField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) CPRegType() (*field.CPRegTypeField, errors.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetCPRegType(f *field.CPRegTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) NoEvents() (*field.NoEventsField, errors.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetNoEvents(f *field.NoEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) DatedDate() (*field.DatedDateField, errors.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetDatedDate(f *field.DatedDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) InterestAccrualDate() (*field.InterestAccrualDateField, errors.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetInterestAccrualDate(f *field.InterestAccrualDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) SecurityStatus() (*field.SecurityStatusField, errors.MessageRejectError) {
	f := &field.SecurityStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetSecurityStatus(f *field.SecurityStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) SettleOnOpenFlag() (*field.SettleOnOpenFlagField, errors.MessageRejectError) {
	f := &field.SettleOnOpenFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetSettleOnOpenFlag(f *field.SettleOnOpenFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethodField, errors.MessageRejectError) {
	f := &field.InstrmtAssignmentMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) StrikeMultiplier() (*field.StrikeMultiplierField, errors.MessageRejectError) {
	f := &field.StrikeMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetStrikeMultiplier(f *field.StrikeMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) StrikeValue() (*field.StrikeValueField, errors.MessageRejectError) {
	f := &field.StrikeValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetStrikeValue(f *field.StrikeValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) MinPriceIncrement() (*field.MinPriceIncrementField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetMinPriceIncrement(f *field.MinPriceIncrementField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) PositionLimit() (*field.PositionLimitField, errors.MessageRejectError) {
	f := &field.PositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetPositionLimit(f *field.PositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) NTPositionLimit() (*field.NTPositionLimitField, errors.MessageRejectError) {
	f := &field.NTPositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetNTPositionLimit(f *field.NTPositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) NoInstrumentParties() (*field.NoInstrumentPartiesField, errors.MessageRejectError) {
	f := &field.NoInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetNoInstrumentParties(f *field.NoInstrumentPartiesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) UnitOfMeasure() (*field.UnitOfMeasureField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetUnitOfMeasure(f *field.UnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) TimeUnit() (*field.TimeUnitField, errors.MessageRejectError) {
	f := &field.TimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetTimeUnit(f *field.TimeUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) MaturityTime() (*field.MaturityTimeField, errors.MessageRejectError) {
	f := &field.MaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetMaturityTime(f *field.MaturityTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) SecurityGroup() (*field.SecurityGroupField, errors.MessageRejectError) {
	f := &field.SecurityGroupField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetSecurityGroup(f *field.SecurityGroupField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) MinPriceIncrementAmount() (*field.MinPriceIncrementAmountField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) UnitOfMeasureQty() (*field.UnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetUnitOfMeasureQty(f *field.UnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) SecurityXMLLen() (*field.SecurityXMLLenField, errors.MessageRejectError) {
	f := &field.SecurityXMLLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetSecurityXMLLen(f *field.SecurityXMLLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) SecurityXML() (*field.SecurityXMLField, errors.MessageRejectError) {
	f := &field.SecurityXMLField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetSecurityXML(f *field.SecurityXMLField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) SecurityXMLSchema() (*field.SecurityXMLSchemaField, errors.MessageRejectError) {
	f := &field.SecurityXMLSchemaField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetSecurityXMLSchema(f *field.SecurityXMLSchemaField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) ProductComplex() (*field.ProductComplexField, errors.MessageRejectError) {
	f := &field.ProductComplexField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetProductComplex(f *field.ProductComplexField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) PriceUnitOfMeasure() (*field.PriceUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.PriceUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.PriceUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) SettlMethod() (*field.SettlMethodField, errors.MessageRejectError) {
	f := &field.SettlMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetSettlMethod(f *field.SettlMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) ExerciseStyle() (*field.ExerciseStyleField, errors.MessageRejectError) {
	f := &field.ExerciseStyleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetExerciseStyle(f *field.ExerciseStyleField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayAmount is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) OptPayAmount() (*field.OptPayAmountField, errors.MessageRejectError) {
	f := &field.OptPayAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayAmount reads a OptPayAmount from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetOptPayAmount(f *field.OptPayAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) PriceQuoteMethod() (*field.PriceQuoteMethodField, errors.MessageRejectError) {
	f := &field.PriceQuoteMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetPriceQuoteMethod(f *field.PriceQuoteMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) ListMethod() (*field.ListMethodField, errors.MessageRejectError) {
	f := &field.ListMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetListMethod(f *field.ListMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) CapPrice() (*field.CapPriceField, errors.MessageRejectError) {
	f := &field.CapPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetCapPrice(f *field.CapPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) FloorPrice() (*field.FloorPriceField, errors.MessageRejectError) {
	f := &field.FloorPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetFloorPrice(f *field.FloorPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) PutOrCall() (*field.PutOrCallField, errors.MessageRejectError) {
	f := &field.PutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetPutOrCall(f *field.PutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) FlexibleIndicator() (*field.FlexibleIndicatorField, errors.MessageRejectError) {
	f := &field.FlexibleIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetFlexibleIndicator(f *field.FlexibleIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicatorField, errors.MessageRejectError) {
	f := &field.FlexProductEligibilityIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FuturesValuationMethod is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) FuturesValuationMethod() (*field.FuturesValuationMethodField, errors.MessageRejectError) {
	f := &field.FuturesValuationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFuturesValuationMethod reads a FuturesValuationMethod from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetFuturesValuationMethod(f *field.FuturesValuationMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) NoLegs() (*field.NoLegsField, errors.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetNoLegs(f *field.NoLegsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) NoUnderlyings() (*field.NoUnderlyingsField, errors.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetNoUnderlyings(f *field.NoUnderlyingsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) NoTradingSessions() (*field.NoTradingSessionsField, errors.MessageRejectError) {
	f := &field.NoTradingSessionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetNoTradingSessions(f *field.NoTradingSessionsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPositions is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) NoPositions() (*field.NoPositionsField, errors.MessageRejectError) {
	f := &field.NoPositionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPositions reads a NoPositions from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetNoPositions(f *field.NoPositionsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AdjustmentType is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) AdjustmentType() (*field.AdjustmentTypeField, errors.MessageRejectError) {
	f := &field.AdjustmentTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAdjustmentType reads a AdjustmentType from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetAdjustmentType(f *field.AdjustmentTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContraryInstructionIndicator is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) ContraryInstructionIndicator() (*field.ContraryInstructionIndicatorField, errors.MessageRejectError) {
	f := &field.ContraryInstructionIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContraryInstructionIndicator reads a ContraryInstructionIndicator from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetContraryInstructionIndicator(f *field.ContraryInstructionIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriorSpreadIndicator is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) PriorSpreadIndicator() (*field.PriorSpreadIndicatorField, errors.MessageRejectError) {
	f := &field.PriorSpreadIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriorSpreadIndicator reads a PriorSpreadIndicator from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetPriorSpreadIndicator(f *field.PriorSpreadIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ThresholdAmount is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) ThresholdAmount() (*field.ThresholdAmountField, errors.MessageRejectError) {
	f := &field.ThresholdAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetThresholdAmount reads a ThresholdAmount from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetThresholdAmount(f *field.ThresholdAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPosAmt is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) NoPosAmt() (*field.NoPosAmtField, errors.MessageRejectError) {
	f := &field.NoPosAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPosAmt reads a NoPosAmt from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetNoPosAmt(f *field.NoPosAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) SettlCurrency() (*field.SettlCurrencyField, errors.MessageRejectError) {
	f := &field.SettlCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) GetSettlCurrency(f *field.SettlCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}
