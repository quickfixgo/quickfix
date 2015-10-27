//Package ordermasscancelrequest msg type = q.
package ordermasscancelrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
)

//Message is a OrderMassCancelRequest wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//ClOrdID is a required field for OrderMassCancelRequest.
func (m Message) ClOrdID() (*field.ClOrdIDField, quickfix.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from OrderMassCancelRequest.
func (m Message) GetClOrdID(f *field.ClOrdIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for OrderMassCancelRequest.
func (m Message) SecondaryClOrdID() (*field.SecondaryClOrdIDField, quickfix.MessageRejectError) {
	f := &field.SecondaryClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from OrderMassCancelRequest.
func (m Message) GetSecondaryClOrdID(f *field.SecondaryClOrdIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MassCancelRequestType is a required field for OrderMassCancelRequest.
func (m Message) MassCancelRequestType() (*field.MassCancelRequestTypeField, quickfix.MessageRejectError) {
	f := &field.MassCancelRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMassCancelRequestType reads a MassCancelRequestType from OrderMassCancelRequest.
func (m Message) GetMassCancelRequestType(f *field.MassCancelRequestTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for OrderMassCancelRequest.
func (m Message) TradingSessionID() (*field.TradingSessionIDField, quickfix.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from OrderMassCancelRequest.
func (m Message) GetTradingSessionID(f *field.TradingSessionIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for OrderMassCancelRequest.
func (m Message) TradingSessionSubID() (*field.TradingSessionSubIDField, quickfix.MessageRejectError) {
	f := &field.TradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from OrderMassCancelRequest.
func (m Message) GetTradingSessionSubID(f *field.TradingSessionSubIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for OrderMassCancelRequest.
func (m Message) Symbol() (*field.SymbolField, quickfix.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from OrderMassCancelRequest.
func (m Message) GetSymbol(f *field.SymbolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for OrderMassCancelRequest.
func (m Message) SymbolSfx() (*field.SymbolSfxField, quickfix.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from OrderMassCancelRequest.
func (m Message) GetSymbolSfx(f *field.SymbolSfxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for OrderMassCancelRequest.
func (m Message) SecurityID() (*field.SecurityIDField, quickfix.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from OrderMassCancelRequest.
func (m Message) GetSecurityID(f *field.SecurityIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for OrderMassCancelRequest.
func (m Message) SecurityIDSource() (*field.SecurityIDSourceField, quickfix.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from OrderMassCancelRequest.
func (m Message) GetSecurityIDSource(f *field.SecurityIDSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for OrderMassCancelRequest.
func (m Message) NoSecurityAltID() (*field.NoSecurityAltIDField, quickfix.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from OrderMassCancelRequest.
func (m Message) GetNoSecurityAltID(f *field.NoSecurityAltIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for OrderMassCancelRequest.
func (m Message) Product() (*field.ProductField, quickfix.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from OrderMassCancelRequest.
func (m Message) GetProduct(f *field.ProductField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for OrderMassCancelRequest.
func (m Message) CFICode() (*field.CFICodeField, quickfix.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from OrderMassCancelRequest.
func (m Message) GetCFICode(f *field.CFICodeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for OrderMassCancelRequest.
func (m Message) SecurityType() (*field.SecurityTypeField, quickfix.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from OrderMassCancelRequest.
func (m Message) GetSecurityType(f *field.SecurityTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for OrderMassCancelRequest.
func (m Message) SecuritySubType() (*field.SecuritySubTypeField, quickfix.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from OrderMassCancelRequest.
func (m Message) GetSecuritySubType(f *field.SecuritySubTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for OrderMassCancelRequest.
func (m Message) MaturityMonthYear() (*field.MaturityMonthYearField, quickfix.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from OrderMassCancelRequest.
func (m Message) GetMaturityMonthYear(f *field.MaturityMonthYearField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for OrderMassCancelRequest.
func (m Message) MaturityDate() (*field.MaturityDateField, quickfix.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from OrderMassCancelRequest.
func (m Message) GetMaturityDate(f *field.MaturityDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for OrderMassCancelRequest.
func (m Message) CouponPaymentDate() (*field.CouponPaymentDateField, quickfix.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from OrderMassCancelRequest.
func (m Message) GetCouponPaymentDate(f *field.CouponPaymentDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for OrderMassCancelRequest.
func (m Message) IssueDate() (*field.IssueDateField, quickfix.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from OrderMassCancelRequest.
func (m Message) GetIssueDate(f *field.IssueDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for OrderMassCancelRequest.
func (m Message) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, quickfix.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from OrderMassCancelRequest.
func (m Message) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for OrderMassCancelRequest.
func (m Message) RepurchaseTerm() (*field.RepurchaseTermField, quickfix.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from OrderMassCancelRequest.
func (m Message) GetRepurchaseTerm(f *field.RepurchaseTermField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for OrderMassCancelRequest.
func (m Message) RepurchaseRate() (*field.RepurchaseRateField, quickfix.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from OrderMassCancelRequest.
func (m Message) GetRepurchaseRate(f *field.RepurchaseRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for OrderMassCancelRequest.
func (m Message) Factor() (*field.FactorField, quickfix.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from OrderMassCancelRequest.
func (m Message) GetFactor(f *field.FactorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for OrderMassCancelRequest.
func (m Message) CreditRating() (*field.CreditRatingField, quickfix.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from OrderMassCancelRequest.
func (m Message) GetCreditRating(f *field.CreditRatingField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for OrderMassCancelRequest.
func (m Message) InstrRegistry() (*field.InstrRegistryField, quickfix.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from OrderMassCancelRequest.
func (m Message) GetInstrRegistry(f *field.InstrRegistryField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for OrderMassCancelRequest.
func (m Message) CountryOfIssue() (*field.CountryOfIssueField, quickfix.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from OrderMassCancelRequest.
func (m Message) GetCountryOfIssue(f *field.CountryOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for OrderMassCancelRequest.
func (m Message) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, quickfix.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from OrderMassCancelRequest.
func (m Message) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for OrderMassCancelRequest.
func (m Message) LocaleOfIssue() (*field.LocaleOfIssueField, quickfix.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from OrderMassCancelRequest.
func (m Message) GetLocaleOfIssue(f *field.LocaleOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for OrderMassCancelRequest.
func (m Message) RedemptionDate() (*field.RedemptionDateField, quickfix.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from OrderMassCancelRequest.
func (m Message) GetRedemptionDate(f *field.RedemptionDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for OrderMassCancelRequest.
func (m Message) StrikePrice() (*field.StrikePriceField, quickfix.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from OrderMassCancelRequest.
func (m Message) GetStrikePrice(f *field.StrikePriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for OrderMassCancelRequest.
func (m Message) StrikeCurrency() (*field.StrikeCurrencyField, quickfix.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from OrderMassCancelRequest.
func (m Message) GetStrikeCurrency(f *field.StrikeCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for OrderMassCancelRequest.
func (m Message) OptAttribute() (*field.OptAttributeField, quickfix.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from OrderMassCancelRequest.
func (m Message) GetOptAttribute(f *field.OptAttributeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for OrderMassCancelRequest.
func (m Message) ContractMultiplier() (*field.ContractMultiplierField, quickfix.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from OrderMassCancelRequest.
func (m Message) GetContractMultiplier(f *field.ContractMultiplierField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for OrderMassCancelRequest.
func (m Message) CouponRate() (*field.CouponRateField, quickfix.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from OrderMassCancelRequest.
func (m Message) GetCouponRate(f *field.CouponRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for OrderMassCancelRequest.
func (m Message) SecurityExchange() (*field.SecurityExchangeField, quickfix.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from OrderMassCancelRequest.
func (m Message) GetSecurityExchange(f *field.SecurityExchangeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for OrderMassCancelRequest.
func (m Message) Issuer() (*field.IssuerField, quickfix.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from OrderMassCancelRequest.
func (m Message) GetIssuer(f *field.IssuerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for OrderMassCancelRequest.
func (m Message) EncodedIssuerLen() (*field.EncodedIssuerLenField, quickfix.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from OrderMassCancelRequest.
func (m Message) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for OrderMassCancelRequest.
func (m Message) EncodedIssuer() (*field.EncodedIssuerField, quickfix.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from OrderMassCancelRequest.
func (m Message) GetEncodedIssuer(f *field.EncodedIssuerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for OrderMassCancelRequest.
func (m Message) SecurityDesc() (*field.SecurityDescField, quickfix.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from OrderMassCancelRequest.
func (m Message) GetSecurityDesc(f *field.SecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for OrderMassCancelRequest.
func (m Message) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, quickfix.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from OrderMassCancelRequest.
func (m Message) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for OrderMassCancelRequest.
func (m Message) EncodedSecurityDesc() (*field.EncodedSecurityDescField, quickfix.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from OrderMassCancelRequest.
func (m Message) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for OrderMassCancelRequest.
func (m Message) Pool() (*field.PoolField, quickfix.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from OrderMassCancelRequest.
func (m Message) GetPool(f *field.PoolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for OrderMassCancelRequest.
func (m Message) ContractSettlMonth() (*field.ContractSettlMonthField, quickfix.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from OrderMassCancelRequest.
func (m Message) GetContractSettlMonth(f *field.ContractSettlMonthField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for OrderMassCancelRequest.
func (m Message) CPProgram() (*field.CPProgramField, quickfix.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from OrderMassCancelRequest.
func (m Message) GetCPProgram(f *field.CPProgramField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for OrderMassCancelRequest.
func (m Message) CPRegType() (*field.CPRegTypeField, quickfix.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from OrderMassCancelRequest.
func (m Message) GetCPRegType(f *field.CPRegTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for OrderMassCancelRequest.
func (m Message) NoEvents() (*field.NoEventsField, quickfix.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from OrderMassCancelRequest.
func (m Message) GetNoEvents(f *field.NoEventsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for OrderMassCancelRequest.
func (m Message) DatedDate() (*field.DatedDateField, quickfix.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from OrderMassCancelRequest.
func (m Message) GetDatedDate(f *field.DatedDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for OrderMassCancelRequest.
func (m Message) InterestAccrualDate() (*field.InterestAccrualDateField, quickfix.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from OrderMassCancelRequest.
func (m Message) GetInterestAccrualDate(f *field.InterestAccrualDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for OrderMassCancelRequest.
func (m Message) SecurityStatus() (*field.SecurityStatusField, quickfix.MessageRejectError) {
	f := &field.SecurityStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from OrderMassCancelRequest.
func (m Message) GetSecurityStatus(f *field.SecurityStatusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for OrderMassCancelRequest.
func (m Message) SettleOnOpenFlag() (*field.SettleOnOpenFlagField, quickfix.MessageRejectError) {
	f := &field.SettleOnOpenFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from OrderMassCancelRequest.
func (m Message) GetSettleOnOpenFlag(f *field.SettleOnOpenFlagField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for OrderMassCancelRequest.
func (m Message) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethodField, quickfix.MessageRejectError) {
	f := &field.InstrmtAssignmentMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from OrderMassCancelRequest.
func (m Message) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for OrderMassCancelRequest.
func (m Message) StrikeMultiplier() (*field.StrikeMultiplierField, quickfix.MessageRejectError) {
	f := &field.StrikeMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from OrderMassCancelRequest.
func (m Message) GetStrikeMultiplier(f *field.StrikeMultiplierField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for OrderMassCancelRequest.
func (m Message) StrikeValue() (*field.StrikeValueField, quickfix.MessageRejectError) {
	f := &field.StrikeValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from OrderMassCancelRequest.
func (m Message) GetStrikeValue(f *field.StrikeValueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for OrderMassCancelRequest.
func (m Message) MinPriceIncrement() (*field.MinPriceIncrementField, quickfix.MessageRejectError) {
	f := &field.MinPriceIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from OrderMassCancelRequest.
func (m Message) GetMinPriceIncrement(f *field.MinPriceIncrementField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for OrderMassCancelRequest.
func (m Message) PositionLimit() (*field.PositionLimitField, quickfix.MessageRejectError) {
	f := &field.PositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from OrderMassCancelRequest.
func (m Message) GetPositionLimit(f *field.PositionLimitField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for OrderMassCancelRequest.
func (m Message) NTPositionLimit() (*field.NTPositionLimitField, quickfix.MessageRejectError) {
	f := &field.NTPositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from OrderMassCancelRequest.
func (m Message) GetNTPositionLimit(f *field.NTPositionLimitField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for OrderMassCancelRequest.
func (m Message) NoInstrumentParties() (*field.NoInstrumentPartiesField, quickfix.MessageRejectError) {
	f := &field.NoInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from OrderMassCancelRequest.
func (m Message) GetNoInstrumentParties(f *field.NoInstrumentPartiesField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for OrderMassCancelRequest.
func (m Message) UnitOfMeasure() (*field.UnitOfMeasureField, quickfix.MessageRejectError) {
	f := &field.UnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from OrderMassCancelRequest.
func (m Message) GetUnitOfMeasure(f *field.UnitOfMeasureField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for OrderMassCancelRequest.
func (m Message) TimeUnit() (*field.TimeUnitField, quickfix.MessageRejectError) {
	f := &field.TimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from OrderMassCancelRequest.
func (m Message) GetTimeUnit(f *field.TimeUnitField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for OrderMassCancelRequest.
func (m Message) MaturityTime() (*field.MaturityTimeField, quickfix.MessageRejectError) {
	f := &field.MaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from OrderMassCancelRequest.
func (m Message) GetMaturityTime(f *field.MaturityTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for OrderMassCancelRequest.
func (m Message) SecurityGroup() (*field.SecurityGroupField, quickfix.MessageRejectError) {
	f := &field.SecurityGroupField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from OrderMassCancelRequest.
func (m Message) GetSecurityGroup(f *field.SecurityGroupField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for OrderMassCancelRequest.
func (m Message) MinPriceIncrementAmount() (*field.MinPriceIncrementAmountField, quickfix.MessageRejectError) {
	f := &field.MinPriceIncrementAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from OrderMassCancelRequest.
func (m Message) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmountField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for OrderMassCancelRequest.
func (m Message) UnitOfMeasureQty() (*field.UnitOfMeasureQtyField, quickfix.MessageRejectError) {
	f := &field.UnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from OrderMassCancelRequest.
func (m Message) GetUnitOfMeasureQty(f *field.UnitOfMeasureQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for OrderMassCancelRequest.
func (m Message) SecurityXMLLen() (*field.SecurityXMLLenField, quickfix.MessageRejectError) {
	f := &field.SecurityXMLLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from OrderMassCancelRequest.
func (m Message) GetSecurityXMLLen(f *field.SecurityXMLLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for OrderMassCancelRequest.
func (m Message) SecurityXML() (*field.SecurityXMLField, quickfix.MessageRejectError) {
	f := &field.SecurityXMLField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from OrderMassCancelRequest.
func (m Message) GetSecurityXML(f *field.SecurityXMLField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for OrderMassCancelRequest.
func (m Message) SecurityXMLSchema() (*field.SecurityXMLSchemaField, quickfix.MessageRejectError) {
	f := &field.SecurityXMLSchemaField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from OrderMassCancelRequest.
func (m Message) GetSecurityXMLSchema(f *field.SecurityXMLSchemaField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for OrderMassCancelRequest.
func (m Message) ProductComplex() (*field.ProductComplexField, quickfix.MessageRejectError) {
	f := &field.ProductComplexField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from OrderMassCancelRequest.
func (m Message) GetProductComplex(f *field.ProductComplexField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for OrderMassCancelRequest.
func (m Message) PriceUnitOfMeasure() (*field.PriceUnitOfMeasureField, quickfix.MessageRejectError) {
	f := &field.PriceUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from OrderMassCancelRequest.
func (m Message) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasureField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for OrderMassCancelRequest.
func (m Message) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQtyField, quickfix.MessageRejectError) {
	f := &field.PriceUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from OrderMassCancelRequest.
func (m Message) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for OrderMassCancelRequest.
func (m Message) SettlMethod() (*field.SettlMethodField, quickfix.MessageRejectError) {
	f := &field.SettlMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from OrderMassCancelRequest.
func (m Message) GetSettlMethod(f *field.SettlMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for OrderMassCancelRequest.
func (m Message) ExerciseStyle() (*field.ExerciseStyleField, quickfix.MessageRejectError) {
	f := &field.ExerciseStyleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from OrderMassCancelRequest.
func (m Message) GetExerciseStyle(f *field.ExerciseStyleField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutAmount is a non-required field for OrderMassCancelRequest.
func (m Message) OptPayoutAmount() (*field.OptPayoutAmountField, quickfix.MessageRejectError) {
	f := &field.OptPayoutAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutAmount reads a OptPayoutAmount from OrderMassCancelRequest.
func (m Message) GetOptPayoutAmount(f *field.OptPayoutAmountField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for OrderMassCancelRequest.
func (m Message) PriceQuoteMethod() (*field.PriceQuoteMethodField, quickfix.MessageRejectError) {
	f := &field.PriceQuoteMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from OrderMassCancelRequest.
func (m Message) GetPriceQuoteMethod(f *field.PriceQuoteMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for OrderMassCancelRequest.
func (m Message) ListMethod() (*field.ListMethodField, quickfix.MessageRejectError) {
	f := &field.ListMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from OrderMassCancelRequest.
func (m Message) GetListMethod(f *field.ListMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for OrderMassCancelRequest.
func (m Message) CapPrice() (*field.CapPriceField, quickfix.MessageRejectError) {
	f := &field.CapPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from OrderMassCancelRequest.
func (m Message) GetCapPrice(f *field.CapPriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for OrderMassCancelRequest.
func (m Message) FloorPrice() (*field.FloorPriceField, quickfix.MessageRejectError) {
	f := &field.FloorPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from OrderMassCancelRequest.
func (m Message) GetFloorPrice(f *field.FloorPriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for OrderMassCancelRequest.
func (m Message) PutOrCall() (*field.PutOrCallField, quickfix.MessageRejectError) {
	f := &field.PutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from OrderMassCancelRequest.
func (m Message) GetPutOrCall(f *field.PutOrCallField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for OrderMassCancelRequest.
func (m Message) FlexibleIndicator() (*field.FlexibleIndicatorField, quickfix.MessageRejectError) {
	f := &field.FlexibleIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from OrderMassCancelRequest.
func (m Message) GetFlexibleIndicator(f *field.FlexibleIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for OrderMassCancelRequest.
func (m Message) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicatorField, quickfix.MessageRejectError) {
	f := &field.FlexProductEligibilityIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from OrderMassCancelRequest.
func (m Message) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ValuationMethod is a non-required field for OrderMassCancelRequest.
func (m Message) ValuationMethod() (*field.ValuationMethodField, quickfix.MessageRejectError) {
	f := &field.ValuationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetValuationMethod reads a ValuationMethod from OrderMassCancelRequest.
func (m Message) GetValuationMethod(f *field.ValuationMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplierUnit is a non-required field for OrderMassCancelRequest.
func (m Message) ContractMultiplierUnit() (*field.ContractMultiplierUnitField, quickfix.MessageRejectError) {
	f := &field.ContractMultiplierUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplierUnit reads a ContractMultiplierUnit from OrderMassCancelRequest.
func (m Message) GetContractMultiplierUnit(f *field.ContractMultiplierUnitField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//FlowScheduleType is a non-required field for OrderMassCancelRequest.
func (m Message) FlowScheduleType() (*field.FlowScheduleTypeField, quickfix.MessageRejectError) {
	f := &field.FlowScheduleTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlowScheduleType reads a FlowScheduleType from OrderMassCancelRequest.
func (m Message) GetFlowScheduleType(f *field.FlowScheduleTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RestructuringType is a non-required field for OrderMassCancelRequest.
func (m Message) RestructuringType() (*field.RestructuringTypeField, quickfix.MessageRejectError) {
	f := &field.RestructuringTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRestructuringType reads a RestructuringType from OrderMassCancelRequest.
func (m Message) GetRestructuringType(f *field.RestructuringTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Seniority is a non-required field for OrderMassCancelRequest.
func (m Message) Seniority() (*field.SeniorityField, quickfix.MessageRejectError) {
	f := &field.SeniorityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSeniority reads a Seniority from OrderMassCancelRequest.
func (m Message) GetSeniority(f *field.SeniorityField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NotionalPercentageOutstanding is a non-required field for OrderMassCancelRequest.
func (m Message) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstandingField, quickfix.MessageRejectError) {
	f := &field.NotionalPercentageOutstandingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNotionalPercentageOutstanding reads a NotionalPercentageOutstanding from OrderMassCancelRequest.
func (m Message) GetNotionalPercentageOutstanding(f *field.NotionalPercentageOutstandingField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OriginalNotionalPercentageOutstanding is a non-required field for OrderMassCancelRequest.
func (m Message) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstandingField, quickfix.MessageRejectError) {
	f := &field.OriginalNotionalPercentageOutstandingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOriginalNotionalPercentageOutstanding reads a OriginalNotionalPercentageOutstanding from OrderMassCancelRequest.
func (m Message) GetOriginalNotionalPercentageOutstanding(f *field.OriginalNotionalPercentageOutstandingField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AttachmentPoint is a non-required field for OrderMassCancelRequest.
func (m Message) AttachmentPoint() (*field.AttachmentPointField, quickfix.MessageRejectError) {
	f := &field.AttachmentPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAttachmentPoint reads a AttachmentPoint from OrderMassCancelRequest.
func (m Message) GetAttachmentPoint(f *field.AttachmentPointField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DetachmentPoint is a non-required field for OrderMassCancelRequest.
func (m Message) DetachmentPoint() (*field.DetachmentPointField, quickfix.MessageRejectError) {
	f := &field.DetachmentPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDetachmentPoint reads a DetachmentPoint from OrderMassCancelRequest.
func (m Message) GetDetachmentPoint(f *field.DetachmentPointField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceDeterminationMethod is a non-required field for OrderMassCancelRequest.
func (m Message) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethodField, quickfix.MessageRejectError) {
	f := &field.StrikePriceDeterminationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceDeterminationMethod reads a StrikePriceDeterminationMethod from OrderMassCancelRequest.
func (m Message) GetStrikePriceDeterminationMethod(f *field.StrikePriceDeterminationMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryMethod is a non-required field for OrderMassCancelRequest.
func (m Message) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethodField, quickfix.MessageRejectError) {
	f := &field.StrikePriceBoundaryMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryMethod reads a StrikePriceBoundaryMethod from OrderMassCancelRequest.
func (m Message) GetStrikePriceBoundaryMethod(f *field.StrikePriceBoundaryMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryPrecision is a non-required field for OrderMassCancelRequest.
func (m Message) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecisionField, quickfix.MessageRejectError) {
	f := &field.StrikePriceBoundaryPrecisionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryPrecision reads a StrikePriceBoundaryPrecision from OrderMassCancelRequest.
func (m Message) GetStrikePriceBoundaryPrecision(f *field.StrikePriceBoundaryPrecisionField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceDeterminationMethod is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethodField, quickfix.MessageRejectError) {
	f := &field.UnderlyingPriceDeterminationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceDeterminationMethod reads a UnderlyingPriceDeterminationMethod from OrderMassCancelRequest.
func (m Message) GetUnderlyingPriceDeterminationMethod(f *field.UnderlyingPriceDeterminationMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutType is a non-required field for OrderMassCancelRequest.
func (m Message) OptPayoutType() (*field.OptPayoutTypeField, quickfix.MessageRejectError) {
	f := &field.OptPayoutTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutType reads a OptPayoutType from OrderMassCancelRequest.
func (m Message) GetOptPayoutType(f *field.OptPayoutTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoComplexEvents is a non-required field for OrderMassCancelRequest.
func (m Message) NoComplexEvents() (*field.NoComplexEventsField, quickfix.MessageRejectError) {
	f := &field.NoComplexEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoComplexEvents reads a NoComplexEvents from OrderMassCancelRequest.
func (m Message) GetNoComplexEvents(f *field.NoComplexEventsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSymbol is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingSymbol() (*field.UnderlyingSymbolField, quickfix.MessageRejectError) {
	f := &field.UnderlyingSymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSymbol reads a UnderlyingSymbol from OrderMassCancelRequest.
func (m Message) GetUnderlyingSymbol(f *field.UnderlyingSymbolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSymbolSfx is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingSymbolSfx() (*field.UnderlyingSymbolSfxField, quickfix.MessageRejectError) {
	f := &field.UnderlyingSymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSymbolSfx reads a UnderlyingSymbolSfx from OrderMassCancelRequest.
func (m Message) GetUnderlyingSymbolSfx(f *field.UnderlyingSymbolSfxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityID is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingSecurityID() (*field.UnderlyingSecurityIDField, quickfix.MessageRejectError) {
	f := &field.UnderlyingSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityID reads a UnderlyingSecurityID from OrderMassCancelRequest.
func (m Message) GetUnderlyingSecurityID(f *field.UnderlyingSecurityIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityIDSource is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingSecurityIDSource() (*field.UnderlyingSecurityIDSourceField, quickfix.MessageRejectError) {
	f := &field.UnderlyingSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityIDSource reads a UnderlyingSecurityIDSource from OrderMassCancelRequest.
func (m Message) GetUnderlyingSecurityIDSource(f *field.UnderlyingSecurityIDSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyingSecurityAltID is a non-required field for OrderMassCancelRequest.
func (m Message) NoUnderlyingSecurityAltID() (*field.NoUnderlyingSecurityAltIDField, quickfix.MessageRejectError) {
	f := &field.NoUnderlyingSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyingSecurityAltID reads a NoUnderlyingSecurityAltID from OrderMassCancelRequest.
func (m Message) GetNoUnderlyingSecurityAltID(f *field.NoUnderlyingSecurityAltIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingProduct is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingProduct() (*field.UnderlyingProductField, quickfix.MessageRejectError) {
	f := &field.UnderlyingProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingProduct reads a UnderlyingProduct from OrderMassCancelRequest.
func (m Message) GetUnderlyingProduct(f *field.UnderlyingProductField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCFICode is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingCFICode() (*field.UnderlyingCFICodeField, quickfix.MessageRejectError) {
	f := &field.UnderlyingCFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCFICode reads a UnderlyingCFICode from OrderMassCancelRequest.
func (m Message) GetUnderlyingCFICode(f *field.UnderlyingCFICodeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityType is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingSecurityType() (*field.UnderlyingSecurityTypeField, quickfix.MessageRejectError) {
	f := &field.UnderlyingSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityType reads a UnderlyingSecurityType from OrderMassCancelRequest.
func (m Message) GetUnderlyingSecurityType(f *field.UnderlyingSecurityTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecuritySubType is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingSecuritySubType() (*field.UnderlyingSecuritySubTypeField, quickfix.MessageRejectError) {
	f := &field.UnderlyingSecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecuritySubType reads a UnderlyingSecuritySubType from OrderMassCancelRequest.
func (m Message) GetUnderlyingSecuritySubType(f *field.UnderlyingSecuritySubTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityMonthYear is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingMaturityMonthYear() (*field.UnderlyingMaturityMonthYearField, quickfix.MessageRejectError) {
	f := &field.UnderlyingMaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityMonthYear reads a UnderlyingMaturityMonthYear from OrderMassCancelRequest.
func (m Message) GetUnderlyingMaturityMonthYear(f *field.UnderlyingMaturityMonthYearField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityDate is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingMaturityDate() (*field.UnderlyingMaturityDateField, quickfix.MessageRejectError) {
	f := &field.UnderlyingMaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityDate reads a UnderlyingMaturityDate from OrderMassCancelRequest.
func (m Message) GetUnderlyingMaturityDate(f *field.UnderlyingMaturityDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCouponPaymentDate is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingCouponPaymentDate() (*field.UnderlyingCouponPaymentDateField, quickfix.MessageRejectError) {
	f := &field.UnderlyingCouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCouponPaymentDate reads a UnderlyingCouponPaymentDate from OrderMassCancelRequest.
func (m Message) GetUnderlyingCouponPaymentDate(f *field.UnderlyingCouponPaymentDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingIssueDate is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingIssueDate() (*field.UnderlyingIssueDateField, quickfix.MessageRejectError) {
	f := &field.UnderlyingIssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingIssueDate reads a UnderlyingIssueDate from OrderMassCancelRequest.
func (m Message) GetUnderlyingIssueDate(f *field.UnderlyingIssueDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepoCollateralSecurityType is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingRepoCollateralSecurityType() (*field.UnderlyingRepoCollateralSecurityTypeField, quickfix.MessageRejectError) {
	f := &field.UnderlyingRepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepoCollateralSecurityType reads a UnderlyingRepoCollateralSecurityType from OrderMassCancelRequest.
func (m Message) GetUnderlyingRepoCollateralSecurityType(f *field.UnderlyingRepoCollateralSecurityTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepurchaseTerm is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingRepurchaseTerm() (*field.UnderlyingRepurchaseTermField, quickfix.MessageRejectError) {
	f := &field.UnderlyingRepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepurchaseTerm reads a UnderlyingRepurchaseTerm from OrderMassCancelRequest.
func (m Message) GetUnderlyingRepurchaseTerm(f *field.UnderlyingRepurchaseTermField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepurchaseRate is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingRepurchaseRate() (*field.UnderlyingRepurchaseRateField, quickfix.MessageRejectError) {
	f := &field.UnderlyingRepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepurchaseRate reads a UnderlyingRepurchaseRate from OrderMassCancelRequest.
func (m Message) GetUnderlyingRepurchaseRate(f *field.UnderlyingRepurchaseRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFactor is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingFactor() (*field.UnderlyingFactorField, quickfix.MessageRejectError) {
	f := &field.UnderlyingFactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFactor reads a UnderlyingFactor from OrderMassCancelRequest.
func (m Message) GetUnderlyingFactor(f *field.UnderlyingFactorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCreditRating is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingCreditRating() (*field.UnderlyingCreditRatingField, quickfix.MessageRejectError) {
	f := &field.UnderlyingCreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCreditRating reads a UnderlyingCreditRating from OrderMassCancelRequest.
func (m Message) GetUnderlyingCreditRating(f *field.UnderlyingCreditRatingField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingInstrRegistry is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingInstrRegistry() (*field.UnderlyingInstrRegistryField, quickfix.MessageRejectError) {
	f := &field.UnderlyingInstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingInstrRegistry reads a UnderlyingInstrRegistry from OrderMassCancelRequest.
func (m Message) GetUnderlyingInstrRegistry(f *field.UnderlyingInstrRegistryField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCountryOfIssue is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingCountryOfIssue() (*field.UnderlyingCountryOfIssueField, quickfix.MessageRejectError) {
	f := &field.UnderlyingCountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCountryOfIssue reads a UnderlyingCountryOfIssue from OrderMassCancelRequest.
func (m Message) GetUnderlyingCountryOfIssue(f *field.UnderlyingCountryOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStateOrProvinceOfIssue is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingStateOrProvinceOfIssue() (*field.UnderlyingStateOrProvinceOfIssueField, quickfix.MessageRejectError) {
	f := &field.UnderlyingStateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStateOrProvinceOfIssue reads a UnderlyingStateOrProvinceOfIssue from OrderMassCancelRequest.
func (m Message) GetUnderlyingStateOrProvinceOfIssue(f *field.UnderlyingStateOrProvinceOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingLocaleOfIssue is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingLocaleOfIssue() (*field.UnderlyingLocaleOfIssueField, quickfix.MessageRejectError) {
	f := &field.UnderlyingLocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingLocaleOfIssue reads a UnderlyingLocaleOfIssue from OrderMassCancelRequest.
func (m Message) GetUnderlyingLocaleOfIssue(f *field.UnderlyingLocaleOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRedemptionDate is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingRedemptionDate() (*field.UnderlyingRedemptionDateField, quickfix.MessageRejectError) {
	f := &field.UnderlyingRedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRedemptionDate reads a UnderlyingRedemptionDate from OrderMassCancelRequest.
func (m Message) GetUnderlyingRedemptionDate(f *field.UnderlyingRedemptionDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStrikePrice is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingStrikePrice() (*field.UnderlyingStrikePriceField, quickfix.MessageRejectError) {
	f := &field.UnderlyingStrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStrikePrice reads a UnderlyingStrikePrice from OrderMassCancelRequest.
func (m Message) GetUnderlyingStrikePrice(f *field.UnderlyingStrikePriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStrikeCurrency is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingStrikeCurrency() (*field.UnderlyingStrikeCurrencyField, quickfix.MessageRejectError) {
	f := &field.UnderlyingStrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStrikeCurrency reads a UnderlyingStrikeCurrency from OrderMassCancelRequest.
func (m Message) GetUnderlyingStrikeCurrency(f *field.UnderlyingStrikeCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingOptAttribute is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingOptAttribute() (*field.UnderlyingOptAttributeField, quickfix.MessageRejectError) {
	f := &field.UnderlyingOptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingOptAttribute reads a UnderlyingOptAttribute from OrderMassCancelRequest.
func (m Message) GetUnderlyingOptAttribute(f *field.UnderlyingOptAttributeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingContractMultiplier is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingContractMultiplier() (*field.UnderlyingContractMultiplierField, quickfix.MessageRejectError) {
	f := &field.UnderlyingContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingContractMultiplier reads a UnderlyingContractMultiplier from OrderMassCancelRequest.
func (m Message) GetUnderlyingContractMultiplier(f *field.UnderlyingContractMultiplierField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCouponRate is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingCouponRate() (*field.UnderlyingCouponRateField, quickfix.MessageRejectError) {
	f := &field.UnderlyingCouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCouponRate reads a UnderlyingCouponRate from OrderMassCancelRequest.
func (m Message) GetUnderlyingCouponRate(f *field.UnderlyingCouponRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityExchange is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingSecurityExchange() (*field.UnderlyingSecurityExchangeField, quickfix.MessageRejectError) {
	f := &field.UnderlyingSecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityExchange reads a UnderlyingSecurityExchange from OrderMassCancelRequest.
func (m Message) GetUnderlyingSecurityExchange(f *field.UnderlyingSecurityExchangeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingIssuer is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingIssuer() (*field.UnderlyingIssuerField, quickfix.MessageRejectError) {
	f := &field.UnderlyingIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingIssuer reads a UnderlyingIssuer from OrderMassCancelRequest.
func (m Message) GetUnderlyingIssuer(f *field.UnderlyingIssuerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingIssuerLen is a non-required field for OrderMassCancelRequest.
func (m Message) EncodedUnderlyingIssuerLen() (*field.EncodedUnderlyingIssuerLenField, quickfix.MessageRejectError) {
	f := &field.EncodedUnderlyingIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingIssuerLen reads a EncodedUnderlyingIssuerLen from OrderMassCancelRequest.
func (m Message) GetEncodedUnderlyingIssuerLen(f *field.EncodedUnderlyingIssuerLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingIssuer is a non-required field for OrderMassCancelRequest.
func (m Message) EncodedUnderlyingIssuer() (*field.EncodedUnderlyingIssuerField, quickfix.MessageRejectError) {
	f := &field.EncodedUnderlyingIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingIssuer reads a EncodedUnderlyingIssuer from OrderMassCancelRequest.
func (m Message) GetEncodedUnderlyingIssuer(f *field.EncodedUnderlyingIssuerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityDesc is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingSecurityDesc() (*field.UnderlyingSecurityDescField, quickfix.MessageRejectError) {
	f := &field.UnderlyingSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityDesc reads a UnderlyingSecurityDesc from OrderMassCancelRequest.
func (m Message) GetUnderlyingSecurityDesc(f *field.UnderlyingSecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingSecurityDescLen is a non-required field for OrderMassCancelRequest.
func (m Message) EncodedUnderlyingSecurityDescLen() (*field.EncodedUnderlyingSecurityDescLenField, quickfix.MessageRejectError) {
	f := &field.EncodedUnderlyingSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingSecurityDescLen reads a EncodedUnderlyingSecurityDescLen from OrderMassCancelRequest.
func (m Message) GetEncodedUnderlyingSecurityDescLen(f *field.EncodedUnderlyingSecurityDescLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingSecurityDesc is a non-required field for OrderMassCancelRequest.
func (m Message) EncodedUnderlyingSecurityDesc() (*field.EncodedUnderlyingSecurityDescField, quickfix.MessageRejectError) {
	f := &field.EncodedUnderlyingSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingSecurityDesc reads a EncodedUnderlyingSecurityDesc from OrderMassCancelRequest.
func (m Message) GetEncodedUnderlyingSecurityDesc(f *field.EncodedUnderlyingSecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCPProgram is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingCPProgram() (*field.UnderlyingCPProgramField, quickfix.MessageRejectError) {
	f := &field.UnderlyingCPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCPProgram reads a UnderlyingCPProgram from OrderMassCancelRequest.
func (m Message) GetUnderlyingCPProgram(f *field.UnderlyingCPProgramField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCPRegType is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingCPRegType() (*field.UnderlyingCPRegTypeField, quickfix.MessageRejectError) {
	f := &field.UnderlyingCPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCPRegType reads a UnderlyingCPRegType from OrderMassCancelRequest.
func (m Message) GetUnderlyingCPRegType(f *field.UnderlyingCPRegTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCurrency is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingCurrency() (*field.UnderlyingCurrencyField, quickfix.MessageRejectError) {
	f := &field.UnderlyingCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCurrency reads a UnderlyingCurrency from OrderMassCancelRequest.
func (m Message) GetUnderlyingCurrency(f *field.UnderlyingCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingQty is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingQty() (*field.UnderlyingQtyField, quickfix.MessageRejectError) {
	f := &field.UnderlyingQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingQty reads a UnderlyingQty from OrderMassCancelRequest.
func (m Message) GetUnderlyingQty(f *field.UnderlyingQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPx is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingPx() (*field.UnderlyingPxField, quickfix.MessageRejectError) {
	f := &field.UnderlyingPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPx reads a UnderlyingPx from OrderMassCancelRequest.
func (m Message) GetUnderlyingPx(f *field.UnderlyingPxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingDirtyPrice is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingDirtyPrice() (*field.UnderlyingDirtyPriceField, quickfix.MessageRejectError) {
	f := &field.UnderlyingDirtyPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingDirtyPrice reads a UnderlyingDirtyPrice from OrderMassCancelRequest.
func (m Message) GetUnderlyingDirtyPrice(f *field.UnderlyingDirtyPriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingEndPrice is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingEndPrice() (*field.UnderlyingEndPriceField, quickfix.MessageRejectError) {
	f := &field.UnderlyingEndPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingEndPrice reads a UnderlyingEndPrice from OrderMassCancelRequest.
func (m Message) GetUnderlyingEndPrice(f *field.UnderlyingEndPriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStartValue is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingStartValue() (*field.UnderlyingStartValueField, quickfix.MessageRejectError) {
	f := &field.UnderlyingStartValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStartValue reads a UnderlyingStartValue from OrderMassCancelRequest.
func (m Message) GetUnderlyingStartValue(f *field.UnderlyingStartValueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCurrentValue is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingCurrentValue() (*field.UnderlyingCurrentValueField, quickfix.MessageRejectError) {
	f := &field.UnderlyingCurrentValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCurrentValue reads a UnderlyingCurrentValue from OrderMassCancelRequest.
func (m Message) GetUnderlyingCurrentValue(f *field.UnderlyingCurrentValueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingEndValue is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingEndValue() (*field.UnderlyingEndValueField, quickfix.MessageRejectError) {
	f := &field.UnderlyingEndValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingEndValue reads a UnderlyingEndValue from OrderMassCancelRequest.
func (m Message) GetUnderlyingEndValue(f *field.UnderlyingEndValueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyingStips is a non-required field for OrderMassCancelRequest.
func (m Message) NoUnderlyingStips() (*field.NoUnderlyingStipsField, quickfix.MessageRejectError) {
	f := &field.NoUnderlyingStipsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyingStips reads a NoUnderlyingStips from OrderMassCancelRequest.
func (m Message) GetNoUnderlyingStips(f *field.NoUnderlyingStipsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingAllocationPercent is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingAllocationPercent() (*field.UnderlyingAllocationPercentField, quickfix.MessageRejectError) {
	f := &field.UnderlyingAllocationPercentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingAllocationPercent reads a UnderlyingAllocationPercent from OrderMassCancelRequest.
func (m Message) GetUnderlyingAllocationPercent(f *field.UnderlyingAllocationPercentField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSettlementType is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingSettlementType() (*field.UnderlyingSettlementTypeField, quickfix.MessageRejectError) {
	f := &field.UnderlyingSettlementTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSettlementType reads a UnderlyingSettlementType from OrderMassCancelRequest.
func (m Message) GetUnderlyingSettlementType(f *field.UnderlyingSettlementTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCashAmount is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingCashAmount() (*field.UnderlyingCashAmountField, quickfix.MessageRejectError) {
	f := &field.UnderlyingCashAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCashAmount reads a UnderlyingCashAmount from OrderMassCancelRequest.
func (m Message) GetUnderlyingCashAmount(f *field.UnderlyingCashAmountField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCashType is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingCashType() (*field.UnderlyingCashTypeField, quickfix.MessageRejectError) {
	f := &field.UnderlyingCashTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCashType reads a UnderlyingCashType from OrderMassCancelRequest.
func (m Message) GetUnderlyingCashType(f *field.UnderlyingCashTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingUnitOfMeasure is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingUnitOfMeasure() (*field.UnderlyingUnitOfMeasureField, quickfix.MessageRejectError) {
	f := &field.UnderlyingUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingUnitOfMeasure reads a UnderlyingUnitOfMeasure from OrderMassCancelRequest.
func (m Message) GetUnderlyingUnitOfMeasure(f *field.UnderlyingUnitOfMeasureField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingTimeUnit is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingTimeUnit() (*field.UnderlyingTimeUnitField, quickfix.MessageRejectError) {
	f := &field.UnderlyingTimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingTimeUnit reads a UnderlyingTimeUnit from OrderMassCancelRequest.
func (m Message) GetUnderlyingTimeUnit(f *field.UnderlyingTimeUnitField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCapValue is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingCapValue() (*field.UnderlyingCapValueField, quickfix.MessageRejectError) {
	f := &field.UnderlyingCapValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCapValue reads a UnderlyingCapValue from OrderMassCancelRequest.
func (m Message) GetUnderlyingCapValue(f *field.UnderlyingCapValueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoUndlyInstrumentParties is a non-required field for OrderMassCancelRequest.
func (m Message) NoUndlyInstrumentParties() (*field.NoUndlyInstrumentPartiesField, quickfix.MessageRejectError) {
	f := &field.NoUndlyInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUndlyInstrumentParties reads a NoUndlyInstrumentParties from OrderMassCancelRequest.
func (m Message) GetNoUndlyInstrumentParties(f *field.NoUndlyInstrumentPartiesField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSettlMethod is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingSettlMethod() (*field.UnderlyingSettlMethodField, quickfix.MessageRejectError) {
	f := &field.UnderlyingSettlMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSettlMethod reads a UnderlyingSettlMethod from OrderMassCancelRequest.
func (m Message) GetUnderlyingSettlMethod(f *field.UnderlyingSettlMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingAdjustedQuantity is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingAdjustedQuantity() (*field.UnderlyingAdjustedQuantityField, quickfix.MessageRejectError) {
	f := &field.UnderlyingAdjustedQuantityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingAdjustedQuantity reads a UnderlyingAdjustedQuantity from OrderMassCancelRequest.
func (m Message) GetUnderlyingAdjustedQuantity(f *field.UnderlyingAdjustedQuantityField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFXRate is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingFXRate() (*field.UnderlyingFXRateField, quickfix.MessageRejectError) {
	f := &field.UnderlyingFXRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFXRate reads a UnderlyingFXRate from OrderMassCancelRequest.
func (m Message) GetUnderlyingFXRate(f *field.UnderlyingFXRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFXRateCalc is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingFXRateCalc() (*field.UnderlyingFXRateCalcField, quickfix.MessageRejectError) {
	f := &field.UnderlyingFXRateCalcField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFXRateCalc reads a UnderlyingFXRateCalc from OrderMassCancelRequest.
func (m Message) GetUnderlyingFXRateCalc(f *field.UnderlyingFXRateCalcField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityTime is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingMaturityTime() (*field.UnderlyingMaturityTimeField, quickfix.MessageRejectError) {
	f := &field.UnderlyingMaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityTime reads a UnderlyingMaturityTime from OrderMassCancelRequest.
func (m Message) GetUnderlyingMaturityTime(f *field.UnderlyingMaturityTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPutOrCall is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingPutOrCall() (*field.UnderlyingPutOrCallField, quickfix.MessageRejectError) {
	f := &field.UnderlyingPutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPutOrCall reads a UnderlyingPutOrCall from OrderMassCancelRequest.
func (m Message) GetUnderlyingPutOrCall(f *field.UnderlyingPutOrCallField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingExerciseStyle is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingExerciseStyle() (*field.UnderlyingExerciseStyleField, quickfix.MessageRejectError) {
	f := &field.UnderlyingExerciseStyleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingExerciseStyle reads a UnderlyingExerciseStyle from OrderMassCancelRequest.
func (m Message) GetUnderlyingExerciseStyle(f *field.UnderlyingExerciseStyleField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingUnitOfMeasureQty is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingUnitOfMeasureQty() (*field.UnderlyingUnitOfMeasureQtyField, quickfix.MessageRejectError) {
	f := &field.UnderlyingUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingUnitOfMeasureQty reads a UnderlyingUnitOfMeasureQty from OrderMassCancelRequest.
func (m Message) GetUnderlyingUnitOfMeasureQty(f *field.UnderlyingUnitOfMeasureQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceUnitOfMeasure is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingPriceUnitOfMeasure() (*field.UnderlyingPriceUnitOfMeasureField, quickfix.MessageRejectError) {
	f := &field.UnderlyingPriceUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceUnitOfMeasure reads a UnderlyingPriceUnitOfMeasure from OrderMassCancelRequest.
func (m Message) GetUnderlyingPriceUnitOfMeasure(f *field.UnderlyingPriceUnitOfMeasureField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceUnitOfMeasureQty is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingPriceUnitOfMeasureQty() (*field.UnderlyingPriceUnitOfMeasureQtyField, quickfix.MessageRejectError) {
	f := &field.UnderlyingPriceUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceUnitOfMeasureQty reads a UnderlyingPriceUnitOfMeasureQty from OrderMassCancelRequest.
func (m Message) GetUnderlyingPriceUnitOfMeasureQty(f *field.UnderlyingPriceUnitOfMeasureQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingContractMultiplierUnit is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingContractMultiplierUnit() (*field.UnderlyingContractMultiplierUnitField, quickfix.MessageRejectError) {
	f := &field.UnderlyingContractMultiplierUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingContractMultiplierUnit reads a UnderlyingContractMultiplierUnit from OrderMassCancelRequest.
func (m Message) GetUnderlyingContractMultiplierUnit(f *field.UnderlyingContractMultiplierUnitField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFlowScheduleType is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingFlowScheduleType() (*field.UnderlyingFlowScheduleTypeField, quickfix.MessageRejectError) {
	f := &field.UnderlyingFlowScheduleTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFlowScheduleType reads a UnderlyingFlowScheduleType from OrderMassCancelRequest.
func (m Message) GetUnderlyingFlowScheduleType(f *field.UnderlyingFlowScheduleTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRestructuringType is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingRestructuringType() (*field.UnderlyingRestructuringTypeField, quickfix.MessageRejectError) {
	f := &field.UnderlyingRestructuringTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRestructuringType reads a UnderlyingRestructuringType from OrderMassCancelRequest.
func (m Message) GetUnderlyingRestructuringType(f *field.UnderlyingRestructuringTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSeniority is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingSeniority() (*field.UnderlyingSeniorityField, quickfix.MessageRejectError) {
	f := &field.UnderlyingSeniorityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSeniority reads a UnderlyingSeniority from OrderMassCancelRequest.
func (m Message) GetUnderlyingSeniority(f *field.UnderlyingSeniorityField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingNotionalPercentageOutstanding is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingNotionalPercentageOutstanding() (*field.UnderlyingNotionalPercentageOutstandingField, quickfix.MessageRejectError) {
	f := &field.UnderlyingNotionalPercentageOutstandingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingNotionalPercentageOutstanding reads a UnderlyingNotionalPercentageOutstanding from OrderMassCancelRequest.
func (m Message) GetUnderlyingNotionalPercentageOutstanding(f *field.UnderlyingNotionalPercentageOutstandingField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingOriginalNotionalPercentageOutstanding is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingOriginalNotionalPercentageOutstanding() (*field.UnderlyingOriginalNotionalPercentageOutstandingField, quickfix.MessageRejectError) {
	f := &field.UnderlyingOriginalNotionalPercentageOutstandingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingOriginalNotionalPercentageOutstanding reads a UnderlyingOriginalNotionalPercentageOutstanding from OrderMassCancelRequest.
func (m Message) GetUnderlyingOriginalNotionalPercentageOutstanding(f *field.UnderlyingOriginalNotionalPercentageOutstandingField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingAttachmentPoint is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingAttachmentPoint() (*field.UnderlyingAttachmentPointField, quickfix.MessageRejectError) {
	f := &field.UnderlyingAttachmentPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingAttachmentPoint reads a UnderlyingAttachmentPoint from OrderMassCancelRequest.
func (m Message) GetUnderlyingAttachmentPoint(f *field.UnderlyingAttachmentPointField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingDetachmentPoint is a non-required field for OrderMassCancelRequest.
func (m Message) UnderlyingDetachmentPoint() (*field.UnderlyingDetachmentPointField, quickfix.MessageRejectError) {
	f := &field.UnderlyingDetachmentPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingDetachmentPoint reads a UnderlyingDetachmentPoint from OrderMassCancelRequest.
func (m Message) GetUnderlyingDetachmentPoint(f *field.UnderlyingDetachmentPointField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a non-required field for OrderMassCancelRequest.
func (m Message) Side() (*field.SideField, quickfix.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from OrderMassCancelRequest.
func (m Message) GetSide(f *field.SideField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for OrderMassCancelRequest.
func (m Message) TransactTime() (*field.TransactTimeField, quickfix.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from OrderMassCancelRequest.
func (m Message) GetTransactTime(f *field.TransactTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for OrderMassCancelRequest.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from OrderMassCancelRequest.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for OrderMassCancelRequest.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, quickfix.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from OrderMassCancelRequest.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for OrderMassCancelRequest.
func (m Message) EncodedText() (*field.EncodedTextField, quickfix.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from OrderMassCancelRequest.
func (m Message) GetEncodedText(f *field.EncodedTextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for OrderMassCancelRequest.
func (m Message) NoPartyIDs() (*field.NoPartyIDsField, quickfix.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from OrderMassCancelRequest.
func (m Message) GetNoPartyIDs(f *field.NoPartyIDsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a non-required field for OrderMassCancelRequest.
func (m Message) MarketID() (*field.MarketIDField, quickfix.MessageRejectError) {
	f := &field.MarketIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from OrderMassCancelRequest.
func (m Message) GetMarketID(f *field.MarketIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for OrderMassCancelRequest.
func (m Message) MarketSegmentID() (*field.MarketSegmentIDField, quickfix.MessageRejectError) {
	f := &field.MarketSegmentIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from OrderMassCancelRequest.
func (m Message) GetMarketSegmentID(f *field.MarketSegmentIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoTargetPartyIDs is a non-required field for OrderMassCancelRequest.
func (m Message) NoTargetPartyIDs() (*field.NoTargetPartyIDsField, quickfix.MessageRejectError) {
	f := &field.NoTargetPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTargetPartyIDs reads a NoTargetPartyIDs from OrderMassCancelRequest.
func (m Message) GetNoTargetPartyIDs(f *field.NoTargetPartyIDsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for OrderMassCancelRequest.
func New(
	clordid *field.ClOrdIDField,
	masscancelrequesttype *field.MassCancelRequestTypeField,
	transacttime *field.TransactTimeField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("q"))
	builder.Body.Set(clordid)
	builder.Body.Set(masscancelrequesttype)
	builder.Body.Set(transacttime)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "q", r
}
