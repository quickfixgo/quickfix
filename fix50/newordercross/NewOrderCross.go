//Package newordercross msg type = s.
package newordercross

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a NewOrderCross wrapper for the generic Message type
type Message struct {
	message.Message
}

//CrossID is a required field for NewOrderCross.
func (m Message) CrossID() (*field.CrossIDField, errors.MessageRejectError) {
	f := &field.CrossIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCrossID reads a CrossID from NewOrderCross.
func (m Message) GetCrossID(f *field.CrossIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CrossType is a required field for NewOrderCross.
func (m Message) CrossType() (*field.CrossTypeField, errors.MessageRejectError) {
	f := &field.CrossTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCrossType reads a CrossType from NewOrderCross.
func (m Message) GetCrossType(f *field.CrossTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CrossPrioritization is a required field for NewOrderCross.
func (m Message) CrossPrioritization() (*field.CrossPrioritizationField, errors.MessageRejectError) {
	f := &field.CrossPrioritizationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCrossPrioritization reads a CrossPrioritization from NewOrderCross.
func (m Message) GetCrossPrioritization(f *field.CrossPrioritizationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSides is a required field for NewOrderCross.
func (m Message) NoSides() (*field.NoSidesField, errors.MessageRejectError) {
	f := &field.NoSidesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSides reads a NoSides from NewOrderCross.
func (m Message) GetNoSides(f *field.NoSidesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for NewOrderCross.
func (m Message) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from NewOrderCross.
func (m Message) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for NewOrderCross.
func (m Message) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from NewOrderCross.
func (m Message) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for NewOrderCross.
func (m Message) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from NewOrderCross.
func (m Message) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for NewOrderCross.
func (m Message) SecurityIDSource() (*field.SecurityIDSourceField, errors.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from NewOrderCross.
func (m Message) GetSecurityIDSource(f *field.SecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for NewOrderCross.
func (m Message) NoSecurityAltID() (*field.NoSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from NewOrderCross.
func (m Message) GetNoSecurityAltID(f *field.NoSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for NewOrderCross.
func (m Message) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from NewOrderCross.
func (m Message) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for NewOrderCross.
func (m Message) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from NewOrderCross.
func (m Message) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for NewOrderCross.
func (m Message) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from NewOrderCross.
func (m Message) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for NewOrderCross.
func (m Message) SecuritySubType() (*field.SecuritySubTypeField, errors.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from NewOrderCross.
func (m Message) GetSecuritySubType(f *field.SecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for NewOrderCross.
func (m Message) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from NewOrderCross.
func (m Message) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for NewOrderCross.
func (m Message) MaturityDate() (*field.MaturityDateField, errors.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from NewOrderCross.
func (m Message) GetMaturityDate(f *field.MaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for NewOrderCross.
func (m Message) CouponPaymentDate() (*field.CouponPaymentDateField, errors.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from NewOrderCross.
func (m Message) GetCouponPaymentDate(f *field.CouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for NewOrderCross.
func (m Message) IssueDate() (*field.IssueDateField, errors.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from NewOrderCross.
func (m Message) GetIssueDate(f *field.IssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for NewOrderCross.
func (m Message) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from NewOrderCross.
func (m Message) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for NewOrderCross.
func (m Message) RepurchaseTerm() (*field.RepurchaseTermField, errors.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from NewOrderCross.
func (m Message) GetRepurchaseTerm(f *field.RepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for NewOrderCross.
func (m Message) RepurchaseRate() (*field.RepurchaseRateField, errors.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from NewOrderCross.
func (m Message) GetRepurchaseRate(f *field.RepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for NewOrderCross.
func (m Message) Factor() (*field.FactorField, errors.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from NewOrderCross.
func (m Message) GetFactor(f *field.FactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for NewOrderCross.
func (m Message) CreditRating() (*field.CreditRatingField, errors.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from NewOrderCross.
func (m Message) GetCreditRating(f *field.CreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for NewOrderCross.
func (m Message) InstrRegistry() (*field.InstrRegistryField, errors.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from NewOrderCross.
func (m Message) GetInstrRegistry(f *field.InstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for NewOrderCross.
func (m Message) CountryOfIssue() (*field.CountryOfIssueField, errors.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from NewOrderCross.
func (m Message) GetCountryOfIssue(f *field.CountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for NewOrderCross.
func (m Message) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from NewOrderCross.
func (m Message) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for NewOrderCross.
func (m Message) LocaleOfIssue() (*field.LocaleOfIssueField, errors.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from NewOrderCross.
func (m Message) GetLocaleOfIssue(f *field.LocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for NewOrderCross.
func (m Message) RedemptionDate() (*field.RedemptionDateField, errors.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from NewOrderCross.
func (m Message) GetRedemptionDate(f *field.RedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for NewOrderCross.
func (m Message) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from NewOrderCross.
func (m Message) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for NewOrderCross.
func (m Message) StrikeCurrency() (*field.StrikeCurrencyField, errors.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from NewOrderCross.
func (m Message) GetStrikeCurrency(f *field.StrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for NewOrderCross.
func (m Message) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from NewOrderCross.
func (m Message) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for NewOrderCross.
func (m Message) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from NewOrderCross.
func (m Message) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for NewOrderCross.
func (m Message) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from NewOrderCross.
func (m Message) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for NewOrderCross.
func (m Message) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from NewOrderCross.
func (m Message) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for NewOrderCross.
func (m Message) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from NewOrderCross.
func (m Message) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for NewOrderCross.
func (m Message) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from NewOrderCross.
func (m Message) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for NewOrderCross.
func (m Message) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from NewOrderCross.
func (m Message) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for NewOrderCross.
func (m Message) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from NewOrderCross.
func (m Message) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for NewOrderCross.
func (m Message) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from NewOrderCross.
func (m Message) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for NewOrderCross.
func (m Message) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from NewOrderCross.
func (m Message) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for NewOrderCross.
func (m Message) Pool() (*field.PoolField, errors.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from NewOrderCross.
func (m Message) GetPool(f *field.PoolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for NewOrderCross.
func (m Message) ContractSettlMonth() (*field.ContractSettlMonthField, errors.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from NewOrderCross.
func (m Message) GetContractSettlMonth(f *field.ContractSettlMonthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for NewOrderCross.
func (m Message) CPProgram() (*field.CPProgramField, errors.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from NewOrderCross.
func (m Message) GetCPProgram(f *field.CPProgramField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for NewOrderCross.
func (m Message) CPRegType() (*field.CPRegTypeField, errors.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from NewOrderCross.
func (m Message) GetCPRegType(f *field.CPRegTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for NewOrderCross.
func (m Message) NoEvents() (*field.NoEventsField, errors.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from NewOrderCross.
func (m Message) GetNoEvents(f *field.NoEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for NewOrderCross.
func (m Message) DatedDate() (*field.DatedDateField, errors.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from NewOrderCross.
func (m Message) GetDatedDate(f *field.DatedDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for NewOrderCross.
func (m Message) InterestAccrualDate() (*field.InterestAccrualDateField, errors.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from NewOrderCross.
func (m Message) GetInterestAccrualDate(f *field.InterestAccrualDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for NewOrderCross.
func (m Message) SecurityStatus() (*field.SecurityStatusField, errors.MessageRejectError) {
	f := &field.SecurityStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from NewOrderCross.
func (m Message) GetSecurityStatus(f *field.SecurityStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for NewOrderCross.
func (m Message) SettleOnOpenFlag() (*field.SettleOnOpenFlagField, errors.MessageRejectError) {
	f := &field.SettleOnOpenFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from NewOrderCross.
func (m Message) GetSettleOnOpenFlag(f *field.SettleOnOpenFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for NewOrderCross.
func (m Message) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethodField, errors.MessageRejectError) {
	f := &field.InstrmtAssignmentMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from NewOrderCross.
func (m Message) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for NewOrderCross.
func (m Message) StrikeMultiplier() (*field.StrikeMultiplierField, errors.MessageRejectError) {
	f := &field.StrikeMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from NewOrderCross.
func (m Message) GetStrikeMultiplier(f *field.StrikeMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for NewOrderCross.
func (m Message) StrikeValue() (*field.StrikeValueField, errors.MessageRejectError) {
	f := &field.StrikeValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from NewOrderCross.
func (m Message) GetStrikeValue(f *field.StrikeValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for NewOrderCross.
func (m Message) MinPriceIncrement() (*field.MinPriceIncrementField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from NewOrderCross.
func (m Message) GetMinPriceIncrement(f *field.MinPriceIncrementField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for NewOrderCross.
func (m Message) PositionLimit() (*field.PositionLimitField, errors.MessageRejectError) {
	f := &field.PositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from NewOrderCross.
func (m Message) GetPositionLimit(f *field.PositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for NewOrderCross.
func (m Message) NTPositionLimit() (*field.NTPositionLimitField, errors.MessageRejectError) {
	f := &field.NTPositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from NewOrderCross.
func (m Message) GetNTPositionLimit(f *field.NTPositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for NewOrderCross.
func (m Message) NoInstrumentParties() (*field.NoInstrumentPartiesField, errors.MessageRejectError) {
	f := &field.NoInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from NewOrderCross.
func (m Message) GetNoInstrumentParties(f *field.NoInstrumentPartiesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for NewOrderCross.
func (m Message) UnitOfMeasure() (*field.UnitOfMeasureField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from NewOrderCross.
func (m Message) GetUnitOfMeasure(f *field.UnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for NewOrderCross.
func (m Message) TimeUnit() (*field.TimeUnitField, errors.MessageRejectError) {
	f := &field.TimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from NewOrderCross.
func (m Message) GetTimeUnit(f *field.TimeUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for NewOrderCross.
func (m Message) MaturityTime() (*field.MaturityTimeField, errors.MessageRejectError) {
	f := &field.MaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from NewOrderCross.
func (m Message) GetMaturityTime(f *field.MaturityTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for NewOrderCross.
func (m Message) NoUnderlyings() (*field.NoUnderlyingsField, errors.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from NewOrderCross.
func (m Message) GetNoUnderlyings(f *field.NoUnderlyingsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for NewOrderCross.
func (m Message) NoLegs() (*field.NoLegsField, errors.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from NewOrderCross.
func (m Message) GetNoLegs(f *field.NoLegsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlType is a non-required field for NewOrderCross.
func (m Message) SettlType() (*field.SettlTypeField, errors.MessageRejectError) {
	f := &field.SettlTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlType reads a SettlType from NewOrderCross.
func (m Message) GetSettlType(f *field.SettlTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for NewOrderCross.
func (m Message) SettlDate() (*field.SettlDateField, errors.MessageRejectError) {
	f := &field.SettlDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from NewOrderCross.
func (m Message) GetSettlDate(f *field.SettlDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HandlInst is a non-required field for NewOrderCross.
func (m Message) HandlInst() (*field.HandlInstField, errors.MessageRejectError) {
	f := &field.HandlInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHandlInst reads a HandlInst from NewOrderCross.
func (m Message) GetHandlInst(f *field.HandlInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecInst is a non-required field for NewOrderCross.
func (m Message) ExecInst() (*field.ExecInstField, errors.MessageRejectError) {
	f := &field.ExecInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecInst reads a ExecInst from NewOrderCross.
func (m Message) GetExecInst(f *field.ExecInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinQty is a non-required field for NewOrderCross.
func (m Message) MinQty() (*field.MinQtyField, errors.MessageRejectError) {
	f := &field.MinQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinQty reads a MinQty from NewOrderCross.
func (m Message) GetMinQty(f *field.MinQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxFloor is a non-required field for NewOrderCross.
func (m Message) MaxFloor() (*field.MaxFloorField, errors.MessageRejectError) {
	f := &field.MaxFloorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxFloor reads a MaxFloor from NewOrderCross.
func (m Message) GetMaxFloor(f *field.MaxFloorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestination is a non-required field for NewOrderCross.
func (m Message) ExDestination() (*field.ExDestinationField, errors.MessageRejectError) {
	f := &field.ExDestinationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExDestination reads a ExDestination from NewOrderCross.
func (m Message) GetExDestination(f *field.ExDestinationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a non-required field for NewOrderCross.
func (m Message) NoTradingSessions() (*field.NoTradingSessionsField, errors.MessageRejectError) {
	f := &field.NoTradingSessionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from NewOrderCross.
func (m Message) GetNoTradingSessions(f *field.NoTradingSessionsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProcessCode is a non-required field for NewOrderCross.
func (m Message) ProcessCode() (*field.ProcessCodeField, errors.MessageRejectError) {
	f := &field.ProcessCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProcessCode reads a ProcessCode from NewOrderCross.
func (m Message) GetProcessCode(f *field.ProcessCodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PrevClosePx is a non-required field for NewOrderCross.
func (m Message) PrevClosePx() (*field.PrevClosePxField, errors.MessageRejectError) {
	f := &field.PrevClosePxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrevClosePx reads a PrevClosePx from NewOrderCross.
func (m Message) GetPrevClosePx(f *field.PrevClosePxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocateReqd is a non-required field for NewOrderCross.
func (m Message) LocateReqd() (*field.LocateReqdField, errors.MessageRejectError) {
	f := &field.LocateReqdField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocateReqd reads a LocateReqd from NewOrderCross.
func (m Message) GetLocateReqd(f *field.LocateReqdField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for NewOrderCross.
func (m Message) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from NewOrderCross.
func (m Message) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStipulations is a non-required field for NewOrderCross.
func (m Message) NoStipulations() (*field.NoStipulationsField, errors.MessageRejectError) {
	f := &field.NoStipulationsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoStipulations reads a NoStipulations from NewOrderCross.
func (m Message) GetNoStipulations(f *field.NoStipulationsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a required field for NewOrderCross.
func (m Message) OrdType() (*field.OrdTypeField, errors.MessageRejectError) {
	f := &field.OrdTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from NewOrderCross.
func (m Message) GetOrdType(f *field.OrdTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for NewOrderCross.
func (m Message) PriceType() (*field.PriceTypeField, errors.MessageRejectError) {
	f := &field.PriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from NewOrderCross.
func (m Message) GetPriceType(f *field.PriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for NewOrderCross.
func (m Message) Price() (*field.PriceField, errors.MessageRejectError) {
	f := &field.PriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from NewOrderCross.
func (m Message) GetPrice(f *field.PriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StopPx is a non-required field for NewOrderCross.
func (m Message) StopPx() (*field.StopPxField, errors.MessageRejectError) {
	f := &field.StopPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStopPx reads a StopPx from NewOrderCross.
func (m Message) GetStopPx(f *field.StopPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for NewOrderCross.
func (m Message) Spread() (*field.SpreadField, errors.MessageRejectError) {
	f := &field.SpreadField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from NewOrderCross.
func (m Message) GetSpread(f *field.SpreadField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for NewOrderCross.
func (m Message) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrencyField, errors.MessageRejectError) {
	f := &field.BenchmarkCurveCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from NewOrderCross.
func (m Message) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for NewOrderCross.
func (m Message) BenchmarkCurveName() (*field.BenchmarkCurveNameField, errors.MessageRejectError) {
	f := &field.BenchmarkCurveNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from NewOrderCross.
func (m Message) GetBenchmarkCurveName(f *field.BenchmarkCurveNameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for NewOrderCross.
func (m Message) BenchmarkCurvePoint() (*field.BenchmarkCurvePointField, errors.MessageRejectError) {
	f := &field.BenchmarkCurvePointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from NewOrderCross.
func (m Message) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPrice is a non-required field for NewOrderCross.
func (m Message) BenchmarkPrice() (*field.BenchmarkPriceField, errors.MessageRejectError) {
	f := &field.BenchmarkPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPrice reads a BenchmarkPrice from NewOrderCross.
func (m Message) GetBenchmarkPrice(f *field.BenchmarkPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPriceType is a non-required field for NewOrderCross.
func (m Message) BenchmarkPriceType() (*field.BenchmarkPriceTypeField, errors.MessageRejectError) {
	f := &field.BenchmarkPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPriceType reads a BenchmarkPriceType from NewOrderCross.
func (m Message) GetBenchmarkPriceType(f *field.BenchmarkPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityID is a non-required field for NewOrderCross.
func (m Message) BenchmarkSecurityID() (*field.BenchmarkSecurityIDField, errors.MessageRejectError) {
	f := &field.BenchmarkSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityID reads a BenchmarkSecurityID from NewOrderCross.
func (m Message) GetBenchmarkSecurityID(f *field.BenchmarkSecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityIDSource is a non-required field for NewOrderCross.
func (m Message) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSourceField, errors.MessageRejectError) {
	f := &field.BenchmarkSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityIDSource reads a BenchmarkSecurityIDSource from NewOrderCross.
func (m Message) GetBenchmarkSecurityIDSource(f *field.BenchmarkSecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldType is a non-required field for NewOrderCross.
func (m Message) YieldType() (*field.YieldTypeField, errors.MessageRejectError) {
	f := &field.YieldTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldType reads a YieldType from NewOrderCross.
func (m Message) GetYieldType(f *field.YieldTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Yield is a non-required field for NewOrderCross.
func (m Message) Yield() (*field.YieldField, errors.MessageRejectError) {
	f := &field.YieldField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYield reads a Yield from NewOrderCross.
func (m Message) GetYield(f *field.YieldField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldCalcDate is a non-required field for NewOrderCross.
func (m Message) YieldCalcDate() (*field.YieldCalcDateField, errors.MessageRejectError) {
	f := &field.YieldCalcDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldCalcDate reads a YieldCalcDate from NewOrderCross.
func (m Message) GetYieldCalcDate(f *field.YieldCalcDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionDate is a non-required field for NewOrderCross.
func (m Message) YieldRedemptionDate() (*field.YieldRedemptionDateField, errors.MessageRejectError) {
	f := &field.YieldRedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionDate reads a YieldRedemptionDate from NewOrderCross.
func (m Message) GetYieldRedemptionDate(f *field.YieldRedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPrice is a non-required field for NewOrderCross.
func (m Message) YieldRedemptionPrice() (*field.YieldRedemptionPriceField, errors.MessageRejectError) {
	f := &field.YieldRedemptionPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPrice reads a YieldRedemptionPrice from NewOrderCross.
func (m Message) GetYieldRedemptionPrice(f *field.YieldRedemptionPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPriceType is a non-required field for NewOrderCross.
func (m Message) YieldRedemptionPriceType() (*field.YieldRedemptionPriceTypeField, errors.MessageRejectError) {
	f := &field.YieldRedemptionPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPriceType reads a YieldRedemptionPriceType from NewOrderCross.
func (m Message) GetYieldRedemptionPriceType(f *field.YieldRedemptionPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for NewOrderCross.
func (m Message) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from NewOrderCross.
func (m Message) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ComplianceID is a non-required field for NewOrderCross.
func (m Message) ComplianceID() (*field.ComplianceIDField, errors.MessageRejectError) {
	f := &field.ComplianceIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetComplianceID reads a ComplianceID from NewOrderCross.
func (m Message) GetComplianceID(f *field.ComplianceIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIID is a non-required field for NewOrderCross.
func (m Message) IOIID() (*field.IOIIDField, errors.MessageRejectError) {
	f := &field.IOIIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIID reads a IOIID from NewOrderCross.
func (m Message) GetIOIID(f *field.IOIIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a non-required field for NewOrderCross.
func (m Message) QuoteID() (*field.QuoteIDField, errors.MessageRejectError) {
	f := &field.QuoteIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from NewOrderCross.
func (m Message) GetQuoteID(f *field.QuoteIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeInForce is a non-required field for NewOrderCross.
func (m Message) TimeInForce() (*field.TimeInForceField, errors.MessageRejectError) {
	f := &field.TimeInForceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeInForce reads a TimeInForce from NewOrderCross.
func (m Message) GetTimeInForce(f *field.TimeInForceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for NewOrderCross.
func (m Message) EffectiveTime() (*field.EffectiveTimeField, errors.MessageRejectError) {
	f := &field.EffectiveTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from NewOrderCross.
func (m Message) GetEffectiveTime(f *field.EffectiveTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireDate is a non-required field for NewOrderCross.
func (m Message) ExpireDate() (*field.ExpireDateField, errors.MessageRejectError) {
	f := &field.ExpireDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireDate reads a ExpireDate from NewOrderCross.
func (m Message) GetExpireDate(f *field.ExpireDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for NewOrderCross.
func (m Message) ExpireTime() (*field.ExpireTimeField, errors.MessageRejectError) {
	f := &field.ExpireTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from NewOrderCross.
func (m Message) GetExpireTime(f *field.ExpireTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GTBookingInst is a non-required field for NewOrderCross.
func (m Message) GTBookingInst() (*field.GTBookingInstField, errors.MessageRejectError) {
	f := &field.GTBookingInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGTBookingInst reads a GTBookingInst from NewOrderCross.
func (m Message) GetGTBookingInst(f *field.GTBookingInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxShow is a non-required field for NewOrderCross.
func (m Message) MaxShow() (*field.MaxShowField, errors.MessageRejectError) {
	f := &field.MaxShowField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxShow reads a MaxShow from NewOrderCross.
func (m Message) GetMaxShow(f *field.MaxShowField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegOffsetValue is a non-required field for NewOrderCross.
func (m Message) PegOffsetValue() (*field.PegOffsetValueField, errors.MessageRejectError) {
	f := &field.PegOffsetValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegOffsetValue reads a PegOffsetValue from NewOrderCross.
func (m Message) GetPegOffsetValue(f *field.PegOffsetValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegMoveType is a non-required field for NewOrderCross.
func (m Message) PegMoveType() (*field.PegMoveTypeField, errors.MessageRejectError) {
	f := &field.PegMoveTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegMoveType reads a PegMoveType from NewOrderCross.
func (m Message) GetPegMoveType(f *field.PegMoveTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegOffsetType is a non-required field for NewOrderCross.
func (m Message) PegOffsetType() (*field.PegOffsetTypeField, errors.MessageRejectError) {
	f := &field.PegOffsetTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegOffsetType reads a PegOffsetType from NewOrderCross.
func (m Message) GetPegOffsetType(f *field.PegOffsetTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegLimitType is a non-required field for NewOrderCross.
func (m Message) PegLimitType() (*field.PegLimitTypeField, errors.MessageRejectError) {
	f := &field.PegLimitTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegLimitType reads a PegLimitType from NewOrderCross.
func (m Message) GetPegLimitType(f *field.PegLimitTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegRoundDirection is a non-required field for NewOrderCross.
func (m Message) PegRoundDirection() (*field.PegRoundDirectionField, errors.MessageRejectError) {
	f := &field.PegRoundDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegRoundDirection reads a PegRoundDirection from NewOrderCross.
func (m Message) GetPegRoundDirection(f *field.PegRoundDirectionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegScope is a non-required field for NewOrderCross.
func (m Message) PegScope() (*field.PegScopeField, errors.MessageRejectError) {
	f := &field.PegScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegScope reads a PegScope from NewOrderCross.
func (m Message) GetPegScope(f *field.PegScopeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegPriceType is a non-required field for NewOrderCross.
func (m Message) PegPriceType() (*field.PegPriceTypeField, errors.MessageRejectError) {
	f := &field.PegPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegPriceType reads a PegPriceType from NewOrderCross.
func (m Message) GetPegPriceType(f *field.PegPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityIDSource is a non-required field for NewOrderCross.
func (m Message) PegSecurityIDSource() (*field.PegSecurityIDSourceField, errors.MessageRejectError) {
	f := &field.PegSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityIDSource reads a PegSecurityIDSource from NewOrderCross.
func (m Message) GetPegSecurityIDSource(f *field.PegSecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityID is a non-required field for NewOrderCross.
func (m Message) PegSecurityID() (*field.PegSecurityIDField, errors.MessageRejectError) {
	f := &field.PegSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityID reads a PegSecurityID from NewOrderCross.
func (m Message) GetPegSecurityID(f *field.PegSecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSymbol is a non-required field for NewOrderCross.
func (m Message) PegSymbol() (*field.PegSymbolField, errors.MessageRejectError) {
	f := &field.PegSymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegSymbol reads a PegSymbol from NewOrderCross.
func (m Message) GetPegSymbol(f *field.PegSymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityDesc is a non-required field for NewOrderCross.
func (m Message) PegSecurityDesc() (*field.PegSecurityDescField, errors.MessageRejectError) {
	f := &field.PegSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityDesc reads a PegSecurityDesc from NewOrderCross.
func (m Message) GetPegSecurityDesc(f *field.PegSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionInst is a non-required field for NewOrderCross.
func (m Message) DiscretionInst() (*field.DiscretionInstField, errors.MessageRejectError) {
	f := &field.DiscretionInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionInst reads a DiscretionInst from NewOrderCross.
func (m Message) GetDiscretionInst(f *field.DiscretionInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffsetValue is a non-required field for NewOrderCross.
func (m Message) DiscretionOffsetValue() (*field.DiscretionOffsetValueField, errors.MessageRejectError) {
	f := &field.DiscretionOffsetValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffsetValue reads a DiscretionOffsetValue from NewOrderCross.
func (m Message) GetDiscretionOffsetValue(f *field.DiscretionOffsetValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionMoveType is a non-required field for NewOrderCross.
func (m Message) DiscretionMoveType() (*field.DiscretionMoveTypeField, errors.MessageRejectError) {
	f := &field.DiscretionMoveTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionMoveType reads a DiscretionMoveType from NewOrderCross.
func (m Message) GetDiscretionMoveType(f *field.DiscretionMoveTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffsetType is a non-required field for NewOrderCross.
func (m Message) DiscretionOffsetType() (*field.DiscretionOffsetTypeField, errors.MessageRejectError) {
	f := &field.DiscretionOffsetTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffsetType reads a DiscretionOffsetType from NewOrderCross.
func (m Message) GetDiscretionOffsetType(f *field.DiscretionOffsetTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionLimitType is a non-required field for NewOrderCross.
func (m Message) DiscretionLimitType() (*field.DiscretionLimitTypeField, errors.MessageRejectError) {
	f := &field.DiscretionLimitTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionLimitType reads a DiscretionLimitType from NewOrderCross.
func (m Message) GetDiscretionLimitType(f *field.DiscretionLimitTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionRoundDirection is a non-required field for NewOrderCross.
func (m Message) DiscretionRoundDirection() (*field.DiscretionRoundDirectionField, errors.MessageRejectError) {
	f := &field.DiscretionRoundDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionRoundDirection reads a DiscretionRoundDirection from NewOrderCross.
func (m Message) GetDiscretionRoundDirection(f *field.DiscretionRoundDirectionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionScope is a non-required field for NewOrderCross.
func (m Message) DiscretionScope() (*field.DiscretionScopeField, errors.MessageRejectError) {
	f := &field.DiscretionScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionScope reads a DiscretionScope from NewOrderCross.
func (m Message) GetDiscretionScope(f *field.DiscretionScopeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TargetStrategy is a non-required field for NewOrderCross.
func (m Message) TargetStrategy() (*field.TargetStrategyField, errors.MessageRejectError) {
	f := &field.TargetStrategyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTargetStrategy reads a TargetStrategy from NewOrderCross.
func (m Message) GetTargetStrategy(f *field.TargetStrategyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TargetStrategyParameters is a non-required field for NewOrderCross.
func (m Message) TargetStrategyParameters() (*field.TargetStrategyParametersField, errors.MessageRejectError) {
	f := &field.TargetStrategyParametersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTargetStrategyParameters reads a TargetStrategyParameters from NewOrderCross.
func (m Message) GetTargetStrategyParameters(f *field.TargetStrategyParametersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ParticipationRate is a non-required field for NewOrderCross.
func (m Message) ParticipationRate() (*field.ParticipationRateField, errors.MessageRejectError) {
	f := &field.ParticipationRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetParticipationRate reads a ParticipationRate from NewOrderCross.
func (m Message) GetParticipationRate(f *field.ParticipationRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CancellationRights is a non-required field for NewOrderCross.
func (m Message) CancellationRights() (*field.CancellationRightsField, errors.MessageRejectError) {
	f := &field.CancellationRightsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCancellationRights reads a CancellationRights from NewOrderCross.
func (m Message) GetCancellationRights(f *field.CancellationRightsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MoneyLaunderingStatus is a non-required field for NewOrderCross.
func (m Message) MoneyLaunderingStatus() (*field.MoneyLaunderingStatusField, errors.MessageRejectError) {
	f := &field.MoneyLaunderingStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMoneyLaunderingStatus reads a MoneyLaunderingStatus from NewOrderCross.
func (m Message) GetMoneyLaunderingStatus(f *field.MoneyLaunderingStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistID is a non-required field for NewOrderCross.
func (m Message) RegistID() (*field.RegistIDField, errors.MessageRejectError) {
	f := &field.RegistIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from NewOrderCross.
func (m Message) GetRegistID(f *field.RegistIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Designation is a non-required field for NewOrderCross.
func (m Message) Designation() (*field.DesignationField, errors.MessageRejectError) {
	f := &field.DesignationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDesignation reads a Designation from NewOrderCross.
func (m Message) GetDesignation(f *field.DesignationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStrategyParameters is a non-required field for NewOrderCross.
func (m Message) NoStrategyParameters() (*field.NoStrategyParametersField, errors.MessageRejectError) {
	f := &field.NoStrategyParametersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoStrategyParameters reads a NoStrategyParameters from NewOrderCross.
func (m Message) GetNoStrategyParameters(f *field.NoStrategyParametersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransBkdTime is a non-required field for NewOrderCross.
func (m Message) TransBkdTime() (*field.TransBkdTimeField, errors.MessageRejectError) {
	f := &field.TransBkdTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransBkdTime reads a TransBkdTime from NewOrderCross.
func (m Message) GetTransBkdTime(f *field.TransBkdTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRootPartyIDs is a non-required field for NewOrderCross.
func (m Message) NoRootPartyIDs() (*field.NoRootPartyIDsField, errors.MessageRejectError) {
	f := &field.NoRootPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRootPartyIDs reads a NoRootPartyIDs from NewOrderCross.
func (m Message) GetNoRootPartyIDs(f *field.NoRootPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchIncrement is a non-required field for NewOrderCross.
func (m Message) MatchIncrement() (*field.MatchIncrementField, errors.MessageRejectError) {
	f := &field.MatchIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMatchIncrement reads a MatchIncrement from NewOrderCross.
func (m Message) GetMatchIncrement(f *field.MatchIncrementField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxPriceLevels is a non-required field for NewOrderCross.
func (m Message) MaxPriceLevels() (*field.MaxPriceLevelsField, errors.MessageRejectError) {
	f := &field.MaxPriceLevelsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxPriceLevels reads a MaxPriceLevels from NewOrderCross.
func (m Message) GetMaxPriceLevels(f *field.MaxPriceLevelsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryDisplayQty is a non-required field for NewOrderCross.
func (m Message) SecondaryDisplayQty() (*field.SecondaryDisplayQtyField, errors.MessageRejectError) {
	f := &field.SecondaryDisplayQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryDisplayQty reads a SecondaryDisplayQty from NewOrderCross.
func (m Message) GetSecondaryDisplayQty(f *field.SecondaryDisplayQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayWhen is a non-required field for NewOrderCross.
func (m Message) DisplayWhen() (*field.DisplayWhenField, errors.MessageRejectError) {
	f := &field.DisplayWhenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayWhen reads a DisplayWhen from NewOrderCross.
func (m Message) GetDisplayWhen(f *field.DisplayWhenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayMethod is a non-required field for NewOrderCross.
func (m Message) DisplayMethod() (*field.DisplayMethodField, errors.MessageRejectError) {
	f := &field.DisplayMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayMethod reads a DisplayMethod from NewOrderCross.
func (m Message) GetDisplayMethod(f *field.DisplayMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayLowQty is a non-required field for NewOrderCross.
func (m Message) DisplayLowQty() (*field.DisplayLowQtyField, errors.MessageRejectError) {
	f := &field.DisplayLowQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayLowQty reads a DisplayLowQty from NewOrderCross.
func (m Message) GetDisplayLowQty(f *field.DisplayLowQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayHighQty is a non-required field for NewOrderCross.
func (m Message) DisplayHighQty() (*field.DisplayHighQtyField, errors.MessageRejectError) {
	f := &field.DisplayHighQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayHighQty reads a DisplayHighQty from NewOrderCross.
func (m Message) GetDisplayHighQty(f *field.DisplayHighQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayMinIncr is a non-required field for NewOrderCross.
func (m Message) DisplayMinIncr() (*field.DisplayMinIncrField, errors.MessageRejectError) {
	f := &field.DisplayMinIncrField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayMinIncr reads a DisplayMinIncr from NewOrderCross.
func (m Message) GetDisplayMinIncr(f *field.DisplayMinIncrField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefreshQty is a non-required field for NewOrderCross.
func (m Message) RefreshQty() (*field.RefreshQtyField, errors.MessageRejectError) {
	f := &field.RefreshQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRefreshQty reads a RefreshQty from NewOrderCross.
func (m Message) GetRefreshQty(f *field.RefreshQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayQty is a non-required field for NewOrderCross.
func (m Message) DisplayQty() (*field.DisplayQtyField, errors.MessageRejectError) {
	f := &field.DisplayQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayQty reads a DisplayQty from NewOrderCross.
func (m Message) GetDisplayQty(f *field.DisplayQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceProtectionScope is a non-required field for NewOrderCross.
func (m Message) PriceProtectionScope() (*field.PriceProtectionScopeField, errors.MessageRejectError) {
	f := &field.PriceProtectionScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceProtectionScope reads a PriceProtectionScope from NewOrderCross.
func (m Message) GetPriceProtectionScope(f *field.PriceProtectionScopeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerType is a non-required field for NewOrderCross.
func (m Message) TriggerType() (*field.TriggerTypeField, errors.MessageRejectError) {
	f := &field.TriggerTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerType reads a TriggerType from NewOrderCross.
func (m Message) GetTriggerType(f *field.TriggerTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerAction is a non-required field for NewOrderCross.
func (m Message) TriggerAction() (*field.TriggerActionField, errors.MessageRejectError) {
	f := &field.TriggerActionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerAction reads a TriggerAction from NewOrderCross.
func (m Message) GetTriggerAction(f *field.TriggerActionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPrice is a non-required field for NewOrderCross.
func (m Message) TriggerPrice() (*field.TriggerPriceField, errors.MessageRejectError) {
	f := &field.TriggerPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPrice reads a TriggerPrice from NewOrderCross.
func (m Message) GetTriggerPrice(f *field.TriggerPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSymbol is a non-required field for NewOrderCross.
func (m Message) TriggerSymbol() (*field.TriggerSymbolField, errors.MessageRejectError) {
	f := &field.TriggerSymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSymbol reads a TriggerSymbol from NewOrderCross.
func (m Message) GetTriggerSymbol(f *field.TriggerSymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityID is a non-required field for NewOrderCross.
func (m Message) TriggerSecurityID() (*field.TriggerSecurityIDField, errors.MessageRejectError) {
	f := &field.TriggerSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityID reads a TriggerSecurityID from NewOrderCross.
func (m Message) GetTriggerSecurityID(f *field.TriggerSecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityIDSource is a non-required field for NewOrderCross.
func (m Message) TriggerSecurityIDSource() (*field.TriggerSecurityIDSourceField, errors.MessageRejectError) {
	f := &field.TriggerSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityIDSource reads a TriggerSecurityIDSource from NewOrderCross.
func (m Message) GetTriggerSecurityIDSource(f *field.TriggerSecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityDesc is a non-required field for NewOrderCross.
func (m Message) TriggerSecurityDesc() (*field.TriggerSecurityDescField, errors.MessageRejectError) {
	f := &field.TriggerSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityDesc reads a TriggerSecurityDesc from NewOrderCross.
func (m Message) GetTriggerSecurityDesc(f *field.TriggerSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceType is a non-required field for NewOrderCross.
func (m Message) TriggerPriceType() (*field.TriggerPriceTypeField, errors.MessageRejectError) {
	f := &field.TriggerPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceType reads a TriggerPriceType from NewOrderCross.
func (m Message) GetTriggerPriceType(f *field.TriggerPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceTypeScope is a non-required field for NewOrderCross.
func (m Message) TriggerPriceTypeScope() (*field.TriggerPriceTypeScopeField, errors.MessageRejectError) {
	f := &field.TriggerPriceTypeScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceTypeScope reads a TriggerPriceTypeScope from NewOrderCross.
func (m Message) GetTriggerPriceTypeScope(f *field.TriggerPriceTypeScopeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceDirection is a non-required field for NewOrderCross.
func (m Message) TriggerPriceDirection() (*field.TriggerPriceDirectionField, errors.MessageRejectError) {
	f := &field.TriggerPriceDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceDirection reads a TriggerPriceDirection from NewOrderCross.
func (m Message) GetTriggerPriceDirection(f *field.TriggerPriceDirectionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerNewPrice is a non-required field for NewOrderCross.
func (m Message) TriggerNewPrice() (*field.TriggerNewPriceField, errors.MessageRejectError) {
	f := &field.TriggerNewPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerNewPrice reads a TriggerNewPrice from NewOrderCross.
func (m Message) GetTriggerNewPrice(f *field.TriggerNewPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerOrderType is a non-required field for NewOrderCross.
func (m Message) TriggerOrderType() (*field.TriggerOrderTypeField, errors.MessageRejectError) {
	f := &field.TriggerOrderTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerOrderType reads a TriggerOrderType from NewOrderCross.
func (m Message) GetTriggerOrderType(f *field.TriggerOrderTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerNewQty is a non-required field for NewOrderCross.
func (m Message) TriggerNewQty() (*field.TriggerNewQtyField, errors.MessageRejectError) {
	f := &field.TriggerNewQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerNewQty reads a TriggerNewQty from NewOrderCross.
func (m Message) GetTriggerNewQty(f *field.TriggerNewQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerTradingSessionID is a non-required field for NewOrderCross.
func (m Message) TriggerTradingSessionID() (*field.TriggerTradingSessionIDField, errors.MessageRejectError) {
	f := &field.TriggerTradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerTradingSessionID reads a TriggerTradingSessionID from NewOrderCross.
func (m Message) GetTriggerTradingSessionID(f *field.TriggerTradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerTradingSessionSubID is a non-required field for NewOrderCross.
func (m Message) TriggerTradingSessionSubID() (*field.TriggerTradingSessionSubIDField, errors.MessageRejectError) {
	f := &field.TriggerTradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerTradingSessionSubID reads a TriggerTradingSessionSubID from NewOrderCross.
func (m Message) GetTriggerTradingSessionSubID(f *field.TriggerTradingSessionSubIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestinationIDSource is a non-required field for NewOrderCross.
func (m Message) ExDestinationIDSource() (*field.ExDestinationIDSourceField, errors.MessageRejectError) {
	f := &field.ExDestinationIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExDestinationIDSource reads a ExDestinationIDSource from NewOrderCross.
func (m Message) GetExDestinationIDSource(f *field.ExDestinationIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds NewOrderCross messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for NewOrderCross.
func Builder(
	crossid *field.CrossIDField,
	crosstype *field.CrossTypeField,
	crossprioritization *field.CrossPrioritizationField,
	nosides *field.NoSidesField,
	transacttime *field.TransactTimeField,
	ordtype *field.OrdTypeField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50))
	builder.Header().Set(field.NewMsgType("s"))
	builder.Body().Set(crossid)
	builder.Body().Set(crosstype)
	builder.Body().Set(crossprioritization)
	builder.Body().Set(nosides)
	builder.Body().Set(transacttime)
	builder.Body().Set(ordtype)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX50, "s", r
}
