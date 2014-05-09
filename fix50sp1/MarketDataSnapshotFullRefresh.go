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

//MarketDataSnapshotFullRefresh msg type = W.
type MarketDataSnapshotFullRefresh struct {
	message.Message
}

//MarketDataSnapshotFullRefreshBuilder builds MarketDataSnapshotFullRefresh messages.
type MarketDataSnapshotFullRefreshBuilder struct {
	message.MessageBuilder
}

//CreateMarketDataSnapshotFullRefreshBuilder returns an initialized MarketDataSnapshotFullRefreshBuilder with specified required fields.
func CreateMarketDataSnapshotFullRefreshBuilder(
	nomdentries *field.NoMDEntriesField) MarketDataSnapshotFullRefreshBuilder {
	var builder MarketDataSnapshotFullRefreshBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header.Set(field.NewMsgType("W"))
	builder.Body.Set(nomdentries)
	return builder
}

//MDReqID is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MDReqID() (*field.MDReqIDField, errors.MessageRejectError) {
	f := &field.MDReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDReqID reads a MDReqID from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetMDReqID(f *field.MDReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecurityIDSource() (*field.SecurityIDSourceField, errors.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSecurityIDSource(f *field.SecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NoSecurityAltID() (*field.NoSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetNoSecurityAltID(f *field.NoSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecuritySubType() (*field.SecuritySubTypeField, errors.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSecuritySubType(f *field.SecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MaturityDate() (*field.MaturityDateField, errors.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetMaturityDate(f *field.MaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CouponPaymentDate() (*field.CouponPaymentDateField, errors.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetCouponPaymentDate(f *field.CouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) IssueDate() (*field.IssueDateField, errors.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetIssueDate(f *field.IssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) RepurchaseTerm() (*field.RepurchaseTermField, errors.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetRepurchaseTerm(f *field.RepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) RepurchaseRate() (*field.RepurchaseRateField, errors.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetRepurchaseRate(f *field.RepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) Factor() (*field.FactorField, errors.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetFactor(f *field.FactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CreditRating() (*field.CreditRatingField, errors.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetCreditRating(f *field.CreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) InstrRegistry() (*field.InstrRegistryField, errors.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetInstrRegistry(f *field.InstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CountryOfIssue() (*field.CountryOfIssueField, errors.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetCountryOfIssue(f *field.CountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) LocaleOfIssue() (*field.LocaleOfIssueField, errors.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetLocaleOfIssue(f *field.LocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) RedemptionDate() (*field.RedemptionDateField, errors.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetRedemptionDate(f *field.RedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) StrikeCurrency() (*field.StrikeCurrencyField, errors.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetStrikeCurrency(f *field.StrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) Pool() (*field.PoolField, errors.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetPool(f *field.PoolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) ContractSettlMonth() (*field.ContractSettlMonthField, errors.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetContractSettlMonth(f *field.ContractSettlMonthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CPProgram() (*field.CPProgramField, errors.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetCPProgram(f *field.CPProgramField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CPRegType() (*field.CPRegTypeField, errors.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetCPRegType(f *field.CPRegTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NoEvents() (*field.NoEventsField, errors.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetNoEvents(f *field.NoEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) DatedDate() (*field.DatedDateField, errors.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetDatedDate(f *field.DatedDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) InterestAccrualDate() (*field.InterestAccrualDateField, errors.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetInterestAccrualDate(f *field.InterestAccrualDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecurityStatus() (*field.SecurityStatusField, errors.MessageRejectError) {
	f := &field.SecurityStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSecurityStatus(f *field.SecurityStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SettleOnOpenFlag() (*field.SettleOnOpenFlagField, errors.MessageRejectError) {
	f := &field.SettleOnOpenFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSettleOnOpenFlag(f *field.SettleOnOpenFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethodField, errors.MessageRejectError) {
	f := &field.InstrmtAssignmentMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) StrikeMultiplier() (*field.StrikeMultiplierField, errors.MessageRejectError) {
	f := &field.StrikeMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetStrikeMultiplier(f *field.StrikeMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) StrikeValue() (*field.StrikeValueField, errors.MessageRejectError) {
	f := &field.StrikeValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetStrikeValue(f *field.StrikeValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MinPriceIncrement() (*field.MinPriceIncrementField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetMinPriceIncrement(f *field.MinPriceIncrementField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) PositionLimit() (*field.PositionLimitField, errors.MessageRejectError) {
	f := &field.PositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetPositionLimit(f *field.PositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NTPositionLimit() (*field.NTPositionLimitField, errors.MessageRejectError) {
	f := &field.NTPositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetNTPositionLimit(f *field.NTPositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NoInstrumentParties() (*field.NoInstrumentPartiesField, errors.MessageRejectError) {
	f := &field.NoInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetNoInstrumentParties(f *field.NoInstrumentPartiesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) UnitOfMeasure() (*field.UnitOfMeasureField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetUnitOfMeasure(f *field.UnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) TimeUnit() (*field.TimeUnitField, errors.MessageRejectError) {
	f := &field.TimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetTimeUnit(f *field.TimeUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MaturityTime() (*field.MaturityTimeField, errors.MessageRejectError) {
	f := &field.MaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetMaturityTime(f *field.MaturityTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecurityGroup() (*field.SecurityGroupField, errors.MessageRejectError) {
	f := &field.SecurityGroupField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSecurityGroup(f *field.SecurityGroupField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MinPriceIncrementAmount() (*field.MinPriceIncrementAmountField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) UnitOfMeasureQty() (*field.UnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetUnitOfMeasureQty(f *field.UnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecurityXMLLen() (*field.SecurityXMLLenField, errors.MessageRejectError) {
	f := &field.SecurityXMLLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSecurityXMLLen(f *field.SecurityXMLLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecurityXML() (*field.SecurityXMLField, errors.MessageRejectError) {
	f := &field.SecurityXMLField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSecurityXML(f *field.SecurityXMLField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecurityXMLSchema() (*field.SecurityXMLSchemaField, errors.MessageRejectError) {
	f := &field.SecurityXMLSchemaField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSecurityXMLSchema(f *field.SecurityXMLSchemaField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) ProductComplex() (*field.ProductComplexField, errors.MessageRejectError) {
	f := &field.ProductComplexField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetProductComplex(f *field.ProductComplexField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) PriceUnitOfMeasure() (*field.PriceUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.PriceUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.PriceUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SettlMethod() (*field.SettlMethodField, errors.MessageRejectError) {
	f := &field.SettlMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSettlMethod(f *field.SettlMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) ExerciseStyle() (*field.ExerciseStyleField, errors.MessageRejectError) {
	f := &field.ExerciseStyleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetExerciseStyle(f *field.ExerciseStyleField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayAmount is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) OptPayAmount() (*field.OptPayAmountField, errors.MessageRejectError) {
	f := &field.OptPayAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayAmount reads a OptPayAmount from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetOptPayAmount(f *field.OptPayAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) PriceQuoteMethod() (*field.PriceQuoteMethodField, errors.MessageRejectError) {
	f := &field.PriceQuoteMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetPriceQuoteMethod(f *field.PriceQuoteMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) ListMethod() (*field.ListMethodField, errors.MessageRejectError) {
	f := &field.ListMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetListMethod(f *field.ListMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CapPrice() (*field.CapPriceField, errors.MessageRejectError) {
	f := &field.CapPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetCapPrice(f *field.CapPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) FloorPrice() (*field.FloorPriceField, errors.MessageRejectError) {
	f := &field.FloorPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetFloorPrice(f *field.FloorPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) PutOrCall() (*field.PutOrCallField, errors.MessageRejectError) {
	f := &field.PutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetPutOrCall(f *field.PutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) FlexibleIndicator() (*field.FlexibleIndicatorField, errors.MessageRejectError) {
	f := &field.FlexibleIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetFlexibleIndicator(f *field.FlexibleIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicatorField, errors.MessageRejectError) {
	f := &field.FlexProductEligibilityIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FuturesValuationMethod is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) FuturesValuationMethod() (*field.FuturesValuationMethodField, errors.MessageRejectError) {
	f := &field.FuturesValuationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFuturesValuationMethod reads a FuturesValuationMethod from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetFuturesValuationMethod(f *field.FuturesValuationMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NoUnderlyings() (*field.NoUnderlyingsField, errors.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetNoUnderlyings(f *field.NoUnderlyingsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NoLegs() (*field.NoLegsField, errors.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetNoLegs(f *field.NoLegsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FinancialStatus is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) FinancialStatus() (*field.FinancialStatusField, errors.MessageRejectError) {
	f := &field.FinancialStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFinancialStatus reads a FinancialStatus from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetFinancialStatus(f *field.FinancialStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CorporateAction is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CorporateAction() (*field.CorporateActionField, errors.MessageRejectError) {
	f := &field.CorporateActionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCorporateAction reads a CorporateAction from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetCorporateAction(f *field.CorporateActionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NetChgPrevDay is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NetChgPrevDay() (*field.NetChgPrevDayField, errors.MessageRejectError) {
	f := &field.NetChgPrevDayField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNetChgPrevDay reads a NetChgPrevDay from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetNetChgPrevDay(f *field.NetChgPrevDayField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoMDEntries is a required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NoMDEntries() (*field.NoMDEntriesField, errors.MessageRejectError) {
	f := &field.NoMDEntriesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoMDEntries reads a NoMDEntries from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetNoMDEntries(f *field.NoMDEntriesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplQueueDepth is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) ApplQueueDepth() (*field.ApplQueueDepthField, errors.MessageRejectError) {
	f := &field.ApplQueueDepthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplQueueDepth reads a ApplQueueDepth from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetApplQueueDepth(f *field.ApplQueueDepthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplQueueResolution is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) ApplQueueResolution() (*field.ApplQueueResolutionField, errors.MessageRejectError) {
	f := &field.ApplQueueResolutionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplQueueResolution reads a ApplQueueResolution from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetApplQueueResolution(f *field.ApplQueueResolutionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MDReportID is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MDReportID() (*field.MDReportIDField, errors.MessageRejectError) {
	f := &field.MDReportIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDReportID reads a MDReportID from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetMDReportID(f *field.MDReportIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) ClearingBusinessDate() (*field.ClearingBusinessDateField, errors.MessageRejectError) {
	f := &field.ClearingBusinessDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetClearingBusinessDate(f *field.ClearingBusinessDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MDBookType is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MDBookType() (*field.MDBookTypeField, errors.MessageRejectError) {
	f := &field.MDBookTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDBookType reads a MDBookType from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetMDBookType(f *field.MDBookTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MDFeedType is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MDFeedType() (*field.MDFeedTypeField, errors.MessageRejectError) {
	f := &field.MDFeedTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDFeedType reads a MDFeedType from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetMDFeedType(f *field.MDFeedTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) TradeDate() (*field.TradeDateField, errors.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetTradeDate(f *field.TradeDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRoutingIDs is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NoRoutingIDs() (*field.NoRoutingIDsField, errors.MessageRejectError) {
	f := &field.NoRoutingIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRoutingIDs reads a NoRoutingIDs from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetNoRoutingIDs(f *field.NoRoutingIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MDSubBookType is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MDSubBookType() (*field.MDSubBookTypeField, errors.MessageRejectError) {
	f := &field.MDSubBookTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDSubBookType reads a MDSubBookType from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetMDSubBookType(f *field.MDSubBookTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketDepth is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MarketDepth() (*field.MarketDepthField, errors.MessageRejectError) {
	f := &field.MarketDepthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketDepth reads a MarketDepth from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetMarketDepth(f *field.MarketDepthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNumReports is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) TotNumReports() (*field.TotNumReportsField, errors.MessageRejectError) {
	f := &field.TotNumReportsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNumReports reads a TotNumReports from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetTotNumReports(f *field.TotNumReportsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefreshIndicator is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) RefreshIndicator() (*field.RefreshIndicatorField, errors.MessageRejectError) {
	f := &field.RefreshIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRefreshIndicator reads a RefreshIndicator from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetRefreshIndicator(f *field.RefreshIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) ApplID() (*field.ApplIDField, errors.MessageRejectError) {
	f := &field.ApplIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetApplID(f *field.ApplIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) ApplSeqNum() (*field.ApplSeqNumField, errors.MessageRejectError) {
	f := &field.ApplSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetApplSeqNum(f *field.ApplSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) ApplLastSeqNum() (*field.ApplLastSeqNumField, errors.MessageRejectError) {
	f := &field.ApplLastSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetApplLastSeqNum(f *field.ApplLastSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) ApplResendFlag() (*field.ApplResendFlagField, errors.MessageRejectError) {
	f := &field.ApplResendFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetApplResendFlag(f *field.ApplResendFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}
